package task

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Task struct {
	Short string     `json:"short"`
	Long  string     `json:"long"`
	Until *time.Time `json:"until,omitempty"`
}

func (task *Task) String() (result string) {
	result += fmt.Sprintf("SHORT DESCRIPTION: %s\n\n", task.Short)

	if task.Long != "" {
		result += fmt.Sprintf("LONG DESCRIPTION:\n%s\n\n", task.Long)
	} else {
		result += "LONG DESCRIPTION: No long description.\n\n"
	}

	if task.Until != nil {
		result += fmt.Sprintf("UNTIL: %s", task.Until)
	} else {
		result += "UNTIL: No deadline"
	}

	return
}

func (task *Task) Read(path string) {
	bytes, err := os.ReadFile(path)

	if err != nil {
		log.Fatal("Could not read file located at " + path)
	}

	err = json.Unmarshal(bytes, task)

	if err != nil {
		log.Fatal("Could not deserialize json into task.")
	}
}

func (task *Task) Parse(flags *pflag.FlagSet) {
	short, _ := flags.GetString("short")
	long, _ := flags.GetString("long")
	task.Short, task.Long = short, long

	untilString, _ := flags.GetString("until")
	untilTime, err := time.Parse("2006-01-02 15:04", untilString)

	if err == nil {
		task.Until = &untilTime
	}
}

func (task *Task) GetTaskPath() string {
	directoryPath := viper.Get("tasks.path").(string)

	return fmt.Sprintf("%v%v.json", directoryPath, getLastAvailableId(directoryPath))
}

func getLastAvailableId(directoryPath string) int {
	for i := 1; ; i++ {
		if _, err := os.Stat(directoryPath + strconv.Itoa(i) + ".json"); err != nil {
			return i
		}
	}
}
