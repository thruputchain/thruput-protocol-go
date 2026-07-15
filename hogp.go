package proto

const ProtocolGitCommit string = "076483f157d23b1c214e3365477b3e738936b2fd"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
