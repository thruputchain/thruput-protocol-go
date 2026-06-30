package proto

const ProtocolGitCommit string = "fbaf885511e8edfc49cc26c5ba87abf07bbe25a6"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
