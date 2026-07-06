package proto

const ProtocolGitCommit string = "0575ad7f6faa29a526a432f2823c7db6b60b1e9b"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
