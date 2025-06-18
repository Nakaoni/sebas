package controller

import (
	"fmt"
	"strconv"

	"github.com/e-felix/sebas/internal/command"
	"github.com/e-felix/sebas/internal/env"
	"github.com/e-felix/sebas/internal/project"
)

var projects map[string]*project.Project

func GenerateFixture() map[string]*project.Project {
	projects = make(map[string]*project.Project)

	for i := 1; i <= 3; i++ {
		new_project := project.NewProject(fmt.Sprintf("Project_%d", i))
		new_project.Id = i

		for j := 0; j < 10; j++ {
			new_project.AddEnv(env.Env{
				Key:   fmt.Sprintf("ENV_%d", j),
				Value: fmt.Sprintf("VALUE_%d", j),
			})
			new_project.AddCmd(command.Command{
				Path: "echo",
				Args: []string{new_project.Name, strconv.Itoa(j)},
			})
		}

		projects[new_project.Name] = new_project
	}

	return projects
}

func AddProject(name string) map[string]*project.Project {
	project := project.NewProject(name)
	project.Id = len(projects)

	projects[project.Name] = project

	return projects
}
