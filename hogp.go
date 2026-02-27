package proto

const ProtocolGitCommit string = "d5af2cac57b8db39acf3db51875deac9b9d4aec6"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
