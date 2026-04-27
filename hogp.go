package proto

const ProtocolGitCommit string = "1a648ede2e50ea080211c05a859d1a2d1d3c572d"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
