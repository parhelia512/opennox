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
extern unsigned int dword_5d4594_1046492;
extern unsigned int dword_5d4594_2650652;
extern unsigned int nox_player_netCode_85319C;
extern nox_object_t* nox_xxx_host_player_unit_3843628;
void nox_xxx_WideScreenDo_515240(bool enable);
static void nox_xxx_printToAll_4D9FD0_go(int a1, wchar_t* str) {
	nox_xxx_printToAll_4D9FD0(a1, str);
}

extern float nox_xxx_warriorMaxHealth_587000_312784;
extern float nox_xxx_warriorMaxMana_587000_312788;

extern float nox_xxx_conjurerMaxHealth_587000_312800;
extern float nox_xxx_conjurerMaxMana_587000_312804;

extern float nox_xxx_wizardMaxHealth_587000_312816;
extern float nox_xxx_wizardMaximumMana_587000_312820;

*/
import "C"
import (
	"encoding"
	"encoding/binary"
	"fmt"
	"image"
	"math"
	"unsafe"

	"github.com/noxworld-dev/opennox-lib/common"
	"github.com/noxworld-dev/opennox-lib/noxnet"
	"github.com/noxworld-dev/opennox-lib/object"
	"github.com/noxworld-dev/opennox-lib/player"
	"github.com/noxworld-dev/opennox-lib/script"
	"github.com/noxworld-dev/opennox-lib/spell"
	"github.com/noxworld-dev/opennox-lib/things"
	"github.com/noxworld-dev/opennox-lib/types"

	"github.com/noxworld-dev/opennox/v1/common/alloc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"github.com/noxworld-dev/opennox/v1/common/sound"
	"github.com/noxworld-dev/opennox/v1/internal/netlist"
	"github.com/noxworld-dev/opennox/v1/server"
)

//export nox_xxx_playerSpell_4FB2A0_magic_plyrspel
func nox_xxx_playerSpell_4FB2A0_magic_plyrspel(up *nox_object_t) {
	noxServer.playerSpell(asUnitC(up))
}

//export nox_xxx_updateSpellRelated_424830
func nox_xxx_updateSpellRelated_424830(p unsafe.Pointer, ph C.int) unsafe.Pointer {
	return unsafe.Pointer((*phonemeLeaf)(p).Next(spell.Phoneme(ph)))
}

//export nox_common_playerInfoGetFirst_416EA0
func nox_common_playerInfoGetFirst_416EA0() *C.nox_playerInfo {
	return noxServer.PlayerFirst().C()
}

//export nox_common_playerInfoGetNext_416EE0
func nox_common_playerInfoGetNext_416EE0(it *C.nox_playerInfo) *C.nox_playerInfo {
	return noxServer.PlayerNext(asPlayer(it)).C()
}

//export nox_common_playerInfoCount_416F40
func nox_common_playerInfoCount_416F40() C.int {
	return C.int(noxServer.Players.Count())
}

//export nox_common_playerInfoGetByID_417040
func nox_common_playerInfoGetByID_417040(id C.int) *C.nox_playerInfo {
	return noxServer.GetPlayerByID(int(id)).C()
}

//export nox_common_playerInfoFromNum_417090
func nox_common_playerInfoFromNum_417090(ind C.int) *C.nox_playerInfo {
	return noxServer.GetPlayerByInd(int(ind)).C()
}

//export nox_common_playerInfoFromNumRaw
func nox_common_playerInfoFromNumRaw(ind C.int) *C.nox_playerInfo {
	return noxServer.GetPlayerByIndRaw(int(ind)).C()
}

//export nox_xxx_playerDisconnByPlrID_4DEB00
func nox_xxx_playerDisconnByPlrID_4DEB00(id C.int) {
	if p := noxServer.GetPlayerByInd(int(id)); p != nil {
		p.Disconnect(4)
	}
}

//export nox_xxx_playerCallDisconnect_4DEAB0
func nox_xxx_playerCallDisconnect_4DEAB0(ind C.int, v C.char) *C.char {
	noxServer.GetPlayerByInd(int(ind)).Disconnect(int(v))
	return nil
}

//export nox_xxx_playerCameraUnlock_4E6040
func nox_xxx_playerCameraUnlock_4E6040(cplayer *nox_object_t) {
	asUnitC(cplayer).ControllingPlayer().CameraUnlock()
}

//export nox_xxx_playerCameraFollow_4E6060
func nox_xxx_playerCameraFollow_4E6060(cplayer, cunit *nox_object_t) {
	asUnitC(cplayer).ControllingPlayer().CameraToggle(asObjectC(cunit))
}

//export nox_xxx_playerGetPossess_4DDF30
func nox_xxx_playerGetPossess_4DDF30(cplayer *nox_object_t) *nox_object_t {
	return asUnitC(cplayer).ControllingPlayer().ObserveTarget().CObj()
}

//export nox_xxx_playerGoObserver_4E6860
func nox_xxx_playerGoObserver_4E6860(pl *C.nox_playerInfo, a2 C.int, a3 C.int) C.int {
	return C.int(bool2int(asPlayer(pl).GoObserver(a2 != 0, a3 != 0)))
}

