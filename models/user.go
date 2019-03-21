package models

import (
	"time"

	"auroraLZDF/beegoBBS/utils"
	_ "github.com/go-sql-driver/mysql"
)

type Users struct {
	Id                int    `gorm:"primary_key"`
	Name              string `gorm:"not null"`
	Email             string `gorm:"not null"`
	Password          string `gorm:"not null"`
	Avatar            string
	Introduction      string
	RememberToken     string
	NotificationCount int
	CreatedAt         time.Time
	UpdatedAt         time.Time
	LastActivedAt     time.Time
	Repy              []Replies
}

// 设置User的表名为`bbs_users`
func (Users) TableName() string {
	return "bbs_users"
}

func AddUser(u Users) error {
	db := DB()
	defer db.Close()

	if err := db.Create(&u).Error; err != nil {
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
