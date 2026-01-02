package proto

func ProtocolGitCommit() string {
	return "394557a7b22cbd1326a4ea3089ccc322f3b1052d"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
