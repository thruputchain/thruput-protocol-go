package proto

func ProtocolGitCommit() string {
	return "5a169d56f1042e09dfcf31647c37146b87e3c463"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
