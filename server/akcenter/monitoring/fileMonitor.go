package monitoring

import (
	"akserver/server/akbase/base"
	"akserver/server/akbase/dbUtil"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//FileEvent 文件监控接口字段
type FileEvent struct {
	Timestamp uint64 `json:"timestamp" xorm:"BigInt notnull 'timestamp'"`      //事件时间戳
	Data_type uint32 `json:"data_type" xorm:"data_type" enum:"1001,1002,1003"` //数据类型
	Pid       uint32 `json:"pid" xorm:"pid"`                                   //进程id
	Ppid      uint32 `json:"ppid" xorm:"ppid"`                                 //父进程ID
	Ptgid     uint32 `json:"ptgid" xorm:"ptgid"`                               //父进程TGID
	Uid       uint32 `json:"uid" xorm:"uid"`                                   //用户ID
	UserName  string `json:"user_name" xorm:"user_name"`                       //用户名
	Gid       uint32 `json:"gid" xorm:"gid"`                                   //用户组ID
	GroupName string `json:"group_name" xorm:"group_name"`                     //组名
	Namespace uint32 `json:"namespace" xorm:"namespace"`                       //进程命名空间ID
	Exe_file  string `json:"exe_file" xorm:"exe_file"`                         //进程文件
	Chg_file  string `json:"chg_file" xorm:"chg_file"`                         //变更文件
}

func (fe *FileEvent) TableName() string {
	return "FileEvent"
}

func (fe *FileEvent) TableSyn() error {
	return dbUtil.Engine().Sync2(fe)
}

func (fe *FileEvent) Insert() error {
	_, err := dbUtil.Engine().Insert(fe)
	return err
}

//Validate 结构体内容校验
func (p *FileEvent) Validate() error {
	return nil
}

//Print 结构体内容打印
func (p *FileEvent) Print() {
	log.Print(p)
}

//文件监控数据上报接口
// @Summary 文件监控数据上报接口
// @Tags monitor
// @accept json
// @Produce  json
// @Param FileEvent body FileEvent true "FileEvent"
// @Success 200	{string} json "{"code": 200,"data" :"","msg":  "success"}"
// @Failure 400 {string} json "{"code": 400,"data" :"","msg":  "errinfo"}"
// @Router /monitor/file [POST]
func fileMonitor(c *gin.Context) {
	var fe FileEvent
	if c.BindJSON(&fe) != nil {
		c.JSON(http.StatusOK, base.FailReturn("请求数据解析错误"))
		return
	}

	fe.TableSyn()

	if err := fe.Validate(); err != nil {
		c.JSON(http.StatusOK, base.FailReturn(err.Error()))
		return
	}

	//入库
	if err := fe.Insert(); err != nil {
		c.JSON(http.StatusOK, base.FailReturn(err.Error()))
		return
	}

	c.JSON(http.StatusOK, base.SuccessReturn("ok."))
	return
}
