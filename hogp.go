package proto

const ProtocolGitCommit string = "86459afcdfa38fab7473fb237804f19b2568aac5"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
