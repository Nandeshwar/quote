package api

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"quote/pkg/model"
)

func (s Server) adminEvent(w http.ResponseWriter, r *http.Request) {
	session, err := s.sessionCookieStore.Get(r, s.cookieName)
	if err != nil {
		logrus.Error("error getting session info", err)
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		logrus.Error("session expired")
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	t, err := template.ParseFiles(s.views.AdminEventDetail)
	if err != nil {
		logrus.Error("error parsing view admin event detail", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		err := t.Execute(w, nil)
		if err != nil {
			logrus.Error("error executing view Admin Event detail view", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		return

	case "POST":
		err := r.ParseForm()
		if err != nil {
			logrus.Error("error parsing form Admin Event Detail", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if len(r.Form["title"]) == 0 || len(r.Form["info"]) == 0 || len(r.Form["eventDate"]) == 0 || len(r.Form["eventType"]) == 0 || len(r.Form["link"]) == 0 || len(r.Form["createdAt"]) == 0 {
			logrus.Errorf("check if these fields(title, info) exist in file view/admin-info.gptl")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		eventDetailForm := model.EventDetailForm{
			Title:     r.Form["title"][0],
			Info:      r.Form["info"][0],
			EventDate: r.Form["eventDate"][0],
			Typ:       r.Form["eventType"][0],
			Link:      r.Form["link"][0],
			CreatedAt: r.Form["createdAt"][0],
		}

		err = s.eventDetailService.ValidateFormEvent(eventDetailForm)
		if err != nil {
			logrus.WithError(err).Error("invalid data for event detail")
			status := map[string]interface{}{"Status": "validation error. check log"}

			err := t.Execute(w, status)
			if err != nil {
				logrus.Error("error executing view Admin Event View", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		}

		id, err := s.eventDetailService.CreateNewEventDetail(eventDetailForm)
		if err != nil {
			logrus.WithError(err).Error("error creating event detail")
			status := map[string]interface{}{"Status": "Did not create record. error. check log"}
			err := t.Execute(w, status)
			if err != nil {
				logrus.Error("error executing view Admin event view", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		}
		logrus.WithField("id", id).Info("created event detail")

		status := map[string]interface{}{"Status": fmt.Sprintf("SUCCESS. ID=%d", id)}
		t, err := template.ParseFiles(s.views.AdminEventDetail)
		if err != nil {
			logrus.Error("error parsing view", err)
			w.WriteHeader(http.StatusNotFound)
		}
		err = t.Execute(w, status)
		if err != nil {
			logrus.Error("error executing adminevent view", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusCreated)
	}
}
