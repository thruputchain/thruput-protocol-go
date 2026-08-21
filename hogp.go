package proto

const ProtocolGitCommit string = "3d80720d875ca6ea2e442c70c934d6cad92550ad"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
