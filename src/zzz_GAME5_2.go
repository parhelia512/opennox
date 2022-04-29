package opennox

import (
	"github.com/gotranspile/cxgo/runtime/cmath"
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/gotranspile/cxgo/runtime/stdio"
	"github.com/noxworld-dev/opennox/v1/common/alloc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"math"
	"unsafe"
)

var nox_client_onLobbyServer_2513928 func(*byte, uint16, *byte, *byte) int32 = nil
var nox_net_struct_arr [128]*nox_net_struct_t = [128]*nox_net_struct_t{}
var nox_net_struct2_arr [128]nox_net_struct2_t = [128]nox_net_struct2_t{}
var dword_5d4594_2523912 unsafe.Pointer = nil
var dword_5d4594_3843632 uint32 = 0

func nox_xxx_netHandlerDefXxx_553D60(a1 uint32, a2 *byte, a3 int32, a4 unsafe.Pointer) int32 {
	return 0
}
func nox_xxx_netHandlerDefYyy_553D70(a1 uint32, a2 *byte, a3 int32, a4 unsafe.Pointer) int32 {
	return 0
}
func nox_xxx_netSendReadPacket_5528B0(a1 uint32, a2 int8) int32 {
	var (
		v2  uint32
		v5  uint32
		v7  int32
		v10 uint32
		v11 int32
		v12 uint32
		v13 *byte
		v14 int32
		v15 uint32
		v18 uint32
	)
	v2 = a1
	if a1 >= NOX_NET_STRUCT_MAX {
		return -3
	}
	var ns *nox_net_struct_t = nox_net_struct_arr[a1]
	if ns == nil {
		return -3
	}
	if ns.id == -1 {
		if ns.field_19 != 0 {
			compatSetEvent(unsafe.Pointer(uintptr(ns.field_34)))
			return 0
		}
		v5 = 0
		v18 = NOX_NET_STRUCT_MAX
	} else {
		v5 = a1
		v18 = a1 + 1
		v2 = uint32(ns.id)
	}
	v15 = v2
	v7 = int32(v18)
	for j := int32(int32(v5)); uint32(j) < v18; j++ {
		var ns2 *nox_net_struct_t = nox_net_struct_arr[j]
		if ns2 == nil || uint32(ns2.id) != v15 {
			continue
		}
		nox_xxx_netSend_5552D0(uint32(j), 0, 0)
		v10 = compatWaitForSingleObject(ns2.mutex_yyy, 1000)
		v7 = int32(v10)
		if v10 == math.MaxUint32 || v10 == 258 {
			return -16
		}
		if (int32(a2) & 1) == 0 {
			v11 = ns2.func_xxx(uint32(j), ns2.data_2_xxx, int32(uintptr(unsafe.Pointer(ns2.data_2_end)))-int32(uintptr(unsafe.Pointer(ns2.data_2_xxx))), ns2.data_3)
			v7 = v11
			if v11 > 0 {
				v12 = uint32(v11) + uint32(uintptr(unsafe.Pointer(ns2.data_2_xxx)))
				if uintptr(unsafe.Pointer(uintptr(v12))) < uintptr(unsafe.Pointer(ns2.data_2_end)) {
					ns2.data_2_xxx = (*byte)(unsafe.Pointer(uintptr(v12)))
				}
			}
		}
		v13 = ns2.data_2_base
		v14 = int32(uint32(uintptr(unsafe.Pointer(ns2.data_2_xxx))) - uint32(uintptr(unsafe.Pointer(v13))))
		if v14 > 2 {
			v7 = nox_xxx_sendto_551F90(ns.sock, v13, v14, &ns2.addr)
			if v7 == -1 {
				return -1
			}
			sub_553F40(v14, 1)
			nox_xxx_netCountData_554030(v14, j)
			ns2.data_2_xxx = ns2.data_2_yyy
		}
		if compatReleaseMutex(ns2.mutex_yyy) == 0 {
			compatReleaseMutex(ns2.mutex_yyy)
		}
	}
	return v7
}
func nox_xxx_servNetInitialPackets_552A80(id uint32, flags int8) int32 {
	var (
		to  nox_net_sockaddr_in
		buf [256]byte
	)
	if id >= NOX_NET_STRUCT_MAX {
		return -3
	}
	var ns *nox_net_struct_t = nox_net_struct_arr[id]
	if ns == nil {
		return -3
	}
	var ns2 *nox_net_struct_t = ns
	var argp uint32 = 0
	if int32(flags)&1 != 0 {
		if nox_net_recv_available(ns.sock, &argp) == -1 {
			return -1
		}
		if argp == 0 {
			return -1
		}
	} else {
		argp = 1
	}
	var v26 int8 = 1
	for {
		var n int32 = nox_xxx_netRecv_552020(ns.sock, ns.data_1_xxx, int32(uintptr(unsafe.Pointer(ns.data_1_end)))-int32(uintptr(unsafe.Pointer(ns.data_1_xxx))), (*nox_net_sockaddr)(unsafe.Pointer(&to)))
		if n == -1 {
			return -1
		}
		sub_553FC0(n, 1)
		if n < 3 {
			ns.data_1_yyy = ns.data_1_base
			ns.data_1_xxx = ns.data_1_base
			if (int32(flags)&1) == 0 || (int32(flags)&4) != 0 {
				return n
			}
			argp = 0
			if nox_net_recv_available(ns.sock, &argp) == -1 {
				return -1
			} else if argp == 0 {
				return n
			}
			continue
		}
		var v7 *uint8 = (*uint8)(unsafe.Pointer(ns.data_1_yyy))
		ns.data_1_xxx = (*byte)(unsafe.Add(unsafe.Pointer(ns.data_1_xxx), n))
		var id2 uint32 = uint32(*(*uint8)(unsafe.Add(unsafe.Pointer(v7), 0)))
		var v9 uint8 = *(*uint8)(unsafe.Add(unsafe.Pointer(v7), 1))
		var op uint8 = *(*uint8)(unsafe.Add(unsafe.Pointer(v7), 2))
		if debugNet {
			stdio.Printf("servNetInitialPackets: op=%d\n", op)
		}
		if int32(op) == 12 {
			if nox_xxx_check_flag_aaa_43AF70() == 0 {
				n = int32(nox_server_makeServerInfoPacket_554040(ns.data_1_yyy, int32(uintptr(unsafe.Pointer(ns.data_1_xxx)))-int32(uintptr(unsafe.Pointer(ns.data_1_yyy))), &buf[0]))
				if n > 0 {
					n = nox_net_sendto(ns.sock, unsafe.Pointer(&buf[0]), uint32(n), &to)
					sub_553F40(n, 1)
				}
			}
			ns.data_1_yyy = ns.data_1_base
			ns.data_1_xxx = ns.data_1_base
			if (int32(flags)&1) == 0 || (int32(flags)&4) != 0 {
				return n
			}
			argp = 0
			if nox_net_recv_available(ns.sock, &argp) == -1 {
				return -1
			} else if argp == 0 {
				return n
			}
			continue
		}
		if int32(op) >= 14 && int32(op) <= 20 {
			v26 = 1
		} else {
			if id2 == math.MaxUint8 {
				if int32(v26) != 1 {
					goto LABEL_48
				}
			} else {
				v26 = 0
				if !sub_551E00(uint8(id2&math.MaxInt8), int32(uintptr(unsafe.Pointer(&to)))) {
					goto LABEL_48
				}
				v26 = 1
			}
			if ns.id == -1 {
				ns2 = nox_net_struct_arr[id2&math.MaxInt8]
			}
			if (id2 & NOX_NET_STRUCT_MAX) == 0 {
				if ns2 == nil {
					goto LABEL_48
				}
				if int32(v9) != int32(ns2.field_28_1) {
					sub_5551F0(id2, int8(v9), 1)
					sub_555360(id2, v9, 1)
					ns2.field_28_1 = int8(v9)
					var v20 bool = false
					if nox_xxx_netRead2Xxx_551EB0(id, id2, v9, int32(uintptr(unsafe.Pointer(ns.data_1_yyy))), int32(uintptr(unsafe.Pointer(ns.data_1_xxx)))-int32(uintptr(unsafe.Pointer(ns.data_1_yyy)))) == 1 {
						v20 = false
					} else {
						v20 = true
					}
					buf[0] = 38
					buf[1] = byte(ns2.field_28_1)
					ns.func_yyy(id2, &buf[0], 2, ns2.data_3)
					if !v20 {
						goto LABEL_48
					}
				}
			} else if id2 == math.MaxUint8 {
				if int32(op) == 0 {
					libc.MemCpy(unsafe.Pointer(&buf[0]), unsafe.Pointer(&to), int(unsafe.Sizeof(nox_net_sockaddr_in{})))
					n = nox_xxx_netBigSwitch_553210(id, (*uint8)(unsafe.Pointer(ns.data_1_yyy)), int32(uintptr(unsafe.Pointer(ns.data_1_xxx)))-int32(uintptr(unsafe.Pointer(ns.data_1_yyy))), int32(uintptr(unsafe.Pointer(&buf[0]))))
					if n > 0 {
						n = nox_xxx_sendto_551F90(ns.sock, &buf[0], n, &to)
						sub_553F40(n, 1)
					}
					goto LABEL_48
				}
			} else {
				*(*byte)(unsafe.Add(unsafe.Pointer(ns.data_1_yyy), 0)) &= math.MaxInt8
				id2 = uint32(*(*byte)(unsafe.Add(unsafe.Pointer(ns.data_1_yyy), 0)))
				if ns2 == nil {
					goto LABEL_48
				}
				if *(*byte)(unsafe.Add(unsafe.Pointer(ns2.data_2_base), 1)) != byte(v9) {
					goto LABEL_48
				}
				*(*byte)(unsafe.Add(unsafe.Pointer(ns2.data_2_base), 1))++
				if nox_xxx_netRead2Xxx_551EB0(id, id2, v9, int32(uintptr(unsafe.Pointer(ns.data_1_yyy))), int32(uintptr(unsafe.Pointer(ns.data_1_xxx)))-int32(uintptr(unsafe.Pointer(ns.data_1_yyy)))) == 1 {
					goto LABEL_48
				}
			}
		}
		if int32(op) < 32 {
			libc.MemCpy(unsafe.Pointer(&buf[0]), unsafe.Pointer(&to), int(unsafe.Sizeof(nox_net_sockaddr_in{})))
			n = nox_xxx_netBigSwitch_553210(id, (*uint8)(unsafe.Pointer(ns.data_1_yyy)), int32(uintptr(unsafe.Pointer(ns.data_1_xxx)))-int32(uintptr(unsafe.Pointer(ns.data_1_yyy))), int32(uintptr(unsafe.Pointer(&buf[0]))))
			if n > 0 {
				n = nox_xxx_sendto_551F90(ns.sock, &buf[0], n, &to)
				sub_553F40(n, 1)
			}
		} else {
			if ns2 != nil && (int32(flags)&2) == 0 {
				ns.func_yyy(id2, (*byte)(unsafe.Add(unsafe.Pointer(ns.data_1_yyy), 2)), n-2, ns2.data_3)
			}
		}
	LABEL_48:
		ns.data_1_yyy = ns.data_1_base
		ns.data_1_xxx = ns.data_1_base
		if (int32(flags)&1) == 0 || (int32(flags)&4) != 0 {
			return n
		}
		argp = 0
		if nox_net_recv_available(ns.sock, &argp) == -1 {
			return -1
		} else if argp == 0 {
			return n
		}
	}
}
func sub_552E70(a1 uint32) int32 {
	var (
		v3 uint32
		v4 uint32
		v5 uint32
		v9 [5]byte
	)
	if a1 >= NOX_NET_STRUCT_MAX {
		return -3
	}
	var ns *nox_net_struct_t = nox_net_struct_arr[a1]
	if ns == nil {
		return -3
	}
	if ns.id == -1 {
		v3 = 0
		v4 = NOX_NET_STRUCT_MAX
		v5 = a1
	} else {
		v5 = uint32(ns.id)
		v3 = a1
		v4 = a1 + 1
	}
	v9[0] = 6
	for ; v3 < v4; v3++ {
		var ns2 *nox_net_struct_t = nox_net_struct_arr[v3]
		if ns2 != nil && uint32(ns2.id) == v5 {
			var v8 int32 = int32(dword_5d4594_2495920)
			ns2.field_22 = dword_5d4594_2495920
			ns2.field_23 = dword_5d4594_2495920
			*(*uint32)(unsafe.Pointer(&v9[1])) = uint32(v8)
			nox_xxx_netSendSock_552640(v3, &v9[0], 5, int8(NOX_NET_SEND_FLAG2))
		}
	}
	return 0
}
func sub_552F20(a1 uint32) int32 {
	var (
		v3 uint32
		v4 uint32
		v5 uint32
		v9 [256]byte
	)
	if a1 >= NOX_NET_STRUCT_MAX {
		return -3
	}
	var ns *nox_net_struct_t = nox_net_struct_arr[a1]
	if ns == nil {
		return -3
	}
	if ns.id == -1 {
		v3 = 0
		v4 = NOX_NET_STRUCT_MAX
		v5 = a1
	} else {
		v5 = uint32(ns.id)
		v3 = a1
		v4 = a1 + 1
	}
	v9[0] = 5
	for ; v3 < v4; v3++ {
		var ns2 *nox_net_struct_t = nox_net_struct_arr[v3]
		if ns2 != nil && uint32(ns2.id) == v5 {
			var v8 int32 = int32(ns2.field_25 + 1)
			ns2.field_25 = uint32(v8)
			ns2.field_26 = dword_5d4594_2495920
			*(*uint32)(unsafe.Pointer(&v9[1])) = uint32(v8)
			nox_xxx_netSendSock_552640(v3, &v9[0], 256, int8(NOX_NET_SEND_FLAG2))
		}
	}
	return 0
}
func sub_552FD0(a1 int32) uint32 {
	var argp uint32 = 0
	if nox_net_recv_available(nox_net_struct_arr[a1].sock, &argp) != 0 {
		return math.MaxUint32
	}
	return argp
}
func nox_xxx_netBigSwitch_553210(id uint32, packet *uint8, packetSz int32, a4 int32) int32 {
	var (
		out       int32  = a4
		pid       int32  = int32(int8(*(*uint8)(unsafe.Add(unsafe.Pointer(packet), 0))))
		p1        int8   = int8(*(*uint8)(unsafe.Add(unsafe.Pointer(packet), 1)))
		packetEnd *uint8 = (*uint8)(unsafe.Add(unsafe.Pointer(packet), packetSz))
		v74       [8]uint8
	)
	*(*uint32)(unsafe.Pointer(&v74[0])) = *(*uint32)(unsafe.Pointer(uintptr(a4 + 0)))
	*(*uint32)(unsafe.Pointer(&v74[4])) = *(*uint32)(unsafe.Pointer(uintptr(a4 + 4)))
	var v75 int32 = int32(*(*uint32)(unsafe.Pointer(uintptr(a4 + 8))))
	var v76 int32 = int32(*(*uint32)(unsafe.Pointer(uintptr(a4 + 12))))
	if packetSz <= 2 {
		return 0
	}
	var packetCur *uint8 = (*uint8)(unsafe.Add(unsafe.Pointer(packet), 2))
	if id > NOX_NET_STRUCT_MAX {
		stdio.Printf("nox_net_struct_arr overflow (1): %d\n", int32(id))
		panic("abort")
	}
	var ns1 *nox_net_struct_t = nox_net_struct_arr[id]
	var pidb uint32 = uint32(pid)
	for 2 != 0 {
		var op int32 = int32(*(*uint8)(unsafe.Add(unsafe.Pointer(packetCur), 0)))
		packetCur = (*uint8)(unsafe.Add(unsafe.Pointer(packetCur), 1))
		if debugNet {
			stdio.Printf("nox_xxx_netBigSwitch_553210: op=%d [%d]\n", op, packetSz)
		}
		switch op {
		case 0:
			if noxflags.HasGame(noxflags.GameHost) && noxflags.HasGame(noxflags.GameFlag4) {
				return 0
			}
			*(*uint8)(unsafe.Pointer(uintptr(out + 0))) = 0
			*(*uint8)(unsafe.Pointer(uintptr(out + 1))) = uint8(p1)
			if ns1.field_21 >= uint32(nox_xxx_servGetPlrLimit_409FA0())+uint32(*memmap.PtrUint8(6112660, 2500076))-1 {
				*(*uint8)(unsafe.Pointer(uintptr(out + 2))) = 2
				return 3
			}
			if pid != -1 {
				*(*uint8)(unsafe.Pointer(uintptr(out + 2))) = 2
				return 3
			}
			for i := int32(0); i < NOX_NET_STRUCT_MAX; i++ {
				var ns9 *nox_net_struct_t = nox_net_struct_arr[i]
				if ns9 == nil {
					pid = i
					break
				}
				if int32(*(*uint16)(unsafe.Pointer(&v74[2]))) == int32(ns9.addr.sin_port) && *(*uint32)(unsafe.Pointer(&v74[4])) == uint32(ns9.addr.sin_addr) {
					stdio.Printf("%d %d\n", *(*uint16)(unsafe.Pointer(&v74[2])), *(*uint32)(unsafe.Pointer(&v74[4])))
					*(*uint8)(unsafe.Pointer(uintptr(out + 2))) = 4
					return 3
				}
			}
			if pid == -1 {
				*(*uint8)(unsafe.Pointer(uintptr(out + 2))) = 2
				return 3
			}
			var narg nox_net_struct_arg_t
			narg = nox_net_struct_arg_t{}
			narg.data_3_size = 4
			narg.data_size = int32(uintptr(unsafe.Pointer(ns1.data_2_end))) - int32(uintptr(unsafe.Pointer(ns1.data_2_base)))
			var ns10 *nox_net_struct_t = nox_xxx_makeNewNetStruct_553000(&narg)
			nox_net_struct_arr[pid] = ns10
			if ns10 == nil {
				*(*uint8)(unsafe.Pointer(uintptr(out + 2))) = 2
				return 3
			}
			ns1.field_21++
			*(*byte)(unsafe.Add(unsafe.Pointer(ns10.data_2_base), 0)) = byte(id)
			var v62 int32 = int32(uintptr(unsafe.Pointer(ns10.data_2_base)))
			var v63 int8 = int8(*(*uint8)(unsafe.Pointer(uintptr(v62 + 1))))
			if int32(v63) == int32(p1) {
				*(*uint8)(unsafe.Pointer(uintptr(v62 + 1))) = uint8(int8(int32(v63) + 1))
			}
			ns10.id = int32(id)
			ns10.sock = ns1.sock
			ns10.func_xxx = ns1.func_xxx
			ns10.func_yyy = ns1.func_yyy
			libc.MemSet(memmap.PtrOff(6112660, id*32+0x2647F4), 0, 32)
			*memmap.PtrUint32(6112660, id*32+0x264810) = 1
			var key int8 = int8(noxRndCounter1.IntClamp(1, math.MaxUint8))
			if nox_net_no_xor {
				key = 0
			}
			ns10.xor_key = 0
			var v66 int32 = int32(uintptr(unsafe.Pointer(&ns10.addr)))
			*(*uint64)(unsafe.Pointer(uintptr(v66))) = *(*uint64)(unsafe.Pointer(&v74[0]))
			*(*uint32)(unsafe.Pointer(uintptr(v66 + 8))) = uint32(v75)
			*(*uint32)(unsafe.Pointer(uintptr(v66 + 12))) = uint32(v76)
			*(*uint8)(unsafe.Pointer(uintptr(out + 0))) = 31
			*(*uint8)(unsafe.Pointer(uintptr(out + 1))) = uint8(p1)
			*(*uint8)(unsafe.Pointer(uintptr(out + 2))) = 1
			*(*uint32)(unsafe.Pointer(uintptr(out + 3))) = uint32(pid)
			*(*uint8)(unsafe.Pointer(uintptr(out + 7))) = uint8(key)
			var v67 int8 = int8(nox_xxx_netSendSock_552640(uint32(pid), (*byte)(unsafe.Pointer(uintptr(out))), 8, int8(NOX_NET_SEND_NO_LOCK|NOX_NET_SEND_FLAG2)))
			ns10.xor_key = uint8(key)
			ns10.field_38 = 1
			ns10.data_39[0] = uint8(v67)
			ns10.field_40 = nox_frame_xxx_2598000
			return 0
		case 1:
			var (
				v11 int32  = int32(*(*uint32)(unsafe.Pointer(packetCur)))
				v12 *uint8 = (*uint8)(unsafe.Pointer(ns1.data_2_base))
				v13 *uint8 = (*uint8)(unsafe.Add(unsafe.Pointer(packetCur), 4))
			)
			ns1.id = v11
			*v12 = uint8(int8(v11))
			packetCur = (*uint8)(unsafe.Add(unsafe.Pointer(v13), 1))
			ns1.xor_key = *v13
			dword_5d4594_3844304 = 1
			if uintptr(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(packetCur)))))) >= uintptr(unsafe.Pointer(packetEnd)) {
				return 0
			}
		case 2:
			ns1.id = -18
			dword_5d4594_3844304 = 1
			if uintptr(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(packetCur)))))) >= uintptr(unsafe.Pointer(packetEnd)) {
				return 0
			}
		case 3:
			ns1.id = -12
			dword_5d4594_3844304 = 1
			if uintptr(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(packetCur)))))) >= uintptr(unsafe.Pointer(packetEnd)) {
				return 0
			}
		case 4:
			ns1.id = -13
			dword_5d4594_3844304 = 1
			if uintptr(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(packetCur)))))) >= uintptr(unsafe.Pointer(packetEnd)) {
				return 0
			}
		case 5:
			*(*uint8)(unsafe.Pointer(uintptr(out + 0))) = ns1.xor_key
			*(*uint8)(unsafe.Pointer(uintptr(out + 2))) = 7
			*(*uint32)(unsafe.Pointer(uintptr(out + 3))) = *(*uint32)(unsafe.Pointer(packetCur))
			return 7
		case 6:
			if pidb > NOX_NET_STRUCT_MAX {
				stdio.Printf("nox_net_struct_arr overflow (2): %d\n", int32(pidb))
				panic("abort")
			}
			var ns2 *nox_net_struct_t = nox_net_struct_arr[pidb]
			var ns3 *nox_net_struct_t = nox_net_struct_arr[pid]
			var a2b3 uint8 = 37
			ns1.func_yyy(uint32(pid), (*byte)(unsafe.Pointer(&a2b3)), 1, ns3.data_3)
			var v18 int32 = 0
			if ns1.id == -1 {
				*(*uint8)(unsafe.Pointer(uintptr(out + 0))) = uint8(int8(ns2.id))
				v18 = int32(uintptr(unsafe.Pointer(ns2.data_2_base)))
			} else {
				*(*uint8)(unsafe.Pointer(uintptr(out + 0))) = uint8(int8(ns1.id))
				v18 = int32(uintptr(unsafe.Pointer(ns1.data_2_base)))
			}
			*(*uint8)(unsafe.Pointer(uintptr(out + 1))) = *(*uint8)(unsafe.Pointer(uintptr(v18 + 1)))
			*memmap.PtrUint32(6112660, uint32(pid*32)+0x2647F8) = *(*uint32)(unsafe.Pointer(packetCur))
			*(*uint8)(unsafe.Pointer(uintptr(out + 2))) = 8
			*(*uint32)(unsafe.Pointer(uintptr(out + 3))) = *(*uint32)(unsafe.Pointer(packetCur))
			return 7
		case 7:
			if pidb > NOX_NET_STRUCT_MAX {
				stdio.Printf("nox_net_struct_arr overflow (2): %d\n", int32(pidb))
				panic("abort")
			}
			var ns4 *nox_net_struct_t = nox_net_struct_arr[pidb]
			if ns4.field_25 == 0 {
				return 0
			}
			var v31 int32 = int32(dword_5d4594_2495920 - uint32(int32(ns4.field_26)) - uint32(int32(ns4.field_24)))
			var v32 int32 = -1
			if v31 >= 1 {
				v32 = int32(256000 / uint32(v31))
			}
			*(*uint8)(unsafe.Pointer(uintptr(out + 0))) = 35
			*(*uint32)(unsafe.Pointer(uintptr(out + 1))) = uint32(v32)
			if ns1.id == -1 {
				ns1.func_yyy(uint32(pid), (*byte)(unsafe.Pointer(uintptr(out))), 5, ns4.data_3)
			} else {
				ns1.func_yyy(uint32(pid), (*byte)(unsafe.Pointer(uintptr(out))), 5, ns1.data_3)
			}
			return 0
		case 8:
			if pidb > NOX_NET_STRUCT_MAX {
				stdio.Printf("nox_net_struct_arr overflow (2): %d\n", int32(pidb))
				panic("abort")
			}
			var ns5 *nox_net_struct_t = nox_net_struct_arr[pidb]
			if *(*uint32)(unsafe.Pointer(packetCur)) != ns5.field_22 {
				return 0
			}
			ns5.field_24 = dword_5d4594_2495920 - ns5.field_23
			*(*uint8)(unsafe.Pointer(uintptr(out + 0))) = 36
			*(*uint32)(unsafe.Pointer(uintptr(out + 1))) = ns5.field_24
			var v19 int32 = 0
			if ns1.id == -1 {
				v19 = int32(uintptr(ns5.data_3))
			} else {
				v19 = int32(uintptr(ns1.data_3))
			}
			ns1.func_yyy(uint32(pid), (*byte)(unsafe.Pointer(uintptr(out))), 5, unsafe.Pointer(uintptr(v19)))
			*(*uint8)(unsafe.Pointer(uintptr(out + 0))) = uint8(*(*byte)(unsafe.Add(unsafe.Pointer(ns1.data_2_base), 0)))
			*(*uint8)(unsafe.Pointer(uintptr(out + 1))) = uint8(*(*byte)(unsafe.Add(unsafe.Pointer(ns5.data_2_base), 1)))
			*(*uint8)(unsafe.Pointer(uintptr(out + 2))) = 9
			*(*uint32)(unsafe.Pointer(uintptr(out + 3))) = dword_5d4594_2495920
			return 7
		case 9:
			var (
				v21 int32 = pid * 32
				v22 int32 = int32(*(*uint32)(unsafe.Pointer(packetCur)) - *memmap.PtrUint32(6112660, uint32(pid*32)+0x2647F8))
			)
			if v22 <= 0 || v22 >= 1000 {
				return 0
			}
			var v23 int32 = int32(*memmap.PtrUint32(6112660, uint32(pid*32)+0x2647F4))
			var v24 int32 = v23 + pid*8
			var v25 int32 = 5
			*memmap.PtrUint32(6112660, uint32(v24*4)+0x2647FC) = uint32(v22)
			var v26 int32 = (v23 + 1) % 5
			var v27 int32 = v26
			if v26 == 0 {
				var v28 *uint8 = (*uint8)(memmap.PtrOff(6112660, uint32(v21)+0x2647FC))
				for {
					{
						var v29 int32 = int32(*(*uint32)(unsafe.Pointer(v28)))
						v28 = (*uint8)(unsafe.Add(unsafe.Pointer(v28), 4))
						v26 += v29
						v25--
					}
					if v25 == 0 {
						break
					}
				}
				*memmap.PtrUint32(6112660, uint32(v21)+0x264810) = uint32(v26 / 5)
			}
			*memmap.PtrUint32(6112660, uint32(v21)+0x2647F4) = uint32(v27)
			return 0
		case 10:
			if pid == math.MaxUint8 {
				return 0
			}
			var ns6 *nox_net_struct_t = nox_net_struct_arr[pid]
			if ns6.field_38 == 1 {
				return 0
			}
			var a2b2 uint8 = 34
			ns1.func_yyy(uint32(pid), (*byte)(unsafe.Pointer(&a2b2)), 1, ns6.data_3)
			libc.MemSet(memmap.PtrOff(6112660, id*32+0x2647F4), 0, 32)
			var v69 *int32 = nox_xxx_findPlayerID_5541D0(pid)
			if v69 != nil {
				sub_425920((**uint32)(unsafe.Pointer(v69)))
				alloc.Free(unsafe.Pointer(v69))
				*memmap.PtrUint8(6112660, 2500076)--
			}
			nox_xxx_netStructReadPackets_5545B0(uint32(pid))
			return 0
		case 11:
			var (
				ns7 *nox_net_struct_t = nox_net_struct_arr[pid]
				a2b uint8             = 33
			)
			ns1.func_yyy(uint32(pid), (*byte)(unsafe.Pointer(&a2b)), 1, ns7.data_3)
			sub_554A50(id)
			return 0
		case 14:
			var (
				v43 int32 = 0
				v78 *byte = (*byte)(sub_416640())
			)
			nox_xxx_cliGamedataGet_416590(0)
			var a4a bool = false
			*(*uint8)(unsafe.Pointer(uintptr(out + 0))) = 0
			*(*uint8)(unsafe.Pointer(uintptr(out + 1))) = uint8(p1)
			if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(packet))), unsafe.Sizeof(uint32(0))*20))) != NOX_CLIENT_VERS_CODE {
				*(*uint8)(unsafe.Pointer(uintptr(out + 2))) = 19
				*(*uint8)(unsafe.Pointer(uintptr(out + 3))) = 13
				return 4
			}
			if ns1.field_21 >= uint32(nox_xxx_servGetPlrLimit_409FA0()-1) {
				a4a = true
			}
			if sub_40A740() != 0 {
				var v46 int32 = nox_xxx_countObserverPlayers_425BF0()
				if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(packet))), unsafe.Sizeof(uint32(0))*21))) == 0 {
					if v46 >= int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v78), 53)))) {
						*(*uint8)(unsafe.Pointer(uintptr(out + 2))) = 19
						*(*uint8)(unsafe.Pointer(uintptr(out + 3))) = 11
						return 4
					}
				} else if nox_server_teamByXxx_418AE0(int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(packet))), unsafe.Sizeof(uint32(0))*21))))) != nil {
					if v46 > 0 {
						v43 = 1
					}
				} else {
					if int32(uint8(sub_417DE0())) >= int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v78), 52)))) {
						if v46 >= int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v78), 53)))) {
							*(*uint8)(unsafe.Pointer(uintptr(out + 2))) = 19
							*(*uint8)(unsafe.Pointer(uintptr(out + 3))) = 11
							return 4
						}
					} else if v46 > 0 {
						v43 = 1
					}
				}
			}
			if a4a {
				if v43 == 0 || *(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v78), 54)))) == 0 {
					*(*uint8)(unsafe.Pointer(uintptr(out + 2))) = 19
					*(*uint8)(unsafe.Pointer(uintptr(out + 3))) = 11
					return 4
				}
				for i := (*byte)(nox_xxx_firstReplaceablePlayer_425C40()); i != nil; i = nox_xxx_nextReplaceablePlayer_425C70(int32(uintptr(unsafe.Pointer(i)))) {
					if nox_xxx_findPlayerID_5541D0(int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(i), 2064))))+1) == nil {
						nox_xxx_playerCallDisconnect_4DEAB0(int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(i), 2064)))), 4)
						var v50 *uint32 = (*uint32)(alloc.Calloc(1, 16))
						*(*uint32)(unsafe.Add(unsafe.Pointer(v50), unsafe.Sizeof(uint32(0))*3)) = uint32(int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(i), 2064)))) + 1)
						nox_common_list_append_4258E0((*nox_list_item_t)(unsafe.Pointer(uintptr(int32(uintptr(memmap.PtrOff(6112660, 2495908)))))), (*nox_list_item_t)(unsafe.Pointer(v50)))
						*memmap.PtrUint8(6112660, 2500076)++
						*(*uint8)(unsafe.Pointer(uintptr(out + 2))) = 21
						return 3
					}
				}
			}
			if *(*byte)(unsafe.Add(unsafe.Pointer(v78), 100))&16 != 0 {
				var v48 *int32 = sub_4168E0()
				if v48 == nil {
					*(*uint8)(unsafe.Pointer(uintptr(out + 2))) = 19
					*(*uint8)(unsafe.Pointer(uintptr(out + 3))) = 4
					return 4
				}
				for _nox_wcsicmp((*libc.WChar)(unsafe.Add(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(v48))), unsafe.Sizeof(libc.WChar(0))*6)), (*libc.WChar)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(packet), 4))))) != 0 {
					v48 = sub_4168F0(v48)
					if v48 == nil {
						*(*uint8)(unsafe.Pointer(uintptr(out + 2))) = 19
						*(*uint8)(unsafe.Pointer(uintptr(out + 3))) = 4
						return 4
					}
				}
			} else {
				for j := (*int32)(sub_416900()); j != nil; j = sub_416910(j) {
					if libc.StrCmp((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(j))), 72)), CString("0")) == 0 {
						if _nox_wcsicmp((*libc.WChar)(unsafe.Add(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(j))), unsafe.Sizeof(libc.WChar(0))*6)), (*libc.WChar)(unsafe.Add(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(packet))), unsafe.Sizeof(libc.WChar(0))*2))) == 0 {
							*(*uint8)(unsafe.Pointer(uintptr(out + 2))) = 19
							*(*uint8)(unsafe.Pointer(uintptr(out + 3))) = 5
							return 4
						}
					} else if libc.StrCaseCmp((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(j))), 72)), (*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(packet))), 56))) == 0 {
						*(*uint8)(unsafe.Pointer(uintptr(out + 2))) = 19
						*(*uint8)(unsafe.Pointer(uintptr(out + 3))) = 5
						return 4
					}
				}
			}
			var v35 *byte = v78
			var v52 int8 = int8(*(*byte)(unsafe.Add(unsafe.Pointer(v78), 100)))
			if int32(v52) != 0 && int32(uint8(int8(1<<int32(*(*uint8)(unsafe.Add(unsafe.Pointer(packet), 54))))))&int32(uint8(v52)) != 0 {
				*(*uint8)(unsafe.Pointer(uintptr(out + 2))) = 19
				*(*uint8)(unsafe.Pointer(uintptr(out + 3))) = 7
				return 4
			}
			if int32(v52)&32 != 0 {
				*(*uint8)(unsafe.Pointer(uintptr(out + 2))) = 15
				return 3
			}
			if int32(*(*int16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v78), 105))))) == -1 && int32(*(*int16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v35), 107))))) == -1 {
				*(*uint8)(unsafe.Pointer(uintptr(out + 2))) = 20
				return 3
			}
			var id53 int32 = sub_553D10()
			if id53 < 0 {
				*(*uint8)(unsafe.Pointer(uintptr(out + 2))) = 20
				return 3
			}
			var nx *nox_net_struct2_t = &nox_net_struct2_arr[id53]
			nx.field_0 = 1
			nx.field_1_1 = 0
			nx.field_1_0 = 0
			*(*uint64)(unsafe.Pointer(&nx.addr)) = *(*uint64)(unsafe.Pointer(&v74[0]))
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&nx.addr))), unsafe.Sizeof(uint32(0))*1))) = uint32(v75)
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&nx.addr))), unsafe.Sizeof(uint32(0))*2))) = uint32(v76)
			return nox_xxx_makePacketTime_552340(id53, (*uint8)(unsafe.Pointer(uintptr(out))))
		case 17:
			var (
				v33 *byte = (*byte)(sub_416640())
				v35 *byte = v33
			)
			*(*uint8)(unsafe.Pointer(uintptr(out + 0))) = 0
			*(*uint8)(unsafe.Pointer(uintptr(out + 1))) = uint8(p1)
			if nox_wcscmp((*libc.WChar)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(packet), 4)))), (*libc.WChar)(unsafe.Add(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(v33))), unsafe.Sizeof(libc.WChar(0))*39))) != 0 {
				*(*uint8)(unsafe.Pointer(uintptr(out + 2))) = 19
				*(*uint8)(unsafe.Pointer(uintptr(out + 3))) = 6
				return 4
			}
			if int32(*(*int16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v35), 105))))) == -1 && int32(*(*int16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v35), 107))))) == -1 {
				*(*uint8)(unsafe.Pointer(uintptr(out + 2))) = 20
				return 3
			}
			var id53 int32 = sub_553D10()
			if id53 < 0 {
				*(*uint8)(unsafe.Pointer(uintptr(out + 2))) = 20
				return 3
			}
			var nx1 *nox_net_struct2_t = &nox_net_struct2_arr[id53]
			nx1.field_0 = 1
			nx1.field_1_1 = 0
			nx1.field_1_0 = 0
			*(*uint64)(unsafe.Pointer(&nx1.addr)) = *(*uint64)(unsafe.Pointer(&v74[0]))
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&nx1.addr))), unsafe.Sizeof(uint32(0))*1))) = uint32(v75)
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&nx1.addr))), unsafe.Sizeof(uint32(0))*2))) = uint32(v76)
			return nox_xxx_makePacketTime_552340(id53, (*uint8)(unsafe.Pointer(uintptr(out))))
		case 18:
			var (
				v39  int32 = int32(platformTicks() - *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(packet))), unsafe.Sizeof(uint32(0))*1))))
				id40 int32 = sub_553D30(int32(uintptr(unsafe.Pointer(&v74[0]))))
			)
			if id40 < 0 {
				return 0
			}
			var nx1 *nox_net_struct2_t = &nox_net_struct2_arr[id40]
			if int32(*((*uint8)(unsafe.Add(unsafe.Pointer(packet), 3)))) != int32(nx1.field_1_1) {
				return 0
			}
			nx1.field_6[nx1.field_1_1] = uint32(v39)
			nx1.field_1_1++
			if int32(nx1.field_1_1) >= 10 {
				return 0
			}
			return nox_xxx_makePacketTime_552340(id40, (*uint8)(unsafe.Pointer(uintptr(out))))
		case 31:
			if pidb > NOX_NET_STRUCT_MAX {
				stdio.Printf("nox_net_struct_arr overflow (2): %d\n", int32(pidb))
				panic("abort")
			}
			var ns8 *nox_net_struct_t = nox_net_struct_arr[pidb]
			var v14 uint8 = *(*uint8)(unsafe.Add(unsafe.Pointer(packetCur), 0))
			var v15 uint8 = uint8(ns8.field_28_1)
			packetCur = (*uint8)(unsafe.Add(unsafe.Pointer(packetCur), 1))
			var a4b uint8 = v14
			stdio.Printf("foo 0x%x 0x%x\n", v14, v15)
			if int32(v14) != int32(v15) {
				sub_5551F0(uint32(pid), int8(a4b), 1)
				sub_555360(uint32(pid), a4b, 1)
				ns8.field_28_1 = int8(a4b)
				*(*uint8)(unsafe.Pointer(uintptr(out + 0))) = 38
				*(*uint8)(unsafe.Pointer(uintptr(out + 1))) = uint8(ns8.field_28_1)
				ns1.func_yyy(uint32(pid), (*byte)(unsafe.Pointer(uintptr(out))), 2, ns8.data_3)
			}
			if uintptr(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(packetCur)))))) >= uintptr(unsafe.Pointer(packetEnd)) {
				return 0
			}
		default:
			return 0
		}
	}
}
func sub_553D10() int32 {
	for i := int32(0); i < NOX_NET_STRUCT_MAX; i++ {
		if nox_net_struct2_arr[i].field_0 == 0 {
			return i
		}
	}
	return -1
}
func sub_553D30(a1 int32) int32 {
	for i := int32(0); i < NOX_NET_STRUCT_MAX; i++ {
		var nx *nox_net_struct2_t = &nox_net_struct2_arr[i]
		if nx.addr.sin_addr == nox_net_in_addr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 4)))) && int32(nx.addr.sin_port) == int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 2)))) {
			return i
		}
	}
	return -1
}
func sub_553D80() int32 {
	return int32(*memmap.PtrUint32(6112660, 2495944))
}
func sub_553D90() int32 {
	return int32(*memmap.PtrUint32(6112660, 2495952))
}
func sub_553DA0() int32 {
	return int32(*memmap.PtrUint32(6112660, 2495948))
}
func sub_553DB0() int32 {
	return int32(*memmap.PtrUint32(6112660, 2495956))
}
func sub_553DC0(a1 int32) int64 {
	var (
		v1 int32
		v2 uint32
		v3 int32
	)
	v1 = int32(dword_5d4594_2496472)
	v2 = 0
	v3 = math.MaxInt8
	for {
		v2 += *memmap.PtrUint32(6112660, uint32(v1*4)+2495960)
		v1 = (v1 + 1) % 128
		v3--
		if v3 == 0 {
			break
		}
	}
	return int64(float64(v2) / (128.0 / float64(a1)))
}
func sub_553E10(a1 int32) uint32 {
	var (
		v1 int32
		v2 uint32
		v3 int32
	)
	v1 = int32(*memmap.PtrUint32(6112660, 2497504))
	v2 = 0
	v3 = math.MaxInt8
	for {
		v2 += *memmap.PtrUint32(6112660, uint32(v1*4)+0x2619E0)
		v1 = (v1 + 1) % 128
		v3--
		if v3 == 0 {
			break
		}
	}
	return v2 / uint32(128/a1)
}
func sub_553E50(a1 int32) uint32 {
	var (
		v1 int32
		v2 uint32
		v3 int32
	)
	v1 = int32(dword_5d4594_2496988)
	v2 = 0
	v3 = math.MaxInt8
	for {
		v2 += *memmap.PtrUint32(6112660, uint32(v1*4)+0x2617DC)
		v1 = (v1 + 1) % 128
		v3--
		if v3 == 0 {
			break
		}
	}
	return v2 / uint32(128/a1)
}
func sub_553E90(a1 int32) uint32 {
	var (
		v1 int32
		v2 uint32
		v3 int32
	)
	v1 = int32(*memmap.PtrUint32(6112660, 2498020))
	v2 = 0
	v3 = math.MaxInt8
	for {
		v2 += *memmap.PtrUint32(6112660, uint32(v1*4)+0x261BE4)
		v1 = (v1 + 1) % 128
		v3--
		if v3 == 0 {
			break
		}
	}
	return v2 / uint32(128/a1)
}
func sub_553F40(a1 int32, a2 int32) {
	*memmap.PtrUint32(6112660, 2495952) += uint32(a1)
	*memmap.PtrUint32(6112660, 2495956) += uint32(a2)
	*memmap.PtrUint32(6112660, *memmap.PtrUint32(6112660, 2497504)*4+0x2619E0) = uint32(a1)
	*memmap.PtrUint32(6112660, *memmap.PtrUint32(6112660, 2498020)*4+0x261BE4) = uint32(a2)
	*memmap.PtrUint32(6112660, 2497504) = (dword_5d4594_2496472 + 1) % 128
	*memmap.PtrUint32(6112660, 2498020) = (dword_5d4594_2496988 + 1) % 128
}
func sub_553FC0(a1 int32, a2 int32) {
	var (
		v2 int32
		v3 int32
	)
	*memmap.PtrUint32(6112660, 2495944) += uint32(a1)
	*memmap.PtrUint32(6112660, 2495948) += uint32(a2)
	v2 = int32(dword_5d4594_2496472)
	*memmap.PtrUint32(6112660, dword_5d4594_2496472*4+2495960) = uint32(a1)
	v3 = int32(dword_5d4594_2496988)
	*memmap.PtrUint32(6112660, dword_5d4594_2496988*4+0x2617DC) = uint32(a2)
	dword_5d4594_2496472 = uint32((v2 + 1) % 128)
	dword_5d4594_2496988 = uint32((v3 + 1) % 128)
}
func nox_xxx_netCountData_554030(a1 int32, a2 int32) {
	*memmap.PtrUint32(6112660, uint32(a2*4)+0x261DE8) += uint32(a1)
}
func nox_server_makeServerInfoPacket_554040(inBuf *byte, inSz int32, out *byte) uint32 {
	var (
		buf  [72]byte
		v3   *byte = (*byte)(sub_416640())
		game *byte = nox_xxx_cliGamedataGet_416590(0)
	)
	if sub_43AF30() == 0 || sub_4D6F30() != 0 {
		return 0
	}
	var playerLimit int8 = int8(nox_xxx_servGetPlrLimit_409FA0())
	var playerCount int8 = int8(nox_common_playerInfoCount_416F40())
	if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING)) {
		playerCount--
		playerLimit--
	}
	var srvName *byte = nox_xxx_serverOptionsGetServername_40A4C0()
	buf[0] = 0
	buf[1] = 0
	buf[2] = 13
	buf[3] = byte(playerCount)
	buf[4] = byte(playerLimit)
	buf[5] = *(*byte)(unsafe.Add(unsafe.Pointer(v3), 101)) & 15
	buf[6] = byte(int8(int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v3), 101)))) >> 4))
	*(*uint32)(unsafe.Pointer(&buf[7])) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(game))), unsafe.Sizeof(uint32(0))*11)))
	libc.StrCpy(&buf[10], nox_xxx_mapGetMapName_409B40())
	buf[19] = byte(int8(int32(*(*byte)(unsafe.Add(unsafe.Pointer(v3), 102))) | sub_43BE50_get_video_mode_id()))
	buf[20] = *(*byte)(unsafe.Add(unsafe.Pointer(v3), 100))
	buf[21] = *(*byte)(unsafe.Add(unsafe.Pointer(v3), 100)) & 16
	*(*uint32)(unsafe.Pointer(&buf[24])) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*12)))
	var gameFlags uint32 = nox_common_gameFlags_getVal_40A5B0()
	if nox_xxx_isQuest_4D6F50() != 0 {
		gameFlags = (gameFlags & 0xFFFFFF7F) | 4096
		*(*uint16)(unsafe.Pointer(&buf[68])) = uint16(int16(nox_game_getQuestStage_4E3CC0()))
	}
	*(*uint32)(unsafe.Pointer(&buf[28])) = gameFlags
	*(*uint32)(unsafe.Pointer(&buf[32])) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(game))), unsafe.Sizeof(uint32(0))*12)))
	*(*uint16)(unsafe.Pointer(&buf[36])) = *(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v3), 105))))
	*(*uint16)(unsafe.Pointer(&buf[38])) = *(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v3), 107))))
	*(*uint32)(unsafe.Pointer(&buf[40])) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*11)))
	*(*uint32)(unsafe.Pointer(&buf[44])) = *(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(inBuf), 8))))
	libc.MemCpy(unsafe.Pointer(&buf[48]), unsafe.Add(unsafe.Pointer(game), 24), 20)
	libc.MemCpy(unsafe.Add(unsafe.Pointer(out), 0), unsafe.Pointer(&buf[0]), 72)
	libc.StrCpy((*byte)(unsafe.Add(unsafe.Pointer(out), 72)), srvName)
	return uint32(libc.StrLen(srvName) + 72 + 1)
}
func nox_xxx_findPlayerID_5541D0(a1 int32) *int32 {
	var result *int32
	result = (*int32)(unsafe.Pointer(nox_common_list_getFirstSafe_425890((*nox_list_item_t)(unsafe.Pointer(memmap.PtrInt32(6112660, 2495908))))))
	if result == nil {
		return nil
	}
	for *(*int32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(int32(0))*3)) != a1 {
		result = (*int32)(unsafe.Pointer(nox_common_list_getNextSafe_4258A0((*nox_list_item_t)(unsafe.Pointer(result)))))
		if result == nil {
			return nil
		}
	}
	return result
}
func nox_xxx_net_getIP_554200(a1 uint32) int32 {
	if a1 > NOX_NET_STRUCT_MAX {
		return 0
	}
	if a1 == 0 {
		return int32(dword_5d4594_3843632)
	}
	var ns *nox_net_struct_t = nox_net_struct_arr[a1]
	if ns == nil {
		return 0
	}
	return int32(ns.addr.sin_addr)
}
func sub_554230() *byte {
	return (*byte)(memmap.PtrOff(0x973F18, 44216))
}
func sub_554290() uint32 {
	var (
		v0     uint32
		v1     int32
		v2     *byte
		v3     uint32
		result uint32
	)
	v0 = math.MaxUint32
	v1 = 0
	v2 = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0()))
	if v2 == nil {
		goto LABEL_13
	}
	for {
		if *(*byte)(unsafe.Add(unsafe.Pointer(v2), 2064)) != 31 && sub_554240(int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v2), 2064))))) > 0 {
			v3 = uint32(sub_554240(int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v2), 2064))))))
			if v3 < v0 {
				v0 = v3
			}
			v1++
		}
		v2 = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))))))
		if v2 == nil {
			break
		}
	}
	if v1 != 0 {
		result = v0
	} else {
	LABEL_13:
		result = 0
	}
	return result
}
func sub_554300() int32 {
	var (
		v0     int32
		v1     int32
		v2     *byte
		result int32
	)
	v0 = 0
	v1 = 0
	v2 = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0()))
	if v2 == nil {
		goto LABEL_11
	}
	for {
		if *(*byte)(unsafe.Add(unsafe.Pointer(v2), 2064)) != 31 && sub_554240(int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v2), 2064))))) > 0 {
			v0 += sub_554240(int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v2), 2064)))))
			v1++
		}
		v2 = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))))))
		if v2 == nil {
			break
		}
	}
	if v1 != 0 {
		result = v0 / v1
	} else {
	LABEL_11:
		result = 0
	}
	return result
}
func sub_554240(a1 int32) int32 {
	var (
		v1     *byte
		result int32
	)
	if a1 != 31 {
		return int32(*memmap.PtrUint32(6112660, uint32(a1*32)+0x264830))
	}
	v1 = (*byte)(sub_416640())
	switch *(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 66)))) {
	case 1:
		result = int32(sub_554290())
	case 2:
		result = sub_554300()
	case 3:
		result = int32(*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 70)))))
	default:
		result = 0
	}
	return result
}
func sub_5549F0(a1 uint32) int32 {
	var v2 int8
	v2 = 10
	if a1 >= NOX_NET_STRUCT_MAX {
		return -3
	}
	if nox_net_struct_arr[a1] != nil {
		nox_xxx_netSendReadPacket_5528B0(a1, 0)
		nox_xxx_netSendSock_552640(a1, (*byte)(unsafe.Pointer(&v2)), 1, 0)
		nox_xxx_netSendReadPacket_5528B0(a1, 0)
		sub_554A50(a1)
	}
	return 0
}
func sub_554A50(a1 uint32) int32 {
	if a1 >= NOX_NET_STRUCT_MAX {
		return -3
	}
	var ns *nox_net_struct_t = nox_net_struct_arr[a1]
	if ns != nil {
		nox_net_close(ns.sock)
		OnLibraryNotice_259(uint32(uintptr(unsafe.Pointer(ns))))
		nox_xxx_netStructFree_5531C0(nox_net_struct_arr[a1])
		nox_net_struct_arr[a1] = nil
		nox_net_stop()
	}
	return 0
}
func sub_555130(a1 uint32, a2 unsafe.Pointer, a3 int32) int32 {
	var v5 *uint32
	if a3 > *memmap.PtrInt32(6112660, 2512884) {
		return -1
	}
	if a2 == nil {
		return -1
	}
	if a1 >= NOX_NET_STRUCT_MAX {
		return -3
	}
	var ns *nox_net_struct_t = nox_net_struct_arr[a1]
	if ns == nil {
		return -3
	}
	v5 = (*uint32)(nox_alloc_class_new_obj_zero(nox_alloc_gQueue_3844300))
	if v5 == nil {
		return -1
	}
	*v5 = uint32(uintptr(ns.field_29))
	ns.field_29 = unsafe.Pointer(v5)
	*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*3)) = 1
	*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v5))), 20))) = uint8(*(*byte)(unsafe.Add(unsafe.Pointer(ns.data_2_base), 0)) | 128)
	*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v5))), 21))) = func() uint8 {
		p := &ns.field_28_0
		x := *p
		*p++
		return x
	}()
	*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*4)) = uint32(a3 + 2)
	libc.MemCpy(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v5))), 22), a2, int(a3))
	return int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v5))), 21))))
}
func sub_5551F0(a1 uint32, a2 int8, a3 int32) int32 {
	var i *int32
	if a1 >= NOX_NET_STRUCT_MAX {
		return -3
	}
	var ns *nox_net_struct_t = nox_net_struct_arr[a1]
	if ns == nil {
		return -3
	}
	for i = (*int32)(ns.field_29); i != nil; i = (*int32)(unsafe.Pointer(uintptr(*i))) {
		if a3 != 0 {
			if int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(i))), 21)))) == int32(a2) {
				*(*int32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(int32(0))*3)) = 1
				continue
			}
		} else if *(*int32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(int32(0))*1)) <= int32(dword_5d4594_2495920) {
			*(*int32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(int32(0))*3)) = 1
			continue
		}
	}
	return 0
}
func sub_555250(a1 uint32, a2 *uint32) int32 {
	var (
		v3     int32
		v4     int32
		result int32
	)
	if a1 >= NOX_NET_STRUCT_MAX {
		return 0
	}
	var ns *nox_net_struct_t = nox_net_struct_arr[a1]
	if ns == nil {
		return 0
	}
	v3 = int32(uintptr(ns.field_29))
	if v3 == 0 {
		return 0
	}
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 16))))
	result = v3 + 22
	*a2 = uint32(v4)
	dword_5d4594_2513932 = *(*uint32)(unsafe.Pointer(uintptr(result - 22)))
	return result
}
func sub_555290(a1 uint32, a2 *uint32) int32 {
	var result int32
	if dword_5d4594_2513932 == 0 || a1 >= NOX_NET_STRUCT_MAX || nox_net_struct_arr[a1] == nil {
		return 0
	}
	result = int32(dword_5d4594_2513932 + 22)
	*a2 = *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2513932 + 16)))
	dword_5d4594_2513932 = *(*uint32)(unsafe.Pointer(uintptr(result - 22)))
	return result
}
func nox_xxx_netSend_5552D0(a1 uint32, a2 int8, a3 int32) int32 {
	var (
		i  *int32
		v6 int32
	)
	if a1 >= NOX_NET_STRUCT_MAX {
		return -3
	}
	var ns *nox_net_struct_t = nox_net_struct_arr[a1]
	if ns == nil {
		return -3
	}
	for i = (*int32)(ns.field_29); i != nil; i = (*int32)(unsafe.Pointer(uintptr(*i))) {
		if a3 != 0 {
			if int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(i))), 21)))) == int32(a2) {
				goto LABEL_10
			}
		} else if *(*int32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(int32(0))*3)) != 0 {
		LABEL_10:
			v6 = *(*int32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(int32(0))*4))
			*(*int32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(int32(0))*3)) = 0
			*(*int32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(int32(0))*1)) = int32(dword_5d4594_2495920 + 2000)
			if nox_xxx_sendto_551F90(ns.sock, (*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(i))), 20)), v6, &ns.addr) == -1 {
				return 0
			}
			continue
		}
	}
	return 0
}
func sub_555360(a1 uint32, a2 uint8, a3 int32) int32 {
	var (
		v5 *byte
		v6 *uint64
		v7 *uint32
		v8 [24]byte
	)
	if a1 >= NOX_NET_STRUCT_MAX {
		return -3
	}
	var ns *nox_net_struct_t = nox_net_struct_arr[a1]
	if ns == nil {
		return -3
	}
	*(*uint32)(unsafe.Pointer(&v8[0])) = uint32(uintptr(ns.field_29))
	v5 = &v8[0]
	for *(*uint32)(unsafe.Pointer(v5)) != 0 {
		if a3 != 0 {
			if a3 == 1 {
				if int32(a2) < 32 || int32(a2) > 224 {
					if int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(v5)) + 21)))) < int32(int8(a2)) {
						goto LABEL_17
					}
				} else if int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(v5)) + 21)))) < int32(a2) {
					goto LABEL_17
				}
			} else if a3 == 2 {
			LABEL_17:
				v6 = *(**uint64)(unsafe.Pointer(v5))
				v7 = (*uint32)(ns.field_29)
				if *(**uint32)(unsafe.Pointer(v5)) == v7 {
					ns.field_29 = unsafe.Pointer(uintptr(*v7))
				}
				*(*uint32)(unsafe.Pointer(v5)) = **(**uint32)(unsafe.Pointer(v5))
				*(*uint32)(unsafe.Pointer(v6)) = 0
				nox_alloc_class_free_obj_first(nox_alloc_gQueue_3844300, unsafe.Pointer(v6))
				continue
			}
		} else if int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(v5)) + 21)))) == int32(a2) {
			goto LABEL_17
		}
		v5 = *(**byte)(unsafe.Pointer(v5))
	}
	return 0
}
func sub_565360(a1 int32, a2 *uint16, a3 *int32, a4 uint32, a5 int32, a6 int32) *int32 {
	var (
		v6     *int32
		v7     int32
		v8     *uint16
		result *int32
		v10    *int32
		v11    int32
		v12    *uint16
		v13    int32
		v14    *int32
		v15    int32
		v16    *uint16
		v17    int32
		v18    uint16
		v19    uint16
		v20    *uint16
		v21    uint16
		v22    uint16
		v23    *int32
		v24    int32
		v25    *int32
		v26    *int32
	)
	v6 = a3
	v7 = a6 * 2
	v8 = a2
	result = (*int32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(a3))), a6*(2*2)*a5))))
	v10 = a3
	v23 = (*int32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(a3))), a6*(2*2)*a5))))
	v26 = a3
	v25 = (*int32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(int32(0))*uintptr(a4*2)))
	if uintptr(unsafe.Pointer(a3)) < uintptr(unsafe.Pointer(result)) {
		for {
			result = (*int32)(unsafe.Pointer(uintptr(int32(*v8) & 8191)))
			v11 = int32(*v8) & 0xE000
			v24 = int32(uintptr(unsafe.Pointer((*uint16)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint16(0))*1)))))
			if int32(*v8)&0xE000 != 0 {
				if v11 == 8192 {
					v12 = (*uint16)(unsafe.Pointer(uintptr(a1 + (int32(*v8)&8191)*16)))
					v13 = 2
					for {
						v14 = v6
						v15 = 4
						for {
							if int32(*v12) != math.MinInt16 {
								*(*uint16)(unsafe.Pointer(v14)) = *v12
							}
							v14 = (*int32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v14))), 2))))
							v12 = (*uint16)(unsafe.Add(unsafe.Pointer(v12), unsafe.Sizeof(uint16(0))*1))
							v15--
							if v15 == 0 {
								break
							}
						}
						v6 = (*int32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v6))), v7))))
						v13--
						if v13 == 0 {
							break
						}
					}
					v10 = v26
					result = (*int32)(unsafe.Pointer(uintptr(v7 * 2)))
					v6 = (*int32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v6))), 8))), -(v7 * 2)))))
				} else if v11 == 0x4000 {
					v6 = (*int32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(int32(0))*2))
				}
			} else {
				v16 = (*uint16)(unsafe.Pointer(uintptr(a1 + (int32(*v8)&8191)*16)))
				v17 = 2
				for {
					v18 = *(*uint16)(unsafe.Add(unsafe.Pointer(v16), unsafe.Sizeof(uint16(0))*1))
					v19 = *v16
					v20 = (*uint16)(unsafe.Add(unsafe.Pointer(v16), unsafe.Sizeof(uint16(0))*2))
					*v6 = int32(v19) | int32(v18)<<16
					v21 = *(*uint16)(unsafe.Add(unsafe.Pointer(v20), unsafe.Sizeof(uint16(0))*1))
					v22 = *v20
					v16 = (*uint16)(unsafe.Add(unsafe.Pointer(v20), unsafe.Sizeof(uint16(0))*2))
					*(*int32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(int32(0))*1)) = int32(v22) | int32(v21)<<16
					v6 = (*int32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v6))), v7))))
					v17--
					if v17 == 0 {
						break
					}
				}
				result = (*int32)(unsafe.Pointer(uintptr(v7 * 2)))
				v6 = (*int32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v6))), 8))), -(v7 * 2)))))
			}
			if v6 == v25 {
				v10 = (*int32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v10))), v7*2))))
				result = (*int32)(unsafe.Pointer(uintptr(a4)))
				v26 = v10
				v6 = v10
				v25 = (*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*uintptr(a4*2)))
			}
			if uintptr(unsafe.Pointer(v6)) >= uintptr(unsafe.Pointer(v23)) {
				break
			}
			v8 = (*uint16)(unsafe.Pointer(uintptr(v24)))
		}
	}
	return result
}
func sub_5654A0(a1 int32, a2 *uint8, a3 *int32, a4 uint32, a5 int32, a6 int32) uint32 {
	var (
		v6     int32
		v7     int32
		v8     *int32
		v9     *uint8
		result uint32
		v11    int16
		v12    int16
		v13    uint32
		v14    *int32
		v15    *uint16
		v16    int32
		v17    int32
		v18    uint16
		v19    uint16
		v20    *uint16
		v21    uint16
		v22    uint16
		v23    *uint8
		v24    *int32
		v25    *uint16
		v26    int32
		v27    uint16
		v28    uint16
		v29    *uint16
		v30    uint16
		v31    uint16
		v32    int32
		v33    *uint16
		v34    int32
		v35    uint16
		v36    uint16
		v37    *uint16
		v38    uint16
		v39    uint16
		v40    int32
		v41    *int16
		v42    *int32
		v43    int16
		v44    *int32
		v45    int32
		v46    int32
		v47    *int32
		v48    int32
		v49    *int32
		v50    *int32
		v51    int32
		v52    int32
		v53    *uint32
		v54    int32
		v55    *uint16
		v56    int32
		v57    *int32
		v58    int32
		v59    *int32
		v60    *int32
		v61    uint32
		v62    int32
		v63    uint32
		v64    int32
		v65    int32
		v66    int32
		v67    uint32
		v68    uint32
	)
	v6 = a6 * 2
	v7 = a6 * 2 * a5
	v8 = a3
	v9 = a2
	v68 = uint32(a6 * 2)
	result = uint32(uintptr(unsafe.Pointer(a3))) + uint32(v7*2)
	v61 = result
	v59 = a3
	v60 = (*int32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(int32(0))*uintptr(a4*2)))
	if uint32(uintptr(unsafe.Pointer(a3))) < result {
		for {
			v11 = int16(*(*uint16)(unsafe.Pointer(v9)))
			v9 = (*uint8)(unsafe.Add(unsafe.Pointer(v9), 2))
			v12 = v11
			result = uint32(int32(v11) & 8191)
			v13 = uint32(int32(v12) & 0xE000)
			v66 = int32(uintptr(unsafe.Pointer(v9)))
			if v13 > 0x6000 {
				switch v13 {
				case 0x8000:
					v55 = (*uint16)(unsafe.Pointer(uintptr(a1 + int32(uint16(result))*16)))
					v56 = 2
					for {
						v57 = v8
						v58 = 4
						for {
							if int32(*v55) != math.MinInt16 {
								*(*uint16)(unsafe.Pointer(v57)) = *v55
							}
							v57 = (*int32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v57))), 2))))
							v55 = (*uint16)(unsafe.Add(unsafe.Pointer(v55), unsafe.Sizeof(uint16(0))*1))
							v58--
							if v58 == 0 {
								break
							}
						}
						v8 = (*int32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v8))), v6))))
						v56--
						if v56 == 0 {
							break
						}
					}
					result = uint32(v6 * 2)
					v8 = (*int32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v8))), 8))), -(v6 * 2)))))
					goto LABEL_52
				case 0xA000:
					goto LABEL_16
				case 0xC000:
					v40 = int32(*v9)
					v66 = int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v9), 1)))))
					v41 = (*int16)(unsafe.Pointer(uintptr(a1 + int32(uint16(result))*16)))
					v62 = 2
					for {
						v42 = v8
						v65 = 4
						for {
							v43 = *v41
							if int32(*v41) != math.MinInt16 {
								v44 = v42
								if v40 != 0 {
									v45 = v40
									for {
										*(*uint16)(unsafe.Pointer(v44)) = uint16(v43)
										v44 = (*int32)(unsafe.Add(unsafe.Pointer(v44), unsafe.Sizeof(int32(0))*2))
										v45--
										if v45 == 0 {
											break
										}
									}
								}
							}
							v42 = (*int32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v42))), 2))))
							v41 = (*int16)(unsafe.Add(unsafe.Pointer(v41), unsafe.Sizeof(int16(0))*1))
							v65--
							if v65 == 0 {
								break
							}
						}
						v8 = (*int32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v8))), v68))))
						v62--
						if v62 == 0 {
							break
						}
					}
					v6 = int32(v68)
					result = v68 * 2
					v8 = (*int32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v8))), v40*8))), -int(v68*2)))))
					goto LABEL_52
				}
			} else if v13 == 0x6000 {
				v33 = (*uint16)(unsafe.Pointer(uintptr(a1 + int32(uint16(result))*16)))
				v34 = 2
				for {
					v35 = *(*uint16)(unsafe.Add(unsafe.Pointer(v33), unsafe.Sizeof(uint16(0))*1))
					v36 = *v33
					v37 = (*uint16)(unsafe.Add(unsafe.Pointer(v33), unsafe.Sizeof(uint16(0))*2))
					*v8 = int32(v36) | int32(v35)<<16
					v38 = *(*uint16)(unsafe.Add(unsafe.Pointer(v37), unsafe.Sizeof(uint16(0))*1))
					v39 = *v37
					v33 = (*uint16)(unsafe.Add(unsafe.Pointer(v37), unsafe.Sizeof(uint16(0))*2))
					*(*int32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(int32(0))*1)) = int32(v39) | int32(v38)<<16
					v8 = (*int32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v8))), v6))))
					v34--
					if v34 == 0 {
						break
					}
				}
				result = uint32(v6 * 2)
				v8 = (*int32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v8))), 8))), -(v6 * 2)))))
			} else if v13 != 0 {
				if v13 == 8192 {
					v32 = ((int32(uint16(result)) >> 7) & 62) + 2
					if ((int32(uint16(result)) >> 7) & 62) == -2 {
					LABEL_16:
						result = uint32(uint16(result))
						v32 = int32(*v9)
						v66 = int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v9), 1)))))
					} else {
						result = uint32(uint8(result))
					}
					v46 = 2
					v47 = (*int32)(unsafe.Pointer(uintptr(uint32(a1) + result*16)))
					for {
						v48 = *v47
						v49 = (*int32)(unsafe.Add(unsafe.Pointer(v47), unsafe.Sizeof(int32(0))*1))
						v50 = v8
						if v32 != 0 {
							v51 = v32
							for {
								*v50 = v48
								v50 = (*int32)(unsafe.Add(unsafe.Pointer(v50), unsafe.Sizeof(int32(0))*2))
								v51--
								if v51 == 0 {
									break
								}
							}
							v6 = int32(v68)
						}
						v52 = *v49
						v47 = (*int32)(unsafe.Add(unsafe.Pointer(v49), unsafe.Sizeof(int32(0))*1))
						v53 = (*uint32)(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(int32(0))*1))))
						if v32 != 0 {
							v54 = v32
							for {
								*v53 = uint32(v52)
								v53 = (*uint32)(unsafe.Add(unsafe.Pointer(v53), unsafe.Sizeof(uint32(0))*2))
								v54--
								if v54 == 0 {
									break
								}
							}
							v6 = int32(v68)
						}
						v8 = (*int32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v8))), v6))))
						v46--
						if v46 == 0 {
							break
						}
					}
					result = uint32(v6 * 2)
					v8 = (*int32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v8))), v32*8))), -(v6 * 2)))))
				LABEL_52:
					v9 = (*uint8)(unsafe.Pointer(uintptr(v66)))
					goto LABEL_53
				}
				if v13 == 0x4000 {
					v14 = v8
					v15 = (*uint16)(unsafe.Pointer(uintptr(a1 + int32(*((*uint8)(unsafe.Add(unsafe.Pointer(v9), -2))))*16)))
					v63 = ((result >> 7) & 62) + 2
					v16 = 2
					v17 = int32((v68 >> 2) * 4)
					for {
						v18 = *(*uint16)(unsafe.Add(unsafe.Pointer(v15), unsafe.Sizeof(uint16(0))*1))
						v19 = *v15
						v20 = (*uint16)(unsafe.Add(unsafe.Pointer(v15), unsafe.Sizeof(uint16(0))*2))
						*v14 = int32(v19) | int32(v18)<<16
						v21 = *(*uint16)(unsafe.Add(unsafe.Pointer(v20), unsafe.Sizeof(uint16(0))*1))
						v22 = *v20
						v15 = (*uint16)(unsafe.Add(unsafe.Pointer(v20), unsafe.Sizeof(uint16(0))*2))
						*(*int32)(unsafe.Add(unsafe.Pointer(v14), unsafe.Sizeof(int32(0))*1)) = int32(v22) | int32(v21)<<16
						v14 = (*int32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v14))), v17))))
						v16--
						if v16 == 0 {
							break
						}
					}
					result = v63
					v23 = (*uint8)(unsafe.Pointer(uintptr(v66)))
					v8 = (*int32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(int32(0))*2))
					if v63 != 0 {
						v67 = v63
						for {
							v24 = v8
							v64 = int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v23), 1)))))
							v25 = (*uint16)(unsafe.Pointer(uintptr(a1 + int32(*v23)*16)))
							v26 = 2
							for {
								v27 = *(*uint16)(unsafe.Add(unsafe.Pointer(v25), unsafe.Sizeof(uint16(0))*1))
								v28 = *v25
								v29 = (*uint16)(unsafe.Add(unsafe.Pointer(v25), unsafe.Sizeof(uint16(0))*2))
								*v24 = int32(v28) | int32(v27)<<16
								v30 = *(*uint16)(unsafe.Add(unsafe.Pointer(v29), unsafe.Sizeof(uint16(0))*1))
								v31 = *v29
								v25 = (*uint16)(unsafe.Add(unsafe.Pointer(v29), unsafe.Sizeof(uint16(0))*2))
								*(*int32)(unsafe.Add(unsafe.Pointer(v24), unsafe.Sizeof(int32(0))*1)) = int32(v31) | int32(v30)<<16
								v24 = (*int32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v24))), v17))))
								v26--
								if v26 == 0 {
									break
								}
							}
							v23 = (*uint8)(unsafe.Pointer(uintptr(v64)))
							v8 = (*int32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(int32(0))*2))
							result = func() uint32 {
								p := &v67
								*p--
								return *p
							}()
							if v67 == 0 {
								break
							}
						}
					}
					v6 = int32(v68)
					v66 = int32(uintptr(unsafe.Pointer(v23)))
					goto LABEL_52
				}
			} else {
				result = uint32(uint8(result))
				v8 = (*int32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(int32(0))*uintptr(int32(uint8(result))*2)))
			}
		LABEL_53:
			if v8 == v60 {
				v8 = (*int32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v59))), v6*2))))
				result = a4
				v59 = v8
				v60 = (*int32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(int32(0))*uintptr(a4*2)))
			}
			if uint32(uintptr(unsafe.Pointer(v8))) >= v61 {
				break
			}
		}
	}
	return result
}
func sub_56F1C0() int32 {
	var (
		v0     int32
		result int32
	)
	v0 = int32(libc.GetTime(nil))
	sub_56FF00(v0)
	dword_5d4594_2516352 = 0
	dword_5d4594_2516348 = nox_frame_xxx_2598000
	dword_5d4594_2516344 = 0
	*memmap.PtrUint16(0x587000, 311204) = 0
	dword_5d4594_2516356 = 0x2734945F
	dword_5d4594_2516348 ^= uint32(nox_xxx_protect_56F240())
	dword_5d4594_2516328 = ^dword_5d4594_2516348
	*memmap.PtrUint32(6112660, 2516340) = uint32(nox_xxx_protectionCreateInt_56F400(0))
	sub_56F250()
	result = nox_xxx_protectionCreateInt_56F400(1)
	*memmap.PtrUint32(6112660, 2516332) = uint32(result)
	return result
}
func nox_xxx_protect_56F240() int32 {
	return sub_56FF80(1, -1)
}
func sub_56F250() int32 {
	var (
		v0     int32
		result int32
	)
	v0 = 7
	for {
		result = nox_xxx_protectionCreateStructForInt_56F280(*(*int32)(unsafe.Pointer(&dword_5d4594_2516356)), 0)
		v0--
		dword_5d4594_2516356++
		if v0 == 0 {
			break
		}
	}
	return result
}
func nox_xxx_protectionCreateStructForInt_56F280(a1 int32, a2 int32) int32 {
	var (
		v2 *uint32
		v3 int32
		v4 int32
	)
	v2 = (*uint32)(alloc.Calloc(1, 16))
	if v2 == nil {
		return 0
	}
	*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*3)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*2)) = 0
	v3 = int32(uint32(a1) ^ dword_5d4594_2516348)
	*v2 = uint32(a1) ^ dword_5d4594_2516348
	dword_5d4594_2516328 ^= uint32(v3)
	*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*1)) = uint32(a2)
	v4 = int32(uint32(a2) ^ dword_5d4594_2516348)
	*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*1)) = uint32(a2) ^ dword_5d4594_2516348
	dword_5d4594_2516328 ^= uint32(v4)
	return sub_56F2F0(v2)
}
func sub_56F2F0(a1 *uint32) int32 {
	var (
		v1     int32
		v2     int16
		result int32
		i      int16
		v5     int32
	)
	v1 = int32(dword_5d4594_2516344)
	v2 = 0
	if int32(*memmap.PtrUint16(0x587000, 311204)) != 0 {
		for i = int16(noxRndCounter1.IntClamp(0, int32(*memmap.PtrUint16(0x587000, 311204))-1)); v1 != 0; v2++ {
			if int32(v2) == int32(i) {
				break
			}
			v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 8))))
		}
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3)) = *(*uint32)(unsafe.Pointer(uintptr(v1 + 12)))
		if dword_5d4594_2516348 == 0 {
			nullsub_31(1)
		}
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2)) = uint32(v1)
		*(*uint32)(unsafe.Pointer(uintptr(v1 + 12))) = uint32(uintptr(unsafe.Pointer(a1)))
		v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3)))
		if v5 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v5 + 8))) = uint32(uintptr(unsafe.Pointer(a1)))
			*memmap.PtrUint16(0x587000, 311204)++
		} else {
			*memmap.PtrUint16(0x587000, 311204)++
			dword_5d4594_2516344 = uint32(uintptr(unsafe.Pointer(a1)))
		}
		result = 1
	} else {
		*memmap.PtrUint16(0x587000, 311204)++
		dword_5d4594_2516352 = uint32(uintptr(unsafe.Pointer(a1)))
		dword_5d4594_2516344 = uint32(uintptr(unsafe.Pointer(a1)))
		result = 1
	}
	return result
}
func sub_56F3B0() *uint32 {
	var (
		result *uint32
		v1     *uint32
	)
	result = *(**uint32)(unsafe.Pointer(&dword_5d4594_2516344))
	if dword_5d4594_2516344 != 0 {
		for {
			v1 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*2)))))
			alloc.Free(unsafe.Pointer(result))
			result = v1
			if v1 == nil {
				break
			}
		}
	}
	dword_5d4594_2516328 = 0
	*memmap.PtrUint16(0x587000, 311204) = 0
	dword_5d4594_2516348 = 0
	dword_5d4594_2516352 = 0
	dword_5d4594_2516344 = 0
	return result
}
func nox_xxx_protectionCreateInt_56F400(a1 int32) int32 {
	if nox_xxx_protectionCreateStructForInt_56F280(int32(dword_5d4594_2516356), a1) != 0 {
		return int32(func() uint32 {
			p := &dword_5d4594_2516356
			x := *p
			*p++
			return x
		}())
	}
	nullsub_31(1)
	return 0
}
func nox_xxx_protectionCreateFloat_56F440(a1 float32) int32 {
	if nox_xxx_protectionCreateStructForFloat_56F480(int32(dword_5d4594_2516356), a1) != 0 {
		return int32(func() uint32 {
			p := &dword_5d4594_2516356
			x := *p
			*p++
			return x
		}())
	}
	nullsub_31(1)
	return 0
}
func nox_xxx_protectionCreateStructForFloat_56F480(a1 int32, val float32) int32 {
	var (
		a2 int32 = *(*int32)(unsafe.Pointer(&val))
		v2 *uint32
		v3 int32
		v4 int32
	)
	v2 = (*uint32)(alloc.Calloc(1, 16))
	if v2 == nil {
		return 0
	}
	*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*3)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*2)) = 0
	v3 = int32(uint32(a1) ^ dword_5d4594_2516348)
	*v2 = uint32(a1) ^ dword_5d4594_2516348
	dword_5d4594_2516328 ^= uint32(v3)
	*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*1)) = uint32(a2)
	v4 = int32(uint32(a2) ^ dword_5d4594_2516348)
	*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*1)) = uint32(a2) ^ dword_5d4594_2516348
	dword_5d4594_2516328 ^= uint32(v4)
	return sub_56F2F0(v2)
}
func sub_56F4F0(a1 *int32) int32 {
	var result int32
	result = sub_56F510(*a1)
	if result != 0 {
		*a1 = 0
	}
	return result
}
func sub_56F510(a1 int32) int32 {
	var (
		v1 *uint32
		v2 int32
		v3 int32
		v4 int32
		v5 int32
	)
	v1 = sub_56F590(a1)
	if v1 == nil {
		return 0
	}
	v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*3)))
	if v2 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = *(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*2))
	} else {
		dword_5d4594_2516344 = *(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*2))
	}
	v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*2)))
	if v3 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(v3 + 12))) = *(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*3))
	} else {
		dword_5d4594_2516352 = *(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*3))
	}
	v4 = int32(*v1 ^ dword_5d4594_2516328)
	dword_5d4594_2516328 = uint32(v4)
	v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*1)) ^ uint32(v4))
	*memmap.PtrUint16(0x587000, 311204)--
	dword_5d4594_2516328 = uint32(v5)
	alloc.Free(unsafe.Pointer(v1))
	return 1
}
func sub_56F590(a1 int32) *uint32 {
	var result *uint32
	result = *(**uint32)(unsafe.Pointer(&dword_5d4594_2516344))
	if dword_5d4594_2516344 != 0 {
		for *result != (uint32(a1) ^ dword_5d4594_2516348) {
			result = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*2)))))
			if result == nil {
				goto LABEL_4
			}
		}
	} else {
	LABEL_4:
		nullsub_31(1)
		result = nil
	}
	return result
}
func nox_xxx_protectData_56F5C0() int32 {
	var (
		v0     int32
		v1     int32
		v2     int32
		v3     int32
		v4     uint32
		i      int32
		v6     int32
		v7     int32
		v8     *int32
		v9     *int32
		v10    int32
		v11    int32
		v12    bool
		result int32
		v14    *int32
	)
	if dword_5d4594_2516348 == 0 {
		nullsub_31(1)
	}
	v0 = int32(nox_frame_xxx_2598000)
	v1 = int32(dword_5d4594_2516348)
	v2 = nox_xxx_protect_56F240() ^ v0
	v3 = v2 ^ v1
	dword_5d4594_2516328 = uint32(^v2)
	v4 = uint32(*memmap.PtrUint16(0x587000, 311204))
	for i = 0; i < (int32(*memmap.PtrUint16(0x587000, 311204)) >> 2); v4 = uint32(*memmap.PtrUint16(0x587000, 311204)) {
		v6 = noxRndCounter1.IntClamp(0, int32(v4>>1))
		v7 = noxRndCounter1.IntClamp((int32(*memmap.PtrUint16(0x587000, 311204))>>1)+1, int32(*memmap.PtrUint16(0x587000, 311204))-1)
		if v6 != v7 {
			v14 = (*int32)(unsafe.Pointer(sub_56F6F0(v7)))
			v8 = (*int32)(unsafe.Pointer(sub_56F6F0(v6)))
			sub_56F720(v8, v14)
		}
		i++
	}
	v9 = *(**int32)(unsafe.Pointer(&dword_5d4594_2516344))
	dword_5d4594_2516348 = 0
	if dword_5d4594_2516344 != 0 {
		for {
			v10 = v3 ^ *v9
			v11 = v3 ^ *(*int32)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(int32(0))*1))
			*v9 = v10
			*(*int32)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(int32(0))*1)) = v11
			dword_5d4594_2516328 ^= uint32(v10)
			dword_5d4594_2516328 ^= uint32(*(*int32)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(int32(0))*1)))
			v9 = (*int32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(int32(0))*2)))))
			if v9 == nil {
				break
			}
		}
	}
	result = int32(uint32(v2) ^ dword_5d4594_2516348)
	v12 = uint32(v2) == dword_5d4594_2516348
	*memmap.PtrUint32(6112660, 2516364)++
	dword_5d4594_2516348 ^= uint32(v2)
	if v12 {
		nullsub_31(1)
	}
	return result
}
func sub_56F6F0(a1 int32) *uint32 {
	var (
		result *uint32
		v2     int32
	)
	result = *(**uint32)(unsafe.Pointer(&dword_5d4594_2516344))
	v2 = 0
	if dword_5d4594_2516344 != 0 {
		for v2 != a1 {
			result = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*2)))))
			v2++
			if result == nil {
				goto LABEL_4
			}
		}
	} else {
	LABEL_4:
		nullsub_31(1)
		result = nil
	}
	return result
}
func sub_56F720(a1 *int32, a2 *int32) {
	var (
		v2 int32
		v3 int32
	)
	if a1 == nil || a2 == nil {
		nullsub_31(1)
		return
	}
	v2 = *a1
	v3 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*1))
	*a1 = *a2
	*(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*1)) = *(*int32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(int32(0))*1))
	*a2 = v2
	*(*int32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(int32(0))*1)) = v3
	*memmap.PtrUint32(6112660, 2516360)++
	if dword_5d4594_2516348 == 0 {
		nullsub_31(1)
	}
}
func sub_56F780(a1 int32, a2 int32) *uint32 {
	var (
		result *uint32
		v3     int32
	)
	result = (*uint32)(unsafe.Pointer(uintptr(a1)))
	if uint32(a1) >= 0x2734945F {
		result = sub_56F590(a1)
		if result != nil {
			dword_5d4594_2516328 ^= *(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*1))
			v3 = int32(uint32(a2) ^ dword_5d4594_2516348)
			*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*1)) = uint32(a2) ^ dword_5d4594_2516348
			dword_5d4594_2516328 ^= uint32(v3)
			result = (*uint32)(unsafe.Pointer(uintptr(nox_xxx_protectData_56F5C0())))
		}
	}
	return result
}
func nox_xxx_playerResetProtectionCRC_56F7D0(a1 int32, a2 int32) *uint32 {
	var (
		result *uint32
		v3     int32
	)
	result = (*uint32)(unsafe.Pointer(uintptr(a1)))
	if uint32(a1) >= 0x2734945F {
		result = sub_56F590(a1)
		if result != nil {
			dword_5d4594_2516328 ^= *(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*1))
			v3 = int32(uint32(a2) ^ dword_5d4594_2516348)
			*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*1)) = uint32(a2) ^ dword_5d4594_2516348
			dword_5d4594_2516328 ^= uint32(v3)
			result = (*uint32)(unsafe.Pointer(uintptr(nox_xxx_protectData_56F5C0())))
		}
	}
	return result
}
func sub_56F820(a1 int32, a2 uint8) *uint32 {
	var (
		result *uint32
		v3     int32
	)
	result = (*uint32)(unsafe.Pointer(uintptr(a1)))
	if uint32(a1) >= 0x2734945F {
		result = sub_56F590(a1)
		if result != nil {
			dword_5d4594_2516328 ^= *(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*1))
			v3 = int32(dword_5d4594_2516348 ^ uint32(a2))
			*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*1)) = uint32(v3)
			dword_5d4594_2516328 ^= uint32(v3)
			result = (*uint32)(unsafe.Pointer(uintptr(nox_xxx_protectData_56F5C0())))
		}
	}
	return result
}
func nox_xxx_protectPlayerHPMana_56F870(a1 int32, a2 uint16) *uint32 {
	var (
		result *uint32
		v3     int32
	)
	result = (*uint32)(unsafe.Pointer(uintptr(a1)))
	if uint32(a1) >= 0x2734945F {
		result = sub_56F590(a1)
		if result != nil {
			dword_5d4594_2516328 ^= *(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*1))
			v3 = int32(dword_5d4594_2516348 ^ uint32(a2))
			*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*1)) = uint32(v3)
			dword_5d4594_2516328 ^= uint32(v3)
			result = (*uint32)(unsafe.Pointer(uintptr(nox_xxx_protectData_56F5C0())))
		}
	}
	return result
}
func sub_56F8C0(a1 int32, a2 float32) *uint32 {
	var (
		result *uint32
		v3     *uint32
		v4     int32
	)
	result = (*uint32)(unsafe.Pointer(uintptr(a1)))
	if uint32(a1) >= 0x2734945F {
		result = sub_56F590(a1)
		v3 = result
		if result != nil {
			dword_5d4594_2516328 ^= *(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*1))
			v4 = int32(uint32(uint64(dword_5d4594_2516348) ^ uint64(int64(a2))))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*1)) = uint32(v4)
			dword_5d4594_2516328 ^= uint32(v4)
			result = (*uint32)(unsafe.Pointer(uintptr(nox_xxx_protectData_56F5C0())))
		}
	}
	return result
}
func sub_56F920(a1 int32, a2 int32) *uint32 {
	var (
		result *uint32
		v3     int32
	)
	result = (*uint32)(unsafe.Pointer(uintptr(a1)))
	if uint32(a1) >= 0x2734945F {
		result = sub_56F590(a1)
		if result != nil {
			dword_5d4594_2516328 ^= *(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*1))
			v3 = int32(dword_5d4594_2516348 ^ (uint32(a2) + (dword_5d4594_2516348 ^ *(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*1)))))
			*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*1)) = uint32(v3)
			dword_5d4594_2516328 ^= uint32(v3)
			result = (*uint32)(unsafe.Pointer(uintptr(nox_xxx_protectData_56F5C0())))
		}
	}
	return result
}
func nox_xxx_protectMana_56F9E0(a1 int32, a2 int16) *uint32 {
	var (
		result *uint32
		v3     int32
	)
	result = (*uint32)(unsafe.Pointer(uintptr(a1)))
	if uint32(a1) >= 0x2734945F {
		result = sub_56F590(a1)
		if result != nil {
			dword_5d4594_2516328 ^= *(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*1))
			v3 = int32(dword_5d4594_2516348 ^ (uint32(a2) + (dword_5d4594_2516348 ^ *(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*1)))))
			*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*1)) = uint32(v3)
			dword_5d4594_2516328 ^= uint32(v3)
			result = (*uint32)(unsafe.Pointer(uintptr(nox_xxx_protectData_56F5C0())))
		}
	}
	return result
}
func sub_56FA40(a1 int32, a2 float32) *uint32 {
	var (
		result *uint32
		v3     *uint32
		v4     int32
	)
	result = (*uint32)(unsafe.Pointer(uintptr(a1)))
	if uint32(a1) >= 0x2734945F {
		result = sub_56F590(a1)
		v3 = result
		if result != nil {
			dword_5d4594_2516328 ^= *(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*1))
			v4 = int32(uint32(uint64(dword_5d4594_2516348) ^ uint64(int64(float64(dword_5d4594_2516348^*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*1)))+float64(a2)))))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*1)) = uint32(v4)
			dword_5d4594_2516328 ^= uint32(v4)
			result = (*uint32)(unsafe.Pointer(uintptr(nox_xxx_protectData_56F5C0())))
		}
	}
	return result
}
func nox_xxx_protectionStringCRC_56FAC0(a1 *int32, a2 uint32) int32 {
	var (
		v2     *int32
		result int32
		i      uint32
		v5     int32
	)
	v2 = a1
	result = 0
	for i = a2 >> 2; i != 0; i-- {
		v5 = *v2
		v2 = (*int32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(int32(0))*1))
		result ^= v5
	}
	return result
}
func nox_xxx_protectionStringCRCLen_56FAE0(a1 *int32, a2 uint32) int32 {
	var result int32
	result = 0
	if a1 != nil {
		result = nox_xxx_protectionStringCRC_56FAC0(a1, a2)
	}
	return result
}
func sub_56FB00(a1 *int32, a2 uint32, a3 int32) int32 {
	var v3 *uint32
	if uint32(a3) >= 0x2734945F {
		v3 = sub_56F590(a3)
		if v3 != nil && (dword_5d4594_2516348^uint32(nox_xxx_protectionStringCRCLen_56FAE0(a1, a2))) == *(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*1)) {
			return 1
		}
		nullsub_31(1)
	}
	return 0
}
func sub_56FB60(item *nox_object_t) int32 {
	var (
		result int32
		v2     int32
		v3     int32
		v4     int32
		v5     *int32
		v6     int32
		v7     *int32
	)
	result = 0
	if item != nil {
		v2 = sub_4E4C00(item)
		v3 = int32(uint16(nox_xxx_unitGetHP_4EE780(item))) ^ v2
		v4 = sub_4E4C10(item) ^ v3
		v5 = (*int32)(unsafe.Pointer(uintptr(nox_object_getInitData_4E4C30(item))))
		v6 = sub_4E4C50(item)
		if v5 != nil && v6 > 0 {
			v4 ^= nox_xxx_protectionStringCRC_56FAC0(v5, uint32(v6))
		}
		v7 = (*int32)(unsafe.Pointer(uintptr(sub_4E4C80(item))))
		if v7 != nil {
			if libc.StrLen((*byte)(unsafe.Pointer(v7))) != 0 {
				v4 ^= nox_xxx_protectionStringCRC_56FAC0(v7, uint32(libc.StrLen((*byte)(unsafe.Pointer(v7)))))
			}
		}
		result = v4
	}
	return result
}
func nox_xxx_protect_56FBF0(a1 int32, item *nox_object_t) int32 {
	var (
		result int32
		v3     *uint32
		v4     *uint32
		v5     int32
	)
	result = a1
	if uint32(a1) >= 0x2734945F {
		v3 = sub_56F590(a1)
		v4 = v3
		if v3 != nil {
			dword_5d4594_2516328 ^= *(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*1))
			v5 = int32(uint32(sub_56FB60(item)) ^ *(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*1)))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*1)) = uint32(v5)
			result = int32(uint32(v5) ^ dword_5d4594_2516328)
			dword_5d4594_2516328 ^= uint32(v5)
		} else {
			nullsub_31(1)
		}
	}
	return result
}
func nox_xxx_protect_56FC50(a1 int32, object *nox_object_t) int32 {
	var (
		a2     *int32 = (*int32)(unsafe.Pointer(object))
		result int32
		v3     *uint32
		v4     *uint32
		v5     int32
	)
	result = a1
	if uint32(a1) >= 0x2734945F {
		v3 = sub_56F590(a1)
		v4 = v3
		if v3 != nil {
			dword_5d4594_2516328 ^= *(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*1))
			v5 = int32(uint32(sub_56FB60((*nox_object_t)(unsafe.Pointer(a2)))) ^ *(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*1)))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*1)) = uint32(v5)
			result = int32(uint32(v5) ^ dword_5d4594_2516328)
			dword_5d4594_2516328 ^= uint32(v5)
		} else {
			nullsub_31(1)
		}
	}
	return result
}
func sub_56FCB0(a1 int32, a2 int32) int32 {
	var result int32
	result = 0
	if a2 != 0 {
		result = 1 << (a1 % 32)
	}
	return result
}
func nox_xxx_playerAwardSpellProtectionCRC_56FCE0(a1 int32, a2 int32, a3 int32) int32 {
	var (
		result int32
		v4     *uint32
		v5     *uint32
	)
	result = a1
	if uint32(a1) >= 0x2734945F {
		v4 = sub_56F590(a1)
		v5 = v4
		if v4 != nil {
			dword_5d4594_2516328 ^= *(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*1))
			result = int32(dword_5d4594_2516348 ^ (dword_5d4594_2516348 ^ *(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*1)) | uint32(sub_56FCB0(a2, a3))))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*1)) = uint32(result)
			dword_5d4594_2516328 ^= uint32(result)
		} else {
			nullsub_31(1)
		}
	}
	return result
}
func nox_xxx_playerApplyProtectionCRC_56FD50(a1 int32, a2p unsafe.Pointer, a3 int32) int32 {
	var (
		a2 int32 = int32(uintptr(a2p))
		v3 int32
		v4 *uint32
		v5 int32
		v6 *int32
		v8 *uint32
	)
	v3 = 0
	if uint32(a1) >= 0x2734945F {
		v4 = sub_56F590(a1)
		v8 = v4
		if v4 != nil {
			v5 = 1
			if a3 > 1 {
				v6 = (*int32)(unsafe.Pointer(uintptr(a2 + 4)))
				for {
					v3 |= sub_56FCB0(func() int32 {
						p := &v5
						x := *p
						*p++
						return x
					}(), *v6)
					v6 = (*int32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(int32(0))*1))
					if v5 >= a3 {
						break
					}
				}
				v4 = v8
			}
			if (uint32(v3) ^ dword_5d4594_2516348) == *(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*1)) {
				return 1
			}
		}
		nullsub_31(1)
	}
	return 0
}
func nox_xxx_unkDoubleSmth_56FE30() float64 {
	var v0 float64
	*memmap.PtrUint64(6112660, 2516412) = *memmap.PtrUint64(6112660, 2516404)
	*memmap.PtrUint64(6112660, 2516404) = *memmap.PtrUint64(6112660, 2516396)
	*memmap.PtrUint32(6112660, 2516396) = *memmap.PtrUint32(6112660, 2516388)
	*memmap.PtrUint32(6112660, 2516400) = *memmap.PtrUint32(6112660, 2516392)
	v0 = *mem_getDoublePtr(6112660, 0x2665A4)**mem_getDoublePtr(0x581450, 0x2C70) + *mem_getDoublePtr(6112660, 0x2665B4)**mem_getDoublePtr(0x581450, 0x2C68) + *mem_getDoublePtr(6112660, 0x2665BC)**mem_getDoublePtr(0x581450, 0x2C60) + *mem_getDoublePtr(6112660, 0x2665BC)**mem_getDoublePtr(0x581450, 0x2C58) + *mem_getDoublePtr(6112660, 2516420)
	math.Floor(v0)
	*mem_getDoublePtr(6112660, 0x2665A4) = v0 - v0
	*mem_getDoublePtr(6112660, 2516420) = v0 * *mem_getDoublePtr(0x581450, 0x2C50)
	return *mem_getDoublePtr(6112660, 0x2665A4)
}
func sub_56FF00(a1 int32) {
	var (
		v1 int32
		v2 *uint8
		v3 uint32
		v4 int32
	)
	v1 = a1
	if a1 == 0 {
		v1 = -1
	}
	v2 = (*uint8)(memmap.PtrOff(6112660, 2516388))
	for {
		v2 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), 8))
		v3 = ((uint32(v1<<13) ^ uint32(v1)) >> 17) ^ uint32(v1<<13) ^ uint32(v1)
		v1 = int32((v3 * 32) ^ v3)
		*((*float64)(unsafe.Add(unsafe.Pointer((*float64)(unsafe.Pointer(v2))), -int(unsafe.Sizeof(float64(0))*1)))) = float64(uint32(v1)) * *mem_getDoublePtr(0x581450, 0x2C50)
		if int32(uintptr(unsafe.Pointer(v2))) >= int32(uintptr(memmap.PtrOff(6112660, 2516428))) {
			break
		}
	}
	v4 = 19
	for {
		nox_xxx_unkDoubleSmth_56FE30()
		v4--
		if v4 == 0 {
			break
		}
	}
	dword_5d4594_2516380 = 0
	*memmap.PtrUint32(6112660, 2516376) = 99
	dword_5d4594_2516372 = 100
}
func sub_56FF80(a1 int32, a2 int32) int32 {
	var (
		v2     int64
		result int32
	)
	*memmap.PtrUint32(6112660, 2516376) = uint32(a2)
	dword_5d4594_2516380 = uint32(a1)
	dword_5d4594_2516372 = uint32(a2 - a1 + 1)
	v2 = int64(nox_xxx_unkDoubleSmth_56FE30() * float64(dword_5d4594_2516372))
	if uint32(int32(v2)) < uint32(*(*int32)(unsafe.Pointer(&dword_5d4594_2516372))) {
		result = int32(int64(dword_5d4594_2516380) + v2)
	} else {
		result = int32(dword_5d4594_2516372 + dword_5d4594_2516380)
	}
	return result
}
func nox_xxx_netGetUnitCodeServ_578AC0(item *nox_object_t) uint32 {
	var result uint32
	if item == nil {
		return 0
	}
	if item.field_9 >= 0x8000 {
		return 0
	}
	result = item.extent
	if result >= 0x8000 {
		return 0
	}
	if (item.obj_class & 0x20400000) == 0 {
		return item.field_9
	}
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&result))), 1)) |= 128
	return result
}
func nox_xxx_netGetUnitCodeCli_578B00(a1 int32) uint32 {
	var result uint32
	if a1 == 0 {
		return 0
	}
	result = *(*uint32)(unsafe.Pointer(uintptr(a1 + 128)))
	if result >= 0x8000 {
		return 0
	}
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + 112)))&0x20400000 != 0 {
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&result))), 1)) |= 128
	}
	return result
}
func nox_xxx_netClearHighBit_578B30(a1 int16) int32 {
	return int32(a1) & math.MaxInt16
}
func nox_xxx_packetDynamicUnitCode_578B40(a1 int32) int32 {
	var (
		result int32
		v2     int32
	)
	result = a1
	if (a1 & 0x8000) == 0x8000 {
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&result))), 1)) &= math.MaxInt8
		v2 = nox_xxx_netGetUnitByExtent_4ED020(result)
		if v2 != 0 {
			result = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 36))))
		} else {
			result = 0
		}
	}
	return result
}
func nox_xxx_netTestHighBit_578B70(a1 uint32) uint32 {
	return (a1 >> 15) & 1
}
func sub_578B80() *uint32 {
	var (
		v0 *uint32
		v2 *uint32
		v3 int32
	)
	v0 = (*uint32)(operator_new(24))
	if v0 == nil {
		return nil
	}
	v2 = v0
	*v0 = uint32(uintptr(operator_new(0x10000)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*1)) = 0
	sub_57DF00((*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*2)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*5)) = uint32(uintptr(operator_new(0x10000)))
	v3 = 0
	for {
		v3 += 2
		*(*uint16)(unsafe.Pointer(uintptr(uint32(v3) + *(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*5)) - 2))) = math.MaxUint16
		if uint32(v3) >= 0x10000 {
			break
		}
	}
	return v2
}
func sub_578BA0(a1 uint32) uint32 {
	return (a1 >> 1) + a1 + 32
}
func sub_578BB0(a1 *unsafe.Pointer, a2 int32, a3 *uint8, a4 int32) int32 {
	return sub_57D1C0(a1, a2, a3, a4)
}
func sub_578BD0(lpMem unsafe.Pointer) {
	if lpMem != nil {
		sub_57D150((*unsafe.Pointer)(lpMem))
		operator_delete(lpMem)
	}
}
func sub_578BF0() *uint32 {
	var (
		v0     *uint32
		result *uint32
	)
	v0 = (*uint32)(operator_new(152))
	if v0 != nil {
		result = sub_57E9A0(v0)
	} else {
		result = nil
	}
	return result
}
func nox_xxx_nxzDecompress_578C10(a1 *uint32, a2 *uint8, a3 *uint32, a4 uint32, a5 *uint32) int32 {
	return nox_xxx_nxzDecompressImpl_57EA80(a1, a2, a3, a4, a5)
}
func sub_578C30(a1 int32) int32 {
	return sub_57EA60(a1)
}
func sub_578C40(lpMem unsafe.Pointer) {
	if lpMem != nil {
		sub_57EA00((*unsafe.Pointer)(lpMem))
		operator_delete(lpMem)
	}
}
func sub_578C60() int32 {
	if sub_44E560() != nil {
		nox_client_lockScreenBriefing_450160(math.MaxUint8, 1, 0)
		sub_4A2530()
	}
	return 1
}
func sub_578C90(a1 int32) int32 {
	dword_587000_311372 = uint32(a1)
	*memmap.PtrUint8(6112660, 2516476) |= uint8(int8(1 << a1))
	nox_xxx_cliPlayMapIntro_44E0B0(1)
	sub_413960()
	sub_477530(0)
	return nox_client_quit_4460C0()
}
func sub_578CD0() int32 {
	var (
		result int32
		v1     uint8
		v2     *byte
	)
	_ = v2
	var v3 [16]byte
	var v4 [128]byte
	result = int32(dword_587000_311372)
	if *(*int32)(unsafe.Pointer(&dword_587000_311372)) != -1 {
		v1 = *memmap.PtrUint8(0x587000, 311380)
		libc.StrCpy(&v3[0], *(**byte)(memmap.PtrOff(0x587000, dword_587000_311372*4+0x7310)))
		v2 = &v3[libc.StrLen(&v3[0])]
		*(*uint32)(unsafe.Pointer(v2)) = *memmap.PtrUint32(0x587000, 311376)
		*(*byte)(unsafe.Add(unsafe.Pointer(v2), 4)) = byte(v1)
		if nox_game_setMovieFile_4CB230(&v3[0], &v4[0]) != 0 {
			sub_4B0300(&v4[0])
			sub_4B0640(unsafe.Pointer(cgoFuncAddr(sub_578C60)))
			result = nox_client_drawGeneral_4B0340(1)
		} else {
			result = sub_578C60()
		}
	}
	return result
}
func nox_xxx_GetEndgameDialog_578D80() *byte {
	if dword_587000_311372 != 0 {
		if dword_587000_311372 == 1 {
			if (int32(*memmap.PtrUint8(6112660, 2516476)) & 1) == 0 {
				return (*byte)(memmap.PtrOff(0x587000, 311416))
			}
			if (int32(*memmap.PtrUint8(6112660, 2516476)) & 4) == 0 {
				return (*byte)(memmap.PtrOff(0x587000, 311432))
			}
		} else if dword_587000_311372 == 2 {
			if (int32(*memmap.PtrUint8(6112660, 2516476)) & 2) == 0 {
				return (*byte)(memmap.PtrOff(0x587000, 311448))
			}
			if (int32(*memmap.PtrUint8(6112660, 2516476)) & 1) == 0 {
				return (*byte)(memmap.PtrOff(0x587000, 311464))
			}
		}
	} else {
		if (int32(*memmap.PtrUint8(6112660, 2516476)) & 2) == 0 {
			return (*byte)(memmap.PtrOff(0x587000, 311384))
		}
		if (int32(*memmap.PtrUint8(6112660, 2516476)) & 4) == 0 {
			return (*byte)(memmap.PtrOff(0x587000, 311400))
		}
	}
	return nil
}
func sub_578DE0(a1 int8) int8 {
	var result int8
	result = a1
	*memmap.PtrUint8(6112660, 2516476) = uint8(a1)
	return result
}
func sub_578DF0() uint8 {
	return *memmap.PtrUint8(6112660, 2516476)
}
func sub_578E00() {
	dword_587000_311372 = math.MaxUint32
}
func nox_xxx_waypointGetList_579860() unsafe.Pointer {
	return *(*unsafe.Pointer)(unsafe.Pointer(&nox_xxx_waypointsHead_2523752))
}
func nox_xxx_waypointNext_579870(a1 int32) int32 {
	var result int32
	if a1 != 0 {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 484))))
	} else {
		result = 0
	}
	return result
}
func sub_579890() int32 {
	return int32(dword_5d4594_2523756)
}
func sub_5798A0(a1 int32) int32 {
	var result int32
	if a1 != 0 {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 484))))
	} else {
		result = 0
	}
	return result
}
func nox_xxx_waypoint_5798C0() uint32 {
	var (
		v0     *uint32
		result uint32
	)
	v0 = *(**uint32)(unsafe.Pointer(&nox_xxx_waypointsHead_2523752))
	result = 1
	if nox_xxx_waypointsHead_2523752 != 0 {
		for {
			if result <= *v0 {
				result = *v0 + 1
			}
			v0 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(uint32(0))*121)))))
			if v0 == nil {
				break
			}
		}
	}
	return result
}
func nox_xxx_waypointNew_5798F0(a1 float32, a2 float32) *uint32 {
	var (
		v2 *uint32
		v3 uint32
		v4 int32
	)
	v2 = (*uint32)(alloc.Calloc(1, 516))
	v3 = nox_xxx_waypoint_5798C0()
	v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*120)))
	*v2 = v3
	*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v2))), unsafe.Sizeof(float32(0))*2))) = a1
	*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v2))), unsafe.Sizeof(float32(0))*3))) = a2
	*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*120)) = uint32(v4 | 1)
	*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*121)) = nox_xxx_waypointsHead_2523752
	if nox_xxx_waypointsHead_2523752 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(nox_xxx_waypointsHead_2523752 + 488))) = uint32(uintptr(unsafe.Pointer(v2)))
	}
	nox_xxx_waypointsHead_2523752 = uint32(uintptr(unsafe.Pointer(v2)))
	if noxflags.HasGame(noxflags.GameHost) {
		nox_xxx_waypointMapRegister_5179B0(int32(uintptr(unsafe.Pointer(v2))))
	}
	return v2
}
func nox_xxx_waypointNewNotMap_579970(a1 int32, a2 float32, a3 float32) *float32 {
	var (
		result *float32
		v4     int32
	)
	result = (*float32)(alloc.Calloc(1, 516))
	*(*uint32)(unsafe.Pointer(result)) = uint32(a1)
	*(*float32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(float32(0))*3)) = a3
	v4 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*120))) | 1)
	*(*float32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(float32(0))*2)) = a2
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*120))) = uint32(v4)
	*(*float32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(float32(0))*121)) = *(*float32)(unsafe.Pointer(&dword_5d4594_2523756))
	dword_5d4594_2523756 = uint32(uintptr(unsafe.Pointer(result)))
	return result
}
func nox_xxx_waypoint_5799C0() *byte {
	var (
		v0 int32
		v1 int32
	)
	v0 = int32(dword_5d4594_2523756)
	if dword_5d4594_2523756 != 0 {
		for {
			v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(v0 + 484))))
			*(*uint32)(unsafe.Pointer(uintptr(v0 + 484))) = nox_xxx_waypointsHead_2523752
			if nox_xxx_waypointsHead_2523752 != 0 {
				*(*uint32)(unsafe.Pointer(uintptr(nox_xxx_waypointsHead_2523752 + 488))) = uint32(v0)
			}
			nox_xxx_waypointsHead_2523752 = uint32(v0)
			if noxflags.HasGame(noxflags.GameHost) {
				nox_xxx_waypointMapRegister_5179B0(v0)
			}
			v0 = v1
			if v1 == 0 {
				break
			}
		}
	}
	dword_5d4594_2523756 = 0
	return sub_579A30()
}
func sub_579A30() *byte {
	var (
		result *byte
		i      *byte
		v2     int32
		v3     *byte
		v4     int8
		j      *byte
		v6     int32
		v7     *uint8
	)
	result = (*byte)(nox_xxx_waypointGetList_579860())
	for i = result; result != nil; i = result {
		*(*byte)(unsafe.Add(unsafe.Pointer(i), 477)) = 0
		v2 = 0
		if *(*byte)(unsafe.Add(unsafe.Pointer(i), 476)) != 0 {
			v3 = (*byte)(unsafe.Add(unsafe.Pointer(i), 96))
			for {
				v4 = int8(*v3)
				v3 = (*byte)(unsafe.Add(unsafe.Pointer(v3), 8))
				*(*byte)(unsafe.Add(unsafe.Pointer(i), 477)) |= byte(v4)
				v2++
				if v2 >= int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(i), 476)))) {
					break
				}
			}
		}
		for j = (*byte)(nox_xxx_waypointGetList_579860()); j != nil; j = (*byte)(unsafe.Pointer(uintptr(nox_xxx_waypointNext_579870(int32(uintptr(unsafe.Pointer(j))))))) {
			v6 = 0
			if *(*byte)(unsafe.Add(unsafe.Pointer(j), 476)) != 0 {
				v7 = (*uint8)(unsafe.Add(unsafe.Pointer(j), 96))
				for {
					if *((**byte)(unsafe.Add(unsafe.Pointer((**byte)(unsafe.Pointer(v7))), -int(unsafe.Sizeof((*byte)(nil))*1)))) == i {
						*(*byte)(unsafe.Add(unsafe.Pointer(i), 477)) |= byte(*v7)
					}
					v6++
					v7 = (*uint8)(unsafe.Add(unsafe.Pointer(v7), 8))
					if v6 >= int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(j), 476)))) {
						break
					}
				}
			}
		}
		result = (*byte)(unsafe.Pointer(uintptr(nox_xxx_waypointNext_579870(int32(uintptr(unsafe.Pointer(i)))))))
	}
	return result
}
func sub_579AD0(a1 float32, a2 float32) *float32 {
	var (
		v2 int32
		v3 int32
		v4 float64
		v5 float64
		v6 float64
		i  float32
	)
	v2 = int32(nox_xxx_waypointsHead_2523752)
	v3 = 0
	for i = 100.0; v2 != 0; v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 484)))) {
		v4 = float64(*(*float32)(unsafe.Pointer(uintptr(v2 + 8))) - a1)
		v5 = float64(*(*float32)(unsafe.Pointer(uintptr(v2 + 12))) - a2)
		v6 = v5*v5 + v4*v4
		if v6 < float64(i) {
			i = float32(v6)
			v3 = v2
		}
	}
	return (*float32)(unsafe.Pointer(uintptr(v3)))
}
func sub_579B30(lpMem unsafe.Pointer) {
	var (
		i  int32
		v2 int32
		v3 int32
		v4 *unsafe.Pointer
		v5 int32
		v6 *uint32
		v7 int32
		v8 int32
	)
	for i = int32(nox_xxx_waypointsHead_2523752); i != 0; i = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 484)))) {
		v2 = 0
		v3 = int32(*(*uint8)(unsafe.Pointer(uintptr(i + 476))))
		if v3 > 0 {
			v4 = (*unsafe.Pointer)(unsafe.Pointer(uintptr(i + 92)))
			for *v4 != lpMem {
				v2++
				v4 = (*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*2))
				if v2 >= v3 {
					goto LABEL_11
				}
			}
			v5 = v2
			if v2 < v3-1 {
				v6 = (*uint32)(unsafe.Pointer(uintptr(i + v2*8 + 92)))
				for {
					v5++
					*v6 = *(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*2))
					*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*1)) = *(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*3))
					v6 = (*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*2))
					if v5 >= int32(*(*uint8)(unsafe.Pointer(uintptr(i + 476))))-1 {
						break
					}
				}
			}
			*(*uint8)(unsafe.Pointer(uintptr(i + 476)))--
		}
	LABEL_11:
	}
	v7 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(lpMem)), unsafe.Sizeof(uint32(0))*121))))
	if v7 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(v7 + 488))) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(lpMem)), unsafe.Sizeof(uint32(0))*122)))
	}
	v8 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(lpMem)), unsafe.Sizeof(uint32(0))*122))))
	if v8 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(v8 + 484))) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(lpMem)), unsafe.Sizeof(uint32(0))*121)))
	} else {
		nox_xxx_waypointsHead_2523752 = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(lpMem)), unsafe.Sizeof(uint32(0))*121)))
	}
	if noxflags.HasGame(noxflags.GameHost) {
		sub_517A70(int32(uintptr(lpMem)))
	}
	lpMem = nil
}
func sub_579C00() *uint32 {
	var (
		result *uint32
		v1     *uint32
	)
	result = *(**uint32)(unsafe.Pointer(&nox_xxx_waypointsHead_2523752))
	if nox_xxx_waypointsHead_2523752 != 0 {
		for {
			v1 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*121)))))
			alloc.Free(unsafe.Pointer(result))
			result = v1
			if v1 == nil {
				break
			}
		}
		nox_xxx_waypointsHead_2523752 = 0
	} else {
		nox_xxx_waypointsHead_2523752 = 0
	}
	return result
}
func nox_server_getWaypointById_579C40(a1 int32) *uint32 {
	var result *uint32
	result = *(**uint32)(unsafe.Pointer(&nox_xxx_waypointsHead_2523752))
	if nox_xxx_waypointsHead_2523752 == 0 {
		return nil
	}
	for *result != uint32(a1) {
		result = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*121)))))
		if result == nil {
			return nil
		}
	}
	return result
}
func sub_579C60(a1 int32) int32 {
	var result int32
	result = int32(dword_5d4594_2523756)
	if dword_5d4594_2523756 == 0 {
		return 0
	}
	for *(*uint32)(unsafe.Pointer(uintptr(result + 4))) != uint32(a1) {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(result + 484))))
		if result == 0 {
			return 0
		}
	}
	return result
}
func sub_579C80(a1 int32) *uint32 {
	var result *uint32
	result = *(**uint32)(unsafe.Pointer(&dword_5d4594_2523756))
	if dword_5d4594_2523756 == 0 {
		return nil
	}
	for *result != uint32(a1) {
		result = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*121)))))
		if result == nil {
			return nil
		}
	}
	return result
}
func sub_579CA0() int32 {
	var (
		v0 *uint32
		v1 *uint32
		v2 int32
		v3 *int32
		v4 *int32
		v5 int32
	)
	v0 = *(**uint32)(unsafe.Pointer(&dword_5d4594_2523756))
	if dword_5d4594_2523756 != 0 {
		for {
			*(*uint32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(uint32(0))*1)) = *v0
			v0 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(uint32(0))*121)))))
			if v0 == nil {
				break
			}
		}
		v0 = *(**uint32)(unsafe.Pointer(&dword_5d4594_2523756))
	}
	v1 = v0
	if v0 == nil {
		return 1
	}
	for {
		v2 = 0
		if int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v1))), 476)))) != 0 {
			break
		}
	LABEL_9:
		v1 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*121)))))
		if v1 == nil {
			return 1
		}
	}
	v3 = (*int32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*23))))
	v4 = (*int32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*87))))
	for {
		v5 = sub_579C60(*v4)
		*v3 = v5
		if v5 == 0 {
			return 0
		}
		v2++
		v4 = (*int32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(int32(0))*1))
		v3 = (*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*2))
		if v2 >= int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v1))), 476)))) {
			goto LABEL_9
		}
	}
}
func sub_579D20() int32 {
	var (
		v0 uint32
		v1 *uint32
		v2 uint32
		v3 *uint32
		v4 int32
		v5 *int32
		v6 *int32
		v7 int32
		v9 int8
	)
	v0 = nox_xxx_waypoint_5798C0()
	v1 = *(**uint32)(unsafe.Pointer(&dword_5d4594_2523756))
	if dword_5d4594_2523756 != 0 {
		for {
			v2 = *v1
			*v1 = v0
			*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*1)) = v2
			v1 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*121)))))
			v0++
			if v1 == nil {
				break
			}
		}
		v1 = *(**uint32)(unsafe.Pointer(&dword_5d4594_2523756))
	}
	v3 = v1
	if v1 == nil {
		return 1
	}
	for {
		v4 = 0
		v9 = 0
		if int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v3))), 476)))) != 0 {
			v5 = (*int32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*23))))
			v6 = (*int32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*87))))
			for {
				v7 = sub_579C60(*v6)
				*v5 = v7
				if v7 != 0 {
					v5 = (*int32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(int32(0))*2))
					v9++
				}
				v4++
				v6 = (*int32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(int32(0))*1))
				if v4 >= int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v3))), 476)))) {
					break
				}
			}
		}
		*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v3))), 476))) = uint8(v9)
		v3 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*121)))))
		if v3 == nil {
			break
		}
	}
	return 1
}
func nox_xxx_waypointDeleteAll_579DD0() {
	var (
		v0 *uint32
		v1 *uint32
	)
	v0 = *(**uint32)(unsafe.Pointer(&nox_xxx_waypointsHead_2523752))
	if nox_xxx_waypointsHead_2523752 != 0 {
		for {
			v1 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(uint32(0))*121)))))
			if noxflags.HasGame(noxflags.GameHost) {
				sub_517A70(int32(uintptr(unsafe.Pointer(v0))))
			}
			alloc.Free(unsafe.Pointer(v0))
			v0 = v1
			if v1 == nil {
				break
			}
		}
	}
	nox_xxx_waypointsHead_2523752 = 0
	dword_5d4594_2523756 = 0
	dword_5d4594_2523760 = 0
}
func nox_xxx_waypointByName_579E30(a1 *byte) *byte {
	var i *byte
	for i = (*byte)(nox_xxx_waypointGetList_579860()); i != nil; i = (*byte)(unsafe.Pointer(uintptr(nox_xxx_waypointNext_579870(int32(uintptr(unsafe.Pointer(i))))))) {
		if nox_server_strcmpWithoutMapname_4DA3F0((*byte)(unsafe.Add(unsafe.Pointer(i), 16)), a1) != 0 {
			break
		}
	}
	return i
}
func sub_579E70() *uint32 {
	var result *uint32
	result = (*uint32)(alloc.Calloc(1, 516))
	if result != nil {
		*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*120)) |= 0x1000000
	}
	return result
}
func sub_579E90(a1 int32) {
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 480))) |= 0x1000000
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 484))) = dword_5d4594_2523756
	if dword_5d4594_2523756 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2523756 + 488))) = uint32(a1)
	}
	dword_5d4594_2523756 = uint32(a1)
	if noxflags.HasGame(noxflags.GameHost) {
		nox_xxx_waypointMapRegister_5179B0(a1)
	}
}
func sub_579EE0(a1 int32, a2 uint8) int32 {
	return bool2int((int32(a2) & int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 477))))) != 0)
}
func nox_xxx_waypoint_579F00(a1 *uint32, a2 int32) int32 {
	var (
		v2  int32
		v3  int32
		i   *uint8
		v5  int32
		v6  unsafe.Pointer
		v7  *float32
		v8  float32
		v9  float32
		v10 float32
		v12 float4
	)
	v2 = 0
	if noxflags.HasGame(noxflags.GameModeCTF) {
		if a2 != 0 {
			v3 = int32(uintptr(unsafe.Pointer(noxServer.firstServerObject())))
			if v3 != 0 {
				for (*(*uint32)(unsafe.Pointer(uintptr(v3 + 8)))&0x10000000) == 0 || nox_xxx_servCompareTeams_419150(a2+48, v3+48) != 0 {
					v3 = int32(uintptr(unsafe.Pointer((*nox_object_t)(unsafe.Pointer(uintptr(v3))).Next())))
					if v3 == 0 {
						goto LABEL_9
					}
				}
				v2 = v3
			}
		}
	}
LABEL_9:
	dword_5d4594_2523760 = 0
	for i = (*uint8)(nox_xxx_waypointGetList_579860()); i != nil; i = (*uint8)(unsafe.Pointer(uintptr(nox_xxx_waypointNext_579870(int32(uintptr(unsafe.Pointer(i))))))) {
		if sub_579EE0(int32(uintptr(unsafe.Pointer(i))), 128) != 0 && int32(*(*uint8)(unsafe.Add(unsafe.Pointer(i), 480)))&1 != 0 {
			dword_5d4594_2523760++
		}
	}
	if dword_5d4594_2523760 == 0 {
		return 0
	}
	v5 = noxRndCounter1.IntClamp(0, int32(dword_5d4594_2523760-1))
	v6 = nox_xxx_waypointGetList_579860()
	if v6 == nil {
		return 0
	}
	for {
		if sub_579EE0(int32(uintptr(v6)), 128) == 0 {
			goto LABEL_24
		}
		if (int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(v6)), 480)))) & 1) == 0 {
			goto LABEL_24
		}
		if noxflags.HasGame(noxflags.GameModeCTF) {
			if v2 != 0 {
				if a2 != 0 {
					v7 = *(**float32)(unsafe.Pointer(uintptr(v2 + 748)))
					v12.field_0 = *v7
					v8 = *((*float32)(unsafe.Add(unsafe.Pointer((*float32)(v6)), unsafe.Sizeof(float32(0))*3)))
					v9 = *(*float32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(float32(0))*1))
					v10 = *((*float32)(unsafe.Add(unsafe.Pointer((*float32)(v6)), unsafe.Sizeof(float32(0))*2)))
					v12.field_4 = v9
					v12.field_8 = v10
					v12.field_C = v8
					if nox_xxx_mapTraceRay_535250(&v12, nil, nil, 9) == 1 {
						goto LABEL_24
					}
				}
			}
		}
		if v5 == 0 {
			break
		}
		v5--
	LABEL_24:
		v6 = unsafe.Pointer(uintptr(nox_xxx_waypointNext_579870(int32(uintptr(v6)))))
		if v6 == nil {
			return 0
		}
	}
	*a1 = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(v6)), unsafe.Sizeof(uint32(0))*2)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(v6)), unsafe.Sizeof(uint32(0))*3)))
	return 1
}
func nox_xxx_playerCanTalkMB_57A160(a1 int32) int32 {
	var result int32
	if a1 != 0 && noxflags.HasGame(noxflags.GameClient) {
		result = int32((*(*uint32)(unsafe.Pointer(uintptr(a1 + 3680))) >> 3) & 1)
	} else {
		result = 0
	}
	return result
}
func nox_xxx_giant_57A190(a1 int32) int32 {
	var result int32
	if a1 != 0 {
		result = int32((*(*uint32)(unsafe.Pointer(uintptr(a1 + 3680))) >> 2) & 1)
	} else {
		result = 0
	}
	return result
}
func sub_57A1B0(a1 int16) *byte {
	var (
		v1 int32
		v2 *uint8
	)
	v1 = 0
	v2 = (*uint8)(memmap.PtrOff(0x587000, 312212))
	for uint32(int32(a1)&6128) != *(*uint32)(unsafe.Pointer(v2)) {
		v2 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), 8))
		v1++
		if int32(uintptr(unsafe.Pointer(v2))) >= int32(uintptr(memmap.PtrOff(0x587000, 312268))) {
			return nil
		}
	}
	return *(**byte)(memmap.PtrOff(0x587000, uint32(v1*8)+0x4C390))
}
func sub_57A1E0(a1 *int32, a2 *byte, a3 *int32, a4 int8, a5 int16) int8 {
	var (
		v5     *int32
		v6     int32
		v7     uint8
		v8     *byte
		v9     int32
		v10    int32
		v11    *byte
		v12    uint8
		result int8
		v14    [256]byte
		v15    [256]byte
	)
	v5 = a3
	v6 = int32(a5) & 6128
	if a3 != nil {
		sub_57ADF0(a3)
	}
	*(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*6)) = -1
	*(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*7)) = -1
	*(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*8)) = -1
	*(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*9)) = -1
	*(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*10)) = -1
	*(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*11)) = -1
	*(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*12)) = -1
	if int32(a4)&3 != 0 {
		v15[0] = 0
		libc.StrCpy(&v14[0], CString("maps\\"))
		libc.StrNCat(&v14[0], (*byte)(unsafe.Pointer(a1)), 8)
		*(*uint16)(unsafe.Pointer(&v14[libc.StrLen(&v14[0])])) = *memmap.PtrUint16(0x587000, 312376)
		if int32(a4)&2 != 0 {
			libc.StrCpy(&v15[0], &v14[0])
			if a2 != nil {
				libc.StrCat(&v15[0], a2)
			} else {
				v7 = *memmap.PtrUint8(0x587000, 312388)
				v8 = &v15[libc.StrLen(&v15[0])+1]
				v9 = int32(*memmap.PtrUint32(0x587000, 312384))
				*(*uint32)(unsafe.Pointer(func() *byte {
					p := &v8
					*p = (*byte)(unsafe.Add(unsafe.Pointer(*p), -1))
					return *p
				}())) = *memmap.PtrUint32(0x587000, 312380)
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v8))), unsafe.Sizeof(uint32(0))*1))) = uint32(v9)
				*(*byte)(unsafe.Add(unsafe.Pointer(v8), 8)) = byte(v7)
			}
			v10 = sub_57A3F0(&v15[0], int32(uintptr(unsafe.Pointer(a1))), int32(uintptr(unsafe.Pointer(a3))), v6)
			v5 = a3
		} else {
			v10 = 0
		}
		if int32(a4)&1 != 0 && v10 == 0 {
			libc.StrNCat(&v14[0], (*byte)(unsafe.Pointer(a1)), 8)
			v11 = &v14[libc.StrLen(&v14[0])+1]
			v12 = *memmap.PtrUint8(0x587000, 312396)
			*(*uint32)(unsafe.Pointer(func() *byte {
				p := &v11
				*p = (*byte)(unsafe.Add(unsafe.Pointer(*p), -1))
				return *p
			}())) = *memmap.PtrUint32(0x587000, 312392)
			*(*byte)(unsafe.Add(unsafe.Pointer(v11), 4)) = byte(v12)
			sub_57A3F0(&v14[0], int32(uintptr(unsafe.Pointer(a1))), int32(uintptr(unsafe.Pointer(v5))), v6)
		}
	}
	if dword_5d4594_2650652 != 0 && int32(a4)&4 != 0 {
		sub_57A3F0((*byte)(memmap.PtrOff(0x587000, 312400)), int32(uintptr(unsafe.Pointer(a1))), int32(uintptr(unsafe.Pointer(v5))), v6)
	}
	result = int8(a5)
	if int32(a5)&64 != 0 {
		result = int8(sub_453FA0(int32(uintptr(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*6))))), 132, 0))
	}
	return result
}
func sub_57A3F0(a1 *byte, a2 int32, a3 int32, a4 int32) int32 {
	var (
		v4 *File
		v5 *File
		v6 *byte
		v8 [256]byte
		v9 [256]libc.WChar
	)
	dword_5d4594_2523764 = 6128
	v4 = nox_fs_open_text(a1)
	v5 = v4
	if v4 == nil {
		return 0
	}
	if !nox_fs_feof(v4) {
		for {
			*(*[256]byte)(unsafe.Pointer(&v8[0])) = [256]byte{}
			nox_fs_fgets(v5, &v8[0], 256)
			v6 = libc.StrChr(&v8[0], 10)
			if v6 != nil {
				*v6 = 0
			}
			if v8[0] != 0 {
				nox_swprintf(&v9[0], libc.CWString("%S"), &v8[0])
				sub_57A4D0(&v9[0], a2, a3, a4)
			}
			if nox_fs_feof(v5) {
				break
			}
		}
	}
	nox_fs_close(v5)
	return 1
}
func sub_57A4D0(a1 *libc.WChar, a2 int32, a3 int32, a4 int32) {
	var (
		v4     uint8
		v5     int32
		result *libc.WChar
		v7     *libc.WChar
		v8     int32
		v9     *libc.WChar
		v10    *libc.WChar
		v11    uint8
		v12    [32]int32
		v13    [256]libc.WChar
	)
	v4 = 0
	v5 = 0
	v11 = 0
	sub_416580()
	result = nox_wcscpy(&v13[0], a1)
	if true {
		if v13[0] == 34 {
			result = nox_wcstok(&v13[1], libc.CWString("\"\n\r"))
			v7 = result
			v5 = 1
		} else {
			result = nox_wcstok(&v13[0], libc.CWString(" \n\t\r"))
			v7 = result
		}
		if v7 != nil {
			for {
				v8 = int32(v11)
				v11 = func() uint8 {
					p := &v4
					*p++
					return *p
				}()
				v12[v8] = int32(uintptr(unsafe.Pointer(v7)))
				v9 = (*libc.WChar)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(libc.WChar(0))*uintptr(nox_wcslen(v7)+1)))
				if v5 != 0 {
					v9 = (*libc.WChar)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(libc.WChar(0))*1))
				}
				if *v9 == 34 {
					result = nox_wcstok((*libc.WChar)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(libc.WChar(0))*1)), libc.CWString("\"\n\r"))
					v7 = result
					v5 = 1
				} else {
					result = nox_wcstok(nil, libc.CWString(" \n\t\r"))
					v7 = result
					v5 = 0
				}
				if v7 == nil {
					break
				}
			}
			if int32(v4) != 0 {
				result = (*libc.WChar)(unsafe.Pointer(uintptr(sub_57A620(v4, (**libc.WChar)(unsafe.Pointer(&v12[0])), a2, a4))))
				if result == nil {
					if a3 != 0 {
						v10 = (*libc.WChar)(alloc.Calloc(1, 524))
						nox_wcscpy((*libc.WChar)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(libc.WChar(0))*6)), a1)
						nox_common_list_append_4258E0((*nox_list_item_t)(unsafe.Pointer(uintptr(a3))), (*nox_list_item_t)(unsafe.Pointer(v10)))
					}
				}
			}
		}
	}
}
func sub_57A620(a1 uint8, a2 **libc.WChar, a3 int32, a4 int32) int32 {
	var (
		v4  **libc.WChar
		v5  uint8
		v6  int32
		v8  *uint16
		v9  int32
		v10 int32
		v11 uint8
		v12 int32
		v13 *uint16
		v14 int32
		v15 int32
		v16 [100]byte
		v17 uint8
		v18 uint8
	)
	v4 = a2
	nox_sprintf(&v16[0], CString("%S"), *a2)
	v5 = 0
	v17 = 0
	for {
		if libc.StrCmp(*(**byte)(memmap.PtrOff(0x587000, uint32(int32(v17)*8)+0x4C390)), &v16[0]) == 0 {
			dword_5d4594_2523764 = *memmap.PtrUint32(0x587000, uint32(int32(v17)*8)+0x4C394)
			return bool2int(uint32(a4) == dword_5d4594_2523764)
		}
		v17 = func() uint8 {
			p := &v5
			*p++
			return *p
		}()
		if int32(v5) >= 7 {
			break
		}
	}
	if (dword_5d4594_2523764&uint32(a4)) == 0 || _nox_wcsicmp(*v4, libc.CWString("set")) != 0 || int32(a1) <= 1 {
		return 0
	}
	if _nox_wcsicmp(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof((*libc.WChar)(nil))*1)), libc.CWString("spell")) != 0 {
		if _nox_wcsicmp(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof((*libc.WChar)(nil))*1)), libc.CWString("weapon")) != 0 {
			if _nox_wcsicmp(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof((*libc.WChar)(nil))*1)), libc.CWString("armor")) != 0 {
				return 0
			}
			if int32(a1) != 4 {
				return 0
			}
			if !noxflags.HasGame(noxflags.GameHost) {
				return 0
			}
			nox_sprintf(&v16[0], CString("%S"), *(**libc.WChar)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof((*libc.WChar)(nil))*2)))
			v13 = (*uint16)(unsafe.Pointer(uintptr(sub_415EC0(&v16[0]))))
			if v13 == nil {
				return 0
			}
			v14 = sub_415D10(int32(uintptr(unsafe.Pointer((*byte)(unsafe.Pointer(uintptr(*v13)))))))
			if v14 == 0 {
				return 0
			}
			if _nox_wcsicmp(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof((*libc.WChar)(nil))*3)), libc.CWString("off")) != 0 {
				v15 = int32(uint32(v14) | *(*uint32)(unsafe.Pointer(uintptr(a3 + 48))))
			} else {
				v15 = int32(uint32(^v14) & *(*uint32)(unsafe.Pointer(uintptr(a3 + 48))))
			}
			*(*uint32)(unsafe.Pointer(uintptr(a3 + 48))) = uint32(v15)
		} else {
			if int32(a1) != 4 {
				return 0
			}
			if !noxflags.HasGame(noxflags.GameHost) {
				return 0
			}
			nox_sprintf(&v16[0], CString("%S"), *(**libc.WChar)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof((*libc.WChar)(nil))*2)))
			v8 = (*uint16)(unsafe.Pointer(uintptr(sub_415A30(&v16[0]))))
			if v8 == nil {
				return 0
			}
			v9 = nox_xxx_ammoCheck_415880(uint16(uintptr(unsafe.Pointer((*byte)(unsafe.Pointer(uintptr(*v8)))))))
			v10 = v9
			if v9 == 0 {
				return 0
			}
			v11 = 0
			v18 = 0
			if v9 > 0 {
				for {
					v12 = v10 >> 8
					if v10>>8 > 0 {
						v10 >>= 8
					}
					v11++
					if v12 <= 0 {
						break
					}
				}
				v18 = v11
			}
			if _nox_wcsicmp(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof((*libc.WChar)(nil))*3)), libc.CWString("off")) != 0 {
				*(*uint8)(unsafe.Pointer(uintptr(int32(v18) + a3 + 43))) |= uint8(int8(v10))
			} else {
				*(*uint8)(unsafe.Pointer(uintptr(int32(v18) + a3 + 43))) &= ^uint8(int8(v10))
			}
		}
	} else {
		if int32(a1) != 4 {
			return 0
		}
		nox_sprintf(&v16[0], CString("%S"), *(**libc.WChar)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof((*libc.WChar)(nil))*2)))
		v6 = things.SpellIndex(&v16[0])
		if v6 == 0 {
			v6 = nox_xxx_spellByTitle_424960(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof((*libc.WChar)(nil))*2)))
			if v6 == 0 {
				return 0
			}
		}
		if !nox_xxx_spellIsValid_424B50(v6) {
			return 0
		}
		if nox_xxx_spellFlags_424A70(v6)&0x7000000 != 0 {
			if _nox_wcsicmp(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof((*libc.WChar)(nil))*3)), libc.CWString("off")) == 0 {
				sub_453FA0(a3+24, v6, 0)
			}
		}
	}
	if uint32(a4) == dword_5d4594_2523764 {
		return 1
	}
	return 0
}
func sub_57A950(a1 *byte) int32 {
	var v2 [256]byte
	libc.StrCpy(&v2[0], CString("maps\\"))
	libc.StrNCat(&v2[0], a1, func() int {
		if libc.StrLen(a1)-4 < 256 {
			return libc.StrLen(a1) - 4
		}
		return 256
	}())
	*(*uint16)(unsafe.Pointer(&v2[libc.StrLen(&v2[0])])) = *memmap.PtrUint16(0x587000, 312564)
	libc.StrCat(&v2[0], a1)
	return sub_4D0550(&v2[0])
}
func sub_57A9F0(a1 *byte, a2 *byte) int32 {
	var FileName [256]byte
	libc.StrCpy(&FileName[0], CString("maps\\"))
	libc.StrCat(&FileName[0], a1)
	*(*uint16)(unsafe.Pointer(&FileName[libc.StrLen(&FileName[0])])) = *memmap.PtrUint16(0x587000, 312576)
	libc.StrCat(&FileName[0], a2)
	return int32(stdio.Remove(libc.GoString(&FileName[0])))
}
func sub_57AAA0(a1 *byte, a2 *byte, a3 *int32) int8 {
	var (
		v3       *File
		v4       *File
		i        *int32
		v6       *byte
		v7       int32
		v8       int32
		v9       *byte
		v10      int32
		v11      int32
		v12      *byte
		v13      int32
		v14      int32
		v15      *byte
		v16      *byte
		v18      int32
		v19      [24]byte
		v20      [36]byte
		v21      [24]byte
		v22      [36]byte
		v23      [256]byte
		FileName [256]byte
	)
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v3))), 0)) = uint8(*(*byte)(unsafe.Add(unsafe.Pointer(a2), 52)))
	if int32(int8(uintptr(unsafe.Pointer(v3)))) >= 0 {
		libc.StrCpy(&FileName[0], CString("maps\\"))
		libc.StrNCat(&FileName[0], a2, 8)
		*(*uint16)(unsafe.Pointer(&FileName[libc.StrLen(&FileName[0])])) = *memmap.PtrUint16(0x587000, 312588)
		libc.StrCat(&FileName[0], a1)
		v3 = nox_fs_create_text(&FileName[0])
		v4 = v3
		if v3 != nil {
			if dword_5d4594_2650652 != 0 {
				libc.StrCpy(&v21[0], a2)
				libc.StrCpy(&v19[0], a2)
				sub_57A1E0((*int32)(unsafe.Pointer(&v21[0])), nil, nil, 4, int16(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(a2))), unsafe.Sizeof(uint16(0))*26)))))
				sub_57A1E0((*int32)(unsafe.Pointer(&v19[0])), nil, nil, 3, int16(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(a2))), unsafe.Sizeof(uint16(0))*26)))))
			}
			if a3 != nil {
				for i = (*int32)(unsafe.Pointer(nox_common_list_getFirstSafe_425890((*nox_list_item_t)(unsafe.Pointer(a3))))); i != nil; i = (*int32)(unsafe.Pointer(nox_common_list_getNextSafe_4258A0((*nox_list_item_t)(unsafe.Pointer(i))))) {
					nox_sprintf(&v23[0], CString("%S\n"), (*int32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(int32(0))*3)))
					nox_fs_fputs(v4, &v23[0])
				}
			}
			v6 = sub_57A1B0(int16(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(a2))), unsafe.Sizeof(uint16(0))*26)))))
			nox_fs_fputs(v4, v6)
			nox_fs_fputs(v4, CString("\n"))
			v7 = 1
			v8 = 136
			for {
				if nox_xxx_spellIsValid_424B50(v7) && sub_454000(int32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(a2), 24))))), v7) == 0 && nox_xxx_spellFlags_424A70(v7)&0x7000000 != 0 && (dword_5d4594_2650652 == 0 || sub_454000(int32(uintptr(unsafe.Pointer(&v22[0]))), v7) != 0 || sub_454000(int32(uintptr(unsafe.Pointer(&v20[0]))), v7) == 0) {
					v9 = nox_xxx_spellNameByN_424870(v7)
					nox_sprintf(&v23[0], CString("%s %s \"%s\" %s\n"), memmap.PtrOff(0x587000, 312616), memmap.PtrOff(0x587000, 312608), v9, memmap.PtrOff(0x587000, 312604))
					nox_fs_fputs(v4, &v23[0])
				}
				v7++
				v8--
				if v8 == 0 {
					break
				}
			}
			v10 = 1
			v11 = 26
			for {
				if (uint32(v10) & *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a2))), unsafe.Sizeof(uint32(0))*12)))) == 0 {
					v12 = sub_415E40((*byte)(unsafe.Pointer(uintptr(v10))))
					if v12 != nil {
						nox_sprintf(&v23[0], CString("%s %s \"%s\" %s\n"), memmap.PtrOff(0x587000, 312648), memmap.PtrOff(0x587000, 312640), v12, memmap.PtrOff(0x587000, 312636))
						nox_fs_fputs(v4, &v23[0])
					}
				}
				v10 *= 2
				v11--
				if v11 == 0 {
					break
				}
			}
			v13 = 1
			v14 = 1
			v18 = 27
			v15 = (*byte)(unsafe.Add(unsafe.Pointer(a2), 44))
			for {
				if (int32(uint8(int8(v13))) & int32(uint8(*v15))) == 0 {
					v16 = sub_4159B0((*byte)(unsafe.Pointer(uintptr(v14))))
					if v16 != nil {
						nox_sprintf(&v23[0], CString("%s %s \"%s\" %s\n"), memmap.PtrOff(0x587000, 312680), memmap.PtrOff(0x587000, 312672), v16, memmap.PtrOff(0x587000, 312668))
						nox_fs_fputs(v4, &v23[0])
					}
				}
				if v13 == 128 {
					v13 = 1
					v15 = (*byte)(unsafe.Add(unsafe.Pointer(v15), 1))
				} else {
					v13 *= 2
				}
				v14 *= 2
				v18--
				if v18 == 0 {
					break
				}
			}
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v3))), 0)) = 0
			nox_fs_close(v4)
		}
	}
	return int8(uintptr(unsafe.Pointer(v3)))
}
func sub_57ADF0(a1 *int32) *int32 {
	var (
		result *int32
		v2     *int32
		v3     *int32
	)
	result = (*int32)(unsafe.Pointer(nox_common_list_getFirstSafe_425890((*nox_list_item_t)(unsafe.Pointer(a1)))))
	v2 = result
	if result != nil {
		for {
			v3 = (*int32)(unsafe.Pointer(nox_common_list_getNextSafe_4258A0((*nox_list_item_t)(unsafe.Pointer(v2)))))
			sub_425920((**uint32)(unsafe.Pointer(v2)))
			alloc.Free(unsafe.Pointer(v2))
			v2 = v3
			if v3 == nil {
				break
			}
		}
	}
	return result
}
func sub_57AE30(a1 *byte) int32 {
	var (
		v1 int32
		v2 **byte
	)
	v1 = 0
	v2 = (**byte)(memmap.PtrOff(0x587000, 312208))
	for libc.StrCmp(*v2, a1) != 0 {
		v2 = (**byte)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof((*byte)(nil))*2))
		v1++
		if int32(uintptr(unsafe.Pointer(v2))) >= int32(uintptr(memmap.PtrOff(0x587000, 312264))) {
			return 0
		}
	}
	return int32(*memmap.PtrUint32(0x587000, uint32(v1*8)+0x4C394))
}
func nox_xxx_playerCheckSpellClass_57AEA0(a1 int32, a2 int32) int32 {
	var (
		v2     int32
		result int32
		v4     int32
	)
	v2 = int32(nox_xxx_spellFlags_424A70(a2))
	if a1 == 1 {
		v4 = 0x2000000
	} else {
		if a1 != 2 {
			return 9
		}
		v4 = 0x4000000
	}
	if uint32(v2)&0x1000000 != 0 || v4&v2 != 0 {
		result = 0
	} else {
		result = 9
	}
	return result
}
func sub_57AEE0(a1 int32, a2 int32) int32 {
	return bool2int(a1 < 75 || a1 > 114 || nox_xxx_countControlledCreatures_500D10(a2) <= 4)
}
func nox_xxx_get_57AF20() int32 {
	return int32(dword_5d4594_2523804)
}
func sub_57B0A0() {
	var (
		result int32
		v1     *uint32
	)
	result = int32(dword_5d4594_2523804)
	if result == 0 {
		return
	}
	v1 = *(**uint32)(unsafe.Pointer(&dword_5d4594_2523780))
	if dword_5d4594_2523780 != 0 && (*memmap.PtrUint32(6112660, 2523772) == 0 || *memmap.PtrUint32(6112660, 2523772) == 1) {
		nox_xxx_netSendPointFx_522FF0(-102, (*float2)(unsafe.Pointer(uintptr(dword_5d4594_2523780+56))))
		v1 = *(**uint32)(unsafe.Pointer(&dword_5d4594_2523780))
	}
	if dword_5d4594_2523776 != 0 {
		nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_2523776))))))
		v1 = *(**uint32)(unsafe.Pointer(&dword_5d4594_2523780))
	}
	dword_5d4594_2523776 = 0
	if v1 != nil {
		nox_xxx_playerSetState_4FA020((*nox_object_t)(unsafe.Pointer(v1)), 13)
	}
	dword_5d4594_2523780 = 0
	if sub_45D9B0() == 0 {
		sub_413A00(0)
	}
	dword_5d4594_2523804 = 0
}
func nox_xxx___Getcvt_57B180() int64 {
	return int64(*memmap.PtrUint64(6112660, 2523788))
}
func sub_57B190(a1 uint16, a2 uint16) int32 {
	var (
		result int32
		v3     float64
		v4     float64
	)
	if int32(a2) == 0 {
		return 4
	}
	if int32(a1) == int32(a2) {
		return 0
	}
	v3 = float64(a1)
	v4 = float64(a2)
	if v3 >= v4**(*float64)(unsafe.Pointer(&qword_581450_9544)) {
		return 1
	}
	result = 2
	if v3 < v4**mem_getDoublePtr(0x581450, 9608) {
		result = 3
	}
	return result
}
func nox_xxx_loadBaseValues_57B200() {
	*mem_getFloatPtr(6112660, 0x2682A4) = float32(nox_xxx_gamedataGetFloat_419D40(CString("BaseHealth")))
	*mem_getFloatPtr(6112660, 0x2682A8) = float32(nox_xxx_gamedataGetFloat_419D40(CString("BaseMana")))
	*mem_getFloatPtr(6112660, 0x2682B0) = float32(nox_xxx_gamedataGetFloat_419D40(CString("BaseStrength")))
	*mem_getFloatPtr(6112660, 2523820) = float32(nox_xxx_gamedataGetFloat_419D40(CString("BaseSpeed")))
	*mem_getFloatPtr(6112660, 0x2682B4) = float32(nox_xxx_gamedataGetFloat_419D40(CString("WarriorMaxHealth")) * float64(*(*float32)(unsafe.Pointer(&nox_xxx_warriorMaxHealth_587000_312784))))
	*mem_getFloatPtr(6112660, 0x2682B8) = float32(nox_xxx_gamedataGetFloat_419D40(CString("WarriorMaxMana")) * float64(*(*float32)(unsafe.Pointer(&nox_xxx_warriorMaxMana_587000_312788))))
	*mem_getFloatPtr(6112660, 0x2682C0) = float32(nox_xxx_gamedataGetFloat_419D40(CString("WarriorMaxStrength")) * float64(*(*float32)(unsafe.Pointer(&nox_xxx_warriorMaxStrength_587000_312792))))
	*mem_getFloatPtr(6112660, 0x2682BC) = float32(nox_xxx_gamedataGetFloat_419D40(CString("WarriorMaxSpeed")) * float64(*(*float32)(unsafe.Pointer(&nox_xxx_warriorMaxSpeed_587000_312796))))
	*mem_getFloatPtr(6112660, 2523860) = float32(nox_xxx_gamedataGetFloat_419D40(CString("ConjurerMaxHealth")) * float64(*(*float32)(unsafe.Pointer(&nox_xxx_conjurerMaxHealth_587000_312800))))
	*mem_getFloatPtr(6112660, 0x2682D8) = float32(nox_xxx_gamedataGetFloat_419D40(CString("ConjurerMaxMana")) * float64(*(*float32)(unsafe.Pointer(&nox_xxx_conjurerMaxMana_587000_312804))))
	*mem_getFloatPtr(6112660, 0x2682E0) = float32(nox_xxx_gamedataGetFloat_419D40(CString("ConjurerMaxStrength")) * float64(*(*float32)(unsafe.Pointer(&nox_xxx_conjurerStrength_587000_312808))))
	*mem_getFloatPtr(6112660, 0x2682DC) = float32(nox_xxx_gamedataGetFloat_419D40(CString("ConjurerMaxSpeed")) * float64(*(*float32)(unsafe.Pointer(&nox_xxx_conjurerSpeed_587000_312812))))
	*mem_getFloatPtr(6112660, 0x2682C4) = float32(nox_xxx_gamedataGetFloat_419D40(CString("WizardMaxHealth")) * float64(*(*float32)(unsafe.Pointer(&nox_xxx_wizardMaxHealth_587000_312816))))
	*mem_getFloatPtr(6112660, 0x2682C8) = float32(nox_xxx_gamedataGetFloat_419D40(CString("WizardMaxMana")) * float64(*(*float32)(unsafe.Pointer(&nox_xxx_wizardMaximumMana_587000_312820))))
	*mem_getFloatPtr(6112660, 0x2682D0) = float32(nox_xxx_gamedataGetFloat_419D40(CString("WizardMaxStrength")) * float64(*(*float32)(unsafe.Pointer(&nox_xxx_wizardStrength_587000_312824))))
	*mem_getFloatPtr(6112660, 0x2682CC) = float32(nox_xxx_gamedataGetFloat_419D40(CString("WizardMaxSpeed")) * float64(*(*float32)(unsafe.Pointer(&nox_xxx_wizardSpeed_587000_312828))))
}
func sub_57B350() *float32 {
	return mem_getFloatPtr(6112660, 0x2682A4)
}
func nox_xxx_plrGetMaxVarsPtr_57B360(a1 int32) *float32 {
	return mem_getFloatPtr(6112660, uint32(a1*16)+0x2682B4)
}
func sub_57B370(a1 int32, a2 uint8, a3 int32) int8 {
	var (
		v3     *uint32
		result int8
	)
	if (uint32(a1) & 0x3001010) == 0 {
		return -1
	}
	if uint32(a1)&0x1001000 != 0 {
		v3 = nox_xxx_getProjectileClassById_413250(a3)
		goto LABEL_4
	}
	if uint32(a1)&0x2000000 != 0 {
		v3 = nox_xxx_equipClothFindDefByTT_413270(a3)
	LABEL_4:
		if v3 != nil {
			result = int8(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v3))), 62))))
		} else {
			result = 0
		}
		return result
	}
	if a1&16 != 0 {
		result = int8(^(int32(a2) >> 5) | 254)
	} else {
		result = int8(a3)
	}
	return result
}

