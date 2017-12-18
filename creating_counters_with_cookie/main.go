package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", tomoka)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func tomoka(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("my_cookie")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "my_cookie",
			Value: "0",
		}
	}
	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatalln(err)
	}
	count += 1
	cookie.Value = strconv.Itoa(count)

	http.SetCookie(w, cookie)
	io.WriteString(w, cookie.Value)
}
