package opennox

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
	"github.com/noxworld-dev/opennox-lib/spell"
	"github.com/noxworld-dev/opennox-lib/things"
	"github.com/noxworld-dev/opennox-lib/types"

	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"github.com/noxworld-dev/opennox/v1/common/ntype"
	"github.com/noxworld-dev/opennox/v1/common/sound"
	"github.com/noxworld-dev/opennox/v1/internal/netlist"
	"github.com/noxworld-dev/opennox/v1/internal/netstr"
	"github.com/noxworld-dev/opennox/v1/legacy"
	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc"
	"github.com/noxworld-dev/opennox/v1/server"
)

func nox_xxx_playerCameraUnlock_4E6040(cplayer *server.Object) {
	cplayer.ControllingPlayer().CameraUnlock()
}

func nox_xxx_playerCameraFollow_4E6060(cplayer, cunit *server.Object) {
	cplayer.ControllingPlayer().CameraToggle(cunit)
}

func nox_xxx_playerGetPossess_4DDF30(cplayer *server.Object) *server.Object {
	return cplayer.ControllingPlayer().ObserveTarget()
}

func nox_xxx_playerGoObserver_4E6860(pl *server.Player, a2 int, a3 int) int {
	return bool2int(noxServer.PlayerGoObserver(pl, a2 != 0, a3 != 0))
}

func nox_xxx_playerObserveClear_4DDEF0(cplayer *server.Object) {
	asObjectS(cplayer).observeClear()
}

func clientPlayer() *server.Player {
	return noxServer.Players.ByID(legacy.ClientPlayerNetCode())
}

func getPlayerClass() player.Class {
	return player.Class(memmap.Uint8(0x85B3FC, 12254))
}

func BlindPlayers(blind bool) {
	noxServer.Nox_xxx_netMsgFadeBegin_4D9800(!blind, false)
}

type Player = server.Player

func (s *Server) PlayerSetPos(p *Player, pos types.Pointf) {
	if p == nil {
		return
	}
	u := p.PlayerUnit
	if u == nil {
		return
	}
	asObjectS(u).SetPos(pos)
}

func (s *Server) PlayerPrint(p *Player, text string) {
	legacy.Nox_xxx_netSendLineMessage_4D9EB0(p.PlayerUnit, text)
}

func (s *Server) PlayerBlind(p *Player, blind bool) {
	legacy.Nox_xxx_netMsgFadeBeginPlayer(p.Index(), bool2int(!blind), 0)
}

func (s *Server) PlayerCinema(p *Player, v bool) {
	// TODO: only works for everyone for now
	s.CinemaPlayers(v)
}

func (s *Server) PlayerDisconnect(p *server.Player, v int) {
	if !p.IsActive() {
		return
	}
	nox_xxx_playerDisconnFinish_4DE530(p.PlayerIndex(), int8(v))
	legacy.Nox_xxx_playerForceDisconnect_4DE7C0(p.PlayerIndex())
	s.nox_xxx_netStructReadPackets2_4DEC50(p.PlayerIndex())
}

func (s *Server) PlayerDisconnectByIndCode4(ind ntype.PlayerInd) { // nox_xxx_playerDisconnByPlrID_4DEB00
	if pl := s.Players.ByInd(ind); pl != nil {
		s.PlayerDisconnect(pl, 4)
	}
}

func (s *Server) PlayerDisconnectByInd(ind ntype.PlayerInd, v int8) { // nox_xxx_playerCallDisconnect_4DEAB0
	if pl := s.Players.ByInd(ind); pl != nil {
		s.PlayerDisconnect(pl, int(v))
	}
}

