package proto

const ProtocolGitCommit string = "020ed23d4e2fedda0d31790370ef532acd3827e2"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
