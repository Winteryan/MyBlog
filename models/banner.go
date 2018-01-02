package models

import (
	"github.com/astaxie/beego/orm"
)

type Banner struct {
	Id       int
	Imgurl   string
	Title    string
	Subtitle string
	Url      string
}

func (a *Banner) TableName() string {
	return "banner"
}

func GetBannerById(id int) (*Banner, error) {
	banner := new(Banner)
	err := orm.NewOrm().QueryTable(TableName("banner")).Filter("id", id).One(banner)
	if err != nil {
		return nil, err
	}
	return banner, err

}

func BannerGetList(page, pageSize int, filters ...interface{}) ([]*Banner, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Banner, 0)
	query := orm.NewOrm().QueryTable(TableName("banner"))
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("id").Limit(pageSize, offset).All(&list)

	return list, total
}

func (a *Banner) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}
