package proto

func ProtocolGitCommit() string {
	return "3bf466e918a254e91d8a4880e84bf0366000b5d0"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
