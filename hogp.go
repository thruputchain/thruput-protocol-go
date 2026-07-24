package proto

const ProtocolGitCommit string = "1971be04a3129fc178e22fb2dd25c3ddfafe5915"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
