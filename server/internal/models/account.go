package models

type Account struct {
	Id        int64
	Name      string
	Email     string
	Phone     string
	Password  string
	Role      string
	Orther_id int64
}
