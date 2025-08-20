package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/restaurent_table_booking/internal/config"
)

var DB *sql.DB

func InitDB() {
	config, err := config.LoadConfig()
	if err != nil {
		return
	}

	uri := config.GetDatabaseUri()

	DB, err = sql.Open("mysql", uri)
	if err != nil {
		fmt.Println("Cannotconnect to database")
		return
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	err = createTable()
	if err != nil {
		panic(err)
	}
}

func createTable() error {
	CustomerQuery := `
	CREATE TABLE IF NOT EXISTS customers (
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		gmail VARCHAR(50) NOT NULL UNIQUE,
		name NVARCHAR(50) NOT NULL, 
		phone VARCHAR(50) NOT NULL,
		password VARCHAR(250) NOT NULL
	)	
	`
	_, err := DB.Exec(CustomerQuery)
	if err != nil {
		return err
	}

	AdminQuery := `
	CREATE TABLE IF NOT EXISTS admin (
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		gmail VARCHAR(50) NOT NULL UNIQUE,
		name NVARCHAR(50) NOT NULL, 
		phone VARCHAR(50) NOT NULL,
		password VARCHAR(250) NOT NULL
	)	
	`
	_, err = DB.Exec(AdminQuery)
	if err != nil {
		return err
	}

	OwnerQuery := `
	CREATE TABLE IF NOT EXISTS owners (
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		gmail VARCHAR(50) NOT NULL UNIQUE,
		name NVARCHAR(50) NOT NULL, 
		phone VARCHAR(50) NOT NULL,
		password VARCHAR(250) NOT NULL
	)	
	`
	_, err = DB.Exec(OwnerQuery)
	if err != nil {
		return err
	}

	RestaurantQuery := `
	CREATE TABLE IF NOT EXISTS restaurants (
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		name NVARCHAR(50) NOT NULL, 
		description VARCHAR(250) NOT NULL,
		time_start TIME NOT NULL,
		time_end TIME NOT NULL,
		owner_id INTEGER NOT NULL,
		FOREIGN KEY (owner_id) REFERENCES owners(id)
	)	
	`
	_, err = DB.Exec(RestaurantQuery)
	if err != nil {
		return err
	}

	StaffQuery := `
	CREATE TABLE IF NOT EXISTS staffs (
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		gmail VARCHAR(50) NOT NULL UNIQUE,
		name NVARCHAR(50) NOT NULL, 
		phone VARCHAR(50) NOT NULL,
		password VARCHAR(250) NOT NULL,
		restaurant_id INTEGER NOT NULL,
		FOREIGN KEY (restaurant_id) REFERENCES restaurants(id)
	)`
	_, err = DB.Exec(StaffQuery)
	if err != nil {
		return err
	}

	TableQuery := `
	CREATE TABLE IF NOT EXISTS tables (
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		name NVARCHAR(50) NOT NULL, 
		type INTEGER NOT NULL,
		seats NVARCHAR(50) NOT NULL,
		restaurant_id INTEGER NOT NULL,
		FOREIGN KEY (restaurant_id) REFERENCES restaurants(id)
	)	
	`
	_, err = DB.Exec(TableQuery)
	if err != nil {
		return err
	}

	StatusQuery := `
	CREATE TABLE IF NOT EXISTS status (
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		name NVARCHAR(50) NOT NULL
	)	
	`
	_, err = DB.Exec(StatusQuery)
	if err != nil {
		return err
	}

	ReservationQuery := `
	CREATE TABLE IF NOT EXISTS reservations (
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		numberOfCustomer VARCHAR(50) NOT NULL,
		book_date DATE NOT NULL, 
		time_start TIME NOT NULL,
		time_end TIME NOT NULL,
		actual_end TIME NOT NULL,
		price FLOAT NOT NULL,
		customer_email VARCHAR(50) NOT NULL,
		table_id INTEGER NOT NULL,
		FOREIGN KEY (table_id) REFERENCES tables(id),
		customer_id INTEGER NOT NULL,
		FOREIGN KEY (customer_id) REFERENCES customers(id),
		status_id INTEGER NOT NULL,
		FOREIGN KEY (status_id) REFERENCES status(id)
	)	
	`
	_, err = DB.Exec(ReservationQuery)
	if err != nil {
		return err
	}

	return nil
}
