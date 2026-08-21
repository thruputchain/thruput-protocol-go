package proto

const ProtocolGitCommit string = "d2534e6e8f6225cb1d67db5be906cf14a26be4ec"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
