# EnrollApp 

EnrollApp es una aplicación web moderna diseñada para la gestión y matrícula de estudiantes. El proyecto está estructurado como un monorrepósito que integra un **Backend en Go (Golang)** de alto rendimiento y un **Frontend en React + TypeScript + Vite** rápido y reactivo, utilizando **PostgreSQL** como base de datos.

---

## Arquitectura del Proyecto

El proyecto se divide en dos componentes principales:

```
EnrollApp/
├── Backend/          # Servidor API REST y lógica de negocio en Go
└── Frontend/         # Interfaz de usuario interactiva en React + TypeScript
```

### Estructura del Backend

El backend sigue un diseño modular y limpio para facilitar la escalabilidad:

- **`cmd/api/main.go`**: Punto de entrada de la aplicación. Inicializa la conexión a la base de datos y arranca el servidor HTTP.
- **`internal/`**: Contiene la lógica interna protegida de la aplicación:
  - **`db/DBconnection.go`**: Gestor de la conexión a PostgreSQL mediante `pgx/v5` y variables de entorno.
  - **`server/server.go`**: Configuración del servidor HTTP y definición inicial de rutas globales (como `/health`).
  - **`models/`**: Definición de las estructuras de datos (ej. Estudiante, Matrícula, Cursos).
  - **`handlers/`**: Controladores que reciben las peticiones HTTP y formatean las respuestas.
  - **`services/`**: Lógica de negocio principal y reglas del sistema.
  - **`routes/`**: Rutas específicas del API.

### Estructura del Frontend

Desarrollado con las tecnologías web más modernas:

- **React 19** & **TypeScript** para un desarrollo seguro y componentes interactivos.
- **Vite 8** como empaquetador para una velocidad de carga y HMR (Hot Module Replacement) instantáneos.
- **Estructura limpia** en `src/` con `App.tsx` como componente principal y estilos globales en `App.css`/`index.css`.

---

## Requisitos Previos

Antes de comenzar, asegúrate de tener instalado:

- [Go](https://go.dev/dl/) (Versión 1.26 o superior)
- [Node.js](https://nodejs.org/) (Versión 18 o superior) y npm
- [PostgreSQL](https://www.postgresql.org/) (Servicio activo localmente o remoto)

---

## Configuración del Sistema

### 1. Base de Datos & Variables de Entorno

El Backend utiliza un archivo `.env` para conectarse a PostgreSQL, el cual no debe subirse al repositorio. Crea o edita el archivo `Backend/.env` con la configuración de tu base de datos local:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=your_database
```

> [!IMPORTANT]
> Asegúrate de haber creado la base de datos `your_database` en tu instancia de PostgreSQL antes de iniciar el backend.

---

## Guía de Ejecución

Sigue estos pasos para levantar la aplicación en modo desarrollo:

### Paso 1: Levantar el Backend (Go)

1. Navega al directorio del Backend:
   ```bash
   cd Backend
   ```
2. Instala las dependencias del módulo Go:
   ```bash
   go mod download
   ```
3. Ejecuta el servidor utilizando el archivo `Makefile` o el comando directo:
   * **Con Makefile:**
     ```bash
     make run
     ```
   * **Comando directo:**
     ```bash
     go run ./cmd/api
     ```
4. El servidor se iniciará en `http://localhost:8080`. Puedes verificar su estado en:
   ```bash
   curl http://localhost:8080/health
   ```
   *(Debería responder `ok`)*

### Paso 2: Levantar el Frontend (React)

1. Abre una nueva terminal y navega al directorio del Frontend:
   ```bash
   cd Frontend
   ```
2. Instala las dependencias necesarias:
   ```bash
   npm install
   ```
3. Inicia el servidor de desarrollo de Vite:
   ```bash
   npm run dev
   ```
4. Abre tu navegador en la URL indicada por la terminal (usualmente `http://localhost:5173`).

---

## Rutas del API del Backend

Actualmente, el backend expone los siguientes endpoints:

| Método | Endpoint | Descripción | Estado |
| :--- | :--- | :--- | :--- |
| **GET** | `/health` | Verificación del estado del servidor | Activo |

---

## Roadmap

Los próximos pasos planeados para el desarrollo se documentan por separado en ROADMAP.md, para mantener este README enfocado en la instalación y el uso del proyecto.
