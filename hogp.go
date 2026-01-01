package proto

func ProtocolGitCommit() string {
	return "93375bd1eb5ad77f6182b2831dd8e389882fcc0f"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
