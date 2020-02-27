package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Attendance struct {
	AttendanceID bson.ObjectId `bson:"_id" json:"id"`
	Student      Student       `bson:"student" json:"student"`
	Attendance   bool          `bson:"attendance" json:"attendance"`
	ScheduleId   string        `bson:"scheduleid" json:"scheduleid"`
	CreatedAt    time.Time     `bson:"created_at" json:"created_at"`
}
