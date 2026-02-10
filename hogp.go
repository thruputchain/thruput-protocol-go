package proto

func ProtocolGitCommit() string {
	return "7adb4396b133e65df33f113a9c644bf4931a0ee5"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
