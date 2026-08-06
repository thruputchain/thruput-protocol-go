package proto

const ProtocolGitCommit string = "07c15ee8e677bf7d4627eb890a837b7b0fa21b81"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
