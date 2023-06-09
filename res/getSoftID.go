package res

type GetSoftIDResponse struct {
	ErrorId int     `json:"errorId"`
	Balance float64 `json:"balance"`
	SoftID  int     `json:"softID"`
}