func (s *Server) PlayerGoObserver(p *server.Player, notify, keepPlayer bool) bool { // nox_xxx_playerGoObserver_4E6860
	if p == nil {
		return true
	}
	u := p.PlayerUnit
	if u == nil {
		return true
	}
	if !keepPlayer && s.Abils.IsAnyActive(u) {
		return false
	}
	if u.Update == legacy.Get_nox_xxx_updatePlayerMonsterBot_4FAB20() {
		return false
	}
	ud := u.UpdateDataPlayer()
	if noxflags.HasGame(noxflags.GameModeKOTR | noxflags.GameModeCTF | noxflags.GameModeFlagBall) {
		crown := s.Types.CrownID()
		ball := s.Types.GameBallID()
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
				asObjectS(u).forceDropAt(it, u.Pos())
			}
		}
	}
	if p.ObserveTarget() != nil {
		asObjectS(u).observeClear()
	}
	legacy.Nox_xxx_netNeedTimestampStatus_4174F0(p, 1)
	v10 := legacy.Nox_xxx_gamePlayIsAnyPlayers_40A8A0() != 0
	if !v10 && !noxflags.HasGame(noxflags.GameModeQuest) {
		legacy.Sub_40A1F0(0)
		legacy.Nox_xxx_playerForceSendLessons_416E50(1)
		s.TeamsResetYyy()
		legacy.Sub_40A970()
	}
	s.NetInformTextMsg(p.PlayerIndex(), 12, bool2int(notify))
	asObjectS(u).ApplyEnchant(server.ENCHANT_INVISIBLE, 0, 5)
	u.ObjFlags |= uint32(object.FlagNoCollide)
	p.SetPos3632(u.Pos())
	p.CameraUnlock()
	if noxflags.HasGame(noxflags.GameModeCoop) {
		p.Field3672 = 1
		p.CameraFollowObj = nil
	} else if noxflags.HasGame(noxflags.GameModeFlagBall) {
		if !keepPlayer {
			s.PlayerLeaveMonsterObserver(p)
		}
	}
	legacy.Nox_xxx_playerRemoveSpawnedStuff_4E5AD0(u)
	ud.CurTraps = 0
	_ = nox_xxx_updatePlayerObserver_4E62F0
	u.Update = legacy.Get_nox_xxx_updatePlayerObserver_4E62F0()
	legacy.Sub_4D7E50(u)
	return true
}

func (s *Server) PlayerLeaveMonsterObserver(p *Player) {
	u := p.PlayerUnit
	if u == nil {
		return
	}
	var targ *server.Object
	if p.ObserveTarget() != nil {
		targ = legacy.Nox_xxx_playerObserverFindGoodSlave0_4E6280(p)
		if targ == nil {
			asObjectS(u).observeClear()
			return
		}
	} else {
		targ = legacy.Sub_4E6150(p)
	}
	p.CameraFollow(targ)
}

func (obj *Object) observeClear() {
	pl := obj.ControllingPlayer()
	if pl.Field3680&2 != 0 {
		legacy.Nox_xxx_playerUnsetStatus_417530(pl, 2)
		pl.CameraUnlock()
		_ = nox_xxx_updatePlayer_4F8100
		obj.Update = legacy.Get_nox_xxx_updatePlayer_4F8100()
	}
}

func nox_xxx_netNewPlayerMakePacket_4DDA90(buf []byte, pl *server.Player) {
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
	alloc.StrCopy(buf[119:], pl.Field2096())
	*(*server.PlayerInfo)(unsafe.Pointer(&buf[3])) = *pl.Info()
}

func sub_459D70() int {
	var v0 uint32
	if legacy.Get_dword_5d4594_1046492() != 0 {
		v0 = math.MaxInt32
	} else {
		v0 = 0
	}
	v0 &= 0xFFFFFFFE
	return int(v0 + 2)
}

func sub_4E4F30(a1 ntype.PlayerInd) {
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
	p.Serial = alloc.GoStringS(data[105:128])
	p.Field2096 = alloc.GoStringS(data[128:138])
	p.Field2068 = int(binary.LittleEndian.Uint32(data[138:142]))
	p.Field2072 = alloc.GoString16B(data[142:152])
	p.Byte152 = data[152]
	return nil
}

func (p *PlayerOpts) MarshalBinary() ([]byte, error) {
	data := make([]byte, 153)
	*(*server.PlayerInfo)(unsafe.Pointer(&data[0])) = p.Info // TODO: set fields individually
	binary.LittleEndian.PutUint32(data[97:101], uint32(p.Screen.X))
	binary.LittleEndian.PutUint32(data[101:105], uint32(p.Screen.Y))
	alloc.StrCopy(data[105:128], p.Serial)
	alloc.StrCopy(data[128:138], p.Field2096)
	binary.LittleEndian.PutUint32(data[138:142], uint32(p.Field2068))
	alloc.StrCopy16B(data[142:152], p.Field2072)
	data[152] = p.Byte152
	return data, nil
}

