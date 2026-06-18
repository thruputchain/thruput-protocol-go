package proto

const ProtocolGitCommit string = "1cf254e7101b49ffa7e1cc9fbf552a3b5521cfa5"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
