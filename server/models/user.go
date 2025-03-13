package models

import (
	"errors"

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
	cumtomersQuery := `
	SELECT id, password FROM users
	WHERE gmail = ?
	`
	rowCustomer := db.DB.QueryRow(cumtomersQuery, u.Email)

	staffsQuery := `
	SELECT id, password FROM staffs
	WHERE gmail = ?
	`
	rowStaff := db.DB.QueryRow(staffsQuery, u.Email)

	adminsQuery := `
	SELECT id, password FROM admin
	WHERE gmail = ?
	`
	rowAdmin := db.DB.QueryRow(adminsQuery, u.Email)

	var retrievedPassword string
	var err error
	err = rowAdmin.Scan(&u.Id, &retrievedPassword)
	if err == nil {
		u.Role = "admin"
	} else {
		err = rowStaff.Scan(&u.Id, &retrievedPassword)
		if err == nil {
			u.Role = "staff"
		} else {
			err = rowCustomer.Scan(&u.Id, &retrievedPassword)
			if err == nil {
				u.Role = "user"
			} else {
				return errors.New("Don't have anyone having this email")
			}
		}
	}
	ok := utils.PasswordVerify(u.Password, retrievedPassword)
	if !ok {
		return errors.New("Invalid Password!")
	}
	return nil
}
