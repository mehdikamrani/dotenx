package pipelineStore

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log"

	"github.com/utopiops/automated-ops/ao-api/db"
	"github.com/utopiops/automated-ops/ao-api/models"
)

func (ps *pipelineStore) GetTaskByExecution(context context.Context, executionId int, taskId int) (task models.TaskDetails, err error) {
	switch ps.db.Driver {
	case db.Postgres:
		conn := ps.db.Connection
		var body interface{}
		err = conn.QueryRow(getTaskByExecution, executionId, taskId).Scan(&task.Id, &task.Name, &task.Type, &task.Integration, &body, &task.Timeout, &task.AccountId)
		if err != nil {
			log.Println(err.Error())
			if err == sql.ErrNoRows {
				err = errors.New("task Details not found")
			}
			return
		}
		var taskBody models.TaskBodyMap
		json.Unmarshal(body.([]byte), &taskBody)
		task.Body = taskBody
	}
	return

}

var getTaskByExecution = `
select t.id, t.name, t.task_type, t.integration, t.body, t.timeout, pv.account_id
from executions e
join pipelines pv on e.pipeline_id = pv.id
join tasks t on t.pipeline_id = pv.id
where e.id = $1 and t.id = $2
LIMIT 1;
`
