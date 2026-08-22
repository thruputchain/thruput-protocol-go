package proto

const ProtocolGitCommit string = "db2688fab2be2919cd47ebcd0fb607be28360ad9"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
