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
		restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))
		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
		restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))

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
