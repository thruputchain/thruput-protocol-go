package proto

const ProtocolGitCommit string = "b66a7855a5f11eac12f4536c796127b856e7d9d8"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
