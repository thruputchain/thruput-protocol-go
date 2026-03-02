package proto

const ProtocolGitCommit string = "d563bdded40f43f9c2bbf527c045d42e215427c8"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
