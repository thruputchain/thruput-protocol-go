package proto

const ProtocolGitCommit string = "a6516e66a6c3f60f6bb8c8e32054a7f52cd74f7d"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
