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

	var page = params["page"].(int)
	if page <= 1 {
		page = 1
	}

	db.Where("status = ?", STATUS_ON)

	id := params["id"].(int)
	if id > 0 {
		db = db.Where("id = ?", id)
	}

	category := params["category"].(int)
	if category > 0 {
		db = db.Where("category_id = ?", category)
	}

	// 获取记录个数
	db.Model(topics).Count(&count)

	order := params["order"]
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

	//db.Where("id=?", id).First(&topic)
	// 关联的关键代码
	//db.Model(&topic).Related(&topic.Category, "CategoryId")
	//result := db.Model(&topic).Related(&topic.User, "UserId")

	result := db.Where("id=?", id).Preload("Category").Preload("User").First(&topic)
	if err := result.Error; err != nil {
		return &topic, err
	}

	if result.RecordNotFound() == true {
		return &topic, utils.Error("文章不存在！")
	}

	return &topic, nil
}
