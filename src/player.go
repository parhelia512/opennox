package opennox

/*
#include "GAME1.h"
#include "GAME1_1.h"
#include "GAME1_2.h"
#include "GAME2.h"
#include "GAME2_1.h"
#include "GAME3_2.h"
#include "GAME3_3.h"
#include "GAME4.h"
#include "GAME4_1.h"
#include "common__net_list.h"
#include "defs.h"
extern unsigned int nox_gameDisableMapDraw_5d4594_2650672;
extern unsigned int dword_5d4594_1046492;
extern unsigned int dword_5d4594_2650652;
extern unsigned int nox_player_netCode_85319C;
extern nox_object_t* nox_xxx_host_player_unit_3843628;
void nox_xxx_WideScreenDo_515240(bool enable);
static void nox_xxx_printToAll_4D9FD0_go(int a1, wchar_t* str) {
	nox_xxx_printToAll_4D9FD0(a1, str);
}
*/
import "C"
import (
	"encoding"
	"encoding/binary"
	"fmt"
	"image"
	"math"
	"unsafe"

	"github.com/noxworld-dev/opennox-lib/noxnet"
	"github.com/noxworld-dev/opennox-lib/object"
	"github.com/noxworld-dev/opennox-lib/player"
	"github.com/noxworld-dev/opennox-lib/script"
	"github.com/noxworld-dev/opennox-lib/things"
	"github.com/noxworld-dev/opennox-lib/types"

	"github.com/noxworld-dev/opennox/v1/common/alloc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
)

const NOX_PLAYERINFO_MAX = int(C.NOX_PLAYERINFO_MAX)

func (s *Server) allocPlayers() {
	p, _ := alloc.Calloc(NOX_PLAYERINFO_MAX, unsafe.Sizeof(Player{}))
	s.players = unsafe.Slice((*Player)(p), NOX_PLAYERINFO_MAX)
}

func clientPlayer() *Player {
	return noxServer.getPlayerByID(clientPlayerNetCode())
}

func clientPlayerNetCode() int {
	return int(C.nox_player_netCode_85319C)
}

func clientSetPlayerNetCode(id int) {
	C.nox_player_netCode_85319C = C.uint(id)
}

func (s *Server) resetAllPlayers() { // nox_xxx_cliResetAllPlayers_416E30
	for i := range s.players {
		s.players[i] = Player{}
	}
}

func (s *Server) playerFirst() *Player {
	for i := range s.players {
		p := &s.players[i]
		if p.IsActive() {
			return p
		}
	}
	return nil
}

func (s *Server) playerNext(it *Player) *Player {
	if it == nil {
		return nil
	}
	for i := it.Index() + 1; i < len(s.players); i++ {
		p := &s.players[i]
		if p.IsActive() {
			return p
		}
	}
	return nil
}

//export nox_common_playerInfoGetFirst_416EA0
func nox_common_playerInfoGetFirst_416EA0() *C.nox_playerInfo {
	return noxServer.playerFirst().C()
}

//export nox_common_playerInfoGetNext_416EE0
func nox_common_playerInfoGetNext_416EE0(it *C.nox_playerInfo) *C.nox_playerInfo {
	return noxServer.playerNext(asPlayer(it)).C()
}

//export nox_common_playerInfoCount_416F40
func nox_common_playerInfoCount_416F40() C.int {
	return C.int(noxServer.cntPlayers())
}

//export nox_common_playerInfoNew_416F60
func nox_common_playerInfoNew_416F60(id C.int) *C.nox_playerInfo {
	return noxServer.newPlayerInfo(int(id)).C()
}

func (s *Server) playerResetInd(ind int) *Player {
	p := &s.players[ind]
	p.Reset(ind)
	return p
}

//export nox_common_playerInfoResetInd_417000
func nox_common_playerInfoResetInd_417000(ind C.int) *C.nox_playerInfo {
	return noxServer.playerResetInd(int(ind)).C()
}

