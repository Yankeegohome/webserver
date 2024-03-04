package app

type App interface {
	User() UserRepository
	LabSample() LabSampleRepository
	LabResearch() LabResearchRepository
	Specimen() SpecimenRepository
}
