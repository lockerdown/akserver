package models

import (
	"akserver/server/akbase/dbUtil"
)

type User struct {
	Uid      int
	userName string
}

func GetUsers(pageNum int, pageSize int, maps interface{}) ([]*User, error) {
	var file []*User
	err := dbUtil.Engine().Table("fileevent").Where("uid = ?", file).Limit(pageSize, 0).Find(&file)

	println(file)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// func GetTotal(maps interface{}) (int, error) {
// 	var count int
// 	if err := dbUtil.Engine().Count(&count); err != nil {
// 		return 0, err
// 	}

// }
