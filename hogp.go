package proto

const ProtocolGitCommit string = "b74b41ca7a487f2d6ddf6f391f37fcbdce9ff4ae"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
