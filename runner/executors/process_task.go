package executors

import (
	"fmt"
	"strings"

	"github.com/utopiops/automated-ops/runner/models"
)

func ProcessTask(task *models.TaskDetails) (processedTask *models.Task) {
	processedTask = &models.Task{}
	processedTask.Details = *task
	processedTask.IsPredifined = true
	if task.Type == "runImage" {
		processedTask.IsPredifined = false
		processedTask.Details.Image = task.Body["image"].(string)
		processedTask.Script = strings.Split(task.Body["script"].(string), " ")
	} else {
		envs := make([]string, 0)
		for _, field := range task.MetaData.Fields {
			if value, ok := task.Body[field.Key]; ok {
				var envVar string
				if field.Type == "text" {
					envVar = field.Key + "=" + value.(string)
				} else {
					envVar = field.Key + "=" + fmt.Sprintf("%v", value)
				}
				envs = append(envs, envVar)
			}
		}
		processedTask.EnvironmentVariables = envs
	}
	return
}
