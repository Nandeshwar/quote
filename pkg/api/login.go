package api

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
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

		db, err := sql.Open("sqlite3", "./db/quote.db")
		if err != nil {
			fmt.Println("Database connection error.....")
		}

		rows, err := db.Query("SELECT user, password FROM login")
		if err != nil {
			fmt.Println("Error in selecting login")
		}

		var dbUser string
		var dbPassword string

		for rows.Next() {
			err = rows.Scan(&dbUser, &dbPassword)
			if err != nil {
				fmt.Println("Db error in select statement", err)
			}

			fmt.Printf("\n Db user=%v, password=%v", dbUser, dbPassword)
		}

		if user[0] == "Radha" && password[0] == "Krishna" {
			session, _ := store.Get(r, "cookie-name")

			session.Values["authenticated"] = true
			session.Save(r, w)

			t, _ := template.ParseFiles("./views/info2.gtpl")
			t.Execute(w, nil)
		} else {
			http.Redirect(w, r, "login", http.StatusSeeOther)
		}
	}
}
