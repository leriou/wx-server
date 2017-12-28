package models

import (
	"github.com/astaxie/beego/orm"
)

type Msg struct {
	Mid int64 `orm:"pk;auto"`
	Msgtype string
	Content string
	Senderid string
	Reciverid string
	State int
	Created	string
}

func init() {
	orm.RegisterModel(new(Msg))
}

func SaveMsg(m *Msg) int64{
	o := orm.NewOrm()
	id, _ := o.Insert(m)
	return id
}
