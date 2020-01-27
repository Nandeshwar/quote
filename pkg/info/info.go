package info

import "time"

//go:generate gofp -destination fp.go -pkg info -type "Info"
type Info struct {
	Title        string
	Info         string
	Link         []string
	CreationDate time.Time
}

func GetAllInfo() []Info {
	var allInfo []Info
	allInfo = append(allInfo, getMiscInfo()...)
	allInfo = append(allInfo, getKripaluJiMaharajInfo()...)
	return allInfo
}
