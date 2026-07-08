package proto

const ProtocolGitCommit string = "8e72064a3e033565721cf5a6fefd7650476ea484"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
