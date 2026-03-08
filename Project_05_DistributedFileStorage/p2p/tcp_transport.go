package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer represent the remote node over a TCP established connection.
type TCPPeer struct {
	// conn is the underlying connection of the peer
	conn net.Conn

	// if we dial and retrieve a conn => outbound = true
	// if we accept and retrieve a conn => outbound = false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransport struct {
	ListenAddress string
	Listener      net.Listener
	HandshakeFunc HandshakeFunc
	Mu            sync.RWMutex
	Peers         map[net.Addr]Peer
}

func NewTCPTransport(listenAddress string) *TCPTransport {
	return &TCPTransport{
		ListenAddress: listenAddress,
		HandshakeFunc: NoPHandshakeFunc,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.Listener, err = net.Listen("tcp", t.ListenAddress)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()
	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.Listener.Accept()
		if err != nil {
			fmt.Println("TCP accept error:", err)
		}

		go t.handlerConn(conn)
	}
}

func (t *TCPTransport) handlerConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)

	if err := t.HandshakeFunc(conn); err != nil {
		fmt.Println("TCP handshake error:", err)
		return
	}

	fmt.Printf("New incoming connection %v\n", peer)
}
