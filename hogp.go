package proto

const ProtocolGitCommit string = "96cd70b2b4f4ff06bd74145f7a5e37d2b7338367"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
