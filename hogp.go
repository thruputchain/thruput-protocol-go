package proto

const ProtocolGitCommit string = "70780a9cf74de76fbc2209477cf76f30cfe88b92"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
