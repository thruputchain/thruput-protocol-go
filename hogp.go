package proto

const ProtocolGitCommit string = "49b58b5a214ec5b0be719b238d2b8ce790d6e82d"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
