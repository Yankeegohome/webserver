package app

import "web/models"

type UserRepository interface {
	Create(user *models.User) error
}

type LabSampleRepository interface {
	CreateLabSample(labSample *models.LabSample) error
}

type LabResearchRepository interface {
	CreateLabResearch(labResearch *models.LabResearch) error
}

type SpecimenRepository interface {
	CreateSpecimen(specimen *models.Specimen) error
}
