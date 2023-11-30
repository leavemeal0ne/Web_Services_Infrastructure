package servicelogs

import (
	"lab4/internal/models"
	"lab4/internal/repository"
)

type LogService struct {
	repo repository.MongoDBLog
}

func NewLogService(repo repository.MongoDBLog) *LogService {
	return &LogService{repo: repo}
}

func (l *LogService) InsertLog(data models.LogData) error {
	return l.repo.InsertLog(data)
}
