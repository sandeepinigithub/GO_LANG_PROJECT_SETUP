package service

import (
	"devsMailGo/api/dto"
	"devsMailGo/repository"
)

type BannedService struct{}

func (s *BannedService) ListBanned() ([]dto.BannedResponse, error) {
	banned, err := repository.GetAllBanned()
	if err != nil {
		return nil, err
	}
	var resp []dto.BannedResponse
	for _, b := range banned {
		resp = append(resp, dto.BannedResponse{IP: b.IP})
	}
	return resp, nil
}

func (s *BannedService) UnbanByID(id uint64) error {
	return repository.UnbanByID(id)
} 