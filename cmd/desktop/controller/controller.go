package controller

import (
	"fmt"
	"strconv"

	"github.com/e-felix/sebas/internal/command"
	"github.com/e-felix/sebas/internal/env"
	"github.com/e-felix/sebas/internal/project"
)

func GetProjects() map[string]*project.Project {
	projects := make(map[string]*project.Project)
	for i := 1; i <= 3; i++ {
		new_project := project.NewProject(fmt.Sprintf("Project_%d", i))

		for j := 0; j < 3; j++ {
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
