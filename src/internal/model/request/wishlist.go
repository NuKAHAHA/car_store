package request

type Wishlist struct {
	UserID string `json:"user_id" bson:"user_id"`

	CarID string `json:"car_id" bson:"car_id"`
}
