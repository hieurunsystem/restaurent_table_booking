package repositories

import (
	"fmt"

	"github.com/restaurent_table_booking/internal/db"
	"github.com/restaurent_table_booking/internal/models"
)

func CheckAccount(a *models.Account) (string, bool, error) {
	queries := map[string]string{
		"customer": "SELECT id, password FROM customers WHERE gmail = ?",
		"staff":    "SELECT id, password FROM staffs WHERE gmail = ?",
		"admin":    "SELECT id, password FROM admin WHERE gmail = ?",
		"owner":    "SELECT id, password FROM owners WHERE gmail = ?",
	}

	var retrievedPassword string
	for role, query := range queries {
		row := db.DB.QueryRow(query, a.Email)
		if err := row.Scan(&a.Id, &retrievedPassword); err == nil {
			a.Role = role
			return retrievedPassword, false, nil
		}
	}
	return "", true, nil
}

func InsertAccount(table string, values ...interface{}) (int64, error) {
	placeholders := ""
	for i := range values {
		if i > 0 {
			placeholders += ","
		}
		placeholders += "?"
	}

	query := fmt.Sprintf("INSERT INTO %s VALUES (%s)", table, placeholders)
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(values...)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func GetAllAccounts() ([]models.Account, error) {
	sqlQuery := `
		SELECT * FROM users
		UNION
		SELECT * FROM admin`

	rows, err := db.DB.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.Account
	for rows.Next() {
		var u models.Account
		err := rows.Scan(&u.Id, &u.Email, &u.Name, &u.Phone, &u.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
