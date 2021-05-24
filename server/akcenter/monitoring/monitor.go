package monitoring

import "log"

type AkEventType uint16

var (
	ProcessFork AkEventType = 1001
	ProcessExec AkEventType = 1002
	ProcessExit AkEventType = 1003

	FileCreate AkEventType = 2001
	FileWrite  AkEventType = 2002
	FileChmod  AkEventType = 2003
	FileDelete AkEventType = 2004
	FileRemove AkEventType = 2005

	TcpConnect AkEventType = 3001
	TcpBind    AkEventType = 3002
	TcpAccept  AkEventType = 3003
	TcpClose   AkEventType = 3004
	DnsSend    AkEventType = 3010
)

func (a AkEventType) String() string {
	switch a {
	case ProcessFork:
		return "ProcessFork"
	case ProcessExec:
		return "ProcessExec"
	case ProcessExit:
		return "ProcessExit"

	case FileCreate:
		return "FileCreate"
	case FileWrite:
		return "FileWrite"
	case FileChmod:
		return "FileChmod"
	case FileDelete:
		return "FileDelete"
	case FileRemove:
		return "FileRemove"

	case TcpConnect:
		return "TcpConnect"
	case TcpBind:
		return "TcpBind"
	case TcpAccept:
		return "TcpAccept"
	case TcpClose:
		return "TcpClose"
	case DnsSend:
		return "DnsSend"
	}
	return ""
}

//ProcessMonitor 进程监控接口字段
type ProcessMonitor struct {
	Exe_file  string `json:"exe_file"`                        //进程文件
	Exe_hash  string `json:"exe_hash"`                        //文件hash
	Pid       uint32 `json:"pid"`                             //进程id
	Ppid      uint32 `json:"ppid"`                            //父进程ID
	Data_type uint32 `json:"data_type" enum:"1001,1002,1003"` //数据类型
	Argv      string `json:"argv"`                            //进程参数
	Uid       uint32 `json:"uid"`                             //用户ID
	UserName  string `json:"user_name"`                       //用户名
	Gid       uint32 `json:"gid"`                             //用户组ID
	GroupName string `json:"group_name"`                      //组名
	Namespace uint32 `json:"namespace"`                       //进程命名空间ID
	Timestamp uint64 `json:"timestamp"`                       //事件时间戳
}

//Validate 结构体内容校验
func (p *ProcessMonitor) Validate() {}

//Print 结构体内容打印
func (p *ProcessMonitor) Print() {
	log.Print(p)
}

//NetWorkMonitor 网络监控接口字段
type NetWorkMonitor struct {
	Exe_file  string `json:"exe_file"`	//进程文件
	Pid       uint32 `json:"pid"`		//进程id
	Data_type uint16 `json:"data_type" enum:"3001,3002,3003,3004,3010"`	//数据类型
	SrcIp     string `json:"srcIp"`		//源地址
	SrcPort   uint16 `json:"srcPort" min:"1" max:"65535"`		//源端口
	DstIp     string `json:"dstIp"`		//目的地址
	DstPort   uint16 `json:"dstPort" min:"1" max:"65535"`		//目的端口
	Uid       uint32 `json:"uid"`		//用户ID
	Gid       uint32 `json:"gid"`		//用户组ID
	Namespace uint32 `json:"namespace"`	//进程命名空间ID
	Timestamp uint64 `json:"timestamp"`	//事件时间戳
}

//Validate 结构体内容校验
func (p *NetWorkMonitor) Validate() {

}