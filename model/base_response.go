package model

type Report struct {
	Message string `json:"message"`
}

type WithCount struct {
	Data  interface{} `json:"data"`
	Count int64       `json:"count"`
}
