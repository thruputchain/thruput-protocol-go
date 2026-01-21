package proto

func ProtocolGitCommit() string {
	return "02882c6194de1ff1d7ec3e39cf226a2dd4b35a65"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
