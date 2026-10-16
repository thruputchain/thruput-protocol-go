package proto

const ProtocolGitCommit string = "b60a5a68d21a0e99d108f9e3e4db24f2f41c26a4"

// hogProtocol is the token peers exchange to check they share a protocol
// revision. build.sh emits it at its final width, so this conversion cannot
// fail.
var hogProtocol = ProtocolCommit([]byte("hogp/b60a5a68d21a0e9"))

// HogProtocol returns the token peers exchange to check they share a protocol
// revision, in the fixed-width form the wire carries. Callers needing text
// take string(t[:]).
func HogProtocol() ProtocolCommit {
	return hogProtocol
}
