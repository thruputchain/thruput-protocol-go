package proto

const ProtocolGitCommit string = "aacb2076217ecc473e2395cc52ea42c7a8bd3987"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
