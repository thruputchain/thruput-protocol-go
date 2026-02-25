package proto

const ProtocolGitCommit string = "40f3cae2d5b9fa90113739f7754cc9757de2f0cd"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
