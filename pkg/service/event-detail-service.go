package service

import (
	"fmt"
	"quote/pkg/constants"
	"quote/pkg/model"
	"regexp"
	"strings"
	"time"
)

//go:generate mockgen -destination "mock/mock_eventdetail.go" -source "event-detail-service.go" IEventDetail
type IEventDetail interface {
	ValidateFormEvent(form model.EventDetailForm) error
	CreateNewEventDetail(form model.EventDetailForm) (int64, error)
	GetEventDetailByTitleOrInfo(searchTxt string) ([]model.EventDetail, error)
	GetEventDetailByMonthDay(month, day int) ([]model.EventDetail, error)
	GetEventDetailByID(ID int64) (model.EventDetail, error)
	UpdateEventDetailByID(eventDetail model.EventDetail) error
}

func (s InfoEventService) ValidateFormEvent(form model.EventDetailForm) error {
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

func (s InfoEventService) CreateNewEventDetail(form model.EventDetailForm) (int64, error) {
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

func (s InfoEventService) GetEventDetailByTitleOrInfo(searchTxt string) ([]model.EventDetail, error) {
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

func (s InfoEventService) GetEventDetailByMonthDay(month, day int) ([]model.EventDetail, error) {
	eventDetailList, err := s.EventDetailRepo.GetEventDetailByMonthDay(month, day)
	if err != nil {
		return nil, err
	}

	return aggregateEventDetailByURL(eventDetailList), nil
}

func (s InfoEventService) EventsInFuture(t time.Time) ([]model.EventDetail, error) {
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

func (s InfoEventService) GetEventDetailByID(ID int64) (model.EventDetail, error) {
	evenDetailList, err := s.EventDetailRepo.GetEventDetailByID(ID)
	if err != nil {
		return model.EventDetail{}, err
	}

	if len(evenDetailList) == 0 {
		return model.EventDetail{}, fmt.Errorf("event detail id=%d not found", ID)
	}

	var links []string
	for i := 0; i < len(evenDetailList); i++ {
		links = append(links, evenDetailList[i].URL)
	}
	evenDetailList[0].Links = links
	return evenDetailList[0], nil
}

func (s InfoEventService) UpdateEventDetailByID(eventDetail model.EventDetail) error {
	updatedAt := time.Now()
	eventDetail.UpdatedAt = updatedAt
	err := s.EventDetailRepo.UpdateEventDetailByID(eventDetail)
	if err != nil {
		return err
	}
	return nil
}