var nox_cheat_allowall int32 = 0

func nox_xxx_playerClassCanUseItem_57B3D0(item *nox_object_t, a2 int8) int32 {
	if nox_cheat_allowall != 0 {
		return 1
	}
	return bool2int((int32(uint8(int8(1<<int32(a2)))) & int32(uint8(sub_57B370(int32(item.obj_class), uint8(item.obj_subclass), int32(item.typ_ind))))) != 0)
}
func nox_xxx_client_57B400(a1 int32) int32 {
	var v1 int32
	v1 = int32(*memmap.PtrUint32(6112660, 2523876))
	if *memmap.PtrUint32(6112660, 2523876) == 0 {
		v1 = nox_xxx_getTTByNameSpriteMB_44CFC0("Glyph")
		*memmap.PtrUint32(6112660, 2523876) = uint32(v1)
	}
	if *memmap.PtrUint32(0x8531A0, 2576) == 0 {
		return 0
	}
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + 108))) != uint32(v1) || int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 2251)))) == 1 {
		return 1
	}
	return 0
}
func sub_57B450(a1p *nox_drawable) int32 {
	var (
		a1 *int32 = (*int32)(unsafe.Pointer(a1p))
		v1 int32
		v2 uint8
	)
	v1 = int32(*memmap.PtrUint32(6112660, 2523880))
	if *memmap.PtrUint32(6112660, 2523880) == 0 {
		v1 = nox_xxx_getTTByNameSpriteMB_44CFC0("Glyph")
		*memmap.PtrUint32(6112660, 2523880) = uint32(v1)
	}
	if a1 == nil || *memmap.PtrUint32(0x852978, 8) == 0 || *memmap.PtrUint32(0x8531A0, 2576) == 0 || *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*27)) == v1 && int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 2251)))) != 1 {
		return 0
	}
	if nox_cheat_allowall != 0 {
		return 1
	}
	v2 = uint8(int8(1 << int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 2251))))))
	return bool2int((int32(v2) & int32(uint8(sub_57B370(*(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*28)), uint8(int8(*(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*29)))), *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*27)))))) != 0)
}
func sub_57B4D0(a1 int32) int32 {
	var result int32
	result = a1
	if a1 != 0 {
		*memmap.PtrUint32(6112660, 2523884) = uint32(a1)
		dword_5d4594_2523888 = 1
	} else {
		dword_5d4594_2523888 = 0
	}
	return result
}
func sub_57B500(a1 int32, a2 int32, a3 int8) int8 {
	var (
		v3     int32
		v4     int8
		v5     int32
		v6     int32
		result int8
		v8     int32
		v9     int32
		v10    bool
		v11    bool
	)
	if a1 < 0 {
		return -1
	}
	if a1 >= 256 {
		return -1
	}
	if a2 < 0 {
		return -1
	}
	if a2 >= 256 {
		return -1
	}
	v3 = nox_xxx_wall_4105E0(a1, a2)
	if v3 == 0 {
		return -1
	}
	v4 = int8(*(*uint8)(unsafe.Pointer(uintptr(v3 + 4))))
	if int32(v4)&16 != 0 {
		if (int32(a3) & 16) == 0 {
			return -1
		}
		v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 28))))
		if v5 == 0 {
			return -1
		}
		v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 748))))
		if int32(a3)&8 != 0 {
			if int32(*(*uint8)(unsafe.Pointer(uintptr(v6 + 1)))) == 0 {
				return -1
			}
			if dword_5d4594_2523888 != 0 && nox_xxx_doorGetSomeKey_4E8910(*memmap.PtrInt32(6112660, 2523884), v5) != 0 {
				dword_5d4594_2523888 = 0
				return -1
			}
		}
		if int32(a3) >= 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(v5 + 12))))&4 != 0 {
			return -1
		}
		v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(v6 + 12))))
		if uint32(v8) != *(*uint32)(unsafe.Pointer(uintptr(v6 + 4))) {
			return -1
		}
		v9 = *memmap.PtrInt32(0x587000, uint32(v8*8)+0x2FE58)
		v10 = v9 < 0
		v11 = v9 <= 0
		if v9 > 0 {
			if *memmap.PtrInt32(0x587000, uint32(v8*8)+0x2FE5C) > 0 {
				return 1
			}
			v10 = v9 < 0
			v11 = v9 <= 0
		}
		if v10 {
			if *memmap.PtrInt32(0x587000, uint32(v8*8)+0x2FE5C) < 0 {
				return 1
			}
			v11 = v9 <= 0
			if v9 < 0 {
				if *memmap.PtrInt32(0x587000, uint32(v8*8)+0x2FE5C) > 0 {
					return 0
				}
				v11 = v9 <= 0
			}
		}
		if v11 || *memmap.PtrInt32(0x587000, uint32(v8*8)+0x2FE5C) >= 0 {
			return -1
		}
		result = 0
	} else {
		if (int32(a3)&64) == 0 && int32(v4)&64 != 0 || int32(v4)&4 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 28))) + 22)))) > 11 {
			return -1
		}
		result = int8(*(*uint8)(unsafe.Pointer(uintptr(v3))))
	}
	return result
}
func sub_57B630(a1 int32, a2 int32, a3 int32) int8 {
	var (
		v3  int32
		v4  int8
		v5  int32
		v6  int32
		v7  int32
		v8  int32
		v9  bool
		v10 bool
		v11 int8
		v13 int32
	)
	if a2 >= 0 && a2 < 256 && (a3 >= 0 || a3 < 256) {
		v3 = nox_xxx_wall_4105E0(a2, a3)
		if v3 != 0 {
			v4 = int8(*(*uint8)(unsafe.Pointer(uintptr(v3 + 4))))
			if int32(v4)&16 != 0 {
				v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 28))))
				if v5 != 0 {
					v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 748))))
					v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(v6 + 12))))
					if v7 == *(*int32)(unsafe.Pointer(uintptr(v6 + 4))) {
						v8 = *memmap.PtrInt32(0x587000, uint32(v7*8)+0x2FE58)
						v9 = v8 < 0
						v10 = v8 <= 0
						if v8 > 0 {
							if *memmap.PtrInt32(0x587000, uint32(v7*8)+0x2FE5C) > 0 {
								v11 = 1
								goto LABEL_22
							}
							v9 = v8 < 0
							v10 = v8 <= 0
						}
						if v9 {
							if *memmap.PtrInt32(0x587000, uint32(v7*8)+0x2FE5C) < 0 {
								v11 = 1
								goto LABEL_22
							}
							v10 = v8 <= 0
							if v8 < 0 {
								if *memmap.PtrInt32(0x587000, uint32(v7*8)+0x2FE5C) > 0 {
								LABEL_21:
									v11 = 0
								LABEL_22:
									if *(*uint32)(unsafe.Pointer(uintptr(v5 + 508))) != 0 {
										if noxRndCounter1.IntClamp(0, 100) >= 50 {
											return v11
										}
									} else if int32(*(*uint8)(unsafe.Pointer(uintptr(v6 + 1)))) != 0 && nox_xxx_doorGetSomeKey_4E8910(a1, v5) == 0 {
										return v11
									}
									return -1
								}
								v10 = v8 <= 0
							}
						}
						if !v10 && *memmap.PtrInt32(0x587000, uint32(v7*8)+0x2FE5C) < 0 {
							goto LABEL_21
						}
					}
				}
			} else if (*(*uint32)(unsafe.Pointer(uintptr(a1 + 16)))&0x4000) == 0 || (int32(v4)&64) == 0 {
				if (int32(v4) & 4) == 0 {
					return int8(*(*uint8)(unsafe.Pointer(uintptr(v3))))
				}
				v13 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 28))))
				if (int32(*(*uint8)(unsafe.Pointer(uintptr(v13 + 20))))&2) == 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(v13 + 22)))) <= 11 {
					return int8(*(*uint8)(unsafe.Pointer(uintptr(v3))))
				}
			}
		}
	}
	return -1
}
func sub_57B770(a1 *float2, a2 *float2) *float2 {
	var (
		result *float2
		v3     float64
		v4     float64
		v5     float64
		v6     float64
		v7     float32
		v8     float32
		v9     float32
		v10    float32
		v11    float32
	)
	result = a2
	v9 = a2.field_0
	v3 = math.Sqrt(float64(a2.field_0*a2.field_0 + a2.field_4*a2.field_4))
	v4 = v3 + 0.1
	v5 = float64(-a2.field_4)
	v6 = float64(a1.field_0*a2.field_0+a1.field_4*a2.field_4) / (v3 + 0.1)
	v11 = float32((float64(a2.field_0*a1.field_4) + v5*float64(a1.field_0)) / v4)
	v7 = float32(v6 * float64(result.field_0) / v4)
	v8 = float32(v6 * float64(result.field_4) / v4)
	v10 = float32(float64(v11) * v5 / v4)
	a1.field_0 = v10 - v7
	a1.field_4 = float32(float64(v11*v9)/v4 - float64(v8))
	return result
}
func nox_xxx_collideReflect_57B810(a1 *float32, a2 int32) int32 {
	var (
		result int32
		v3     float64
		v4     int32
	)
	result = a2
	v3 = float64(*(*float32)(unsafe.Pointer(uintptr(a2))))
	if float64(*(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*1))**a1) <= 0.0 {
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 4))))
		*(*float32)(unsafe.Pointer(uintptr(a2 + 4))) = *(*float32)(unsafe.Pointer(uintptr(a2)))
		*(*uint32)(unsafe.Pointer(uintptr(a2))) = uint32(v4)
	} else {
		*(*float32)(unsafe.Pointer(uintptr(a2))) = -*(*float32)(unsafe.Pointer(uintptr(a2 + 4)))
		*(*float32)(unsafe.Pointer(uintptr(a2 + 4))) = float32(-v3)
	}
	return result
}
func nox_xxx_map_57B850(a1 *float2, a2 *float32, a3 *float2) int32 {
	var (
		result int32
		v4     float32
		v5     float32
		v6     float32
		v7     float32
	)
	v4 = *(*float32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(float32(0))*5)) + a1.field_0
	v5 = *(*float32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(float32(0))*6)) + a1.field_4
	result = 0
	if float64(v5-v4+a3.field_0-a3.field_4)*0.70709997 < 0.0 && float64(*(*float32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(float32(0))*8))+a1.field_4-(*(*float32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(float32(0))*7))+a1.field_0)+a3.field_0-a3.field_4)*0.70709997 > 0.0 {
		v6 = *(*float32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(float32(0))*9)) + a1.field_0
		v7 = *(*float32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(float32(0))*10)) + a1.field_4
		if float64(v7+v6-a3.field_0-a3.field_4)*0.70709997 > 0.0 && float64(v5+v4-a3.field_0-a3.field_4)*0.70709997 < 0.0 {
			result = 1
		}
	}
	return result
}
func sub_57B920(a1 unsafe.Pointer) int32 {
	var result int32
	result = 0
	libc.MemSet(a1, 0, 2040)
	return result
}
func sub_57B930(a1 int32, a2 int32, a3 int32, a4 uint32) int8 {
	var (
		v4 int32
		v5 int32
	)
	v4 = int32(uint8(int8(a2)))
	v5 = int32(uint8(int8(a2)))
	if int32(uint8(int8(a2))) == math.MaxUint8 || int32(uint8(int8(a2))) == 0 {
		v4 = 1
		v5 = 1
	}
	for int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + v4*8)))) != a2 || int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + v4*8 + 2)))) != a3 {
		if func() int32 {
			p := &v4
			*p++
			return *p
		}() == math.MaxUint8 {
			v4 = 1
		}
		if v4 == v5 {
			goto LABEL_11
		}
	}
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + v4*8 + 4))) >= a4 {
		return int8(v4)
	}
