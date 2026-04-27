package proto

const ProtocolGitCommit string = "407da379a0ecc4df1544f26157e070d7130ffd1c"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
