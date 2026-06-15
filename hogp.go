package proto

const ProtocolGitCommit string = "595adb97fde8219f41ac8cad3cb475a90e4dc81c"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
