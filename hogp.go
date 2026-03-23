package proto

const ProtocolGitCommit string = "95cb0efd2ecf9597e7a4f3e9ad4dc949d73be7f4"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
