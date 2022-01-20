package app

func mapUrls() {

	router.GET("/api/students/")
	router.POST("/api/students/")
	router.GET("/api/students/:student_id")
	router.PUT("/api/students/:student_id")

}
