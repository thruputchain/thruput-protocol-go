package proto

const ProtocolGitCommit string = "cc14514eacb15abfc036f61073f91ecde83368ae"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
