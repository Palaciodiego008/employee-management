package internal

import "github.com/jmoiron/sqlx"

func AddSampleEmployees(db *sqlx.DB) error {
	_, err := db.Exec("INSERT INTO Employee (name, version) VALUES (?, ?)", "Employee1", 1)
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO Employee (name, version) VALUES (?, ?)", "Employee2", 1)
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO Employee (name, version) VALUES (?, ?)", "Employee3", 1)
	return err
}
