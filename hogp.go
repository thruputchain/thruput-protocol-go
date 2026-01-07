package proto

func ProtocolGitCommit() string {
	return "d9524d0d585df0630dd6fa5c65383733f60a0dab"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
