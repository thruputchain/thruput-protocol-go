package proto

const ProtocolGitCommit string = "c19a1132d4cc2557f7193563d5a72f535dee4188"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
