package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	First string
	Last  string
	Items []string
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/mshl", mshl)
	http.HandleFunc("/encd", encd)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	s := `<!Doctype html>
        <html lang="en">
        <head>
        <meta charset="UTF-8>
        <title>Foo</title>
        </head>
        <body>
        your are at foo
        </body>
        </html>`
	w.Write([]byte(s))
}
func mshl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		"James",
		"Bond",
		[]string{"suit", "gun", "Wry sense of humor"},
	}
	json, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}

func encd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		"James",
		"Bond",
		[]string{"suit", "gun", "Wry sense of humor"},
	}
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println(err)
	}
}
