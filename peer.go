package enet

// #include <enet/enet.h>
import "C"

// Peer is a peer which data packets may be sent or received from
type Peer interface {
	GetAddress() Address

	Disconnect(data uint32)
	DisconnectNow(data uint32)
	DisconnectLater(data uint32)

	SendBytes(data []byte, channel uint8, flags PacketFlags) error
	SendString(str string, channel uint8, flags PacketFlags) error
	SendPacket(packet Packet, channel uint8) error
}

type enetPeer struct {
	cPeer *C.struct__ENetPeer
}

func (peer enetPeer) GetAddress() Address {
	return &enetAddress{
		cAddr: peer.cPeer.address,
	}
}

func (peer enetPeer) Disconnect(data uint32) {
	C.enet_peer_disconnect(
		peer.cPeer,
		(C.enet_uint32)(data),
	)
}

func (peer enetPeer) DisconnectNow(data uint32) {
	C.enet_peer_disconnect_now(
		peer.cPeer,
		(C.enet_uint32)(data),
	)
}

func (peer enetPeer) DisconnectLater(data uint32) {
	C.enet_peer_disconnect_later(
		peer.cPeer,
		(C.enet_uint32)(data),
	)
}

func (peer enetPeer) SendBytes(data []byte, channel uint8, flags PacketFlags) error {
	packet, err := NewPacket(data, flags)
	if err != nil {
		return err
	}
	return peer.SendPacket(packet, channel)
}

func (peer enetPeer) SendString(str string, channel uint8, flags PacketFlags) error {
	packet, err := NewPacket([]byte(str), flags)
	if err != nil {
		return err
	}
	return peer.SendPacket(packet, channel)
}

func (peer enetPeer) SendPacket(packet Packet, channel uint8) error {
	C.enet_peer_send(
		peer.cPeer,
		(C.enet_uint8)(channel),
		packet.(enetPacket).cPacket,
	)
	return nil
}
