package proto

func ProtocolGitCommit() string {
	return "97cca4cf5099675c07b4a5e657211014eaee82e1"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
