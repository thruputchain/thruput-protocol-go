package proto

const ProtocolGitCommit string = "81e4b2f7b632249e3f46cfd7f0126fca5fcb8efb"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
