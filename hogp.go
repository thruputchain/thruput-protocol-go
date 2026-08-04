package proto

const ProtocolGitCommit string = "8e1aaa5d00fce0b30881ec8dfce2cf9a2e8e0e08"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
