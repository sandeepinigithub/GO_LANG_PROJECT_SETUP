package service

import (
	"devsMailGo/models"
	"devsMailGo/repository"
	"devsMailGo/api/dto"
)

type WblistService struct{}

func (s *WblistService) ListWblist() ([]dto.WblistResponse, error) {
	entries, err := repository.GetAllWblist()
	if err != nil {
		return nil, err
	}
	var resp []dto.WblistResponse
	for _, e := range entries {
		resp = append(resp, dto.WblistResponse{Rid: e.Rid, Sid: e.Sid, Wb: e.Wb})
	}
	return resp, nil
}

func (s *WblistService) GetWblistByRid(rid uint64) (*dto.WblistResponse, error) {
	entry, err := repository.GetWblistByRid(rid)
	if err != nil {
		return nil, err
	}
	resp := &dto.WblistResponse{Rid: entry.Rid, Sid: entry.Sid, Wb: entry.Wb}
	return resp, nil
}

func (s *WblistService) CreateWblistDTO(req dto.WblistRequest) (*dto.WblistResponse, error) {
	entry := models.Wblist{
		Sid: req.Sid,
		Wb:  req.Wb,
	}
	if err := repository.CreateWblist(&entry); err != nil {
		return nil, err
	}
	return &dto.WblistResponse{Rid: entry.Rid, Sid: entry.Sid, Wb: entry.Wb}, nil
}

func (s *WblistService) UpdateWblistDTO(rid uint64, req dto.WblistRequest) (*dto.WblistResponse, error) {
	entry, err := repository.GetWblistByRid(rid)
	if err != nil {
		return nil, err
	}
	if req.Sid != 0 {
		entry.Sid = req.Sid
	}
	if req.Wb != "" {
		entry.Wb = req.Wb
	}
	if err := repository.UpdateWblist(rid, &entry); err != nil {
		return nil, err
	}
	return &dto.WblistResponse{Rid: entry.Rid, Sid: entry.Sid, Wb: entry.Wb}, nil
}

func (s *WblistService) DeleteWblist(rid uint64) error {
	return repository.DeleteWblist(rid)
} 