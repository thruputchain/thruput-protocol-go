package proto

const ProtocolGitCommit string = "78e402a78dd7b1bcfca095d308d4a51bd6919ddb"

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
