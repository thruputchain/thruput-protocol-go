package proto

func ProtocolGitCommit() string {
	return "3af2abcc843c3d9a197ed2f7b3490f821af3d764"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
