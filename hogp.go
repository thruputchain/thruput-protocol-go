package proto

const ProtocolGitCommit string = "d9ee802204b9d4ebdb4b6af18afb292dd07106b7"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
