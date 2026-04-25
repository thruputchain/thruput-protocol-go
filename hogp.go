package proto

const ProtocolGitCommit string = "dcd4e8a384ed674ab2997e85765be0b2643b25b2"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
