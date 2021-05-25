package list

import (
	"akserver/server/akbase/base"
	"net/http"

	"akserver/service/user_service"

	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {
	userService := user_service.User{
		PageNum:  base.GetPage(c),
		PageSize: base.AppSetting.PageSize,
	}

	user, err := userService.GetAll()
	println(user)
	if err != nil {
		c.JSON(http.StatusOK, base.FailReturn(err))
		return
	}
	c.JSON(http.StatusOK, base.SuccessReturn(user))
}
