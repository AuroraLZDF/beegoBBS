package models

import (
	"time"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/auroraLZDF/beegoBBS/utils"
)

type Users struct {
	Id                 int
	Name               string
	Email              string
	Password           string
	Avatar             string
	Introduction       string
	RememberToken     string
	NotificationCount int
	CreatedAt         time.Time
	UpdatedAt         time.Time
	LastActivedAt    time.Time
}

// 设置User的表名为`profiles`
func (Users) TableName() string {
	return "bbs_users"
}


func AddUser(u Users) (bool, error) {
	user := u

	db := DB()
	defer db.Close()

	if err := db.Create(&user).Error; err != nil {
		return false, err
	}

	return true, nil
}

func FindUserByFields(u Users) (bool , Users, string) {
	db := DB()
	defer db.Close()

	user := Users{}

	result := db.Where(u).First(&user)

	if err := result.Error; err != nil {
		//utils.ShowErr(err)
		return false, user, err.Error()
	}

	if result.RecordNotFound() == true {
		//utils.ShowErr("数据不存在！")
		return false, user, "数据不存在！"
	}

	return true, user, "用户已存在！"
}

