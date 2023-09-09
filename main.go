package main

import (
	"github.com/SarletSKY/go-iris-e-commerce/web/controllers"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.RegisterView(iris.HTML("./web/views", ".html"))
	mvc.New(app.Party("/hello")).Handle(new(controllers.MovieController))
	mvc.New(app.Party("/student")).Handle(new(controllers.StudentController))
	app.Run(iris.Addr("localhost:8080"))
}
