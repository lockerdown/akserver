package monitoring

import (
	"akserver/server/akbase/base"
	"akserver/server/akbase/dbUtil"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//ProcessEvent 进程监控接口字段
type ProcessEvent struct {
	Exe_file   string `json:"exe_file" xorm:"exe_file"`                         //进程文件
	Exe_hash   string `json:"exe_hash" xorm:"exe_hash"`                         //文件hash
	Pid        uint32 `json:"pid" xorm:"pid"`                                   //进程id
	Ppid       uint32 `json:"ppid" xorm:"ppid"`                                 //父进程ID
	Ptgid      uint32 `json:"ptgid" xorm:"ptgid"`                               //父进程TGID
	Data_type  uint32 `json:"data_type" xorm:"data_type" enum:"1001,1002,1003"` //数据类型
	Argv       string `json:"argv" xorm:"argv"`                                 //进程参数
	ParentName string `json:"parent_name" xorm:"parent_name"`                   //父进程名
	Uid        uint32 `json:"uid" xorm:"uid"`                                   //用户ID
	UserName   string `json:"user_name" xorm:"user_name"`                       //用户名
	Gid        uint32 `json:"gid" xorm:"gid"`                                   //用户组ID
	GroupName  string `json:"group_name" xorm:"group_name"`                     //组名
	Namespace  uint32 `json:"namespace" xorm:"namespace"`                       //进程命名空间ID
	Timestamp  uint64 `json:"timestamp" xorm:"timestamp"`                       //事件时间戳
}

func (pe *ProcessEvent) TableName() string {
	return "ProcessEvent"
}

func (pe *ProcessEvent) TableSyn() error {
	return dbUtil.Engine().Sync2(pe)
}

func (pe *ProcessEvent) Insert() error {
	_, err := dbUtil.Engine().Insert(pe)
	return err
}

//Validate 结构体内容校验
func (pe *ProcessEvent) Validate() error {
	return nil
}

//Print 结构体内容打印
func (pe *ProcessEvent) Print() {
	log.Print(pe)
}

//进程监控数据上报接口
// @Summary 进程监控数据上报接口
// @Tags monitor
// @accept json
// @Produce  json
// @Param ProcessEvent body ProcessEvent true "ProcessEvent"
// @Success 200	{string} json "{"code": 200,"data" :"","msg":  "success"}"
// @Failure 400 {string} json "{"code": 400,"data" :"","msg":  "errinfo"}"
// @Router /monitor/process [POST]
func processMonitor(c *gin.Context){
	var pe ProcessEvent
	if c.BindJSON(&pe) != nil {
		c.JSON(http.StatusBadRequest, base.FailReturn("请求数据解析错误"))
		return
	}

	if err := pe.Validate(); err != nil {
		c.JSON(http.StatusOK, base.FailReturn(err.Error()))
		return
	}

	//入库
	if err := pe.Insert(); err != nil {
		c.JSON(http.StatusOK, base.FailReturn(err.Error()))
		return
	}

	c.JSON(http.StatusOK, base.SuccessReturn("ok."))
	return
}