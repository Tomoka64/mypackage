package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/abundance", abundance)
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "some value",
	})
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
}

func read(w http.ResponseWriter, r *http.Request) {
	c1, err := r.Cookie("my-cookie")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #1:", c1)
	}
	c2, err := r.Cookie("general")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #2", c2)
	}
	c3, err := r.Cookie("specific")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #3", c3)
	}
}

func abundance(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "general",
		Value: "general values",
	})

	http.SetCookie(w, &http.Cookie{
		Name:  "specific",
		Value: "some specific value",
	})

	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
}
