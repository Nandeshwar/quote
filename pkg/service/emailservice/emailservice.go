package emailservice

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"html/template"
	"math/rand"
	"os"
	"quote/pkg/excel"
	"quote/pkg/model"
	"quote/pkg/service"
	"quote/pkg/service/quote"
	"strings"
	"time"

	"github.com/logic-building/functional-go/fp"
	gomail "gopkg.in/mail.v2"

	"github.com/sirupsen/logrus"

	"quote/pkg/repo"
)

type IEmailQuote interface {
	SendEmailForEventDetail()
}

type EmailQuote struct {
	emailStatusRepo       repo.IEmailStatusRepo
	emailStatusInMemory   EmailStatusInMemory
	eventDetailService    service.IEventDetail
	emailServer           string
	emailServerPort       int
	emailFrom             string
	emailFromPwd          string
	emailToForEvents      []string
	emailToForImageQuotes []string
}

func NewEmailQuote(
	sqlite3DB repo.SQLite3Repo,
	emailStatusInMemory EmailStatusInMemory,
	eventDetailService service.IEventDetail,
	emailServer string,
	emailServerPort int,
	emailFrom string,
	emailFromPwd string,
	emailToForEvents []string,
	emailToForImageQuotes []string,
) EmailQuote {
	return EmailQuote{
		emailStatusRepo:       sqlite3DB,
		emailStatusInMemory:   emailStatusInMemory,
		eventDetailService:    eventDetailService,
		emailServer:           emailServer,
		emailServerPort:       emailServerPort,
		emailFrom:             emailFrom,
		emailFromPwd:          emailFromPwd,
		emailToForEvents:      emailToForEvents,
		emailToForImageQuotes: emailToForImageQuotes,
	}
}

func (e *EmailQuote) SendEmailForEventDetail(ctx context.Context) {

	subject := "Quote: Events for the next 3 days. Radhe Krishna"
	logrus.Debugf("email service started")
	nowUTC := time.Now().UTC()
	typ := "event-detail"

	if e.emailStatusInMemory.Exists(EmailStatus{
		typ:    typ,
		sentAt: time.Now().Local(),
	}) {
		logrus.Info("Record already exist in memory. Email-Quote already sent for the day")
		return
	}

	found := e.emailStatusRepo.EmailSentForEvents(ctx, time.Now().Local(), typ)

	if found {
		logrus.Infof("email is already sent for events")
		return
	}

	eventsFor7DaysMap := make(map[int][]model.EventDetail, 7)
	var EventsInFuture7days []model.EventDetail

	for _, day := range fp.RangeInt(0, 7) {
		today := time.Now()
		futureTime := today.AddDate(0, 0, day)

		eventsInFuture, err := e.eventDetailService.EventsInFuture(futureTime)
		if err != nil {
			logrus.Errorf("error while getting EventsInFuture, error=%v", err)
		}

		eventsFor7DaysMap[day] = eventsInFuture
		EventsInFuture7days = append(EventsInFuture7days, eventsInFuture...)

	}

	todayEvents := eventsFor7DaysMap[0]
	tomorrowEvents := eventsFor7DaysMap[1]
	dayAfterTomorrowEvents := eventsFor7DaysMap[2]

	if len(todayEvents)+len(tomorrowEvents)+len(dayAfterTomorrowEvents) == 0 {
		logrus.Infof("no events to notify")
		return
	}

	excelFilePath, excelErr := excel.CreateExcelEventList("event_list.xlsx", "events next 7 days", EventsInFuture7days)
	if excelErr != nil {
		logrus.Errorf("error creating excel file=%v", excelErr)
	}

	type Data struct {
		Day1                   string
		TodayEvents            []model.EventDetail
		Day2                   string
		TomorrowEvents         []model.EventDetail
		Day3                   string
		DayAfterTomorrowEvents []model.EventDetail
	}

	data := Data{
		Day1:                   "Today",
		TodayEvents:            todayEvents,
		Day2:                   "Tomorrow",
		TomorrowEvents:         tomorrowEvents,
		Day3:                   "Day after tomorrow",
		DayAfterTomorrowEvents: dayAfterTomorrowEvents,
	}

	body, err := ParseTemplate("./views/email.gtpl", data)
	if err != nil {
		logrus.Errorf("error parsing email template. error=%v", err)
		return
	}

	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", e.emailFrom)

	m.SetHeader("To", e.emailToForEvents...)

	// Set E-Mail subject
	m.SetHeader("Subject", subject)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/html", body)

	if excelFilePath != "" {
		m.Attach(excelFilePath)
	}

	// Settings for SMTP server
	d := gomail.NewDialer(e.emailServer, e.emailServerPort, e.emailFrom, e.emailFromPwd)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		logrus.Error("\nerror dialing and sending emails=", err)
		return
	}

	emailStauts := model.EmailStatusGORM{
		Status: "sent",
		Typ:    typ,
		SentAt: nowUTC,
	}
	logrus.Infof("Email sent successfully for events")

	_, err = e.emailStatusRepo.CreateEmailStatus(ctx, emailStauts)
	if err != nil {
		logrus.Errorf("error creating email status. error=%v", err)
		return
	}

	err = os.Remove(excelFilePath)
	if err != nil {
		logrus.Errorf("error deleting excel file=%s", excelFilePath)
	}
	logrus.Infof("excel file=%s deleted successfully", excelFilePath)

	e.emailStatusInMemory.Add(EmailStatus{
		typ:    typ,
		sentAt: time.Now().Local(),
	})

	return
}

func ParseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (e *EmailQuote) SendEmailForQuoteImage(ctx context.Context) {
	_, allImages := quote.AllQuotesImage()
	s2 := rand.NewSource(int64(time.Now().Nanosecond()))
	r2 := rand.New(s2)
	ind := r2.Intn(len(allImages))
	image := allImages[ind]
	imagePath := "./" + image
	logrus.Infof("Image path=%v", imagePath)

	i := strings.LastIndex(imagePath, "/")
	imageName := imagePath[i:]

	nowUTC := time.Now().UTC()
	typ := "quote-image"

	if e.emailStatusInMemory.Exists(EmailStatus{
		typ:    typ,
		sentAt: time.Now().Local(),
	}) {
		logrus.Info("Record already exist in memory. Email-Quote already sent for the day")
		return
	}

	found := e.emailStatusRepo.EmailSentForEvents(ctx, time.Now().Local(), typ)

	if found {
		logrus.Infof("email is already sent for quote image")
		return
	}

	subject := "Quote-Image for the day"

	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", e.emailFrom)

	m.SetHeader("To", e.emailToForImageQuotes...)

	// Set E-Mail subject
	m.SetHeader("Subject", subject)

	// Set E-Mail body. You can set plain text or html with text/html
	m.Embed(imagePath)
	m.SetBody("text/html", fmt.Sprintf(`<img src="cid:%s" alt="जय श्री कृपालु जी महाराज" />`, imageName))

	// Settings for SMTP server
	d := gomail.NewDialer(e.emailServer, e.emailServerPort, e.emailFrom, e.emailFromPwd)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		logrus.Error("\nerror dialing and sending emails=", err)
		return
	}

	emailStauts := model.EmailStatusGORM{
		Status: "sent",
		Typ:    typ,
		SentAt: nowUTC,
	}
	logrus.Infof("Email sent successfully for quote-image")

	id, err := e.emailStatusRepo.CreateEmailStatus(ctx, emailStauts)
	if err != nil {
		logrus.Errorf("error creating email status. error=%v", err)
		return
	}

	if id > 0 {
		logrus.Info("Email status saved to database successfully for quote image")
	}

	e.emailStatusInMemory.Add(EmailStatus{
		typ:    typ,
		sentAt: time.Now().Local(),
	})

}
