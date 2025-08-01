package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nakaoni/sebas/internal/util"
	repo "github.com/nakaoni/sebas/repository"
)

func main() {

	db, err := util.GetDatabase()
	if err != nil {
		panic(err)
	}

	result, err := db.Exec(`
	DROP TABLE commands;
	`)
	if err != nil {
		panic(err)
	}
	i, _ := result.RowsAffected()
	fmt.Println("drop table: ", i)

	result, err = db.Exec(`
	CREATE TABLE commands (
		id INT PRIMARY KEY,
		cmd VARCHAR(500),
		args VARCHAR(500)
	)
	`)
	if err != nil {
		panic(err)
	}

	i, _ = result.RowsAffected()
	fmt.Println("create table: ", i)

	result, err = db.Exec(`
	INSERT INTO commands (cmd, args)
	VALUES
	("echo", "hello world"),
	("ls", "-la")
	`)
	if err != nil {
		panic(err)
	}

	i, _ = result.RowsAffected()
	fmt.Println("rows: ", i)

	cmds := repo.CommandAll()

	for _, c := range cmds {
		fmt.Println(c)
	}
}
