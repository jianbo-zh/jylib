package typec

import (
	"context"
	"net"
	"time"

	"github.com/quic-go/quic-go"
)

// QUIC 监听器适配器（将 quic.Listener 转换为 net.Listener）
type QuicListener struct {
	quic.Listener
}

func (q *QuicListener) Accept() (net.Conn, error) {
	sess, err := q.Listener.Accept(context.Background())
	if err != nil {
		return nil, err
	}
	stream, err := sess.AcceptStream(context.Background())
	if err != nil {
		return nil, err
	}
	return &QuicStreamConn{Session: sess, Stream: stream}, nil
}

func (q *QuicListener) Close() error {
	return q.Listener.Close()
}

func (q *QuicListener) Addr() net.Addr {
	return q.Listener.Addr()
}

// QuicStreamConn quic 流包装成 net.Conn
type QuicStreamConn struct {
	Session quic.Connection
	Stream  quic.Stream
}

func (c *QuicStreamConn) Read(b []byte) (int, error)         { return c.Stream.Read(b) }
func (c *QuicStreamConn) Write(b []byte) (int, error)        { return c.Stream.Write(b) }
func (c *QuicStreamConn) Close() error                       { return c.Stream.Close() }
func (c *QuicStreamConn) LocalAddr() net.Addr                { return c.Session.LocalAddr() }
func (c *QuicStreamConn) RemoteAddr() net.Addr               { return c.Session.RemoteAddr() }
func (c *QuicStreamConn) SetDeadline(t time.Time) error      { return nil }
func (c *QuicStreamConn) SetReadDeadline(t time.Time) error  { return c.Stream.SetReadDeadline(t) }
func (c *QuicStreamConn) SetWriteDeadline(t time.Time) error { return c.Stream.SetWriteDeadline(t) }
