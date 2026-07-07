package proto

const ProtocolGitCommit string = "6041f93eef9788eac649f8c06adaa2e27ade8224"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
