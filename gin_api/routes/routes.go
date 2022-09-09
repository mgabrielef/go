package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mgabrielef/GOLANG/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/:name", controllers.Greeting)
	r.GET("/students", controllers.ReturnStudents)
	r.POST("/students", controllers.CreateStudent)
	r.GET("/students/:id", controllers.ReturnStudentByID)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.PATCH("/students/:id", controllers.EditStudent)
	r.GET("/students/cpf/:cpf", controllers.SearchStudentByCPF)
	r.Run()
}
