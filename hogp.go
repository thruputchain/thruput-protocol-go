package proto

func ProtocolGitCommit() string {
	return "bb2a595a39712f121afe07508dc096a8c93ceecb"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
