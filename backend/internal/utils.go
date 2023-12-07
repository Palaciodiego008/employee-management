package internal

import "github.com/jmoiron/sqlx"

func AddSampleEmployees(db *sqlx.DB) error {
	employees := []struct {
		Name      string
		Version   int
		ManagerID int
	}{
		{"Employee1", 1, 0}, // 0 porque estos son los empleados iniciales sin manager
		{"Employee2", 1, 0},
		{"Employee3", 1, 0},
	}

	existsQuery := "SELECT COUNT(*) FROM Employee WHERE name = ? AND version = ?"

	for _, emp := range employees {
		var count int
		err := db.Get(&count, existsQuery, emp.Name, emp.Version)
		if err != nil {
			return err
		}

		if count == 0 {
			_, err = db.Exec("INSERT INTO Employee (name, manager_id, version) VALUES (?, ?, ?)", emp.Name, emp.ManagerID, emp.Version)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