//export nox_common_playerInfoGetByID_417040
func nox_common_playerInfoGetByID_417040(id C.int) *C.nox_playerInfo {
	return noxServer.getPlayerByID(int(id)).C()
}

//export nox_common_playerInfoFromNum_417090
func nox_common_playerInfoFromNum_417090(ind C.int) *C.nox_playerInfo {
	return noxServer.getPlayerByInd(int(ind)).C()
}

//export nox_common_playerInfoFromNumRaw
func nox_common_playerInfoFromNumRaw(ind C.int) *C.nox_playerInfo {
	return noxServer.players[ind].C()
}

//export nox_xxx_playerNew_4DD320
func nox_xxx_playerNew_4DD320(ind C.int, data *C.uchar) C.int {
	return C.int(noxServer.newPlayerFromPacket(int(ind), unsafe.Slice((*byte)(unsafe.Pointer(data)), 153)))
}

func (s *Server) newPlayerInfo(id int) *Player {
	if p := s.getPlayerByID(id); p != nil {
		return p
	}
	for i := range s.players {
		p := &s.players[i]
		if !p.IsActive() {
			p.Reset(i)
			p.netCode = C.uint(id)
			return p
		}
	}
	return nil
}

func (p *Player) Reset(ind int) { // nox_common_playerInfoReset_416FD0
	*p = Player{
		playerInd:  C.uchar(ind),
		active:     1,
		field_3648: 4,
	}
}

//export nox_xxx_playerDisconnByPlrID_4DEB00
func nox_xxx_playerDisconnByPlrID_4DEB00(id C.int) {
	if p := noxServer.getPlayerByInd(int(id)); p != nil {
		p.Disconnect(4)
	}
}

//export nox_xxx_playerCallDisconnect_4DEAB0
func nox_xxx_playerCallDisconnect_4DEAB0(ind C.int, v C.char) *C.char {
	noxServer.getPlayerByInd(int(ind)).Disconnect(int(v))
	return nil
}

func getPlayerClass() player.Class {
	return player.Class(memmap.Uint8(0x85B3FC, 12254))
}

func asPlayer(p *C.nox_playerInfo) *Player {
	return (*Player)(p)
}

func BlindPlayers(blind bool) {
	noxServer.nox_xxx_netMsgFadeBegin_4D9800(!blind, false)
}

func CinemaPlayers(cinema bool) {
	nox_xxx_WideScreenDo_515240(C.bool(cinema))
}

func PrintToPlayers(text string) {
	cstr, free := CWString(text)
	defer free()
	C.nox_xxx_printToAll_4D9FD0_go(0, cstr)
}

func HostPlayer() *Player {
	// TODO: better way
	for _, p := range noxServer.getPlayers() {
		if p.IsHost() {
			return p
		}
	}
	return nil
}

var _ noxObject = (*Player)(nil) // proxies Unit

type PlayerInfo C.nox_playerInfo2

func (p *PlayerInfo) field(off uintptr) unsafe.Pointer {
	return unsafe.Pointer(uintptr(unsafe.Pointer(p)) + off)
}

func (p *PlayerInfo) PlayerClass() player.Class {
	if p == nil {
		return 0
	}
	return player.Class(p.playerClass)
}

func (p *PlayerInfo) IsFemale() bool {
	if p == nil {
		return false
	}
	return p.isFemale != 0
}

func (p *PlayerInfo) Name() string {
	return GoWString(&p.name[0])
}

func (p *PlayerInfo) SetName(v string) {
	WStrCopy(&p.name[0], 25, v)
}

func (p *PlayerInfo) SetNameSuff(v string) {
	WStrCopy((*C.wchar_t)(p.field(89)), 4, v)
}

func (p *PlayerInfo) Field2235() uint32 {
	return *(*uint32)(p.field(50))
}

func (p *PlayerInfo) Field2239() uint32 {
	return *(*uint32)(p.field(54))
}

func (p *PlayerInfo) SetField2235(v uint32) {
	*(*uint32)(p.field(50)) = v
}

