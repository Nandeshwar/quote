package service

import (
	"fmt"
	"quote/pkg/constants"
	"quote/pkg/model"
	"strings"
	"time"
)

type IEventDetail interface {
	ValidateFormEvent(form model.EventDetailForm) error
	CreateNewEventDetail(form model.EventDetailForm) (int64, error)
	GetEventDetailByTitleOrInfo(searchTxt string) ([]model.EventDetail, error)
	GetEventDetailByMonthDay(month, day int) ([]model.EventDetail, error)
}

func (s QuoteService) ValidateFormEvent(form model.EventDetailForm) error {
	err := validateEventDate(form.EventDate)
	if err != nil {
		return err
	}

	err = validateEventType(form.Typ)
	if err != nil {
		return err
	}

	err = validateCreatedAt(form.CreatedAt)
	if err != nil {
		return err
	}

	err = validateLink(form.Link)
	if err != nil {
		return err
	}

	return nil
}

func validateEventDate(eventDate string) error {
	eventDate = strings.TrimSpace(eventDate)
	if len(eventDate) < 0 {
		return fmt.Errorf("event date should not be empty")
	}

	if len(eventDate) != 10 || eventDate[4] != '-' || eventDate[7] != '-' {
		return fmt.Errorf("wrong date and time format for event date. please provide date in this format yyyy-mm-dd tt:mm")
	}
	_, err := time.Parse(constants.DATE_FORMAT_EVENT_DATE, eventDate)
	if err != nil {
		return err
	}

	return nil
}

func validateEventType(typ string) error {
	typ = strings.TrimSpace(typ)
	if len(typ) > 0 {
		typ = strings.ToLower(typ)
		if strings.ToLower(typ) != "same" && strings.ToLower(typ) != "different" {
			return fmt.Errorf("invalid value for event type. expected value: same/different")
		}
	}
	return nil
}

func validateCreatedAt(createdAt string) error {
	createdAt = strings.TrimSpace(createdAt)
	if len(createdAt) > 0 {
		if len(createdAt) != 16 || createdAt[4] != '-' || createdAt[7] != '-' || createdAt[13] != ':' {
			return fmt.Errorf("wrong date and time format for createdAt. given date=%s, please provide date in this format yyyy-mm-dd tt:mm", createdAt)
		}
		_, err := time.Parse(constants.DATE_FORMAT, createdAt)
		if err != nil {
			return err
		}
	}
	return nil
}

func validateLink(link string) error {
	link = strings.TrimSpace(link)
	if len(link) > 0 {
		for _, link := range strings.Split(link, "|") {
			link = strings.TrimSpace(link)
			if len(link) < 4 {
				return fmt.Errorf("pipeline(|) seperated links value must start with http or https. link could not be less than 4")
			}
			if link[0:4] != "http" {
				return fmt.Errorf("pipeline(|) seperated links value must start with http or https")
			}

			if link[len(link)-1] == '"' || link[len(link)-1] == '\'' || link[len(link)-1] == '.' {
				return fmt.Errorf("pipeline(|) seperated link's value should not ended with (\", ', .)")
			}
		}
	}
	return nil
}

func (s QuoteService) CreateNewEventDetail(form model.EventDetailForm) (int64, error) {
	var createdAt time.Time
	var err error

	eventDate, err := time.Parse(constants.DATE_FORMAT_EVENT_DATE, form.EventDate)
	if err != nil {
		return 0, err
	}

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
		links = strings.Split(link, ",")
	}
	eventDetail := model.EventDetail{
		Title:        form.Title,
		Info:         form.Info,
		Day:          eventDate.Day(),
		Month:        int(eventDate.Month()),
		Year:         eventDate.Year(),
		Type:         form.Typ,
		Links:        links,
		CreationDate: createdAt,
		UpdatedAt:    time.Now(),
	}

	if len(eventDetail.Type) == 0 {
		eventDetail.Type = "same"
	}

	id, err := s.EventDetailRepo.CreateEventDetail(eventDetail)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s QuoteService) GetEventDetailByTitleOrInfo(searchTxt string) ([]model.EventDetail, error) {
	eventDetailList, err := s.EventDetailRepo.GetEventDetailByTitleOrInfo(searchTxt)
	if err != nil {
		return nil, err
	}

	eventDetailListSorted := model.SortEventDetailByID(eventDetailList)

	var distinctEventDetailList []model.EventDetail
	var links []string
	found := false
	for i := 0; i < len(eventDetailListSorted); i++ {
		if i+1 < len(eventDetailListSorted) && eventDetailListSorted[i].ID == eventDetailListSorted[i+1].ID {
			links = append(links, eventDetailListSorted[i].URL)
			found = true
		} else {
			found = false
		}

		if !found {
			links = append(links, eventDetailListSorted[i].URL)
			eventDetailListSorted[i].Links = make([]string, len(links))
			copy(eventDetailListSorted[i].Links, links)
			links = nil
			distinctEventDetailList = append(distinctEventDetailList, eventDetailListSorted[i])
		}
	}
	return distinctEventDetailList, nil
}

func (s QuoteService) GetEventDetailByMonthDay(month, day int) ([]model.EventDetail, error) {
	eventDetailList, err := s.EventDetailRepo.GetEventDetailByMonthDay(month, day)
	if err != nil {
		return nil, err
	}

	eventDetailListSorted := model.SortEventDetailByID(eventDetailList)

	var distinctEventDetailList []model.EventDetail
	var links []string
	found := false
	for i := 0; i < len(eventDetailListSorted); i++ {
		if i+1 < len(eventDetailListSorted) && eventDetailListSorted[i].ID == eventDetailListSorted[i+1].ID {
			links = append(links, eventDetailListSorted[i].URL)
			found = true
		} else {
			found = false
		}

		if !found {
			links = append(links, eventDetailListSorted[i].URL)
			eventDetailListSorted[i].Links = make([]string, len(links))
			copy(eventDetailListSorted[i].Links, links)
			links = nil
			distinctEventDetailList = append(distinctEventDetailList, eventDetailListSorted[i])
		}
	}
	return distinctEventDetailList, nil
}

func (s QuoteService) EventsInFuture(t time.Time) ([]model.EventDetail, error) {
	year, month, day := t.Date()

	events, err := s.GetEventDetailByMonthDay(int(month), day)
	if err != nil {
		return nil, err
	}

	findTodayEvent := func(event model.EventDetail) bool {
		if event.Type == "different" && event.Year != year {
			return false
		}
		return true
	}

	todayEvents := model.FilterEventDetail(findTodayEvent, events)

	return todayEvents, nil
}
