package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/mattn/go-sqlite3"

	"encoding/json"
)

type Page struct {
	Name     string
	DbStatus bool
}

type SearchResult struct {
	Title  string
	Author string
	Year   string
	Id     string
}

func main() {
	templates := template.Must(template.ParseFiles("templates/index.html"))
	db, _ := sql.Open("sqlite3", "/home/aufather/dev.db")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := Page{Name: "Gopher"}
		if name := r.FormValue("name"); name != "" {
			p.Name = name
		}
		p.DbStatus = db.Ping() == nil
		if err := templates.ExecuteTemplate(w, "index.html", p); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		results := []SearchResult{
			SearchResult{"Hello, World", "K&R", "1971", "1234"},
			SearchResult{"Hello, Sam", "A&A", "2015", "1111"},
			SearchResult{"Hello, Ben", "A&A", "2019", "2222"},
		}

		encoder := json.NewEncoder(w)
		if err := encoder.Encode(results); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	fmt.Println(http.ListenAndServe(":8080", nil))
}
