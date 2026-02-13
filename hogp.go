package proto

func ProtocolGitCommit() string {
	return "0266d2ad9748bf2f4985f4b9595542717b0d6aec"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
