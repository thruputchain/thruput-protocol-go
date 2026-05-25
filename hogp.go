package proto

const ProtocolGitCommit string = "cfeadf405b670a7387586f58439ea78ae4ff7acc"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
