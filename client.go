package yescaptcha

import (
	"bytes"
	"encoding/json"
	"github.com/deanxv/yescaptcha-go/req"
	"github.com/deanxv/yescaptcha-go/res"
	"github.com/deanxv/yescaptcha-go/url"
	"io"
	"net/http"
)

type Client struct {
	ClientKey string
	BaseUrl   string
	SoftID    string
}

func NewClient(clientKey, softID, baseUrl string) *Client {
	if baseUrl == "" {
		baseUrl = url.BaseUrlCN
	}
	return &Client{
		ClientKey: clientKey,
		BaseUrl:   baseUrl,
		SoftID:    softID,
	}
}

// 查询帐户余额
func (c *Client) GetBalance() (getBalanceResponse res.GetBalanceResponse, err error) {
	resBytes, err := c.post(c.BaseUrl+url.GetBalance, &req.GetBalanceRequest{ClientKey: c.ClientKey})
	if err != nil {
		return
	}
	err = decodeResponse(resBytes, &getBalanceResponse)
	return
}

// 创建识别任务-ImageToTextTask : 图片不定长英文数字
// refs: https://yescaptcha.atlassian.net/wiki/spaces/YESCAPTCHA/pages/164300/ImageToTextTask
func (c *Client) CreateImageToText(task *req.ImageToTextTaskRequest) (createTaskResponse res.CreateTaskResponse, err error) {
	resBytes, err := c.post(c.BaseUrl+url.CreateTask, &req.CreateTaskRequest{ClientKey: c.ClientKey, Task: task, SoftID: c.SoftID})
	if err != nil {
		return
	}
	err = decodeResponse(resBytes, &createTaskResponse)
	return
}

// 创建识别任务-NoCaptchaTaskProxyless : reCaptcha V2 协议接口
// refs: https://yescaptcha.atlassian.net/wiki/spaces/YESCAPTCHA/pages/229796/NoCaptchaTaskProxyless+reCaptcha+V2
func (c *Client) CreateNoCaptchaTaskProxyless(task *req.NoCaptchaTaskProxylessRequest) (createTaskResponse res.CreateTaskResponse, err error) {
	resBytes, err := c.post(c.BaseUrl+url.CreateTask, &req.CreateTaskRequest{ClientKey: c.ClientKey, Task: task, SoftID: c.SoftID})
	if err != nil {
		return
	}
	err = decodeResponse(resBytes, &createTaskResponse)
	return
}

// 创建识别任务-RecaptchaV3TaskProxyless: reCaptcha V3 协议接口
// refs: https://yescaptcha.atlassian.net/wiki/spaces/YESCAPTCHA/pages/655381/RecaptchaV3TaskProxyless+reCaptcha+V3
func (c *Client) CreateRecaptchaV3TaskProxyless(task *req.RecaptchaV3TaskProxylessRequest) (createTaskResponse res.CreateTaskResponse, err error) {
	resBytes, err := c.post(c.BaseUrl+url.CreateTask, &req.CreateTaskRequest{ClientKey: c.ClientKey, Task: task, SoftID: c.SoftID})
	if err != nil {
		return
	}
	err = decodeResponse(resBytes, &createTaskResponse)
	return
}

// 创建识别任务-RecaptchaV2EnterpriseTaskProxyless: 企业版 reCaptcha V2
// refs: https://yescaptcha.atlassian.net/wiki/spaces/YESCAPTCHA/pages/2916435/RecaptchaV2EnterpriseTaskProxyless+reCaptcha+V2
func (c *Client) CreateRecaptchaV2EnterpriseTaskProxyless(task *req.RecaptchaV2EnterpriseTaskProxylessRequest) (createTaskResponse res.CreateTaskResponse, err error) {
	resBytes, err := c.post(c.BaseUrl+url.CreateTask, &req.CreateTaskRequest{ClientKey: c.ClientKey, Task: task, SoftID: c.SoftID})
	if err != nil {
		return
	}
	err = decodeResponse(resBytes, &createTaskResponse)
	return
}

// 创建识别任务-ReCaptchaV2Classification: reCaptcha V2 图像识别
// refs: https://yescaptcha.atlassian.net/wiki/spaces/YESCAPTCHA/pages/18055169/ReCaptchaV2Classification+reCaptcha+V2
func (c *Client) CreateReCaptchaV2Classification(task *req.ReCaptchaV2ClassificationRequest) (createTaskResponse res.CreateTaskResponse, err error) {
	resBytes, err := c.post(c.BaseUrl+url.CreateTask, &req.CreateTaskRequest{ClientKey: c.ClientKey, Task: task, SoftID: c.SoftID})
	if err != nil {
		return
	}
	err = decodeResponse(resBytes, &createTaskResponse)
	return
}

