package proto

func ProtocolGitCommit() string {
	return "9015ab295ee2a5d1b3155ed57022f3409c74ddfc"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
