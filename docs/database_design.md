# Data Model & Database Design

## Overview
The database schema is the foundation of the EnrollApp application. It is designed to support academic administrative processes, including career programs, course catalogs, class sections (groups), classroom allocation, professor assignments, student registration, and academic periods — all structured cleanly to ensure integrity and consistency.

## Entity Relationship Diagram

```
┌──────────────────────┐       ┌──────────────────────┐       ┌──────────────────────┐
│     carreers_tb      │       │      courses_tb      │       │     classrooms_tb    │
├──────────────────────┤       ├──────────────────────┤       ├──────────────────────┤
│ id_carreer (PK)      │──1:N──│ id_course (PK)       │       │ id_classroom (PK)    │
│ carreer_name         │       │ id_carreer (FK)      │       │ classroom_code       │
│ carreer_code         │       │ course_code          │       │ capacity             │
└──────────┬───────────┘       │ course_name          │       └──────────┬───────────┘
           │                   │ credits              │                  │
           │ 1:N               │ quarter              │                  │ 1:N
           │                   └──────────┬───────────┘                  │
┌──────────▼───────────┐                  │                              │
│     students_tb      │                  │ 1:N                          │
├──────────────────────┤                  │                              │
│ id_student (PK)      │                  │        ┌─────────────────────┼─────────────────────┐
│ id_carreer (FK)      │                  │        │                     │                     │
│ identification       │                  │        │                     │                     │
│ first_name           │                  │        │                     │                     │
│ last_name1           │                  │        │                     │                     │
│ last_name2           │                  │        │                     │                     │
│ email                │                  │        │                     │                     │
│ phone                │                  │        │                     │                     │
│ registration_date    │                  │        │                     │                     │
└──────────┬───────────┘                  │        │                     │                     │
           │                              │        │                     │                     │
           │ 1:N                          │        │                     │                     │
           │                     ┌────────▼────────▼────────┐            │                     │
           │                     │        groups_tb         │◄───────────┘                     │
           │                     ├──────────────────────────┤                                  │
           │                     │ id_group (PK)            │                                  │
           │                     │ id_course (FK)           │                                  │
           │                     │ id_professor (FK)        │◄───────────────────────────────┐ │
           │                     │ id_classroom (FK)        │                                │ │
           │                     │ id_period (FK)           │◄─────────────────────────────┐ │ │
           │                     │ group_number             │                              │ │ │
           │                     │ max_capacity             │                              │ │ │
           │                     │ current_capacity         │                              │ │ │
           │                     └────────┬─────────────────┘                              │ │ │
           │                              │                                                │ │ │
           │                              │ 1:N                                            │ │ │
┌──────────▼───────────┐                  │                                                │ │ │
│      enroll_tb       │◄─────────────────┘                                                │ │ │
├──────────────────────┤                                                                   │ │ │
│ id_enroll (PK)       │                                                                   │ │ │
│ id_student (FK)      │                                                                   │ │ │
│ id_group (FK)        │                                                                   │ │ │
│ enroll_date          │                                                                   │ │ │
│ enroll_state         │                                                                   │ │ │
└──────────────────────┘                                                                   │ │ │
                                                                                           │ │ │
┌──────────────────────┐                                                                   │ │ │
│     professors_tb    │───────────────────────────────────────────────────────────────────┘ │ │
├──────────────────────┤                                                                     │ │
│ id_professor (PK)    │                                                                     │ │
│ identification       │                                                                     │ │
│ firt_name            │                                                                     │ │
│ last_name1           │                                                                     │ │
│ last_name2           │                                                                     │ │
│ email                │                                                                     │ │
│ phone                │                                                                     │ │
└──────────────────────┘                                                                     │ │
                                                                                             │ │
┌──────────────────────┐                                                                     │ │
│      periods_tb      │─────────────────────────────────────────────────────────────────────┘ │
├──────────────────────┤                                                                       │
│ id_period (PK)       │                                                                       │
│ period_name          │                                                                       │
│ start_date           │                                                                       │
│ end_date             │                                                                       │
└──────────────────────┘                                                                       │
                                                                                               │
┌──────────────────────┐                                                                       │
│       users_tb       │                                                                       │
├──────────────────────┤                                                                       │
│ id_user (PK)         │                                                                       │
│ username             │                                                                       │
│ password             │                                                                       │
│ role                 │                                                                       │
│ id_student (FK)      │◄──────────────────────────────────────────────────────────────────────┘
│ id_professor (FK)    │
└──────────────────────┘
```

---

## Tables Dictionary

