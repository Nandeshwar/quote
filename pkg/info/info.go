package info

import (
	"math/rand"
	"strings"
	"time"
)

//go:generate gofp -destination fp.go -pkg info -type "Info, string"
type Info struct {
	ID           int
	Title        string
	Info         string
	Link         []string
	CreationDate time.Time
	UpdatedDate  time.Time
}

func GetAllInfo() []Info {
	var allInfo []Info
	allInfo = append(allInfo, getMiscInfo()...)
	allInfo = append(allInfo, getKripaluJiMaharajInfo()...)
	allInfo = append(allInfo, getMyMeditaionInfo()...)
	return allInfo
}

func GetRandomTwoWordsFromTitle() (word1, word2 string) {
	allInfo := GetAllInfo()

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s) // initialize local pseudorandom generator
	ind := r.Intn(len(allInfo))

	info := allInfo[ind]
	words := strings.Split(info.Title, " ")
	return words[0], words[1]
}
