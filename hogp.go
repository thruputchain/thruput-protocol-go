package proto

func ProtocolGitCommit() string {
	return "6617ccce1f4ca6c7c8ba5242f168dcab68ccbcb5"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
