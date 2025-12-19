package proto

func ProtocolGitCommit() string {
	return "bee5b6ee94d1bb6990e757573d1ae9969f99a1f5"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
