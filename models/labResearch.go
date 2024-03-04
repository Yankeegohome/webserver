package models

type LabResearch struct {
	ID          int    `json:"id"`
	LabSampleID int    `json:"labSampleID"`
	Code        string `json:"code"`
	Text        string `json:"text"`
}
