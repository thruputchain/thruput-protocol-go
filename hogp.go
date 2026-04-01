package proto

const ProtocolGitCommit string = "a4fe8b41398510893f2a66190c1ae68d935c90b9"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
