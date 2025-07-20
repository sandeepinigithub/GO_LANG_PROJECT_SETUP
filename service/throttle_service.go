package service

import (
	"GO_LANG_PROJECT_SETUP/models"
	"GO_LANG_PROJECT_SETUP/repository"
	"GO_LANG_PROJECT_SETUP/api/dto"
)

type ThrottleService struct{}

func (s *ThrottleService) ListThrottle() ([]dto.ThrottleResponse, error) {
	entries, err := repository.GetAllThrottle()
	if err != nil {
		return nil, err
	}
	var resp []dto.ThrottleResponse
	for _, e := range entries {
		resp = append(resp, dto.ThrottleResponse{
			ID: e.ID, Account: e.Account, Kind: e.Kind, Priority: e.Priority, Period: e.Period, MsgSize: e.MsgSize, MaxMsgs: e.MaxMsgs, MaxQuota: e.MaxQuota, MaxRcpts: e.MaxRcpts,
		})
	}
	return resp, nil
}

func (s *ThrottleService) GetThrottleByID(id uint64) (*dto.ThrottleResponse, error) {
	entry, err := repository.GetThrottleByID(id)
	if err != nil {
		return nil, err
	}
	resp := &dto.ThrottleResponse{
		ID: entry.ID, Account: entry.Account, Kind: entry.Kind, Priority: entry.Priority, Period: entry.Period, MsgSize: entry.MsgSize, MaxMsgs: entry.MaxMsgs, MaxQuota: entry.MaxQuota, MaxRcpts: entry.MaxRcpts,
	}
	return resp, nil
}

func (s *ThrottleService) CreateThrottleDTO(req dto.ThrottleRequest) (*dto.ThrottleResponse, error) {
	entry := models.Throttle{
		Account:  req.Account,
		Kind:     req.Kind,
		Priority: req.Priority,
		Period:   req.Period,
		MsgSize:  req.MsgSize,
		MaxMsgs:  req.MaxMsgs,
		MaxQuota: req.MaxQuota,
		MaxRcpts: req.MaxRcpts,
	}
	if err := repository.CreateThrottle(&entry); err != nil {
		return nil, err
	}
	return &dto.ThrottleResponse{
		ID: entry.ID, Account: entry.Account, Kind: entry.Kind, Priority: entry.Priority, Period: entry.Period, MsgSize: entry.MsgSize, MaxMsgs: entry.MaxMsgs, MaxQuota: entry.MaxQuota, MaxRcpts: entry.MaxRcpts,
	}, nil
}

func (s *ThrottleService) UpdateThrottleDTO(id uint64, req dto.ThrottleRequest) (*dto.ThrottleResponse, error) {
	entry, err := repository.GetThrottleByID(id)
	if err != nil {
		return nil, err
	}
	if req.Account != "" {
		entry.Account = req.Account
	}
	if req.Kind != "" {
		entry.Kind = req.Kind
	}
	if req.Priority != 0 {
		entry.Priority = req.Priority
	}
	if req.Period != 0 {
		entry.Period = req.Period
	}
	if req.MsgSize != 0 {
		entry.MsgSize = req.MsgSize
	}
	if req.MaxMsgs != 0 {
		entry.MaxMsgs = req.MaxMsgs
	}
	if req.MaxQuota != 0 {
		entry.MaxQuota = req.MaxQuota
	}
	if req.MaxRcpts != 0 {
		entry.MaxRcpts = req.MaxRcpts
	}
	if err := repository.UpdateThrottle(id, &entry); err != nil {
		return nil, err
	}
	return &dto.ThrottleResponse{
		ID: entry.ID, Account: entry.Account, Kind: entry.Kind, Priority: entry.Priority, Period: entry.Period, MsgSize: entry.MsgSize, MaxMsgs: entry.MaxMsgs, MaxQuota: entry.MaxQuota, MaxRcpts: entry.MaxRcpts,
	}, nil
}

func (s *ThrottleService) DeleteThrottle(id uint64) error {
	return repository.DeleteThrottle(id)
} 