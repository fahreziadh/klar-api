package dao

import (
	"time"

	. "github.com/klar/models"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	ATTENDANCE = "attendance"
)

//insert new class
func (m *DAO) InsertAttendance(attendance Attendance) error {
	attendance.CreatedAt = time.Now()
	var oldAttendance []Attendance
	var err error
	//check attendance first
	err = db.C(ATTENDANCE).Find(bson.M{
		"$and": []bson.M{
			bson.M{"student.username": attendance.Student.UserName},
			bson.M{"scheduleid": attendance.ScheduleId},
		},
	}).All(&oldAttendance)

	if len(oldAttendance) == 0 { //if nil result, will be create new attendance
		err = db.C(ATTENDANCE).Insert(&attendance)
	} else { //if not nil result, will be update new attendance
		err = db.C(ATTENDANCE).Update(
			bson.M{
				"$and": []bson.M{
					bson.M{"student.username": attendance.Student.UserName},
					bson.M{"scheduleid": attendance.ScheduleId},
				},
			},
			bson.M{
				"$set": bson.M{"attendance": attendance.Attendance},
			})
	}

	return err
}

func (m *DAO) FindAttendanceBySchedule(scheduleid string) ([]Attendance, error) {
	var attendance []Attendance
	err := db.C(ATTENDANCE).Find(bson.M{"scheduleid": scheduleid}).All(&attendance)
	return attendance, err
}
