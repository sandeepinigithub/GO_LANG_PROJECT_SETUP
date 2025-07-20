package service

import (
	"GO_LANG_PROJECT_SETUP/models"
	"GO_LANG_PROJECT_SETUP/repository"
	"GO_LANG_PROJECT_SETUP/utils"
	"GO_LANG_PROJECT_SETUP/api/dto"
	"errors"
)

type UserService struct{}

func (s *UserService) ListUsers() ([]dto.UserResponse, error) {
	users, err := repository.GetAllUsers()
	if err != nil {
		return nil, err
	}
	var resp []dto.UserResponse
	for _, u := range users {
		resp = append(resp, dto.UserResponse{
			ID:     u.ID, 
			Email:  u.Email, 
			Name:   u.Name,
			Domain: u.Domain,
		})
	}
	return resp, nil
}

func (s *UserService) GetUserByID(id uint) (*dto.UserResponse, error) {
	user, err := repository.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	resp := &dto.UserResponse{
		ID:     user.ID, 
		Email:  user.Email, 
		Name:   user.Name,
		Domain: user.Domain,
	}
	return resp, nil
}

func (s *UserService) RegisterUserDTO(req dto.UserRequest) (*dto.UserResponse, error) {
	if req.Password == "" || req.Email == "" {
		return nil, errors.New("email and password required")
	}
	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	user := models.User{
		Email:    req.Email,
		Name:     req.Name,
		Password: hash,
		Domain:   req.Domain,
		Quota:    req.Quota,
		Language: req.Language,
	}
	_, err = repository.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return &dto.UserResponse{
		ID:     user.ID, 
		Email:  user.Email, 
		Name:   user.Name,
		Domain: user.Domain,
	}, nil
}

func (s *UserService) UpdateUserDTO(id uint, req dto.UserRequest) (*dto.UserResponse, error) {
	user, err := repository.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Password != "" {
		hash, err := utils.HashPassword(req.Password)
		if err != nil {
			return nil, err
		}
		user.Password = hash
	}
	if req.Domain != "" {
		user.Domain = req.Domain
	}
	if req.Quota != 0 {
		user.Quota = req.Quota
	}
	if req.Language != "" {
		user.Language = req.Language
	}
	_, err = repository.UpdateUser(id, user)
	if err != nil {
		return nil, err
	}
	return &dto.UserResponse{
		ID:     user.ID, 
		Email:  user.Email, 
		Name:   user.Name,
		Domain: user.Domain,
	}, nil
}

func (s *UserService) DeleteUser(id uint) error {
	return repository.DeleteUser(id)
}

func (s *UserService) AuthenticateUser(email, password string) (*dto.UserResponse, error) {
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	
	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid credentials")
	}
	
	return &dto.UserResponse{
		ID:     user.ID,
		Email:  user.Email,
		Name:   user.Name,
		Domain: user.Domain,
	}, nil
} 