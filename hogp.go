package proto

func ProtocolGitCommit() string {
	return "be4e81386a52b2099ce27bb67d6279afdd55ad04"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
