package model

type Wishlist struct {
	UserID string `json:"user_id"`
	User   User   `gorm:"foreignKey:UserID" json:"user,omitempty"`

	CarID string `json:"car_id"`
	Car   Car    `gorm:"foreignKey:CarID" json:"car,omitempty"`
}
