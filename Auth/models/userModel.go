package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	Id            primitive.ObjectID `bson:"id"`
	FirstName     *string            `json:"first_name" validate: "required, min=2, max=100"`
	LastName      *string            `json:"last_name" validate: "required, min=2, max=100"`
	Password      *string            `json:"password" validate: "required, min=6, max=12"`
	Email         *string            `json:"email" validate: "email, min=6, max=12"`
	Phone         *string            `json:"phone" validate: "required, min=10, max=12"`
	Token         *string            `json:"token"`
	User_type     *string            `json:"user_type" validate: "required, eq=ADMIN | eq=USER"`
	Refresh_token *string            `json:"refresh_Token"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
	User_Id       string             `json:"user_id"`
}

//min, max= character count validation
