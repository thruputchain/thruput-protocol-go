package proto

func ProtocolGitCommit() string {
	return "40796df7cedff9300228e4d85e7d471cd0fc4d65"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
