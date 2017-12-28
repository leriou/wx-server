package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id      int64 	`orm:"pk;auto"`
	Openid string
	Username  string	`orm:"unique"`
    Age  int64 	
    Mobile string	`orm:unique`
    Email string
    Salt string
    State int 
    Password string
	Created string
    Updated string
}

func init() {
	orm.RegisterModel(new (User))
}

func SaveUser(u *User) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(u)
	return id
}

func UpdateUser(u *User) {
	o := orm.NewOrm()
	o.Update(u)
}

func DeleteUser(id int) {
	o := orm.NewOrm()
	o.Raw("delete from user where id = ?", id).Exec()
}

func FindUserById(id int) (bool,User){
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("id", id).One(&user)
	return err != orm.ErrNoRows, user
}

func FindUserByUserName(username string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Username", username).One(&user)
	return err != orm.ErrNoRows, user
}

func FindAllUser() []*User {
	o := orm.NewOrm()
	var u1 User
	var users []*User
	o.QueryTable(u1).All(&users)
	return users
}
