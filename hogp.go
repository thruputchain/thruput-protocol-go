package proto

const ProtocolGitCommit string = "c8ebf52df8de818c4a386727d8f742ce66bf3d01"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