func (p *PlayerInfo) SetField2239(v uint32) {
	*(*uint32)(p.field(54)) = v
}

func (p *PlayerInfo) SetField2256(v uint16) {
	*(*uint16)(p.field(71)) = v
}

func (p *PlayerInfo) SetField2262(v uint16) {
	*(*uint16)(p.field(77)) = v
}

type Player C.nox_playerInfo

func (p *Player) getServer() *Server {
	return noxServer // TODO: attach to object
}

func (p *Player) field(off uintptr) unsafe.Pointer {
	return unsafe.Add(unsafe.Pointer(p), off)
}

func (p *Player) Pos() types.Pointf {
	if p == nil {
		return types.Pointf{}
	}
	u := p.Unit()
	if u == nil {
		return types.Pointf{}
	}
	return u.Pos()
}

func (p *Player) SetPos(pos types.Pointf) {
	if p == nil {
		return
	}
	u := p.Unit()
	if u == nil {
		return
	}
	u.SetPos(pos)
}

func (p *Player) OrigName() string {
	return p.Info().Name()
}

func (p *Player) SetName(v string) {
	WStrCopy(&p.name_final[0], 30, v) // TODO: size is a wild guess
}

func (p *Player) Name() string {
	return GoWStringN(&p.name_final[0], 30) // TODO: size is a wild guess
}

func (p *Player) Serial() string {
	return GoStringN(&p.serial[0], 22)
}

func (p *Player) SetSerial(v string) {
	StrCopy(&p.serial[0], 22, v)
}

func (p *Player) Field2096() string {
	return GoStringN(&p.field_2096[0], 12)
}

func (p *Player) SetField2096(v string) {
	StrCopy(&p.field_2096[0], 12, v)
}

func (p *Player) String() string {
	return fmt.Sprintf("Player(%q)", p.Name())
}

func (p *Player) Gold() int {
	return int(p.gold)
}

func (p *Player) IsHost() bool {
	// TODO: better way
	return p.UnitC() == HostPlayerUnit()
}

func (p *Player) Print(text string) {
	nox_xxx_netSendLineMessage_4D9EB0(p.UnitC(), text)
}

func (p *Player) Blind(blind bool) {
	nox_xxx_netMsgFadeBeginPlayer(C.int(p.Index()), C.int(bool2int(!blind)), 0)
}

func (p *Player) Cinema(v bool) {
	// TODO: only works for everyone for now
	CinemaPlayers(v)
}

func (p *Player) CObj() *nox_object_t {
	u := p.UnitC()
	if u == nil {
		return nil
	}
	return u.CObj()
}

func (p *Player) GetObject() script.Object {
	u := p.Unit()
	if u == nil {
		return nil
	}
	return u
}

func (p *Player) Unit() script.Unit {
	if p == nil {
		return nil
	}
	return p.UnitC()
}

func HostPlayerUnit() *Unit {
	return asUnit(unsafe.Pointer(C.nox_xxx_host_player_unit_3843628))
}

func (p *Player) C() *C.nox_playerInfo {
	return (*C.nox_playerInfo)(p)
}

func (p *Player) Index() int {
	if p == nil {
		return -1
	}
	return int(p.playerInd)
}

func (p *Player) NetCode() int {
	if p == nil {
		return -1
	}
	return int(p.netCode)
}

func (p *Player) IsActive() bool {
	return p != nil && p.active != 0
}

func (p *Player) UnitC() *Unit {
	if p == nil {
		return nil
	}
	return asUnitC(p.playerUnit)
}

func (p *Player) Info() *PlayerInfo {
	if p == nil {
		return nil
	}
	return (*PlayerInfo)(p.field(2185)) // inaccessible due to alignment issues
}

func (p *Player) PlayerClass() player.Class {
	if p == nil {
		return 0
	}
	return p.Info().PlayerClass()
}

