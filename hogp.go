package proto

func ProtocolGitCommit() string {
	return "ba581349d4d4c4fa4dd00ea950bb5afeab22c5d5"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
