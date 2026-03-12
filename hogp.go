package proto

const ProtocolGitCommit string = "de5ccf154a9e965721197dcf6af2228d045fda60"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
