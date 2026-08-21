package proto

const ProtocolGitCommit string = "e8833a1dd0d04e47605f7a67d3d24758071aff61"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
