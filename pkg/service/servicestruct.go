package service

import "quote/pkg/repo"

type InfoEventService struct {
	SQLite3Repo     repo.IRepo
	InfoRepo        repo.IInfoRepo
	EventDetailRepo repo.IEventDetailRepo
}

func NewInfoEventService(sqlite3DB repo.SQLite3Repo) InfoEventService {
	return InfoEventService{
		SQLite3Repo:     sqlite3DB,
		InfoRepo:        sqlite3DB,
		EventDetailRepo: sqlite3DB,
	}
}
