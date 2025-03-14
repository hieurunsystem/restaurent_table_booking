package models

import (
	"errors"
	"fmt"

	"github.com/restaurent_table_booking/db"
	"github.com/restaurent_table_booking/utils"
)

type Account struct {
	Id        int64
	Name      string
	Email     string
	Phone     string
	Password  string
	Role      string
	orther_id int64
}

func (u *Account) RegisterCustomer() error {
	_, check := checkAccount(u)
	if !check {
		return errors.New("This gmail already create account before!!")
	}
	query := `INSERT INTO customers(name,gmail, phone, password) 
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

func (u *Account) RegisterOwner() error {
	_, check := checkAccount(u)
	if !check {
		return errors.New("This gmail already create account before!!")
	}
	query := `INSERT INTO owners(name, gmail, phone, password) 
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

func (u *Account) RegisterAdmin() error {
	_, check := checkAccount(u)
	if !check {
		return errors.New("This gmail already create account before!!")
	}
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

func (u *Account) RegisterStaff() error {
	_, check := checkAccount(u)
	if !check {
		return errors.New("This gmail already create account before!!")
	}
	query := `INSERT INTO customers(name, gmail, phone, password, restaurant_id) 
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

func (u *Account) Login() error {
	retrievedPassword, _ := checkAccount(u)
	ok := utils.PasswordVerify(u.Password, retrievedPassword)
	if !ok {
		return errors.New("Invalid Password!")
	}
	return nil
}

func checkAccount(a *Account) (string, bool) {
	CumtomersQuery := `
	SELECT id, password FROM customers
	WHERE gmail = ?
	`
	rowCustomer := db.DB.QueryRow(CumtomersQuery, a.Email)

	staffsQuery := `
	SELECT id, password FROM staffs
	WHERE gmail = ?
	`
	rowStaff := db.DB.QueryRow(staffsQuery, a.Email)

	adminsQuery := `
	SELECT id, password FROM admin
	WHERE gmail = ?
	`
	rowAdmin := db.DB.QueryRow(adminsQuery, a.Email)

	ownersQuery := `
	SELECT id, password FROM owners
	WHERE gmail = ?
	`
	rowOwner := db.DB.QueryRow(ownersQuery, a.Email)

	var retrievedPassword string
	var err error
	err = rowAdmin.Scan(&a.Id, &retrievedPassword)
	if err == nil {
		a.Role = "admin"
		return retrievedPassword, false
	} else {
		err = rowStaff.Scan(&a.Id, &retrievedPassword)
		if err == nil {
			a.Role = "staff"
			return retrievedPassword, false
		} else {
			err = rowOwner.Scan(&a.Id, &retrievedPassword)
			if err == nil {
				a.Role = "owner"
				return retrievedPassword, false
			} else {
				err = rowCustomer.Scan(&a.Id, &retrievedPassword)
				if err == nil {
					a.Role = "customer"
					return retrievedPassword, false
				} else {
					fmt.Printf("Don't have any account like this")
					return "", true
				}
			}
		}
	}
	return "", false
}
