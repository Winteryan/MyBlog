package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id          int
	Username    string
	Password    string
	Sex         int
	Age         int
	Email       string
	Imgurl      string
	Status      int
	Personalkey string
}

func (a *User) TableName() string {
	return "user"
}

func GetUserById(id int) (*User, error) {
	user := new(User)
	err := orm.NewOrm().QueryTable(TableName("user")).Filter("id", id).One(user)
	if err != nil {
		return nil, err
	}
	return user, err

}

func UserGetList(page, pageSize int, filters ...interface{}) ([]*User, int64) {
	offset := (page - 1) * pageSize
	list := make([]*User, 0)
	query := orm.NewOrm().QueryTable(TableName("user"))
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)

	return list, total
}
