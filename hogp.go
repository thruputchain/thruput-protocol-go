package proto

const ProtocolGitCommit string = "f8d3703a388e7b4332b9642fdc6007545411a2e5"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
