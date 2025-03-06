package usecase

import (
	"database/sql"
	db_errors "remote_db/errors"

	_ "github.com/lib/pq"
)

type Student struct {
	ID     int
	Name   string
	Email  string
	Course string
}

func GetStudentsByCourse(db *sql.DB, course string) ([]Student, error) {
	rows, err := db.Query("SELECT * FROM get_students_by_course($1)", course)
	if err != nil {
		return nil, db_errors.ErrGetRecord
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var s Student
		if err := rows.Scan(&s.ID, &s.Name, &s.Email, &s.Course); err != nil {
			return nil, db_errors.ErrScanRow
		}
		students = append(students, s)
	}

	if err := rows.Err(); err != nil {
		return nil, db_errors.ErrIteration
	}

	return students, nil
}

func AddStudent(db *sql.DB, name, email, course string) error {
	_, err := db.Exec("CALL add_student($1, $2, $3)", name, email, course)
	if err != nil {
		return db_errors.ErrAddNewRow
	}
	return nil
}

func UpdateStudentEmail(db *sql.DB, studentName, newEmail string) error {
	_, err := db.Exec("CALL update_student_email($1, $2)", studentName, newEmail)
	if err != nil {
		return db_errors.ErrUpdateRow
	}
	return nil
}

func DeleteStudentByName(db *sql.DB, studentName string) error {
	_, err := db.Exec("CALL delete_student_by_name($1)", studentName)
	if err != nil {
		return db_errors.ErrDeleteRow
	}
	return nil
}

func TruncateStudentsTable(db *sql.DB) error {
	_, err := db.Exec("CALL truncate_students_table()")
	if err != nil {
		return db_errors.ErrTruncateTable
	}
	return nil
}

func CreateStudentsTable(db *sql.DB) error {
	_, err := db.Exec("CALL create_students_table()")
	if err != nil {
		return db_errors.ErrCreateTable
	}
	return nil
}

func DropStudentsTable(db *sql.DB) error {
	_, err := db.Exec("CALL drop_students_table()")
	if err != nil {
		return db_errors.ErrDropTable
	}
	return nil
}
