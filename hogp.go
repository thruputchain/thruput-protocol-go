package proto

func ProtocolGitCommit() string {
	return "96ff4b538169277bf890944f675d502be874612b"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
