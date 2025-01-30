package handlers

import "net/http"

func (app *Application) Routes() {
	http.HandleFunc("/", app.home)
	http.HandleFunc("/register", app.register)
}
