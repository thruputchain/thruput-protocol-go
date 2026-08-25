package proto

const ProtocolGitCommit string = "c2a41277a434944f439e6c46d27d74d2fe39d29b"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
