package proto

const ProtocolGitCommit string = "0cd3c454cdeb3298a720c0e992a00e46b2b325f3"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
