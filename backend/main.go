package main

import (
	"employee_system_management/config"
	"employee_system_management/internal"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// Employee representa la estructura de un empleado.
type Employee struct {
	ID      int       `json:"id" db:"id"`
	Name    string    `json:"name" db:"name"`
	Manager *Employee `json:"manager,omitempty" db:"-"`
	Version int       `json:"version" db:"version"`
}

// EmployeeService representa el servicio de gestión de empleados.
type EmployeeService struct {
	mutex sync.Mutex
}

func (s *EmployeeService) GetHierarchy(w http.ResponseWriter, r *http.Request) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	var employees []Employee

	err := db.Select(&employees, `
		SELECT id, name, version
		FROM Employee
	`)
	if err != nil {
		http.Error(w, "Error fetching employees", http.StatusInternalServerError)
		return
	}

	// Construir jerarquía de empleados asignando instancias de Manager
	for i, e := range employees {
		if e.Manager != nil {
			employees[i].Manager.Version = e.Manager.Version
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

// Agregar un nuevo empleado
func (s *EmployeeService) AddEmployee(w http.ResponseWriter, r *http.Request) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Decodificar el cuerpo JSON de la solicitud para obtener los datos del nuevo empleado
	var newEmployee Employee
	err := json.NewDecoder(r.Body).Decode(&newEmployee)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	// Insertar el nuevo empleado en la base de datos
	_, err = db.Exec("INSERT INTO Employee (name, version) VALUES (?, ?)", newEmployee.Name, newEmployee.Version)
	if err != nil {
		http.Error(w, "Error adding new employee", http.StatusInternalServerError)
		return
	}

	// Enviar respuesta exitosa
	w.WriteHeader(http.StatusCreated)
}

func (s *EmployeeService) UpdateManager(w http.ResponseWriter, r *http.Request) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	vars := mux.Vars(r)
	employeeID := vars["employeeID"]
	managerID := vars["managerID"]

	eID, err := strconv.Atoi(employeeID)
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	mID, err := strconv.Atoi(managerID)
	if err != nil {
		http.Error(w, "Invalid manager ID", http.StatusBadRequest)
		return
	}

	var employee Employee
	err = db.Get(&employee, "SELECT * FROM Employee WHERE id = ?", eID)
	if err != nil {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	var manager Employee
	err = db.Get(&manager, "SELECT * FROM Employee WHERE id = ?", mID)
	if err != nil {
		http.Error(w, "Manager not found", http.StatusNotFound)
		return
	}

	// Comprobar si el empleado ya tiene asignado el mismo manager
	if employee.Manager != nil && employee.Manager.ID == manager.ID {
		http.Error(w, "Employee already has the same manager", http.StatusBadRequest)
		return
	}

	// Actualizar el jefe en la base de datos
	_, err = db.Exec("UPDATE Employee SET version = version + 1 WHERE id = ?", eID)
	if err != nil {
		http.Error(w, "Error updating employee version", http.StatusInternalServerError)
		return
	}

	// Actualizar el jefe en el objeto Employee
	employee.Manager = &manager
	employee.Version++

	// Enviar respuesta exitosa
	w.WriteHeader(http.StatusOK)
}

func main() {
	var err error
	db, err = config.InitializeDatabase()
	if err != nil {
		fmt.Println("Error initializing database:", err)
		return
	}
	defer db.Close()

	r := mux.NewRouter()
	service := &EmployeeService{}

	// Añadir empleados de ejemplo
	err = internal.AddSampleEmployees(db)
	if err != nil {
		fmt.Println("Error adding sample employees:", err)
		return
	}

	// Establecer rutas
	r.HandleFunc("/hierarchy", service.GetHierarchy).Methods("GET")
	r.HandleFunc("/update/{employeeID}/{managerID}", service.UpdateManager).Methods("PUT")
	r.HandleFunc("/employees", service.AddEmployee).Methods("POST")

	// Configurar opciones CORS
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// Habilitar CORS para todas las rutas
	corsEnabledHandler := handlers.CORS(headersOk, originsOk, methodsOk)(r)

	// Iniciar el servidor en el puerto 8080 con soporte CORS
	fmt.Println("Server listening on port :8080")
	http.ListenAndServe(":8080", corsEnabledHandler)
}
