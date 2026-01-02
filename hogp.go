package proto

func ProtocolGitCommit() string {
	return "1a51643191d53ea365771bbc5bf392325b109b49"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
