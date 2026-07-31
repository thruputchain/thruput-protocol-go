package proto

const ProtocolGitCommit string = "bd86a953ae024827fcd320fc12f5ca8ebf3cfa28"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
