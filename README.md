# Online School Database Management

This project is a desktop application for managing a student database in an online school.  
It provides user-friendly access to student records with features such as adding, updating, searching, and deleting student data, along with basic role-based access control.

---

## Features

### Authorization
- Login with **username and password**
- Two access levels:
  - **Administrator** — full access to all operations
  - **Guest** — read-only access

### Core Functionality
- Create a table of students
- Add a new student
- Delete the entire table
- Clear all student data
- Get students by course
- Delete a student by name
- Update a student’s email

---

## Technologies Used
- **Go** — core backend and application logic
- **PostgreSQL** — relational database for persistent storage
- **Fyne** — cross-platform GUI library for Go
- **pq** — PostgreSQL driver for Go

---

## User Roles:
- **Administrator**  
  Full access to create, delete, and update student records and tables.

- **Guest**  
  Can only view and search student data without modifying it.

---

## Author

**Maxim Sokolov**  
Student of Computer Science and Technology at HSE University  
This project was developed as a desktop GUI-based pet project using Go and PostgreSQL.
