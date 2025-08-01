package project

import (
	"fmt"
	"reflect"

	. "github.com/nakaoni/sebas/internal/command"
	. "github.com/nakaoni/sebas/internal/env"
)

type Project struct {
	Id   int
	Name string
	Envs []Env
	Cmds []Command
}

func NewProject(name string) *Project {
	return &Project{
		Id:   1,
		Name: name,
		Envs: make([]Env, 0),
		Cmds: make([]Command, 0),
	}
}

func (p *Project) Print() {
	fmt.Printf("Id: %d\n", p.Id)
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Println("Envs:")

	for _, env := range p.Envs {
		fmt.Printf("\t%v: %v\n", env.Key, env.Value)
	}
}

func (p *Project) Update(name string) {
	p.Name = name
}

func (p *Project) AddEnv(newEnv Env) bool {
	envs := p.Envs

	envExists := false
	for _, env := range envs {
		if env.Key == newEnv.Key {
			envExists = true
			break
		}

	}

	if !envExists {
		p.Envs = append(p.Envs, newEnv)
		return true
	}

	return false
}

func (p *Project) RemoveEnv(envToRemove Env) bool {
	newEnvList := make([]Env, 0)

	exists := false
	for _, env := range p.Envs {
		if reflect.DeepEqual(env, envToRemove) {
			exists = true
			continue
		}

		newEnvList = append(newEnvList, env)
	}

	if !exists {
		return false
	}

	p.Envs = newEnvList

	return true
}

func (p *Project) AddCmd(newCmd Command) bool {
	p.Cmds = append(p.Cmds, newCmd)

	return true
}

func (p *Project) RemoveCmd(cmdToRemove Command) bool {
	newCmdList := make([]Command, 0)

	exists := false
	for _, cmd := range p.Cmds {
		if reflect.DeepEqual(cmd, cmdToRemove) {
			exists = true
			continue
		}

		newCmdList = append(newCmdList, cmd)
	}

	if !exists {
		return false
	}

	p.Cmds = newCmdList

	return true
}
