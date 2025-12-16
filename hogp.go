package proto

func ProtocolGitCommit() string {
	return "7c650fa8672d63d0c7b803f369ec894440ad4cef"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
