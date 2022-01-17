package model

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"jrtt/drivers/mysql"
)

type List struct {
	Id int `gorm:"primaryKey"`
	SearchId string
	QueryId string
	SearchResultId string
	IsWinner bool `gorm:"type:char"`
	LotteryTime string `gorm:"type:datetime"`
	ParticipateType int
	Reward string
	Status int
	WinnerType int
	Detail string
	LuckData string
	IsOk int
	IsRepost int
	IsLike int
	UserId int
	UserName string
	ZhuanfaUid string
}

func (*List) List(where string)[]List {
	var list []List
	mysql.Db.Where(where).Find(&list)
	return list
}

func (*List) Edit(params List)  {
	var info List
	err := mysql.Db.Where("search_result_id = ? and zhuanfa_uid = ?", params.SearchResultId,params.ZhuanfaUid,params.ZhuanfaUid).First(&info).Error
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