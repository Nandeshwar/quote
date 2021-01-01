package emailservice

import (
	"time"
)

//go:generate gofp -destination fpEmailStatus.go -pkg emailservice -type "EmailStatus"
type EmailStatus struct {
	typ    string
	sentAt time.Time
}

type EmailStatusInMemory struct {
	emailStatusList []EmailStatus
}

func (e *EmailStatusInMemory) Add(emailStatus EmailStatus) {
	e.emailStatusList = append(e.emailStatusList, emailStatus)
}

func (e *EmailStatusInMemory) Exists(emailStatus EmailStatus) bool {
	emailStatusExists := func(es EmailStatus) bool {
		return es.sentAt.Year() == emailStatus.sentAt.Year() &&
			es.sentAt.Month() == emailStatus.sentAt.Month() &&
			es.sentAt.Day() == emailStatus.sentAt.Day() &&
			es.typ == emailStatus.typ
	}

	return SomeEmailStatus(emailStatusExists, e.emailStatusList)
}

func (e *EmailStatusInMemory) Delete(emailStatus EmailStatus) {
	emailStatusExists := func(es EmailStatus) bool {
		return es.sentAt.Year() == emailStatus.sentAt.Year() &&
			es.sentAt.Month() == emailStatus.sentAt.Month() &&
			es.sentAt.Day() == emailStatus.sentAt.Day() &&
			es.typ == emailStatus.typ
	}

	e.emailStatusList = RemoveEmailStatus(emailStatusExists, e.emailStatusList)
}
