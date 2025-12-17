package proto

func ProtocolGitCommit() string {
	return "29541f76b1f20557cc788312b5e03a75e8a0e552"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
