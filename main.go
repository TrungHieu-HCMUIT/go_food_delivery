package main

import (
	"github.com/gin-gonic/gin"
	"go_restaurant/component"
	"go_restaurant/modules/restaurant/restauranttransport/ginrestaurant"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	dsn := os.Getenv("DBConnectionStr")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	if err := runService(db); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB) error {
	r := gin.Default()

	appCtx := component.NewAppContext(db)

	restaurants := r.Group("/restaurants")
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))

		restaurants.GET("/:id", func(context *gin.Context) {
			id, err := strconv.Atoi(context.Param("id"))

			if err != nil {
				context.JSON(http.StatusUnauthorized, gin.H{
					"error": err.Error(),
				})

				return
			}

			var data Restaurant

			if err := db.Table(Restaurant{}.TableName()).
				Where("id = ?", id).
				First(&data).Error; err != nil {
				context.JSON(http.StatusUnauthorized, gin.H{
					"error": err.Error(),
				})

				return
			}

			context.JSON(http.StatusOK, data)
		})

		restaurants.GET("", func(context *gin.Context) {
			var data []Restaurant

			type Filter struct {
				CityId int `json:"city_id" form:"city_id"`
			}

			var filter Filter

			if err := context.ShouldBind(&filter); err != nil {
				context.JSON(http.StatusUnauthorized, gin.H{
					"error": err.Error(),
				})

				return
			}

			newDb := db
			if filter.CityId > 0 {
				newDb = db.Where("city_id = ?", filter.CityId)
			}

			if err := newDb.Find(&data).Error; err != nil {
				context.JSON(http.StatusUnauthorized, gin.H{
					"error": err.Error(),
				})

				return
			}

			context.JSON(http.StatusOK, data)
		})

		restaurants.PATCH("/:id", func(context *gin.Context) {
			id, err := strconv.Atoi(context.Param("id"))
			if err != nil {
				context.JSON(http.StatusUnauthorized, gin.H{
					"error": err.Error(),
				})

				return
			}

			var data RestaurantUpdate
			if err := context.ShouldBind(&data); err != nil {
				context.JSON(http.StatusUnauthorized, gin.H{
					"error": err.Error(),
				})
				return
			}

			if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
				context.JSON(http.StatusUnauthorized, gin.H{
					"error": err.Error(),
				})
				return
			}

			context.JSON(http.StatusOK, gin.H{"ok": 1})
		})

		restaurants.DELETE("", func(context *gin.Context) {
			if err := db.Table(Restaurant{}.TableName()).Where("status = 1").Delete(nil).Error; err != nil {
				context.JSON(http.StatusUnauthorized, gin.H{
					"error": err.Error(),
				})
				return
			}

			context.JSON(http.StatusOK, gin.H{"ok": 1})
		})
	}

	return r.Run()
}

type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"address" gorm:"column:addr;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"address" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantCreate struct {
	Id   int    `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"address" gorm:"column:addr;"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}
