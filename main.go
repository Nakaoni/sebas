package main

import (
	"database/sql"
	"os"
	"os/exec"
	"path"
	"plugin"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nakaoni/sebas/internal/util"
)

func initDB(db *sql.DB) {
	cwd, _ := os.Getwd()
	basePath := path.Join(cwd, "migrations")
	dirEntry, _ := os.ReadDir(basePath)

	for _, entry := range dirEntry {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		versionPath := path.Join(basePath, name)

		compiledVersionPath := path.Join(os.TempDir(), strings.Replace(name, ".go", ".so", 1))
		cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", compiledVersionPath, versionPath)
		if err := cmd.Run(); err != nil {
			panic(err)
		}

		p, err := plugin.Open(compiledVersionPath)
		if err != nil {
			panic(err)
		}

		f, err := p.Lookup("Up")
		if err != nil {
			panic(err)
		}

		tx, err := db.Begin()
		if err != nil {
			panic(err)
		}
		defer tx.Rollback()

		up := f.(func() string)()
		if _, err = tx.Exec(up); err != nil {
			panic(err)
		}

		_, err = db.Exec(`
		INSERT INTO migrations (version, created_at, executed_at)
		VALUES ($1, NOW(), NOW())
		`, name)
		if err != nil {
			panic(err)
		}

		if err := tx.Commit(); err != nil {
			panic(err)
		}
	}
}

func main() {
	db, err := util.GetDatabase()
	if err != nil {
		panic(err)
	}

	initDB(db)

	// result, err := db.Exec(`
	// DROP TABLE IF EXISTS commands;
	// `)
	// if err != nil {
	// 	panic(err)
	// }
	// i, _ := result.RowsAffected()
	// fmt.Println("drop table: ", i)

	// result, err = db.Exec(`
	// CREATE TABLE commands (
	// 	id INT PRIMARY KEY,
	// 	cmd VARCHAR(500),
	// 	args VARCHAR(500)
	// )
	// `)
	// if err != nil {
	// 	panic(err)
	// }

	// i, _ = result.RowsAffected()
	// fmt.Println("create table: ", i)

	// result, err = db.Exec(`
	// INSERT INTO commands (cmd, args)
	// VALUES
	// ("echo", "hello world"),
	// ("ls", "-la")
	// `)
	// if err != nil {
	// 	panic(err)
	// }

	// i, _ = result.RowsAffected()
	// fmt.Println("rows: ", i)

	// cmds := repo.CommandAll()

	// for _, c := range cmds {
	// 	fmt.Println(c)
	// }
}
