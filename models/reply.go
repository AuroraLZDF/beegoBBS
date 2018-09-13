package models

import "time"

type Replies struct {
	Id        int    `gorm:"primary_key"`
	TopicId   int    `gorm:"not null"`
	UserId    int    `gorm:"not null"`
	content   string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Topics Topics `gorm:"ForeignKey:TopicId"`
	Users  Users  `gorm:"ForeignKey:UserId"`
}
