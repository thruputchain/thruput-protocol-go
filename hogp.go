package proto

const ProtocolGitCommit string = "3e4564dda6476943985b9f959d6db910f0077d80"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
