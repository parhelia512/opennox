package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func nox_xxx_netRecv_552020(s nox_socket_t, buf *byte, len_ int32, from *nox_net_sockaddr_in) int32 {
	var n int32 = mix_recvfrom(s, buf, len_, (*nox_net_sockaddr)(unsafe.Pointer(from)))
	if n != -1 {
		var ns *nox_net_struct_t = nox_xxx_netStructByAddr_551E60(from)
		if ns != nil {
			if int32(ns.xor_key) != 0 {
				nox_xxx_cryptXor_56FDD0(int8(ns.xor_key), (*uint8)(unsafe.Pointer(buf)), n)
			}
		}
	}
	if noxflags.HasGame(noxflags.GameHost) {
		return n
	}
	var r int32 = randomIntMinMax(1, 99)
	if r >= *memmap.PtrInt32(6112660, 2495940) {
		return n
	}
	return 0
}
