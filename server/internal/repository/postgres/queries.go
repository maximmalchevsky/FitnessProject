package postgres

const (
	createTableTrainings = `
		CREATE TABLE IF NOT EXISTS trainings (
		    id SERIAL PRIMARY KEY,
		    name VARCHAR NOT NULL,
		    phone VARCHAR NOT NULL,
		    confirmation VARCHAR NOT NULL DEFAULT '-'
		);
`
	createTablePasses = `
		CREATE TABLE IF NOT EXISTS passes (
		    id SERIAL PRIMARY KEY,
		    name VARCHAR NOT NULL,
		    phone VARCHAR NOT NULL,
		    type VARCHAR NOT NULL,
		    Duration VARCHAR NOT NULL
		)
`
)
