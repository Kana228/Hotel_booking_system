package handlers

import (
	"net/http"
	"text/template"
)

func (app *Application) register(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("ui/html/register.html")
	_ = t.Execute(w, nil)
}
