package models

type Specimen struct {
	ID     int    `json:"id"`
	Code   string `json:"code"`
	Text   string `json:"text"`
	status int    `json:"status"`
}
