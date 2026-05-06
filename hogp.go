package proto

const ProtocolGitCommit string = "5d7c0a4099e4d449637c4b2ac5a2a267100b5846"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
