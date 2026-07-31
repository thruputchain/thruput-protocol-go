package proto

const ProtocolGitCommit string = "f9b80a5f9842bc6d63321d89de7c54b3d4b4f904"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
