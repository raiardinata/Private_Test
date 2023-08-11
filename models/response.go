package models

type Response struct {
	Status   int         `json:"status"`
	Messages string      `json:"message"`
	Data     interface{} `json:"data"`
}
