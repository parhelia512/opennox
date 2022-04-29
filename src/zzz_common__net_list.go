package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/alloc"
	"unsafe"
)

const NOX_NETBUF_MAX_PACKETS = 512
const NOX_NETBUF_MAX_SIZE = 2032

type nox_net_list_item_t struct {
	buf  *uint8
	size uint32
	prev *nox_net_list_item_t
	next *nox_net_list_item_t
}
type nox_net_list_t struct {
	first   *nox_net_list_item_t
	last    *nox_net_list_item_t
	field_2 uint32
	alloc   *nox_alloc_class
	count   uint32
	size    uint32
	field_6 uint32
	field_7 uint32
}
type nox_net_lists_buf_t struct {
	buf [2048]uint8
	cur uint32
}

var nox_net_lists [3][32]*nox_net_list_t = [3][32]*nox_net_list_t{}
var nox_net_lists_buf [2048]uint8 = [2048]uint8{}
var nox_net_lists_buf_arr [3][32]nox_net_lists_buf_t = [3][32]nox_net_lists_buf_t{}

func nox_netlist_count_420BC0(p *nox_net_list_t) int32 {
	if p == nil {
		return 0
	}
	return int32(p.count)
}
func nox_netlist_size_420BD0(p *nox_net_list_t) int32 {
	if p == nil {
		return 0
	}
	return int32(p.size)
}
func nox_netlist_countByInd_40E9D0(ind1 int32, ind2 int32) int32 {
	return nox_netlist_count_420BC0(nox_net_lists[ind2][ind1])
}
func nox_netlist_sizeByInd_40E9F0(ind1 int32, ind2 int32) int32 {
	return nox_netlist_size_420BD0(nox_net_lists[ind2][ind1])
}
func nox_netlist_countByInd2_40F0B0(ind int32) int32 {
	return nox_netlist_countByInd_40E9D0(ind, 2)
}
func nox_netlist_sizeByInd2_40F0D0(ind int32) int32 {
	return nox_netlist_sizeByInd_40E9F0(ind, 2)
}
func nox_netlist_newMsgList_420890(cnt int32) *nox_net_list_t {
	var p *nox_net_list_t = new(nox_net_list_t)
	if p == nil {
		return nil
	}
	var aclass *nox_alloc_class = nox_new_alloc_class(CString("CreateMsgList"), int32(unsafe.Sizeof(nox_net_list_item_t{})), cnt)
	if aclass == nil {
		return nil
	}
	p.first = nil
	p.last = nil
	p.field_2 = 0
	p.alloc = aclass
	p.count = 0
	p.size = 0
	p.field_6 = 0
	p.field_7 = 0
	return p
}
func nox_netlist_freeMsgList_4208F0(p *nox_net_list_t) {
	nox_free_alloc_class(p.alloc)
	alloc.Free(unsafe.Pointer(p))
}
func nox_netlist_init_40EA10() {
	for i := int32(0); i < NOX_PLAYERINFO_MAX; i++ {
		nox_net_lists[1][i] = nox_netlist_newMsgList_420890(NOX_NETBUF_MAX_PACKETS)
		nox_net_lists[0][i] = nil
		nox_net_lists[2][i] = nox_netlist_newMsgList_420890(NOX_NETBUF_MAX_PACKETS)
	}
	nox_net_lists[0][NOX_PLAYERINFO_MAX-1] = nox_netlist_newMsgList_420890(NOX_NETBUF_MAX_PACKETS)
}
func nox_netlist_free_40EA70() {
	for i := int32(0); i < NOX_PLAYERINFO_MAX; i++ {
		if nox_net_lists[1][i] != nil {
			nox_netlist_freeMsgList_4208F0(nox_net_lists[1][i])
		}
		if nox_net_lists[0][i] != nil {
			nox_netlist_freeMsgList_4208F0(nox_net_lists[0][i])
		}
		if nox_net_lists[2][i] != nil {
			nox_netlist_freeMsgList_4208F0(nox_net_lists[2][i])
		}
	}
}
func nox_netlist_checkSizes_40EAC0(ind1 int32, ind2 int32, sz int32) bool {
	if ind2 == 1 {
		var psz int32 = nox_netlist_sizeByInd_40E9F0(ind1, 1)
		if psz+sz+nox_netlist_sizeByInd_40E9F0(ind1, 2) > NOX_NETBUF_MAX_SIZE {
			return false
		}
	} else {
		if sz+nox_netlist_sizeByInd_40E9F0(ind1, ind2) > NOX_NETBUF_MAX_SIZE {
			return false
		}
	}
	return nox_netlist_count_420BC0(nox_net_lists[ind2][ind1]) < NOX_NETBUF_MAX_PACKETS
}
func nox_netlist_checkSizesExt_40EB60(ind1 int32, ind2 int32, sz int32, sz2 int32) bool {
	var (
		p   *nox_net_list_t = nox_net_lists[ind2][ind1]
		psz int32           = nox_netlist_size_420BD0(p)
	)
	if sz+psz+sz2 > NOX_NETBUF_MAX_SIZE {
		return false
	}
	return nox_netlist_count_420BC0(p) < NOX_NETBUF_MAX_PACKETS
}
func nox_netlist_sendByInd_40EC30(ind1 int32, ind2 int32, buf *uint8, sz int32) *uint8 {
	if buf == nil || sz <= 0 {
		return nil
	}
	var p *nox_net_lists_buf_t = &nox_net_lists_buf_arr[ind2+1][ind1]
	var i int32 = int32(p.cur)
	if i+sz > 2048 {
		return nil
	}
	libc.MemCpy(unsafe.Pointer(&p.buf[i]), unsafe.Pointer(buf), int(sz))
	p.cur += uint32(sz)
	return &p.buf[i]
}
func nox_netlist_add_420940(list *nox_net_list_t, buf int32, sz int32, append bool) int32 {
	var item *nox_net_list_item_t = (*nox_net_list_item_t)(nox_alloc_class_new_obj_zero(list.alloc))
	if item == nil {
		return 0
	}
	item.buf = (*uint8)(unsafe.Pointer(uintptr(buf)))
	item.size = uint32(sz)
	if append {
		item.prev = list.last
		item.next = nil
		if list.last != nil {
			list.last.next = item
		} else {
			list.first = item
		}
		list.last = item
	} else {
		item.prev = nil
		item.next = list.first
		if list.first != nil {
			list.first.prev = item
		} else {
			list.last = item
		}
		list.first = item
	}
	list.count++
	list.size += uint32(sz)
	return 1
}
func nox_netlist_addToMsgListCli_40EBC0(ind1 int32, ind2 int32, buf *uint8, sz int32) int32 {
	var p *nox_net_list_t = nox_net_lists[ind2][ind1]
	if sz <= 0 {
		return 1
	}
	if !nox_netlist_checkSizes_40EAC0(ind1, ind2, sz) {
		return 0
	}
	var out *uint8 = nox_netlist_sendByInd_40EC30(ind1, ind2, buf, sz)
	if out == nil {
		return 0
	}
	nox_netlist_add_420940(p, int32(uintptr(unsafe.Pointer(out))), sz, true)
	return 1
}
func nox_netlist_clientSend_0_40ECA0(ind1 int32, ind2 int32, buf *uint8, sz int32, sz2 int32) int32 {
	var p *nox_net_list_t = nox_net_lists[ind2][ind1]
	if sz <= 0 {
		return 1
	}
	if !nox_netlist_checkSizesExt_40EB60(ind1, ind2, sz, sz2) {
		return 0
	}
	var out *uint8 = nox_netlist_sendByInd_40EC30(ind1, ind2, buf, sz)
	if out == nil {
		return 0
	}
	nox_netlist_add_420940(p, int32(uintptr(unsafe.Pointer(out))), sz, true)
	return 1
}
func nox_xxx_netBufsSetZero_40ED40(ind int32, ind2 int32) {
	nox_net_lists_buf_arr[ind2+1][ind].cur = 0
}
func sub_40F040(ind int32) {
	nox_net_lists_buf_arr[0][ind].cur = 0
}
func nox_netlist_resetList_420830(p *nox_net_list_t) {
	p.first = nil
	p.last = nil
	p.field_2 = 0
	p.count = 0
	p.size = 0
	nox_alloc_class_free_all(p.alloc)
}
func nox_netlist_resetByInd_40ED10(ind1 int32, ind2 int32) {
	var p *nox_net_list_t = nox_net_lists[ind2][ind1]
	if p != nil {
		nox_netlist_resetList_420830(p)
		nox_xxx_netBufsSetZero_40ED40(ind1, ind2)
	}
}
func nox_netlist_initPlayerBufs_40F020(ind int32) {
	nox_netlist_resetList_420830(nox_net_lists[2][ind])
	sub_40F040(ind)
}
func nox_netlist_resetAllInList_40EE90(ind int32) {
	for i := int32(0); i < NOX_PLAYERINFO_MAX; i++ {
		nox_netlist_resetByInd_40ED10(i, ind)
	}
}
func nox_netlist_resetAll_40EE60() {
	for i := int32(0); i < NOX_PLAYERINFO_MAX; i++ {
		nox_netlist_resetByInd_40ED10(i, 1)
		nox_netlist_resetByInd_40ED10(i, 0)
		nox_netlist_initPlayerBufs_40F020(i)
	}
}
func nox_netlist_free_item_4209C0(list *nox_net_list_t, item *nox_net_list_item_t) {
	nox_alloc_class_free_obj_first(list.alloc, unsafe.Pointer(item))
}
func nox_netlist_get_420A90(list *nox_net_list_t, outSz *uint32) *uint32 {
	if list.first == nil {
		return nil
	}
	var item *nox_net_list_item_t = list.first
	*outSz = item.size
	list.count--
	list.size -= item.size
	var next *nox_net_list_item_t = item.next
	if next != nil {
		next.prev = item.prev
	} else {
		list.last = item.prev
	}
	var prev *nox_net_list_item_t = item.prev
	if prev != nil {
		prev.next = item.next
	} else {
		list.first = item.next
	}
	var buf *uint8 = item.buf
	nox_netlist_free_item_4209C0(list, item)
	return (*uint32)(unsafe.Pointer(buf))
}
func nox_netlist_copyPacketList_40ED60(ind1 int32, ind2 int32, outSz *uint32) *uint8 {
	var (
		list *nox_net_list_t = nox_net_lists[ind2][ind1]
		cnt  uint32          = 0
	)
	*(*[2048]uint8)(unsafe.Pointer(&nox_net_lists_buf[0])) = [2048]uint8{}
	for {
		var (
			sz  uint32 = 0
			src *uint8 = (*uint8)(unsafe.Pointer(nox_netlist_get_420A90(list, &sz)))
		)
		if src == nil {
			break
		}
		if cnt+sz > uint32(unsafe.Sizeof([2048]uint8{})) {
			break
		}
		libc.MemCpy(unsafe.Pointer(&nox_net_lists_buf[cnt]), unsafe.Pointer(src), int(sz))
		cnt += sz
	}
	*outSz = cnt
	return &nox_net_lists_buf[0]
}
func nox_netlist_getInd_40EEB0(ind1 int32, ind2 int32, outSz *uint32) *uint8 {
	return (*uint8)(unsafe.Pointer(nox_netlist_get_420A90(nox_net_lists[ind2][ind1], outSz)))
}
func sub_40EEF0(ind int32, a2 int32) bool {
	if uint32(a2+nox_netlist_size_420BD0(nox_net_lists[2][ind])) > NOX_NETBUF_MAX_SIZE {
		return false
	}
	return nox_netlist_count_420BC0(nox_net_lists[2][ind]) < NOX_NETBUF_MAX_PACKETS
}
func nox_netlist_getByInd2_40F080(ind int32, outSz *uint32) *uint8 {
	return (*uint8)(unsafe.Pointer(nox_netlist_get_420A90(nox_net_lists[2][ind], outSz)))
}
func nox_netlist_findAndFreeBuf_4209E0(list *nox_net_list_t, buf *uint8) {
	if list.first == nil {
		return
	}
	var item *nox_net_list_item_t = nil
	for p := (*nox_net_list_item_t)(list.first); p != nil; p = p.next {
		if p.buf == buf {
			item = p
			break
		}
	}
	if item == nil {
		return
	}
	list.count--
	list.size -= item.size
	var next *nox_net_list_item_t = item.next
	if next != nil {
		next.prev = item.prev
	} else {
		list.last = item.prev
	}
	var prev *nox_net_list_item_t = item.prev
	if prev != nil {
		prev.next = item.next
	} else {
		list.first = item.next
	}
	nox_netlist_free_item_4209C0(list, item)
}
func nox_netlist_findAndFreeBuf_40F000(ind int32, buf *uint8) {
	nox_netlist_findAndFreeBuf_4209E0(nox_net_lists[2][ind], buf)
}
func nox_xxx_netlistAdd_40EFA0(ind int32, buf *uint8, sz int32) *byte {
	if buf == nil || sz <= 0 {
		return nil
	}
	var p *nox_net_lists_buf_t = &nox_net_lists_buf_arr[0][ind]
	var i int32 = int32(p.cur)
	if i+sz > 2048 {
		return nil
	}
	libc.MemCpy(unsafe.Pointer(&p.buf[i]), unsafe.Pointer(buf), int(sz))
	p.cur += uint32(sz)
	return (*byte)(unsafe.Pointer(&p.buf[i]))
}
func sub_40F060() {
	for i := int32(0); i < NOX_PLAYERINFO_MAX; i++ {
		nox_netlist_initPlayerBufs_40F020(i)
	}
}
func sub_420A60(list *nox_net_list_t, fnc func(uint32, int32) int32, a3 int32) {
	if list.first == nil {
		return
	}
	for p := (*nox_net_list_item_t)(list.first); p != nil; p = p.next {
		if fnc(uint32(uintptr(unsafe.Pointer(p.buf))), a3) != 0 {
			break
		}
	}
}
func nox_netlist_forEach2_40F0F0(ind int32, fnc func(uint32, int32) int32, a3 int32) {
	if fnc != nil && a3 != 0 {
		sub_420A60(nox_net_lists[2][ind], fnc, a3)
	}
}
func nox_netlist_addToMsgListSrv_40EF40(ind int32, buf *uint8, sz int32) bool {
	if sz <= 0 {
		return true
	}
	var out *uint8
	var list *nox_net_list_t = nox_net_lists[2][ind]
	if !sub_40EEF0(ind, sz) || (func() *uint8 {
		out = (*uint8)(unsafe.Pointer(nox_xxx_netlistAdd_40EFA0(ind, buf, sz)))
		return out
	}()) == nil {
		var (
			p    *nox_net_lists_buf_t = &nox_net_lists_buf_arr[0][ind]
			len1 uint32
			len2 uint32
		)
		len1 = list.first.size
		len2 = list.first.next.size
		if ind == 31 {
			nox_netlist_receiveCli_494E90(ind)
		} else {
			nox_xxx_netSendReadPacket_5528B0(*((*uint32)(unsafe.Add(unsafe.Pointer(&nox_common_playerInfoFromNum_417090(ind).field_0), unsafe.Sizeof(uint32(0))*516)))+1, 0)
		}
		p.cur = len1 + len2
		nox_netlist_add_420940(list, int32(uintptr(unsafe.Pointer(&p.buf[0]))), int32(len1), true)
		nox_netlist_add_420940(list, int32(uintptr(unsafe.Pointer(&p.buf[len1]))), int32(len2), true)
		out = (*uint8)(unsafe.Pointer(nox_xxx_netlistAdd_40EFA0(ind, buf, sz)))
	}
	if out == nil {
		return false
	}
	nox_netlist_add_420940(list, int32(uintptr(unsafe.Pointer(out))), sz, true)
	return true
}
func nox_netlist_copyPacketList2_40F120(ind int32, outSz *uint32) *uint8 {
	var (
		list *nox_net_list_t = nox_net_lists[2][ind]
		cnt  uint32          = 0
	)
	*(*[2048]uint8)(unsafe.Pointer(&nox_net_lists_buf[0])) = [2048]uint8{}
	for {
		var (
			sz  uint32 = 0
			src *uint8 = (*uint8)(unsafe.Pointer(nox_netlist_get_420A90(list, &sz)))
		)
		if src == nil {
			break
		}
		if cnt+sz > uint32(unsafe.Sizeof([2048]uint8{})) {
			nox_netlist_add_420940(list, int32(uintptr(unsafe.Pointer(src))), int32(sz), false)
			break
		}
		libc.MemCpy(unsafe.Pointer(&nox_net_lists_buf[cnt]), unsafe.Pointer(src), int(sz))
		cnt += sz
	}
	*outSz = cnt
	return &nox_net_lists_buf[0]
}
