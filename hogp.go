package proto

const ProtocolGitCommit string = "06a247e3991c67c82dad34c7277928e2d99d026e"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
