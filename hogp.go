package proto

const ProtocolGitCommit string = "042a6625b0185e7167d68180ebb3d74a5bd27c48"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
