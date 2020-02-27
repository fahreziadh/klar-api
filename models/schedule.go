package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type ScheduleRequest struct {
	Class      Class              `bson:"class" json:"class"`
	TotalMonth int                `bson:"totalmonth" json:"totalmonth"`
	StartDate  string             `bson:"startdate" json:"startdate"`
	DetailTime DetailTimeSchedule `bson:"detailtime" json:"detailtime"`
	Status     string             `bson:"status" json:"status"`
	Teacher    Teacher            `bson:"teacher" json:"teacher"`
}

type Schedule struct {
	ScheduleID bson.ObjectId `bson:"_id" json:"id"`
	Class      Class         `bson:"class" json:"class"`
	StartDate  string        `bson:"startdate" json:"startdate"`
	StartTime  string        `bson:"starttime" json:"starttime"`
	CreatedAt  time.Time     `bson:"created_at" json:"created_at"`
}

type DetailTimeSchedule struct {
	Sunday    string `bson:"sunday" json:"sunday"`
	Monday    string `bson:"monday" json:"monday"`
	Tuesday   string `bson:"tuesday" json:"tuesday"`
	Wednesday string `bson:"wednesday" json:"wednesday"`
	Thursday  string `bson:"thursday" json:"thursday"`
	Friday    string `bson:"Friday" json:"Friday"`
	Saturday  string `bson:"Saturday" json:"Saturday"`
}
