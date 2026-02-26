package proto

const ProtocolGitCommit string = "405bb592aa5be4ccc8ac4c2c4eb0173570f01d33"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