LABEL_11:
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v4))), 0)) = math.MaxUint8
	return int8(v4)
}
func nox_xxx_cliGenerateAlias_57B9A0(a1 int32, a2 int32, a3 int32, a4 uint32) int8 {
	var (
		v4 int32
		v5 int32
	)
	v4 = int32(uint8(int8(a2)))
	v5 = int32(uint8(int8(a2)))
	if int32(uint8(int8(a2))) == math.MaxUint8 || int32(uint8(int8(a2))) == 0 {
		v4 = 1
		v5 = 1
	}
	for (int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + v4*8)))) != a2 || int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + v4*8 + 2)))) != a3) && *(*uint32)(unsafe.Pointer(uintptr(a1 + v4*8 + 4))) >= a4 {
		if func() int32 {
			p := &v4
			*p++
			return *p
		}() == math.MaxUint8 {
			v4 = 1
		}
		if v4 == v5 {
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v4))), 0)) = math.MaxUint8
			return int8(v4)
		}
	}
	return int8(v4)
}
func sub_57BA10(a1 int32, a2 int16, a3 int16, a4 int32) int32 {
	var result int32
	result = a1
	*(*uint16)(unsafe.Pointer(uintptr(a1))) = uint16(a2)
	*(*uint16)(unsafe.Pointer(uintptr(a1 + 2))) = uint16(a3)
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 4))) = uint32(a4)
	return result
}
func sub_57BA30(a1 *int2, a2 *int2, a3 *int4) int32 {
	var (
		v3     int32
		v4     int32
		v5     int32
		v6     int32
		v7     int32
		v8     int32
		result int32
		v10    int32
		v11    int32
		v12    int32
	)
	v12 = 0
	for {
		if a1.field_0 >= a3.field_0 {
			if a1.field_0 <= a3.field_8 {
				v3 = 0
			} else {
				v3 = 4
			}
		} else {
			v3 = 8
		}
		v4 = a1.field_4
		v5 = a3.field_4
		if v4 >= v5 {
			if v4 > a3.field_C {
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v3))), 0)) = uint8(int8(v3 | 1))
			}
		} else {
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v3))), 0)) = uint8(int8(v3 | 2))
		}
		if a2.field_0 >= a3.field_0 {
			if a2.field_0 <= a3.field_8 {
				v6 = 0
			} else {
				v6 = 4
			}
		} else {
			v6 = 8
		}
		v7 = a2.field_4
		if v7 >= v5 {
			if v7 > a3.field_C {
				v6 |= 1
			}
		} else {
			v6 |= 2
		}
		if v6&v3 != 0 {
			result = 0
			goto LABEL_30
		}
		if v3 == 0 {
			break
		}
	LABEL_20:
		if v3&1 != 0 {
			a1.field_0 += (a3.field_C - a1.field_4) * (a2.field_0 - a1.field_0) / (a2.field_4 - a1.field_4)
			a1.field_4 = a3.field_C
		} else if v3&2 != 0 {
			a1.field_0 += (a3.field_4 - a1.field_4) * (a2.field_0 - a1.field_0) / (a2.field_4 - a1.field_4)
			a1.field_4 = a3.field_4
		} else if v3&4 != 0 {
			a1.field_4 += (a3.field_8 - a1.field_0) * (a2.field_4 - a1.field_4) / (a2.field_0 - a1.field_0)
			a1.field_0 = a3.field_8
		} else if v3&8 != 0 {
			a1.field_4 += (a2.field_4 - a1.field_4) * (a3.field_0 - a1.field_0) / (a2.field_0 - a1.field_0)
			a1.field_0 = a3.field_0
		}
	}
	if v6 != 0 {
		v3 = a1.field_0
		v8 = a1.field_4
		*a1 = *a2
		a2.field_0 = v3
		a2.field_4 = v8
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v3))), 0)) = uint8(int8(v6))
		v12 = bool2int(v12 == 0)
		goto LABEL_20
	}
	result = 1
