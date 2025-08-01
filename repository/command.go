package repository

import (
	"log"

	"github.com/nakaoni/sebas/internal/command"
	"github.com/nakaoni/sebas/internal/util"
)

func CommandAll() []command.Command {
	db, err := util.GetDatabase()
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM commands")
	if err != nil {
		log.Fatalln(err)
		return []command.Command{}
	}
	defer rows.Close()

	commands := make([]command.Command, 0)
	for rows.Next() {
		var id any
		var cmd string
		var args string
		if err := rows.Scan(&id, &cmd, &args); err != nil {
			log.Fatalln(err)
			continue
		}

		commands = append(commands, command.Command{
			Cmd:  cmd,
			Args: []string{args},
		})
	}

	if err := rows.Err(); err != nil {
		log.Fatalln(err)
		return []command.Command{}
	}

	return commands
}
