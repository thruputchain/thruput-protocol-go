package proto

const ProtocolGitCommit string = "68837ec7da243f7701a0067281e17d12a34b9956"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
