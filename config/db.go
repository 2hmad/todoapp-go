package config

import "github.com/jmoiron/sqlx"

func Connect() (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", "root:root_password@tcp(localhost:3306)/todoapp")
	if err != nil {
		panic(err)
	}

	return db, nil
}
