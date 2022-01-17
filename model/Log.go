package model

import (
	"jrtt/drivers/mysql"
)

type Log struct {
	Id int `gorm:"primaryKey"`
	Log string `gorm:"type:mediumtext"`
}

func (*Log)Add(log string)  {
	var model Log
	model.Log = log
	mysql.Db.Create(&model)
}