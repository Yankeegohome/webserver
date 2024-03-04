package models

type LabSample struct {
	ID         int    `json:"ID"`
	Ids        string `json:"ids"`
	SpecimenID int    `json:"specimenID"`
	Text       string `json:"text"`
}
