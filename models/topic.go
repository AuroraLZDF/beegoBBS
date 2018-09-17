package models

import (
	"fmt"
	"time"

	"github.com/auroraLZDF/beegoBBS/utils"
)

type Topics struct {
	Id              int    `gorm:"primary_key"`
	Title           string `gorm:"not null"`
	Body            string `gorm:"not null"`
	UserId          int    `gorm:"not null"`
	CategoryId      int    `gorm:"not null"`
	ReplyCount      int
	ViewCount       int
	LastReplyUserId int
	Order           int
	Excerpt         string
	Slug            string
	Status          int `gorm:"default:1"`
	CreatedAt       time.Time
	UpdatedAt       time.Time

	Category Categories `gorm:"foreignkey:CategoryId"`
	User     Users      `gorm:"foreignkey:UserId"`
}

const (
	STATUS_ON  = 1
	STATUS_OFF = 2
)

// 设置User的表名为`bbs_topics`
func (Topics) TableName() string {
	return "bbs_topics"
}

// 创建话题
func (Topics) Create(val Topics) error {
	db := DB()
	defer db.Close()

	res := db.Create(&val)
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

// 获得所有话题
func (Topics) FindAll() (*[]Topics, error) {
	db := DB()
	defer db.Close()

	topic := []Topics{}
	result := db.Find(&topic)

	if err := result.Error; err != nil {
		return &topic, err
	}

	if result.RecordNotFound() == true {
		return &topic, utils.Error("未查找到记录")
	}

	return &topic, nil
}

// 获得话题分页列表
func (Topics) TopicLists(params map[string]interface{}) (map[string]interface{}, error) {
	db := DB()
	defer db.Close()

	var limit = 5
	var count int64
	var topics []Topics

	var page = 0
	if params["page"] != nil {
		page = params["page"].(int)
	}

	if page <= 1 {
		page = 1
	}

	db.Where("status = ?", STATUS_ON)

	if params["id"] != nil {
		db = db.Where("id = ?", params["id"].(int))
	}

	if params["category"] != nil {
		db = db.Where("category_id = ?", params["category"].(int))
	}

	// 获取记录个数
	db.Model(topics).Count(&count)
	fmt.Println("count:", count)

	var order = ""
	if params["order"] != nil {
		order = params["order"].(string)
	}

	if order == "recent" {
		db = db.Order("created_at desc")
	} else {
		db = db.Order("updated_at desc")
	}

	db = db.Limit(limit).Offset((page - 1) * limit)

	// 获取记录数据
	result := db.Model(topics).Preload("Category").Preload("User").Find(&topics)
	if err := result.Error; err != nil {
		fmt.Println(err)
	}

	pageNate := utils.Paginator(page, limit, count, params["currentPath"].(string))

	return map[string]interface{}{
		"pageNate": pageNate,
		"topics":   &topics,
	}, nil
}

// 根据话题 id 检索
func (Topics) TopicByID(id int) (*Topics, error) {
	db := DB()
	defer db.Close()

	var topic Topics

	result := db.Where("id=?", id).Preload("Category").Preload("User").First(&topic)
	if err := result.Error; err != nil {
		return &topic, err
	}

	if result.RecordNotFound() == true {
		return &topic, utils.Error("文章不存在！")
	}

	return &topic, nil
}

// 修改话题内容
func (Topics) UpdateById(id int, T Topics) error {
	db := DB()
	defer db.Close()

	var topic Topics

	if err := db.Model(topic).Where("id=?", id).Updates(T).Error; err != nil {
		return err
	}

	return nil
}

// 删除话题
func (Topics) DeleteById(id int) error {
	db := DB()
	defer db.Close()

	var topic Topics
	if _, err := topic.TopicByID(id); err != nil {
		return err
	}

	if err := db.Where("id=?", id).Delete(topic).Error; err != nil {
		return err
	}

	return nil
}
