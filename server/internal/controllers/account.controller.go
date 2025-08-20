package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/restaurent_table_booking/internal/models"
	"github.com/restaurent_table_booking/internal/services"
	"github.com/restaurent_table_booking/internal/utils"
)

func RegisterHandler(context *gin.Context) {
	var u models.Account
	err := context.ShouldBindJSON(&u)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Can't read your input information"})
		return
	}

	switch u.Role {
	case "customer":
		err = services.RegisterCustomer(&u)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
	case "admin":
		err = services.RegisterAdmin(&u)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
	case "owner":
		err = services.RegisterOwner(&u)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
	case "staff":
		err = services.RegisterStaff(&u)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
	default:
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid role"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"Message": "Register successfully"})
}

func LoginHandler(context *gin.Context) {
	var u models.Account
	err := context.ShouldBindBodyWithJSON(&u)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Can't read your input information"})
		return
	}
	err = services.Login(&u)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Can't login"})
		return
	}
	// create token
	token, err := utils.GenerateToken(u.Id, u.Email, u.Role)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"Message": "Can't generate token"})
		return
	}

	// save token into cookie
	context.SetCookie("token", token, 7200, "/", "localhost", false, true)

	context.JSON(http.StatusOK, gin.H{"Message": "Login successfully !!", "tokens": token, "role": u.Role})
}

func GetAllAccountsHandler(context *gin.Context) {
	u, err := services.GetAllAccounts()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"Message": err.Error()})
	}
	context.JSON(http.StatusOK, gin.H{"users": u})
}

func LogoutHandler(context *gin.Context) {
	// Xóa cookie bằng cách đặt giá trị ròng và thời gian hết hazole
	context.SetCookie("token", "", -1, "/", "localhost", false, true)

	// Trả về phản hồi JSON
	context.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

func GetProfile(context *gin.Context) {
	role, exists := context.Get("role")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"role": role})
}
