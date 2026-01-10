package proto

func ProtocolGitCommit() string {
	return "4707ff236ed28d7a44d9394e880b947f266e81eb"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
