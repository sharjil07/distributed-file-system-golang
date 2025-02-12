package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer represents the remote node over a TCP established connection.
type TCPPeer struct {
	//conn is the underlying connection of the peer
	conn net.Conn
	//if we dial and retrieve a conn=>outbound==true
	//if we accpet and retrieve a conn=>outbound==false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPtransport struct {
	listenAddress string
	listener      net.Listener
	shakeHands    Handshakerfunc
	decoder       Decoder
	mu            sync.Mutex
	peers         map[net.Addr]Peer
}

func NewTCPtransport(listenAddr string) *TCPtransport {
	return &TCPtransport{
		shakeHands:    NOPHandshakerfunc,
		listenAddress: listenAddr,
	}
}

func (t *TCPtransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}
	go t.startAcceptLoop()

	return nil

}

func (t *TCPtransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()

		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
		}
		fmt.Printf("new incoming connection %+v\n", conn)
		go t.handleConn(conn)
	}
}

type Temp struct{}

func (t *TCPtransport) handleConn(conn net.Conn) {

	peer := NewTCPPeer(conn, true)

	if err := t.shakeHands(conn); err != nil {
       conn.Close()
	   fmt.Printf("TCP handshake err: %s\n",err)
	   return
	}
	msg := &Temp{}

	for {
		if err := t.decoder.Decode(conn, msg); err != nil {
			fmt.Printf("TCP error: %s\n", err)
			continue
		}
	}

	fmt.Printf("new incoming connection %+v\n", peer)
}
