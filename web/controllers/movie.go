package controllers

import (
	"github.com/SarletSKY/go-iris-e-commerce/repositories"
	"github.com/SarletSKY/go-iris-e-commerce/services"
	"github.com/kataras/iris/mvc"
)

type MovieController struct {
}

func (c *MovieController) Get() mvc.View {
	movieRepo := repositories.NewMovieManager()
	movieService := services.NewMovieServiceManager(movieRepo)
	movieResult := movieService.ShowMovieName()
	return mvc.View{
		Name: "movie/index.html",
		Data: movieResult,
	}
}
