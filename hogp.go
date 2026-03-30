package proto

const ProtocolGitCommit string = "5fef82c8de4e6b4b1aaa48632feeb85502d2ed8f"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
