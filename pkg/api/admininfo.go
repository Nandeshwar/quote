package api

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"quote/pkg/model"
)

func (s Server) adminInfo(w http.ResponseWriter, r *http.Request) {
	session, _ := s.sessionCookieStore.Get(r, "cookie-name")

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		logrus.Errorf("session expired")
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	if r.Method == "GET" {
		t, _ := template.ParseFiles("./views/admin-info.gtpl")
		t.Execute(w, nil)
		return
	}

	r.ParseForm()
	infoForm := model.InfoForm{
		Title:     r.Form["title"][0],
		Info:      r.Form["info"][0],
		Link:      r.Form["link"][0],
		CreatedAt: r.Form["createdAt"][0],
	}

	err := s.infoService.ValidateForm(infoForm)
	if err != nil {
		logrus.WithError(err).Errorf("invalid data for info")
		status := map[string]interface{}{"Status": "validation error. check log"}
		t, _ := template.ParseFiles("./views/admin-info.gtpl")
		t.Execute(w, status)
		//http.Redirect(w, r, "admin-info", http.StatusSeeOther)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	id, err := s.infoService.CreateNewInfo(infoForm)
	if err != nil {
		logrus.WithError(err).Error("error creating info")
		status := map[string]interface{}{"Status": "Did not create record. error. check log"}
		t, _ := template.ParseFiles("./views/admin-info.gtpl")
		t.Execute(w, status)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	logrus.WithField("id", id).Info("created info")

	status := map[string]interface{}{"Status": fmt.Sprintf("SUCCESS. ID=%d", id)}
	t, _ := template.ParseFiles("./views/admin-info.gtpl")
	t.Execute(w, status)
	w.WriteHeader(http.StatusCreated)
}
