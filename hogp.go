package proto

const ProtocolGitCommit string = "6d1e8e6eca4eaecb54ba6d72d84887333f7be149"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
