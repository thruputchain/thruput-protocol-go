package proto

const ProtocolGitCommit string = "d505437361fd26714046e817b64f072d9532b89a"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
