package proto

const ProtocolGitCommit string = "4f499ecfa238a35608804482efd1b4ba8f9e271a"

// hogProtocol is the token peers exchange to check they share a protocol
// revision. build.sh emits it at its final width, so this conversion cannot
// fail.
var hogProtocol = ProtocolCommit([]byte("hogp/4f499ecfa238a35"))

// HogProtocol returns the token peers exchange to check they share a protocol
// revision, in the fixed-width form the wire carries. Callers needing text
// take string(t[:]).
func HogProtocol() ProtocolCommit {
	return hogProtocol
}
