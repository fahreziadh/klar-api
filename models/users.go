package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	UserID       bson.ObjectId `bson:"_id" json:"id"`
	UserName     string        `bson:"username" json:"username"`
	FullName     string        `bson:"fullname" json:"fullname"`
	Email        string        `bson:"email" json:"email"`
	Password     string        `bson:"password" json:"password"`
	ImageProfile string        `bson:"imageprofile" json:"imageprofile"`
	CreatedAt    time.Time     `bson:"created_at" json:"created_at"`
}

type Teacher struct {
	UserID       string `bson:"userid" json:"userid"`
	UserName     string `bson:"username" json:"username"`
	FullName     string `bson:"fullname" json:"fullname"`
	ImageProfile string `bson:"imageprofile" json:"imageprofile"`
}

type Student struct {
	ClassID      string    `bson:"classid" json:"classid"`
	UserName     string    `bson:"username" json:"username"`
	FullName     string    `bson:"fullname" json:"fullname"`
	ImageProfile string    `bson:"imageprofile" json:"imageprofile"`
	CreatedAt    time.Time `bson:"created_at" json:"created_at"`
}

type AuthRequest struct {
	UserName string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
}

type AuthResponse struct {
	AccessToken string `bson:"accesstoken" json:"accesstoken"`
	User        User   `bson:"user" json:"user"`
}