LABEL_30:
	if v12 != 0 {
		v10 = a1.field_0
		v11 = a1.field_4
		*a1 = *a2
		a2.field_0 = v10
		a2.field_4 = v11
	}
	return result
}
func nox_xxx_mapNxzDecompress_57BC50(a1 *byte, a2 *byte) int32 {
	var (
		v2  *File
		v3  *File
		v4  int32
		v5  *byte
		v6  *byte
		v7  *byte
		v8  *byte
		v9  *uint32
		v10 *File
		v12 uint32
		v13 uint32
	)
	v12 = 0
	if a1 == nil {
		return 0
	}
	if a2 == nil {
		return 0
	}
	v2 = nox_fs_open(a1)
	v3 = v2
	if v2 == nil {
		return 0
	}
	v4 = nox_fs_fsize(v3)
	v5 = (*byte)(unsafe.Pointer(uintptr(v4 - 4)))
	nox_binfile_fread_raw_40ADD0((*byte)(unsafe.Pointer(&v12)), 1, 4, v3)
	v6 = (*byte)(alloc.Calloc(1, int(uint32(uintptr(unsafe.Pointer(v5))))))
	v7 = (*byte)(alloc.Calloc(1, int(v12)))
	v8 = v7
	if v6 == nil || v7 == nil {
		return 0
	}
	nox_binfile_fread_raw_40ADD0(v6, 1, uint32(uintptr(unsafe.Pointer(v5))), v3)
	nox_fs_close(v3)
	a1 = v5
	v13 = v12
	v9 = sub_578BF0()
	if nox_xxx_nxzDecompress_578C10(v9, (*uint8)(unsafe.Add(unsafe.Pointer(v8), v12-v13)), &v13, uint32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v5), int32(uintptr(unsafe.Pointer(v6))-uintptr(unsafe.Pointer(a1)))))))), (*uint32)(unsafe.Pointer(&a1))) != 0 {
		for int32(uintptr(unsafe.Pointer(a1))) > 0 && nox_xxx_nxzDecompress_578C10(v9, (*uint8)(unsafe.Add(unsafe.Pointer(v8), v12-v13)), &v13, uint32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v5), int32(uintptr(unsafe.Pointer(v6))-uintptr(unsafe.Pointer(a1)))))))), (*uint32)(unsafe.Pointer(&a1))) != 0 {
		}
	}
	sub_578C40(unsafe.Pointer(v9))
	v10 = nox_fs_create(a2)
	if v10 == nil {
		return 0
	}
	nox_fs_fwrite(v10, unsafe.Pointer(v8), int32(v12))
	nox_fs_close(v10)
	alloc.Free(unsafe.Pointer(v6))
	alloc.Free(unsafe.Pointer(v8))
	return 1
}
func nox_xxx_mapFile_57BDD0(lpMem unsafe.Pointer, a2 int32) int32 {
	var (
		v2     uint32
		v3     *File
		v4     *File
		v5     *byte
		v6     uint32
		v7     unsafe.Pointer
		v8     *unsafe.Pointer
		v9     uint32
		i      uint32
		v11    uint32
		v12    *File
		v13    *File
		v15    uint32
		lpMema unsafe.Pointer
	)
	v2 = 0
	v15 = 0
	if lpMem == nil {
		return 0
	}
	if a2 == 0 {
		return 0
	}
	v3 = nox_fs_open((*byte)(lpMem))
	v4 = v3
	if v3 == nil {
		return 0
	}
	v15 = uint32(nox_fs_fsize(v4))
	v5 = (*byte)(alloc.Calloc(1, int(v15)))
	v6 = sub_578BA0(v15)
	v7 = alloc.Calloc(1, int(v6))
	lpMema = v7
	if v5 == nil || v7 == nil {
		return 0
	}
	nox_binfile_fread_raw_40ADD0(v5, 1, v15, v4)
	nox_fs_close(v4)
	v8 = (*unsafe.Pointer)(unsafe.Pointer(sub_578B80()))
	v9 = v15
	for i = 0; i < v15; i += 500000 {
		v11 = v9 - i
		if v11 > 500000 {
			v11 = 500000
		}
		v2 += uint32(sub_578BB0(v8, int32(uint32(int32(uintptr(lpMema)))+v2), (*uint8)(unsafe.Add(unsafe.Pointer(v5), i)), int32(v11)))
		v9 = v15
	}
	sub_578BD0(unsafe.Pointer(v8))
	v12 = nox_fs_create((*byte)(unsafe.Pointer(uintptr(a2))))
	v13 = v12
	if v12 == nil {
		return 0
	}
	nox_fs_fwrite(v12, unsafe.Pointer(&v15), 4)
	nox_fs_fwrite(v13, lpMema, int32(v2))
	nox_fs_close(v13)
	alloc.Free(unsafe.Pointer(v5))
	lpMema = nil
	return 1
}
func sub_57BF20(a1 int32, a2 int32) int32 {
	return bool2int(a1 != 0 && a2 != 0)
}
func sub_57BF40(a1 int32, a2 int32) int32 {
	return bool2int(a1 != 0 && a2 != 0)
}
func sub_57BF60() int32 {
	return 0
}
func sub_57BF70() int32 {
	return int32(dword_5d4594_2523904)
}
func sub_57BF80() uint32 {
	var (
		v0 uint32
		i  int32
		v2 uint32
	)
	v0 = 0
	for i = nox_server_getFirstMapGroup_57C080(); i != 0; i = nox_server_getNextMapGroup_57C090(i) {
		v2 = *(*uint32)(unsafe.Pointer(uintptr(i + 4)))
		if v2 >= v0 {
			v0 = v2 + 1
		}
	}
	return v0
}
func nox_xxx_allocGroupRelatedArrays_57BFB0() int32 {
	dword_5d4594_2523904 = 0
	var result *byte = (*byte)(unsafe.Pointer(nox_new_alloc_class(CString("ItemGroupInfo"), 96, 512)))
	nox_alloc_groupInfo_2523892 = unsafe.Pointer(result)
	if result == nil {
		return 0
	}
	nox_alloc_itemGroupElem_2523896 = unsafe.Pointer(nox_new_alloc_class(CString("ItemGroupElement"), 16, 5000))
	return bool2int(nox_alloc_itemGroupElem_2523896 != nil)
}
func sub_57C000() {
	dword_5d4594_2523904 = 0
	nox_alloc_class_free_all((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_itemGroupElem_2523896)))))
	nox_alloc_class_free_all((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_groupInfo_2523892)))))
	nox_server_mapGroupsHead_2523900 = 0
}
func sub_57C030() int32 {
	dword_5d4594_2523904 = 0
	if nox_alloc_groupInfo_2523892 != nil {
		nox_free_alloc_class((*nox_alloc_class)(nox_alloc_groupInfo_2523892))
		nox_alloc_groupInfo_2523892 = nil
	}
	if nox_alloc_itemGroupElem_2523896 != nil {
		nox_free_alloc_class((*nox_alloc_class)(nox_alloc_itemGroupElem_2523896))
		nox_alloc_itemGroupElem_2523896 = nil
	}
	nox_server_mapGroupsHead_2523900 = 0
	return 1
}
func nox_server_getFirstMapGroup_57C080() int32 {
	return int32(nox_server_mapGroupsHead_2523900)
}
func nox_server_getNextMapGroup_57C090(a1 int32) int32 {
	var result int32
	if a1 != 0 {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 88))))
	} else {
		result = 0
	}
	return result
}
func nox_server_scriptGetGroup_57C0A0(a1 int32) int32 {
	var result int32
	result = int32(nox_server_mapGroupsHead_2523900)
	if nox_server_mapGroupsHead_2523900 == 0 {
		return 0
	}
	for *(*uint32)(unsafe.Pointer(uintptr(result + 4))) != uint32(a1) {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(result + 88))))
		if result == 0 {
			return 0
		}
	}
	return result
}
func nox_server_mapLoadAddGroup_57C0C0(a1 *byte, a2 int32, a3 int8) int32 {
	var (
		result int32
		v4     int32
		v5     int32
	)
	result = int32(uintptr(sub_57C330()))
	v4 = result
	if result != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(result + 4))) = uint32(a2)
		*(*uint8)(unsafe.Pointer(uintptr(result))) = uint8(a3)
		libc.StrNCpy((*byte)(unsafe.Pointer(uintptr(result+8))), a1, 76)
		*(*uint8)(unsafe.Pointer(uintptr(v4 + 83))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(v4 + 84))) = 0
		v5 = int32(nox_server_mapGroupsHead_2523900)
		*(*uint32)(unsafe.Pointer(uintptr(v4 + 92))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(v4 + 88))) = uint32(v5)
		if nox_server_mapGroupsHead_2523900 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(nox_server_mapGroupsHead_2523900 + 92))) = uint32(v4)
		}
		nox_server_mapGroupsHead_2523900 = uint32(v4)
		result = 1
	}
	return result
}
func sub_57C130(a1 *uint32, a2 int32) int32 {
	var (
		result int32
		v3     *byte
		v4     int8
		v5     int32
		v6     int32
	)
	if a1 == nil {
		return 0
	}
	v3 = *(**byte)(unsafe.Pointer(&nox_server_mapGroupsHead_2523900))
	if nox_server_mapGroupsHead_2523900 == 0 {
		return 0
	}
	for *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*1))) != uint32(a2) {
		v3 = (*byte)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*22))))))
		if v3 == nil {
			return 0
		}
	}
	if v3 == nil {
		return 0
	}
	v4 = int8(*v3)
	result = int32(uintptr(sub_57C360()))
	if result != 0 {
		if int32(v4) != 0 && int32(v4) != 1 && int32(v4) != 3 {
			if int32(v4) != 2 {
				nox_alloc_class_free_obj_first((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_itemGroupElem_2523896)))), unsafe.Pointer(uintptr(result)))
				return 0
			}
			*(*uint32)(unsafe.Pointer(uintptr(result))) = *a1
			*(*uint32)(unsafe.Pointer(uintptr(result + 4))) = *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1))
		} else {
			*(*uint32)(unsafe.Pointer(uintptr(result))) = *a1
		}
		v5 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*21))))
		*(*uint32)(unsafe.Pointer(uintptr(result + 12))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(result + 8))) = uint32(v5)
		v6 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*21))))
		if v6 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v6 + 12))) = uint32(result)
		}
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*21))) = uint32(result)
		result = 1
	}
	return result
}
func sub_57C1E0(a1 *uint32, a2 int32) int32 {
	var (
		v2 int32
		v3 *byte
		i  int32
		v5 int8
		v7 int32
	)
	v2 = 0
	v3 = *(**byte)(unsafe.Pointer(&nox_server_mapGroupsHead_2523900))
	if nox_server_mapGroupsHead_2523900 == 0 {
		return 0
	}
	for *(*uint32)(unsafe.Pointer(uintptr(nox_server_mapGroupsHead_2523900 + 4))) != uint32(a2) {
	}
	for i = int32(*(*uint32)(unsafe.Pointer(uintptr(nox_server_mapGroupsHead_2523900 + 84)))); i != 0; i = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 8)))) {
		v5 = int8(*v3)
		if *v3 != 0 && int32(v5) != 1 && int32(v5) != 3 {
			if int32(v5) == 2 && *(*uint32)(unsafe.Pointer(uintptr(i))) == *a1 && *(*uint32)(unsafe.Pointer(uintptr(i + 4))) == *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)) {
				v2 = 1
			LABEL_15:
				if *(*uint32)(unsafe.Pointer(uintptr(i + 12))) != 0 {
					*(*uint32)(unsafe.Pointer(uintptr(i + 12))) = *(*uint32)(unsafe.Pointer(uintptr(i + 8)))
				}
				v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 8))))
				if v7 != 0 {
					*(*uint32)(unsafe.Pointer(uintptr(v7 + 12))) = *(*uint32)(unsafe.Pointer(uintptr(i + 12)))
				}
				sub_57C390((*uint64)(unsafe.Pointer(uintptr(i))))
				continue
			}
		} else if *(*uint32)(unsafe.Pointer(uintptr(i))) == *a1 {
			v2 = 1
			goto LABEL_15
		}
		if v2 == 1 {
			goto LABEL_15
		}
	}
	return v2
}
func nox_server_scriptGetMapGroupByName_57C280(a1 *byte, a2 int32) unsafe.Pointer {
	var i int32
	for i = nox_server_getFirstMapGroup_57C080(); i != 0; i = nox_server_getNextMapGroup_57C090(i) {
		if a2 == nox_server_scriptGetGroupId_57C2D0((**int32)(unsafe.Pointer(uintptr(i)))) && nox_server_strcmpWithoutMapname_4DA3F0((*byte)(unsafe.Pointer(uintptr(i+8))), a1) != 0 {
			break
		}
	}
	return unsafe.Pointer(uintptr(i))
}
func sub_57C330() unsafe.Pointer {
	var (
		v0     unsafe.Pointer
		result unsafe.Pointer
	)
	v0 = nil
	if !noxflags.HasGame(noxflags.GameHost|noxflags.GameFlag22) || (func() bool {
		result = nox_alloc_class_new_obj_zero((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_groupInfo_2523892)))))
		return (func() unsafe.Pointer {
			v0 = result
			return v0
		}()) != nil
	}()) {
		result = v0
	}
	return result
}
func sub_57C360() unsafe.Pointer {
	return nox_alloc_class_new_obj_zero((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_itemGroupElem_2523896)))))
}
func sub_57C370(a1 *uint64) {
	nox_alloc_class_free_obj_first((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_groupInfo_2523892)))), unsafe.Pointer(a1))
}
func sub_57C390(a1 *uint64) {
	nox_alloc_class_free_obj_first((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_itemGroupElem_2523896)))), unsafe.Pointer(a1))
}
func nox_server_addNewMapGroup_57C3B0(a1 int32) int32 {
	var result int32
	result = a1
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 92))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 88))) = nox_server_mapGroupsHead_2523900
	if nox_server_mapGroupsHead_2523900 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(nox_server_mapGroupsHead_2523900 + 92))) = uint32(a1)
	}
	nox_server_mapGroupsHead_2523900 = uint32(a1)
	return result
}
func nox_xxx_getDebugData_57C3E0() unsafe.Pointer {
	return dword_5d4594_2523912
}
func nox_xxx_nextDebugObject_57C3F0(a1 unsafe.Pointer) unsafe.Pointer {
	var result int32
	if a1 != nil {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(a1)) + 336))))
	} else {
		result = 0
	}
	return unsafe.Pointer(uintptr(result))
}
func nox_xxx_allocDebugDataArray_57C410() int32 {
	nox_alloc_debugData_2523908 = unsafe.Pointer(nox_new_alloc_class(CString("DebugData"), 344, 256))
	return bool2int(nox_alloc_debugData_2523908 != nil)
}
func sub_57C440() {
	nox_alloc_class_free_all((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_debugData_2523908)))))
	dword_5d4594_2523912 = nil
}
func sub_57C460() int32 {
	if nox_alloc_debugData_2523908 != nil {
		nox_free_alloc_class((*nox_alloc_class)(nox_alloc_debugData_2523908))
		nox_alloc_debugData_2523908 = nil
	}
	dword_5d4594_2523912 = nil
	return 1
}
func sub_57C490(a1 *byte) int32 {
	var v1 int32
	v1 = int32(uintptr(dword_5d4594_2523912))
	if dword_5d4594_2523912 == nil {
		return 0
	}
	for libc.StrCmp(a1, (*byte)(unsafe.Pointer(uintptr(v1)))) != 0 {
		v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 336))))
		if v1 == 0 {
			return 0
		}
	}
	return v1 + 80
}
func sub_57C500(a1 *byte, a2 *byte) int32 {
	var v2 *byte
	v2 = (*byte)(nox_alloc_class_new_obj_zero((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_debugData_2523908))))))
	if v2 == nil {
		return 0
	}
	libc.StrCpy(v2, a1)
	libc.StrCpy((*byte)(unsafe.Add(unsafe.Pointer(v2), 80)), a2)
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*85))) = 0
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*84))) = uint32(uintptr(dword_5d4594_2523912))
	if dword_5d4594_2523912 != nil {
		*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_5d4594_2523912)) + 340))) = uint32(uintptr(unsafe.Pointer(v2)))
	}
	dword_5d4594_2523912 = unsafe.Pointer(v2)
	return 1
}
func sub_57C5A0(a1 *byte) {
	var (
		v1 int32
		v2 int32
		v3 int32
	)
	v1 = int32(uintptr(dword_5d4594_2523912))
	if dword_5d4594_2523912 != nil {
		for libc.StrCmp(a1, (*byte)(unsafe.Pointer(uintptr(v1)))) != 0 {
			v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 336))))
			if v1 == 0 {
				return
			}
		}
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 340))))
		if v2 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v2 + 336))) = *(*uint32)(unsafe.Pointer(uintptr(v1 + 336)))
		} else {
			dword_5d4594_2523912 = unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 336)))))
		}
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 336))))
		if v3 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v3 + 340))) = *(*uint32)(unsafe.Pointer(uintptr(v1 + 340)))
		}
		nox_alloc_class_free_obj_first((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_debugData_2523908)))), unsafe.Pointer(uintptr(v1)))
	}
}
func sub_57C650(a1 *float32, a2 *float32, a3 *float32) {
	var (
		v3  *float32
		v4  float64
		v5  float64
		v6  float64
		v7  *float32
		v8  float64
		v9  float64
		v10 float32
		v11 float32
		v12 float32
		v13 float32
		v14 float32
		v15 float32
		v16 float32
	)
	v3 = a1
	v4 = float64(*(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*2)) - *a1)
	v13 = float32(v4)
	v5 = float64(*(*float32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(float32(0))*3)) - *(*float32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(float32(0))*1)))
	v6 = math.Sqrt(v5*v5 + v4*v4)
	if v6 < 0.0099999998 {
		v6 = 0.0099999998
	}
	v12 = float32(v5)
	v7 = a3
	v10 = float32(v5)
	v8 = float64(v10*(*(*float32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(float32(0))*1))-*(*float32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(float32(0))*1)))+v13*(*a2-*v3)) / (v6 * v6)
	v11 = float32(v8 * float64(v13))
	*a3 = v11 + *v3
	*(*float32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(float32(0))*1)) = float32(v8*float64(v12) + float64(*(*float32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(float32(0))*1))))
	if float64(*v3) >= float64(*(*float32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(float32(0))*2))) {
		v15 = *v3
		v14 = *(*float32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(float32(0))*2))
	} else {
		v14 = *v3
		v15 = *(*float32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(float32(0))*2))
	}
	if float64(*(*float32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(float32(0))*1))) >= float64(*(*float32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(float32(0))*3))) {
		v9 = float64(*(*float32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(float32(0))*3)))
		v16 = *(*float32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(float32(0))*1))
	} else {
		v9 = float64(*(*float32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(float32(0))*1)))
		v16 = *(*float32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(float32(0))*3))
	}
	if float64(*v7) >= float64(v14) {
		if float64(*v7) > float64(v15) {
			*v7 = v15
		}
	} else {
		*v7 = v14
	}
	if v9 <= float64(*(*float32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(float32(0))*1))) {
		if float64(*(*float32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(float32(0))*1))) > float64(v16) {
			*(*float32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(float32(0))*1)) = v16
		}
	} else {
		*(*float32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(float32(0))*1)) = float32(v9)
	}
}
func sub_57C790(a1 *float4, a2 *float2, a3 *float2, a4 float32) {
	var (
		v4  float64
		v5  float64
		v6  float64
		v7  float64
		v8  float32
		v9  float32
		v10 float32
		v11 float32
		v12 float32
		v13 float32
	)
	v8 = a1.field_8 - a1.field_0
	v4 = float64(a1.field_C - a1.field_4)
	v9 = float32(v4)
	v5 = v4*float64(a2.field_4-a1.field_4) + float64(v8*(a2.field_0-a1.field_0))
	v6 = float64(a4 * a4)
	v10 = float32(v5 * float64(v8) / v6)
	a3.field_0 = v10 + a1.field_0
	a3.field_4 = float32(v5*float64(v9)/v6 + float64(a1.field_4))
	if float64(a1.field_0) >= float64(a1.field_8) {
		v12 = a1.field_0
		v13 = a1.field_8
	} else {
		v13 = a1.field_0
		v12 = a1.field_8
	}
	if float64(a1.field_4) >= float64(a1.field_C) {
		v7 = float64(a1.field_C)
		v11 = a1.field_4
	} else {
		v7 = float64(a1.field_4)
		v11 = a1.field_C
	}
	if float64(a3.field_0) >= float64(v13) {
		if float64(a3.field_0) > float64(v12) {
			a3.field_0 = v12
		}
	} else {
		a3.field_0 = v13
	}
	if v7 <= float64(a3.field_4) {
		if float64(a3.field_4) > float64(v11) {
			a3.field_4 = v11
		}
	} else {
		a3.field_4 = float32(v7)
	}
}
func nox_xxx_mathPointOnTheLine_57C8A0(a1 *float4, a2 *float2, a3 *float2) int32 {
	var (
		v3  *float4
		v4  *float2
		v5  float64
		v6  float64
		v7  float64
		v8  float64
		v10 float32
		v11 float32
		v12 float32
		v13 float32
		v14 float32
		v15 float32
		v16 float32
		v17 float32
	)
	v3 = a1
	v4 = a3
	v5 = float64(a1.field_8 - a1.field_0)
	v6 = float64(a1.field_C - a1.field_4)
	v12 = float32(v6)
	v10 = float32(v6*float64(v12) + v5*v5)
	v7 = float64((a2.field_4-a1.field_4)*v12) + float64(a2.field_0-a1.field_0)*v5
	v14 = float32(v7)
	v13 = v14 * v12 / v10
	a3.field_0 = float32(v5*v7/float64(v10) + float64(v3.field_0))
	v15 = v13 + v3.field_4
	a3.field_4 = v15
	if float64(v3.field_0) >= float64(v3.field_8) {
		v8 = float64(v3.field_8)
		v16 = v3.field_0
	} else {
		v8 = float64(v3.field_0)
		v16 = v3.field_8
	}
	if float64(v3.field_4) >= float64(v3.field_C) {
		v11 = v3.field_4
		v17 = v3.field_C
	} else {
		v17 = v3.field_4
		v11 = v3.field_C
	}
	return bool2int(v8 <= float64(v4.field_0) && float64(v4.field_0) <= float64(v16) && float64(v15) >= float64(v17) && float64(v15) <= float64(v11))
}
func nox_xxx_mapTraceRayImpl_57C9A0(a1 int32, a2 int32, a3 *float32, a4 *float32, a5 int8) int8 {
	var (
		v5  int32
		v6  int8
		v7  float64
		v8  float64
		v9  int32
		v10 int32
		v11 float32
		v12 float32
		v13 float32
		v14 float32
		v15 int8
		v16 float64
		v17 float64
		v18 float32
		v19 float32
		v20 float32
		a3a float2
		v23 float2
		a2a float2
		v25 float2
		v26 float32
		v27 float32
		v28 float32
		v29 float32
		v30 int8
		v31 uint8
		v32 uint8
	)
	v5 = a1
	if a1 < 0 || a1 >= 256 || a2 < 0 || a2 >= 256 {
		return 0
	}
	v6 = a5
	if int32(a5)&8 != 0 {
		v7 = float64(a1)*float64(*mem_getFloatPtr(0x587000, 0x4C8C0)) + 11.5 - float64(*a3)
		v8 = float64(a2)*float64(*mem_getFloatPtr(0x587000, 0x4C8C0)) + 11.5 - float64(*(*float32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(float32(0))*1)))
		if v8*v8+v7*v7 < 3600.0 {
			v6 = int8(int32(a5) | 4)
		}
	}
	v30 = int8((int32(^v6) & 4) * 16)
	if int32(v6)&1 != 0 {
		v30 = int8(((int32(^v6) & 4) * 16) | 16)
	}
	v31 = uint8(sub_57B500(v5, a2, v30))
	if int32(v31) == math.MaxUint8 {
		return 0
	}
	if int32(v6)&1 != 0 {
		v9 = nox_xxx_wall_4105E0(v5, a2)
	} else {
		v9 = int32(uintptr(nox_server_getWallAtGrid_410580(v5, a2)))
	}
	if v9 == 0 || int32(v6) < 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(v9 + 4))))&4 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v9 + 28))) + 20))))&2 != 0 {
		return 0
	}
	v10 = int32(*memmap.PtrUint32(0x85B3FC, uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(v9 + 1))))*0x302C+0xA844)))
	if v10&2 != 0 || int32(v6)&64 != 0 && (v10&1) == 0 {
		return 0
	}
	if float64(*a3) >= float64(*(*float32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(float32(0))*2))) {
		v12 = *(*float32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(float32(0))*2))
		v28 = *a3
		v26 = v12
	} else {
		v11 = *(*float32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(float32(0))*2))
		v26 = *a3
		v28 = v11
	}
	if float64(*(*float32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(float32(0))*1))) >= float64(*(*float32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(float32(0))*3))) {
		v14 = *(*float32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(float32(0))*3))
		v29 = *(*float32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(float32(0))*1))
		v27 = v14
	} else {
		v13 = *(*float32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(float32(0))*3))
		v27 = *(*float32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(float32(0))*1))
		v29 = v13
	}
	a2a.field_0 = float32(float64(v5 * 23))
	a2a.field_4 = float32(float64(a2 * 23))
	if int32(v31) != 0 {
		sub_57CD30((*float4)(unsafe.Pointer(a3)), &a2a, &a3a)
	}
	if int32(v31) != 1 {
		v25.field_0 = float32(float64(a2a.field_0) + 23.0)
		v25.field_4 = a2a.field_4
		sub_57CD70((*float4)(unsafe.Pointer(a3)), &v25, &v23)
	}
	v15 = 0
	v32 = 0
	if int32(*memmap.PtrUint8(0x587000, uint32(int32(v31)*24)+0x4C7B8)) != 0 && a2a.field_0+*mem_getFloatPtr(0x587000, uint32(int32(v31)*24)+0x4C7BC) <= a3a.field_0 && a2a.field_0+*mem_getFloatPtr(0x587000, uint32(int32(v31)*24)+0x4C7C0) >= a3a.field_0 && float64(a3a.field_0) >= float64(v26) && float64(a3a.field_0) <= float64(v28) && float64(a3a.field_4) >= float64(v27) && float64(a3a.field_4) <= float64(v29) {
		v15 = 1
		*(*float2)(unsafe.Pointer(a4)) = a3a
		v32 = 1
	}
	if int32(*memmap.PtrUint8(0x587000, uint32(int32(v31)*24)+0x4C7C4)) != 0 && a2a.field_0+*mem_getFloatPtr(0x587000, uint32(int32(v31)*24)+0x4C7C8) <= v23.field_0 && a2a.field_0+*mem_getFloatPtr(0x587000, uint32(int32(v31)*24)+0x4C7CC) >= v23.field_0 && float64(v23.field_0) >= float64(v26) && float64(v23.field_0) <= float64(v28) && float64(v23.field_4) >= float64(v27) && float64(v23.field_4) <= float64(v29) {
		v15++
		*(*float2)(unsafe.Pointer((*float32)(unsafe.Add(unsafe.Pointer(a4), unsafe.Sizeof(float32(0))*uintptr(int32(v32)*2))))) = v23
	}
	if int32(v15) == 2 {
		v16 = float64(*(*float32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(float32(0))*1)) - a3a.field_4)
		v17 = float64(*(*float32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(float32(0))*1)) - v23.field_4)
		if v17*v17+float64((*a3-v23.field_0)*(*a3-v23.field_0)) < v16*v16+float64((*a3-a3a.field_0)*(*a3-a3a.field_0)) {
			v18 = *a4
			v19 = *(*float32)(unsafe.Add(unsafe.Pointer(a4), unsafe.Sizeof(float32(0))*1))
			*a4 = *(*float32)(unsafe.Add(unsafe.Pointer(a4), unsafe.Sizeof(float32(0))*2))
			v20 = *(*float32)(unsafe.Add(unsafe.Pointer(a4), unsafe.Sizeof(float32(0))*3))
			*(*float32)(unsafe.Add(unsafe.Pointer(a4), unsafe.Sizeof(float32(0))*2)) = v18
			*(*float32)(unsafe.Add(unsafe.Pointer(a4), unsafe.Sizeof(float32(0))*1)) = v20
			*(*float32)(unsafe.Add(unsafe.Pointer(a4), unsafe.Sizeof(float32(0))*3)) = v19
		}
	}
	return v15
}
func sub_57CD30(a1 *float4, a2 *float2, a3 *float2) *float2 {
	var (
		v3     float64
		v4     float64
		result *float2
		v6     float64
	)
	v3 = float64(a1.field_8 - a1.field_0)
	v4 = float64(a1.field_C - a1.field_4)
	result = a3
	v6 = (float64(a2.field_4-a1.field_4)*v3 - float64(a2.field_0-a1.field_0)*v4) / (v4 - v3)
	a3.field_0 = float32(v6 + float64(a2.field_0))
	a3.field_4 = float32(v6 + float64(a2.field_4))
	return result
}
func sub_57CD70(a1 *float4, a2 *float2, a3 *float2) *float2 {
	var (
		v3     float64
		v4     float64
		result *float2
		v6     float64
	)
	v3 = float64(a1.field_8 - a1.field_0)
	v4 = float64(a1.field_C - a1.field_4)
	result = a3
	v6 = (float64(a2.field_4-a1.field_4)*v3 - float64(a2.field_0-a1.field_0)*v4) / (-v4 - v3)
	a3.field_0 = float32(float64(a2.field_0) - v6)
	a3.field_4 = float32(v6 + float64(a2.field_4))
	return result
}
func sub_57CDB0(a1 *int2, a2 *float32, a3 *float2) int32 {
	var (
		v3     *int2
		v4     int8
		result int32
		v6     *float2
		v7     *float2
		v8     *float2
		v9     int8
	)
	v3 = a1
	v9 = int8(sub_57F2A0((*float2)(unsafe.Pointer(a2)), a1.field_0, a1.field_4))
	v4 = sub_57F1D0((*float2)(unsafe.Add(unsafe.Pointer((*float2)(unsafe.Pointer(a2))), unsafe.Sizeof(float2{})*1)))
	switch sub_57B500(v3.field_0, v3.field_4, 64) {
	case 0:
		if int32(v9) != 1 && int32(v9) != 0 {
			goto LABEL_44
		}
		a3.field_0 = -0.70709997
		a3.field_4 = -0.70709997
		return 1
	case 1:
		if int32(v9) == 1 || int32(v9) == 2 {
			goto LABEL_43
		}
		a3.field_0 = -0.70709997
		a3.field_4 = 0.70709997
		return 1
	case 2:
		switch v9 {
		case 0:
			goto LABEL_40
		case 1:
			v6 = a3
			if (int32(v4) & 1) == 0 {
				goto LABEL_23
			}
			a3.field_0 = -0.70709997
			a3.field_4 = -0.70709997
			result = 1
		case 2:
			goto LABEL_30
		case 3:
			goto LABEL_35
		default:
			return 1
		}
		return result
	case 3:
		if int32(v9) == 0 {
			goto LABEL_40
		}
		if int32(v9) != 1 {
			goto LABEL_44
		}
		goto LABEL_25
	case 4:
		if int32(v9) == 1 {
			goto LABEL_25
		}
		if int32(v9) == 2 {
			goto LABEL_30
		}
		a3.field_0 = -0.70709997
		a3.field_4 = 0.70709997
		return 1
	case 5:
		if int32(v9) == 2 {
			goto LABEL_30
		}
		if int32(v9) == 3 {
			goto LABEL_35
		}
		a3.field_0 = -0.70709997
		a3.field_4 = -0.70709997
		return 1
	case 6:
		if int32(v9) == 0 {
			goto LABEL_40
		}
		if int32(v9) == 3 {
			goto LABEL_35
		}
		v6 = a3
	LABEL_23:
		v6.field_0 = 0.70709997
		v6.field_4 = -0.70709997
		return 1
	case 7:
		if int32(v9) == 1 {
		LABEL_25:
			v7 = a3
			if (int32(v4) & 1) == 0 {
				goto LABEL_38
			}
			goto LABEL_26
		}
		v8 = a3
		if int32(v4)&1 != 0 {
			goto LABEL_45
		}
		a3.field_0 = -0.70709997
		a3.field_4 = 0.70709997
		return 1
	case 8:
		if int32(v9) == 2 {
		LABEL_30:
			v8 = a3
			a3.field_0 = 0.70709997
			if int32(v4)&1 != 0 {
				goto LABEL_46
			}
			a3.field_4 = -0.70709997
			result = 1
		} else {
			v8 = a3
			a3.field_0 = -0.70709997
			if (int32(v4) & 1) == 0 {
				goto LABEL_46
			}
			a3.field_4 = -0.70709997
			result = 1
		}
		return result
	case 9:
		if int32(v9) == 3 {
		LABEL_35:
			v8 = a3
			if (int32(v4) & 4) == 0 {
				goto LABEL_45
			}
			a3.field_0 = -0.70709997
			a3.field_4 = 0.70709997
			result = 1
		} else {
			v7 = a3
			if int32(v4)&4 != 0 {
			LABEL_38:
				v7.field_0 = 0.70709997
				v7.field_4 = -0.70709997
				result = 1
			} else {
			LABEL_26:
				v7.field_0 = -0.70709997
				v7.field_4 = -0.70709997
				result = 1
			}
		}
		return result
	case 10:
		if int32(v9) != 0 {
			if int32(v4)&2 != 0 {
			LABEL_43:
				a3.field_0 = 0.70709997
				a3.field_4 = -0.70709997
				return 1
			}
		LABEL_44:
			v8 = a3
		LABEL_45:
			v8.field_0 = 0.70709997
		} else {
		LABEL_40:
			v8 = a3
			a3.field_0 = -0.70709997
			if (int32(v4) & 2) == 0 {
				a3.field_4 = -0.70709997
				return 1
			}
		}
	LABEL_46:
		v8.field_4 = 0.70709997
		return 1
	default:
		return 0
	}
}
func sub_57D150(this *unsafe.Pointer) {
	var v1 *unsafe.Pointer
	v1 = this
	operator_delete(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(this), unsafe.Sizeof(unsafe.Pointer(nil))*5)))
	sub_57DF70((*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(unsafe.Pointer(nil))*2)))
	operator_delete(*v1)
}
func sub_57D1C0(this *unsafe.Pointer, a2 int32, a3 *uint8, a4 int32) int32 {
	var (
		v4  *unsafe.Pointer
		v5  unsafe.Pointer
		v6  int32
		v7  *uint8
		v8  int32
		v9  int32
		v10 int32
		v11 int32
		v12 int32
		v13 *uint16
		v14 uint32
		v15 int32
		v16 *uint16
	)
	_ = v16
	var v17 int32
	var v18 int32
	var v19 int32
	var v20 int32
	var i *uint8
	var v22 int32
	var v23 int32
	var v24 int32
	var v25 int32
	var v26 int32
	var v27 int32
	var v28 *uint8
	var v29 int32
	var v30 int32
	var v31 uint8
	var v32 int32
	var v33 int32
	var v34 int32
	var v35 *byte
	var v36 int32
	var v37 int32
	var v38 int32
	var v39 *uint8
	var v40 *uint8
	var v41 *uint8
	var v42 uint32
	var v43 int32
	var v44 int32
	var v45 int32
	var v46 int32
	var v47 int32
	var v48 int32
	var v49 *byte
	var v50 *uint8
	var v51 uint32
	var v52 *byte
	var v53 *uint8
	var v54 int8
	var v55 int32
	var v56 int32
	var v57 int32
	var v58 *uint8
	var v59 int32
	var v60 int32
	var v61 *uint8
	var v62 uint32
	var v63 int32
	var v64 int32
	var v65 int32
	var v66 int32
	var v67 int32
	var v68 int32
	var v69 int32
	var v70 int32
	var v71 int32
	var v72 int32
	var v73 *uint8
	var v74 *uint8
	var v75 uint32
	var v76 int32
	var v77 int32
	var v78 int32
	var v79 int32
	var v80 int32
	var v81 *byte
	var v82 *uint8
	var v83 uint32
	var v84 *byte
	var v85 int32
	var v86 *uint8
	var v87 int32
	var v88 int32
	var v89 *uint8
	var v90 *uint8
	var v91 uint32
	var v92 int32
	var v93 int32
	var v94 int32
	var v95 int32
	var v96 int32
	var v97 int32
	var v98 int32
	var v99 int32
	var v100 int32
	var v101 *uint8
	var v102 int32
	var v103 int32
	var v104 *uint8
	var v105 uint32
	var v106 int32
	var v107 *uint8
	var v108 int32
	var v110 int32
	var v111 int32
	var v112 int32
	var v113 int32
	var v114 int32
	var v115 int32
	var v116 int32
	var v117 int32
	var v118 *uint8
	var v119 int32
	var v120 int32
	var v121 int32
	var v122 uint16
	var v123 int32
	var v124 int32
	var v125 int32
	var v126 int32
	var v127 int32
	var v128 int16
	var v129 int32
	var v130 int32
	var v131 int32
	var v132 int32
	var v133 unsafe.Pointer
	var v134 int32
	var v135 [10]int32
	var v136 int32
	var v137 *uint8
	var v138 *uint8
	var v139 *uint8
	var v140 *uint8
	var v141 int32
	var v142 *uint8
	var v143 *uint8
	v4 = this
	v5 = *(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(this), unsafe.Sizeof(unsafe.Pointer(nil))*2))
	v135[0] = int32(uintptr(unsafe.Pointer((*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*2)))))
	v135[2] = int32(uintptr(unsafe.Pointer((*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*3)))))
	v135[4] = a2
	v135[3] = a2
	v6 = 0
	v135[1] = int32(uintptr(v5))
	v135[5] = 0
	v135[6] = 0
	v135[9] = 0
	v136 = 0
	v118 = (*uint8)(unsafe.Add(unsafe.Pointer(a3), a4))
	if uint32(a4) >= 5 {
		v7 = a3
		if uintptr(unsafe.Pointer(a3)) < uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(a3), 5)))) {
			for {
				v8 = int32(*func() *uint8 {
					p := &v7
					x := *p
					*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), 1))
					return x
				}())
				v6 = int32(compat_rotl(uint32(v8^v6), 5))
				if uintptr(unsafe.Pointer(v7)) >= uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(a3), 5)))) {
					break
				}
			}
			v136 = v6
		}
	}
	v9 = a4
	v111 = a4
	if a4 < 5 {
		goto LABEL_144
	}
	v10 = v136
	for 2 != 0 {
		v11 = 0
		v126 = v9 - 5
		if v9-5 >= 64 {
			v126 = 64
		}
		v110 = 0
		v125 = 0
		for {
			*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v15))), unsafe.Sizeof(uint16(0))*1)) = 0
			v12 = 0
			v13 = (*uint16)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*5)))
			v14 = (uint32(v10)*0x343FD + 0x269EC3) >> 17
			*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v15))), unsafe.Sizeof(uint16(0))*0)) = *(*uint16)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(uint16(0))*uintptr(v14)))
			v16 = (*uint16)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(uint16(0))*uintptr(v14)))
			v17 = int32(uint32(uintptr(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*1)))) & math.MaxUint16)
			v117 = v17
			*v16 = uint16(int16(v17))
			if int32(uint16(int16(v15))) == math.MaxUint16 || v15 == v17 {
				goto LABEL_62
			}
			v18 = int32(uint16(int16(v17 - v15)))
			v19 = v18
			if v18 >= v111-v11 {
				v19 = v111 - v11
			}
			if v19 >= 521 {
				v18 = 521
			} else if v18 >= v111-v11 {
				v18 = v111 - v11
			}
			v112 = v18
			v138 = (*uint8)(unsafe.Add(unsafe.Pointer(a3), v11))
			v119 = int32(0x10000 - uint32(v15))
			if 0x10000-uint32(v15) < uint32(v18) {
				v22 = v15
				if uint32(v15) >= 0x10000 {
				LABEL_33:
					v23 = 0
					if uint32(v18+v15)-0x10000 <= 0 {
					LABEL_23:
						v12 = v18
						goto LABEL_24
					}
					for int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(*v4)), v23)))) == int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v138), v119+v23))) {
						v18 = v112
						if uint32(func() int32 {
							p := &v23
							*p++
							return *p
						}()) >= uint32(v112+v15)-0x10000 {
							goto LABEL_23
						}
					}
					v12 = v119 + v23
				} else {
					for int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(*v4)), v22)))) == int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v138), v22-v15))) {
						if uint32(func() int32 {
							p := &v22
							*p++
							return *p
						}()) >= 0x10000 {
							v10 = v136
							goto LABEL_33
						}
					}
					v10 = v136
					v12 = v22 - v15
				}
			} else {
				v12 = 0
				if v18 <= 0 {
					goto LABEL_23
				}
				for int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(*v4)), v15))), v12)))) == int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a3), v11+v12))) {
					v18 = v112
					if func() int32 {
						p := &v12
						*p++
						return *p
					}() >= v112 {
						goto LABEL_23
					}
				}
			}
		LABEL_24:
			if int32(uint16(int16(v15+v12))) == v117 {
				v20 = v111 + (-v12) - v11
				if 521-v12 < v20 {
					v20 = 521 - v12
				}
				v120 = v20
				v113 = 0
				if v20 > 0 {
					for i = (*uint8)(unsafe.Add(unsafe.Pointer(a3), v11)); int32(*i) == int32(*(*uint8)(unsafe.Add(unsafe.Pointer(i), v12))); i = (*uint8)(unsafe.Add(unsafe.Pointer(i), 1)) {
						if func() int32 {
							p := &v113
							*p++
							return *p
						}() >= v120 {
							break
						}
					}
				}
				v12 += v113
			}
			if v12 >= 3 {
				v114 = 521 - v12
				if 521-v12 >= v11 {
					v114 = v11
				}
				v24 = int32(uint16(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*1)))), -v15))))))
				if v114 >= v24-v12 {
					v121 = int32(uint16(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*1)))), -v15)))))) - v12
				} else {
					v121 = v114
				}
				v25 = int32(0x10000 - uint32(v24))
				if uint32(v121) < 0x10000-uint32(v24) {
					v25 = v114
					if v114 < v24-v12 {
					LABEL_53:
						v26 = 0
						if v25 <= 0 {
							goto LABEL_152
						}
						v122 = uint16(int16(v15 - 1))
						v139 = (*uint8)(unsafe.Add(unsafe.Pointer(v138), -1))
						for {
							if int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(*v4)), v122)))) != int32(*v139) {
								break
							}
							v26++
							v122--
							v139 = (*uint8)(unsafe.Add(unsafe.Pointer(v139), -1))
							if v26 >= v114 {
								break
							}
						}
						if v26 <= 0 {
						LABEL_152:
							v10 = v136
						} else {
							v11 -= v26
							v27 = int32(uintptr(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*1)))) - v26
							*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v15))), unsafe.Sizeof(uint16(0))*0)) = uint16(int16(v15 - v26))
							v12 += v26
							*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*1)) = unsafe.Pointer(uintptr(v27))
							v28 = (*uint8)(unsafe.Add(unsafe.Pointer(a3), v11))
							*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v117))), unsafe.Sizeof(uint16(0))*0)) = uint16(int16(v27))
							v10 = 0
							if uintptr(unsafe.Pointer(v28)) < uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v28), 5)))) {
								for {
									v29 = int32(*func() *uint8 {
										p := &v28
										x := *p
										*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), 1))
										return x
									}()) ^ v10
									v10 = int32(compat_rotl(uint32(v29), 5))
									if uintptr(unsafe.Pointer(v28)) >= uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(a3), v11+5)))) {
										break
									}
								}
							}
							v136 = v10
							v125 = 1
						}
						goto LABEL_62
					}
					v25 = int32(uint16(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*1)))), -v15)))))) - v12
				}
				v114 = v25
				goto LABEL_53
			}
		LABEL_62:
			v30 = v110
			if v110 >= 4 {
				v115 = v11
				if v12 <= v110 {
					*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*1)) = v133
					v137 = (*uint8)(unsafe.Add(unsafe.Pointer(a3), v127))
					v55 = int32(compat_rotl(uint32(v134^int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a3), v127+5))))^compat_rotl(uint32(*(*uint8)(unsafe.Add(unsafe.Pointer(a3), v127))), 25), 5))
					sub_57E4C0((**uint32)(unsafe.Pointer(&v135[0])), uint32(uintptr(unsafe.Pointer(a3))), v127, uint32(v110-4), uint32(uint16(int16(int32(uint16(uintptr(v133)))-int32(v128)))))
					v56 = int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v118), -v127)))) - uintptr(unsafe.Pointer(a3)))
					v57 = v110 - 2
					if v110-2 >= v56-2 {
						v57 = v56 - 2
					}
					v58 = (*uint8)(unsafe.Add(unsafe.Pointer(a3), v127))
					v59 = int32(uintptr(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*1)))) + 2
					v141 = int32(uintptr(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*1)))) + 2
					if v57 > 0 {
						if v57 <= 1024 {
							v64 = int32(compat_rotl(uint32(v55^int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v137), 6))))^compat_rotl(uint32(*(*uint8)(unsafe.Add(unsafe.Pointer(v137), 1))), 25), 5))
							v65 = 0
							v130 = int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v137), 7)))))
							for {
								*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*5)))), unsafe.Sizeof(uint16(0))*uintptr((uint32(v64)*0x343FD+0x269EC3)>>17)))) = uint16(int16(v65 + v59))
								v66 = int32(*(*uint8)(unsafe.Pointer(uintptr(v130 + v65))))
								v67 = int32(compat_rotl(uint32(*(*uint8)(unsafe.Pointer(uintptr(v130 + func() int32 {
									p := &v65
									x := *p
									*p++
									return x
								}() - 5)))), 25))
								v64 = int32(compat_rotl(uint32(v64^v66^v67), 5))
								if v65 >= v57 {
									break
								}
								*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v59))), unsafe.Sizeof(uint16(0))*0)) = uint16(int16(v141))
							}
							v58 = (*uint8)(unsafe.Add(unsafe.Pointer(a3), v127))
							v136 = v64
						} else {
							v60 = 0
							v61 = (*uint8)(unsafe.Add(unsafe.Pointer(v137), v57+2))
							v62 = uint32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v61), 5)))))
							if uintptr(unsafe.Pointer(v61)) < uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v61), 5)))) {
								for {
									v63 = int32(*func() *uint8 {
										p := &v61
										x := *p
										*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), 1))
										return x
									}())
									v60 = int32(compat_rotl(uint32(v63^v60), 5))
									if uint32(uintptr(unsafe.Pointer(v61))) >= v62 {
										break
									}
								}
							}
							v136 = v60
						}
					} else {
						v136 = 0
					}
					v68 = int32(uint32(uintptr(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*1)))) & math.MaxUint16)
					if uint32(v68+v110) <= 0x10000 {
						v69 = v110
						libc.MemCpy(unsafe.Add(unsafe.Pointer((*byte)(*v4)), v68), unsafe.Pointer(v58), int(v110))
					} else {
						libc.MemCpy(unsafe.Add(unsafe.Pointer((*byte)(*v4)), v68), unsafe.Pointer(v58), int(0x10000-uint32(v68)))
						v69 = v110
						libc.MemCpy(*v4, unsafe.Add(unsafe.Pointer(v58), 0x10000-uint32(v68)), int(uint32(v110)-(0x10000-uint32(v68))))
					}
					*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*1)) = unsafe.Add(unsafe.Pointer((*byte)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*1)))), v69)
					a3 = (*uint8)(unsafe.Add(unsafe.Pointer(a3), v69+v127))
					goto LABEL_143
				}
				sub_57E4C0((**uint32)(unsafe.Pointer(&v135[0])), uint32(uintptr(unsafe.Pointer(a3))), v11, uint32(v12-4), uint32(uint16(int16(v117-v15))))
				v36 = int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v118), -v11))))-uintptr(unsafe.Pointer(a3))) - 1
				if v12-1 < v36 {
					v36 = v12 - 1
				}
				v37 = 0
				v38 = int32(uintptr(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*1)))) + 1
				v39 = (*uint8)(unsafe.Add(unsafe.Pointer(a3), v11))
				v123 = v36
				v129 = int32(uintptr(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*1)))) + 1
				v140 = (*uint8)(unsafe.Add(unsafe.Pointer(a3), v11))
				v40 = (*uint8)(unsafe.Add(unsafe.Pointer(a3), v11))
				if v36 > 0 {
					if v36 <= 1024 {
						v44 = int32(compat_rotl(uint32(v136^int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v40), 5))))^compat_rotl(uint32(*v40), 25), 5))
						v45 = int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v40), 1)))))
						if v123 > 0 {
							for {
								*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*5)))), unsafe.Sizeof(uint16(0))*uintptr((uint32(v44)*0x343FD+0x269EC3)>>17)))) = uint16(int16(v37 + v38))
								v46 = int32(uint32(v44^int32(*(*uint8)(unsafe.Pointer(uintptr(v45 + 5 + v37))))) ^ compat_rotl(uint32(*(*uint8)(unsafe.Pointer(uintptr(v45 + v37)))), 25))
								v37++
								v44 = int32(compat_rotl(uint32(v46), 5))
								if v37 >= v123 {
									break
								}
								*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v38))), unsafe.Sizeof(uint16(0))*0)) = uint16(int16(v129))
							}
						}
						v136 = v44
					} else {
						v41 = (*uint8)(unsafe.Add(unsafe.Pointer(v40), v36+1))
						v42 = uint32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v41), 5)))))
						if uintptr(unsafe.Pointer(v41)) < uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v41), 5)))) {
							for {
								v43 = int32(*func() *uint8 {
									p := &v41
									x := *p
									*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), 1))
									return x
								}())
								v37 = int32(compat_rotl(uint32(v43^v37), 5))
								if uint32(uintptr(unsafe.Pointer(v41))) >= v42 {
									break
								}
							}
						}
						v136 = v37
					}
					v39 = v140
				} else {
					v136 = 0
				}
				v47 = int32(uint32(uintptr(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*1)))) & math.MaxUint16)
				if uint32(v47+v12) > 0x10000 {
					v48 = int32(0x10000 - uint32(v47))
					libc.MemCpy(unsafe.Add(unsafe.Pointer((*byte)(*v4)), v47), unsafe.Pointer(v39), int(0x10000-uint32(v47)))
					v49 = (*byte)(*v4)
					v50 = (*uint8)(unsafe.Add(unsafe.Pointer(v140), v48))
					v51 = uint32(v12 - v48)
					libc.MemCpy(*v4, unsafe.Pointer(v50), int((v51>>2)*4))
					v53 = (*uint8)(unsafe.Add(unsafe.Pointer(v50), (v51>>2)*4))
					v52 = (*byte)(unsafe.Add(unsafe.Pointer(v49), (v51>>2)*4))
					v54 = int8(uint8(v51))
				LABEL_118:
					libc.MemCpy(unsafe.Pointer(v52), unsafe.Pointer(v53), int(int32(v54)&3))
					*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*1)) = unsafe.Add(unsafe.Pointer((*byte)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*1)))), v12)
					a3 = (*uint8)(unsafe.Add(unsafe.Pointer(a3), v115+v12))
					goto LABEL_143
				}
			LABEL_117:
				v84 = (*byte)(unsafe.Add(unsafe.Pointer((*byte)(*v4)), v47))
				libc.MemCpy(unsafe.Pointer(v84), unsafe.Pointer(v39), int((uint32(v12)>>2)*4))
				v53 = (*uint8)(unsafe.Add(unsafe.Pointer(v39), (uint32(v12)>>2)*4))
				v52 = (*byte)(unsafe.Add(unsafe.Pointer(v84), (uint32(v12)>>2)*4))
				v54 = int8(v12)
				goto LABEL_118
			}
			if v12 < 4 {
				goto LABEL_66
			}
			if v125 != 0 {
				v115 = v11
				sub_57E4C0((**uint32)(unsafe.Pointer(&v135[0])), uint32(uintptr(unsafe.Pointer(a3))), v11, uint32(v12-4), uint32(uint16(int16(v117-v15))))
				v70 = int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v118), -v11))))-uintptr(unsafe.Pointer(a3))) - 1
				if v12-1 < v70 {
					v70 = v12 - 1
				}
				v71 = 0
				v72 = int32(uintptr(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*1)))) + 1
				v39 = (*uint8)(unsafe.Add(unsafe.Pointer(a3), v11))
				v124 = v70
				v131 = int32(uintptr(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*1)))) + 1
				v142 = (*uint8)(unsafe.Add(unsafe.Pointer(a3), v11))
				v73 = (*uint8)(unsafe.Add(unsafe.Pointer(a3), v11))
				if v70 > 0 {
					if v70 <= 1024 {
						v77 = int32(compat_rotl(uint32(v136^int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v73), 5))))^compat_rotl(uint32(*v73), 25), 5))
						v78 = int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v73), 1)))))
						if v124 > 0 {
							for {
								*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*5)))), unsafe.Sizeof(uint16(0))*uintptr((uint32(v77)*0x343FD+0x269EC3)>>17)))) = uint16(int16(v71 + v72))
								v79 = int32(uint32(v77^int32(*(*uint8)(unsafe.Pointer(uintptr(v78 + 5 + v71))))) ^ compat_rotl(uint32(*(*uint8)(unsafe.Pointer(uintptr(v78 + v71)))), 25))
								v71++
								v77 = int32(compat_rotl(uint32(v79), 5))
								if v71 >= v124 {
									break
								}
								*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v72))), unsafe.Sizeof(uint16(0))*0)) = uint16(int16(v131))
							}
						}
						v136 = v77
					} else {
						v74 = (*uint8)(unsafe.Add(unsafe.Pointer(v73), v70+1))
						v75 = uint32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v74), 5)))))
						if uintptr(unsafe.Pointer(v74)) < uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v74), 5)))) {
							for {
								v76 = int32(*func() *uint8 {
									p := &v74
									x := *p
									*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), 1))
									return x
								}())
								v71 = int32(compat_rotl(uint32(v76^v71), 5))
								if uint32(uintptr(unsafe.Pointer(v74))) >= v75 {
									break
								}
							}
						}
						v136 = v71
					}
					v39 = v142
				} else {
					v136 = 0
				}
				v47 = int32(uint32(uintptr(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*1)))) & math.MaxUint16)
				if uint32(v47+v12) > 0x10000 {
					v80 = int32(0x10000 - uint32(v47))
					libc.MemCpy(unsafe.Add(unsafe.Pointer((*byte)(*v4)), v47), unsafe.Pointer(v39), int(0x10000-uint32(v47)))
					v81 = (*byte)(*v4)
					v82 = (*uint8)(unsafe.Add(unsafe.Pointer(v142), v80))
					v83 = uint32(v12 - v80)
					libc.MemCpy(*v4, unsafe.Pointer(v82), int((v83>>2)*4))
					v53 = (*uint8)(unsafe.Add(unsafe.Pointer(v82), (v83>>2)*4))
					v52 = (*byte)(unsafe.Add(unsafe.Pointer(v81), (v83>>2)*4))
					v54 = int8(uint8(v83))
					goto LABEL_118
				}
				goto LABEL_117
			}
			v30 = v12
			v110 = v12
			v128 = int16(v15)
			v127 = v11
			v133 = *(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*1))
			v134 = v10
		LABEL_66:
			if v11+1 > v126 {
				break
			}
			v31 = *(*uint8)(unsafe.Add(unsafe.Pointer(a3), v11))
			v32 = int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a3), v11+5)))
			v33 = int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a3), func() int32 {
				p := &v11
				x := *p
				*p++
				return x
			}())))
			v34 = int32(uint32(v32) ^ compat_rotl(uint32(v33), 25))
			v35 = (*byte)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*1)))
			v10 = int32(compat_rotl(uint32(v10^v34), 5))
			*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*1)) = unsafe.Add(unsafe.Pointer(v35), 1)
			v136 = v10
			*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(*v4)), uint16(uintptr(unsafe.Pointer(v35)))))) = v31
		}
		v116 = v11
		if v30 < 4 {
			if v11+5 >= v111 && v111 <= 64 {
				v100 = v111 - v11
				v101 = (*uint8)(unsafe.Add(unsafe.Pointer(a3), v11))
				v102 = int32(uint32(uintptr(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*1)))) & math.MaxUint16)
				if uint32(v102+v100) <= 0x10000 {
					v105 = uint32(v100)
					v104 = (*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(*v4)), v102))
				} else {
					v103 = int32(0x10000 - uint32(v102))
					libc.MemCpy(unsafe.Add(unsafe.Pointer((*byte)(*v4)), v102), unsafe.Pointer(v101), int(0x10000-uint32(v102)))
					v104 = (*uint8)(*v4)
					v105 = uint32(v100 - v103)
					v101 = (*uint8)(unsafe.Add(unsafe.Pointer(v101), v103))
				}
				libc.MemCpy(unsafe.Pointer(v104), unsafe.Pointer(v101), int(v105))
				*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*1)) = unsafe.Add(unsafe.Pointer((*byte)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*1)))), v100)
				v11 = v111
			}
			sub_57E3F0((**uint32)(unsafe.Pointer(&v135[0])), uint32(uintptr(unsafe.Pointer(a3))), v11)
			a3 = (*uint8)(unsafe.Add(unsafe.Pointer(a3), v11))
		} else {
			sub_57E4C0((**uint32)(unsafe.Pointer(&v135[0])), uint32(uintptr(unsafe.Pointer(a3))), v11, uint32(v110-4), uint32(uint16(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*1)))), -v128)))))))
			v85 = v110 - 1
			if v110-1 >= int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v118), -v11))))-uintptr(unsafe.Pointer(a3)))-1 {
				v85 = int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v118), -v11))))-uintptr(unsafe.Pointer(a3))) - 1
			}
			v86 = (*uint8)(unsafe.Add(unsafe.Pointer(a3), v11))
			v87 = 0
			v88 = int32(uintptr(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*1)))) + 1
			v132 = int32(uintptr(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*1)))) + 1
			v143 = (*uint8)(unsafe.Add(unsafe.Pointer(a3), v11))
			v89 = (*uint8)(unsafe.Add(unsafe.Pointer(a3), v11))
			if v85 > 0 {
				if v85 <= 1024 {
					v93 = int32(uint32(v136^int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v89), 5)))) ^ compat_rotl(uint32(*v86), 25))
					v94 = int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v89), 1)))))
					v95 = int32(compat_rotl(uint32(v93), 5))
					for {
						*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*5)))), unsafe.Sizeof(uint16(0))*uintptr((uint32(v95)*0x343FD+0x269EC3)>>17)))) = uint16(int16(v87 + v88))
						v96 = int32(*(*uint8)(unsafe.Pointer(uintptr(v94 + 5 + v87))))
						v97 = int32(compat_rotl(uint32(*(*uint8)(unsafe.Pointer(uintptr(v94 + func() int32 {
							p := &v87
							x := *p
							*p++
							return x
						}())))), 25))
						v95 = int32(compat_rotl(uint32(v95^v96^v97), 5))
						if v87 >= v85 {
							break
						}
						*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v88))), unsafe.Sizeof(uint16(0))*0)) = uint16(int16(v132))
					}
					v86 = v143
					v136 = v95
				} else {
					v90 = (*uint8)(unsafe.Add(unsafe.Pointer(v86), v85+1))
					v91 = uint32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v90), 5)))))
					if uintptr(unsafe.Pointer(v90)) < uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v90), 5)))) {
						for {
							v92 = int32(*func() *uint8 {
								p := &v90
								x := *p
								*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), 1))
								return x
							}())
							v87 = int32(compat_rotl(uint32(v92^v87), 5))
							if uint32(uintptr(unsafe.Pointer(v90))) >= v91 {
								break
							}
						}
					}
					v136 = v87
				}
			} else {
				v136 = 0
			}
			v98 = int32(uint32(uintptr(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*1)))) & math.MaxUint16)
			if uint32(v98+v110) <= 0x10000 {
				v99 = v110
				libc.MemCpy(unsafe.Add(unsafe.Pointer((*byte)(*v4)), v98), unsafe.Pointer(v86), int(v110))
			} else {
				libc.MemCpy(unsafe.Add(unsafe.Pointer((*byte)(*v4)), v98), unsafe.Pointer(v86), int(0x10000-uint32(v98)))
				v99 = v110
				libc.MemCpy(*v4, unsafe.Add(unsafe.Pointer(v86), 0x10000-uint32(v98)), int(uint32(v110)-(0x10000-uint32(v98))))
			}
			*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*1)) = unsafe.Add(unsafe.Pointer((*byte)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*1)))), v99)
			a3 = (*uint8)(unsafe.Add(unsafe.Pointer(a3), v116+v99))
		}
	LABEL_143:
		v9 = int32(uintptr(unsafe.Pointer(v118)) - uintptr(unsafe.Pointer(a3)))
		v111 = int32(uintptr(unsafe.Pointer(v118)) - uintptr(unsafe.Pointer(a3)))
		if int32(uintptr(unsafe.Pointer(v118))-uintptr(unsafe.Pointer(a3))) >= 5 {
			v10 = v136
			continue
		}
		break
	}
