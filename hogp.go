package proto

const ProtocolGitCommit string = "9a6db0a237bd89de8a59ad42492e665ea483edc4"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
