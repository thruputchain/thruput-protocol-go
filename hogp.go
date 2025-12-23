package proto

func ProtocolGitCommit() string {
	return "6abfedebdf11d2f41bddde90aea15d3a916834cb"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
