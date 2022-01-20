package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dto "github.com/tupt0101/student-api-service/src/domain/students"
	"github.com/tupt0101/student-api-service/src/oauth"
	"github.com/tupt0101/student-api-service/src/services/students"
	"github.com/tupt0101/student-api-service/src/utils/errors"
)

func Get(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err)
		return
	}

	students, err := students.StudentsService.GetStudents()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	var result dto.Students = students
	c.JSON(http.StatusOK, result.Marshal(true))
}

func Create(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err)
		return
	}

	var student dto.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := students.StudentsService.CreateStudent(student)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result.Marshal(true))
}

func GetById(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err)
		return
	}

	studentId := c.Param("student_id")
	if studentId == "" {
		c.JSON(http.StatusBadRequest, "id is not valid")
		return
	}

	student, getErr := students.StudentsService.GetStudent(studentId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, student.Marshal(true))
}

func Update(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err)
		return
	}

	studentId := c.Param("student_id")
	if studentId == "" {
		c.JSON(http.StatusBadRequest, "id is not valid")
		return
	}

	var student dto.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	student.StudentId = studentId
	result, err := students.StudentsService.UpdateStudent(student)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusNoContent, result.Marshal(true))
}
