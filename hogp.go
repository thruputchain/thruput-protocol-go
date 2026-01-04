package proto

func ProtocolGitCommit() string {
	return "f8cedfb367e653751e8acd964979b9de8e0db26d"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
