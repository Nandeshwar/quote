package api

import (
	"fmt"
	"html/template"
	"net/http"
	"quote/pkg/model"
)

func (s Server) adminEvent(w http.ResponseWriter, r *http.Request) {
	session, _ := s.sessionCookieStore.Get(r, "cookie-name")

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		fmt.Println("Error......session.....")
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	if r.Method == "GET" {
		t, _ := template.ParseFiles("./views/admin-event-detail.gtpl")
		t.Execute(w, nil)
		return
	}

	r.ParseForm()
	eventDetailForm := model.EventDetailForm{
		Title:     r.Form["title"][0],
		Info:      r.Form["info"][0],
		EventDate: r.Form["eventDate"][0],
		Typ:       r.Form["eventType"][0],
		Link:      r.Form["link"][0],
		CreatedAt: r.Form["createdAt"][0],
	}

	err := s.eventDetailService.ValidateFormEvent(eventDetailForm)
	if err != nil {
		fmt.Println(err)
		status := map[string]interface{}{"Status": "validation error. check log"}
		t, _ := template.ParseFiles("./views/admin-event-detail.gtpl")
		t.Execute(w, status)
		return
	}

	id, err := s.eventDetailService.CreateNewEventDetail(eventDetailForm)
	if err != nil {
		fmt.Println(err)
		status := map[string]interface{}{"Status": "Did not create record. error. check log"}
		t, _ := template.ParseFiles("./views/admin-event-detail.gtpl")
		t.Execute(w, status)
		return
	}

	status := map[string]interface{}{"Status": fmt.Sprintf("SUCCESS. ID=%d", id)}
	t, _ := template.ParseFiles("./views/admin-event-detail.gtpl")
	t.Execute(w, status)
}
