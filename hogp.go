package proto

const ProtocolGitCommit string = "6e4931a795210b1cdd86ace89b5a30f735f3d2d2"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