LABEL_144:
	if v111 != 0 {
		v106 = int32(uint32(uintptr(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*1)))) & math.MaxUint16)
		if uint32(v106+v111) <= 0x10000 {
			v108 = v111
			libc.MemCpy(unsafe.Add(unsafe.Pointer((*byte)(*v4)), v106), unsafe.Pointer(a3), int(v111))
			v107 = a3
		} else {
			v107 = a3
			libc.MemCpy(unsafe.Add(unsafe.Pointer((*byte)(*v4)), v106), unsafe.Pointer(a3), int(0x10000-uint32(v106)))
			v108 = v111
			libc.MemCpy(*v4, unsafe.Add(unsafe.Pointer(a3), 0x10000-uint32(v106)), int(uint32(v111)-(0x10000-uint32(v106))))
		}
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*1)) = unsafe.Add(unsafe.Pointer((*byte)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*1)))), v108)
		sub_57E3F0((**uint32)(unsafe.Pointer(&v135[0])), uint32(uintptr(unsafe.Pointer(v107))), v108)
	}
	return sub_57E7D0((**uint32)(unsafe.Pointer(&v135[0])))
}
func sub_57DD70(this *unsafe.Pointer) {
	operator_delete(*this)
}
func sub_57DD90(this *uint32) *uint32 {
	var (
		v1 *uint32
		v2 unsafe.Pointer
	)
	v1 = this
	v2 = operator_new(548)
	*v1 = uint32(uintptr(v2))
	libc.MemSet(v2, 0, 548)
	return v1
}
func sub_57DDC0(this *unsafe.Pointer) {
	operator_delete(*this)
}
func sub_57DDD0(this *unsafe.Pointer) int32 {
	var result int32
	result = 0
	libc.MemSet(*this, 0, 548)
	return result
}
func sub_57DDE0(a1 int32, a2 int32) uint32 {
	var (
		v2     int32
		v3     int32
		i      int32
		v5     int32
		v6     int32
		v7     int32
		v8     int32
		result uint32
		v10    int32
		v11    int32
		v12    int32
	)
	v2 = a1
	v3 = a2
	for i = 40; i > 0; i /= 3 {
		v5 = i + 1
		v11 = i + 1
		if i+1 <= v3 {
			for {
				v6 = v5 * 4
				v10 = v5
				v12 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5*4 + v2))))
				if v5 > i {
					v7 = i * 4
					for {
						v8 = int32(*(*int16)(unsafe.Pointer(uintptr(v6 - v7 + v2 + 2)))) - int32(*(*int16)(unsafe.Add(unsafe.Pointer((*int16)(unsafe.Pointer(&v12))), unsafe.Sizeof(int16(0))*1)))
						if v8 == 0 {
							v8 = int32(*(*int16)(unsafe.Pointer(uintptr(v6 - v7 + v2)))) - int32(int16(v12))
						}
						if v8 >= 0 {
							break
						}
						*(*uint32)(unsafe.Pointer(uintptr(v6 + v2))) = *(*uint32)(unsafe.Pointer(uintptr(v6 - v7 + v2)))
						v6 -= v7
						v10 -= i
						if v10 <= i {
							break
						}
					}
					v5 = v11
				}
				v5++
				*(*uint32)(unsafe.Pointer(uintptr(v2 + v10*4))) = uint32(v12)
				v3 = a2
				v11 = v5
				if v5 > a2 {
					break
				}
			}
		}
		result = uint32(i/3) >> 31
	}
	return result
}
func sub_57DEA0(this *uint32, a2 *uint16) int32 {
	var (
		v2 *uint16
		v3 int32
		i  int32
		v5 *int16
		v6 int16
	)
	v2 = a2
	v3 = 0
	for i = 0; i < 274; i++ {
		*v2 = uint16(int16(i))
		v2 = (*uint16)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint16(0))*2))
		*((*uint16)(unsafe.Add(unsafe.Pointer(v2), -int(unsafe.Sizeof(uint16(0))*1)))) = *(*uint16)(unsafe.Pointer(uintptr(*this + uint32(i*2))))
		v5 = (*int16)(unsafe.Pointer(uintptr(*this + uint32(i*2))))
		v6 = *v5
		v3 += int32(v6)
		*v5 = int16(int32(v6) >> 1)
	}
	sub_57DDE0(int32(uintptr(unsafe.Pointer((*uint16)(unsafe.Add(unsafe.Pointer(a2), -int(unsafe.Sizeof(uint16(0))*2)))))), 274)
	return v3
}
func sub_57DF00(this *uint32) *uint32 {
	var (
		v1     *uint32
		v2     *uint8
		v3     *uint8
		result *uint32
	)
	v1 = this
	sub_57DD90(this)
	*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*1)) = 4096
	v2 = (*uint8)(operator_new(1096))
	v3 = v2
	*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*2)) = uint32(uintptr(unsafe.Pointer(v2)))
	result = v1
	libc.MemCpy(unsafe.Pointer(v3), memmap.PtrOff(0x587000, 313544), 1096)
	return result
}
func sub_57DF70(this *unsafe.Pointer) {
	var v1 *unsafe.Pointer
	v1 = this
	operator_delete(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(this), unsafe.Sizeof(unsafe.Pointer(nil))*2)))
	sub_57DDC0(v1)
}
func sub_57DFC0(this *uint32, a2 *int32) int32 {
	var (
		v2     *uint32
		v3     int32
		v4     int32
		v5     int32
		v6     int32
		v7     int32
		v8     int32
		v9     int32
		v10    int32
		v11    int32
		v12    *byte
		v13    int32
		v14    *int32
		v15    int32
		v16    int32
		v17    int32
		v18    *byte
		v19    int32
		v20    int32
		v21    int32
		v22    int32
		v23    int32
		v24    *byte
		v25    int32
		v26    int32
		v27    *int32
		v28    int32
		v29    *int32
		v30    int32
		v31    *int32
		v32    int32
		v33    int32
		v34    int32
		v35    int32
		result int32
		v37    int32
		v38    int16
		v39    int32
		v40    *byte
		v41    int32
		v42    int16
		v43    *int32
		v44    int32
		v45    int32
		v46    int32
		v47    int32
		v48    *int32
		v49    int32
		v50    int32
		v51    int32
		v52    int32
		v53    int32
		v54    int32
		v55    int32
		v56    int32
		v57    int32
		v58    int32
		v59    int32
		v60    [1100]byte
	)
	v2 = this
	*(*uint32)(unsafe.Pointer(&v60[0])) = uint32(uintptr(unsafe.Pointer(this)))
	v3 = sub_57DEA0(this, (*uint16)(unsafe.Pointer(&v60[4])))
	v4 = 0
	v5 = 0
	v6 = 16
	v56 = v3
	*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*1)) = 4096
	v55 = 0
	v46 = 0
	v52 = 16
	v43 = (*int32)(unsafe.Add(unsafe.Pointer(a2), -int(unsafe.Sizeof(int32(0))*1)))
	for {
		v7 = 0
		v58 = 0
		v8 = 0
		v59 = (v56 - v55) / v6
		for {
			v49 = 0
			v9 = 1 << v4
			if v5+(1<<v4)+v7 > 274 {
				v49 = 1
				v9 = 274 - v5
			}
			if v7 < v9 {
				v10 = v7 + v5
				v11 = v9 - v7
				v7 += v11
				v12 = &v60[v10*4+6]
				for {
					v8 += int32(*(*int16)(unsafe.Pointer(v12)))
					v12 = (*byte)(unsafe.Add(unsafe.Pointer(v12), 4))
					v11--
					if v11 == 0 {
						break
					}
				}
				v6 = v52
			}
			if v49 != 0 || v4 >= 8 || v8 > v59 {
				break
			}
			v58 = v8
			v4++
		}
		if v4 != 0 && cmath.Abs(int64(v58-v59)) <= cmath.Abs(int64(v8-v59)) {
			v8 = v58
			v4--
		}
		v13 = v46
		if v6 < 16 {
			v14 = v43
			for {
				if v4 >= *v14 {
					break
				}
				*(*int32)(unsafe.Add(unsafe.Pointer(v14), unsafe.Sizeof(int32(0))*1)) = *v14
				v13--
				v14 = (*int32)(unsafe.Add(unsafe.Pointer(v14), -int(unsafe.Sizeof(int32(0))*1)))
				if v13 <= 0 {
					break
				}
			}
		}
		*(*int32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(int32(0))*uintptr(v13))) = v4
		v55 += v8
		v43 = (*int32)(unsafe.Add(unsafe.Pointer(v43), unsafe.Sizeof(int32(0))*1))
		v5 += 1 << v4
		v6--
		v46++
		v52 = v6
		if v6 <= 2 {
			break
		}
		v4 = 0
	}
	v15 = 0
	v16 = 0
	v17 = 0
	v44 = 0
	v47 = 0
	v50 = math.MaxInt32
	v53 = 0
	if v5 < 274 {
		v18 = &v60[v5*4+6]
		v19 = 274 - v5
		for {
			v17 += int32(*(*int16)(unsafe.Pointer(v18)))
			v18 = (*byte)(unsafe.Add(unsafe.Pointer(v18), 4))
			v19--
			if v19 == 0 {
				break
			}
		}
		v53 = v17
	}
	v20 = 0
	v21 = 1
	if v5+1 <= 274 {
		for {
			if v15 < v21 {
				v22 = v15 + v5
				v23 = v21 - v15
				v15 += v23
				v24 = &v60[v22*4+6]
				for {
					v16 += int32(*(*int16)(unsafe.Pointer(v24)))
					v24 = (*byte)(unsafe.Add(unsafe.Pointer(v24), 4))
					v23--
					if v23 == 0 {
						break
					}
				}
			}
			v25 = 0
			v26 = 274 - v15 - v5
			if v26 > 1 {
				for {
					v25++
					if 1<<v25 >= v26 {
						break
					}
				}
			}
			if v20 <= 8 && v25 <= 8 {
				if v16*v20+v25*(v53-v16) >= v50 {
					break
				}
				v50 = v16*v20 + v25*(v53-v16)
				v44 = v20
				v47 = v25
			}
			v21 = 1 << func() int32 {
				p := &v20
				*p++
				return *p
			}()
			if v5+(1<<v20)+v15 > 274 {
				break
			}
		}
	}
	v27 = a2
	v28 = 14
	v29 = (*int32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(int32(0))*13))
	for {
		if v44 >= *v29 {
			break
		}
		*(*int32)(unsafe.Add(unsafe.Pointer(v29), unsafe.Sizeof(int32(0))*1)) = *v29
		v28--
		v29 = (*int32)(unsafe.Add(unsafe.Pointer(v29), -int(unsafe.Sizeof(int32(0))*1)))
		if v28 <= 0 {
			break
		}
	}
	*(*int32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(int32(0))*uintptr(v28))) = v44
	v30 = 15
	v31 = (*int32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(int32(0))*14))
	for {
		if v47 >= *v31 {
			break
		}
		*(*int32)(unsafe.Add(unsafe.Pointer(v31), unsafe.Sizeof(int32(0))*1)) = *v31
		v30--
		v31 = (*int32)(unsafe.Add(unsafe.Pointer(v31), -int(unsafe.Sizeof(int32(0))*1)))
		if v30 <= 0 {
			break
		}
	}
	*(*int32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(int32(0))*uintptr(v30))) = v47
	v32 = 0
	v33 = 0
	v51 = 0
	v45 = 0
	v48 = a2
	for {
		v34 = *v27
		v35 = 1 << *v27
		result = 274 - v32
		v57 = v35
		if v35 < 274-v32 {
			result = v35
		}
		v37 = 0
		v54 = result
		if result > 0 {
			v38 = int16(v34 + 4)
			v39 = v33 << v34
			v40 = &v60[v32*4+4]
			for {
				v41 = int32(*(*int16)(unsafe.Pointer(v40)))
				v40 = (*byte)(unsafe.Add(unsafe.Pointer(v40), 4))
				*(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(&v60[0])) + 8))) + uint32(v41*4)))) = uint16(v38)
				v42 = int16(func() int32 {
					p := &v37
					x := *p
					*p++
					return x
				}() | v39)
				*(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(&v60[0])) + 8))) + uint32(v41*4) + 2))) = uint16(v42)
				result = v54
				if v37 >= v54 {
					break
				}
			}
			v35 = v57
			v33 = v45
			v32 = v51
		}
		v32 += v35
		v33++
		v27 = (*int32)(unsafe.Add(unsafe.Pointer(v48), unsafe.Sizeof(int32(0))*1))
		v51 = v32
		v45 = v33
		v48 = (*int32)(unsafe.Add(unsafe.Pointer(v48), unsafe.Sizeof(int32(0))*1))
		if v33 >= 16 {
			break
		}
	}
	return result
}
func sub_57E2C0(this **uint32) int32 {
	var (
		v1     **uint32
		v2     *uint32
		v3     uint32
		v4     int32
		v5     int32
		v6     uint16
		v7     int32
		v8     *uint32
		v9     *uint8
		v10    int32
		v11    *uint32
		v12    int32
		v13    *byte
		v14    *uint32
		v15    int32
		v16    int8
		v17    int32
		v18    int32
		v19    uint32
		v20    int32
		v21    *uint8
		v22    uint32
		v23    *uint32
		v24    int32
		result int32
		v26    int32
		v27    [64]byte
	)
	v1 = this
	**(**uint32)(unsafe.Add(unsafe.Pointer(this), unsafe.Sizeof((*uint32)(nil))*2)) = 2
	if int32(func() uint32 {
		p := *(**uint32)(unsafe.Add(unsafe.Pointer(this), unsafe.Sizeof((*uint32)(nil))*2))
		*p--
		return *p
	}()) <= 0 {
		sub_57E2C0(this)
	}
	*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*1))))), unsafe.Sizeof(uint16(0))*272)))++
	v2 = *(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*6))
	v3 = uint32(uintptr(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*5)))))
	v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(*v1), unsafe.Sizeof(uint32(0))*2)))
	v5 = int32(*(*int16)(unsafe.Pointer(uintptr(v4 + 1088))))
	v6 = *(*uint16)(unsafe.Pointer(uintptr(v4 + 1090)))
	*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*6)) = (*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v2))), v5))))
	v7 = int32(uint32(int32(v6)<<(32-int32(uint8(uintptr(unsafe.Pointer(v2))))-v5)) | v3)
	*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*5)) = (*uint32)(unsafe.Pointer(uintptr(v7)))
	if int32(uintptr(unsafe.Pointer(v2)))+v5 >= 16 {
		*(*uint8)(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*3)))) = *(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v7))), unsafe.Sizeof(int32(0))-1))
		v8 = *(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*5))
		v9 = (*uint8)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*3))))), 1))
		*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*3)) = (*uint32)(unsafe.Pointer(v9))
		*v9 = *(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v8))), 2))
		v10 = int32(uintptr(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*6))), -int(unsafe.Sizeof(uint32(0))*4))))))
		v11 = (*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*5))))) << 16)))
		*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*3)) = (*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*3))))), 1))))
		*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*6)) = (*uint32)(unsafe.Pointer(uintptr(v10)))
		*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*5)) = v11
	}
	sub_57DFC0(*v1, (*int32)(unsafe.Pointer(&v27[0])))
	v12 = 0
	v13 = &v27[0]
	v26 = 16
	for {
		v14 = *(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*6))
		v15 = int32(*(*uint32)(unsafe.Pointer(v13)) - uint32(v12))
		v12 = int32(*(*uint32)(unsafe.Pointer(v13)))
		v16 = int8(32 - int32(uint8(uintptr(unsafe.Pointer(v14)))) - func() int32 {
			p := &v15
			*p++
			return *p
		}())
		v17 = int32(uintptr(unsafe.Pointer(v14))) + v15
		v18 = 1 << int32(v16)
		v19 = uint32(uintptr(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*5)))))
		*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*6)) = (*uint32)(unsafe.Pointer(uintptr(v17)))
		v20 = int32(uint32(v18) | v19)
		*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*5)) = (*uint32)(unsafe.Pointer(uintptr(v20)))
		if v17 >= 16 {
			*(*uint8)(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*3)))) = *(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v20))), unsafe.Sizeof(int32(0))-1))
			v21 = (*uint8)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*3))))), 1))
			v22 = uint32(uintptr(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*5))))) >> 16
			*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*3)) = (*uint32)(unsafe.Pointer(v21))
			*v21 = uint8(v22)
			v23 = *(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*5))
			v24 = int32(uintptr(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*6))), -int(unsafe.Sizeof(uint32(0))*4))))))
			*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*3)) = (*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*3))))), 1))))
			*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*6)) = (*uint32)(unsafe.Pointer(uintptr(v24)))
			*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*5)) = (*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(v23))) << 16)))
		}
		v13 = (*byte)(unsafe.Add(unsafe.Pointer(v13), 4))
		result = func() int32 {
			p := &v26
			*p--
			return *p
		}()
		if v26 == 0 {
			break
		}
	}
	return result
}
func sub_57E3F0(this **uint32, a2 uint32, a3 int32) uint32 {
	var (
		v3     *uint8
		result uint32
		v5     **uint32
		v6     bool
		v7     uint16
		v8     int32
		v9     *uint32
		v10    int32
		v11    uint16
		v12    int32
		v13    *uint32
		v14    *uint8
		v15    int32
		v16    *uint32
		v17    uint32
	)
	v3 = (*uint8)(unsafe.Pointer(uintptr(a2)))
	result = a2 + uint32(a3)
	v5 = this
	v6 = a2 < a2+uint32(a3)
	v17 = a2 + uint32(a3)
	if v6 {
		for {
			v7 = uint16(*v3)
			if int32(func() uint32 {
				p := *(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*2))
				*p--
				return *p
			}()) <= 0 {
				sub_57E2C0(v5)
			}
			v8 = int32(v7)
			*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*1))))), unsafe.Sizeof(uint16(0))*uintptr(v7))))++
			v9 = *(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*6))
			v10 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(*v5), unsafe.Sizeof(uint32(0))*2)))
			v11 = *(*uint16)(unsafe.Pointer(uintptr(v10 + v8*4 + 2)))
			v12 = int32(*(*int16)(unsafe.Pointer(uintptr(v10 + v8*4))))
			*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*5)) = (*uint32)(unsafe.Pointer(uintptr(uint32(int32(v11)<<(32-int32(uint8(uintptr(unsafe.Pointer(v9))))-v12)) | uint32(uintptr(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*5))))))))
			*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*6)) = (*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v9))), v12))))
			if int32(uintptr(unsafe.Pointer(v9)))+v12 >= 16 {
				*(*uint8)(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*3)))) = uint8(uint32(uintptr(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*5))))) >> 24)
				v13 = *(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*5))
				v14 = (*uint8)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*3))))), 1))
				*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*3)) = (*uint32)(unsafe.Pointer(v14))
				*v14 = *(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v13))), 2))
				v15 = int32(uintptr(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*6))), -int(unsafe.Sizeof(uint32(0))*4))))))
				v16 = (*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*5))))) << 16)))
				*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*3)) = (*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*3))))), 1))))
				*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*6)) = (*uint32)(unsafe.Pointer(uintptr(v15)))
				*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*5)) = v16
			}
			result = v17
			v3 = (*uint8)(unsafe.Add(unsafe.Pointer(v3), 1))
			if uint32(uintptr(unsafe.Pointer(v3))) >= v17 {
				break
			}
		}
	}
	return result
}
func sub_57E4C0(this **uint32, a2 uint32, a3 int32, a4 uint32, a5 uint32) int32 {
	var (
		v5     **uint32
		v6     uint16
		v7     int32
		v8     *uint32
		v9     uint32
		v10    int32
		v11    int32
		v12    uint16
		v13    uint16
		v14    uint32
		v15    int32
		v16    int32
		v17    *uint32
		v18    int32
		v19    int32
		v20    uint32
		v21    int32
		v22    int32
		v23    *uint8
		v24    uint16
		v25    *uint32
		v26    int32
		v27    uint32
		v28    int32
		v29    int32
		v30    *uint32
		v31    *uint8
		v32    int32
		v33    *uint32
		v34    int32
		v35    uint32
		v36    *uint32
		v37    uint32
		v38    int32
		v39    int32
		v40    *uint32
		v41    *uint8
		v42    int32
		v43    *uint32
		v44    *uint32
		v45    int32
		result int32
		v47    *uint32
		v48    int8
		v49    int32
		v50    uint32
		v51    *uint32
		v52    *uint8
		v53    int32
	)
	v5 = this
	sub_57E3F0(this, a2, a3)
	if a4 >= 8 {
		if a4 >= 38 {
			v23 = (*uint8)(memmap.PtrOff(0x587000, ((a4-38)>>5)*12+0x4CD10))
			v24 = uint16(int16(int32(*memmap.PtrUint16(0x587000, ((a4-38)>>5)*12+0x4CD10)) + 4))
			if int32(func() uint32 {
				p := *(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*2))
				*p--
				return *p
			}()) <= 0 {
				sub_57E2C0(v5)
			}
			*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*1))))), unsafe.Sizeof(uint16(0))*uintptr(v24))))++
			sub_57F160(int32(uintptr(unsafe.Pointer(v5))), int32(*(*int16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(*v5), unsafe.Sizeof(uint32(0))*2)) + uint32(int32(v24)*4))))), int32(*(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(*v5), unsafe.Sizeof(uint32(0))*2)) + uint32(int32(v24)*4) + 2)))))
			v25 = *(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*6))
			v26 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v23))), unsafe.Sizeof(uint32(0))*1))) + 4)
			v27 = uint32(uintptr(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*5)))))
			v28 = ((int32(uint8(a4))-38)&31 | int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v23))), unsafe.Sizeof(uint16(0))*4))))*16) << (32 - int32(uint8(uintptr(unsafe.Pointer(v25)))) - v26)
			v29 = int32(uintptr(unsafe.Pointer(v25))) + v26
			*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*6)) = (*uint32)(unsafe.Pointer(uintptr(v29)))
			*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*5)) = (*uint32)(unsafe.Pointer(uintptr(uint32(v28) | v27)))
			if v29 >= 16 {
				goto LABEL_14
			}
		} else {
			v13 = *memmap.PtrUint16(0x587000, ((a4-8)>>1)*12+0x4CD10)
			v14 = (a4-8)&1 | uint32(*memmap.PtrUint16(0x587000, ((a4-8)>>1)*12+0x4CD18))
			v15 = int32(*memmap.PtrUint32(0x587000, ((a4-8)>>1)*12+0x4CD14))
			if int32(func() uint32 {
				p := *(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*2))
				*p--
				return *p
			}()) <= 0 {
				sub_57E2C0(v5)
			}
			*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*1))))), unsafe.Sizeof(uint16(0))*uintptr(v13))))++
			v16 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(*v5), unsafe.Sizeof(uint32(0))*2)))
			v17 = *(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*6))
			v18 = v15 + int32(*(*int16)(unsafe.Pointer(uintptr(v16 + int32(v13)*4))))
			v19 = int32(*(*uint16)(unsafe.Pointer(uintptr(v16 + int32(v13)*4 + 2)))) << v15
			v20 = uint32(uintptr(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*5)))))
			v21 = int32(32 - uint32(uintptr(unsafe.Pointer(v17))) - uint32(v18))
			v22 = int32(uintptr(unsafe.Pointer(v17))) + v18
			*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*6)) = (*uint32)(unsafe.Pointer(uintptr(v22)))
			*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*5)) = (*uint32)(unsafe.Pointer(uintptr(((v14 | uint32(v19)) << uint32(v21)) | v20)))
			if v22 >= 16 {
				goto LABEL_14
			}
		}
	} else {
		v6 = uint16(a4 + 256)
		if int32(func() uint32 {
			p := *(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*2))
			*p--
			return *p
		}()) <= 0 {
			sub_57E2C0(v5)
		}
		v7 = int32(v6)
		*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*1))))), unsafe.Sizeof(uint16(0))*uintptr(v6))))++
		v8 = *(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*6))
		v9 = uint32(uintptr(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*5)))))
		v10 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(*v5), unsafe.Sizeof(uint32(0))*2)))
		v11 = int32(*(*int16)(unsafe.Pointer(uintptr(v10 + v7*4))))
		v12 = *(*uint16)(unsafe.Pointer(uintptr(v10 + v7*4 + 2)))
		*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*6)) = (*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v8))), v11))))
		*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*5)) = (*uint32)(unsafe.Pointer(uintptr(uint32(int32(v12)<<(32-int32(uint8(uintptr(unsafe.Pointer(v8))))-v11)) | v9)))
		if int32(uintptr(unsafe.Pointer(v8)))+v11 >= 16 {
		LABEL_14:
			*(*uint8)(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*3)))) = uint8(uint32(uintptr(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*5))))) >> 24)
			v30 = *(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*5))
			v31 = (*uint8)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*3))))), 1))
			*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*3)) = (*uint32)(unsafe.Pointer(v31))
			*v31 = *(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v30))), 2))
			v32 = int32(uintptr(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*6))), -int(unsafe.Sizeof(uint32(0))*4))))))
			v33 = (*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*5))))) << 16)))
			*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*3)) = (*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*3))))), 1))))
			*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*6)) = (*uint32)(unsafe.Pointer(uintptr(v32)))
			*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*5)) = v33
			goto LABEL_15
		}
	}
