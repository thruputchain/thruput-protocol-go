package proto

const ProtocolGitCommit string = "021c335cd81054ed1499791f5d727a9d88622209"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
