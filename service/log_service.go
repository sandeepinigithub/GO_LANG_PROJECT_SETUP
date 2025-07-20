package service

import (
	"devsMailGo/models"
	"devsMailGo/repository"
	"devsMailGo/api/dto"
)

type LogService struct{}

func (s *LogService) ListLogs() ([]dto.LogResponse, error) {
	logs, err := repository.GetAllLogs()
	if err != nil {
		return nil, err
	}
	var resp []dto.LogResponse
	for _, l := range logs {
		resp = append(resp, dto.LogResponse{
			ID: l.ID, Timestamp: l.Timestamp, Admin: l.Admin, IP: l.IP, Domain: l.Domain, Username: l.Username, Event: l.Event, Loglevel: l.Loglevel, Msg: l.Msg,
		})
	}
	return resp, nil
}

func (s *LogService) GetLogByID(id uint64) (*dto.LogResponse, error) {
	logEntry, err := repository.GetLogByID(id)
	if err != nil {
		return nil, err
	}
	resp := &dto.LogResponse{
		ID: logEntry.ID, Timestamp: logEntry.Timestamp, Admin: logEntry.Admin, IP: logEntry.IP, Domain: logEntry.Domain, Username: logEntry.Username, Event: logEntry.Event, Loglevel: logEntry.Loglevel, Msg: logEntry.Msg,
	}
	return resp, nil
}

func (s *LogService) CreateLogDTO(req dto.LogRequest) (*dto.LogResponse, error) {
	logEntry := models.Log{
		Timestamp: req.Timestamp,
		Admin:     req.Admin,
		IP:        req.IP,
		Domain:    req.Domain,
		Username:  req.Username,
		Event:     req.Event,
		Loglevel:  req.Loglevel,
		Msg:       req.Msg,
	}
	if err := repository.CreateLog(&logEntry); err != nil {
		return nil, err
	}
	return &dto.LogResponse{
		ID: logEntry.ID, Timestamp: logEntry.Timestamp, Admin: logEntry.Admin, IP: logEntry.IP, Domain: logEntry.Domain, Username: logEntry.Username, Event: logEntry.Event, Loglevel: logEntry.Loglevel, Msg: logEntry.Msg,
	}, nil
}

func (s *LogService) UpdateLogDTO(id uint64, req dto.LogRequest) (*dto.LogResponse, error) {
	logEntry, err := repository.GetLogByID(id)
	if err != nil {
		return nil, err
	}
	if req.Timestamp != "" {
		logEntry.Timestamp = req.Timestamp
	}
	if req.Admin != "" {
		logEntry.Admin = req.Admin
	}
	if req.IP != "" {
		logEntry.IP = req.IP
	}
	if req.Domain != "" {
		logEntry.Domain = req.Domain
	}
	if req.Username != "" {
		logEntry.Username = req.Username
	}
	if req.Event != "" {
		logEntry.Event = req.Event
	}
	if req.Loglevel != "" {
		logEntry.Loglevel = req.Loglevel
	}
	if req.Msg != "" {
		logEntry.Msg = req.Msg
	}
	if err := repository.UpdateLog(id, &logEntry); err != nil {
		return nil, err
	}
	return &dto.LogResponse{
		ID: logEntry.ID, Timestamp: logEntry.Timestamp, Admin: logEntry.Admin, IP: logEntry.IP, Domain: logEntry.Domain, Username: logEntry.Username, Event: logEntry.Event, Loglevel: logEntry.Loglevel, Msg: logEntry.Msg,
	}, nil
}

func (s *LogService) DeleteLog(id uint64) error {
	return repository.DeleteLog(id)
} 