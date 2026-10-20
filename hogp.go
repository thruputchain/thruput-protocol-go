package proto

const ProtocolGitCommit string = "e62b87726d16bbb42e019df4b11ccb27365383ee"

// hogProtocol is the token peers exchange to check they share a protocol
// revision. build.sh emits it at its final width, so this conversion cannot
// fail.
var hogProtocol = ProtocolCommit([]byte("hogp/e62b87726d16bbb"))

// HogProtocol returns the token peers exchange to check they share a protocol
// revision, in the fixed-width form the wire carries. Callers needing text
// take string(t[:]).
func HogProtocol() ProtocolCommit {
	return hogProtocol
}
