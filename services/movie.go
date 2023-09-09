package services

import "github.com/SarletSKY/go-iris-e-commerce/repositories"

type MovieService interface {
	ShowMovieName() string
}
type MovieServiceManager struct {
	repo repositories.MovieRepository
}

func NewMovieServiceManager(repo repositories.MovieRepository) MovieService {
	return &MovieServiceManager{repo: repo}
}

func (m *MovieServiceManager) ShowMovieName() string {
	return "我们获取到的视频为：" + m.repo.GetMovieName()
}
