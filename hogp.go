package proto

const ProtocolGitCommit string = "076631e4622a6a9136bf8f9b110130023a8f61b2"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
