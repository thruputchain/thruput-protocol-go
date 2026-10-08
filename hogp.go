package proto

const ProtocolGitCommit string = "6bfe776bf040665ce064e81087750aa1f6f4767b"

// hogProtocol is the token peers exchange to check they share a protocol
// revision. build.sh emits it at its final width, so this conversion cannot
// fail.
var hogProtocol = ProtocolCommit([]byte("hogp/6bfe776bf040665"))

// HogProtocol returns the token peers exchange to check they share a protocol
// revision, in the fixed-width form the wire carries. Callers needing text
// take string(t[:]).
func HogProtocol() ProtocolCommit {
	return hogProtocol
}
