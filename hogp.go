package proto

func ProtocolGitCommit() string {
	return "941e8b8545b18e4fbe05df1a8731cdefa2a177ea"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
