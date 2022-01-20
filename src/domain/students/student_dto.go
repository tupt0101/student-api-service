package students

import (
	"strings"

	"github.com/tupt0101/student-api-service/src/utils/errors"
)

const (
	StatusActive = "active"
)

type Student struct {
	StudentId string `json:"student_id"`
	Name      string `json:"name"`
	Class     string `json:"class"`
	CreatedOn string `json:"created_on"`
	Status    string `json:"status"`
}

type Students []Student

func (student *Student) Validate() *errors.RestErr {
	student.StudentId = strings.TrimSpace(student.StudentId)
	if student.StudentId == "" {
		return errors.NewBadRequestError("invalid student id")
	}

	student.Name = strings.TrimSpace(student.Name)
	if student.Name == "" {
		return errors.NewBadRequestError("invalid student name")
	}

	student.Class = strings.TrimSpace(student.Class)
	if student.Class == "" {
		return errors.NewBadRequestError("invalid student class")
	}

	return nil
}
