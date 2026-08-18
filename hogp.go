package proto

const ProtocolGitCommit string = "d76ec609fc8b26553b1e30906b8dd68c5a854716"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
