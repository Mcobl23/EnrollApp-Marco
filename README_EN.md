# EnrollApp

EnrollApp is a modern web application designed for student record and enrollment management. The project is structured as a monorepo integrating a high-performance **Go (Golang) Backend** and a fast, reactive **React + TypeScript + Vite Frontend**, using **PostgreSQL** as the database.

---

## Project Architecture

The project is divided into two main components:

```
EnrollApp/
├── Backend/          # REST API server and business logic in Go
└── Frontend/         # Interactive user interface in React + TypeScript
```

### Backend Structure

The backend follows a clean, modular design to facilitate scalability:

- **`cmd/api/main.go`**: Application entry point. Initializes the database connection and starts the HTTP server.
- **`internal/`**: Contains the application's protected internal logic:
  - **`db/DBconnection.go`**: Manages the PostgreSQL connection using `pgx/v5` and environment variables.
  - **`server/server.go`**: Configures the HTTP server and defines initial global routes (such as `/health`).
  - **`models/`**: Data structure definitions (e.g., Student, Enrollment, Courses).
  - **`handlers/`**: Controllers that handle HTTP requests and format responses.
  - **`services/`**: Core business logic and system rules.
  - **`routes/`**: API-specific routes.

### Frontend Structure

Developed using the latest web technologies:

- **React 19** & **TypeScript** for safe development and interactive components.
- **Vite 8** as the bundler for instant load times and HMR (Hot Module Replacement).
- **Clean structure** in `src/` with `App.tsx` as the main component and global styles in `App.css`/`index.css`.

---

## Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://go.dev/dl/) (Version 1.26 or higher)
- [Node.js](https://nodejs.org/) (Version 18 or higher) and npm
- [PostgreSQL](https://www.postgresql.org/) (Running local or remote service)

---

## ⚙️ System Configuration

### 1. Database & Environment Variables

The Backend uses a `.env` file to connect to PostgreSQL, which shouldn't be uploated to the repository. Create or edit the `Backend/.env` file with your local database configuration:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=your_database
```

> [!IMPORTANT]
> Make sure you have created the `your_database` database in your PostgreSQL instance before starting the backend.

---

## Execution Guide

Follow these steps to run the application in development mode:

### Step 1: Start the Backend (Go)

1. Navigate to the Backend directory:
   ```bash
   cd Backend
   ```
2. Download the Go module dependencies:
   ```bash
   go mod download
   ```
3. Run the server using the `Makefile` or the direct command:
   * **Using Makefile:**
     ```bash
     make run
     ```
   * **Direct command:**
     ```bash
     go run ./cmd/api
     ```
4. The server will start at `http://localhost:8080`. You can verify its status at:
   ```bash
   curl http://localhost:8080/health
   ```
   *(Should respond with `ok`)*

### Step 2: Start the Frontend (React)

1. Open a new terminal and navigate to the Frontend directory:
   ```bash
   cd Frontend
   ```
2. Install the required dependencies:
   ```bash
   npm install
   ```
3. Start the Vite development server:
   ```bash
   npm run dev
   ```
4. Open your browser at the URL shown in the terminal (usually `http://localhost:5173`).

---

## Backend API Routes

Currently, the backend exposes the following endpoints:

| Method | Endpoint | Description | Status |
| :--- | :--- | :--- | :--- |
| **GET** | `/health` | Server health status check | Active |

---

## Roadmap

The next planned steps for development are documented separately in ROADMAP.md, to keep this README focused on the installation and use of the project.
