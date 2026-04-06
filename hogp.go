package proto

const ProtocolGitCommit string = "1bb23b9e8579ebcba4c2505ffd5dc9f213390877"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
