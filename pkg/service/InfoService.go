package service

import (
	"fmt"
	"quote/pkg/constants"
	"quote/pkg/model"
	"strings"
	"time"
)

type IInfo interface {
	ValidateForm(form model.InfoForm) error
}

func (s QuoteService) ValidateForm(form model.InfoForm) error {
	createdAt := strings.TrimSpace(form.CreateAt)
	if len(createdAt) > 0 {
		if len(createdAt) != 16 || createdAt[4] != '-' || createdAt[7] != '-' || createdAt[13] != ':' {
			return fmt.Errorf("wrong date and time format. please provide date in this format yyyy-mm-dd tt:mm")
		}
	}

	_, err := time.Parse(constants.DATE_FORMAT, createdAt)
	if err != nil {
		return err
	}
	return nil
}
