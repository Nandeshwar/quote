package model

import (
	"time"
)

//go:generate gofp -destination fp.go -pkg model -type "Info, string"
type Info struct {
	ID           int
	Title        string
	Info         string
	Links        []string
	Link         string
	CreationDate time.Time
	UpdatedDate  time.Time
}
