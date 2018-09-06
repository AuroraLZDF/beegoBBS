package models

import (
	"time"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/auroraLZDF/beegoBBS/utils"
	"github.com/auroraLZDF/beegoBBS/utils"
)

type Users struct {
	Id                 int
	Name               string
	Email              string
	Password           string
	Avatar             string
	Introduction       string
	RememberToken      string	//`gorm:column:remember_token`
	NotificationCount int
	CreatedAt         time.Time
	UpdatedAt         time.Time
	LastActivedAt    time.Time
}

// 设置User的表名为`profiles`
func (Users) TableName() string {
	return "bbs_users"
}


func AddUser(u Users) error {
	user := u

	db := DB()
	defer db.Close()

	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func LastUser() Users {
	db := DB()
	defer db.Close()

	user := Users{}
	db.Last(&user)

	return user
}

func FindUserByFields(u Users) (Users, error) {
	db := DB()
	defer db.Close()

	user := Users{}

	result := db.Where(u).First(&user)

	if err := result.Error; err != nil {
		return user, err
	}

	if result.RecordNotFound() == true {

		return user, utils.Error("用户不存在！")
	}

	return user, nil
}

func UpdateUserByEmail(email string, u Users) error {
	db := DB()
	defer db.Close()

	user := Users{}

	if err := db.Model(&user).Where("email=?", email).Updates(u).Error; err != nil {
		return err
	}

	return nil
}

func UpdateUserById(id int, u Users) error {
	db := DB()
	defer db.Close()

	user := Users{}

	if err := db.Model(&user).Where("id=?", id).Updates(u).Error; err != nil {
		return err
	}

	return nil
}

