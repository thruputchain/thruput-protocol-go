package proto

func ProtocolGitCommit() string {
	return "b9f13efa6eda7bbc276e2f1bd0fc67ac14fcc677"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
