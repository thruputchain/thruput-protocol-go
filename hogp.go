package proto

const ProtocolGitCommit string = "830b7810ee934cd0fb35c7c6b94aaa185a6c554b"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
