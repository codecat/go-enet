package enet

// #include <enet/enet.h>
import "C"

// Peer is a peer which data packets may be sent or received from
type Peer interface {
	GetAddress() Address
}

type enetPeer struct {
	cPeer *C.struct__ENetPeer
}

func (peer enetPeer) GetAddress() Address {
	return &enetAddress{
		cAddr: peer.cPeer.address,
	}
}
