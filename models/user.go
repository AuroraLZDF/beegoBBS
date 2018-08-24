package models

import (
	"time"
	_ "github.com/go-sql-driver/mysql"
	"log"
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


func AddUser() {

}

func FindUserByEmail(email string) Users {
	db := DB()
	defer db.Close()

	user := Users{}
	//db.First(&user, "email = ?", email)
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		log.Fatal(err)
	}

	return user
}
