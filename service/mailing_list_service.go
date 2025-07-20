package service

import (
	"GO_LANG_PROJECT_SETUP/models"
	"GO_LANG_PROJECT_SETUP/repository"
	"GO_LANG_PROJECT_SETUP/api/dto"
	"errors"
)

type MailingListService struct{}



func (s *MailingListService) ListMailingLists() ([]dto.MailingListResponse, error) {
	lists, err := repository.GetAllMailingLists()
	if err != nil {
		return nil, err
	}
	var resp []dto.MailingListResponse
	for _, l := range lists {
		resp = append(resp, dto.MailingListResponse{ID: uint64(l.ID), Address: l.Address, Domain: l.Domain, Members: l.Members, Type: l.Type, Active: l.Active})
	}
	return resp, nil
}

func (s *MailingListService) GetMailingListByAddress(address string) (*dto.MailingListResponse, error) {
	list, err := repository.GetMailingListByAddress(address)
	if err != nil {
		return nil, err
	}
	resp := &dto.MailingListResponse{ID: uint64(list.ID), Address: list.Address, Domain: list.Domain, Members: list.Members, Type: list.Type, Active: list.Active}
	return resp, nil
}

func (s *MailingListService) CreateMailingListDTO(address string, req dto.MailingListRequest) (*dto.MailingListResponse, error) {
	if address == "" {
		return nil, errors.New("address required")
	}
	list := models.MailingList{
		Address: address,
		Domain:  req.Domain,
		Members: req.Members,
		Type:    req.Type,
		Active:  req.Active,
	}
	if err := repository.CreateMailingList(&list); err != nil {
		return nil, err
	}
	return &dto.MailingListResponse{ID: uint64(list.ID), Address: list.Address, Domain: list.Domain, Members: list.Members, Type: list.Type, Active: list.Active}, nil
}

func (s *MailingListService) UpdateMailingListDTO(address string, req dto.MailingListRequest) (*dto.MailingListResponse, error) {
	list, err := repository.GetMailingListByAddress(address)
	if err != nil {
		return nil, err
	}
	if req.Domain != "" {
		list.Domain = req.Domain
	}
	if req.Members != "" {
		list.Members = req.Members
	}
	if req.Type != "" {
		list.Type = req.Type
	}
	list.Active = req.Active
	if err := repository.UpdateMailingList(address, &list); err != nil {
		return nil, err
	}
	return &dto.MailingListResponse{ID: uint64(list.ID), Address: list.Address, Domain: list.Domain, Members: list.Members, Type: list.Type, Active: list.Active}, nil
}

func (s *MailingListService) DeleteMailingList(address string) error {
	return repository.DeleteMailingList(address)
} 