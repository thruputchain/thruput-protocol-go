package proto

const ProtocolGitCommit string = "1ad984965cfb63fd24b9d76cd5bcdb89bd703e5b"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
