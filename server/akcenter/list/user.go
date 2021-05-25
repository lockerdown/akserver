package list

import (
	"akserver/server/akbase/base"
	"net/http"

	"akserver/service/user_service"

	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {
	article, err := user_service.User()
	if err != nil {
		c.JSON(http.StatusOK, base.FailReturn(err))
		return
	}
	c.JSON(http.StatusOK, base.SuccessReturn(article))
}
