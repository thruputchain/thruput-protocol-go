package proto

func ProtocolGitCommit() string {
	return "aae04738cde1f65fa36fb25ff2e8c87efc994608"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
