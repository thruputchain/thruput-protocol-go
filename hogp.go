package proto

const ProtocolGitCommit string = "833d749f5fd9e22bdd7a993a1bccb702dc49082e"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
