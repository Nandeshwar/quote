package service

import (
	"fmt"
	"quote/pkg/constants"
	"regexp"
	"strings"
	"time"
)

func validateLink(link string) error {
	link = strings.TrimSpace(link)
	if len(link) > 0 {
		for _, link := range strings.Split(link, "|") {
			link = strings.TrimSpace(link)
			if len(link) < 4 {
				return fmt.Errorf("pipeline(|) separated links value must start with http or https. link could not be less than 4")
			}
			if link[0:4] != "http" {
				return fmt.Errorf("pipeline(|) separated links value must start with http or https")
			}

			if link[len(link)-1] == '"' || link[len(link)-1] == '\'' || link[len(link)-1] == '.' {
				return fmt.Errorf("pipeline(|) separated link's value should not ended with (\", ', .)")
			}
		}
	}
	return nil
}

func validateCreatedAt(createdAt string) error {
	createdAt = strings.TrimSpace(createdAt)
	if len(createdAt) > 0 {
		re := regexp.MustCompile(`[0-9]{4}[-][0-9]{2}[-][0-9]{2} [0-9]{2}:[0-9]{2}`)

		if !re.MatchString(createdAt) {
			return fmt.Errorf("wrong date and time format for createdAt. given date=%s, please provide date in this format yyyy-mm-dd tt:mm", createdAt)
		}

		_, err := time.Parse(constants.DATE_FORMAT, createdAt)
		if err != nil {
			return err
		}
	}
	return nil
}
