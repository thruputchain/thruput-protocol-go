package proto

const ProtocolGitCommit string = "99d7dfc89a6a88e6ed99710e50b5bb729179993a"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
