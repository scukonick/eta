package calculator

import (
	"github.com/scukonick/eta/external"
	"github.com/scukonick/eta/repos"
)

type Service struct {
	tasksRepo   repos.Tasks
	resultsRepo repos.Results
	locator     external.Locator
	timer       external.Timer
}

func NewService(tasksRepo repos.Tasks,
	resultsRepo repos.Results,
	locatior external.Locator,
	timer external.Timer,
) *Service {
	return &Service{
		tasksRepo:   tasksRepo,
		resultsRepo: resultsRepo,
		locator:     locatior,
		timer:       timer,
	}
}
