package proto

func ProtocolGitCommit() string {
	return "f9f5eb2d2e9d8165ca4bd6ccbbb7c0399cebf97e"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
