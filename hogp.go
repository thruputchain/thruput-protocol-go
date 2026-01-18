package proto

func ProtocolGitCommit() string {
	return "e76d788414d68e19fcb5d584d5e0787e20888997"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
