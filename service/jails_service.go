package service

import (
	"devsMailGo/api/dto"
	"devsMailGo/repository"
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