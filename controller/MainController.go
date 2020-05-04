package controller

import (
	"Gimic/database"
	"Gimic/helpers"
	"Gimic/model"
	"github.com/gin-gonic/gin"
)

func GetAllUser(c *gin.Context) {
		var user []model.User
		db := database.Connect()
		err := db.Find(&user)

		if err != nil {
			helpers.RespondJSON(c,200,user)
		}

}
func CreateUser(c *gin.Context) {
	db := database.Connect()
	var user model.User
	c.Bind(&user)
	newUser := model.User{Fullname:user.Fullname,Email:user.Email,Phone:user.Phone}

	err1 := db.Create(&newUser)
	if err1 != nil {
		helpers.RespondJSON(c,200,user)
	}

}
func Halo(c *gin.Context) {
		helpers.RespondJSON(c,200,"Halo :)")

}