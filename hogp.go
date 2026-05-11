package proto

const ProtocolGitCommit string = "4d6568f6b3e65ab6037daf08317033099056c74d"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
