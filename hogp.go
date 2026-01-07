package proto

func ProtocolGitCommit() string {
	return "ffa48f6e1e8b6c407275be875c1cf3f4ea8b9c35"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
