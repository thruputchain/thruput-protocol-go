package proto

func ProtocolGitCommit() string {
	return "+656e88604abb1056cb74ae74007ad169c3e35e07"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
