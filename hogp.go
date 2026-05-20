package proto

const ProtocolGitCommit string = "6eec13c0a30558e7b77f7d4303ad5a1c2497423f"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
