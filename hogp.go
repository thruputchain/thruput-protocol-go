package proto

func ProtocolGitCommit() string {
	return "55efdffd46673548e652ff52dfef68a961f7b6df"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