func (s *Server) newPlayerFromPacket(ind ntype.PlayerInd, data []byte) int {
	var opts PlayerOpts
	if err := opts.UnmarshalBinary(data); err != nil {
		panic(err)
	}
	return s.newPlayer(ind, &opts)
}

func (s *Server) newPlayer(ind ntype.PlayerInd, opts *PlayerOpts) int {
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
	s.NetList.ResetByInd(ind, netlist.Kind1)
	legacy.Nox_xxx_playerResetImportantCtr_4E4F40(ind)
	sub_4E4F30(ind)

	var ptyp string
	if opts.Info.IsFemale() {
		ptyp = "PlayerFemale"
	} else if memmap.Uint32(0x5D4594, 1563280) != 0 {
		ptyp = "Player"
	} else {
		ptyp = "NewPlayer"
	}
	punit := s.NewObjectByTypeID(ptyp)
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
	pl := s.Players.ResetInd(ind)
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
		v12 := legacy.Sub_425A70(int(pl.Field2068))
		if v12 == nil {
			v12 = legacy.Sub_425AD0(int(pl.Field2068), &pl.Field2072[0])
		}
		legacy.Sub_425B30(v12, ind)
	}
	pl.Frame3596 = s.Frame()
	pl.Field3676 = 2
	pl.Field3680 = 0
	info := pl.Info()
	*info = opts.Info
	info.SetNameSuff("")
	pl.SetName(pl.OrigName())
	s.Players.CheckName(pl)
	legacy.Nox_xxx_playerInitColors_461460(pl)
	pl.PlayerUnit = punit
	pl.Field2152 = 0
	pl.NetCodeVal = punit.NetCode
	pl.Field2156 = legacy.Nox_xxx_scavengerTreasureMax_4D1600()
	udata := punit.UpdateDataPlayer()
	h := punit.HealthData
	udata.Player = pl
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
	pl.ProtPlayerLevel = protectInt(int(pl.Level)) // level
	pl.Field4648 = -1
	pl.Field4700 = 1
	if legacy.Get_dword_5d4594_2650652() != 0 {
		legacy.Sub_41D670(pl.Field2096())
	}
	legacy.Nox_xxx_netNotifyRate_4D7F10(ind)
	if noxflags.HasGame(noxflags.GameModeQuest) {
		s.PlayerGoObserver(pl, false, true)
	} else if noxflags.HasGame(noxflags.GameModeCoopTeam) {
		legacy.Nox_xxx_netReportPlayerStatus_417630(pl)
	} else if pl.Index() == common.MaxPlayers-1 && noxflags.HasEngine(noxflags.EngineNoRendering) {
		s.PlayerGoObserver(pl, false, true)
	} else if noxflags.HasGame(noxflags.GameModeChat) {
		if legacy.Sub_40A740() != 0 {
			if legacy.Sub_40AA70(pl) == 0 {
				s.PlayerGoObserver(pl, false, true)
			}
		} else if noxflags.HasGamePlay(4) {
			legacy.Sub_4DF3C0(pl)
		}
	} else if !noxflags.HasGame(noxflags.GameModeCoop) {
		s.PlayerGoObserver(pl, true, true)
	}
	s.sendSettings(punit)
	if pl.Index() == common.MaxPlayers-1 {
		s.Players.HostUnit = punit
	}
	var v30 [132]byte
	nox_xxx_netNewPlayerMakePacket_4DDA90(v30[:], pl)
	s.NetSendPacketXxx(int(ind)|0x80, v30[:129], 0, 0, 0)
	pl.Field3676 = 2
	if false && !noxflags.HasGame(noxflags.GameModeChat) {
		legacy.Sub_425F10(pl)
	}
	s.CreateObjectAt(punit, nil, types.Pointf{X: 2944.0, Y: 2944.0})
	s.ObjectsAddPending()
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
	asObjectS(punit).SetPos(p28)
	pl.Sub422140()
	if ind != common.MaxPlayers-1 {
		if sub_459D70() == 2 {
			v24 := nox_xxx_cliGamedataGet_416590(1)
			legacy.Nox_xxx_netGuiGameSettings_4DD9B0(1, unsafe.Pointer(&v24[0]), pl.Index())
		} else {
			v29, v29free := alloc.Make([]byte{}, 60)
			defer v29free()
			legacy.Sub_459AA0(unsafe.Pointer(&v29[0]))
			legacy.Nox_xxx_netGuiGameSettings_4DD9B0(1, unsafe.Pointer(&v29[0]), pl.Index())
		}
	}
	if noxflags.HasGame(noxflags.GameFlag15 | noxflags.GameFlag16) {
		if (pl.Field3680 & 1) == 0 {
			legacy.Sub_509C30(pl)
		}
	}
	if !noxflags.HasGame(noxflags.GameModeCoop) {
		if noxflags.HasGame(noxflags.GameModeQuest) {
			legacy.Nox_game_sendQuestStage_4D6960(ind)
			return int(punit.NetCode)
		}
		var buf [3]byte
		buf[0] = byte(noxnet.MSG_FADE_BEGIN)
		buf[1] = 1
		buf[2] = 1
		s.NetSendPacketXxx(int(ind), buf[:], 0, 0, 0)
	}
	s.CallOnPlayerJoin(scrPlayer{pl})
	return int(punit.NetCode)
}

