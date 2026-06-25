package proto

const ProtocolGitCommit string = "959c4b24effdf4b2f20e228c76f53bdd34ec7566"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
