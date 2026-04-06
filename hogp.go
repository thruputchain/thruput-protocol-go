package proto

const ProtocolGitCommit string = "ccdcc4d4f3a56fae632fe3d0e3a17baae0542d74"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
