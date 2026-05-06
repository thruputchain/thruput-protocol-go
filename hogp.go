package proto

const ProtocolGitCommit string = "20b280b393e46d810c247f9e86bdecfba1a6d47c"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
