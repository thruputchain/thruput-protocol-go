package proto

const ProtocolGitCommit string = "43f401e6cdf7bb9d0c0361ceab06541caa7714ca"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
