package proto

const ProtocolGitCommit string = "5873430685f3efb0121fb4acda95f484e63cad0c"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
