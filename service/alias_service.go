package service

import (
	"devsMailGo/models"
	"devsMailGo/repository"
	"errors"
	"devsMailGo/api/dto"
)

type AliasService struct{}

func (s *AliasService) ListAliases() ([]dto.AliasResponse, error) {
	aliases, err := repository.GetAllAliases()
	if err != nil {
		return nil, err
	}
	var resp []dto.AliasResponse
	for _, a := range aliases {
		resp = append(resp, dto.AliasResponse{Address: a.Address, Domain: a.Domain, Goto: a.Goto, Active: a.Active})
	}
	return resp, nil
}

func (s *AliasService) GetAliasByAddress(address string) (*dto.AliasResponse, error) {
	alias, err := repository.GetAliasByAddress(address)
	if err != nil {
		return nil, err
	}
	resp := &dto.AliasResponse{Address: alias.Address, Domain: alias.Domain, Goto: alias.Goto, Active: alias.Active}
	return resp, nil
}

func (s *AliasService) CreateAliasDTO(address string, req dto.AliasRequest) (*dto.AliasResponse, error) {
	if address == "" {
		return nil, errors.New("address required")
	}
	alias := models.Alias{
		Address: address,
		Domain:  req.Domain,
		Goto:    req.Goto,
		Active:  req.Active,
	}
	if err := repository.CreateAlias(&alias); err != nil {
		return nil, err
	}
	return &dto.AliasResponse{Address: alias.Address, Domain: alias.Domain, Goto: alias.Goto, Active: alias.Active}, nil
}

func (s *AliasService) UpdateAliasDTO(address string, req dto.AliasRequest) (*dto.AliasResponse, error) {
	alias, err := repository.GetAliasByAddress(address)
	if err != nil {
		return nil, err
	}
	if req.Domain != "" {
		alias.Domain = req.Domain
	}
	if req.Goto != "" {
		alias.Goto = req.Goto
	}
	alias.Active = req.Active
	if err := repository.UpdateAlias(address, &alias); err != nil {
		return nil, err
	}
	return &dto.AliasResponse{Address: alias.Address, Domain: alias.Domain, Goto: alias.Goto, Active: alias.Active}, nil
}

func (s *AliasService) DeleteAlias(address string) error {
	return repository.DeleteAlias(address)
} 