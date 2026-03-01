package proto

const ProtocolGitCommit string = "09bddb1686e6e26fb336eaeacbb1f505ec3a4e98"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
