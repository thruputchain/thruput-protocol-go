package proto

func ProtocolGitCommit() string {
	return "+844283d5a35c4d39b91b18c3dc88bf5dc627114b"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
