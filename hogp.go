package proto

func ProtocolGitCommit() string {
	return "af97108f5b96814380604a63e11424a45a6109c2"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
