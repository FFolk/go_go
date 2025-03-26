package controller

import (
	"go_go/dto"
	"go_go/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

var p_db *gorm.DB

func ProductController(route *gin.Engine, db *gorm.DB) {
	p_db = db
	routes := route.Group("/product")
	{
		routes.POST("/search", searchProduct)
	}
}

func searchProduct(c *gin.Context) {
	s_product := dto.SearchProduct{}
	c.ShouldBindJSON(&s_product)
	products := []model.Product{}
	if s_product.Pricemin == "" {
		s_product.Pricemin = "0"
	}
	if s_product.Pricemax == "" {
		s_product.Pricemax = "99999999"
	}
	p_db.Where("price BETWEEN ? AND ? AND description LIKE ?", s_product.Pricemin, s_product.Pricemax, "%"+s_product.Description+"%").Find(&products)
	products_dto := []dto.Product{}
	copier.Copy(&products_dto, &products)
	c.JSON(http.StatusOK, products_dto)
}
