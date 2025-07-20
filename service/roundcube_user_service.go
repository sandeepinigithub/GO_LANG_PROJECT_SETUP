package service

import (
	"devsMailGo/models"
	"devsMailGo/repository"
	"devsMailGo/api/dto"
)

type RoundcubeUserService struct{}

func (s *RoundcubeUserService) ListRoundcubeUsers() ([]dto.RoundcubeUserResponse, error) {
	users, err := repository.GetAllRoundcubeUsers()
	if err != nil {
		return nil, err
	}
	var resp []dto.RoundcubeUserResponse
	for _, u := range users {
		resp = append(resp, dto.RoundcubeUserResponse{
			UserID: u.UserID, Username: u.Username, MailHost: u.MailHost, Created: u.Created, LastLogin: u.LastLogin, FailedLogin: u.FailedLogin, FailedLoginCounter: u.FailedLoginCounter, Language: u.Language, Preferences: u.Preferences,
		})
	}
	return resp, nil
}

func (s *RoundcubeUserService) GetRoundcubeUserByID(id uint64) (*dto.RoundcubeUserResponse, error) {
	user, err := repository.GetRoundcubeUserByID(id)
	if err != nil {
		return nil, err
	}
	resp := &dto.RoundcubeUserResponse{
		UserID: user.UserID, Username: user.Username, MailHost: user.MailHost, Created: user.Created, LastLogin: user.LastLogin, FailedLogin: user.FailedLogin, FailedLoginCounter: user.FailedLoginCounter, Language: user.Language, Preferences: user.Preferences,
	}
	return resp, nil
}

func (s *RoundcubeUserService) CreateRoundcubeUserDTO(req dto.RoundcubeUserRequest) (*dto.RoundcubeUserResponse, error) {
	user := models.RoundcubeUser{
		Username:           req.Username,
		MailHost:           req.MailHost,
		Created:            req.Created,
		LastLogin:          req.LastLogin,
		FailedLogin:        req.FailedLogin,
		FailedLoginCounter: req.FailedLoginCounter,
		Language:           req.Language,
		Preferences:        req.Preferences,
	}
	if err := repository.CreateRoundcubeUser(&user); err != nil {
		return nil, err
	}
	return &dto.RoundcubeUserResponse{
		UserID: user.UserID, Username: user.Username, MailHost: user.MailHost, Created: user.Created, LastLogin: user.LastLogin, FailedLogin: user.FailedLogin, FailedLoginCounter: user.FailedLoginCounter, Language: user.Language, Preferences: user.Preferences,
	}, nil
}

func (s *RoundcubeUserService) UpdateRoundcubeUserDTO(id uint64, req dto.RoundcubeUserRequest) (*dto.RoundcubeUserResponse, error) {
	user, err := repository.GetRoundcubeUserByID(id)
	if err != nil {
		return nil, err
	}
	if req.Username != "" {
		user.Username = req.Username
	}
	if req.MailHost != "" {
		user.MailHost = req.MailHost
	}
	if req.Created != "" {
		user.Created = req.Created
	}
	if req.LastLogin != nil {
		user.LastLogin = req.LastLogin
	}
	if req.FailedLogin != nil {
		user.FailedLogin = req.FailedLogin
	}
	if req.FailedLoginCounter != nil {
		user.FailedLoginCounter = req.FailedLoginCounter
	}
	if req.Language != nil {
		user.Language = req.Language
	}
	if req.Preferences != nil {
		user.Preferences = req.Preferences
	}
	if err := repository.UpdateRoundcubeUser(id, &user); err != nil {
		return nil, err
	}
	return &dto.RoundcubeUserResponse{
		UserID: user.UserID, Username: user.Username, MailHost: user.MailHost, Created: user.Created, LastLogin: user.LastLogin, FailedLogin: user.FailedLogin, FailedLoginCounter: user.FailedLoginCounter, Language: user.Language, Preferences: user.Preferences,
	}, nil
}

func (s *RoundcubeUserService) DeleteRoundcubeUser(id uint64) error {
	return repository.DeleteRoundcubeUser(id)
} 