package proto

func ProtocolGitCommit() string {
	return "c738324745ec3ce366fd542a79ac76b7b7ec4c1e"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
