package proto

const ProtocolGitCommit string = "d982d6de82909f4a6c3feeed3a0f6111a3befb2a"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
