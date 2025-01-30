package main

import (
	"dbm_proj/internal/handlers"
	"fmt"
	"net/http"
)

func main() {
	app := handlers.NewApp()
	app.Routes()

	port := ":8080"

	fmt.Printf("Starting server on http://localhost%s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

//execute go/template
//net/http
