package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
)

type Page struct {
	Name     string
	DbStatus bool
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
		db.Close()
	})
	fmt.Println(http.ListenAndServe(":8080", nil))
}
