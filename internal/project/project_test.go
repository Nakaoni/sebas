package project

import (
	"testing"

	. "github.com/e-felix/sebas/internal/command"
	. "github.com/e-felix/sebas/internal/env"
	Assert "github.com/e-felix/sebas/internal/util/assert"
)

func TestNewProject(t *testing.T) {
	expected := &Project{Id: 1, Name: "MyProject", Envs: make([]Env, 0), Cmds: make([]Command, 0)}

	project := NewProject("MyProject")

	Assert.DeepEqual(project, expected)
}

func TestUpdateProject(t *testing.T) {
	newName := "MyNewName"
	expected := &Project{Id: 1, Name: newName, Envs: make([]Env, 0), Cmds: make([]Command, 0)}

	project := NewProject("MyProject")
	project.Update(newName)

	Assert.DeepEqual(project, expected)
}

func TestAddEnv(t *testing.T) {
	env := Env{Key: "FOO", Value: "BAR"}
	expected := &Project{Id: 1, Name: "MyProject", Envs: []Env{env}, Cmds: make([]Command, 0)}

	project := NewProject("MyProject")
	project.AddEnv(env)

	Assert.DeepEqual(project, expected)
}

func TestRemoveEnv(t *testing.T) {
	env := Env{Key: "FOO", Value: "BAR"}
	expected := &Project{Id: 1, Name: "MyProject", Envs: make([]Env, 0), Cmds: make([]Command, 0)}

	project := NewProject("MyProject")
	project.AddEnv(env)
	project.RemoveEnv(env)

	Assert.DeepEqual(project, expected)
}

func TestAddCommand(t *testing.T) {
	cmd := Command{Path: "echo", Args: []string{"Hello", "World"}}
	expected := &Project{Id: 1, Name: "MyProject", Envs: make([]Env, 0), Cmds: []Command{cmd}}

	project := NewProject("MyProject")
	project.AddCmd(cmd)

	Assert.DeepEqual(project, expected)
}

func TestRemoveCommand(t *testing.T) {
	cmd := Command{Path: "echo", Args: []string{"Hello", "World"}}
	expected := &Project{Id: 1, Name: "MyProject", Envs: make([]Env, 0), Cmds: make([]Command, 0)}

	project := NewProject("MyProject")
	project.AddCmd(cmd)
	project.RemoveCmd(cmd)

	Assert.DeepEqual(project, expected)
}
