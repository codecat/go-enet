package enet

// #cgo LDFLAGS: -lenet
// #cgo windows CFLAGS: -Ienet/include/
// #cgo windows LDFLAGS: -Lenet/ -lWs2_32 -lWinmm
// #include <enet/enet.h>
import "C"

// Initialize enet
func Initialize() {
	C.enet_initialize()
}

// Deinitialize enet
func Deinitialize() {
	C.enet_deinitialize()
}
