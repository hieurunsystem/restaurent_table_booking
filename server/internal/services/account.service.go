package services

import (
	"errors"

	"github.com/restaurent_table_booking/internal/models"
	"github.com/restaurent_table_booking/internal/repositories"
	"github.com/restaurent_table_booking/internal/utils"
)

func RegisterCustomer(u *models.Account) error {
	_, check, _ := repositories.CheckAccount(u)
	if !check {
		return errors.New("this gmail already created account before")
	}

	hashPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	id, err := repositories.InsertAccount("customers(name, gmail, phone, password, status)",
		u.Name, u.Email, u.Phone, hashPassword, "active")
	if err != nil {
		return err
	}

	u.Id = id
	return nil
}

func RegisterOwner(u *models.Account) error {
	_, check, _ := repositories.CheckAccount(u)
	if !check {
		return errors.New("this gmail already created account before")
	}

	hashPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	id, err := repositories.InsertAccount("owners(name, gmail, phone, password)",
		u.Name, u.Email, u.Phone, hashPassword)
	if err != nil {
		return err
	}

	u.Id = id
	return nil
}

func RegisterAdmin(u *models.Account) error {
	_, check, _ := repositories.CheckAccount(u)
	if !check {
		return errors.New("this gmail already created account before")
	}

	hashPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	id, err := repositories.InsertAccount("admin(name, gmail, phone, password)",
		u.Name, u.Email, u.Phone, hashPassword)
	if err != nil {
		return err
	}

	u.Id = id
	return nil
}

func RegisterStaff(u *models.Account) error {
	_, check, _ := repositories.CheckAccount(u)
	if !check {
		return errors.New("this gmail already created account before")
	}

	hashPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	id, err := repositories.InsertAccount("staffs(name, gmail, phone, password, restaurant_id)",
		u.Name, u.Email, u.Phone, hashPassword, u.Orther_id)
	if err != nil {
		return err
	}

	u.Id = id
	return nil
}

func Login(u *models.Account) error {
	retrievedPassword, _, _ := repositories.CheckAccount(u)
	ok := utils.PasswordVerify(u.Password, retrievedPassword)
	if !ok {
		return errors.New("invalid password")
	}
	return nil
}

func GetAllAccounts() ([]models.Account, error) {
	return repositories.GetAllAccounts()
}
