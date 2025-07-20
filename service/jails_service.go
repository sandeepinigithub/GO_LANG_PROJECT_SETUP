package service

import (
	"GO_LANG_PROJECT_SETUP/api/dto"
	"GO_LANG_PROJECT_SETUP/repository"
)

type JailsService struct{}

func (s *JailsService) ListJails() ([]dto.JailResponse, error) {
	jails, err := repository.GetAllJails()
	if err != nil {
		return nil, err
	}
	var resp []dto.JailResponse
	for range jails {
		resp = append(resp, dto.JailResponse{})
	}
	return resp, nil
} 