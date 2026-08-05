package proto

const ProtocolGitCommit string = "b26a10d826bb9e478bdff440257909480c5d8cdd"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
