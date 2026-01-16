package proto

func ProtocolGitCommit() string {
	return "c282cb10a726a6e6922058decd2e9c9ce5ceaf4f"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
