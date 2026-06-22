package proto

const ProtocolGitCommit string = "93375655862ad5b7d905e57e55c4a4ed87fcc7ec"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
