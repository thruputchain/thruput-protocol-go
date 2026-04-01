package proto

const ProtocolGitCommit string = "c9543de8a11358bbac0cb55585232290c79ae96b"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
