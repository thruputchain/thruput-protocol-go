package proto

func ProtocolGitCommit() string {
	return "3032b89e4ef30d1b061cdbefee9fd4fafbb82c4b"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
