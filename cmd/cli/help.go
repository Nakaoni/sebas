package main

import "fmt"

func help() {
	fmt.Print(`
Usage: sebas [command] [arguments]

commands:
	- project:init    create a blank project.
	- project:load    load a project in current context.
	- cmd:new         create a command for the project in current context. 
	- cmd:update      update a command for the project in current context. 
	- cmd:delete      delete a command for the project in current context. 
	- env:show        show the content of an env file for the project in current context. 
`)
}
