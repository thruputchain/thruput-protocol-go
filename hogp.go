package proto

const ProtocolGitCommit string = "f49d183a02dafe608b15e6ad49c3c8179abf467d"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
