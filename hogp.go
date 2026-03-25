package proto

const ProtocolGitCommit string = "e357df1421961050a75a124458d114d9c81a97f4"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
