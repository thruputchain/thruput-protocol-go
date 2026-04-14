package proto

const ProtocolGitCommit string = "ddf43a2a8dca78c97d9169cb8b34f8d9bdc7a979"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
