package proto

const ProtocolGitCommit string = "0c2f1fa9acf35470b8031b0aa2a6ee9aae7e45c8"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
