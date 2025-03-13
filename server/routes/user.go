package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/restaurent_table_booking/models"
	"github.com/restaurent_table_booking/utils"
)

func Login(context *gin.Context) {
	var u models.Users
	err := context.ShouldBindBodyWithJSON(&u)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Can't read your input information"})
		return
	}
	err = u.Login()
	// panic(err)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Can't login"})
		return
	}
	token, err := utils.GenarateToken(u.Id, u.Email, u.Role)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"Message": "Can't login"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"Message": "Login successfully !!", "tokens": token, "role": u.Role})
}

func Register(context *gin.Context) {
	var u models.Users
	err := context.ShouldBindBodyWithJSON(&u)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Can't read your input information"})
		return
	}

	if u.Role == "user" {
		err = u.RegisterUser()
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Can't create account for user"})
			return
		}
	} else if u.Role == "admin" {
		err = u.RegisterAdmin()
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Can't create account for admin"})
			return
		}
	} else {
		err = u.RegisterStaff()
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Can't create accountfor staff"})
			return
		}
	}
	context.JSON(http.StatusCreated, gin.H{"Message": "Register successfully !!"})
}
