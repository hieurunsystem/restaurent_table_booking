package models

import (
	"errors"

	"github.com/restaurent_table_booking/db"
)

type Restaurant struct {
	Id          int64
	Name        string
	Description string
	Admin_id    int
}

func GetAllRestaurants() ([]Restaurant, error) {
	var res []Restaurant
	query := `SELECT * FROM restaurants`
	rows, err := db.DB.Query(query)
	if err != nil {
		return res, errors.New("Can't catch any information")
	}
	defer rows.Close()

	for rows.Next() {
		var e Restaurant
		err = rows.Scan(&e.Id, &e.Name, &e.Description, &e.Admin_id)
		if err != nil {
			return res, errors.New("Can't catch any information")
		}
		res = append(res, e)
	}
	return res, nil
}

func (r *Restaurant) CreateRestaurant() error {
	query := `
	INSERT INTO restaurants(name, description, admin_id) 
	VALUES (?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic(err)
		return errors.New("Can't catch any information")
	}
	defer stmt.Close()
	result, err := stmt.Exec(r.Name, r.Description, r.Admin_id)
	if err != nil {
		// panic(err)
		// panic(r.name)
		return errors.New("Can't catch any information")
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
		return errors.New("Can't catch any information")
	}
	r.Id = id
	return nil
}
