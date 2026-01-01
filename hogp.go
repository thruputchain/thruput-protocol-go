package proto

func ProtocolGitCommit() string {
	return "c9dde7adeaf7bddf0cadd68a5b845c13ceafee43"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
