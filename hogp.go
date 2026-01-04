package proto

func ProtocolGitCommit() string {
	return "a69ccb35141f5dd5cfe2e7ae34a53d004ae39f87"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
