package proto

const ProtocolGitCommit string = "adb998a6490cf6eea41161102d3ed0d7b4c552d3"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
