package proto

const ProtocolGitCommit string = "2117a0677c3daa269c60bcad6262d8ac9706dc2d"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
