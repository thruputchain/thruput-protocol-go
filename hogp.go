package proto

const ProtocolGitCommit string = "a6913f07a5a76379fbb72ade72d2548035938d0c"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
