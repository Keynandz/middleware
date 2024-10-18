package service

import (
	"fmt"
	"go-authorization/internal/login/repository"
	roleRepository "go-authorization/internal/role/repository"
	"go-authorization/middleware"
	"go-authorization/model/dto"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type serviceImpl struct {
	DB             *gorm.DB
	Repository     repository.Repository
	RoleRepository roleRepository.Repository
}

func NewService(
	DB *gorm.DB,
	Repository repository.Repository,
	RoleRepository roleRepository.Repository,
) Service {
	return &serviceImpl{
		DB:             DB,
		Repository:     Repository,
		RoleRepository: RoleRepository,
	}
}

func (s *serviceImpl) Get(c echo.Context, payload dto.LoginRequest) (response dto.UserResponse, err error) {
	user, err := s.Repository.Find(c, s.DB, payload)
	if err != nil {
		return response, fmt.Errorf("Invalid Email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return response, fmt.Errorf("Invalid Password")
	}

	role, err := s.RoleRepository.Find(c, s.DB, int(user.RoleID))
	if err != nil {
		return response, err
	}

	user.Password = ""
	user.OauthUid = ""

	token, err := middleware.GenerateToken(user, role.Name)
	if err != nil {
		return response, fmt.Errorf("Failed to generate token: %v", err)
	}

	response = dto.UserResponse{
		MasterUserModel: user,
		Token: dto.TokenResponse{
			JWT: token,
		},
		RoleResponse: dto.RoleResponse{
			ID:       int(role.ID),
			RoleName: role.Name,
		},
	}

	return response, nil
}

func (s *serviceImpl) GetUser(c echo.Context) (response []dto.ListUserResponse, err error) {
	user, err := s.Repository.FindListUser(c, s.DB)
	if err != nil {
		return response, err
	}

	for _, data := range user {
		data.Password = ""
		data.OauthUid = ""
		response = append(response, dto.ListUserResponse{
			MasterUserModel: data,
		})
	}

	return response, nil
}
