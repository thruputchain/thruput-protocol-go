package proto

const ProtocolGitCommit string = "4b2849e5a6a0dfdf62a7e355e738d5644c6d63ca"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
