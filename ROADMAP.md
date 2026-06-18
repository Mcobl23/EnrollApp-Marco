# EnrollApp Roadmap

This document collects the recommended next steps for continuing the development of **EnrollApp**. It is kept separate from `README.md` so that the latter stays focused on installation and usage, while work in progress or planned is documented here.

---

## 1. Design the Student Model

**Location:** `Backend/internal/models/`

Create a `student.go` file to define the data structure:

```go
package models

type Student struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
    Code  string `json:"code"`
}
```

---

## 2. Implement Handlers and Routes

**Location:** `Backend/internal/handlers/` and `Backend/internal/routes/`

Create CRUD endpoints to manage students:

- `GET /api/students` — List students
- `POST /api/students` — Create a student
- `GET /api/students/{id}` — Get a specific student
- `PUT /api/students/{id}` — Update a student
- `DELETE /api/students/{id}` — Delete a student

---

## 3. Frontend-Backend Connection

Use `fetch` or `axios` on the frontend to query the backend data at `http://localhost:8080/api/students`.

Don't forget to configure **CORS** on the Go backend to allow requests from the frontend's origin (`http://localhost:5173`).

---

## 4. Other Suggested Improvements (Future)

- Add unit and integration tests for both Go (`testing`) and the frontend (Vitest/Jest).
- Set up a CI pipeline (GitHub Actions) that runs build and tests on every Pull Request.
- Document the full data model (Student, Enrollment, Courses) as it gets implemented.
- Evaluate authentication and authorization before exposing the API publicly.