package proto

const ProtocolGitCommit string = "08b10717cf63f6eeddbc2c83eb56b7d4341911f7"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
