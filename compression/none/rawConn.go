package none

import (
	"net"
	"time"

	"github.com/libp2p/go-libp2p-core/compression"
)

// ID is the protocol ID for uncompressed connections (for convenience).
const ID = "/compression/none"

var _ compression.CompressedTransport = &Transport{}

type compConn struct {
	raw net.Conn
}

// Transport mock
type Transport struct{}

// New mocks the creation of a new compressed transport.
func New(level int) *Transport {
	return &Transport{}
}

// NewConn wraps the raw connection in a CompressedConn. It does nothing.
func (t *Transport) NewConn(raw net.Conn, isServer bool) (compression.CompressedConn, error) {
	return &compConn{
		raw: raw,
	}, nil
}

func (c *compConn) Write(b []byte) (int, error) {
	return c.raw.Write(b)
}

func (c *compConn) Read(b []byte) (int, error) {
	return c.raw.Read(b)
}

func (c *compConn) Close() error {
	return c.raw.Close()
}

func (c *compConn) LocalAddr() net.Addr {
	return c.raw.LocalAddr()
}

func (c *compConn) RemoteAddr() net.Addr {
	return c.raw.RemoteAddr()
}

func (c *compConn) SetDeadline(t time.Time) error {
	return c.raw.SetDeadline(t)
}

func (c *compConn) SetReadDeadline(t time.Time) error {
	return c.raw.SetReadDeadline(t)
}

func (c *compConn) SetWriteDeadline(t time.Time) error {
	return c.raw.SetWriteDeadline(t)
}
