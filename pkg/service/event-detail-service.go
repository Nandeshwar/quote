package service

import (
	"fmt"
	"quote/pkg/constants"
	"quote/pkg/model"
	"regexp"
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
	if len(eventDate) <= 0 {
		return fmt.Errorf("event date should not be empty")
	}

	re := regexp.MustCompile("[0-9]{4}-[0-9]{2}-[0-9]{2}")
	if !re.MatchString(eventDate) {
		return fmt.Errorf("wrong date and time format for event date. given date=%s, please provide date in this format yyyy-mm-dd", eventDate)
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
			return fmt.Errorf("invalid value for event type. given type=%s expected value: same/different", typ)
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
		links = strings.Split(link, "|")
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

	return aggregateEventDetailByURL(eventDetailList), nil
}

func aggregateEventDetailByURL(eventDetailList []model.EventDetail) []model.EventDetail {
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
	return distinctEventDetailList
}

func (s QuoteService) GetEventDetailByMonthDay(month, day int) ([]model.EventDetail, error) {
	eventDetailList, err := s.EventDetailRepo.GetEventDetailByMonthDay(month, day)
	if err != nil {
		return nil, err
	}

	return aggregateEventDetailByURL(eventDetailList), nil
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

	eventsInFuture := model.FilterEventDetail(findTodayEvent, events)

	return eventsInFuture, nil
}
