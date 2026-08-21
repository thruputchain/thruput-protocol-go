package proto

const ProtocolGitCommit string = "db12ec69228c496e75da3a4fc187a684b2d027e2"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