### 1. `carreers_tb`
Stores professional degree plans and major information.
* **`id_carreer`**: `integer` (PK) - Unique career identifier.
* **`carreer_name`**: `character varying` - Full name of the career program.
* **`carreer_code`**: `character varying` - Unique academic code assigned to the career.

### 2. `students_tb`
Registers student profile data.
* **`id_student`**: `integer` (PK) - Unique student identifier.
* **`id_carreer`**: `integer` (FK) - Career program the student is enrolled in. Refers to `carreers_tb(id_carreer)`.
* **`identification`**: `character varying` (Not Null) - Legal identification number.
* **`first_name`**: `character varying` (Not Null) - First name.
* **`last_name1`**: `character varying` (Not Null) - First last name.
* **`last_name2`**: `character varying` - Second last name.
* **`email`**: `character varying` (Not Null) - Student's email address.
* **`phone`**: `character varying` - Contact telephone number.
* **`registration_date`**: `timestamp with time zone` - Date of registration.

### 3. `professors_tb`
Stores faculty members' credentials and info.
* **`id_professor`**: `integer` (PK) - Unique professor identifier.
* **`identification`**: `character varying` (Not Null) - Legal identification number.
* **`firt_name`**: `character varying` (Not Null) - First name (mapped from schema's spelling).
* **`last_name1`**: `character varying` (Not Null) - First last name.
* **`last_name2`**: `character varying` - Second last name.
* **`email`**: `character varying` (Not Null) - Email address.
* **`phone`**: `character varying` - Contact telephone number.

### 4. `courses_tb`
Catalog of courses in the curriculum.
* **`id_course`**: `integer` (PK) - Unique course identifier.
* **`id_carreer`**: `integer` (FK) - Career curriculum this course belongs to. Refers to `carreers_tb(id_carreer)`.
* **`course_code`**: `integer` (Not Null) - Numeric course system code.
* **`course_name`**: `character varying` (Not Null) - Name of the course.
* **`credits`**: `integer` (Not Null) - Number of credits awarded.
* **`quarter`**: `integer` (Not Null) - Academic quarter number.

### 5. `classrooms_tb`
Physical space allocation table.
* **`id_classroom`**: `integer` (PK) - Unique classroom identifier.
* **`classroom_code`**: `character varying` (Not Null) - Classroom tag/code.
* **`capacity`**: `integer` (Not Null) - Maximum seat capacity.

### 6. `periods_tb`
Terms and quarters calendar.
* **`id_period`**: `integer` (PK) - Unique period identifier.
* **`period_name`**: `character varying` (Not Null) - Name of the semester/quarter.
* **`start_date`**: `date` (Not Null) - Beginning of the term.
* **`end_date`**: `date` (Not Null) - End of the term.

### 7. `groups_tb`
Class sections scheduled for a course.
* **`id_group`**: `integer` (PK) - Unique section/group identifier.
* **`id_course`**: `integer` (FK) - Subject being taught. Refers to `courses_tb(id_course)`.
* **`id_professor`**: `integer` (FK) - Assigned teacher. Refers to `professors_tb(id_professor)`.
* **`id_classroom`**: `integer` (FK) - Classroom location. Refers to `classrooms_tb(id_classroom)`.
* **`id_period`**: `integer` (FK) - Active academic term. Refers to `periods_tb(id_period)`.
* **`group_number`**: `integer` (Not Null) - Section index number (e.g. Group 1, Group 2).
* **`max_capacity`**: `integer` (Not Null) - Class limit cap.
* **`current_capacity`**: `integer` - Count of enrolled students.

### 8. `enroll_tb`
The transactional table mapping students to their class sections.
* **`id_enroll`**: `integer` (PK) - Unique enrollment transaction identifier.
* **`id_student`**: `integer` (FK) - Refers to `students_tb(id_student)`.
* **`id_group`**: `integer` (FK) - Refers to `groups_tb(id_group)`.
* **`enroll_date`**: `timestamp with time zone` - Timestamp of when the student enrolled.
* **`enroll_state`**: `character varying` - Active/Inactive state status.

### 9. `users_tb`
Access control and security credentials.
* **`id_user`**: `integer` (PK) - Unique user identifier.
* **`username`**: `character varying` (Not Null) - Unique username.
* **`password`**: `character varying` (Not Null) - Hashed password string.
* **`role`**: `character varying` (Not Null) - Role permission level (e.g. Admin, Student, Professor).
* **`id_student`**: `integer` (FK, Nullable) - Linked student profile (if applicable).
* **`id_professor`**: `integer` (FK, Nullable) - Linked professor profile (if applicable).