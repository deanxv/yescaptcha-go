package res

type CreateTaskResponse struct {
	ErrorId          int    `json:"errorId"`
	ErrorCode        string `json:"errorCode"`
	ErrorDescription string `json:"errorDescription"`
	TaskId           string `json:"taskId"`
	Status           string `json:"status"`
	Solution         struct {
		Text      string        `json:"text"`
		Type      string        `json:"type"`
		Objects   []interface{} `json:"objects"`
		HasObject bool          `json:"hasObject"`
		Box       []string      `json:"box"`
		Labels    []string      `json:"labels"`
	} `json:"solution"`
}
