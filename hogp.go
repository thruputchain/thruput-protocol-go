package proto

const ProtocolGitCommit string = "dd299e6a9750b88cfc7cf1352059c670d5119a34"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
