package proto

const ProtocolGitCommit string = "d9c1e599a5cd4288ade651b2baedf7b691862ed1"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
