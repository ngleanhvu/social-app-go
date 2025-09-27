package main

import (
	"crud-go/component/appctx"
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

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// create new restaurant

	appContext := appctx.NewAppContext(db)

	v1 := r.Group("/api/v1/restaurant")

	v1.POST("", ginrestaurant2.CreateRestaurant(appContext))
	v1.DELETE("/:id", ginrestaurant2.DeleteRestaurant(appContext))
	v1.PUT("/:id", ginrestaurant2.UpdateRestaurant(appContext))
	v1.GET("", ginrestaurant2.ListRestaurant(appContext))

	//v1.GET("/restaurant/:id", func(c *gin.Context) {
	//
	//	id, err := strconv.Atoi(c.Param("id"))
	//
	//	if err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//		return
	//	}
	//
	//	var data Restaurant
	//
	//	db.Where("id = ?", id).First(&data)
	//
	//	c.JSON(200, gin.H{
	//		"data": data,
	//	})
	//})
	//
	//v1.GET("/restaurant", func(c *gin.Context) {
	//	var data []Restaurant
	//
	//	type Paging struct {
	//		Page  int `json:"page" form:"page"`
	//		Limit int `json:"limit" form:"limit"`
	//	}
	//
	//	var pagingData Paging
	//
	//	if err := c.ShouldBind(&pagingData); err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//		return
	//	}
	//
	//	if pagingData.Page <= 0 {
	//		pagingData.Page = 1
	//	}
	//
	//	if pagingData.Limit <= 0 {
	//		pagingData.Limit = 10
	//	}
	//
	//	db.Offset((pagingData.Page - 1) * pagingData.Limit).Order("id desc").Limit(pagingData.Limit).Find(&data)
	//
	//	c.JSON(200, gin.H{
	//		"data": data,
	//	})
	//})

	//v1.PUT("/restaurant/:id", func(c *gin.Context) {
	//	id, err := strconv.Atoi(c.Param("id"))
	//
	//	if err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//		return
	//	}
	//
	//	var data RestaurantUpdate
	//
	//	if err := c.ShouldBind(&data); err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	}
	//
	//	db.Where("id = ?", id).Updates(&data)
	//
	//	c.JSON(200, gin.H{
	//		"data": data,
	//	})
	//
	//})

	//v1.DELETE("/restaurant/:id", func(c *gin.Context) {
	//	id, err := strconv.Atoi(c.Param("id"))
	//
	//	if err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//		return
	//	}
	//
	//	db.Where("id = ?", id).Delete(&Restaurant{})
	//
	//	c.JSON(200, gin.H{
	//		"message": "delete success",
	//	})
	//})

	r.Run()

	//var myRestaurant Restaurant
	//
	//if err := db.Where("id = ?", 3).First(&myRestaurant).Error; err != nil {
	//	log.Fatalln(err.Error())
	//}
	//
	//newName := "200Lab"
	//updateData := RestaurantUpdate{Name: &newName}
	//
	//if err := db.Where("id = ?", 3).Updates(&updateData).Error; err != nil {
	//	log.Fatalln(err.Error())
	//}

	//if err := db.Table(myRestaurant.TableName()).Where("id = ?", 2).Delete(nil).Error; err != nil {
	//	log.Fatalln(err.Error())
	//}
}
