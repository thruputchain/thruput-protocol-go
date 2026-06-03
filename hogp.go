package proto

const ProtocolGitCommit string = "7874f05decb9b427c723de93e487b1d64a7a9c98"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
