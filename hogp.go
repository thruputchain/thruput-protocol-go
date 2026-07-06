package proto

const ProtocolGitCommit string = "0914d97c32c9f5a2671b3238b62b7dd69383fe75"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
