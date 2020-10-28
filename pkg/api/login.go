package api

import (
	"github.com/dgrijalva/jwt-go"
	"html/template"
	"net/http"
	"time"

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
		err = t.Execute(w, nil)
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

		expirationTime := time.Now().Add(5 * time.Minute)
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
