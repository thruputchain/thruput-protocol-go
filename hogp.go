package proto

const ProtocolGitCommit string = "561d2617ace29b223f42d9e184965aad6229a2a2"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
