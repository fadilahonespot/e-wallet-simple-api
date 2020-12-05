package model

type ResponWrapper struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponWrapperMessage struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
}

