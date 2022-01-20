package students

import (
	"github.com/tupt0101/student-api-service/src/domain/students"
	"github.com/tupt0101/student-api-service/src/utils/date_utils"
	"github.com/tupt0101/student-api-service/src/utils/errors"
)

var (
	StudentsService studentsServiceInterface = &studentsService{}
)

type studentsService struct {
}

type studentsServiceInterface interface {
	GetStudents() ([]students.Student, *errors.RestErr)
	CreateStudent(students.Student) (*students.Student, *errors.RestErr)
	GetStudent() (*students.Student, *errors.RestErr)
	UpdateStudent(students.Student) (*students.Student, *errors.RestErr)
}

func (*studentsService) GetStudents() ([]students.Student, *errors.RestErr) {
	student := &students.Student{}
	results, err := student.GetStudents(students.StatusActive)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (*studentsService) CreateStudent(student students.Student) (*students.Student, *errors.RestErr) {
	if err := student.Validate(); err != nil {
		return nil, err
	}

	student.CreatedOn = date_utils.GetNowDBFormat()
	student.Status = students.StatusActive
	if err := student.Save(); err != nil {
		return nil, err
	}

	return &student, nil

}

func (*studentsService) GetStudent() (*students.Student, *errors.RestErr) {
	result := &students.Student{}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

func (*studentsService) UpdateStudent(student students.Student) (*students.Student, *errors.RestErr) {
	current := &students.Student{StudentId: student.StudentId}
	if err := current.Get(); err != nil {
		return nil, err
	}

	current.Name = student.Name
	current.Class = student.Class

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}
