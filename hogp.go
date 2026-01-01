package proto

func ProtocolGitCommit() string {
	return "c125db1dc07ac59766c61063010ef5a719b4bb45"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
