package service

import (
	"devsMailGo/models"
	"devsMailGo/repository"
	"errors"
	"devsMailGo/api/dto"
)

type QuotaService struct{}

func (s *QuotaService) ListQuota() ([]dto.QuotaResponse, error) {
	quotas, err := repository.GetAllUsedQuota()
	if err != nil {
		return nil, err
	}
	var resp []dto.QuotaResponse
	for _, q := range quotas {
		resp = append(resp, dto.QuotaResponse{Username: q.Username, Bytes: q.Bytes, Messages: q.Messages, Domain: q.Domain})
	}
	return resp, nil
}

func (s *QuotaService) GetQuotaByUsername(username string) (*dto.QuotaResponse, error) {
	quota, err := repository.GetUsedQuotaByUsername(username)
	if err != nil {
		return nil, err
	}
	resp := &dto.QuotaResponse{Username: quota.Username, Bytes: quota.Bytes, Messages: quota.Messages, Domain: quota.Domain}
	return resp, nil
}

func (s *QuotaService) CreateQuotaDTO(username string, req dto.QuotaRequest) (*dto.QuotaResponse, error) {
	if username == "" {
		return nil, errors.New("username required")
	}
	quota := models.UsedQuota{
		Username: username,
		Bytes:    req.Bytes,
		Messages: req.Messages,
		Domain:   req.Domain,
	}
	if err := repository.CreateUsedQuota(&quota); err != nil {
		return nil, err
	}
	return &dto.QuotaResponse{Username: quota.Username, Bytes: quota.Bytes, Messages: quota.Messages, Domain: quota.Domain}, nil
}

func (s *QuotaService) UpdateQuotaDTO(username string, req dto.QuotaRequest) (*dto.QuotaResponse, error) {
	quota, err := repository.GetUsedQuotaByUsername(username)
	if err != nil {
		return nil, err
	}
	if req.Bytes != 0 {
		quota.Bytes = req.Bytes
	}
	if req.Messages != 0 {
		quota.Messages = req.Messages
	}
	if req.Domain != "" {
		quota.Domain = req.Domain
	}
	if err := repository.UpdateUsedQuota(username, &quota); err != nil {
		return nil, err
	}
	return &dto.QuotaResponse{Username: quota.Username, Bytes: quota.Bytes, Messages: quota.Messages, Domain: quota.Domain}, nil
}

func (s *QuotaService) DeleteQuota(username string) error {
	return repository.DeleteUsedQuota(username)
} 