package app

import (
	"github.com/magic-of-gnu/crm-local-v8/server/internal/repos"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/services"
)

type App struct {
	// repos
	CentresRepo repos.CentresRepo

	// services
	CentresService services.CentresService

	MethodNames map[string]string
}

func NewApp(
	centresRepo repos.CentresRepo,
	centresService services.CentresService,
	methodNames map[string]string) *App {
	return &App{
		CentresRepo:    centresRepo,
		CentresService: centresService,
		MethodNames:    methodNames,
	}
}
