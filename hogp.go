package proto

const ProtocolGitCommit string = "0784e2adc9effe56e110d97466854c36274fbfad"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