func (s *Server) sub_4E8210(u *server.Object) (types.Pointf, bool) {
	var (
		max uint32
		v2  unsafe.Pointer
	)
	for _, u2 := range s.Players.ListUnits() {
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
	out := s.RandomReachablePointAround(60.0, *(*types.Pointf)(unsafe.Add(v2, 56)))
	return out, true
}

func nox_xxx_plrSetSpellType_4F9B90(u *server.Object) {
	s := u.Server()
	ud := u.UpdateDataPlayer()
	ud.SpellPhonemeLeaf = s.Spells.PhonemeTree()
	ud.SpellCastStart = s.Frame()
}

func (s *Server) PlayerSpell(u *server.Object) {
	ok2 := true
	ud := u.UpdateDataPlayer()
	pl := ud.Player
	var a1 int
	if u != nil {
		a1 = 1
	}
	if leaf := ud.SpellPhonemeLeaf; leaf == s.Spells.PhonemeTree() {
		ok2 = false
	} else if leaf != nil && leaf.Ind != 0 {
		spellInd := spell.ID(leaf.Ind)
		if !noxflags.HasGame(noxflags.GameModeQuest) {
			targ := ud.CursorObj
			if s.Spells.HasFlags(spellInd, things.SpellOffensive) {
				if targ != nil && !s.IsEnemyTo(u, targ) {
					return
				}
			}
		}
		if pl.SpellLvl[spellInd] != 0 || spellInd == spell.SPELL_GLYPH {
			ok2 = false
			a1 = legacy.Sub_4FD0E0(u, spellInd)
			if a1 == 0 {
				a1 = legacy.Nox_xxx_checkPlrCantCastSpell_4FD150(u, spellInd, 0)
			}
			if a1 != 0 {
				s.NetInformTextMsg(pl.PlayerIndex(), 0, a1)
				s.Audio.EventObj(sound.SoundPermanentFizzle, u, 0, 0)
			} else {
				mana := legacy.Sub_4FCF90(u, spellInd, 1)
				if mana < 0 {
					a1 = 11
					s.NetInformTextMsg(pl.PlayerIndex(), 0, a1)
					s.Audio.EventObj(sound.SoundManaEmpty, u, 0, 0)
				} else {
					arg, v14free := alloc.New(server.SpellAcceptArg{})
					defer v14free()
					arg.Obj = pl.Obj3640
					if noxflags.HasGame(noxflags.GameModeQuest) && s.Spells.HasFlags(spellInd, things.SpellOffensive) {
						if pl.Obj3640 != nil && !s.IsEnemyTo(u, pl.Obj3640) {
							arg.Obj = nil
						}
					}
					arg.Pos = pl.CursorPos()
					if s.nox_xxx_castSpellByUser4FDD20(spellInd, -1, u, arg) {
						s.NetInformTextMsg(pl.PlayerIndex(), 1, int(spellInd))
					} else {
						sub_4FD030(u, mana)
						a1 = 8
					}
				}
			}
		}
	}
	if ud.State == server.PlayerState2 {
		nox_xxx_playerSetState_4FA020(u, server.PlayerState13)
	}
	if ok2 {
		v13 := s.Strings().GetStringInFile("SpellUnknown", "plyrspel.c")
		legacy.Nox_xxx_netSendLineMessage_4D9EB0(u, v13)
	} else if a1 != 0 {
		v4 := ud.SpellPhonemeLeaf
		s.NetReportSpellStat(pl.Index(), spell.ID(v4.Ind), 0)
	} else {
		v4 := ud.SpellPhonemeLeaf
		if !s.Spells.HasFlags(spell.ID(v4.Ind), things.SpellFlagUnk21) {
			s.NetReportSpellStat(pl.Index(), spell.ID(v4.Ind), 15)
		}
	}
}

func sub_4FD030(u *server.Object, v int) {
	if u.Class().Has(object.ClassPlayer) {
		legacy.Nox_xxx_playerManaAdd_4EEB80(u, v)
	}
}

func nox_xxx_playerSubLessons_4D8EC0(u *server.Object, val int) {
	if !u.Class().Has(object.ClassPlayer) {
		return
	}
	pl := u.ControllingPlayer()
	pl.Lessons -= int32(val)
}

func nox_xxx_changeScore_4D8E90(u *server.Object, val int) {
	if !u.Class().Has(object.ClassPlayer) {
		return
	}
	pl := u.ControllingPlayer()
	pl.Lessons += int32(val)
}

func (s *Server) LoadClassStats() {
	s.Players.Stats.Base = server.ClassStats{
		Health:   float32(s.Balance.Float("BaseHealth")),
		Mana:     float32(s.Balance.Float("BaseMana")),
		Speed:    float32(s.Balance.Float("BaseSpeed")),
		Strength: float32(s.Balance.Float("BaseStrength")),
	}

	_ = player.Warrior
	s.Players.Stats.Warrior = server.ClassStats{
		Health:   float32(s.Balance.Float("WarriorMaxHealth")) * legacy.Get_nox_xxx_warriorMaxHealth_587000_312784(),
		Mana:     float32(s.Balance.Float("WarriorMaxMana")) * legacy.Get_nox_xxx_warriorMaxMana_587000_312788(),
		Speed:    float32(s.Balance.Float("WarriorMaxSpeed")) * s.Players.Mult.Warrior.Speed,
		Strength: float32(s.Balance.Float("WarriorMaxStrength")) * s.Players.Mult.Warrior.Strength,
	}

	_ = player.Wizard
	s.Players.Stats.Wizard = server.ClassStats{
		Health:   float32(s.Balance.Float("WizardMaxHealth")) * legacy.Get_nox_xxx_wizardMaxHealth_587000_312816(),
		Mana:     float32(s.Balance.Float("WizardMaxMana")) * legacy.Get_nox_xxx_wizardMaximumMana_587000_312820(),
		Speed:    float32(s.Balance.Float("WizardMaxSpeed")) * s.Players.Mult.Wizard.Speed,
		Strength: float32(s.Balance.Float("WizardMaxStrength")) * s.Players.Mult.Wizard.Strength,
	}

	_ = player.Conjurer
	s.Players.Stats.Conjurer = server.ClassStats{
		Health:   float32(s.Balance.Float("ConjurerMaxHealth")) * legacy.Get_nox_xxx_conjurerMaxHealth_587000_312800(),
		Mana:     float32(s.Balance.Float("ConjurerMaxMana")) * legacy.Get_nox_xxx_conjurerMaxMana_587000_312804(),
		Speed:    float32(s.Balance.Float("ConjurerMaxSpeed")) * s.Players.Mult.Conjurer.Speed,
		Strength: float32(s.Balance.Float("ConjurerMaxStrength")) * s.Players.Mult.Conjurer.Strength,
	}
}

func (s *Server) sub_4D6B10(a1 bool) {
	legacy.Set_nox_xxx_warriorMaxHealth_587000_312784(memmap.Float32(0x5D4594, 1556076))
	legacy.Set_nox_xxx_warriorMaxMana_587000_312788(memmap.Float32(0x5D4594, 1556084))
	s.Players.Mult.Warrior.Strength = memmap.Float32(0x5D4594, 1556064)
	s.Players.Mult.Warrior.Speed = memmap.Float32(0x5D4594, 1556072)
	legacy.Set_nox_xxx_conjurerMaxHealth_587000_312800(memmap.Float32(0x5D4594, 1556060))
	legacy.Set_nox_xxx_conjurerMaxMana_587000_312804(memmap.Float32(0x5D4594, 1556096))
	s.Players.Mult.Conjurer.Strength = memmap.Float32(0x5D4594, 1550932)
	s.Players.Mult.Conjurer.Speed = memmap.Float32(0x5D4594, 1556080)
	legacy.Set_nox_xxx_wizardMaxHealth_587000_312816(memmap.Float32(0x5D4594, 1556088))
	legacy.Set_nox_xxx_wizardMaximumMana_587000_312820(memmap.Float32(0x5D4594, 1556068))
	s.Players.Mult.Wizard.Strength = memmap.Float32(0x5D4594, 1556100)
	s.Players.Mult.Wizard.Speed = memmap.Float32(0x5D4594, 1556092)
	s.LoadClassStats()
	for _, it := range s.Players.ListUnits() {
		legacy.Nox_xxx_plrReadVals_4EEDC0(it, 0)
		if a1 {
			s.nox_xxx_netStatsMultiplier_4D9C20(it)
		}
	}
}

func sub_4D6A60() {
	s := noxServer
	legacy.Set_nox_xxx_warriorMaxHealth_587000_312784(3)
	legacy.Set_nox_xxx_warriorMaxMana_587000_312788(1)
	s.Players.Mult.Warrior.Strength = 1
	s.Players.Mult.Warrior.Speed = 1
	legacy.Set_nox_xxx_conjurerMaxHealth_587000_312800(3)
	legacy.Set_nox_xxx_conjurerMaxMana_587000_312804(3)
	s.Players.Mult.Conjurer.Strength = 1
	s.Players.Mult.Conjurer.Speed = 1
	legacy.Set_nox_xxx_wizardMaxHealth_587000_312816(3)
	legacy.Set_nox_xxx_wizardMaximumMana_587000_312820(3)
	s.Players.Mult.Wizard.Strength = 1
	s.Players.Mult.Wizard.Speed = 1
	s.LoadClassStats()
	for _, it := range s.Players.ListUnits() {
		legacy.Nox_xxx_plrReadVals_4EEDC0(it, 0)
		s.nox_xxx_netStatsMultiplier_4D9C20(it)
	}
}

func sub_4D6BE0() {
	s := noxServer
	*memmap.PtrFloat32(0x5D4594, 1556076) = legacy.Get_nox_xxx_warriorMaxHealth_587000_312784()
	*memmap.PtrFloat32(0x5D4594, 1556084) = legacy.Get_nox_xxx_warriorMaxMana_587000_312788()
	*memmap.PtrFloat32(0x5D4594, 1556064) = s.Players.Mult.Warrior.Strength
	*memmap.PtrFloat32(0x5D4594, 1556072) = s.Players.Mult.Warrior.Speed
	*memmap.PtrFloat32(0x5D4594, 1556060) = legacy.Get_nox_xxx_conjurerMaxHealth_587000_312800()
	*memmap.PtrFloat32(0x5D4594, 1556096) = legacy.Get_nox_xxx_conjurerMaxMana_587000_312804()
	*memmap.PtrFloat32(0x5D4594, 1550932) = s.Players.Mult.Conjurer.Strength
	*memmap.PtrFloat32(0x5D4594, 1556080) = s.Players.Mult.Conjurer.Speed
	*memmap.PtrFloat32(0x5D4594, 1556088) = legacy.Get_nox_xxx_wizardMaxHealth_587000_312816()
	*memmap.PtrFloat32(0x5D4594, 1556068) = legacy.Get_nox_xxx_wizardMaximumMana_587000_312820()
	*memmap.PtrFloat32(0x5D4594, 1556100) = s.Players.Mult.Wizard.Strength
	*memmap.PtrFloat32(0x5D4594, 1556092) = s.Players.Mult.Wizard.Speed
}

func nox_client_onClassStats(data []byte) {
	s := noxServer
	switch getPlayerClass() {
	case player.Warrior:
		legacy.Set_nox_xxx_warriorMaxHealth_587000_312784(math.Float32frombits(binary.LittleEndian.Uint32(data[1:])))
		legacy.Set_nox_xxx_warriorMaxMana_587000_312788(math.Float32frombits(binary.LittleEndian.Uint32(data[5:])))
		s.Players.Mult.Warrior.Strength = math.Float32frombits(binary.LittleEndian.Uint32(data[9:]))
		s.Players.Mult.Warrior.Speed = math.Float32frombits(binary.LittleEndian.Uint32(data[13:]))
	case player.Wizard:
		legacy.Set_nox_xxx_wizardMaxHealth_587000_312816(math.Float32frombits(binary.LittleEndian.Uint32(data[1:])))
		legacy.Set_nox_xxx_wizardMaximumMana_587000_312820(math.Float32frombits(binary.LittleEndian.Uint32(data[5:])))
		s.Players.Mult.Wizard.Strength = math.Float32frombits(binary.LittleEndian.Uint32(data[9:]))
		s.Players.Mult.Wizard.Speed = math.Float32frombits(binary.LittleEndian.Uint32(data[13:]))
	case player.Conjurer:
		legacy.Set_nox_xxx_conjurerMaxHealth_587000_312800(math.Float32frombits(binary.LittleEndian.Uint32(data[1:])))
		legacy.Set_nox_xxx_conjurerMaxMana_587000_312804(math.Float32frombits(binary.LittleEndian.Uint32(data[5:])))
		s.Players.Mult.Conjurer.Strength = math.Float32frombits(binary.LittleEndian.Uint32(data[9:]))
		s.Players.Mult.Conjurer.Speed = math.Float32frombits(binary.LittleEndian.Uint32(data[13:]))
	}
	s.LoadClassStats()
}

func nox_xxx_playerObserveMonster_4DDE80(cplayer, cunit *server.Object) {
	pu := asObjectS(cplayer)
	ud := pu.UpdateDataPlayer()
	pl := ud.Player

	targ := cunit

	if pl.Field3680&0x1 != 0 {
		legacy.Nox_xxx_playerLeaveObserver_0_4E6AA0(pl)
	}
	if pl.Field3680&0x2 != 0 {
		pu.observeClear()
	}
	legacy.Nox_xxx_netNeedTimestampStatus_4174F0(pl, 2)
	pl.CameraFollow(targ)
	_ = nox_xxx_updatePlayerObserver_4E62F0
	pu.Update = legacy.Get_nox_xxx_updatePlayerObserver_4E62F0()
}

func (s *Server) nox_xxx_playerLeaveObsByObserved_4E60A0(obj server.Obj) {
	cobj := server.ToObject(obj)
	for pl := s.Players.First(); pl != nil; pl = s.Players.Next(pl) {
		if pl.CameraTarget() == cobj {
			s.PlayerLeaveMonsterObserver(pl)
		}
	}
}

func nox_xxx_playerDisconnFinish_4DE530(pli ntype.PlayerInd, a2 int8) {
	s := noxServer
	legacy.Sub_4D79A0(pli)
	sub_419EB0(pli, 0)
	legacy.Sub_4E80C0(pli)
	pl := s.Players.ByInd(pli)
	if pl != nil {
		if pl.PlayerUnit != nil {
			legacy.Sub_506740(pl.PlayerUnit)
		}
		if pl.Field4792 != 0 && pl.PlayerUnit != nil && noxflags.HasGame(noxflags.GameModeQuest) {
			s.Audio.EventObj(sound.SoundQuestPlayerExitGame, pl.PlayerUnit, 0, 0)
			pl.Field4792 = 0
			if legacy.Nox_xxx_player_4E3CE0() == 0 {
				v5 := legacy.Nox_xxx_getQuestMapFile_4D0F60()
				s.nox_game_setQuestStage_4E3CD0(0)
				s.SwitchMap(v5)
			}
			s.sub_4D7390(pl.PlayerUnit)
			s.delayedDeleteAllItems(pl.PlayerUnit)
		}
	}
	if false && !noxflags.HasGame(noxflags.GameModeChat) {
		legacy.Sub_425E90(pl, a2)
	}
	for it := s.Players.First(); it != nil; it = s.Players.Next(it) {
		u := it.PlayerUnit
		if u != nil && it != pl {
			s.Players.Nox_xxx_netUnmarkMinimapObj_417300(pli, u, 3)
			s.Players.Nox_xxx_netUnmarkMinimapObj_417300(it.PlayerIndex(), pl.PlayerUnit, 3)
		}
	}
	if pl.PlayerUnit != nil {
		for it := s.Objs.First(); it != nil; it = it.Next() {
			if it.Class().Has(object.ClassDoor) && it.ObjOwner == pl.PlayerUnit {
				it.ObjOwner = nil
			}
		}
	}
	sub_40C0E0(netstr.Global.ByPlayerInd(pli))
	legacy.Sub_4DE410(pli)
	if pl != nil {
		var buf [3]byte
		buf[0] = byte(noxnet.MSG_OUTGOING_CLIENT)
		// TODO: this can be nil for some reason; should we sent this at all?
		if pl.PlayerUnit != nil {
			binary.LittleEndian.PutUint16(buf[1:], uint16(pl.PlayerUnit.NetCode))
		}
		s.NetSendPacketXxx0(int(pli)|0x80, buf[:3], 0, 0)
	}
	if int32(a2) == 4 {
		var buf [1]byte
		buf[0] = byte(noxnet.MSG_KICK_NOTIFICATION)
		netstr.Global.ByPlayerInd(pli).Send(buf[:1], netstr.SendQueue|netstr.SendFlush)
	}
	legacy.Sub_4E55F0(pli)
	if pl != nil {
		if pl.PlayerUnit != nil {
			legacy.Nox_xxx_playerRemoveSpawnedStuff_4E5AD0(pl.PlayerUnit)
			if noxflags.HasGame(noxflags.GameModeQuest) {
				s.delayedDeleteAllItems(pl.PlayerUnit)
			} else {
				legacy.Nox_xxx_dropAllItems_4EDA40(pl.PlayerUnit)
			}
		}
		pl.Field2140 = 0
		pl.Lessons = 0
		*legacy.Get_dword_5d4594_1599592_ptr() &^= 1 << pli
		pl.Field3676 = 2
		legacy.Nox_xxx_playerUnsetStatus_417530(pl, 16)
	}
}

func (s *Server) delayedDeleteAllItems(u *server.Object) {
	var next *server.Object
	for it := u.FirstItem(); it != nil; it = next {
		next = it.NextItem()
		s.DelayedDelete(it)
	}
}

func (s *Server) sub_4D7390(u *server.Object) {
	if u == nil {
		return
	}
	if legacy.Nox_xxx_player_4E3CE0() == 0 {
		sub_4EDD00(u, object.ClassKey)
		return
	}
	var next *server.Object
	for it := u.FirstItem(); it != nil; it = next {
		next = it.NextItem()
		if !it.Class().Has(object.ClassKey) {
			continue
		}
		if to := findPlayerWithFewerKeys(); to != nil {
			legacy.Sub_4ED0C0(u, it)
			legacy.Nox_xxx_inventoryPutImpl_4F3070(to, it, 1)
			s.NetPriMsgToPlayer(to, "GeneralPrint:GainedKey", 0)
			s.Audio.EventObj(sound.SoundKeyPickup, to, 0, 0)
		}
	}
	sub_4EDD00(u, object.ClassKey)
}

func sub_4EDD00(u *server.Object, cl object.Class) {
	s := noxServer
	var next *server.Object
	for it := u.FirstItem(); it != nil; it = next {
		next = it.NextItem()
		if it.Class().Has(cl) {
			legacy.Sub_4ED0C0(u, it)
			pos := s.RandomReachablePointAround(60.0, u.Pos())
			s.CreateObjectAt(it, nil, pos)
		}
	}
}

func findPlayerWithFewerKeys() *server.Object {
	s := noxServer
	var (
		found   *server.Object
		minKeys = 99999
	)
	for pl := s.Players.First(); pl != nil; pl = s.Players.Next(pl) {
		if !pl.IsActive() || pl.PlayerUnit == nil {
			continue
		}
		u := pl.PlayerUnit
		ud := u.UpdateDataPlayer()
		if ud.Player.Field4792 != 1 {
			continue
		}
		hasKeys := 0
		for it := u.FirstItem(); it != nil; it = it.NextItem() {
			if it.Class().Has(object.ClassKey) {
				hasKeys++
			}
		}
		if hasKeys <= minKeys {
			minKeys = hasKeys
			found = u
		}
	}
	return found
}
