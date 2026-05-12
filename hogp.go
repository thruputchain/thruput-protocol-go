package proto

const ProtocolGitCommit string = "fbb513b53e88838d1689dbec94cafab45f9959c0"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
