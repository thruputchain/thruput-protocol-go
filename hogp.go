package proto

const ProtocolGitCommit string = "034b8b7d30d63c18fa53ddf41ea3ef75396cd7ef"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
