package proto

const ProtocolGitCommit string = "567533feb9f00bb98eeb5dbe884fb34899fb9ef5"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
