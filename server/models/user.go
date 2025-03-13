package models

import (
	"database/sql"
	"errors"
	"log"

	"github.com/restaurent_table_booking/db"
	"github.com/restaurent_table_booking/utils"
)

type Users struct {
	Id        int64
	Name      string
	Email     string
	Phone     string
	Password  string
	Role      string
	orther_id int64
}

func (u *Users) RegisterUser() error {
	query := `INSERT INTO users(name,gmail, phone, password) 
		VALUES (?,?,?,?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic(err)
		return err
	}
	defer stmt.Close()
	hashPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		panic(err)
		return err
	}
	result, err := stmt.Exec(u.Name, u.Email, u.Phone, hashPassword)
	if err != nil {
		panic(err)
		return err
	}
	id, err := result.LastInsertId()
	u.Id = id
	return nil
}

func (u *Users) RegisterAdmin() error {
	query := `INSERT INTO admin(name, gmail, phone, password) 
		VALUES (?,?,?,?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic(err)
		return err
	}
	defer stmt.Close()
	hashPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		panic(err)
		return err
	}
	result, err := stmt.Exec(u.Name, u.Email, u.Phone, hashPassword)
	if err != nil {
		panic(err)
		return err
	}
	id, err := result.LastInsertId()
	u.Id = id
	return nil
}

func (u *Users) RegisterStaff() error {
	query := `INSERT INTO users(name, gmail, phone, password, restaurant_id) 
		VALUES (?,?,?,?,?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic(err)
		return err
	}
	defer stmt.Close()
	hashPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		panic(err)
		return err
	}
	result, err := stmt.Exec(u.Name, u.Email, u.Phone, hashPassword, u.orther_id)
	if err != nil {
		panic(err)
		return err
	}
	id, err := result.LastInsertId()
	u.Id = id
	return nil
}

func (u *Users) Login() error {
	var currentRole string
	cumtomersQuery := `
	SELECT id, password FROM users
	WHERE gmail = ?
	`
	rowCustomer := logRole(cumtomersQuery, u)

	staffsQuery := `
	SELECT id, password FROM staffs
	WHERE gmail = ?
	`
	rowStaff := logRole(staffsQuery, u)

	adminsQuery := `
	SELECT id, password FROM admin
	WHERE gmail = ?
	`
	rowAdmin := logRole(adminsQuery, u)

	var rows *sql.Row

	if rowCustomer != nil {
		rows = rowCustomer
		currentRole = "user"
	} else if rowStaff != nil {
		rows = rowStaff
		currentRole = "staff"
	} else if rowAdmin != nil {
		rows = rowAdmin
		currentRole = "admin"
	} else {
		return errors.New("Don't have anyone having this email")
	}

	log.Fatal(currentRole)

	var retrievedPassword string
	err := rows.Scan(&u.Id, retrievedPassword)
	if err != nil {
		panic(err)
		return err
	}
	ok := utils.PasswordVerify(u.Password, retrievedPassword)
	if !ok {
		return errors.New("Invalid Password!")
	}
	return nil
}

func logRole(query string, u *Users) *sql.Row {
	rows := db.DB.QueryRow(query, u.Email)
	return rows
}
