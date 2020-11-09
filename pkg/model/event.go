package model

import (
	"encoding/json"
	"fmt"
	"quote/pkg/constants"
	"strings"
	"time"

	"github.com/gookit/color"
)

//go:generate gofp -destination fpEventDetail.go -pkg model -type "EventDetail"
type EventDetail struct {
	ID           int64     `json:"id"`
	Day          int       `json:"day"`
	Month        int       `json:"month"`
	Year         int       `json:"year"`
	Title        string    `json:"title"`
	Info         string    `json:"info"`
	URL          string    `json:"-"`
	Links        []string  `json:"links"`
	Type         string    `json:"type"` // Value can be same|different: different - event occurs different day in each year. Example Ram birthday is different each year as per Bikram Sambat
	CreationDate time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func (e EventDetail) DisplayEvent() {
	blue := color.FgBlue.Render
	fmt.Println(e.Title)
	fmt.Println(e.Info)
	fmt.Println(e.URL)
	fmt.Printf("%d-%d-%d\n", e.Year, e.Month, e.Day)
	for i, url := range strings.Split(e.URL, ";") {
		fmt.Printf("\n%d. %s ", i+1, blue(url))
	}
}

func (i *EventDetail) ToJson() ([]byte, error) {
	//jsonBytes, err := json.Marshal(i)
	jsonBytes, err := i.MarshalJSON()
	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}

func (i *EventDetail) ToJsonIndent() ([]byte, error) {
	//jsonBytes, err := json.Marshal(i)
	jsonBytes, err := i.MarshalJSONIndent()
	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}

// Custom JSON Marshalling to output time field in required format
// http://choly.ca/post/go-json-marshalling/
func (d *EventDetail) MarshalJSON() ([]byte, error) {
	type Alias EventDetail
	return json.Marshal(&struct {
		*Alias
		UpdatedDate  string `json:"updatedAt"`
		CreationDate string `json:"createdAt"`
	}{
		Alias:        (*Alias)(d),
		UpdatedDate:  d.UpdatedAt.Format(constants.DATE_FORMAT),
		CreationDate: d.CreationDate.Format(constants.DATE_FORMAT),
	})
}

// Custom JSON Marshalling to output time field in required format
// http://choly.ca/post/go-json-marshalling/
func (d *EventDetail) MarshalJSONIndent() ([]byte, error) {
	type Alias EventDetail
	return json.MarshalIndent(&struct {
		*Alias
		UpdatedDate  string `json:"updatedAt"`
		CreationDate string `json:"createdAt"`
	}{
		Alias:        (*Alias)(d),
		UpdatedDate:  d.UpdatedAt.Format(constants.DATE_FORMAT),
		CreationDate: d.CreationDate.Format(constants.DATE_FORMAT),
	}, "", " ")
}

func (i *EventDetail) UnmarshalJSON(data []byte) error {
	type Alias EventDetail
	aux := &struct {
		*Alias
		UpdatedAt    string `json:"updatedAt"`
		CreationDate string `json:"createdAt"`
	}{
		Alias: (*Alias)(i),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	newCreatedAt, err := time.Parse(constants.DATE_FORMAT, aux.CreationDate)
	if err != nil {
		return err
	}

	newUpdatedAt, err := time.Parse(constants.DATE_FORMAT, aux.UpdatedAt)
	if err != nil {
		return err
	}

	i.CreationDate = newCreatedAt
	i.UpdatedAt = newUpdatedAt
	return nil
}
