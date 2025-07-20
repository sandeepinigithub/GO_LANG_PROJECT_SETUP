package service

import (
	"GO_LANG_PROJECT_SETUP/models"
	"GO_LANG_PROJECT_SETUP/repository"
	"GO_LANG_PROJECT_SETUP/api/dto"
)

type GreylistingService struct{}

func (s *GreylistingService) ListGreylisting() ([]dto.GreylistingResponse, error) {
	entries, err := repository.GetAllGreylisting()
	if err != nil {
		return nil, err
	}
	var resp []dto.GreylistingResponse
	for _, e := range entries {
		resp = append(resp, dto.GreylistingResponse{
			ID: e.ID, Account: e.Account, Priority: e.Priority, Sender: e.Sender, SenderPriority: e.SenderPriority, Comment: e.Comment, Active: e.Active,
		})
	}
	return resp, nil
}

func (s *GreylistingService) GetGreylistingByID(id uint64) (*dto.GreylistingResponse, error) {
	entry, err := repository.GetGreylistingByID(id)
	if err != nil {
		return nil, err
	}
	resp := &dto.GreylistingResponse{
		ID: entry.ID, Account: entry.Account, Priority: entry.Priority, Sender: entry.Sender, SenderPriority: entry.SenderPriority, Comment: entry.Comment, Active: entry.Active,
	}
	return resp, nil
}

func (s *GreylistingService) CreateGreylistingDTO(req dto.GreylistingRequest) (*dto.GreylistingResponse, error) {
	entry := models.Greylisting{
		Account:        req.Account,
		Priority:       req.Priority,
		Sender:         req.Sender,
		SenderPriority: req.SenderPriority,
		Comment:        req.Comment,
		Active:         req.Active,
	}
	if err := repository.CreateGreylisting(&entry); err != nil {
		return nil, err
	}
	return &dto.GreylistingResponse{
		ID: entry.ID, Account: entry.Account, Priority: entry.Priority, Sender: entry.Sender, SenderPriority: entry.SenderPriority, Comment: entry.Comment, Active: entry.Active,
	}, nil
}

func (s *GreylistingService) UpdateGreylistingDTO(id uint64, req dto.GreylistingRequest) (*dto.GreylistingResponse, error) {
	entry, err := repository.GetGreylistingByID(id)
	if err != nil {
		return nil, err
	}
	if req.Account != "" {
		entry.Account = req.Account
	}
	if req.Priority != 0 {
		entry.Priority = req.Priority
	}
	if req.Sender != "" {
		entry.Sender = req.Sender
	}
	if req.SenderPriority != 0 {
		entry.SenderPriority = req.SenderPriority
	}
	if req.Comment != "" {
		entry.Comment = req.Comment
	}
	entry.Active = req.Active
	if err := repository.UpdateGreylisting(id, &entry); err != nil {
		return nil, err
	}
	return &dto.GreylistingResponse{
		ID: entry.ID, Account: entry.Account, Priority: entry.Priority, Sender: entry.Sender, SenderPriority: entry.SenderPriority, Comment: entry.Comment, Active: entry.Active,
	}, nil
}

func (s *GreylistingService) DeleteGreylisting(id uint64) error {
	return repository.DeleteGreylisting(id)
} 