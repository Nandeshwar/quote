package api

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"quote/pkg/model"
)

func (s Server) adminInfo(w http.ResponseWriter, r *http.Request) {
	session, err := s.sessionCookieStore.Get(r, s.cookieName)
	if err != nil {
		logrus.Error("error getting session info", err)
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		logrus.Errorf("session expired")
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	t, err := template.ParseFiles(s.views.AdminInfo)
	if err != nil {
		logrus.Error("error parsing view AdminInfo", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		err := t.Execute(w, nil)
		if err != nil {
			logrus.Error("error executing view AdminInfo", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		return

	case "POST":
		err := r.ParseForm()
		if err != nil {
			logrus.Error("error parsing form AdminInfo", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if len(r.Form["title"]) == 0 || len(r.Form["info"]) == 0 || len(r.Form["link"]) == 0 || len(r.Form["createdAt"]) == 0 {
			logrus.Errorf("check if these fields(title, info, link,createdAt) exist in file view/admin-info.gptl")
			w.WriteHeader(http.StatusNotFound)
			return
		}
		infoForm := model.InfoForm{
			Title:     r.Form["title"][0],
			Info:      r.Form["info"][0],
			Link:      r.Form["link"][0],
			CreatedAt: r.Form["createdAt"][0],
		}

		err = s.infoService.ValidateForm(infoForm)
		if err != nil {
			logrus.WithError(err).Errorf("invalid data for info")
			status := map[string]interface{}{"Status": "validation error. check log"}
			// Also return 200
			err := t.Execute(w, status)
			if err != nil {
				logrus.Error("error executing view AdminInfo", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		}

		IDs, err := s.infoService.GetInfoLinkIDs(infoForm.Link)
		if err != nil {
			logrus.Errorf("error checking existence of links=%v", err)
		}

		id, err := s.infoService.CreateNewInfo(r.Context(), infoForm)
		if err != nil {
			logrus.WithError(err).Error("error creating info")
			status := map[string]interface{}{"Status": "Did not create record. error. check log"}

			// Also set header to 200
			err := t.Execute(w, status)
			if err != nil {
				logrus.Error("error executing view AdminInfo", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		}

		logrus.WithField("id", id).Info("created info")

		status := map[string]interface{}{"Status": fmt.Sprintf("SUCCESS. ID=%d. Link IDs=%v", id, IDs)}
		err = t.Execute(w, status)
		if err != nil {
			logrus.Error("error executing view AdminInfo", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}
