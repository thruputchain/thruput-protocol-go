package proto

func ProtocolGitCommit() string {
	return "e22c9fd737690bfd3cf8e8341c99e35e7628d4dc"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
