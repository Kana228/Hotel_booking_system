package handlers

import (
	"net/http"
	"text/template"
)

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("ui/html/home.html")
	_ = t.Execute(w, nil)
}
