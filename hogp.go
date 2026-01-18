package proto

func ProtocolGitCommit() string {
	return "e413dc9f7174cf4c6a5d87a1f113761348da4030"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
