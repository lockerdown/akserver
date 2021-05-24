package login

import (
	"akserver/server/akbase/base"
	"akserver/setting"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func Router(r *gin.Engine) {
	api := r.Group("/account")
	{
		api.POST("/login", login)
		api.GET("/getPublicKey", getPublicKey)
	}
}

// @Summary 获取密钥
// @tags account
// @accept json
// @Produce  json
// @Success 200	{string} json "{"code": 200,"data" :string,"msg":  "success"}"
// @Failure 500 {string} json "{"code": 500,"data" :"","msg":  "errinfo"}"
// @Router /account/getPublicKey [GET]
func getPublicKey(c *gin.Context) {
	Publickey, err := ioutil.ReadFile(setting.PublicKey)
	if err != nil {
		c.JSON(http.StatusOK, base.ServerFailReturn(err.Error()))
		return
	}
	c.JSON(http.StatusOK, base.SuccessReturn(string(Publickey)))
	return
}
