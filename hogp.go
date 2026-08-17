package proto

const ProtocolGitCommit string = "24984fdf5f61d23d12a2afa6e1e1aa569e5c63cd"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
