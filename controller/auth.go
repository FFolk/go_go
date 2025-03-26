package controller

import (
	"go_go/dto"
	"go_go/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var a_db *gorm.DB

func AuthController(route *gin.Engine, db *gorm.DB) {
	a_db = db
	routes := route.Group("/auth")
	{
		routes.POST("/login", login)
		routes.PUT("/password", changePassword)
	}
}

func login(c *gin.Context) {
	_customer := model.Customer{}
	customer := model.Customer{}

	c.ShouldBindJSON(&customer)

	a_db.Where("email = ? ", customer.Email).Find(&_customer)

	err := bcrypt.CompareHashAndPassword([]byte(_customer.Password), []byte(customer.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	customer_dto := dto.Customer{}
	copier.Copy(&customer_dto, &_customer)
	c.JSON(http.StatusOK, customer_dto)
}

func changePassword(c *gin.Context) {
	newpassword := dto.NewPassword{}
	customer := model.Customer{}
	c.ShouldBindJSON(&newpassword)
	a_db.Where("email = ? ", newpassword.Email).Find(&customer)

	err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(newpassword.OldPassword))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(newpassword.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	a_db.Model(model.Customer{}).Where("email = ?", newpassword.Email).Update("password", string(hash))
	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}
