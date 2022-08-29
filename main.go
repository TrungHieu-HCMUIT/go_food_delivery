package main

import (
	"github.com/gin-gonic/gin"
	"go_restaurant/component"
	"go_restaurant/middleware"
	"go_restaurant/modules/restaurant/restauranttransport/ginrestaurant"
	"go_restaurant/modules/upload"
	"go_restaurant/modules/user/usertransport/ginuser"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	dsn := os.Getenv("DBConnectionStr")
	secretKey := os.Getenv("SystemSecretKey")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()

	if err := runService(db, secretKey); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB, secretKey string) error {
	r := gin.Default()

	appCtx := component.NewAppContext(db, secretKey)

	v1 := r.Group("/v1")

	v1.Use(middleware.Recover(appCtx))

	v1.POST("/upload", upload.Upload(appCtx))

	v1.POST("/register", ginuser.Register(appCtx))
	v1.POST("/login", ginuser.Login(appCtx))
	v1.GET("/profile", middleware.RequiredAuth(appCtx), ginuser.GetProfile(appCtx))

	restaurants := v1.Group("/restaurants")
	{
		restaurants.POST("", middleware.RequiredAuth(appCtx), ginrestaurant.CreateRestaurant(appCtx))
		restaurants.GET("/:id", middleware.RequiredAuth(appCtx), ginrestaurant.GetRestaurant(appCtx))
		restaurants.GET("", middleware.RequiredAuth(appCtx), ginrestaurant.ListRestaurant(appCtx))
		restaurants.PATCH("/:id", middleware.RequiredAuth(appCtx), ginrestaurant.UpdateRestaurant(appCtx))
		restaurants.DELETE("/:id", middleware.RequiredAuth(appCtx), ginrestaurant.DeleteRestaurant(appCtx))
	}

	return r.Run(":3000")
}
