package proto

const ProtocolGitCommit string = "44912a4ff22ceb6e965ca4632a2057571c772d28"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
