package proto

const ProtocolGitCommit string = "42805b51055a0f66a04d463fbac46ff8428b69ab"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
