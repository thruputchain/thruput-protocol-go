package proto

func ProtocolGitCommit() string {
<<<<<<< HEAD
	return "02882c6194de1ff1d7ec3e39cf226a2dd4b35a65"
=======
	return "0efbdea6e702edaaf0f0ec89a9e3316de18b37b8"
>>>>>>> 81b8850 (bump to protocol@0efbdea)
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
