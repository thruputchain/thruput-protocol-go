package proto

const ProtocolGitCommit string = "81e1cb073cfbae98d190ca089f6b764c06ac44cb"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
