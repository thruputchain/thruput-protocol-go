package proto

const ProtocolGitCommit string = "e1053a3ea9dff92fab6b82baeef66a33f91ed4a7"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
