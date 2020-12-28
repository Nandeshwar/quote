package emailservice

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"github.com/logic-building/functional-go/fp"
	gomail "gopkg.in/mail.v2"
	"html/template"
	"math/rand"
	"quote/pkg/model"
	"quote/pkg/service"
	"quote/pkg/service/quote"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"quote/pkg/repo"
)

type IEmailQuote interface {
	SendEmailForEventDetail()
}

type EmailQuote struct {
	emailStatusRepo       repo.IEmailStatusRepo
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
		eventDetailService:    eventDetailService,
		emailServer:           emailServer,
		emailServerPort:       emailServerPort,
		emailFrom:             emailFrom,
		emailFromPwd:          emailFromPwd,
		emailToForEvents:      emailToForEvents,
		emailToForImageQuotes: emailToForImageQuotes,
	}
}

func (e EmailQuote) SendEmailForEventDetail(ctx context.Context) {
	eventsFor7Days := make(map[int][]model.EventDetail, 7)

	for _, day := range fp.RangeInt(0, 7) {
		today := time.Now()
		futureTime := today.AddDate(0, 0, day)

		eventsInFuture, err := e.eventDetailService.EventsInFuture(futureTime)
		if err != nil {
			logrus.Errorf("error while getting EventsInFuture, error=%v", err)
		}

		eventsFor7Days[day] = eventsInFuture
	}

	subject := "Quote: Events for the next 3 days. Radhe Krishna"
	logrus.Debugf("email service started")
	now := time.Now()
	typ := "event-detail"
	found := e.emailStatusRepo.EmailSentForEvents(ctx, now, typ)

	if found {
		logrus.Infof("email is already sent for events")
		return
	}

	todayEvents := eventsFor7Days[0]
	tomorrowEvents := eventsFor7Days[1]
	dayAfterTomorrowEvents := eventsFor7Days[2]

	if len(todayEvents)+len(tomorrowEvents)+len(dayAfterTomorrowEvents) == 0 {
		logrus.Infof("no events to notify")
		return
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
		SentAt: now,
	}
	logrus.Infof("Email sent successfully for events")

	_, err = e.emailStatusRepo.CreateEmailStatus(ctx, emailStauts)
	if err != nil {
		logrus.Errorf("error creating email status", err)
	}

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

func (e EmailQuote) SendEmailForQuoteImage(ctx context.Context) {
	_, allImages := quote.AllQuotesImage()
	s2 := rand.NewSource(int64(time.Now().Nanosecond()))
	r2 := rand.New(s2)
	ind := r2.Intn(len(allImages))
	image := allImages[ind]
	imagePath := "./" + image
	logrus.Infof("Image path=%v", imagePath)

	i := strings.LastIndex(imagePath, "/")
	imageName := imagePath[i:]

	now := time.Now()
	typ := "quote-image"

	found := e.emailStatusRepo.EmailSentForEvents(ctx, now, typ)

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
		SentAt: now,
	}
	logrus.Infof("Email sent successfully for quote-image")

	_, err := e.emailStatusRepo.CreateEmailStatus(ctx, emailStauts)
	if err != nil {
		logrus.Errorf("error creating email status", err)
	}
}
