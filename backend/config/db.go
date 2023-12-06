package config

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func InitializeDatabase() (*sqlx.DB, error) {
	dsn := "root:password@tcp(127.0.0.1:3306)/system_employee_management"

	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	// Comprobar la conexión a la base de datos
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	// Crear tabla si no existe
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS Employee (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			version INT NOT NULL
		);
	`)
	if err != nil {
		return nil, fmt.Errorf("error creating table: %w", err)
	}

	return db, nil
}
