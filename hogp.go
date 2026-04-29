package proto

const ProtocolGitCommit string = "850287355957af222bbd2e300a3095f6aaf2d61c"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
