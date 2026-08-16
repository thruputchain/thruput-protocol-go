package proto

const ProtocolGitCommit string = "66372d1710dea4ace2df7b1fd047e6a484c9dfd6"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
