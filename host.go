package enet

// #include <enet/enet.h>
import "C"
import (
	"errors"
)

// Host for communicating with peers
type Host interface {
	Destroy()
	Service(timeout uint32) Event

	Connect(addr Address, channelCount int, data uint32) (Peer, error)
}

type enetHost struct {
	cHost *C.struct__ENetHost
}

func (host *enetHost) Destroy() {
	C.enet_host_destroy(host.cHost)
}

func (host *enetHost) Service(timeout uint32) Event {
	ret := &enetEvent{}
	C.enet_host_service(
		host.cHost,
		&ret.cEvent,
		(C.enet_uint32)(timeout),
	)
	return ret
}

func (host *enetHost) Connect(addr Address, channelCount int, data uint32) (Peer, error) {
	peer := C.enet_host_connect(
		host.cHost,
		&(addr.(*enetAddress)).cAddr,
		(C.size_t)(channelCount),
		(C.enet_uint32)(data),
	)

	if peer == nil {
		return nil, errors.New("couldn't connect to foreign peer")
	}

	return enetPeer{
		cPeer: peer,
	}, nil
}

// NewHost creats a host for communicating to peers
func NewHost(addr Address, peerCount, channelLimit uint64, incomingBandwidth, outgoingBandwidth uint32) (Host, error) {
	var cAddr *C.struct__ENetAddress
	if addr != nil {
		cAddr = &(addr.(*enetAddress)).cAddr
	}

	host := C.enet_host_create(
		cAddr,
		(C.size_t)(peerCount),
		(C.size_t)(channelLimit),
		(C.enet_uint32)(incomingBandwidth),
		(C.enet_uint32)(outgoingBandwidth),
	)

	if host == nil {
		return nil, errors.New("unable to create host")
	}

	return &enetHost{
		cHost: host,
	}, nil
}
