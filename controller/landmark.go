package controller

import (
	"go_go/dto"
	"go_go/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

var l_db *gorm.DB

func LandmarkController(route *gin.Engine, db *gorm.DB) {
	l_db = db
	routes := route.Group("/landmark")
	{
		routes.GET("/", getLandmark)
		routes.GET("/:id", getLandmarkbyid)
		routes.GET("/name", getLandmarkbyname)
		routes.POST("/", postLandmark)
		routes.PUT("/", putLandmark)
		routes.PUT("/put/", putsLandmark)
	}
}

func getLandmark(c *gin.Context) {
	landmarks := []model.Landmark{}
	l_db.Find(&landmarks)

	landmarks_dto := []dto.Landmark{}
	copier.Copy(&landmarks_dto, &landmarks)
	c.JSON(200, landmarks_dto)
}

func postLandmark(c *gin.Context) {
	landmarks_dto := []dto.Landmark{}
	c.ShouldBindJSON(&landmarks_dto)

	landmark := model.Landmark{}
	copier.Copy(&landmark, &landmarks_dto)

	l_db.Create(&landmark)
	c.JSON(200, "Success")
}

func getLandmarkbyid(c *gin.Context) {
	id := c.Param("id")
	idx, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	landmark := model.Landmark{}
	l_db.Preload("Countrydata").Where("Idx = ?", idx).Find(&landmark)
	// landmark_dto := dto.Landmark{}
	// copier.Copy(&landmark_dto, &landmark)
	c.JSON(200, landmark)
}
func getLandmarkbyname(c *gin.Context) {
	name := c.Query("name")
	landmark := []model.Landmark{}
	name = "%" + name + "%"
	l_db.Preload("Countrydata").Where("Name like ?", name).Find(&landmark)
	// landmark_dto := dto.Landmark{}
	// copier.Copy(&landmark_dto, &landmark)
	c.JSON(200, landmark)
}

func putLandmark(c *gin.Context) {
	landmark := []model.Landmark{}
	c.ShouldBindJSON(&landmark)
	for _, l := range landmark {
		l_db.Model(model.Landmark{}).Where("Idx = ?", l.Idx).Update("Name", "จองแล้วนะจ๊ะ")
	}
	c.JSON(200, "Success")
}

func putsLandmark(c *gin.Context) {
	_landmark := model.Landmark{}
	c.ShouldBindJSON(&_landmark)
	l_db.Model(model.Landmark{}).Where("Idx = ?", _landmark.Idx).Updates(&_landmark)
	c.JSON(200, "Success")
}
