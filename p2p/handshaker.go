package p2p

type Handshakerfunc func(any)error

func NOPHandshakerfunc (any) error {return nil}