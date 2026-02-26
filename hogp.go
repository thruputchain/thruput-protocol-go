package proto

const ProtocolGitCommit string = "a3de2bf4d05955f0a03c2a6269bb7c220b9b78fc"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
