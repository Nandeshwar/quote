package api

import (
	"fmt"
	"net/http"
)

func adminInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method Nandeshar..:", r.Method) //get request method
	session, _ := store.Get(r, "cookie-name")

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		fmt.Println("Error......session.....")
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	fmt.Println("*******************Inside Info")

	r.ParseForm()
	info := r.Form["info"]
	// logic part of log in
	fmt.Println("Info:", info)
}
