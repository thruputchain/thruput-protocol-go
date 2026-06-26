package proto

const ProtocolGitCommit string = "7a82d71d369c41853626fa01112ea117066e3738"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
