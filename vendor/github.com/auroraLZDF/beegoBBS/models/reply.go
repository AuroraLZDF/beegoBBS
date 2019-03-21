package models

import (
	"time"
)

type Replies struct {
	Id        int    `gorm:"primary_key"`
	TopicId   int    `gorm:"not null"`
	UserId    int    `gorm:"not null"`
	Content   string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Replies) TableName() string {
	return "bbs_replies"
}

func (Replies) AddReply(r Replies) error {
	db := DB()
	defer db.Close()

	if err := db.Create(&r).Error; err != nil {
		return err
	}

	return nil
}

func (Replies) Delete(id int, topic_id int) error {
	db := DB()
	defer db.Close()

	var reply Replies
	db = db.Where("id=? and topic_id=?", id, topic_id)
	if err := db.First(&reply).Error; err != nil {
		return err
	}

	if err := db.Delete(&reply).Error; err != nil {
		return err
	}

	return nil
}
