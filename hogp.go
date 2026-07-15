package proto

const ProtocolGitCommit string = "ffc78290e43d83c3ac665e841e4e69717753a142"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
