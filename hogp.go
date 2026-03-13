package proto

const ProtocolGitCommit string = "a0735303a86dfbff572a224ed9437bfe970fed48"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
