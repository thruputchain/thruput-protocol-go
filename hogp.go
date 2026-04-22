package proto

const ProtocolGitCommit string = "047a8651709b4f8fa39d054e197395fe5ec1ee52"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