func (p *Player) Disconnect(v int) {
	if !p.IsActive() {
		return
	}
	nox_xxx_playerDisconnFinish_4DE530(C.int(p.Index()), C.char(v))
	nox_xxx_playerForceDisconnect_4DE7C0(C.int(p.Index()))
	p.getServer().nox_xxx_netStructReadPackets2_4DEC50(p.Index())
}

func (p *Player) CameraTarget() *Object {
	if p == nil {
		return nil
	}
	return asObjectC(p.camera_follow)
}

func (p *Player) CameraUnlock() {
	if p == nil {
		return
	}
	p.camera_follow = nil
}

func (p *Player) CameraFollow(obj *Object) {
	if p == nil {
		return
	}
	p.camera_follow = obj.CObj()
}

func (p *Player) CameraToggle(obj *Object) {
	if p == nil {
		return
	}
	if p.CameraTarget() == obj {
		p.CameraUnlock()
	} else {
		p.CameraFollow(obj)
	}
}

func (p *Player) GoObserver(notify, keepPlayer bool) bool {
	if p == nil {
		return true
	}
	return nox_xxx_playerGoObserver_4E6860(p.C(), C.int(bool2int(notify)), C.int(bool2int(keepPlayer))) != 0
}

func (s *Server) cntPlayers() (n int) {
	for i := range s.players {
		p := &s.players[i]
		if p.IsActive() {
			n++
		}
	}
	return n
}

func (s *Server) getAllPlayerStructs() (out []*Player) {
	for i := range s.players {
		p := &s.players[i]
		out = append(out, p)
	}
	return out
}

func (s *Server) getPlayers() (out []*Player) {
	for i := range s.players {
		p := &s.players[i]
		if p.IsActive() {
			out = append(out, p)
		}
	}
	return out
}

func (s *Server) getPlayerUnits() (out []*Unit) {
	for _, p := range s.getPlayers() {
		if u := p.UnitC(); u != nil {
			out = append(out, u)
		}
	}
	return out
}

func (s *Server) getPlayerByInd(i int) *Player {
	if i < 0 || i >= len(s.players) {
		return nil
	}
	p := &s.players[i]
	if !p.IsActive() {
		return nil
	}
	p.playerInd = C.uchar(i)
	return p
}

func (s *Server) getPlayerByID(id int) *Player {
	for i := range s.players {
		p := &s.players[i]
		if p.IsActive() && int(p.netCode) == id {
			return p
		}
	}
	return nil
}

func (s *Server) hasPlayerUnits() bool {
	for i := range s.players {
		p := &s.players[i]
		if p.UnitC() != nil {
			return true
		}
	}
	return false
}

func nox_xxx_netNewPlayerMakePacket_4DDA90(buf []byte, pl *Player) {
	buf[0] = byte(noxnet.MSG_NEW_PLAYER)
	binary.LittleEndian.PutUint16(buf[1:], uint16(pl.NetCode()))
	binary.LittleEndian.PutUint16(buf[100:], uint16(pl.field_2136))
	binary.LittleEndian.PutUint16(buf[102:], uint16(pl.field_2140))
	binary.LittleEndian.PutUint32(buf[104:], uint32(pl.field_0))
	binary.LittleEndian.PutUint32(buf[108:], uint32(pl.field_4))
	buf[116] = byte(pl.field_2152)
	buf[117] = byte(pl.field_2156)
	buf[118] = byte(bool2int(pl.field_3676 == 3))
	binary.LittleEndian.PutUint32(buf[112:], uint32(pl.field_3680)&0x423)
	StrCopyBytes(buf[119:], pl.Field2096())
	*(*PlayerInfo)(unsafe.Pointer(&buf[3])) = *pl.Info()
}

func sub_422140(pl *Player) {
	if pl != nil {
		pl.field_3660 = 0xDEADFACE
		pl.field_3664 = 0xDEADFACE
	}
}

func sub_459D70() int {
	var v0 uint32
	if C.dword_5d4594_1046492 != 0 {
		v0 = math.MaxInt32
	} else {
		v0 = 0
	}
	v0 &= 0xFFFFFFFE
	return int(v0 + 2)
}

