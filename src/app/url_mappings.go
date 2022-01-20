package app

import "github.com/tupt0101/student-api-service/src/controller/students"

func mapUrls() {

	router.GET("/api/students/", students.Get)
	router.POST("/api/students/", students.Create)
	router.GET("/api/students/:student_id", students.GetById)
	router.PUT("/api/students/:student_id", students.Update)

}
