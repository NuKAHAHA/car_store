package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Car struct {
	//gorm.Model

	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID      string             `json:"user_id"`
	Model_Car   string             `gorm:"size:100" json:"model"`
	Year        int                `json:"year"`
	Cost        float64            `json:"cost"`
	Description string             `gorm:"size:255" json:"description"`
	Image       string             `gorm:"type:varchar(255)" json:"image"`
	Brand       string             `gorm:"size:100" json:"brand"`
}