func (s *Server) nox_xxx_playerCheckName_4DDA00(pl *Player) {
	for i := 2; ; i++ {
		ok := true
		for _, pl2 := range s.getPlayers() {
			if pl.Index() == pl2.Index() {
				continue
			}
			if pl.Name() == pl2.Name() {
				ok = false
				break
			}
		}
		if ok {
			return
		}
		pl.Info().SetNameSuff(fmt.Sprintf(" %d", i))
		pl.SetName(fmt.Sprintf("%s %d", pl.OrigName(), i))
	}
}

func sub_4E4F30(a1 int) {
	*memmap.PtrUint16(0x5D4594, 1565524+2*uintptr(a1)) = 0
}

var (
	_ encoding.BinaryMarshaler   = &PlayerOpts{}
	_ encoding.BinaryUnmarshaler = &PlayerOpts{}
)

type PlayerOpts struct {
	Info      PlayerInfo
	Screen    image.Point
	Serial    string
	Field2096 string
	Field2068 int
	Field2072 string
	Byte152   byte
}

func (p *PlayerOpts) UnmarshalBinary(data []byte) error {
	*p = PlayerOpts{}
	if len(data) < 153 {
		return fmt.Errorf("cannot unmarshal player opts: message too short: %d < %d", len(data), 153)
	}
	p.Info = *(*PlayerInfo)(unsafe.Pointer(&data[0])) // TODO: set fields individually
	p.Screen = image.Point{
		X: int(binary.LittleEndian.Uint32(data[97:101])),
		Y: int(binary.LittleEndian.Uint32(data[101:105])),
	}
	p.Serial = GoStringS(data[105:128])
	p.Field2096 = GoStringS(data[128:138])
	p.Field2068 = int(binary.LittleEndian.Uint32(data[138:142]))
	p.Field2072 = GoWStringBytes(data[142:152])
	p.Byte152 = data[152]
	return nil
}

func (p *PlayerOpts) MarshalBinary() ([]byte, error) {
	data := make([]byte, 153)
	*(*PlayerInfo)(unsafe.Pointer(&data[0])) = p.Info // TODO: set fields individually
	binary.LittleEndian.PutUint32(data[97:101], uint32(p.Screen.X))
	binary.LittleEndian.PutUint32(data[101:105], uint32(p.Screen.Y))
	StrCopyBytes(data[105:128], p.Serial)
	StrCopyBytes(data[128:138], p.Field2096)
	binary.LittleEndian.PutUint32(data[138:142], uint32(p.Field2068))
	WStrCopyBytes(data[142:152], p.Field2072)
	data[152] = p.Byte152
	return data, nil
}

func (s *Server) newPlayerFromPacket(ind int, data []byte) int {
	var opts PlayerOpts
	if err := opts.UnmarshalBinary(data); err != nil {
		panic(err)
	}
	return s.newPlayer(ind, &opts)
}

