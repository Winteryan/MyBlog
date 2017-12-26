package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Blog struct {
	Id           int
	Title        string
	Keywords     string
	Catalogid    string
	Content      string
	Introduction string
	Lastupdate   time.Time
	Type         int
	Status       int
	Views        int
	Imgurl       string
	Createtime   time.Time `orm:"auto_now_add;type(datetime)"`
}

func (a *Blog) TableName() string {
	return "blog"
}

func GetBlogById(id int) (*Blog, error) {
	a := new(Blog)
	err := orm.NewOrm().QueryTable(TableName("blog")).Filter("id", id).One(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func BlogGetList(page, pageSize int, filters ...interface{}) ([]*Blog, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Blog, 0)
	query := orm.NewOrm().QueryTable(TableName("blog"))
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
