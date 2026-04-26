package proto

const ProtocolGitCommit string = "8baa92205b45ef3056639a54922757795d2a8264"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
