package proto

func ProtocolGitCommit() string {
	return "e304d40b98d7e4eba6c5a9ddd2957f2f15ed1e67"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
