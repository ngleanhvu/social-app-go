package main

import (
	"crud-go/component/appctx"
	"crud-go/component/uploadprovider"
	"crud-go/middleware"
	ginrestaurant2 "crud-go/module/restaurant/transport/ginrestaurant"
	"crud-go/module/restaurantlike/transport/ginrestaurantlike"
	"crud-go/module/user/transport/ginuser"
	"crud-go/pubsub/pblocal"
	subscriber2 "crud-go/subscriber"
	"fmt"
	"log"
	"os"

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
	pb := pblocal.NewLocalPubSub()
	if err != nil {
		log.Fatalln(fmt.Errorf("mysql connect err: %v", err))
	}
	fmt.Printf("Mysql connect success\n")

	s3BucketName := os.Getenv("S3_BUCKET_NAME")
	s3Region := os.Getenv("S3_REGION")
	s3ApiKey := os.Getenv("S3_API_KEY")
	s3Secret := os.Getenv("S3_SECRET")
	s3Domain := os.Getenv("S3_DOMAIN")

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3ApiKey, s3Secret, s3Domain)

	appContext := appctx.NewAppContext(db, s3Provider, pb)

	// setup subscriber
	//subscriber.Setup(appContext, context.Background())
	subscriber := subscriber2.NewEngine(appContext)
	subscriber.Start()

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
	v1.POST("register", ginuser.Register(appContext))
	v1.POST("login", ginuser.Login(appContext))
	v1.GET("profile", middleware.RequireAuth(appContext), ginuser.Profile(appContext))
	v1.POST("/:id/like",
		middleware.RequireAuth(appContext),
		ginrestaurantlike.UserLikeRestaurant(appContext),
	)
	v1.POST("/:id/dislike",
		middleware.RequireAuth(appContext),
		ginrestaurantlike.UserDislikeRestaurant(appContext),
	)
	v1.GET("/:id/user-liked", ginrestaurantlike.ListUserLikeRestaurant(appContext))

	r.Run()
}
