package proto

const ProtocolGitCommit string = "3f6788aa81bc933ca3d77dae2af4f69e196a2216"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
