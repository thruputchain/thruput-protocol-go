package proto

const ProtocolGitCommit string = "42f5bbd9a34b6dcbf699a429c2b608b1a07e3f5e"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
