package proto

const ProtocolGitCommit string = "65d78d04c6e23d5d6c68317bbf88ad48160b1500"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
