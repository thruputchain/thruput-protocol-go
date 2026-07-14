package proto

const ProtocolGitCommit string = "53d22febcf94f68703a1dd651049df0eb15a5e8d"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
