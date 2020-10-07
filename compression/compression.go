// Package compression provides compressed connection and transport interfaces for libp2p.
package compression

import (
	"net"
)

// CompressedConn wraps in a compressed connection a raw connection.
type CompressedConn interface {
	net.Conn
}

// A CompressedTransport compresses and uncompresses
// the data exchanged with the raw connection.
type CompressedTransport interface {
	// NewConn wraps a raw connection and applies compression.
	NewConn(raw net.Conn, isServer bool) (CompressedConn, error)
}
