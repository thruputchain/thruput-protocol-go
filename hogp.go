package proto

const ProtocolGitCommit string = "35c154e2a302875e472b6836df2efa0cb8dd1977"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
