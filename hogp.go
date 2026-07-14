package proto

const ProtocolGitCommit string = "4540b0ba8e9babab63c75427ffe2c677f2ac6625"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
