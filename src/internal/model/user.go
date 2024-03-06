package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Mail     string             `gorm:"type:varchar(40);unique" json:"mail,omitempty"`
	Password string             `gorm:"size:255" json:"password,omitempty"`
	Name     string             `gorm:"size:100" json:"name"`
	Surname  string             `gorm:"size:100" json:"surname"`
	Birthday time.Time          `gorm:"type:date" json:"birthday"`
}
