package service

import (
	"GO_LANG_PROJECT_SETUP/api/dto"
	"GO_LANG_PROJECT_SETUP/repository"
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