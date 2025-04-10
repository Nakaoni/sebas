package controller

import (
	"fmt"
	"github.com/e-felix/sebas/internal/command"
	"github.com/e-felix/sebas/internal/env"
	"github.com/e-felix/sebas/internal/project"
	"strconv"
)

var projects = make(map[string]*project.Project)

func GetProjects() []string {
	projects_list := make([]string, 0)

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
		projects_list = append(projects_list, new_project.Name)
	}

	return projects_list
}

type CommandDto struct {
	Id   uint
	Path string
	Args []string
}

var commandDtos = make([]*CommandDto, 0)
var commands = make([]*command.Command, 0)

func GetCmds(project_name string) []*CommandDto {
	if projects[project_name] == nil {
		return commandDtos
	}

	selected_project := projects[project_name]
	for i, cmd := range selected_project.Cmds {
		cmdDto := CommandDto{
			Id:   uint(i + 1),
			Path: cmd.Path,
			Args: cmd.Args,
		}
		commandDtos = append(commandDtos, &cmdDto)
		commands = append(commands, &cmd)
	}

	return commandDtos
}
func RunCmd(command_id uint) (string, error) {
	if nil == commandDtos[command_id-1] {
		return "", fmt.Errorf("could not find actual_command")
	}

	actual_command := commands[command_id-1]

	c := make(chan string)
	go actual_command.Run(c)

	return <-c, nil
}
