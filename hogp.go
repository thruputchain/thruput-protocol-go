package proto

func ProtocolGitCommit() string {
	return "7df9b194215b20ab42acebf213aa57bf4c3ec34"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
