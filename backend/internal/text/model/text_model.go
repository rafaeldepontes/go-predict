package model

type TextReq struct {
	Body     string `json:"body"`
	TeamSize int    `json:"size"`
	Stack    string `json:"stack"`
	Level    string `json:"level"`
}
