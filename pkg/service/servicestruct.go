package service

import "quote/pkg/repo"

type QuoteService struct {
	SQLite3Repo repo.IRepo
}

func NewQuoteService(sqlite3DB repo.IRepo) QuoteService {
	return QuoteService{SQLite3Repo: sqlite3DB}
}
