package base

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

type App struct {
	PageSize int
}

var AppSetting = &App{}

// GetPage get page parameters
func GetPage(c *gin.Context) int {
	result := 0
	page := com.StrTo(c.Query("page")).MustInt()
	if page > 0 {
		result = (page - 1) * AppSetting.PageSize
	}

	return result
}
