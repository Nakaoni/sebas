package migrations

type Version interface {
	Up() string
	Down() string
}

type Version0 struct{}

func (i Version0) Up() string {
	return `
	CREATE TABLE IF NOT EXISTS migrations (
		id INT PRIMARY KEY,
		version VARCHAR(255) NOT NULL,
		created_at DATETIME NOT NULL,
		executed_at DATETIME NOT NULL
	)
	`
}

func (i Version0) Down() string {
	return `
	DROP TABLE migrations
	`
}
