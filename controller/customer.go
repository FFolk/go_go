package controller

import (
	"go_go/dto"
	"go_go/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

var c_db *gorm.DB

func CustomerController(route *gin.Engine, db *gorm.DB) {
	c_db = db
	routes := route.Group("/customer")
	{
		routes.GET("/:id", getCustomer)
		routes.PUT("/address/:id", updatedaddress)

	}
}

func getCustomer(c *gin.Context) {
	id := c.Param("id")
	idx, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	customer := model.Customer{}
	a_db.Where("customer_id = ? ", idx).Find(&customer)
	customer_dto := dto.Customer{}
	copier.Copy(&customer_dto, &customer)
	c.JSON(http.StatusOK, customer_dto)
}

func updatedaddress(c *gin.Context) {
	id := c.Param("id")
	idx, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}

	address := dto.Customer{}
	c.ShouldBindJSON(&address)

	c_db.Model(model.Customer{}).Where("customer_id = ?", idx).Update("address", address.Address)
	c.JSON(http.StatusOK, gin.H{"message": "Address updated successfully"})
}
