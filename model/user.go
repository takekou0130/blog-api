package model

type User struct {
	Id        int64  `xorm:"pk autoincr int(64)" form:"id" json:"id"`
	FirstName string `xorm:"first_name varchar(40)" json:"firstName" form:"firstName"`
	LastName  string `xorm:"last_name varchar(40)" json:"lastName" form:"lastName"`
}
