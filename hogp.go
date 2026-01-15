package proto

func ProtocolGitCommit() string {
	return "7b4aa2d670ec4bd06b81ceecf587935f6ba81ff3"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
