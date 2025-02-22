package models

import "encoding/json"

type Notice struct {
	LocalList []NoticeData `json:"local_list"`
}

type NoticeData struct {
	Language string `json:"language",omitempty`
	Title    string `json:"title"`
	Body     string `json:"body"`
}

