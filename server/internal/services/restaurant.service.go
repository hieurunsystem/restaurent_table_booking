package services

import (
	"errors"
	"fmt"

	"github.com/restaurent_table_booking/internal/db"
	"github.com/restaurent_table_booking/internal/models"
)

func GetAllRestaurants() ([]models.Restaurant, error) {
	var res []models.Restaurant
	query := `SELECT * FROM restaurants`
	rows, err := db.DB.Query(query)
	if err != nil {
		return res, errors.New("can't catch any information")
	}
	defer rows.Close()

	for rows.Next() {
		var e models.Restaurant
		err = rows.Scan(&e.Id, &e.Name, &e.Description, &e.Started, &e.Ended, &e.Owner_id)
		if err != nil {
			return res, errors.New("can't catch any information")
		}
		res = append(res, e)
	}
	return res, nil
}

func CreateRestaurant(r *models.Restaurant) error {
	query := `
	INSERT INTO restaurants(name, description, time_start, time_end, owner_id) 
	VALUES (?, ?, ?,?,?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return errors.New("can't catch any information")
	}
	defer stmt.Close()
	result, err := stmt.Exec(r.Name, r.Description, r.Started, r.Ended, r.Owner_id)
	fmt.Print(r.Owner_id)
	if err != nil {
		// panic(r.name)
		return errors.New("can't catch any information")
	}
	id, err := result.LastInsertId()
	if err != nil {
		return errors.New("can't catch any information")
	}
	r.Id = id
	return nil
}
