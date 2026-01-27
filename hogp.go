package proto

func ProtocolGitCommit() string {
	return "0fbf2fb8e42d73168da8e949f355a27d6793b9a5"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
