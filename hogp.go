package proto

const ProtocolGitCommit string = "f2d5fe44b474520300e09559ae61efffd28109be"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
