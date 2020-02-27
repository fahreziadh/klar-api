package dao

import (
	"time"

	. "github.com/klar/models"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	CLASS   = "class"
	STUDENT = "student"
)

//insert new class
func (m *DAO) InsertClass(class Class) error {
	class.CreatedAt = time.Now()
	err := db.C(CLASS).Insert(&class)
	return err
}

//find all class
func (m *DAO) FindAllClass() ([]Class, error) {
	var class []Class
	err := db.C(CLASS).Find(bson.M{}).All(&class)
	return class, err
}

func (m *DAO) FindClassByClassId(classid string) (Class, error) {
	var class Class
	err := db.C(CLASS).Find(bson.M{"classid": classid}).One(&class)
	return class, err
}

func (m *DAO) FindClassByOwner(username string) ([]Class, error) {
	var class []Class
	err := db.C(CLASS).Find(bson.M{"owner.username": username}).All(&class)
	return class, err
}

func (m *DAO) InsertStudent(student Student) error {
	student.CreatedAt = time.Now()
	var s []Student
	var err error

	err = db.C(STUDENT).Find(bson.M{
		"$and": []bson.M{
			bson.M{"username": student.UserName},
			bson.M{"classid": student.ClassID},
		},
	}).All(&s)

	if len(s) == 0 {
		err = db.C(STUDENT).Insert(&student)
	} else {
		err = db.C(STUDENT).Insert(&err)
	}

	return err
}

func (m *DAO) FindClassByStudent(username string) ([]Student, error) {
	var res []Student
	err := db.C(STUDENT).Find(bson.M{"username": username}).All(&res)
	return res, err
}
