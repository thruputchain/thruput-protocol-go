package proto

const ProtocolGitCommit string = "2cc7b4f7116fe8385c85930cb8e580ab8320d28f"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
