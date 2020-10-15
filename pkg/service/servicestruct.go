package service

import "quote/pkg/repo"

type QuoteService struct {
	SQLite3Repo     repo.IRepo
	InfoRepo        repo.IInfoRepo
	EventDetailRepo repo.IEventDetailRepo
}

func NewQuoteService(sqlite3DB repo.SQLite3Repo) QuoteService {
	return QuoteService{
		SQLite3Repo:     sqlite3DB,
		InfoRepo:        sqlite3DB,
		EventDetailRepo: sqlite3DB,
	}
}
