package students

import (
	"fmt"
	"log"

	"github.com/tupt0101/student-api-service/src/datasource/postgres"
	"github.com/tupt0101/student-api-service/src/utils/errors"
)

const (
	queryGetStudents   = "SELECT student_id, name, class FROM students WHERE status=$1;"
	queryInsertStudent = "INSERT INTO students(student_id, name, class, created_on, status) VALUES($1,$2,$3,$4,$5);"
	queryGetStudent    = "SELECT student_id, name, class FROM students WHERE student_id=$1;"
	queryUpdateStudent = "UPDATE students SET name=$1, class=$2 WHERE student_id=$3;"
)

func (student *Student) GetStudents(status string) ([]Student, *errors.RestErr) {
	stmt, err := postgres.Client.Prepare(queryGetStudents)
	if err != nil {
		log.Println("Error when trying to prepare statement get students by status", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		log.Println("Error when trying to get students by status", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer rows.Close()

	results := make([]Student, 0)
	for rows.Next() {
		var student Student
		if err := rows.Scan(&student.StudentId, &student.Name, &student.Class); err != nil {
			log.Println("Error when trying to scan student row into student struct", err)
			return nil, errors.NewInternalServerError("database error")
		}
		results = append(results, student)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no students matching status %s", status))
	}
	return results, nil
}

func (student *Student) Save() *errors.RestErr {
	stmt, err := postgres.Client.Prepare(queryInsertStudent)
	if err != nil {
		log.Println("Error when trying to prepare statement insert student", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(student.StudentId, student.Name, student.Class, student.CreatedOn, student.Status)
	if err != nil {
		log.Println("Error when trying to save student", err)
		return errors.NewInternalServerError("database error")
	}

	return nil

}

func (student *Student) Get() *errors.RestErr {
	stmt, err := postgres.Client.Prepare(queryGetStudent)
	if err != nil {
		log.Println("Error when trying to prepare statement get student by id", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(student.StudentId)
	if getErr := result.Scan(&student.StudentId, &student.Name, &student.Class); getErr != nil {
		log.Println("Error when trying to get student by id", getErr)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (student *Student) Update() *errors.RestErr {
	stmt, err := postgres.Client.Prepare(queryUpdateStudent)
	if err != nil {
		log.Println("Error when trying to prepare statement update student by id", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(&student.Name, &student.Class, &student.StudentId)
	if err != nil {
		log.Println("Error when trying to update student by id", err)
		return errors.NewInternalServerError("database error")
	}

	return nil
}
