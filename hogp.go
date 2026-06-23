package proto

const ProtocolGitCommit string = "afb6b5114f03aac3458af2022b5dc1ee88d004a0"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
