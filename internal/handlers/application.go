//insert all interfaces from models here

package handlers

import (
	"dbm_proj/models"
)

type Application struct {
	Admin models.AdminModelInterface
}

func NewApp() *Application {
	var app *Application
	return app
}

//func (app Application) home(w http.ResponseWriter, r http.Request){
//app.Admin.GetALlO()
//}
