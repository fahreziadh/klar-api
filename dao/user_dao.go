package dao

import (
	"time"

	. "github.com/klar/models"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	USER = "user"
)

func (m *DAO) SignIn(username string) (User, error) {
	var user User
	err := db.C(USER).Find(bson.M{"username": username}).One(&user)
	return user, err
}

func (m *DAO) InsertUser(user User) error {
	user.CreatedAt = time.Now()
	err := db.C(USER).Insert(&user)
	return err
}

func (m *DAO) FindByEmail(user User) error {
	var res User
	err := db.C(USER).Find(bson.M{"email": user.Email}).One(&res)
	return err
}

func (m *DAO) FindByUserName(user User) error {
	var res User
	err := db.C(USER).Find(bson.M{"username": user.UserName}).One(&res)
	return err
}

func (m *DAO) FindUserProfile(username string) (User, error) {
	var res User
	err := db.C(USER).Find(bson.M{"username": username}).One(&res)
	res.Password = ""
	return res, err
}

func (m *DAO) ChangeProfile(username string, fullname string, imageProfile string) error {
	err := db.C(USER).Update(bson.M{"username": username}, bson.M{"$set": bson.M{"fullname": fullname, "imageprofile": imageProfile}})
	return err
}

func (m *DAO) FindStudentByClass(classid string) ([]Student, error) {
	var students []Student
	err := db.C(STUDENT).Find(bson.M{"classid": classid}).All(&students)
	return students, err
}
