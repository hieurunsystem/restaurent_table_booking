package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/restaurent_table_booking/models"
)

func CreateRestaurant(context *gin.Context) {
	var r models.Restaurant
	err := context.ShouldBindBodyWithJSON(&r)
	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"message": "Can't take any input information"})
		return
	}
	err = r.CreateRestaurant()
	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"message": "Can't create restaurant"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Create succesfully", "restaurant": r})
}

func GetAllRestaurants(context *gin.Context) {
	res, err := models.GetAllRestaurants()
	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"message": "Can't take any restaurants"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"restaurants": res})
}
