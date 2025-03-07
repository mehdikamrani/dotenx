import axios from 'axios'
export const API_URL = process.env.REACT_APP_API_URL

const api = axios.create({
	baseURL: API_URL,
})

export function getPipelines() {
	return api.get<Pipeline[]>('/pipeline')
}

export function addPipeline(payload: AddPipelinePayload) {
	return api.post<void>('/pipeline', payload)
}

export function startPipeline(payload: StartPipelinePayload) {
	return api.post<void>(`/execution/ep/${payload.endpoint}/start`, {})
}

export function getTasks() {
	return api.get<TasksData>('/task')
}

export function getTaskFields(taskName: string) {
	return api.get<TaskFields>(`/task/${taskName}/fields`)
}

export function getPipeline(name: string) {
	return api.get<PipelineData>(`/pipeline/name/${name}`)
}

export function getResult(executionId: string, taskName: string) {
	return api.get<TaskResult>(`/execution/id/${executionId}/task_name/${taskName}/result`)
}

export function getExecutions(pipelineName: string) {
	return api.get<Execution[]>(`/pipeline/name/${pipelineName}/executions`)
}

export function getIntegrationTypes() {
	return api.get<string[]>('/integration/avaliable')
}

export function getIntegrationTypeFields(integrationType: string) {
	return api.get<string[]>(`/integration/type/${integrationType}/fields`)
}

export function getIntegrations() {
	return api.get<IntegrationData[]>(`/integration`)
}

export function addIntegration(payload: AddIntegrationPayload) {
	return api.post<void>('/integration', payload)
}

export function getTriggers() {
	return api.get<TriggerData[]>('/trigger')
}

export function getTriggerTypes() {
	return api.get<string[]>('/trigger/avaliable')
}

export function getTriggerDefinition(type: string) {
	return api.get<TriggerDefinition>(`/trigger/type/${type}/definition`)
}

export function addTrigger(payload: AddTriggerPayload) {
	return api.post<void>('/trigger', payload)
}

export function getPipelineTriggers(pipelineName: string) {
	return api.get<TriggerData[]>(`/trigger`, { params: { pipeline: pipelineName } })
}

export function deletePipeline(name: string) {
	return api.delete<void>(`/pipeline/name/${name}`)
}

export function deleteIntegration(name: string) {
	return api.delete<void>(`/integration/name/${name}`)
}

export function deleteTrigger(name: string) {
	return api.delete<void>(`/trigger/name/${name}`)
}

export function getIntegrationsByType(type: string) {
	return api.get<IntegrationData[]>(`/integration/type/${type}`)
}

export enum QueryKey {
	GetPipelines = 'get-pipelines',
	GetTasks = 'get-tasks',
	GetTaskFields = 'get-task-fields',
	GetPipeline = 'get-pipeline',
	GetResult = 'get-result',
	GetExecutions = 'get-executions',
	GetIntegrationTypes = 'get-integration-types',
	GetIntegrationTypeFields = 'get-integration-type-fields',
	GetIntegrations = 'get-integrations',
	GetTriggers = 'get-triggers',
	GetTriggerTypes = 'get-trigger-types',
	GetTriggerDefinition = 'get-trigger-definition',
	GetPipelineTriggers = 'get-pipeline-triggers',
	GetIntegrationsByType = 'get-integration-by-type',
}

export enum Status {
	Success = 'success',
	Failed = 'failed',
	Timedout = 'timedout',
	Started = 'started',
	Cancelled = 'cancelled',
	Completed = 'completed',
	Waiting = 'waiting',
}

export enum TaskType {
	Text = 'text',
}

export interface TriggerData {
	name: string
	account_id: string
	type: string
	endpoint: string
	pipeline_name: string
	integration: string
	credentials: Record<string, string>
}

export interface AddTriggerPayload {
	name: string
	type: string
	pipeline_name: string
	integration: string
	credentials: Record<string, string>
}

export interface TriggerDefinition {
	type: string
	integration: string
	image: string
	credentials: FieldType[]
}

export interface FieldType {
	Key: string
	Type: string
}

export interface IntegrationData {
	name: string
	account_id: string
	type: string
	url: string
	key: string
	secret: string
	access_token: string
}

export interface AddIntegrationPayload {
	name: string
	type: string
}

export interface Execution {
	Id: number
	StartedAt: string
	InitialData: unknown | null
}

export interface PipelineEventMessage {
	execution_id: string
	tasks: {
		name: string
		status: Status
	}[]
}

export interface TaskResult {
	status: Status
	log: string
	return_value: string
}

export interface PipelineData {
	serviceAccount: string
	endpoint: string
	manifest: Manifest
}

export interface TaskFields {
	fields: {
		key: string
		type: string
	}[]
	integration_type: string
}

export interface TasksData {
	tasks: string[]
}

export interface StartPipelinePayload {
	endpoint: string
}

export interface Pipeline {
	name: string
	endpoint: string
}

export interface AddPipelinePayload {
	name: string
	manifest: Manifest
}

export interface Manifest {
	triggers: Record<string, Trigger>
	tasks: Tasks
}

export type Tasks = Record<string, Task>

export interface Trigger {
	type: string
}

export interface Task {
	type: string
	executeAfter: Record<string, string[]>
	body: Record<string, string>
	integration: string
}
