package sqlapp

import (
	"database/sql"
	"web/iternal/app"
)

type App struct {
	db                    *sql.DB
	userRepository        *UserRepository
	labResearchRepository *LabResearchRepository
	labSampleRepository   *LabSampleRepository
	specimen              *SpecimenRepository
}

func New(db *sql.DB) *App {
	return &App{db: db}
}

func (a *App) User() app.UserRepository {
	if a.userRepository != nil {
		return a.userRepository
	}
	a.userRepository = &UserRepository{
		app: a,
	}
	return a.userRepository
}

func (a *App) LabSample() app.LabSampleRepository {
	if a.labSampleRepository != nil {
		return a.labSampleRepository
	}
	a.labSampleRepository = &LabSampleRepository{
		app: a,
	}
	return a.labSampleRepository
}

func (a *App) LabResearch() app.LabResearchRepository {
	if a.labResearchRepository != nil {
		return a.labResearchRepository
	}
	a.labResearchRepository = &LabResearchRepository{
		app: a,
	}
	return a.labResearchRepository
}
func (a *App) Specimen() app.SpecimenRepository {
	if a.Specimen() != nil {
		return a.specimen
	}
	a.specimen = &SpecimenRepository{
		app: a,
	}
	return a.specimen
}
