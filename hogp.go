package proto

const ProtocolGitCommit string = "8fd7eb24ceb2fe0bdf18d28f9a296e27e2bce50d"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
