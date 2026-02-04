package proto

func ProtocolGitCommit() string {
	return "f1a0480d44808c3a1869eceb98489a9ef87b238b"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
