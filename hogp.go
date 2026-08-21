package proto

const ProtocolGitCommit string = "1046ed1ba1e1894fc092e2b23ab50e36ff9b2f05"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
