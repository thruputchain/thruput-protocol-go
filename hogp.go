package proto

const ProtocolGitCommit string = "64c7fdda31bfc700670d4ef7740b28940933fdf2"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
