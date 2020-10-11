package api

import (
	"fmt"
	"html/template"
	"net/http"
)

func (s Server) admin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method Nandeshar..:", r.Method) //get request method
	session, _ := s.sessionCookieStore.Get(r, "cookie-name")

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		fmt.Println("Error......session.....")
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	if r.Method == "GET" {
		t, _ := template.ParseFiles("./views/admin.gtpl")
		t.Execute(w, nil)
		return
	}
}
