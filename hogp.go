package proto

const ProtocolGitCommit string = "8bab2528f46425b2b448fa85560c9c1b897f5631"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
