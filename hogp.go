package proto

func ProtocolGitCommit() string {
	return "54edfdd57ee5e16c6d77473efeea84f3f7dcee74"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
