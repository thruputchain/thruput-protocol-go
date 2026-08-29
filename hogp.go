package proto

const ProtocolGitCommit string = "bd860533b6e27bba8d0dac0bf7045f5e935a08ce"

// HogProtocol returns the token peers exchange to check they share a protocol
// revision. It is sized to fill proto.ProtocolCommit exactly: the XDR type is
// a fixed 20-byte array, and the "hogp/" prefix leaves 15 for the commit id.
func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:15]
}

// HogProtocolCommit returns the same token in the fixed-width form the wire
// carries. It never fails: the conversion is total for a 20-byte string.
func HogProtocolCommit() ProtocolCommit {
	return ProtocolCommit([]byte(HogProtocol()))
}
