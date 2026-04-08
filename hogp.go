package proto

const ProtocolGitCommit string = "5fac47350f39536fd7fbf9156b03e718b938bbda"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
