package controller

import (
	"database/sql"
	"go_go/dto"
	"go_go/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var cart_db *gorm.DB

func CartController(route *gin.Engine, db *gorm.DB) {
	cart_db = db
	routes := route.Group("/cart")
	{
		routes.GET("/:id", getCartbycustomer_id)
		routes.POST("/:id", addCart)
	}
}

func getCartbycustomer_id(c *gin.Context) {
	id := c.Param("id")
	customerID, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}

	customer := model.Customer{}
	cart_db.First(&customer, customerID)

	carts := []model.Cart{}
	cart_db.Where("customer_id = ?", customerID).Find(&carts)

	r_carts := []dto.R_Cart{}
	for _, cart := range carts {
		cartItems := []model.CartItem{}
		cart_db.Where("cart_id = ?", cart.CartID).Find(&cartItems)

		cartItemWithProducts := []dto.CartItemWithProduct{}
		for _, cartItem := range cartItems {
			product := model.Product{}
			cart_db.First(&product, cartItem.ProductID)
			itemTotal, _ := strconv.ParseFloat(product.Price, 64)
			cartItemWithProducts = append(cartItemWithProducts, dto.CartItemWithProduct{
				CartItemID:  cartItem.CartItemID,
				ProductID:   cartItem.ProductID,
				ProductName: product.ProductName,
				Description: product.Description,
				Price:       product.Price,
				Quantity:    cartItem.Quantity,
				ItemTotal:   itemTotal * float64(cartItem.Quantity),
			})
		}

		r_carts = append(r_carts, dto.R_Cart{
			CartID:     cart.CartID,
			CustomerID: cart.CustomerID,
			CartName:   cart.CartName,
			CreatedAt:  cart.CreatedAt.Time,
			UpdatedAt:  cart.UpdatedAt.Time,
			CartItem:   cartItemWithProducts,
		})
	}
	c.JSON(http.StatusOK, r_carts)
}

func addCart(c *gin.Context) {
	cart_item := dto.AddCartItem{}
	c.ShouldBindJSON(&cart_item)
	id := c.Param("id")

	cart_name := c.Query("name")
	idx, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}

	cart := model.Cart{}
	cart_db.Where("customer_id = ? AND cart_name = ?", idx, cart_name).First(&cart)
	if cart.CartID == 0 {
		cart.CustomerID = idx
		cart.CartName = cart_name
		cart.CreatedAt = sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}
		cart.UpdatedAt = sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}
		cart_db.Create(&cart)
	}
	item := model.CartItem{}
	product := model.Product{}
	cart_db.Where("cart_id = ? AND product_id = ?", cart.CartID, cart_item.ProductID).First(&item)
	if item.CartItemID == 0 {
		item.CartID = cart.CartID
		item.ProductID = cart_item.ProductID
		item.Quantity = cart_item.Quantity
		item.CreatedAt = sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}
		item.UpdatedAt = sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}
		cart_db.Create(&item)
	} else {
		item.Quantity += cart_item.Quantity
		item.UpdatedAt = sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}
		cart_db.Save(&item)
	}
	cart_db.Where("product_id = ?", cart_item.ProductID).First(&product)
	if product.StockQuantity < cart_item.Quantity {
		c.JSON(400, gin.H{"message": "Out of stock"})
		return
	}
	product.StockQuantity -= cart_item.Quantity
	cart_db.Save(&product)

	c.JSON(200, gin.H{"message": "Cart updated successfully"})
}
