package request

type WishlistRequest struct {
	UserID string `json:"user_id"`

	CarID string `json:"car_id"`
}
