package proto

const ProtocolGitCommit string = "4debaa098bf15b957e156ec141cf3e0c6abad242"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
