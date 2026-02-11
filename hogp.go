package proto

func ProtocolGitCommit() string {
	return "8bdb3ee444a1da777f5292f8ff620c438acf4964"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
