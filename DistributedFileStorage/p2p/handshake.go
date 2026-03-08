package p2p

//type Handshake interface {
//	Handshake() error
//}
//
//
//type DefaultHandshake struct {
//}

type HandshakeFunc func(any2 any) error

func NoPHandshakeFunc(any) error {
	return nil
}
