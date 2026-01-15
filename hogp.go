package proto

func ProtocolGitCommit() string {
	return "aa9f0fe6d65bfeb7abe53badad627f12c8e64721"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
