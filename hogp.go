package proto

func ProtocolGitCommit() string {
	return "69f99a09866c9168ac120b1d459c674a3288f689"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
