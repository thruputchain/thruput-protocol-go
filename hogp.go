package proto

const ProtocolGitCommit string = "02555a4d21e8d4c5e9390e10c2c4138c03a523fd"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
