package proto

const ProtocolGitCommit string = "da38733c7610d7b53446a1e5ac0d0f08a97320c0"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