func (s *Server) newPlayer(ind int, opts *PlayerOpts) int {
	v2 := opts.Byte152
	opts.Byte152 &= 0x7F
	v3 := v2 >> 7
	if ind != 31 {
		if !noxflags.HasGame(noxflags.GameModeQuest) && v3 == 1 {
			return 0
		}
		if noxflags.HasGame(noxflags.GameModeQuest) && v3 == 0 {
			return 0
		}
	}
	v5 := sub_416640()
	nox_netlist_resetByInd_40ED10(C.int(ind), 1)
	nox_xxx_playerResetImportantCtr_4E4F40(C.int(ind))
	sub_4E4F30(ind)

	var ptyp string
	if opts.Info.IsFemale() {
		ptyp = "PlayerFemale"
	} else if memmap.Uint32(0x5D4594, 1563280) != 0 {
		ptyp = "Player"
	} else {
		ptyp = "NewPlayer"
	}
	punit := s.newObjectByTypeID(ptyp).AsUnit()
	if punit == nil {
		return 0
	}
	if ind != 31 {
		if v5[100] != 0 {
			if (1<<opts.Info.PlayerClass())&v5[100] != 0 {
				return 0
			}
		}
	}
	pl := s.playerResetInd(ind)
	if int8(v5[102]) >= 0 {
		pl.field_10 = C.ushort(opts.Screen.X / 2)
		pl.field_12 = C.ushort(opts.Screen.Y / 2)
	} else {
		if nox_win_width >= opts.Screen.X {
			pl.field_10 = C.ushort(opts.Screen.X / 2)
			pl.field_12 = C.ushort(opts.Screen.Y / 2)
		} else {
			pl.field_10 = C.ushort(nox_win_width / 2)
			pl.field_12 = C.ushort(nox_win_height / 2)
		}
	}
	pl.SetSerial(opts.Serial)
	pl.field_2135 = C.uchar(opts.Byte152)
	WStrCopy(&pl.field_2072[0], 10, opts.Field2072)
	pl.SetField2096(opts.Field2096)
	pl.field_2068 = C.uint(opts.Field2068)
	if pl.field_2068 != 0 {
		v12 := unsafe.Pointer(sub_425A70(C.int(pl.field_2068)))
		if v12 == nil {
			v12 = unsafe.Pointer(sub_425AD0(C.int(pl.field_2068), &pl.field_2072[0]))
		}
		sub_425B30(v12, C.int(ind))
	}
	pl.frame_3596 = C.uint(gameFrame())
	pl.field_3676 = 2
	pl.field_3680 = 0
	info := pl.Info()
	*info = opts.Info
	info.SetNameSuff("")
	pl.SetName(pl.OrigName())
	s.nox_xxx_playerCheckName_4DDA00(pl)
	nox_xxx_playerInitColors_461460(pl.C())
	pl.playerUnit = punit.CObj()
	pl.field_2152 = 0
	pl.netCode = punit.field_9
	pl.field_2156 = C.uint(nox_xxx_scavengerTreasureMax_4D1600())
	udata := punit.updateDataPlayer()
	xxx := punit.ptrXxx()
	udata.player = pl.C()
	pl.prot_unit_hp_cur = C.uint(protectUint16(*(*uint16)(unsafe.Add(xxx, 0))))
	pl.prot_unit_hp_max = C.uint(protectUint16(*(*uint16)(unsafe.Add(xxx, 4))))
	pl.prot_unit_mana_cur = C.uint(protectUint16(uint16(udata.mana_cur)))
	pl.prot_unit_mana_max = C.uint(protectUint16(uint16(udata.mana_max)))
	pl.prot_unit_experience = C.uint(protectFloat32(float32(punit.experience)))
	pl.prot_unit_mass = C.uint(protectFloat32(punit.Mass()))
	pl.prot_unit_field_85 = C.uint(protectInt(int(*(*uint32)(&punit.field_85))))
	pl.prot_player_class = C.uint(protectInt(int(pl.PlayerClass())))
	pl.prot_player_field_2235 = C.uint(protectUint32(pl.Info().Field2235()))
	pl.prot_player_field_2239 = C.uint(protectUint32(pl.Info().Field2239()))
	pl.prot_player_orig_name = C.uint(protectUint32(protectWStr(pl.OrigName())))
	pl.prot_4632 = C.uint(protectInt(0))
	pl.prot_4636 = C.uint(protectInt(0))
	pl.prot_4640 = C.uint(protectInt(0))
	pl.prot_player_gold = C.uint(protectInt(int(pl.gold)))
	pl.prot_player_level = C.uint(protectInt(int(pl.field_3684))) // level
	pl.field_4648 = -1
	pl.field_4700 = 1
	if C.dword_5d4594_2650652 != 0 {
		sub_41D670(internCStr(pl.Field2096()))
	}
	nox_xxx_netNotifyRate_4D7F10(C.int(ind))
	if noxflags.HasGame(noxflags.GameModeQuest) {
		pl.GoObserver(false, true)
	} else if noxflags.HasGame(noxflags.GameModeSolo10) {
		nox_xxx_netReportPlayerStatus_417630(pl.C())
	} else if pl.Index() == 31 && noxflags.HasEngine(noxflags.EngineNoRendering) {
		pl.GoObserver(false, true)
	} else if noxflags.HasGame(noxflags.GameModeChat) {
		if sub_40A740() != 0 {
			if sub_40AA70(pl.C()) == 0 {
				pl.GoObserver(false, true)
			}
		} else if checkGameplayFlags(4) {
			sub_4DF3C0(pl.C())
		}
	} else if !noxflags.HasGame(noxflags.GameModeCoop) {
		pl.GoObserver(true, true)
	}
	nox_xxx_servSendSettings_4DDB40(punit.CObj())
	if pl.Index() == 31 {
		C.nox_xxx_host_player_unit_3843628 = punit.CObj()
	}
	var v30 [132]byte
	nox_xxx_netNewPlayerMakePacket_4DDA90(v30[:], pl)
	s.nox_xxx_netSendPacket_4E5030(ind|0x80, v30[:129], 0, 0, 0)
	pl.field_3676 = 2
	if nox_xxx_check_flag_aaa_43AF70() == 1 && !noxflags.HasGame(noxflags.GameModeChat) {
		sub_425F10(pl.C())
	}
	nox_xxx_createAt_4DAA50(punit, nil, types.Pointf{X: 2944.0, Y: 2944.0})
	nox_xxx_unitsNewAddToList_4DAC00()
	var p28 types.Pointf
	if noxflags.HasGame(noxflags.GameModeQuest) {
		if p, ok := s.sub_4E8210(punit); !ok {
			p28 = nox_xxx_mapFindPlayerStart_4F7AB0(punit)
		} else {
			p28 = p
		}
	} else {
		p28 = nox_xxx_mapFindPlayerStart_4F7AB0(punit)
	}
	punit.SetPos(p28)
	sub_422140(pl)
	if ind != 31 {
		if sub_459D70() == 2 {
			v24 := nox_xxx_cliGamedataGet_416590(1)
			nox_xxx_netGuiGameSettings_4DD9B0(1, unsafe.Pointer(&v24[0]), C.int(pl.Index()))
		} else {
			v29, v29free := alloc.Make([]byte{}, 60)
			defer v29free()
			sub_459AA0(unsafe.Pointer(&v29[0]))
			nox_xxx_netGuiGameSettings_4DD9B0(1, unsafe.Pointer(&v29[0]), C.int(pl.Index()))
		}
	}
	if noxflags.HasGame(noxflags.GameFlag15 | noxflags.GameFlag16) {
		if (pl.field_3680 & 1) == 0 {
			sub_509C30(pl.C())
		}
	}
	if !noxflags.HasGame(noxflags.GameModeCoop) {
		if noxflags.HasGame(noxflags.GameModeQuest) {
			nox_game_sendQuestStage_4D6960(C.int(ind))
			return int(punit.field_9)
		}
		var buf [3]byte
		buf[0] = byte(noxnet.MSG_FADE_BEGIN)
		buf[1] = 1
		buf[2] = 1
		s.nox_xxx_netSendPacket_4E5030(ind, buf[:], 0, 0, 0)
	}
	s.callOnPlayerJoin(pl)
	return int(punit.field_9)
}