//export nox_xxx_playerObserveClear_4DDEF0
func nox_xxx_playerObserveClear_4DDEF0(cplayer *nox_object_t) {
	asUnitC(cplayer).observeClear()
}

func clientPlayer() *Player {
	return noxServer.GetPlayerByID(clientPlayerNetCode())
}

func clientPlayerNetCode() int {
	return int(C.nox_player_netCode_85319C)
}

func clientSetPlayerNetCode(id int) {
	C.nox_player_netCode_85319C = C.uint(id)
}

func (s *Server) PlayerFirst() *Player {
	return asPlayerS(s.Players.First())
}

func (s *Server) PlayerNext(it *Player) *Player {
	return asPlayerS(s.Players.Next(it.S()))
}

func (s *Server) PlayerResetInd(ind int) *Player {
	return asPlayerS(s.Players.ResetInd(ind))
}

func (s *Server) NewPlayerInfo(id int) *Player {
	return asPlayerS(s.Players.NewRaw(id))
}

func getPlayerClass() player.Class {
	return player.Class(memmap.Uint8(0x85B3FC, 12254))
}

func asPlayer(p *C.nox_playerInfo) *Player {
	return (*Player)(unsafe.Pointer(p))
}

func asPlayerS(p *server.Player) *Player {
	return (*Player)(p)
}

func BlindPlayers(blind bool) {
	noxServer.nox_xxx_netMsgFadeBegin_4D9800(!blind, false)
}

func PrintToPlayers(text string) {
	cstr, free := CWString(text)
	defer free()
	C.nox_xxx_printToAll_4D9FD0_go(0, cstr)
}

func HostPlayer() *Player {
	// TODO: better way
	for _, p := range noxServer.GetPlayers() {
		if p.IsHost() {
			return p
		}
	}
	return nil
}

var _ noxObject = (*Player)(nil) // proxies Unit

type Player server.Player

func (p *Player) getServer() *Server {
	return noxServer // TODO: attach to object
}

func (p *Player) field(off uintptr) unsafe.Pointer {
	return unsafe.Add(unsafe.Pointer(p), off)
}

