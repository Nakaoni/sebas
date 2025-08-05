package migrations

type Version1754416414 struct{}

func (v Version1754416414) Up() string {
	return `
	CREATE TABLE IF NOT EXISTS commands (
		id INT PRIMARY KEY,
		cmd VARCHAR(500) NOT NULL,
		args VARCHAR(500),
		created_at DATETIME,
		updated_at DATETIME
	)
	`
}

func (v Version1754416414) Down() string {
	return `
	DROP TABLE commands
	`
}
