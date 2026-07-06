package proto

const ProtocolGitCommit string = "171120cc6f6372b837429e6230a23fde80d6c505"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
