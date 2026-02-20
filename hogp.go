package proto

const ProtocolGitCommit string = "62830dc03a63f17ae6ff4cfa1c11e98b6139719f"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
