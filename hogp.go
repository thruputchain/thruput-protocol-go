package proto

func ProtocolGitCommit() string {
	return "a4d690d41a8146147854d545d1d3c8ed993938b6"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
