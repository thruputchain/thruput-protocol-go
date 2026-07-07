package proto

const ProtocolGitCommit string = "f2ec1ccb1b1b1dc6c0a5109e8a6307407d88bd2a"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
