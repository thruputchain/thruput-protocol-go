package proto

const ProtocolGitCommit string = "77397967490b18f2531b734687b2a0ddde7513da"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
