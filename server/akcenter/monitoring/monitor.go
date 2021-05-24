package monitoring

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