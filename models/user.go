package models

import (
	"akserver/server/akbase/dbUtil"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID    int    `json:"tag_id" gorm:"index"`
	Title string `json:"title"`
}

func GetUsers(pageNum int, pageSize int, maps interface{}) ([]*User, error) {
	var users []*User
	err := dbUtil.Engine().Where("name = ?", maps).Limit(pageSize, 0).Find(&users)
	// err := db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles).Error
	println(users)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return users, nil
}

// func GetTotal(maps interface{}) (int, error) {
// 	var count int
// 	if err := dbUtil.Engine().Count(&count); err != nil {
// 		return 0, err
// 	}

// }
