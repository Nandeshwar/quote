package api

import (
	"fmt"
	"net/http"
)

func (s Server) adminEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method

	session, _ := s.sessionCookieStore.Get(r, "cookie-name")

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	r.ParseForm()
	info := r.Form["event"]
	// logic part of log in
	fmt.Println("event:", r.Form["event"])
	fmt.Println("event:", info)
}
