package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/logic-building/functional-go/fp"

	"quote/pkg/constants"
	"quote/pkg/model"
)

//go:generate mockgen -source "InfoService.go" -destination "mock/mock_iinfo.go" IInfo
type IInfo interface {
	ValidateForm(form model.InfoForm) error
	CreateNewInfo(ctx context.Context, form model.InfoForm) (int64, error)
	GetInfoByTitleOrInfo(searchTxt string) ([]model.Info, error)
	UpdateInfoByID(info model.Info) error
	GetInfoByID(ctx context.Context, ID int64) (model.Info, error)
	GetInfoLinkIDs(link string) ([]int64, error)
}

func (s InfoEventService) ValidateForm(form model.InfoForm) error {
	createdAt := strings.TrimSpace(form.CreatedAt)
	err := validateCreatedAt(createdAt)
	if err != nil {
		return err
	}

	err = validateLink(form.Link)
	if err != nil {
		return err
	}

	return nil
}

func (s InfoEventService) CreateNewInfo(ctx context.Context, form model.InfoForm) (int64, error) {
	var createdAt time.Time
	var err error

	if len(strings.TrimSpace(form.CreatedAt)) > 0 {
		createdAt, err = time.Parse(constants.DATE_FORMAT, form.CreatedAt)
		if err != nil {
			return 0, err
		}
	} else {
		createdAt = time.Now()
	}

	link := strings.TrimSpace(form.Link)
	var links []string
	if len(link) > 0 {
		links = strings.Split(link, "|")
	}
	info := model.Info{
		Title:        form.Title,
		Info:         form.Info,
		Links:        links,
		CreationDate: createdAt,
		UpdatedDate:  time.Now(),
	}
	id, err := s.InfoRepo.CreateInfo(ctx, info)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s InfoEventService) GetInfoByTitleOrInfo(searchTxt string) ([]model.Info, error) {
	infoList, err := s.InfoRepo.GetInfoByTitleOrInfo(searchTxt)
	if err != nil {
		return nil, err
	}

	infoListSorted := model.SortInfoByID(infoList)

	var distinctInfoList []model.Info
	var links []string
	found := false
	for i := 0; i < len(infoListSorted); i++ {
		if i+1 < len(infoListSorted) && infoListSorted[i].ID == infoListSorted[i+1].ID {
			links = append(links, infoListSorted[i].Link)
			found = true
		} else {
			found = false
		}

		if !found {
			links = append(links, infoListSorted[i].Link)
			infoListSorted[i].Links = links
			links = nil
			distinctInfoList = append(distinctInfoList, infoListSorted[i])
		}
	}
	return distinctInfoList, nil
}

func (s InfoEventService) GetInfoByID(ctx context.Context, ID int64) (model.Info, error) {
	infoList, err := s.InfoRepo.GetInfoByID(ctx, ID)
	if err != nil {
		return model.Info{}, err
	}

	if len(infoList) == 0 {
		return model.Info{}, fmt.Errorf("info id=%d not found", ID)
	}

	var links []string
	for i := 0; i < len(infoList); i++ {
		links = append(links, infoList[i].Link)
	}
	infoList[0].Links = links
	return infoList[0], nil
}

func (s InfoEventService) GetInfoLinkIDs(link string) ([]int64, error) {
	link = strings.TrimSpace(link)
	var links []string
	if len(link) > 0 {
		links = strings.Split(link, "|")
	}

	links = fp.MapStr(strings.TrimSpace, links)

	linkIds, err := s.InfoRepo.GetInfoLinkIDs(links)
	if err != nil {
		return nil, err
	}
	return linkIds, nil
}

func (s InfoEventService) UpdateInfoByID(info model.Info) error {
	updatedAt := time.Now()
	info.UpdatedDate = updatedAt
	err := s.InfoRepo.UpdateInfoByID(info)
	if err != nil {
		return err
	}
	return nil
}