LABEL_15:
	v34 = int32(*memmap.PtrUint32(0x587000, (a5>>9)*8+0x4CDC8) + 9)
	v35 = a5&511 | uint32(int32(*memmap.PtrUint16(0x587000, (a5>>9)*8+0x4CDCC))<<9)
	if v34 <= 16 {
		v47 = *(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*6))
		v48 = int8(32 - int32(uint8(uintptr(unsafe.Pointer(v47)))) - v34)
		v49 = int32(uintptr(unsafe.Pointer(v47))) + v34
		result = int32(v35 << uint32(v48))
		v50 = uint32(uintptr(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*5)))))
		*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*6)) = (*uint32)(unsafe.Pointer(uintptr(v49)))
		*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*5)) = (*uint32)(unsafe.Pointer(uintptr(uint32(result) | v50)))
		if v49 < 16 {
			return result
		}
		goto LABEL_21
	}
	v36 = *(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*6))
	v37 = uint32(uintptr(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*5)))))
	v38 = int32(*memmap.PtrUint32(0x587000, (a5>>9)*8+0x4CDC8) - 7)
	*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*6)) = (*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v36))), v38))))
	v39 = int32((v35 >> 16 << uint32(32-int32(uint8(uintptr(unsafe.Pointer(v36))))-v38)) | v37)
	*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*5)) = (*uint32)(unsafe.Pointer(uintptr(v39)))
	if int32(uintptr(unsafe.Pointer(v36)))+v38 >= 16 {
		*(*uint8)(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*3)))) = *(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v39))), unsafe.Sizeof(int32(0))-1))
		v40 = *(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*5))
		v41 = (*uint8)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*3))))), 1))
		*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*3)) = (*uint32)(unsafe.Pointer(v41))
		*v41 = *(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v40))), 2))
		v42 = int32(uintptr(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*6))), -int(unsafe.Sizeof(uint32(0))*4))))))
		v43 = (*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*5))))) << 16)))
		*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*3)) = (*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*3))))), 1))))
		*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*6)) = (*uint32)(unsafe.Pointer(uintptr(v42)))
		*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*5)) = v43
	}
	v44 = *(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*6))
	v45 = int32(uint32(int32(uint16(v35))<<(16-int32(uint8(uintptr(unsafe.Pointer(v44)))))) | uint32(uintptr(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*5))))))
	result = int32(uintptr(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(v44), unsafe.Sizeof(uint32(0))*4)))))
	*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*5)) = (*uint32)(unsafe.Pointer(uintptr(v45)))
	*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*6)) = (*uint32)(unsafe.Add(unsafe.Pointer(v44), unsafe.Sizeof(uint32(0))*4))
	if int32(uintptr(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(v44), unsafe.Sizeof(uint32(0))*4))))) >= 16 {
	LABEL_21:
		*(*uint8)(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*3)))) = uint8(uint32(uintptr(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*5))))) >> 24)
		v51 = *(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*5))
		v52 = (*uint8)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*3))))), 1))
		*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*3)) = (*uint32)(unsafe.Pointer(v52))
		*v52 = *(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v51))), 2))
		v53 = int32(uintptr(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*6))), -int(unsafe.Sizeof(uint32(0))*4))))))
		result = int32(uint32(uintptr(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*5))))) << 16)
		*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*3)) = (*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*3))))), 1))))
		*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*6)) = (*uint32)(unsafe.Pointer(uintptr(v53)))
		*(**uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof((*uint32)(nil))*5)) = (*uint32)(unsafe.Pointer(uintptr(result)))
	}
	return result
}
func sub_57E7D0(this **uint32) int32 {
	var (
		v1 **uint32
		v2 *uint32
		v3 int32
		v4 *uint32
		v5 *uint8
		v6 int32
		v7 *uint32
		v8 *uint32
		v9 int32
	)
	v1 = this
	if int32(func() uint32 {
		p := *(**uint32)(unsafe.Add(unsafe.Pointer(this), unsafe.Sizeof((*uint32)(nil))*2))
		*p--
		return *p
	}()) <= 0 {
		sub_57E2C0(this)
	}
	*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*1))))), unsafe.Sizeof(uint16(0))*273)))++
	v2 = *(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*6))
	v3 = int32(*(*int16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(*v1), unsafe.Sizeof(uint32(0))*2)) + 1092))))
	*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*5)) = (*uint32)(unsafe.Pointer(uintptr(uint32(int32(*(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(*v1), unsafe.Sizeof(uint32(0))*2)) + 1094))))<<(32-int32(uint8(uintptr(unsafe.Pointer(v2))))-v3)) | uint32(uintptr(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*5))))))))
	*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*6)) = (*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v2))), v3))))
	if int32(uintptr(unsafe.Pointer(v2)))+v3 >= 16 {
		*(*uint8)(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*3)))) = uint8(uint32(uintptr(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*5))))) >> 24)
		v4 = *(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*5))
		v5 = (*uint8)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*3))))), 1))
		*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*3)) = (*uint32)(unsafe.Pointer(v5))
		*v5 = *(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v4))), 2))
		v6 = int32(uintptr(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*6))), -int(unsafe.Sizeof(uint32(0))*4))))))
		v7 = (*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*5))))) << 16)))
		*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*3)) = (*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*3))))), 1))))
		*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*6)) = (*uint32)(unsafe.Pointer(uintptr(v6)))
		*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*5)) = v7
	}
	if int32(uintptr(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*6))))) > 0 {
		for {
			*(*uint8)(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*3)))) = uint8(uint32(uintptr(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*5))))) >> 24)
			v8 = *(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*5))
			v9 = int32(uintptr(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*6))), -int(unsafe.Sizeof(uint32(0))*2))))))
			*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*3)) = (*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*3))))), 1))))
			*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*6)) = (*uint32)(unsafe.Pointer(uintptr(v9)))
			*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*5)) = (*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(v8))) << 8)))
			if v9 <= 0 {
				break
			}
		}
	}
	return int32(uintptr(unsafe.Pointer((*byte)(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*3)))))) - uintptr(unsafe.Pointer((*byte)(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*4)))))))
}
func sub_57E8A0(this *uint32) *uint32 {
	var (
		v1 *uint32
		v2 *uint8
	)
	v1 = this
	sub_57DD90(this)
	v2 = (*uint8)(operator_new(548))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*33)) = uint32(uintptr(unsafe.Pointer(v2)))
	libc.MemCpy(unsafe.Pointer(v2), memmap.PtrOff(0x587000, 315976), 548)
	libc.MemCpy(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*1))), memmap.PtrOff(0x587000, 315848), 128)
	return v1
}
func sub_57E910(this *unsafe.Pointer) {
	var v1 *unsafe.Pointer
	v1 = this
	operator_delete(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(this), unsafe.Sizeof(unsafe.Pointer(nil))*33)))
	sub_57DDC0(v1)
}
func sub_57E970(this *unsafe.Pointer) int32 {
	libc.MemCpy(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(this), unsafe.Sizeof(unsafe.Pointer(nil))*33)), memmap.PtrOff(0x587000, 315976), 548)
	libc.MemCpy(unsafe.Pointer((*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(this), unsafe.Sizeof(unsafe.Pointer(nil))*1))), memmap.PtrOff(0x587000, 315848), 128)
	return sub_57DDD0(this)
}
func sub_57E9A0(this *uint32) *uint32 {
	var v1 *uint32
	v1 = this
	*this = uint32(uintptr(operator_new(0x10000)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*1)) = 0
	sub_57E8A0((*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*2)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*37)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*36)) = 0
	return v1
}
func sub_57EA00(this *unsafe.Pointer) {
	var (
		v1 *unsafe.Pointer
		v2 *unsafe.Pointer
		v4 *unsafe.Pointer
	)
	v1 = this
	v4 = this
	v2 = nil
	if v4 != nil {
		v2 = (*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(unsafe.Pointer(nil))*2))
	}
	sub_57E910(v2)
	operator_delete(*v1)
}
func sub_57EA60(this int32) int32 {
	*(*uint32)(unsafe.Pointer(uintptr(this + 148))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(this + 144))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(this + 4))) = 0
	return sub_57E970((*unsafe.Pointer)(unsafe.Pointer(uintptr(this + 8))))
}
func nox_xxx_nxzDecompressImpl_57EA80(this *uint32, a2 *uint8, a3 *uint32, a4 uint32, a5 *uint32) int32 {
	var (
		v5  *uint8
		v6  *uint32
		v7  uint32
		v8  int32
		v9  int32
		v10 int32
		v11 uint32
		v12 int32
		v13 int32
		v14 int32
		v15 int32
		v16 int32
		v17 int32
		v18 int32
		v19 int32
		v20 *byte
		v21 int16
		v22 int32
		v23 *uint32
		v24 int32
		v25 int32
		v26 int32
		v27 int32
		v28 int32
		v29 int32
		v30 int32
		v31 int32
		v32 int32
		v33 int32
		v34 int32
		v35 int32
		v36 uint32
		v37 int32
		v38 int32
		v39 int32
		v40 int32
		v41 uint32
		v42 int32
		v43 int32
		v44 int32
		v45 *uint8
		v46 int32
		v47 int32
		v48 int32
		v49 int32
		v50 int32
		v51 int32
		v52 unsafe.Pointer
		v53 uint32
		v54 int32
		v55 int32
		v56 *uint8
		v57 int32
		v58 int32
		v59 unsafe.Pointer
		v60 uint32
		v61 *uint8
		v63 int32
		i   int32
		v65 int32
		v66 int32
		v67 int32
		v68 *uint8
		v69 uint32
		v70 int32
		v71 int32
		v72 *uint8
		v73 uint32
		v74 *uint8
		v75 [1096]byte
	)
	v5 = (*uint8)(unsafe.Pointer(uintptr(a4)))
	v6 = this
	v7 = a4 + *a5
	v74 = a2
	v73 = a4
	v69 = a4 + *a5
	v72 = (*uint8)(unsafe.Add(unsafe.Pointer(a2), *a3))
	*(*uint32)(unsafe.Add(unsafe.Pointer(this), unsafe.Sizeof(uint32(0))*37)) = 0
	if a4 >= v7 {
		return 0
	}
	for {
		v8 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*37)))
		if v8 < 4 {
			if uint32(uintptr(unsafe.Pointer(v5))) >= v7 {
				v9 = -1
				*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*37)) = 0
				v63 = -1
				goto LABEL_9
			}
			v10 = int32(uint32(int32(*func() *uint8 {
				p := &v5
				x := *p
				*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), 1))
				return x
			}())<<(24-v8)) | *(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*36)))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*36)) = uint32(v10)
			a4 = uint32(uintptr(unsafe.Pointer(v5)))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*37)) = uint32(v8 + 8)
		}
		v11 = *(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*36))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*36)) = v11 * 16
		v63 = int32(v11 >> 28)
		*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*37)) -= 4
		v9 = int32(v11 >> 28)
	LABEL_9:
		v12 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*uintptr(v9*2+3))))
		if v12 == 0 {
			v13 = int32(*(*int16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*35)) + *(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*uintptr(v9*2+4)))*2))))
			goto LABEL_18
		}
		v14 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*37)))
		if v14 >= v12 {
			goto LABEL_15
		}
		if uint32(uintptr(unsafe.Pointer(v5))) < v7 {
			v16 = int32(uint32(int32(*func() *uint8 {
				p := &v5
				x := *p
				*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), 1))
				return x
			}())<<(24-v14)) | *(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*36)))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*36)) = uint32(v16)
			a4 = uint32(uintptr(unsafe.Pointer(v5)))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*37)) = uint32(v14 + 8)
		LABEL_15:
			v15 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*36)) >> uint32(32-v12))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*36)) <<= uint32(v12)
			*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*37)) -= uint32(v12)
			v9 = v63
			goto LABEL_16
		}
		*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*37)) = 0
		v15 = -1
	LABEL_16:
		v17 = int32(uint32(v15) + *(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*uintptr(v9*2+4))))
		if v17 >= 274 {
			return 0
		}
		v13 = int32(*(*int16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*35)) + uint32(v17*2)))))
	LABEL_18:
		*(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*2)) + uint32(v13*2))))++
		if v13 < 256 {
			if uintptr(unsafe.Pointer(a2)) < uintptr(unsafe.Pointer(v72)) {
				*func() *uint8 {
					p := &a2
					x := *p
					*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), 1))
					return x
				}() = uint8(int8(v13))
				v18 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*1)))
				*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*1)) = uint32(v18 + 1)
				*(*uint8)(unsafe.Pointer(uintptr(uint32(uint16(int16(v18))) + *v6))) = uint8(int8(v13))
				goto LABEL_73
			}
			return 0
		}
		if v13 == 272 {
			sub_57DEA0((*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*2)), (*uint16)(unsafe.Pointer(&v75[0])))
			v19 = 0
			v20 = &v75[0]
			for {
				v21 = int16(*(*uint16)(unsafe.Pointer(v20)))
				v20 = (*byte)(unsafe.Add(unsafe.Pointer(v20), 4))
				*(*uint16)(unsafe.Pointer(uintptr(uint32(v19) + *(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*35))))) = uint16(v21)
				v19 += 2
				if v19 >= 548 {
					break
				}
			}
			v22 = 0
			v23 = (*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*4))
			v70 = 0
			v66 = 16
			for {
				for i = 0; ; i++ {
					v24 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*37)))
					if v24 >= 1 {
						goto LABEL_29
					}
					if uint32(uintptr(unsafe.Pointer(v5))) >= v69 {
						break
					}
					v25 = int32(uint32(int32(*func() *uint8 {
						p := &v5
						x := *p
						*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), 1))
						return x
					}())<<(24-v24)) | *(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*36)))
					*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*36)) = uint32(v25)
					*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*37)) = uint32(v24 + 8)
				LABEL_29:
					v26 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*36)) >> 31)
					v27 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*37)) - 1)
					*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*36)) *= 2
					*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*37)) = uint32(v27)
					if v26 != 0 {
						goto LABEL_32
					}
				}
				*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*37)) = 0
			LABEL_32:
				v22 += i
				*((*uint32)(unsafe.Add(unsafe.Pointer(v23), -int(unsafe.Sizeof(uint32(0))*1)))) = uint32(v22)
				*v23 = uint32(v70)
				v23 = (*uint32)(unsafe.Add(unsafe.Pointer(v23), unsafe.Sizeof(uint32(0))*2))
				v70 += 1 << v22
				if func() int32 {
					p := &v66
					*p--
					return *p
				}() == 0 {
					a4 = uint32(uintptr(unsafe.Pointer(v5)))
					goto LABEL_73
				}
			}
		}
		if v13 == 273 {
			break
		}
		if v13 < 264 {
			v28 = v13 - 256
			goto LABEL_43
		}
		v29 = int32(*memmap.PtrUint32(0x587000, uint32(v13*8)+0x4CC30))
		v30 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*37)))
		if v30 >= v29 {
			goto LABEL_41
		}
		if uint32(uintptr(unsafe.Pointer(v5))) < v69 {
			v32 = int32(uint32(int32(*func() *uint8 {
				p := &v5
				x := *p
				*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), 1))
				return x
			}())<<(24-v30)) | *(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*36)))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*36)) = uint32(v32)
			a4 = uint32(uintptr(unsafe.Pointer(v5)))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*37)) = uint32(v30 + 8)
		LABEL_41:
			v31 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*36)) >> uint32(32-v29))
			v33 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*36)) << uint32(v29))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*37)) -= uint32(v29)
			*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*36)) = uint32(v33)
			goto LABEL_42
		}
		*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*37)) = 0
		v31 = -1
	LABEL_42:
		v28 = int32(uint32(v31) + *memmap.PtrUint32(0x587000, uint32(v13*8)+314420))
	LABEL_43:
		v34 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*37)))
		v67 = v28
		if v34 < 3 {
			if uint32(uintptr(unsafe.Pointer(v5))) >= v69 {
				*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*37)) = 0
				v71 = -1
				goto LABEL_48
			}
			v35 = int32(uint32(int32(*func() *uint8 {
				p := &v5
				x := *p
				*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), 1))
				return x
			}())<<(24-v34)) | *(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*36)))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*36)) = uint32(v35)
			a4 = uint32(uintptr(unsafe.Pointer(v5)))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*37)) = uint32(v34 + 8)
		}
		v36 = *(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*36))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*36)) = v36 * 8
		v71 = int32(v36 >> 29)
		*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*37)) -= 3
	LABEL_48:
		v65 = 0
		v37 = int32(*memmap.PtrUint32(0x587000, uint32(v71*8)+0x4D4B0) + 9)
		if v37 <= 8 {
			goto LABEL_55
		}
		v38 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*37)))
		v37 = int32(*memmap.PtrUint32(0x587000, uint32(v71*8)+0x4D4B0) + 1)
		if v38 >= 8 {
			goto LABEL_53
		}
		if uint32(uintptr(unsafe.Pointer(v5))) < v69 {
			v40 = int32(uint32(int32(*func() *uint8 {
				p := &v5
				x := *p
				*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), 1))
				return x
			}())<<(24-v38)) | *(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*36)))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*36)) = uint32(v40)
			a4 = uint32(uintptr(unsafe.Pointer(v5)))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*37)) = uint32(v38 + 8)
		LABEL_53:
			v41 = *(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*36))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*36)) = v41 << 8
			*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*37)) -= 8
			v39 = int32(v41 >> 24)
			goto LABEL_54
		}
		*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*37)) = 0
		v39 = -1
	LABEL_54:
		v65 = v39 << v37
	LABEL_55:
		v42 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*37)))
		if v42 >= v37 {
			goto LABEL_59
		}
		if uint32(uintptr(unsafe.Pointer(v5))) < v69 {
			*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*36)) |= uint32(int32(*v5) << (24 - v42))
			a4 = uint32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v5), 1)))))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*37)) = uint32(v42 + 8)
		LABEL_59:
			v43 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*36)) >> uint32(32-v37))
			v44 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*36)) << uint32(v37))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*37)) -= uint32(v37)
			*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*36)) = uint32(v44)
			goto LABEL_60
		}
		*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*37)) = 0
		v43 = -1
	LABEL_60:
		v45 = a2
		v46 = int32((*memmap.PtrUint32(0x587000, uint32(v71*8)+0x4D4B4) << 9) + uint32(v65|v43))
		v47 = v67 + 4
		v68 = (*uint8)(unsafe.Add(unsafe.Pointer(a2), v67+4))
		if uintptr(unsafe.Pointer(v68)) > uintptr(unsafe.Pointer(v72)) {
			return 0
		}
		v48 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*1)) - uint32(v46))
		if v47 >= v46 {
			v50 = int32(uint16(int16(v48)))
			if uint32(int32(uint16(int16(v48)))+v46) <= 0x10000 {
				v53 = uint32(v46)
				v52 = unsafe.Pointer(uintptr(*v6 + uint32(v50)))
			} else {
				v51 = int32(0x10000 - uint32(uint16(int16(v48))))
				libc.MemCpy(unsafe.Pointer(a2), unsafe.Pointer(uintptr(*v6+uint32(uint16(int16(v48))))), int(0x10000-uint32(uint16(int16(v48)))))
				v52 = unsafe.Pointer(uintptr(*v6))
				v53 = uint32(v46 - v51)
				v45 = (*uint8)(unsafe.Add(unsafe.Pointer(a2), v51))
			}
			libc.MemCpy(unsafe.Pointer(v45), v52, int(v53))
			v54 = 0
			v55 = v47 - v46
			if v47-v46 > 0 {
				v56 = (*uint8)(unsafe.Add(unsafe.Pointer(a2), v46))
				for {
					v54++
					*(*uint8)(unsafe.Add(unsafe.Pointer(v56), v54-1)) = *(*uint8)(unsafe.Add(unsafe.Pointer(a2), v54-1))
					if v54 >= v55 {
						break
					}
				}
			}
		} else {
			v49 = int32(uint16(int16(v48)))
			if uint32(v49+v47) <= 0x10000 {
				libc.MemCpy(unsafe.Pointer(a2), unsafe.Pointer(uintptr(*v6+uint32(v49))), int(v47))
			} else {
				libc.MemCpy(unsafe.Pointer(a2), unsafe.Pointer(uintptr(*v6+uint32(v49))), int(0x10000-uint32(v49)))
				libc.MemCpy(unsafe.Add(unsafe.Pointer(a2), 0x10000-uint32(v49)), unsafe.Pointer(uintptr(*v6)), int(uint32(v47)-(0x10000-uint32(v49))))
			}
		}
		v57 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*1)) & math.MaxUint16)
		if uint32(v57+v47) <= 0x10000 {
			v61 = a2
			v60 = uint32(v47)
			v59 = unsafe.Pointer(uintptr(*v6 + uint32(v57)))
		} else {
			v58 = int32(0x10000 - uint32(v57))
			libc.MemCpy(unsafe.Pointer(uintptr(*v6+uint32(v57))), unsafe.Pointer(a2), int(0x10000-uint32(v57)))
			v59 = unsafe.Pointer(uintptr(*v6))
			v60 = uint32(v47 - v58)
			v61 = (*uint8)(unsafe.Add(unsafe.Pointer(a2), v58))
		}
		v5 = (*uint8)(unsafe.Pointer(uintptr(a4)))
		libc.MemCpy(v59, unsafe.Pointer(v61), int(v60))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*1)) += uint32(v47)
		a2 = v68
	LABEL_73:
		if uint32(uintptr(unsafe.Pointer(v5))) >= v69 {
			return 0
		}
		v7 = v69
	}
	if a3 != nil {
		*a3 += uint32(int32(uintptr(unsafe.Pointer(v74)) - uintptr(unsafe.Pointer(a2))))
	}
	if a5 != nil {
		*a5 += v73 - uint32(uintptr(unsafe.Pointer(v5)))
	}
	return 1
}
func sub_57F160(this int32, a2 int32, a3 int32) *uint32 {
	var (
		result *uint32
		v4     int32
		v5     *uint8
		v6     int32
		v7     int32
		v8     int32
	)
	result = (*uint32)(unsafe.Pointer(uintptr(this)))
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(this + 24))))
	*(*uint32)(unsafe.Pointer(uintptr(this + 20))) |= uint32(a3 << (32 - v4 - a2))
	*(*uint32)(unsafe.Pointer(uintptr(this + 24))) = uint32(v4 + a2)
	if v4+a2 >= 16 {
		**(**uint8)(unsafe.Pointer(uintptr(this + 12))) = uint8(*(*uint32)(unsafe.Pointer(uintptr(this + 20))) >> 24)
		v5 = (*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(this + 12))) + 1)))
		v6 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*5)) >> 16)
		*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*3)) = uint32(uintptr(unsafe.Pointer(v5)))
		*v5 = uint8(int8(v6))
		v7 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*5)))
		v8 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*6)) - 16)
		*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*3))++
		*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*6)) = uint32(v8)
		*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*5)) = uint32(v7 << 16)
	}
	return result
}
func sub_57F1D0(a1 *float2) int8 {
	var (
		v1 int8
		v2 int32
		v3 float64
		v4 float64
		v5 float64
		v7 uint8
	)
	v1 = 0
	v2 = int(a1.field_0)
	v7 = uint8(int8(int(a1.field_4) % 23))
	v3 = float64(uint8(int8(v2 % 23)))
	if v3 >= 11.5 {
		v4 = float64(v7)
		if v4 >= 11.5 {
			v1 = 4
		}
		if v4 <= 11.5 {
			v1 |= 1
		}
	}
	if v3 <= 11.5 {
		v5 = float64(v7)
		if v5 >= 11.5 {
			v1 |= 8
		}
		if v5 <= 11.5 {
			v1 |= 2
		}
	}
	return v1
}
func sub_57F2A0(a1 *float2, a2 int32, a3 int32) int32 {
	var (
		v3     int32
		v4     int32
		result int32
		v6     int32
		v7     float32
		v8     float32
	)
	v7 = float32(float64(a1.field_0) - float64(a2*23))
	v3 = int(v7)
	v8 = float32(float64(a1.field_4) - float64(a3*23))
	v4 = int(v8)
	if v3 <= v4 {
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v4))), 0)) = uint8(int8(bool2int(22-v3 <= v4)))
		v6 = v4 - 1
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v6))), 0)) = uint8(int8(v6 & 253))
		result = v6 + 3
	} else {
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v4))), 0)) = uint8(int8(bool2int(22-v3 <= v4)))
		result = v4 + 1
	}
	return result
}
func nullsub_2() {
}
func nullsub_4(a1 uint32, a2 uint32, a3 uint32, a4 uint32) {
}
func nox_xxx_set_sage(a1 uint32) {
}
func nullsub_5() {
}
func nullsub_10(a1 uint32) {
}
func sub_42CC50(this *unsafe.Pointer) int32 {
	return sub_42C770(this)
}
func nox_xxx_j_resetNPCRenderData_49A2E0() {
	nox_alloc_npcs()
}
func nullsub_16() {
}
func nullsub_12() {
}
func nullsub_17() {
}
func nullsub_18() {
}
func nullsub_11() {
}
func nullsub_13() {
}
func nullsub_15() {
}
func nullsub_14() {
}
func nullsub_3() {
}
func nullsub_6() {
}
func nullsub_19() {
}
func nox_xxx_j_inventoryNameSignInit_467460() int32 {
	return nox_xxx_inventoryNameSignInit_4671E0()
}
func nox_alloc_npcs_2() {
	nox_alloc_npcs()
}
func nullsub_8(a1 int32, a2 int32) int32 {
	return 0
}
func nullsub_27(a1 uint32) {
}
func nullsub_28(a1 uint32) {
}
func nullsub_30(a1 uint32) {
}
func nullsub_36() {
}
func nullsub_29() {
}
func nullsub_35(a1 uint32, a2 uint32) {
}
func nullsub_34(a1 uint32, a2 uint32, a3 uint32) {
}
func nullsub_20() {
}
func nullsub_24(a1 uint32) {
}
func nullsub_31(a1 uint32) {
}
func nullsub_22() {
}
func nullsub_23() {
}
func nullsub_9(a1 uint32) {
}
func nullsub_33(a1 uint32, a2 uint32, a3 uint32) {
}
func sub_558800(this *func() int32) int32 {
	return (cgoAsFunc(*(*uint32)(unsafe.Add(unsafe.Pointer(this), unsafe.Sizeof(uintptr(0))*385)), (*func() int32)(nil)))()
}
func sub_558810(this *func() int32) int32 {
	return (cgoAsFunc(*(*uint32)(unsafe.Add(unsafe.Pointer(this), unsafe.Sizeof(uintptr(0))*386)), (*func() int32)(nil)))()
}
func nox_xxx_j_allocHitArray_511840() {
	nox_xxx_allocHitArray_5486D0()
}
