package proto

func ProtocolGitCommit() string {
	return "ddde85cbf0572ee66b2e86efe1ce602efd1f3875"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
