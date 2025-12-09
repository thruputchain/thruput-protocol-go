package proto

func ProtocolGitCommit() string {
	return "+508012f05d45f5027068a34432f1c277c4c30aa9"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
