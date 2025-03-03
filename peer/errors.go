package peer

import "errors"

var (
	// ErrEmptyPeerID is returned if the peer id cannot be empty.
	ErrEmptyPeerID = errors.New("peer id cannot be empty")
	// ErrBodyEmpty is returned if the message body was empty.
	ErrBodyEmpty = errors.New("message body cannot be empty")
	// ErrSignatureInvalid is returned for an invalid signature.
	ErrSignatureInvalid = errors.New("message signature invalid")
	// ErrShortMessage is returned if a message is too short.
	ErrShortMessage = errors.New("message too short")
	// ErrNoPrivKey is returned if the private key is not available.
	ErrNoPrivKey = errors.New("private key not available for peer")
	// ErrInvalidEd25519PubKeyForCurve25519 is returned if a public key cannot be used for curve25519.
	ErrInvalidEd25519PubKeyForCurve25519 = errors.New("invalid ed25519 public key for curve25519")
)
