package proto

const ProtocolGitCommit string = "8331a3a50b8807fdc59b41a1ddda25f71dd6531a"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
