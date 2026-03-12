package proto

const ProtocolGitCommit string = "81698e9a1a2c13e12818b6649edb862b9672abe8"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
