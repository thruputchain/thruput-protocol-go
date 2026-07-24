package proto

const ProtocolGitCommit string = "c6645d4a34c34f2f88b9049bca005426edfceca4"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
