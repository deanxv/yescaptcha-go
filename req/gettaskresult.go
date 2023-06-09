package req

type GetTaskResultRequest struct {
	ClientKey string `json:"clientKey"`
	TaskId    string `json:"taskId"`
}
