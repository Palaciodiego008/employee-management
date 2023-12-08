# employee-management


A continuación, proporcionaré un README.md para ambos, el backend (escrito en Go) y el frontend (escrito en Vue.js). Este es un ejemplo básico y puedes personalizarlo según tus necesidades.

---

# Employee Management System

Este es un sistema simple de gestión de empleados con un backend escrito en Go y un frontend escrito en Vue.js.

## Backend (Go)

### Requisitos previos

- Go (versión 1.16 o superior)
- MySQL

### Configuración

1. Clona este repositorio.

```bash
git clone <URL_DEL_REPO>
```

2. Cambia al directorio del backend.

```bash
cd backend
```

3. Instala las dependencias.

```bash
go get -u github.com/gorilla/mux
go get -u github.com/gorilla/handlers
go get -u github.com/jmoiron/sqlx
go get -u github.com/go-sql-driver/mysql
```

4. Configura la base de datos.

   - Asegúrate de tener una base de datos MySQL creada.
   - Copia el archivo `.env.example` a `.env` y configura las variables de entorno según tu entorno.

5. Inicia el servidor.

```bash
go run main.go
```

El servidor estará disponible en `http://localhost:8080`.

### Endpoints

- Obtener la jerarquía de empleados: `GET http://localhost:8080/api/employees/hierarchy`
- Actualizar el jefe de un empleado: `PUT http://localhost:8080/api/employees/{employeeID}/managers/{managerID}`
- Agregar un nuevo empleado: `POST http://localhost:8080/api/employees`

## Frontend (Vue.js)

### Requisitos previos

- Node.js (versión 14 o superior)
- npm (viene con Node.js)

### Configuración

1. Cambia al directorio del frontend.

```bash
cd frontend
```

2. Instala las dependencias.

```bash
npm install
```

3. Inicia la aplicación.

```bash
npm run serve
```

La aplicación estará disponible en `http://localhost:8081`. Abre tu navegador y ve a esa dirección.

### Uso

- La página principal muestra la jerarquía de empleados.
- Puedes navegar a "Add Employee" y "Update Manager" desde la página de la jerarquía de empleados.

---

Este README proporciona una guía básica para la instalación y ejecución del sistema. Puedes personalizarlo y agregar más detalles según tus necesidades. Además, asegúrate de mantener tus dependencias actualizadas y sigue las mejores prácticas de seguridad para las configuraciones en producción.
