package proto

const ProtocolGitCommit string = "cf9275b0e179fb2402e5d2a6534613ffe6baa243"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
