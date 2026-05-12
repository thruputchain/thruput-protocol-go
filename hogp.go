package proto

const ProtocolGitCommit string = "e282f483a650a7d81327ffcad1ed2df8b82037a2"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
