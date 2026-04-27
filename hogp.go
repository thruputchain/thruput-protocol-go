package proto

const ProtocolGitCommit string = "783d069a2fc4ce8bb0c8ae2a347314eeb27ec59e"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
