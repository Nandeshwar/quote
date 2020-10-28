package model

import (
	"encoding/json"
	"quote/pkg/constants"
	"time"
)

//go:generate gofp -destination fp.go -pkg model -type "Info, string"

type Info struct {
	ID           int64     `json:"id"`
	Title        string    `json:"title"`
	Info         string    `json:"info"`
	Links        []string  `json:"links"`
	Link         string    `json:"-"`
	CreationDate time.Time `json:"createdAt"`
	UpdatedDate  time.Time `json:"updatedAt"`
}

func (i *Info) ToJson() ([]byte, error) {
	//jsonBytes, err := json.Marshal(i)
	jsonBytes, err := i.MarshalJSON()
	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}

// Custom JSON Marshalling to output time field in required format
// http://choly.ca/post/go-json-marshalling/
func (d *Info) MarshalJSON() ([]byte, error) {
	type Alias Info
	return json.Marshal(&struct {
		*Alias
		UpdatedDate  string `json:"updatedAt"`
		CreationDate string `json:"createdAt"`
	}{
		Alias:        (*Alias)(d),
		UpdatedDate:  d.UpdatedDate.Format(constants.DATE_FORMAT),
		CreationDate: d.CreationDate.Format(constants.DATE_FORMAT),
	})
}

func (i *Info) UnmarshalJSON(data []byte) error {
	type Alias Info
	aux := &struct {
		*Alias
		UpdatedDate  string `json:"updatedAt"`
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

	newUpdatedAt, err := time.Parse(constants.DATE_FORMAT, aux.UpdatedDate)
	if err != nil {
		return err
	}

	i.CreationDate = newCreatedAt
	i.UpdatedDate = newUpdatedAt
	return nil
}
