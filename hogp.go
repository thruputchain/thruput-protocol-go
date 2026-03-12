package proto

const ProtocolGitCommit string = "b72e4898be8bc68f9940e2597c08c6b9cd95a229"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
