package proto

func ProtocolGitCommit() string {
	return "c5838bf894e21dcdc1f14bfc5887d0a8e1d392ed"
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
