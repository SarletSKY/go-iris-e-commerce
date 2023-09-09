package repositories

import "github.com/SarletSKY/go-iris-e-commerce/models"

type MovieRepository interface {
	GetMovieName() string
}
type MovieManager struct {
}

func NewMovieManager() MovieRepository {
	return &MovieManager{}
}

func (m *MovieManager) GetMovieName() string {
	movies := &models.Movie{Name: "hello world"}
	return movies.Name
}
