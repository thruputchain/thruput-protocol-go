package proto

const ProtocolGitCommit string = "1b74ddb653500ff696a38c5fe55639e17d163941"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
