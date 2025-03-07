package pipelineStore

import (
	"context"

	"github.com/utopiops/automated-ops/ao-api/db"
	"github.com/utopiops/automated-ops/ao-api/models"
)

func New(db *db.DB) PipelineStore {
	return &pipelineStore{db}
}

// NOTE: Many of these endpoints don't get accountId as one of their inputs meaning they don't check if the operation is being performed on the
// jwt token subject's own data or not, this is a BIG VULNERABILITY fix it
type PipelineStore interface {
	// pipelines
	DeletePipeline(context context.Context, accountId, name string) (err error)
	GetPipelineId(context context.Context, accountId, name string) (id int, err error)
	GetPipelineIdByExecution(context context.Context, executionId int) (id int, err error)
	// Create pipelineStore a new pipeline
	Create(context context.Context, base *models.Pipeline, pipeline *models.PipelineVersion) error // todo: return the endpoint
	// Get All pipelines for accountId
	GetPipelines(context context.Context, accountId string) ([]models.Pipeline, error)
	// Retrieve a pipeline based on name
	GetByName(context context.Context, accountId string, name string) (pipeline models.PipelineVersion, endpoint string, err error)
	// Check if the endpoint is valid return the pipeline id
	GetPipelineIdByEndpoint(context context.Context, accountId string, endpoint string) (pipelineId int, err error)

	// tasks
	GetNumberOfTasksForPipeline(context context.Context, pipelineId int) (count int, err error)
	GetTaskByPipelineId(context context.Context, pipelineVersionId int, taskName string) (id int, err error)
	GetTasksWithStatusForExecution(noContext context.Context, executionId int) ([]models.TaskStatusSummery, error)
	GetTaskNameById(noContext context.Context, taskId int) (string, error)
	// Get task details based on execution id and task id
	GetTaskByExecution(context context.Context, executionId int, taskId int) (task models.TaskDetails, err error)
	GetTaskResultDetails(context context.Context, executionId int, taskId int) (res interface{}, err error)
	// Set the status of a task to timed out if it's status is not already set
	SetTaskStatusToTimedout(context context.Context, executionId int, taskId int) (err error)
	// Set the result of a task
	SetTaskResult(context context.Context, executionId int, taskId int, status string) (err error)
	SetTaskResultDetails(context context.Context, executionId int, taskId int, status string, returnValue models.ReturnValueMap, log string) (err error)

	// executions
	GetAllExecutions(context context.Context, pipelineId int) ([]models.Execution, error)
	GetLastExecution(context context.Context, pipelineId int) (id int, err error)
	// Add execution
	CreateExecution(context context.Context, execution models.Execution) (id int, err error)
	// Get initial job of an execution
	GetInitialTask(context context.Context, executionId int) (taskId int, err error)
	// GetInitialData retrieves the initial data of an execution
	GetInitialData(context context.Context, executionId int, accountId string) (InitialData models.InputData, err error)
	// Get next job in an execution based on the status of a task in the execution
	GetNextTasks(context context.Context, executionId int, taskId int, status string) (taskIds []int, err error)
}

type pipelineStore struct {
	db *db.DB
}
