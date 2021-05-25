package user_service

import (
	"akserver/models"
)

type User struct {
	ID       int
	Name     string
	PageNum  int
	PageSize int
}

func (a *User) GetAll() ([]*models.User, error) {

	users, err := models.GetUsers(a.PageNum, a.PageSize, a.getMaps())
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (a *User) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	// maps["deleted_on"] = 0
	// if a.State != -1 {
	// 	maps["state"] = a.State
	// }
	// if a.TagID != -1 {
	// 	maps["tag_id"] = a.TagID
	// }

	return maps
}
