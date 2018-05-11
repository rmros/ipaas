package models

import "time"

// User user info
type User struct {
	ID             int32 `gorm:"primary_key"`
	Name           string
	Displayname    string
	Password       string
	Email          string
	Phone          string
	CreationTime   time.Time
	LastLoginTime  time.Time
	LoginFrequency int
	Active         int8
	APIToken       string
	Role           int32
	Type           int
	Company        string
	Teams          []*Team `gorm:"many2many:user_teams;"`
}

// Team team info
type Team struct {
	ID           string `gorm:"primary_key"`
	TeamName     string
	Description  string
	CreatorID    int32
	CreationTime time.Time
	Users        []*User `gorm:"many2many:team_users;"`
}

// import (
// 	"errors"
// 	"strconv"
// 	"time"
// )

// var (
// 	UserList map[string]*User
// )

// func init() {
// 	UserList = make(map[string]*User)
// 	u := User{"user_11111", "astaxie", "11111", Profile{"male", 20, "Singapore", "astaxie@gmail.com"}}
// 	UserList["user_11111"] = &u
// }

// type User struct {
// 	Id       string
// 	Username string
// 	Password string
// 	Profile  Profile
// }

// type Profile struct {
// 	Gender  string
// 	Age     int
// 	Address string
// 	Email   string
// }

// func AddUser(u User) string {
// 	u.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
// 	UserList[u.Id] = &u
// 	return u.Id
// }

// func GetUser(uid string) (u *User, err error) {
// 	if u, ok := UserList[uid]; ok {
// 		return u, nil
// 	}
// 	return nil, errors.New("User not exists")
// }

// func GetAllUsers() map[string]*User {
// 	return UserList
// }

// func UpdateUser(uid string, uu *User) (a *User, err error) {
// 	if u, ok := UserList[uid]; ok {
// 		if uu.Username != "" {
// 			u.Username = uu.Username
// 		}
// 		if uu.Password != "" {
// 			u.Password = uu.Password
// 		}
// 		if uu.Profile.Age != 0 {
// 			u.Profile.Age = uu.Profile.Age
// 		}
// 		if uu.Profile.Address != "" {
// 			u.Profile.Address = uu.Profile.Address
// 		}
// 		if uu.Profile.Gender != "" {
// 			u.Profile.Gender = uu.Profile.Gender
// 		}
// 		if uu.Profile.Email != "" {
// 			u.Profile.Email = uu.Profile.Email
// 		}
// 		return u, nil
// 	}
// 	return nil, errors.New("User Not Exist")
// }

// func Login(username, password string) bool {
// 	for _, u := range UserList {
// 		if u.Username == username && u.Password == password {
// 			return true
// 		}
// 	}
// 	return false
// }

// func DeleteUser(uid string) {
// 	delete(UserList, uid)
// }
