package model

import "time"

type Topic struct {
	Id int64            `json:"id"` // auto-increment by-default by xorm
	UserId int64        `xorm:"bigint(20)" json:"user_id"`
	Content string	    `xorm:"text" json:"content"`
	CreatedAt time.Time `xorm:"created" json:"created_at"`
}

type Comment struct {
	Id int64            `json:"id"` // auto-increment by-default by xorm
	UserId int64        `xorm:"bigint(20)" json:"user_id"`
	Content string	    `xorm:"text" json:"content"`
	CreatedAt time.Time `xorm:"created" json:"created_at"`
}

type TopicLike struct {
	Id int64            `json:"id"` // auto-increment by-default by xorm
	UserId int64        `xorm:"bigint(20) index" json:"user_id"`
	TopicId int64       `xorm:"bigint(20) index" json:"topic_id"`
	CreatedAt time.Time `xorm:"created" json:"created_at"`
}
