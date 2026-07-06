package proto

const ProtocolGitCommit string = "8eee1b768177db34c3ef03410275679ad96504e7"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