// 创建识别任务-HCaptchaTaskProxyless : HCaptcha 协议接口
// refs: https://yescaptcha.atlassian.net/wiki/spaces/YESCAPTCHA/pages/7929858/HCaptchaTaskProxyless+HCaptcha
func (c *Client) CreateHCaptchaTaskProxyless(task *req.HCaptchaTaskProxylessRequest) (createTaskResponse res.CreateTaskResponse, err error) {
	resBytes, err := c.post(c.BaseUrl+url.CreateTask, &req.CreateTaskRequest{ClientKey: c.ClientKey, Task: task, SoftID: c.SoftID})
	if err != nil {
		return
	}
	err = decodeResponse(resBytes, &createTaskResponse)
	return
}

// 创建识别任务-HCaptchaClassification: Hcaptcha图像识别
// refs: https://yescaptcha.atlassian.net/wiki/spaces/YESCAPTCHA/pages/24543233/HCaptchaClassification+Hcaptcha
func (c *Client) CreateHCaptchaClassification(task *req.HCaptchaClassificationRequest) (createTaskResponse res.CreateTaskResponse, err error) {
	resBytes, err := c.post(c.BaseUrl+url.CreateTask, &req.CreateTaskRequest{ClientKey: c.ClientKey, Task: task, SoftID: c.SoftID})
	if err != nil {
		return
	}
	err = decodeResponse(resBytes, &createTaskResponse)
	return
}

// 创建识别任务-FuncaptchaTaskProxyless : Funcaptcha 协议接口
// refs: https://yescaptcha.atlassian.net/wiki/spaces/YESCAPTCHA/pages/17825793/FuncaptchaTaskProxyless+Funcaptcha
func (c *Client) CreateFuncaptchaTaskProxyless(task *req.FuncaptchaTaskProxylessRequest) (createTaskResponse res.CreateTaskResponse, err error) {
	resBytes, err := c.post(c.BaseUrl+url.CreateTask, &req.CreateTaskRequest{ClientKey: c.ClientKey, Task: task, SoftID: c.SoftID})
	if err != nil {
		return
	}
	err = decodeResponse(resBytes, &createTaskResponse)
	return
}

// 创建识别任务-FunCaptchaClassification: Funcaptcha 图像识别
// refs: https://yescaptcha.atlassian.net/wiki/spaces/YESCAPTCHA/pages/34209793/FunCaptchaClassification+Funcaptcha
func (c *Client) CreateFunCaptchaClassification(task *req.FunCaptchaClassificationRequest) (createTaskResponse res.CreateTaskResponse, err error) {
	resBytes, err := c.post(c.BaseUrl+url.CreateTask, &req.CreateTaskRequest{ClientKey: c.ClientKey, Task: task, SoftID: c.SoftID})
	if err != nil {
		return
	}
	err = decodeResponse(resBytes, &createTaskResponse)
	return
}

// 创建识别任务-TurnstileTaskProxyless: CloudflareTurnstile协议接口
// refs: https://yescaptcha.atlassian.net/wiki/spaces/YESCAPTCHA/pages/61734913/TurnstileTaskProxyless+CloudflareTurnstile
func (c *Client) CreateTurnstileTaskProxyless(task *req.TurnstileTaskProxylessRequest) (createTaskResponse res.CreateTaskResponse, err error) {
	resBytes, err := c.post(c.BaseUrl+url.CreateTask, &req.CreateTaskRequest{ClientKey: c.ClientKey, Task: task, SoftID: c.SoftID})
	if err != nil {
		return
	}
	err = decodeResponse(resBytes, &createTaskResponse)
	return
}

// 获取识别结果
func (c *Client) GetTaskResult(taskId string) (getTaskResultResponse res.GetTaskResultResponse, err error) {
	resBytes, err := c.post(c.BaseUrl+url.GetTaskResult, &req.GetTaskResultRequest{ClientKey: c.ClientKey, TaskId: taskId})
	if err != nil {
		return
	}
	err = decodeResponse(resBytes, &getTaskResultResponse)
	return
}

// 查询开发者ID和金额 21585
func (c *Client) GetSoftID() (getSoftIDResponse res.GetSoftIDResponse, err error) {
	resBytes, err := c.post(c.BaseUrl+url.GetSoftID, &req.GetSoftIDRequest{ClientKey: c.ClientKey})
	if err != nil {
		return
	}
	err = decodeResponse(resBytes, &getSoftIDResponse)
	return
}

// 定义一个通用的POST方法
func (c *Client) post(url string, requestBody interface{}) ([]byte, error) {
	jsonRequest, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonRequest))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// 定义一个通用的GET方法
func (c *Client) get(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func decodeResponse(bytes []byte, v any) error {
	if v == nil {
		return nil
	}

	if result, ok := v.(*string); ok {
		return decodeString(bytes, result)
	}
	return json.Unmarshal(bytes, v)
}

func decodeString(bytes []byte, output *string) error {
	*output = string(bytes)
	return nil
}
