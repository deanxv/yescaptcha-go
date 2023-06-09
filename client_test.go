package yescaptcha

import (
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	client := NewClient("", "21585", "")
	response, err := client.GetBalance()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(response.Balance)
	}
	//taskResponse, err := client.CreateImageToText(&req.ImageToTextTaskRequest{Type: task.ImageToTextTaskM1, Body: "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"})
	//if err != nil {
	//	fmt.Println(err.Error())
	//} else {
	//	fmt.Println(taskResponse.TaskId)
	//}
	// 6ba4e134-08d2-11ee-865a-e2369b0517b2
	resultResponse, err := client.GetTaskResult("6ba4e134-08d2-11ee-865a-e2369b0517b2")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(resultResponse.Solution)
	}
}
