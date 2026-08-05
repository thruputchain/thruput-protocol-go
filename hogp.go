package proto

const ProtocolGitCommit string = "0b20f225f6f497d9c0b70302147b554e5ff85e31"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
