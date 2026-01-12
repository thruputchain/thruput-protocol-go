package proto

func ProtocolGitCommit() string {
	return "382589d1b73b9b9352ecc06347e0cd306b480b94"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
