package proto

const ProtocolGitCommit string = "f18a44980df16d7354f13921c7c22085f7ec6ee3"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
