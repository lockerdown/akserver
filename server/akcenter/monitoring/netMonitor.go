package monitoring

import (
	"akserver/server/akbase/base"
	"akserver/server/akbase/dbUtil"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//NetEvent 网络监控接口字段
type NetEvent struct {
	Timestamp uint64 `json:"timestamp" xorm:"timestamp"`                       //事件时间戳
	Data_type uint32 `json:"data_type" xorm:"data_type" enum:"1001,1002,1003"` //数据类型
	Pid       uint32 `json:"pid" xorm:"pid"`                                   //进程id
	Uid       uint32 `json:"uid" xorm:"uid"`                                   //用户ID
	UserName  string `json:"user_name" xorm:"user_name"`                       //用户名
	Gid       uint32 `json:"gid" xorm:"gid"`                                   //用户组ID
	GroupName string `json:"group_name" xorm:"group_name"`                     //组名
	Namespace uint32 `json:"namespace" xorm:"namespace"`                       //进程命名空间ID
	SrcIp     string `json:"src_ip" xorm:"src_ip"`                             //源地址
	DstIp     string `json:"dst_ip" xorm:"dst_ip"`                             //目的地址
	SrcPort   uint16 `json:"src_port" xorm:"src_port"`                         //源端口
	DstPort   uint16 `json:"dst_port" xorm:"dst_port"`                         //目的端口
	Exe_file  string `json:"exe_file" xorm:"exe_file"`                         //进程文件
}

func (ne *NetEvent) TableName() string {
	return "NetEvent"
}

func (ne *NetEvent) TableSyn() error {
	return dbUtil.Engine().Sync2(ne)
}

func (ne *NetEvent) Insert() error {
	_, err := dbUtil.Engine().Insert(ne)
	return err
}

//Validate 结构体内容校验
func (ne *NetEvent) Validate() error {
	return nil
}

//Print 结构体内容打印
func (ne *NetEvent) Print() {
	log.Print(ne)
}


//主机网络监控数据上报接口
// @Summary 主机网络监控数据上报接口
// @Tags monitor
// @accept json
// @Produce  json
// @Param NetEvent body NetEvent true "NetEvent"
// @Success 200	{string} json "{"code": 200,"data" :"","msg":  "success"}"
// @Failure 400 {string} json "{"code": 400,"data" :"","msg":  "errinfo"}"
// @Router /monitor/net [POST]
func netMonitor(c *gin.Context){
	var ne NetEvent
	if c.BindJSON(&ne) != nil {
		c.JSON(http.StatusBadRequest, base.FailReturn("请求数据解析错误"))
		return
	}

	if err := ne.Validate(); err != nil {
		c.JSON(http.StatusOK, base.FailReturn(err.Error()))
		return
	}

	//入库
	if err := ne.Insert(); err != nil {
		c.JSON(http.StatusOK, base.FailReturn(err.Error()))
		return
	}

	c.JSON(http.StatusOK, base.SuccessReturn("ok."))
	return
}