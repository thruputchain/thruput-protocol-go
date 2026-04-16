package proto

const ProtocolGitCommit string = "c916509bcd94c32e333514f091d14c28c6bcebcb"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
