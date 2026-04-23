package proto

const ProtocolGitCommit string = "f325f6cda8ffcd12c086fba81979d11c663b290f"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
