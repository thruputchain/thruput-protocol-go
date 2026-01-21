package proto

func ProtocolGitCommit() string {
	return "0efbdea6e702edaaf0f0ec89a9e3316de18b37b8"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
