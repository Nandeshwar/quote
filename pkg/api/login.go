package api

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"net/http"
)

func (s Server) login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./views/login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		user := r.Form["username"]
		password := r.Form[("password")]
		// logic part of log in
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])

		err := s.loginService.Login(user[0], password[0])
		if err != nil {
			fmt.Println("error=", err)
			http.Redirect(w, r, "login", http.StatusSeeOther)
		}

		session, _ := s.sessionCookieStore.Get(r, "cookie-name")

		session.Values["authenticated"] = true
		session.Save(r, w)

		t, _ := template.ParseFiles("./views/info2.gtpl")
		t.Execute(w, nil)
	}
}
