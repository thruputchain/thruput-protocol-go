package proto

const ProtocolGitCommit string = "9453f5005813c1dca37993647a0d4cffa7b07211"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
