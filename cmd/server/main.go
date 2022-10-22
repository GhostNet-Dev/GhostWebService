package main

import (
	"log"
	"net/http"
	"html/template"
	"encoding/json"
	"io/ioutil"
)

type TestCase struct {
	Title string
}

func cssHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("bootstrap.css")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, nil)
}

func jsHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("bootstrap.js")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadFile("./result.json")
	if err != nil {
		log.Fatal(err)
	}
	var tc []TestCase
	json.Unmarshal(data, &tc)
	t.Execute(w, tc)
}

func main() {
    http.HandleFunc("/", indexHandler)
	http.HandleFunc("/bootstrap.css", cssHandler)
	http.HandleFunc("/bootstrap.js", jsHandler)

	log.Fatal(http.ListenAndServe(":9123", nil))
}