package proto

const ProtocolGitCommit string = "499e04995548c2ada0fda1279f7097a31bc21f62"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
