package model

import "net/http"

type HttpContext struct {
	Request   *http.Request
	Url       string
	Method    string
	StartTime int64
	EndTime   int64
	TraceId   string
	Ip        string
	Status    int
}
