package model

type Report struct {
	Message string
}

type WithCount struct {
	Results interface{}
	Count   int64
}
