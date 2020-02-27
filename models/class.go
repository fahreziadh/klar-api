package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Class struct {
	ID             bson.ObjectId `bson:"_id" json:"id"`
	ClassName      string        `bson:"classname" json:"classname"`
	ClassID        string        `bson:"classid" json:"classid"`
	Major          string        `bson:"major" json:"major"`
	Grade          string        `bson:"grade" json:"grade"`
	Location       string        `bson:"location" json:"location"`
	LocationDetail string        `bson:"locaiondetail" json:"locaiondetail"`
	Owner          Owner         `bson:"owner" json:"owner"`
	CreatedAt      time.Time     `bson:"created_at" json:"created_at"`
}

type JoinClassRequest struct {
	ClassID  string `bson:"classid" json:"classid"`
	UserName string `bson:"username" json:"username"`
}

type Owner struct {
	Username     string `bson:"username" json:"username"`
	FullName     string `bson:"fullname" json:"fullname"`
	ImageProfile string `bson:"imageprofile" json:"imageprofile"`
}
