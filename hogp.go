package proto

const ProtocolGitCommit string = "e64690207c6901039b2fd3e4d927e068bd23c2ff"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
