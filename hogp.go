package proto

const ProtocolGitCommit string = "9925e9f66699dbd43f30d323454f1b7d2140e20b"

// hogProtocol is the token peers exchange to check they share a protocol
// revision. build.sh emits it at its final width.
const hogProtocol = "hogp/9925e9f66699dbd"

// HogProtocol returns the token peers exchange to check they share a protocol
// revision. It is always len(ProtocolCommit) bytes, so callers that need the
// wire form can convert it without checking: ProtocolCommit([]byte(...)).
func HogProtocol() string {
	return hogProtocol
}
