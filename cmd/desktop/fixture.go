package main

import (
	"github.com/e-felix/sebas/cmd/desktop/controller"
	"github.com/e-felix/sebas/internal/project"
)

func initFixture() map[string]*project.Project {
	return controller.GenerateFixture()
}