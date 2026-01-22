package proto

func ProtocolGitCommit() string {
	return "867fa728560427cc2dba5e2d96976dd3a426d0d9"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
