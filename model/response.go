package model

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type FailedResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