func (s *Server) sub_4E8210(u *Unit) (types.Pointf, bool) {
	var (
		max uint32
		v2  unsafe.Pointer
	)
	for _, u2 := range s.getPlayerUnits() {
		ptr := u2.updateDataPlayer()
		ptr2 := ptr.field_77
		if ptr2 == nil {
			continue
		}
		if val := **(**uint32)(unsafe.Add(ptr2, 700)); val > max {
			max = val
			v2 = ptr2
		}
	}
	if v2 == nil {
		return types.Pointf{}, false
	}
	ud := u.updateDataPlayer()
	ud.field_77 = v2
	out := randomReachablePointAround(60.0, asPointf(unsafe.Add(v2, 7*8)))
	return out, true
}

//export nox_xxx_playerSpell_4FB2A0_magic_plyrspel
func nox_xxx_playerSpell_4FB2A0_magic_plyrspel(up *nox_object_t) {
	u := asUnitC(up)

	ok2 := true
	ud := u.updateDataPlayer()
	pl := ud.Player()
	var a1 int
	if u != nil {
		a1 = 1
	}
	if leaf := (*phonemeLeaf)(ud.spell_phoneme_leaf); leaf == getPhonemeTree() {
		ok2 = false
	} else if leaf != nil && leaf.Ind != 0 {
		spellInd := things.SpellID(leaf.Ind)
		if !noxflags.HasGame(noxflags.GameModeQuest) {
			targ := ud.CursorObj()
			if noxServer.spellHasFlags(spellInd, things.SpellOffensive) {
				if targ != nil && !u.isEnemyTo(targ) {
					return
				}
			}
		}
		if pl.spell_lvl[spellInd] != 0 || spellInd == 34 {
			ok2 = false
			a1 = int(sub_4FD0E0(u.CObj(), C.int(spellInd)))
			if a1 == 0 {
				a1 = int(nox_xxx_checkPlrCantCastSpell_4FD150(u.CObj(), C.int(spellInd), 0))
			}
			if a1 != 0 {
				nox_xxx_netInformTextMsg_4DA0F0(pl.Index(), 0, a1)
				nox_xxx_aud_501960(231, u, 0, 0)
			} else {
				mana := int(sub_4FCF90(u.CObj(), C.int(spellInd), 1))
				if mana < 0 {
					a1 = 11
					nox_xxx_netInformTextMsg_4DA0F0(pl.Index(), 0, a1)
					nox_xxx_aud_501960(232, u, 0, 0)
				} else {
					arg, v14free := alloc.New(spellAcceptArg{})
					defer v14free()
					arg.Obj = pl.obj_3640
					if noxflags.HasGame(noxflags.GameModeQuest) && noxServer.spellHasFlags(spellInd, 32) {
						if pl.obj_3640 != nil && !u.isEnemyTo(asObjectC(pl.obj_3640)) {
							arg.Obj = nil
						}
					}
					arg.Arg1 = float32(pl.field_2284)
					arg.Arg2 = float32(pl.field_2288)
					if nox_xxx_castSpellByUser4FDD20(spellInd, u, arg) {
						nox_xxx_netInformTextMsg_4DA0F0(pl.Index(), 1, int(spellInd))
					} else {
						sub_4FD030(u, mana)
						a1 = 8
					}
				}
			}
		}
	}
	if ud.field_22_0 == 2 {
		nox_xxx_playerSetState_4FA020(u.CObj(), 13)
	}
	if ok2 {
		v13 := noxServer.Strings().GetStringInFile("SpellUnknown", "C:\\NoxPost\\src\\Server\\Magic\\plyrspel.c")
		nox_xxx_netSendLineMessage_4D9EB0(u, v13)
	} else if a1 != 0 {
		v4 := (*phonemeLeaf)(ud.spell_phoneme_leaf)
		nox_xxx_netReportSpellStat_4D9630(pl.Index(), things.SpellID(v4.Ind), 0)
	} else {
		v4 := (*phonemeLeaf)(ud.spell_phoneme_leaf)
		if !noxServer.spellHasFlags(things.SpellID(v4.Ind), things.SpellFlagUnk21) {
			nox_xxx_netReportSpellStat_4D9630(pl.Index(), things.SpellID(v4.Ind), 15)
		}
	}
}

func sub_4FD030(a1 *Unit, a2 int) {
	if a1.Class().Has(object.ClassPlayer) {
		nox_xxx_playerManaAdd_4EEB80(a1.CObj(), C.short(a2))
	}
}
