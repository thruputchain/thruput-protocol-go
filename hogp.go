package proto

const ProtocolGitCommit string = "8bddeaf8c2c267daf4648ec9bffd8dd1c452a743"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
