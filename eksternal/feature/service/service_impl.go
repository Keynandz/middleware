package service

import (
	"go-base-structure/eksternal/feature/repository"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type serviceImpl struct {
	DB         *gorm.DB
	Repository repository.Repository
}

func NewService(
	DB *gorm.DB,
	Repository repository.Repository,
) Service {
	return &serviceImpl{
		DB:         DB,
		Repository: Repository,
	}
}

func (s *serviceImpl) Get(c echo.Context) (string, error) {
	response, err := s.Repository.Find(c, s.DB)
	if err != nil {
		return response, err
	}

	return response, nil
}
