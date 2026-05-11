package proto

const ProtocolGitCommit string = "48cfcd6d43697310829202dd4e893ae20b559584"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
