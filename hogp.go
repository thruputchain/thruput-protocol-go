package proto

func ProtocolGitCommit() string {
	return "5d9117a402c4b0ba42368c5fe58df43cb8832bc6"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
