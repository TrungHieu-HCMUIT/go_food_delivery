package restaurantlikemodel

import "time"

type RestaurantLike struct {
	RestaurantId int       `json:"restaurant_id" gorm:"restaurant_id;"`
	UserId       int       `json:"user_id" gorm:"user_id;"`
	CreatedAt    time.Time `json:"created_at" gorm:"created_at;"`
}

func (l RestaurantLike) TableName() string { return "restaurant_likes" }
