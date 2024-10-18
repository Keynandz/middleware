package dto

import (
	model "go-authorization/model/entity"
	res "go-authorization/pkg/util/response"
)

// request
type (
	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

// response
type (
	RoleResponse struct {
		ID       int    `json:"id"`
		RoleName string `json:"role_name"`
	}
	TokenResponse struct {
		JWT string `json:"jwt"`
	}
	UserResponse struct {
		model.MasterUserModel
		RoleResponse
		Token TokenResponse `json:"token"`
	}
	ListUserResponse struct {
		model.MasterUserModel
	}
	UserResponseDoc struct {
		Body struct {
			Meta  res.Meta      `json:"meta"`
			Data  UserResponse  `json:"data"`
		} `json:"body"`
	}
)
