package proto

const ProtocolGitCommit string = "2c7f458fb5e52414dd387535c046832ae3d9ef75"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
