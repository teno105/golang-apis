package models

import "encoding/json"

type Notice struct {
	LocalList []NoticeLanguageData `json:"local_list"`
}

type NoticeLanguageData struct {
	Language string `json:"language",omitempty`
	Title    string `json:"title"`
	Body     string `json:"body"`
}