package model

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"jrtt/drivers/mysql"
)

type User struct {
	Id  int `gorm:"primaryKey"`
	Uid       int
	Name        string
	Url string
	IsModify       int
	ZhuanfaUid    string
	Token string
}

func (*User) List(where string)[]User {
	var user_list []User
	mysql.Db.Where(where).Find(&user_list)
	return user_list
}


func (*User) Edit(params User)  {
	var info User
	err := mysql.Db.Where("uid = ? and zhuanfa_uid = ?", params.Uid,params.ZhuanfaUid).First(&info).Error
	if errors.Is(err,gorm.ErrRecordNotFound) {
		mysql.Db.Create(&params)
		fmt.Println("添加成功")
	}else if params.Id != 0{
		mysql.Db.Updates(&params)
		fmt.Println("修改成功")
	}else{
		fmt.Println("记录已经存在")
	}
}

func (*User) IsTrue(where string)bool  {
	var info User
	err := mysql.Db.Where(where).First(&info).Error
	if errors.Is(err,gorm.ErrRecordNotFound) {
		return false
	} else {
		return true
	}
}