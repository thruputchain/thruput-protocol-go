package proto

const ProtocolGitCommit string = "cff5b0b2fc949092f5fd4fdf74f62ce78dbddb95"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
