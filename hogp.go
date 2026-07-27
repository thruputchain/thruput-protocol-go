package proto

const ProtocolGitCommit string = "edec479948831f22da913b5f9a779464c8daa820"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
