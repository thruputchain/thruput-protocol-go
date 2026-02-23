package proto

const ProtocolGitCommit string = "9b03aecd91f6a77dbf2fe3455401be918f448ec2"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
