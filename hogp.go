package proto

func ProtocolGitCommit() string {
	return "15bfef5895daa6f8de5e8b22a3b2d9292f724686"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
