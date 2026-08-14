package proto

const ProtocolGitCommit string = "ad2f4e49ddb1355d805d420dd3ae0a429968e09f"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