func (p *Player) net16() *[255]server.PlayerNetData {
	return &p.NetData16
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

func (p *Player) CursorPos() types.Pointf {
	return p.S().CursorPos()
}

func (p *Player) SetCursorPos(pos image.Point) {
	p.S().SetCursorPos(pos)
}

func (p *Player) Pos3632() types.Pointf {
	return p.S().Pos3632()
}

func (p *Player) SetPos3632(pt types.Pointf) {
	p.S().SetPos3632(pt)
}

func (p *Player) OrigName() string {
	return p.S().OrigName()
}

func (p *Player) SetName(v string) {
	p.S().SetName(v)
}

func (p *Player) Name() string {
	return p.S().Name()
}

func (p *Player) SaveName() string {
	return p.S().SaveName()
}

func (p *Player) Serial() string {
	return p.S().Serial()
}

func (p *Player) SetSerial(v string) {
	p.S().SetSerial(v)
}

func (p *Player) Field2096() string {
	return p.S().Field2096()
}

func (p *Player) SetField2096(v string) {
	p.S().SetField2096(v)
}

func (p *Player) String() string {
	return p.S().String()
}

func (p *Player) Gold() int {
	return p.S().Gold()
}

func (p *Player) IsHost() bool {
	// TODO: better way
	return p.UnitC() == HostPlayerUnit()
}

func (p *Player) Print(text string) {
	nox_xxx_netSendLineMessage_4D9EB0(p.UnitC(), text)
}

func (p *Player) Blind(blind bool) {
	C.nox_xxx_netMsgFadeBeginPlayer(C.int(p.Index()), C.int(bool2int(!blind)), 0)
}

func (p *Player) Cinema(v bool) {
	// TODO: only works for everyone for now
	p.getServer().CinemaPlayers(v)
}

func (p *Player) CObj() *nox_object_t {
	u := p.UnitC()
	if u == nil {
		return nil
	}
	return u.CObj()
}

func (p *Player) SObj() *server.Object {
	return p.S().SObj()
}

func (p *Player) AsObject() *Object {
	u := p.UnitC()
	if u == nil {
		return nil
	}
	return u.AsObject()
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
	return (*C.nox_playerInfo)(unsafe.Pointer(p))
}

func (p *Player) S() *server.Player {
	return (*server.Player)(p)
}

func (p *Player) Index() int {
	return p.S().Index()
}

func (p *Player) NetCode() int {
	return p.S().NetCode()
}

func (p *Player) IsActive() bool {
	return p.S().IsActive()
}

func (p *Player) UnitC() *Unit {
	if p == nil {
		return nil
	}
	return asUnitS(p.PlayerUnit)
}

func (p *Player) Info() *server.PlayerInfo {
	return p.S().Info()
}

func (p *Player) PlayerClass() player.Class {
	return p.S().PlayerClass()
}

func (p *Player) Disconnect(v int) {
	if !p.IsActive() {
		return
	}
	C.nox_xxx_playerDisconnFinish_4DE530(C.int(p.Index()), C.char(v))
	C.nox_xxx_playerForceDisconnect_4DE7C0(C.int(p.Index()))
	p.getServer().nox_xxx_netStructReadPackets2_4DEC50(p.Index())
}

func (p *Player) CameraTarget() *Object {
	return asObjectS(p.S().CameraTarget())
}

func (p *Player) ObserveTarget() *Object { // nox_xxx_playerGetPossess_4DDF30
	return asObjectS(p.S().ObserveTarget())
}

func (p *Player) Sub422140() {
	p.S().Sub422140()
}

func (p *Player) CameraUnlock() { // nox_xxx_playerCameraUnlock_4E6040
	p.S().CameraUnlock()
}

func (p *Player) CameraFollow(obj server.Obj) {
	p.S().CameraFollow(obj)
}

func (p *Player) CameraToggle(obj server.Obj) { // nox_xxx_playerCameraFollow_4E6060
	p.S().CameraToggle(obj)
}

func (p *Player) GoObserver(notify, keepPlayer bool) bool { // nox_xxx_playerGoObserver_4E6860
	if p == nil {
		return true
	}
	u := p.UnitC()
	if u == nil {
		return true
	}
	s := u.getServer()
	if !keepPlayer && s.abilities.IsAnyActive(u) {
		return false
	}
	if u.Update == unsafe.Pointer(C.nox_xxx_updatePlayerMonsterBot_4FAB20) {
		return false
	}
	ud := u.UpdateDataPlayer()
	if noxflags.HasGame(noxflags.GameModeKOTR | noxflags.GameModeCTF | noxflags.GameModeFlagBall) {
		crown := s.ObjectTypeID("Crown")
		ball := s.ObjectTypeID("GameBall")
		for it := u.FirstOwned516(); it != nil; it = it.NextOwned512() {
			typ := int(it.TypeInd)
			if typ == crown {
				u.CallDrop(it, u.Pos())
			} else if typ == ball {
				it.Obj130 = nil
				it.ObjFlags &= 0xFFFFFFBF
				it.SetOwner(nil)
				sub_4E8290(1, 0)
			} else if it.Class().Has(object.ClassFlag) {
				u.forceDropAt(it, u.Pos())
			}
		}
	}
	if p.ObserveTarget() != nil {
		u.observeClear()
	}
	C.nox_xxx_netNeedTimestampStatus_4174F0(p.C(), 1)
	v10 := C.nox_xxx_gamePlayIsAnyPlayers_40A8A0() != 0
	if !v10 && !noxflags.HasGame(noxflags.GameModeQuest) {
		C.sub_40A1F0(0)
		C.nox_xxx_playerForceSendLessons_416E50(1)
		nox_server_teamsResetYyy_417D00()
		C.sub_40A970()
	}
	nox_xxx_netInformTextMsg_4DA0F0(p.Index(), 12, bool2int(notify))
	u.ApplyEnchant(server.ENCHANT_INVISIBLE, 0, 5)
	u.ObjFlags |= uint32(object.FlagNoCollide)
	p.SetPos3632(u.Pos())
	p.CameraUnlock()
	if noxflags.HasGame(noxflags.GameModeCoop) {
		p.Field3672 = 1
		p.CameraFollowObj = nil
	} else if noxflags.HasGame(noxflags.GameModeFlagBall) {
		if !keepPlayer {
			p.leaveMonsterObserver()
		}
	}
	C.nox_xxx_playerRemoveSpawnedStuff_4E5AD0(u.CObj())
	ud.Field61 = 0
	_ = nox_xxx_updatePlayerObserver_4E62F0
	u.Update = C.nox_xxx_updatePlayerObserver_4E62F0
	C.sub_4D7E50(u.CObj())
	return true
}

func (p *Player) leaveMonsterObserver() {
	u := p.UnitC()
	if u == nil {
		return
	}
	var targ *Object
	if p.ObserveTarget() != nil {
		targ = asObjectC(C.nox_xxx_playerObserverFindGoodSlave0_4E6280(p.C()))
		if targ == nil {
			u.observeClear()
			return
		}
	} else {
		targ = asObjectC(C.sub_4E6150(p.C()))
	}
	p.CameraFollow(targ)
}

func (u *Unit) observeClear() {
	pl := u.ControllingPlayer()
	if pl.Field3680&2 != 0 {
		C.nox_xxx_playerUnsetStatus_417530(pl.C(), 2)
		pl.CameraUnlock()
		_ = nox_xxx_updatePlayer_4F8100
		u.Update = C.nox_xxx_updatePlayer_4F8100
	}
}

func (s *Server) GetAllPlayerStructs() (out []*Player) {
	arr := s.Players.ListSlots()
	out = make([]*Player, 0, len(arr))
	for _, p := range arr {
		out = append(out, asPlayerS(p))
	}
	return out
}

func (s *Server) GetPlayers() (out []*Player) {
	arr := s.Players.List()
	out = make([]*Player, 0, len(arr))
	for _, p := range arr {
		out = append(out, asPlayerS(p))
	}
	return out
}

func (s *Server) getPlayerUnits() (out []*Unit) {
	for _, p := range s.GetPlayers() {
		if u := p.UnitC(); u != nil {
			out = append(out, u)
		}
	}
	return out
}

func (s *Server) GetPlayerByInd(i int) *Player {
	return asPlayerS(s.Players.ByInd(i))
}

func (s *Server) GetPlayerByIndRaw(i int) *Player {
	return asPlayerS(s.Players.ByIndRaw(i))
}

func (s *Server) GetPlayerByID(id int) *Player {
	return asPlayerS(s.Players.ByID(id))
}

func nox_xxx_netNewPlayerMakePacket_4DDA90(buf []byte, pl *Player) {
	buf[0] = byte(noxnet.MSG_NEW_PLAYER)
	binary.LittleEndian.PutUint16(buf[1:], uint16(pl.NetCode()))
	binary.LittleEndian.PutUint16(buf[100:], uint16(pl.Lessons))
	binary.LittleEndian.PutUint16(buf[102:], uint16(pl.Field2140))
	binary.LittleEndian.PutUint32(buf[104:], uint32(pl.Field0))
	binary.LittleEndian.PutUint32(buf[108:], uint32(pl.Field4))
	buf[116] = byte(pl.Field2152)
	buf[117] = byte(pl.Field2156)
	buf[118] = byte(bool2int(pl.Field3676 == 3))
	binary.LittleEndian.PutUint32(buf[112:], uint32(pl.Field3680)&0x423)
	StrCopyBytes(buf[119:], pl.Field2096())
	*(*server.PlayerInfo)(unsafe.Pointer(&buf[3])) = *pl.Info()
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

func sub_4E4F30(a1 int) {
	*memmap.PtrUint16(0x5D4594, 1565524+2*uintptr(a1)) = 0
}

var (
	_ encoding.BinaryMarshaler   = &PlayerOpts{}
	_ encoding.BinaryUnmarshaler = &PlayerOpts{}
)

type PlayerOpts struct {
	Info      server.PlayerInfo
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
	p.Info = *(*server.PlayerInfo)(unsafe.Pointer(&data[0])) // TODO: set fields individually
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
	*(*server.PlayerInfo)(unsafe.Pointer(&data[0])) = p.Info // TODO: set fields individually
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
	if ind != common.MaxPlayers-1 {
		if !noxflags.HasGame(noxflags.GameModeQuest) && v3 == 1 {
			return 0
		}
		if noxflags.HasGame(noxflags.GameModeQuest) && v3 == 0 {
			return 0
		}
	}
	v5 := sub_416640()
	netlist.ResetByInd(ind, netlist.Kind1)
	C.nox_xxx_playerResetImportantCtr_4E4F40(C.int(ind))
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
	if ind != common.MaxPlayers-1 {
		if v5[100] != 0 {
			if (1<<opts.Info.PlayerClass())&v5[100] != 0 {
				return 0
			}
		}
	}
	pl := s.PlayerResetInd(ind)
	if int8(v5[102]) >= 0 {
		pl.Field10 = uint16(opts.Screen.X / 2)
		pl.Field12 = uint16(opts.Screen.Y / 2)
	} else {
		if nox_win_width >= opts.Screen.X {
			pl.Field10 = uint16(opts.Screen.X / 2)
			pl.Field12 = uint16(opts.Screen.Y / 2)
		} else {
			pl.Field10 = uint16(nox_win_width / 2)
			pl.Field12 = uint16(nox_win_height / 2)
		}
	}
	pl.SetSerial(opts.Serial)
	pl.Field2135 = opts.Byte152
	alloc.StrCopy16(pl.Field2072[:], opts.Field2072)
	pl.SetField2096(opts.Field2096)
	pl.Field2068 = uint32(opts.Field2068)
	if pl.Field2068 != 0 {
		v12 := unsafe.Pointer(C.sub_425A70(C.int(pl.Field2068)))
		if v12 == nil {
			v12 = unsafe.Pointer(C.sub_425AD0(C.int(pl.Field2068), (*C.ushort)(unsafe.Pointer(&pl.Field2072[0]))))
		}
		C.sub_425B30(v12, C.int(ind))
	}
	pl.Frame3596 = s.Frame()
	pl.Field3676 = 2
	pl.Field3680 = 0
	info := pl.Info()
	*info = opts.Info
	info.SetNameSuff("")
	pl.SetName(pl.OrigName())
	s.Players.CheckName(pl.S())
	C.nox_xxx_playerInitColors_461460(pl.C())
	pl.PlayerUnit = punit.SObj()
	pl.Field2152 = 0
	pl.NetCodeVal = punit.NetCode
	pl.Field2156 = uint32(C.nox_xxx_scavengerTreasureMax_4D1600())
	udata := punit.UpdateDataPlayer()
	h := punit.HealthData
	udata.Player = pl.S()
	pl.ProtUnitHPCur = protectUint16(h.Cur)
	pl.ProtUnitHPMax = protectUint16(h.Max)
	pl.ProtUnitManaCur = protectUint16(udata.ManaCur)
	pl.ProtUnitManaMax = protectUint16(udata.ManaMax)
	pl.ProtUnitExperience = protectFloat32(punit.Experience)
	pl.ProtUnitMass = protectFloat32(punit.Mass)
	pl.ProtUnitBuffs = protectInt(int(punit.Buffs))
	pl.ProtPlayerClass = protectInt(int(pl.PlayerClass()))
	pl.ProtPlayerField2235 = protectUint32(pl.Info().Field2235())
	pl.ProtPlayerField2239 = protectUint32(pl.Info().Field2239())
	pl.ProtPlayerOrigName = protectUint32(protectWStr(pl.OrigName()))
	pl.Prot4632 = protectInt(0)
	pl.Prot4636 = protectInt(0)
	pl.Prot4640 = protectInt(0)
	pl.ProtPlayerGold = protectInt(int(pl.GoldVal))
	pl.ProtPlayerLevel = protectInt(int(pl.Field3684)) // level
	pl.Field4648 = -1
	pl.Field4700 = 1
	if C.dword_5d4594_2650652 != 0 {
		C.sub_41D670(internCStr(pl.Field2096()))
	}
	C.nox_xxx_netNotifyRate_4D7F10(C.int(ind))
	if noxflags.HasGame(noxflags.GameModeQuest) {
		pl.GoObserver(false, true)
	} else if noxflags.HasGame(noxflags.GameModeSolo10) {
		C.nox_xxx_netReportPlayerStatus_417630(pl.C())
	} else if pl.Index() == common.MaxPlayers-1 && noxflags.HasEngine(noxflags.EngineNoRendering) {
		pl.GoObserver(false, true)
	} else if noxflags.HasGame(noxflags.GameModeChat) {
		if C.sub_40A740() != 0 {
			if C.sub_40AA70(pl.C()) == 0 {
				pl.GoObserver(false, true)
			}
		} else if checkGameplayFlags(4) {
			C.sub_4DF3C0(pl.C())
		}
	} else if !noxflags.HasGame(noxflags.GameModeCoop) {
		pl.GoObserver(true, true)
	}
	s.sendSettings(punit)
	if pl.Index() == common.MaxPlayers-1 {
		C.nox_xxx_host_player_unit_3843628 = punit.CObj()
	}
	var v30 [132]byte
	nox_xxx_netNewPlayerMakePacket_4DDA90(v30[:], pl)
	s.nox_xxx_netSendPacket_4E5030(ind|0x80, v30[:129], 0, 0, 0)
	pl.Field3676 = 2
	if false && !noxflags.HasGame(noxflags.GameModeChat) {
		C.sub_425F10(pl.C())
	}
	s.createObjectAt(punit, nil, types.Pointf{X: 2944.0, Y: 2944.0})
	s.objectsNewAdd()
	var p28 types.Pointf
	if noxflags.HasGame(noxflags.GameModeQuest) {
		if p, ok := s.sub_4E8210(punit); !ok {
			p28 = s.nox_xxx_mapFindPlayerStart_4F7AB0(punit)
		} else {
			p28 = p
		}
	} else {
		p28 = s.nox_xxx_mapFindPlayerStart_4F7AB0(punit)
	}
	punit.SetPos(p28)
	pl.Sub422140()
	if ind != common.MaxPlayers-1 {
		if sub_459D70() == 2 {
			v24 := nox_xxx_cliGamedataGet_416590(1)
			C.nox_xxx_netGuiGameSettings_4DD9B0(1, unsafe.Pointer(&v24[0]), C.int(pl.Index()))
		} else {
			v29, v29free := alloc.Make([]byte{}, 60)
			defer v29free()
			C.sub_459AA0(unsafe.Pointer(&v29[0]))
			C.nox_xxx_netGuiGameSettings_4DD9B0(1, unsafe.Pointer(&v29[0]), C.int(pl.Index()))
		}
	}
	if noxflags.HasGame(noxflags.GameFlag15 | noxflags.GameFlag16) {
		if (pl.Field3680 & 1) == 0 {
			C.sub_509C30(pl.C())
		}
	}
	if !noxflags.HasGame(noxflags.GameModeCoop) {
		if noxflags.HasGame(noxflags.GameModeQuest) {
			C.nox_game_sendQuestStage_4D6960(C.int(ind))
			return int(punit.NetCode)
		}
		var buf [3]byte
		buf[0] = byte(noxnet.MSG_FADE_BEGIN)
		buf[1] = 1
		buf[2] = 1
		s.nox_xxx_netSendPacket_4E5030(ind, buf[:], 0, 0, 0)
	}
	s.callOnPlayerJoin(pl)
	return int(punit.NetCode)
}

func (s *Server) sub_4E8210(u *Unit) (types.Pointf, bool) {
	var (
		max uint32
		v2  unsafe.Pointer
	)
	for _, u2 := range s.getPlayerUnits() {
		ptr := u2.UpdateDataPlayer()
		ptr2 := ptr.Field77
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
	ud := u.UpdateDataPlayer()
	ud.Field77 = v2
	out := s.randomReachablePointAround(60.0, asPointf(unsafe.Add(v2, 7*8)))
	return out, true
}

func nox_xxx_plrSetSpellType_4F9B90(u *Unit) {
	ud := u.UpdateDataPlayer()
	ud.SpellPhonemeLeaf = unsafe.Pointer(getPhonemeTree())
	ud.SpellCastStart = noxServer.Frame()
}

func (s *Server) playerSpell(u *Unit) {
	ok2 := true
	ud := u.UpdateDataPlayer()
	pl := asPlayerS(ud.Player)
	var a1 int
	if u != nil {
		a1 = 1
	}
	if leaf := (*phonemeLeaf)(ud.SpellPhonemeLeaf); leaf == getPhonemeTree() {
		ok2 = false
	} else if leaf != nil && leaf.Ind != 0 {
		spellInd := spell.ID(leaf.Ind)
		if !noxflags.HasGame(noxflags.GameModeQuest) {
			targ := asObjectS(ud.CursorObj)
			if s.spellHasFlags(spellInd, things.SpellOffensive) {
				if targ != nil && !u.isEnemyTo(targ) {
					return
				}
			}
		}
		if pl.SpellLvl[spellInd] != 0 || spellInd == spell.SPELL_GLYPH {
			ok2 = false
			a1 = int(C.sub_4FD0E0(u.CObj(), C.int(spellInd)))
			if a1 == 0 {
				a1 = int(C.nox_xxx_checkPlrCantCastSpell_4FD150(u.CObj(), C.int(spellInd), 0))
			}
			if a1 != 0 {
				nox_xxx_netInformTextMsg_4DA0F0(pl.Index(), 0, a1)
				s.AudioEventObj(sound.SoundPermanentFizzle, u, 0, 0)
			} else {
				mana := int(C.sub_4FCF90(u.CObj(), C.int(spellInd), 1))
				if mana < 0 {
					a1 = 11
					nox_xxx_netInformTextMsg_4DA0F0(pl.Index(), 0, a1)
					s.AudioEventObj(sound.SoundManaEmpty, u, 0, 0)
				} else {
					arg, v14free := alloc.New(spellAcceptArg{})
					defer v14free()
					arg.Obj = toCObjS(pl.Obj3640)
					if noxflags.HasGame(noxflags.GameModeQuest) && s.spellHasFlags(spellInd, things.SpellOffensive) {
						if pl.Obj3640 != nil && !u.isEnemyTo(asObjectS(pl.Obj3640)) {
							arg.Obj = nil
						}
					}
					arg.Pos = pl.CursorPos()
					if s.nox_xxx_castSpellByUser4FDD20(spellInd, u, arg) {
						nox_xxx_netInformTextMsg_4DA0F0(pl.Index(), 1, int(spellInd))
					} else {
						sub_4FD030(u, mana)
						a1 = 8
					}
				}
			}
		}
	}
	if ud.Field22_0 == 2 {
		nox_xxx_playerSetState_4FA020(u, 13)
	}
	if ok2 {
		v13 := s.Strings().GetStringInFile("SpellUnknown", "plyrspel.c")
		nox_xxx_netSendLineMessage_4D9EB0(u, v13)
	} else if a1 != 0 {
		v4 := (*phonemeLeaf)(ud.SpellPhonemeLeaf)
		nox_xxx_netReportSpellStat_4D9630(pl.Index(), spell.ID(v4.Ind), 0)
	} else {
		v4 := (*phonemeLeaf)(ud.SpellPhonemeLeaf)
		if !s.spellHasFlags(spell.ID(v4.Ind), things.SpellFlagUnk21) {
			nox_xxx_netReportSpellStat_4D9630(pl.Index(), spell.ID(v4.Ind), 15)
		}
	}
}

func sub_4FD030(a1 *Unit, a2 int) {
	if a1.Class().Has(object.ClassPlayer) {
		C.nox_xxx_playerManaAdd_4EEB80(a1.CObj(), C.short(a2))
	}
}

func nox_xxx_playerSubLessons_4D8EC0(u *Unit, val int) {
	if !u.Class().Has(object.ClassPlayer) {
		return
	}
	pl := u.ControllingPlayer()
	pl.Lessons -= int32(val)
}

func nox_xxx_changeScore_4D8E90(u *Unit, val int) {
	if !u.Class().Has(object.ClassPlayer) {
		return
	}
	pl := u.ControllingPlayer()
	pl.Lessons += int32(val)
}

func nox_xxx_loadBaseValues_57B200() {
	*memmap.PtrFloat32(0x5D4594, 2523812) = float32(gamedataFloat("BaseHealth"))
	*memmap.PtrFloat32(0x5D4594, 2523816) = float32(gamedataFloat("BaseMana"))
	*memmap.PtrFloat32(0x5D4594, 2523824) = float32(gamedataFloat("BaseStrength"))
	*memmap.PtrFloat32(0x5D4594, 2523820) = float32(gamedataFloat("BaseSpeed"))
	*memmap.PtrFloat32(0x5D4594, 2523828) = float32(gamedataFloat("WarriorMaxHealth")) * float32(C.nox_xxx_warriorMaxHealth_587000_312784)
	*memmap.PtrFloat32(0x5D4594, 2523832) = float32(gamedataFloat("WarriorMaxMana")) * float32(C.nox_xxx_warriorMaxMana_587000_312788)
	*memmap.PtrFloat32(0x5D4594, 2523840) = float32(gamedataFloat("WarriorMaxStrength")) * noxServer.Players.Mult.Warrior.Strength
	*memmap.PtrFloat32(0x5D4594, 2523836) = float32(gamedataFloat("WarriorMaxSpeed")) * noxServer.Players.Mult.Warrior.Speed
	*memmap.PtrFloat32(0x5D4594, 2523860) = float32(gamedataFloat("ConjurerMaxHealth")) * float32(C.nox_xxx_conjurerMaxHealth_587000_312800)
	*memmap.PtrFloat32(0x5D4594, 2523864) = float32(gamedataFloat("ConjurerMaxMana")) * float32(C.nox_xxx_conjurerMaxMana_587000_312804)
	*memmap.PtrFloat32(0x5D4594, 2523872) = float32(gamedataFloat("ConjurerMaxStrength")) * noxServer.Players.Mult.Conjurer.Strength
	*memmap.PtrFloat32(0x5D4594, 2523868) = float32(gamedataFloat("ConjurerMaxSpeed")) * noxServer.Players.Mult.Conjurer.Speed
	*memmap.PtrFloat32(0x5D4594, 2523844) = float32(gamedataFloat("WizardMaxHealth")) * float32(C.nox_xxx_wizardMaxHealth_587000_312816)
	*memmap.PtrFloat32(0x5D4594, 2523848) = float32(gamedataFloat("WizardMaxMana")) * float32(C.nox_xxx_wizardMaximumMana_587000_312820)
	*memmap.PtrFloat32(0x5D4594, 2523856) = float32(gamedataFloat("WizardMaxStrength")) * noxServer.Players.Mult.Wizard.Strength
	*memmap.PtrFloat32(0x5D4594, 2523852) = float32(gamedataFloat("WizardMaxSpeed")) * noxServer.Players.Mult.Wizard.Speed
}

func sub_4D6B10(a1 bool) {
	C.nox_xxx_warriorMaxHealth_587000_312784 = C.float(memmap.Float32(0x5D4594, 1556076))
	C.nox_xxx_warriorMaxMana_587000_312788 = C.float(memmap.Float32(0x5D4594, 1556084))
	noxServer.Players.Mult.Warrior.Strength = memmap.Float32(0x5D4594, 1556064)
	noxServer.Players.Mult.Warrior.Speed = memmap.Float32(0x5D4594, 1556072)
	C.nox_xxx_conjurerMaxHealth_587000_312800 = C.float(memmap.Float32(0x5D4594, 1556060))
	C.nox_xxx_conjurerMaxMana_587000_312804 = C.float(memmap.Float32(0x5D4594, 1556096))
	noxServer.Players.Mult.Conjurer.Strength = memmap.Float32(0x5D4594, 1550932)
	noxServer.Players.Mult.Conjurer.Speed = memmap.Float32(0x5D4594, 1556080)
	C.nox_xxx_wizardMaxHealth_587000_312816 = C.float(memmap.Float32(0x5D4594, 1556088))
	C.nox_xxx_wizardMaximumMana_587000_312820 = C.float(memmap.Float32(0x5D4594, 1556068))
	noxServer.Players.Mult.Wizard.Strength = memmap.Float32(0x5D4594, 1556100)
	noxServer.Players.Mult.Wizard.Speed = memmap.Float32(0x5D4594, 1556092)
	nox_xxx_loadBaseValues_57B200()
	for _, it := range noxServer.getPlayerUnits() {
		C.nox_xxx_plrReadVals_4EEDC0(it.CObj(), 0)
		if a1 {
			nox_xxx_netStatsMultiplier_4D9C20(it.CObj())
		}
	}
}

func sub_4D6A60() {
	C.nox_xxx_warriorMaxHealth_587000_312784 = 3
	C.nox_xxx_warriorMaxMana_587000_312788 = 1
	noxServer.Players.Mult.Warrior.Strength = 1
	noxServer.Players.Mult.Warrior.Speed = 1
	C.nox_xxx_conjurerMaxHealth_587000_312800 = 3
	C.nox_xxx_conjurerMaxMana_587000_312804 = 3
	noxServer.Players.Mult.Conjurer.Strength = 1
	noxServer.Players.Mult.Conjurer.Speed = 1
	C.nox_xxx_wizardMaxHealth_587000_312816 = 3
	C.nox_xxx_wizardMaximumMana_587000_312820 = 3
	noxServer.Players.Mult.Wizard.Strength = 1
	noxServer.Players.Mult.Wizard.Speed = 1
	nox_xxx_loadBaseValues_57B200()
	for _, it := range noxServer.getPlayerUnits() {
		C.nox_xxx_plrReadVals_4EEDC0(it.CObj(), 0)
		nox_xxx_netStatsMultiplier_4D9C20(it.CObj())
	}
}

func sub_4D6BE0() {
	*memmap.PtrFloat32(0x5D4594, 1556076) = float32(C.nox_xxx_warriorMaxHealth_587000_312784)
	*memmap.PtrFloat32(0x5D4594, 1556084) = float32(C.nox_xxx_warriorMaxMana_587000_312788)
	*memmap.PtrFloat32(0x5D4594, 1556064) = noxServer.Players.Mult.Warrior.Strength
	*memmap.PtrFloat32(0x5D4594, 1556072) = noxServer.Players.Mult.Warrior.Speed
	*memmap.PtrFloat32(0x5D4594, 1556060) = float32(C.nox_xxx_conjurerMaxHealth_587000_312800)
	*memmap.PtrFloat32(0x5D4594, 1556096) = float32(C.nox_xxx_conjurerMaxMana_587000_312804)
	*memmap.PtrFloat32(0x5D4594, 1550932) = noxServer.Players.Mult.Conjurer.Strength
	*memmap.PtrFloat32(0x5D4594, 1556080) = noxServer.Players.Mult.Conjurer.Speed
	*memmap.PtrFloat32(0x5D4594, 1556088) = float32(C.nox_xxx_wizardMaxHealth_587000_312816)
	*memmap.PtrFloat32(0x5D4594, 1556068) = float32(C.nox_xxx_wizardMaximumMana_587000_312820)
	*memmap.PtrFloat32(0x5D4594, 1556100) = noxServer.Players.Mult.Wizard.Strength
	*memmap.PtrFloat32(0x5D4594, 1556092) = noxServer.Players.Mult.Wizard.Speed
}

//export nox_client_onClassStats
func nox_client_onClassStats(cbuf *C.uchar, sz C.int) {
	data := unsafe.Slice((*byte)(unsafe.Pointer(cbuf)), int(sz))
	switch player.Class(memmap.Uint8(0x85B3FC, 12254)) {
	case player.Warrior:
		C.nox_xxx_warriorMaxHealth_587000_312784 = C.float(math.Float32frombits(binary.LittleEndian.Uint32(data[1:])))
		C.nox_xxx_warriorMaxMana_587000_312788 = C.float(math.Float32frombits(binary.LittleEndian.Uint32(data[5:])))
		noxServer.Players.Mult.Warrior.Strength = math.Float32frombits(binary.LittleEndian.Uint32(data[9:]))
		noxServer.Players.Mult.Warrior.Speed = math.Float32frombits(binary.LittleEndian.Uint32(data[13:]))
	case player.Wizard:
		C.nox_xxx_wizardMaxHealth_587000_312816 = C.float(math.Float32frombits(binary.LittleEndian.Uint32(data[1:])))
		C.nox_xxx_wizardMaximumMana_587000_312820 = C.float(math.Float32frombits(binary.LittleEndian.Uint32(data[5:])))
		noxServer.Players.Mult.Wizard.Strength = math.Float32frombits(binary.LittleEndian.Uint32(data[9:]))
		noxServer.Players.Mult.Wizard.Speed = math.Float32frombits(binary.LittleEndian.Uint32(data[13:]))
	case player.Conjurer:
		C.nox_xxx_conjurerMaxHealth_587000_312800 = C.float(math.Float32frombits(binary.LittleEndian.Uint32(data[1:])))
		C.nox_xxx_conjurerMaxMana_587000_312804 = C.float(math.Float32frombits(binary.LittleEndian.Uint32(data[5:])))
		noxServer.Players.Mult.Conjurer.Strength = math.Float32frombits(binary.LittleEndian.Uint32(data[9:]))
		noxServer.Players.Mult.Conjurer.Speed = math.Float32frombits(binary.LittleEndian.Uint32(data[13:]))
	}
	nox_xxx_loadBaseValues_57B200()
}

//export nox_xxx_playerObserveMonster_4DDE80
func nox_xxx_playerObserveMonster_4DDE80(cplayer, cunit *nox_object_t) {
	pu := asUnitC(cplayer)
	ud := pu.UpdateDataPlayer()
	pl := asPlayerS(ud.Player)

	targ := asObjectC(cunit)

	if pl.Field3680&0x1 != 0 {
		C.nox_xxx_playerLeaveObserver_0_4E6AA0(pl.C())
	}
	if pl.Field3680&0x2 != 0 {
		pu.observeClear()
	}
	C.nox_xxx_netNeedTimestampStatus_4174F0(pl.C(), 2)
	pl.CameraFollow(targ)
	_ = nox_xxx_updatePlayerObserver_4E62F0
	pu.Update = C.nox_xxx_updatePlayerObserver_4E62F0
}

func (s *Server) nox_xxx_playerLeaveObsByObserved_4E60A0(obj noxObject) {
	cobj := toCObj(obj)
	for pl := s.PlayerFirst(); pl != nil; pl = s.PlayerNext(pl) {
		if pl.CameraTarget().CObj() == cobj {
			pl.leaveMonsterObserver()
		}
	}
}
