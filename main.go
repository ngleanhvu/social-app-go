package main

import (
	"crud-go/component/appctx"
	"crud-go/middleware"
	ginrestaurant2 "crud-go/module/restaurant/transport/ginrestaurant"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
	Addr string `json:"addr" gorm:"column:addr"`
}

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name"`
	Addr *string `json:"addr" gorm:"column:addr"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

func main() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/food_delivery_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db = db.Debug()
	if err != nil {
		log.Fatalln(fmt.Errorf("mysql connect err: %v", err))
	}
	fmt.Printf("Mysql connect success\n")

	appContext := appctx.NewAppContext(db)

	r := gin.Default()
	r.Use(middleware.Recover(appContext))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/api/v1/restaurant")

	v1.POST("", ginrestaurant2.CreateRestaurant(appContext))
	v1.DELETE("/:id", ginrestaurant2.DeleteRestaurant(appContext))
	v1.PUT("/:id", ginrestaurant2.UpdateRestaurant(appContext))
	v1.GET("", ginrestaurant2.ListRestaurant(appContext))

	r.Run()
}
