package controllers

import (
	"github.com/SarletSKY/go-iris-e-commerce/repositories"
	"github.com/SarletSKY/go-iris-e-commerce/services"
	"github.com/kataras/iris/v12/mvc"
)

type StudentController struct {
}

func (c *StudentController) Get() mvc.View {
	studentRepo := repositories.NewStudentManager()
	studentService := services.NewStudentServiceManager(studentRepo)
	result := studentService.ShowStudentName()
	return mvc.View{
		Name: "student/student.html",
		Data: result,
	}
}
