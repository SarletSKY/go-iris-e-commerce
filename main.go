package main

import (
	"github.com/SarletSKY/go-iris-e-commerce/web/controllers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.RegisterView(iris.HTML("./web/views", ".html"))
	mvc.New(app.Party("/hello")).Handle(new(controllers.MovieController))
	app.Run(iris.Addr("localhost:8080"))
}
