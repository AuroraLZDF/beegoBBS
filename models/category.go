package models

import (
	"time"

	"auroraLZDF/beegoBBS/utils"
)

type Categories struct {
	Id          int    `gorm:"primary_key"`
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	PostCount   int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

//var menu = map[int]map[string]string{}

// 设置User的表名为`profiles`
func (Categories) TableName() string {
	return "bbs_categories"
}

func (Categories) FindAll() (*[]Categories, error) {
	db := DB()
	defer db.Close()

	categories := []Categories{}

	result := db.Find(&categories)
	if err := result.Error; err != nil {
		return &categories, err
	}

	if result.RecordNotFound() == true {
		return &categories, utils.Error("未查找到记录")
	}

	return &categories, nil
}
