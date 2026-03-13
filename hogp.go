package proto

const ProtocolGitCommit string = "8197325dfbc8e6fe56fa471d00b416be0f608799"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
