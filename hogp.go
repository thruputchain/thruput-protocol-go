package proto

const ProtocolGitCommit string = "da37ff4c469a7ef3bd83f80f7fc81904f41485f4"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
