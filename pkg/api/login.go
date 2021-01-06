package api

import (
	"fmt"
	"html/template"
	"net/http"
	"quote/pkg/model"
	"time"

	"github.com/dgrijalva/jwt-go"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

func (s Server) login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		t, err := template.ParseFiles(s.views.Login)
		if err != nil {
			logrus.Error("error parsing file ./views/login.gtpl error=", err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		today := time.Now()
		todayTime := today.AddDate(0, 0, 0)
		eventsToday, err := s.eventDetailService.EventsInFuture(todayTime)
		if err != nil {
			logrus.Errorf("error getting events in future %v", err)
		}

		tomorrow := today.AddDate(0, 0, 1)
		eventsTomorrow, err := s.eventDetailService.EventsInFuture(tomorrow)
		if err != nil {
			logrus.Errorf("error getting events in future %v", err)
		}

		dayAfterTomorrow := today.AddDate(0, 0, 2)
		eventsDayAfterTomorrow, err := s.eventDetailService.EventsInFuture(dayAfterTomorrow)
		if err != nil {
			logrus.Errorf("error getting events in future %v", err)
		}

		day4 := today.AddDate(0, 0, 3)
		eventsDay4, err := s.eventDetailService.EventsInFuture(day4)
		if err != nil {
			logrus.Errorf("error getting events in future %v", err)
		}

		day5 := today.AddDate(0, 0, 4)
		eventsDay5, err := s.eventDetailService.EventsInFuture(day5)
		if err != nil {
			logrus.Errorf("error getting events in future %v", err)
		}

		day6 := today.AddDate(0, 0, 5)
		eventsDay6, err := s.eventDetailService.EventsInFuture(day6)
		if err != nil {
			logrus.Errorf("error getting events in future %v", err)
		}

		day7 := today.AddDate(0, 0, 6)
		eventsDay7, err := s.eventDetailService.EventsInFuture(day7)
		if err != nil {
			logrus.Errorf("error getting events in future %v", err)
		}

		type Event struct {
			Day       string
			EventList []model.EventDetail
		}
		type Data struct {
			Events   []Event
			HTTPPort int
		}
		data := Data{
			HTTPPort: s.httpPort,
			Events: []Event{
				{
					Day:       fmt.Sprintf("Today,  %s", todayTime.Format("Monday Jan _2, 2006")),
					EventList: eventsToday,
				},
				Event{
					Day:       fmt.Sprintf("Tomorrow,  %s", tomorrow.Format("Monday Jan _2, 2006")),
					EventList: eventsTomorrow,
				},

				Event{
					Day:       fmt.Sprintf("Day After Tomorrow,  %s", dayAfterTomorrow.Format("Monday Jan _2, 2006")),
					EventList: eventsDayAfterTomorrow,
				},

				Event{
					Day:       fmt.Sprintf("%s", day4.Format("Monday Jan _2, 2006")),
					EventList: eventsDay4,
				},

				Event{
					Day:       fmt.Sprintf("%s", day5.Format("Monday Jan _2, 2006")),
					EventList: eventsDay5,
				},

				Event{
					Day:       fmt.Sprintf("%s", day6.Format("Monday Jan _2, 2006")),
					EventList: eventsDay6,
				},

				Event{
					Day:       fmt.Sprintf("%s", day7.Format("Monday Jan _2, 2006")),
					EventList: eventsDay7,
				},
			},
		}

		err = t.Execute(w, data)
		if err != nil {
			logrus.Error("error executing login view", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

	case "POST":
		err := r.ParseForm()
		if err != nil {
			logrus.Error("error parsing form", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		user := r.Form["username"]
		password := r.Form[("password")]

		err = s.loginService.Login(user[0], password[0])
		if err != nil {
			logrus.WithError(err).Error("invalid login id and password")
			http.Redirect(w, r, "login", http.StatusSeeOther)
			return
		}

		session, err := s.sessionCookieStore.Get(r, s.cookieName)
		if err != nil {
			logrus.WithError(err).Errorf("error getting cookie name")
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		session.Values["authenticated"] = true
		// session will be expired in given seconds
		session.Options.MaxAge = s.sessionExpireSeconds
		session.Save(r, w)

		expirationTime := time.Now().Add(15 * time.Minute)
		claims := &Claims{
			Username: user[0],
			StandardClaims: jwt.StandardClaims{
				// In JWT, the expiry time is expressed as unix milliseconds
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		var jwtKey = []byte("my_secret_key")
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			logrus.WithError(err).Error("jwt error")
			// If there is an error in creating the JWT return an internal server error
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})

		t, err := template.ParseFiles(s.views.Admin)
		if err != nil {
			logrus.Error("error executing login view", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = t.Execute(w, nil)
		if err != nil {
			logrus.Error("error executing login view", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

	}
}
