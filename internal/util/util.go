package util

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strconv"
	"strings"
)

func Map[T, E any](arr []T, f func(T) E) []E {
	result := make([]E, 0, cap(arr))

	for _, elem := range arr {
		result = append(result, f(elem))
	}

	return result
}

func Filter[T any](arr []T, f func(T) bool) []T {
	result := make([]T, 0)

	for _, elem := range arr {
		if f(elem) {
			result = append(result, elem)
		}
	}

	return result
}

func Reorganize(directoryPath string) {
	entries, err := os.ReadDir(directoryPath)

	if err != nil {
		log.Fatal(err)
	}

	onlyJsons := Filter(entries, func(entry fs.DirEntry) bool {
		return strings.HasSuffix(entry.Name(), ".json")
	})

	names := Map(onlyJsons, func(entry fs.DirEntry) string {
		return strings.TrimSuffix(entry.Name(), ".json")
	})

	expectedId := 1

	for _, name := range names {
		id, err := strconv.Atoi(name)

		if err != nil {
			continue
		}

		if id != expectedId {
			oldPath := fmt.Sprintf("%v%v.json", directoryPath, id)
			newPath := fmt.Sprintf("%v%v.json", directoryPath, expectedId)
			err := os.Rename(oldPath, newPath)

			if err != nil {
				log.Fatal(err)
			}
		}

		expectedId++
	}
}

func ValidateEntry(entry fs.DirEntry) bool {
	if !strings.HasSuffix(entry.Name(), ".json") {
		return false
	}

	var id int

	_, err := fmt.Sscanf(entry.Name(), "%d.json", &id)

	return err == nil
}
