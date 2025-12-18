package proto

func ProtocolGitCommit() string {
	return "07fb347c22d737e2e04bae0081496fc84b46b1ca"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
