package res

type GetTaskResultResponse struct {
	ErrorId          int    `json:"errorId"`
	ErrorCode        string `json:"errorCode"`
	ErrorDescription string `json:"errorDescription"`
	Solution         struct {
		Text               string `json:"text"`
		GRecaptchaResponse string `json:"gRecaptchaResponse"`
		UserAgent          string `json:"userAgent"`
		RespKey            string `json:"respKey"`
		Token              string `json:"token"`
	} `json:"solution"`
	Status string `json:"status"`
}
