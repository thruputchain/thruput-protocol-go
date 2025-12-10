package proto

func ProtocolGitCommit() string {
	return "+af8159777b201bedb9f84bf3f12677f5359f80ad"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
