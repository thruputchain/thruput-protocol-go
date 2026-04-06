package proto

const ProtocolGitCommit string = "ec3fbdb17bc58cfbb3a01cb7a5b6f64c482d336e"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
