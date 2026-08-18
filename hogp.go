package proto

const ProtocolGitCommit string = "ff9c6213c0a60c26a6afa85ae0c2cc6aca9553be"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
