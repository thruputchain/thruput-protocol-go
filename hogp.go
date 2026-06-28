package proto

const ProtocolGitCommit string = "f9043f5aeb5093951f9bebcf383ed62c73befd34"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
