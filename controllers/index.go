package controllers

import (
	"net/http"

	"github.com/GURUAKASH-MUTHURAJAN/swagger/model"
	"github.com/GURUAKASH-MUTHURAJAN/swagger/service"
	"github.com/gin-gonic/gin"
)

// GetUsers return list of all user fom the database
// @Summary return list of all
// @Description return list of all user fom the database
// @Tags Users
// @Success 200 {object} model.UserResponse
// @Router /users [get]
func GetUsers(c *gin.Context) {
	users := service.ListUser()
	c.JSON(http.StatusOK, model.UserResponse{Data: users})
}
// CreateUser create new user
// @Summary return created user
// @Description create and return user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body model.User true "User"
// @Success 200 {object} model.User
// @Failure 400 {object} model.ErrorResponse
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}
	if err := service.CreateUser(req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated,req)
}
