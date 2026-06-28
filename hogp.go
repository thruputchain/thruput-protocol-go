package proto

const ProtocolGitCommit string = "b88e89137ae7c34e4324ba2d253d7bb7f6c7aa0b"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
