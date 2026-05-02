package proto

const ProtocolGitCommit string = "500463c1d97fca03308961a6f2b1781404628882"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
