package proto

const ProtocolGitCommit string = "5778ec23fb93d960488bff89e09fdf813c6fc834"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
