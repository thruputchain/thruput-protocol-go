package proto

const ProtocolGitCommit string = "bc8d9dc63bd3e40e399cc5c2ad0f69a69cae85e0"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
