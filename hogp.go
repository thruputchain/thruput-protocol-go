package proto

const ProtocolGitCommit string = "936a2322ecb046f0db346ad8f5b09ce1b2c72373"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
