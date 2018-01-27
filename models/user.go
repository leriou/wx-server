package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id      int64 	`orm:"pk;auto"`
	Openid string
	Username  string	`orm:"unique"`
    Age  int64 	
    Mobile string	`orm:"unique"`
    Email string
    Salt string
    State int 
    Password string
	Created string
    Updated string
}

func init() {
	orm.RegisterModel(new(User))
}

func SaveUser(u *User) int64 {
	id, _ := orm.NewOrm().Insert(u)
	return id
}

func UpdateUser(u *User) {
	orm.NewOrm().Update(u)
}

func DeleteUser(id int64) {
	orm.NewOrm().Raw("delete from user where id = ?", id).Exec()
}

func FindUserById(id int64) (bool,User){
	var user User
	err := orm.NewOrm().QueryTable(user).Filter("id", id).One(&user)
	return err != orm.ErrNoRows, user
}

func FindUserByUserName(username string) (bool, User) {
	var user User
	err := orm.NewOrm().QueryTable(user).Filter("Username", username).One(&user)
	return err != orm.ErrNoRows, user
}

func FindAllUser() []*User {
	var u User
	var users []*User
	orm.NewOrm().QueryTable(u).All(&users)
	return users
}
