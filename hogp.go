package proto

const ProtocolGitCommit string = "5111bf1150ca75479647fdb09f85158456c69579"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
