package proto

const ProtocolGitCommit string = "875423e7d5852660aa623db5cc7584599b8e3319"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
