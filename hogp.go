package proto

const ProtocolGitCommit string = "d9fd7494e9102a55bfb90e79cdcabf631c30ddb4"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
