package proto

const ProtocolGitCommit string = "06f3d9fd715ff96ac4fa58fddeae6a7ce253bf1c"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
