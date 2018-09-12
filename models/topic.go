package models

import (
	"fmt"
	"time"

	"github.com/auroraLZDF/beegoBBS/utils"
)

type Topics struct {
	Id              int
	Title           string
	Body            string
	UserId          int
	CategoryId      int
	ReplyCount      int
	ViewCount       int
	LastReplyUserId int
	Order           int
	Excerpt         string
	Slug            string
	Status          int `gorm:"default:1"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

const (
	STATUS_ON  = 1
	STATUS_OFF = 2
)

// 设置User的表名为`profiles`
func (Topics) TableName() string {
	return "bbs_topics"
}

func (Topics) Create(val Topics) error {
	db := DB()
	defer db.Close()

	//fmt.Println(val)

	res := db.Create(&val)
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

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

func (Topics) TopicLists(params map[string]interface{}) (*[]map[string]string, error) {
	db := DB()
	defer db.Close()

	var page = params["page"].(int)
	if page <= 1 {
		page = 1
	}

	limit := 10

	//_db := db.Where("status = ?", STATUS_ON)
	var where = " t.status=" + utils.IntToString(STATUS_ON)

	id := params["id"].(int)
	if id > 0 {
		//_db = _db.Where("id = ?", id)
		where = where + " and t.id=" + utils.IntToString(id)
	}

	category := params["category"].(int)
	if category > 0 {
		//_db = _db.Where("category_id = ?", category)
		where = where + " and t.category_id=?" + utils.IntToString(category)
	}

	order := params["order"]
	if order == "default" {
		//_db.Order("updated_at desc, order asc, reply_count desc")
		where = where + " order by t.updated_at desc, t.order asc, t.reply_count desc"
	} else if order == "recent" {
		//_db.Order("created_at desc, order asc, reply_count desc")
		where = where + " order by t.created_at desc, t.order asc, t.reply_count desc"
	}

	//_limit := utils.IntToString(page) + "," + limit
	// _db.Limit(_limit)
	where = where + " limit " + utils.IntToString(page-1) + ", " + utils.IntToString(limit)

	//topic := []Topics{}
	//result := _db.Find(&topic)
	sql := "select * from bbs_topics t left join bbs_users u on t.user_id=u.id left join bbs_categories c on c.id=t.category_id where " + where
	fmt.Println(sql)
	result, err := utils.GetAll(sql)
	if err != nil {
		return result, err
	}

	return result, nil
}
