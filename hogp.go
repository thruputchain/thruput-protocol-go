package proto

const ProtocolGitCommit string = "e4a1ca88f0ca463159ca8ff2fda3a281a761ae9c"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
