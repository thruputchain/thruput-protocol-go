package proto

const ProtocolGitCommit string = "c484d53fd613a406e2c2ee42765b0cb94013b06f"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
