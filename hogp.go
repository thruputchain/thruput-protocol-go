package proto

const ProtocolGitCommit string = "a8320d8c4d93a25c547f0bfb471a4d075841af3d"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
