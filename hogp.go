package proto

func ProtocolGitCommit() string {
	return "+4c84720a61100cefe346ac0007d441d60ef40bea"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
