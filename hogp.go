package proto

const ProtocolGitCommit string = "5b73b22e761092f414edf3c0fc40c8f4542fb49c"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
