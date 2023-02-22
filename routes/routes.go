package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/victorradael/rest_api_go_gin/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/api/students", controllers.GetAllStudents)
	r.GET("/api/students/:id", controllers.GetOneStudentById)
	r.GET("/api/students/cpf/:cpf", controllers.SearchStudentsByCpf)
	r.POST("/api/students", controllers.CreateNewStudent)
	r.DELETE("/api/students/:id", controllers.DeleteOneStudentById)
	r.PATCH("/api/students/:id", controllers.UpdateOneStudentById)
	r.Run(":8000")
}
