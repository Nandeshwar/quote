package model

import (
	"encoding/json"
	"quote/pkg/constants"
	"time"
)

//go:generate gofp -destination fp.go -pkg model -type "Info, string"

type Info struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Info      string    `json:"info"`
	Links     []string  `json:"links"`
	Link      string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type InfoGORM struct {
	ID        int64
	Title     string
	Info      string
	CreatedAt time.Time
	UpdatedAt time.Time

	InfoLinks []InfoLinkGORM `gorm:"foreignKey:LinkID;references:ID"`
}

type InfoLinkGORM struct {
	ID        int64
	LinkID    int64
	Link      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Tabler interface {
	TableName() string
}

// override table name with info
func (InfoGORM) TableName() string {
	return "info"
}

//override table name with info_link
func (InfoLinkGORM) TableName() string {
	return "info_link"
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
		UpdatedDate:  d.UpdatedAt.Format(constants.DATE_FORMAT),
		CreationDate: d.CreatedAt.Format(constants.DATE_FORMAT),
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

	i.CreatedAt = newCreatedAt
	i.UpdatedAt = newUpdatedAt
	return nil
}
