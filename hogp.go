package proto

const ProtocolGitCommit string = "53c406dfb3b3e36c37b04140d22b168036b0ae61"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
