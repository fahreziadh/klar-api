package dao

import (
	"fmt"
	"strings"
	"time"

	. "github.com/klar/models"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	SCHEDULE = "schedule"
)

func (m *DAO) InsertSchedule(schedule Schedule) error {
	schedule.CreatedAt = time.Now()
	err := db.C(SCHEDULE).Insert(&schedule)
	return err
}

func (m *DAO) FindAllSchedule(isToday bool, username string, class string) ([]Schedule, error) {
	var schedule []Schedule
	var err error
	dateNow := strings.Split(time.Now().String(), " ")
	startDate := dateNow[0] + " 00:00:00 +0000 UTC"
	fmt.Println(startDate)
	if isToday {
		err = db.C(SCHEDULE).Find(bson.M{
			"$or": []bson.M{
				bson.M{"class.owner.username": username},
				bson.M{"class.classid": class},
			},
			"$and": []bson.M{
				bson.M{"startdate": startDate},
			},
		}).All(&schedule)
	} else {
		err = db.C(SCHEDULE).Find(bson.M{
			"$or": []bson.M{
				bson.M{"class.owner.username": username},
				bson.M{"class.classid": class},
			},
		}).All(&schedule)
	}

	return schedule, err
}
