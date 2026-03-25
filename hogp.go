package proto

const ProtocolGitCommit string = "0c8553e902fa9dcbc127d8bc6eff8f0a471a9f85"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
