package service

import (
	"GO_LANG_PROJECT_SETUP/models"
	"GO_LANG_PROJECT_SETUP/repository"
	"errors"
	"GO_LANG_PROJECT_SETUP/api/dto"
)

type DomainService struct{}

func (s *DomainService) ListDomains() ([]dto.DomainResponse, error) {
	domains, err := repository.GetAllDomains()
	if err != nil {
		return nil, err
	}
	var resp []dto.DomainResponse
	for _, d := range domains {
		resp = append(resp, dto.DomainResponse{ Name: d.Name, Description: d.Description, Quota: d.Quota })
	}
	return resp, nil
}

func (s *DomainService) GetDomainByName(name string) (*dto.DomainResponse, error) {
	domain, err := repository.GetDomainByName(name)
	if err != nil {
		return nil, err
	}
	resp := &dto.DomainResponse{ Name: domain.Name, Description: domain.Description, Quota: domain.Quota }
	return resp, nil
}

func (s *DomainService) CreateDomainDTO(name string, req dto.DomainRequest) (*dto.DomainResponse, error) {
	if name == "" {
		return nil, errors.New("domain name required")
	}
	domain := models.Domain{
		Name:            name,
		Description:     req.Description,
		Quota:           req.Quota,
		Language:        req.Language,
		Transport:       req.Transport,
		DefaultQuota:    req.DefaultQuota,
		MaxUserQuota:    req.MaxUserQuota,
		NumberOfUsers:   req.NumberOfUsers,
		NumberOfAliases: req.NumberOfAliases,
		NumberOfLists:   req.NumberOfLists,
		AccountStatus:   req.AccountStatus,
		SenderBcc:       req.SenderBcc,
		RecipientBcc:    req.RecipientBcc,
	}
	if err := repository.CreateDomain(&domain); err != nil {
		return nil, err
	}
	return &dto.DomainResponse{ Name: domain.Name, Description: domain.Description, Quota: domain.Quota }, nil
}

func (s *DomainService) UpdateDomainDTO(name string, req dto.DomainRequest) (*dto.DomainResponse, error) {
	domain, err := repository.GetDomainByName(name)
	if err != nil {
		return nil, err
	}
	if req.Description != "" {
		domain.Description = req.Description
	}
	if req.Quota != 0 {
		domain.Quota = req.Quota
	}
	if req.Language != "" {
		domain.Language = req.Language
	}
	if req.Transport != "" {
		domain.Transport = req.Transport
	}
	if req.DefaultQuota != 0 {
		domain.DefaultQuota = req.DefaultQuota
	}
	if req.MaxUserQuota != 0 {
		domain.MaxUserQuota = req.MaxUserQuota
	}
	if req.NumberOfUsers != 0 {
		domain.NumberOfUsers = req.NumberOfUsers
	}
	if req.NumberOfAliases != 0 {
		domain.NumberOfAliases = req.NumberOfAliases
	}
	if req.NumberOfLists != 0 {
		domain.NumberOfLists = req.NumberOfLists
	}
	if req.AccountStatus != "" {
		domain.AccountStatus = req.AccountStatus
	}
	if req.SenderBcc != "" {
		domain.SenderBcc = req.SenderBcc
	}
	if req.RecipientBcc != "" {
		domain.RecipientBcc = req.RecipientBcc
	}
	if err := repository.UpdateDomain(name, &domain); err != nil {
		return nil, err
	}
	return &dto.DomainResponse{ Name: domain.Name, Description: domain.Description, Quota: domain.Quota }, nil
}

func (s *DomainService) DeleteDomain(name string) error {
	return repository.DeleteDomain(name)
} 