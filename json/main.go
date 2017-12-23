package main

import ...

type person struct {
  First string
  Last string
}

func init() {
  tpl = template.Must(template.ParseGlob("templates/*"))
}
var tpl *template.Template

func main() {
  http.HandleFunc("/", index)
  http.HandleFunc("/send", send)
  http.HandleFunc("/catch", catch)
  http.ListenAndServe(":8080", nil)
}
func index(w http.ResponseWriter, r *http.Request) {
  p1 := person {
    First: "AAA",
    Last: "BBB",
  }
  p2 := person {
    First: "CCC",
    Last: "DDD",
  }
  xp := []person{p1, p2}

  err := json.NewEncoder(w),Encode(xp)
  if err != nil {
    fmt.Println(err)
  }
}


func send(w http.ResponseWriter, r *http.Request) {
  tpl.ExecuteTemplate(w, "send.gohtml", nil)
}

var xp []person

func catch(w http.ResponseWriter, r *http.Request){
  if r.Method == http.MethodPost {
    err := json.NewDecoder(r.Body)
    if err != nil {
      fmt.Println(err)
    }
    fmt.Println(xp)
  }

  if r.Method == http.MethodGet {
    tpl.ExecuteTemplate(w, "catch.gohotml", xp)
  }
}
