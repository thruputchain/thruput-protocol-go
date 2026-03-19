package proto

const ProtocolGitCommit string = "fee889d92ef698a6ad27fd944095dd6c4a55d1c4"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
