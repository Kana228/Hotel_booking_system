package models

import "database/sql"

type Admin struct {
	Phone    string
	Password string
	Email    string
	Name     string
	AdminID  int
}

type AdminModelInterface interface {
	Insert(phone string, password string, email string, name string, adminid int) (int, error)
	GetById(id int) (*Admin, error)
	GetByToken(token string) (*Admin, error)
	GetALlO([]*Admin, error)

	// UpdateRole(id int, role string) error dlya usera
}

type AdminModel struct {
	DB *sql.DB
}
