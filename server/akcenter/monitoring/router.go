/* package monitoring

A web framework includes app server, logger, panicer, util and so on.
*/
package monitoring

import (
	"akserver/server/akbase/base"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router(r *gin.Engine) {
	monitor := r.Group("/monitor")
	{
		monitor.POST("/file",fileMonitor)
		monitor.POST("/process",processMonitor)
		monitor.POST("/net",netMonitor)
	}
}


func fileMonitor(c *gin.Context){
}

//进程监控数据上报接口
// @Summary 进程监控数据上报接口
// @Tags monitor
// @accept json
// @Produce  json
// @Param ProcessMonitor body ProcessMonitor true "ProcessMonitor"
// @Success 200	{string} json "{"code": 200,"data" :"","msg":  "success"}"
// @Failure 400 {string} json "{"code": 400,"data" :"","msg":  "errinfo"}"
// @Router /monitor/file [POST]
func processMonitor(c *gin.Context){
	var processMonitor ProcessMonitor
	if c.BindJSON(&processMonitor) != nil {
		c.JSON(http.StatusBadRequest, base.FailReturn(errors.New("请求数据解析错误")))
		return
	}
	processMonitor.Print()
	//入库

	c.JSON(http.StatusOK, base.SuccessReturn("ok."))
	return
}

//主机网络监控数据上报接口
// @Summary 主机网络监控数据上报接口
// @Tags monitor
// @accept json
// @Produce  json
// @Param NetWorkMonitor body NetWorkMonitor true "NetWorkMonitor"
// @Success 200	{string} json "{"code": 200,"data" :"","msg":  "success"}"
// @Failure 400 {string} json "{"code": 400,"data" :"","msg":  "errinfo"}"
// @Router /monitor/net [POST]
func netMonitor(c *gin.Context){
	var netWorkMonitor NetWorkMonitor
	if c.BindJSON(&netWorkMonitor) != nil {
		c.JSON(http.StatusBadRequest, base.FailReturn(errors.New("请求数据解析错误")))
		return
	}

	//入库
}