package common

import "log"

const (
	DbTypeRestaurant = 1
	DbTypeUser       = 2
)

type Requester interface {
	GetUserId() int
}

const (
	CurrentUser = "user"
	SecretKey   = "12l3j12o31ij31oi3j12ij3i13jo1j"
)

func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Recover err:", err)
	}
}

const (
	TopicUserLikeRestaurant    = "TopicUserLikeRestaurant"
	TopicUserDislikeRestaurant = "TopicUserDislikeRestaurant"
)
