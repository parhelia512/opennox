package opennox

/*
#include "server__script__script.h"
#include "server__script__internal.h"
#include "GAME4_1.h" // for nox_xxx_scriptPrepareFoundUnit_511D70 and nox_xxx_script_511C50
extern int nox_script_count_xxx_1599640;
extern nox_script_xxx_t* nox_script_arr_xxx_1599636;
extern char* nox_script_strings[1024];
extern unsigned int nox_script_strings_xxx;
extern unsigned int nox_script_strings_cnt;
int sub_516570();
int nox_xxx_gameIsSwitchToSolo_4DB240();
size_t nox_script_readWriteWww_5417C0(FILE* a1, FILE* a2, FILE* a3);
*/
import "C"
import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"math"
	"os"
	"strings"
	"unsafe"

	"github.com/noxworld-dev/opennox-lib/common"
	"github.com/noxworld-dev/opennox-lib/ifs"
	"github.com/noxworld-dev/opennox-lib/object"
	"github.com/noxworld-dev/opennox-lib/script"
	asm "github.com/noxworld-dev/opennox-lib/script/noxscript/noxasm"
	"github.com/noxworld-dev/opennox-lib/script/noxscript/ns"
	"github.com/noxworld-dev/opennox-lib/spell"
	"github.com/noxworld-dev/opennox-lib/types"

	"github.com/noxworld-dev/opennox/v1/client/noxrender"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"github.com/noxworld-dev/opennox/v1/internal/binfile"
	"github.com/noxworld-dev/opennox/v1/internal/cryptfile"
	"github.com/noxworld-dev/opennox/v1/server"
	"github.com/noxworld-dev/opennox/v1/server/noxscript"
)

//export nox_script_activatorCancelAll_51AC60
func nox_script_activatorCancelAll_51AC60() {
	noxServer.Activators.CancelAll()
}

//export nox_script_activatorSave_51AEA0
func nox_script_activatorSave_51AEA0() C.int {
	return C.int(noxServer.SaveActivators(cryptfile.Global()))
}

//export nox_script_activatorLoad_51AF80
func nox_script_activatorLoad_51AF80() C.int {
	return C.int(noxServer.LoadActivators(cryptfile.Global()))
}

//export nox_script_activatorResolveObjs_51B0C0
func nox_script_activatorResolveObjs_51B0C0() {
	noxServer.noxScript.actResolveObjs()
}

//export nox_xxx_netGetUnitCodeServ_578AC0
func nox_xxx_netGetUnitCodeServ_578AC0(cobj *nox_object_t) C.uint {
	return C.uint(noxServer.getUnitNetCode(asObjectC(cobj)))
}

//export nox_server_scriptValToObjectPtr_511B60
func nox_server_scriptValToObjectPtr_511B60(val C.int) *C.nox_object_t {
	return noxServer.noxScript.scriptToObject(int(val)).CObj()
}

//export nox_script_get_caller
func nox_script_get_caller() unsafe.Pointer {
	return unsafe.Pointer(noxServer.noxScript.Caller().CObj())
}

//export nox_script_get_trigger
func nox_script_get_trigger() unsafe.Pointer {
	return unsafe.Pointer(noxServer.noxScript.Trigger().CObj())
}

//export nox_script_push
func nox_script_push(v C.int) {
	noxServer.noxScript.PushI32(int32(v))
}

//export nox_script_pop
func nox_script_pop() C.int {
	return C.int(noxServer.noxScript.PopI32())
}

//export nox_script_pushf
func nox_script_pushf(v C.float) {
	noxServer.noxScript.PushF32(float32(v))
}

//export nox_script_popf
func nox_script_popf() C.float {
	return C.float(noxServer.noxScript.PopF32())
}

//export nox_script_indexByEvent
func nox_script_indexByEvent(cname *C.char) int {
	name := GoString(cname)
	scripts := noxServer.noxScript.scripts()
	for i := range scripts {
		s := &scripts[i]
		if s.Name() == name {
			return i
		}
	}
	return -1
}

//export nox_script_getString_512E40
func nox_script_getString_512E40(i int) *C.char {
	if i < 0 || i >= int(C.nox_script_strings_cnt) {
		return nil
	}
	return C.nox_script_strings[i]
}

var (
	nox_script_objTelekinesisHand  int
	nox_script_objCinemaRemove     []int
	nox_xxx_imagCasterUnit_1569664 *Unit
)

var _ noxscript.VM = (*noxScript)(nil)

type noxScript struct {
	s        *Server
	dpos     image.Point
	nameSuff string
	vm       struct {
		stack   []uint32
		caller  *Object
		trigger *Object
	}
	panic noxScriptPanic
}

func (s *noxScript) Init(srv *Server) {
	s.s = srv
}

var _ = [1]struct{}{}[48-unsafe.Sizeof(noxScriptCode{})]

type noxScriptCode struct {
	name       *C.char        // 0, 0
	stack_size uint32         // 1, 4
	size_28    uint32         // 2, 8
	field_12   uint32         // 3, 12; len field_20 and field_24
	field_16   uint32         // 4, 16
	field_20   *uint32        // 5, 20
	field_24   *uint32        // 6, 24
	field_28   *uint32        // 7, 28
	data       unsafe.Pointer // 8, 32
	suff       *C.char        // 9, 36
	field_40   uint32         // 10, 40
	field_44   uint32         // 11, 44
}

func (s *noxScriptCode) Name() string {
	return GoString(s.name)
}

func (s *noxScriptCode) field28() []uint32 {
	sz := int(s.size_28) // TODO: these are only the arguments
	if sz == 0 {
		return nil
	}
	return unsafe.Slice(s.field_28, sz)
}

func (s *noxScript) scripts() []noxScriptCode {
	if C.nox_script_arr_xxx_1599636 == nil {
		return nil
	}
	return unsafe.Slice((*noxScriptCode)(unsafe.Pointer(C.nox_script_arr_xxx_1599636)), int(C.nox_script_count_xxx_1599640))
}

func (s *noxScript) Caller() *Object {
	return s.vm.caller
}

func (s *noxScript) Trigger() *Object {
	return s.vm.trigger
}

func (s *noxScript) resetStack() {
	s.vm.stack = s.vm.stack[:0]
}

func (s *noxScript) saveStack() int {
	return len(s.vm.stack)
}

func (s *noxScript) adjustStack(prev, sz int) {
	if n := len(s.vm.stack); n != prev+sz {
		if sz != 0 {
			if n != 0 {
				v := s.PopU32()
				s.vm.stack = s.vm.stack[:prev]
				s.PushU32(v)
			} else {
				s.vm.stack = s.vm.stack[:prev]
				s.PushU32(0)
			}
		} else {
			s.vm.stack = s.vm.stack[:prev]
		}
	}
}

func (s *noxScript) stackAt(i int) uint32 {
	if i < 0 || i >= len(s.vm.stack) {
		return 0
	}
	return s.vm.stack[i]
}

func (s *noxScript) PushU32(v uint32) {
	s.vm.stack = append(s.vm.stack, v)
}

func (s *noxScript) PushI32(v int32) {
	s.vm.stack = append(s.vm.stack, uint32(v))
}

func (s *noxScript) PushF32(v float32) {
	s.vm.stack = append(s.vm.stack, math.Float32bits(v))
}

func (s *noxScript) PushBool(v bool) {
	s.vm.stack = append(s.vm.stack, uint32(bool2int(v)))
}

func (s *noxScript) AddString(str string) uint32 {
	if int(C.nox_script_strings_cnt) >= len(C.nox_script_strings) {
		return uint32(C.nox_script_strings_cnt - 1)
	}
	i := uint32(C.nox_script_strings_cnt)
	C.nox_script_strings[i] = CString(str)
	C.nox_script_strings_cnt++
	return i
}

func (s *noxScript) PushString(str string) {
	s.vm.stack = append(s.vm.stack, s.AddString(str))
}

func (s *noxScript) PopI32() int32 {
	n := len(s.vm.stack)
	if n == 0 {
		return 0
	}
	v := s.vm.stack[n-1]
	s.vm.stack = s.vm.stack[:n-1]
	return int32(v)
}

func (s *noxScript) PopU32() uint32 {
	n := len(s.vm.stack)
	if n == 0 {
		return 0
	}
	v := s.vm.stack[n-1]
	s.vm.stack = s.vm.stack[:n-1]
	return v
}

func (s *noxScript) PopF32() float32 {
	return math.Float32frombits(s.PopU32())
}

func (s *noxScript) PopBool() bool {
	return s.PopU32() != 0
}

func (s *noxScript) GetString(i uint32) string {
	if i < 0 || i >= uint32(C.nox_script_strings_cnt) {
		return ""
	}
	return GoString(C.nox_script_strings[i])
}

func (s *noxScript) PopString() string {
	return s.GetString(s.PopU32())
}

func (s *noxScript) PopPoint() image.Point {
	y := s.PopI32()
	x := s.PopI32()
	return image.Point{X: int(x), Y: int(y)}
}

func (s *noxScript) PopPointf() types.Pointf {
	y := s.PopF32()
	x := s.PopF32()
	return types.Pointf{X: x, Y: y}
}

func (s *noxScript) PopObject() *Object {
	return s.scriptToObject(int(nox_script_pop()))
}

func (s *noxScript) PopGroup() *mapGroup {
	return s.s.mapGroups.GroupByInd(int(s.PopI32()))
}

func (s *noxScript) nox_xxx_scriptRunFirst_507290() {
	s.resetBuiltin()
	if scripts := s.scripts(); len(scripts) >= 2 {
		sc := unsafe.Slice((*int32)(unsafe.Pointer(scripts[1].field_28)), 4)
		sc[0] = -2
		sc[1] = -1
		sc[2] = 1
		sc[3] = 0
		if nox_xxx_gameIsSwitchToSolo_4DB240() == 0 {
			if err := s.callByIndex(1, nil, nil); err != nil {
				scriptLog.Println(err)
			}
		}
	} else {
		scriptLog.Printf("noxscript: no init function")
	}
}

func (s *noxScript) OnEvent(event script.EventType) {
	scripts := s.scripts()
	for i := range scripts {
		sc := &scripts[i]
		name := GoString(sc.name)
		if strings.HasPrefix(name, string(event)) {
			if err := s.callByIndex(i, nil, nil); err != nil {
				scriptLog.Println(err)
			}
		}
	}
}

func (s *Server) getUnitNetCode(p noxObject) int {
	obj := toObject(p)
	if obj == nil {
		return 0
	}
	if obj.NetCode >= 0x8000 {
		return 0
	}
	ext := int(obj.Extent)
	if ext >= 0x8000 {
		return 0
	}
	if !obj.Class().HasAny(object.ClassClientPersist | object.ClassImmobile) {
		return int(obj.NetCode)
	}
	ext |= 0x8000
	return ext
}

func (s *noxScript) scriptToObject(val int) *Object {
	if val == -1 {
		obj := s.Caller()
		if obj == nil || obj.Flags().Has(object.FlagDestroyed) {
			return nil
		}
		return obj
	}
	if val == -2 {
		obj := s.Trigger()
		if obj == nil || obj.Flags().Has(object.FlagDestroyed) {
			return nil
		}
		return obj
	}
	if obj := asObjectC(C.nox_xxx_script_511C50(C.int(val))); obj != nil {
		return obj
	}

	for obj := asObjectS(s.s.Objs.List); obj != nil; obj = obj.Next() {
		if !obj.Flags().Has(object.FlagDestroyed) && obj.ScriptID == val {
			C.nox_xxx_scriptPrepareFoundUnit_511D70(obj.CObj())
			return obj
		}
		for sub := obj.FirstItem(); sub != nil; sub = sub.NextItem() {
			if !sub.Flags().Has(object.FlagDestroyed) && sub.ScriptID == val {
				C.nox_xxx_scriptPrepareFoundUnit_511D70(sub.CObj())
				return sub
			}
		}
	}
	for obj := asObjectS(s.s.Objs.Pending); obj != nil; obj = obj.Next() {
		if !obj.Flags().Has(object.FlagDestroyed) && obj.ScriptID == val {
			C.nox_xxx_scriptPrepareFoundUnit_511D70(obj.CObj())
			return obj
		}
	}
	return nil
}

func (s *noxScript) resetBuiltin() {
	s.dpos = image.Point{}
	s.nameSuff = ""
}

func (s *noxScript) NameSuff() string {
	return s.nameSuff
}

func (s *noxScript) DPos() image.Point {
	return s.dpos
}

func (s *noxScript) DPosf() types.Pointf {
	return types.Point2f(s.dpos)
}

func (s *noxScript) scriptNameSuff(i int) string {
	scripts := s.scripts()
	return GoString(scripts[i].suff)
}

func (s *noxScript) scriptField40(i int) int {
	scripts := s.scripts()
	return int(scripts[i].field_40)
}

func (s *noxScript) scriptField44(i int) int {
	scripts := s.scripts()
	return int(scripts[i].field_44)
}

func (s *noxScript) noxScriptEndGame(v int) {
	dword_587000_311372 = v
	dword_5d4594_2516476 |= 1 << v
	nox_xxx_cliPlayMapIntro_44E0B0(1)
	sub_413960()
	sub_477530(false)
	nox_client_quit_4460C0()
}

//export sub_5165D0
func sub_5165D0() {
	s := &noxServer.noxScript
	sub5165D0(int(s.PopU32()))
}

func sub5165D0(which int) {
	*memmap.PtrUint32(0x5D4594, 2386828) = uint32(which - 1)
	sub_413A00(1)
	noxClient.r.FadeInScreen(25, true, func() {
		C.sub_516570()
	})
}

func (s *Server) CinemaPlayers(enable bool) {
	if nox_script_objTelekinesisHand == 0 {
		nox_script_objTelekinesisHand = s.ObjectTypeID("TelekinesisHand")
	}
	const (
		perc       = 0.16
		fadeOutDur = 30
		fadeInDur  = 10
	)
	if !enable {
		if noxClient.r.FadeOutCinema(perc, fadeOutDur, color.Black) {
			sub_477530(false)
		}
		for it := s.FirstServerObject(); it != nil; it = it.Next() {
			if int(it.TypeInd) == nox_script_objTelekinesisHand {
				if f := it.Flags(); f.Has(object.FlagNoCollide) {
					it.ObjFlags = uint32(f &^ object.FlagNoCollide)
				}
			}
		}
		return
	}
	inFade := noxClient.r.CheckFade(noxrender.FadeInCinemaKey)
	if noxClient.r.FadeInCinema(perc, fadeInDur, color.Black) {
		sub_477530(true)
	}
	if inFade {
		return
	}
	if len(nox_script_objCinemaRemove) == 0 {
		for _, name := range []string{
			"ToxicCloud", "SmallToxicCloud", "Meteor", "SmallFist", "MediumFist", "LargeFist", "Pixie",
		} {
			nox_script_objCinemaRemove = append(nox_script_objCinemaRemove, s.ObjectTypeID(name))
		}
	}

	var next *Object
	for it := asObjectS(s.Objs.UpdatableList2); it != nil; it = next {
		next = it.Next()
		if int(it.TypeInd) != int(memmap.Uint32(0x5D4594, 2386900)) {
			it.Delete()
		}
	}

	next = nil
	for it := s.FirstServerObject(); it != nil; it = next {
		next = it.Next()
		if it.OwnerC() != nil {
			for _, id := range nox_script_objCinemaRemove {
				if int(it.TypeInd) == id {
					it.Delete()
					break
				}
			}
		} else {
			if int(it.TypeInd) == nox_script_objTelekinesisHand {
				if f := it.Flags(); !f.Has(object.FlagNoCollide) {
					it.ObjFlags = uint32(f | object.FlagNoCollide)
				}
			}
			s.spells.duration.CancelFor(spell.SPELL_WALL, it)
			s.spells.duration.CancelFor(spell.SPELL_MANA_BOMB, it)
			s.spells.duration.CancelFor(spell.SPELL_LIGHTNING, it)
			s.spells.duration.CancelFor(spell.SPELL_CHAIN_LIGHTNING, it)
			s.spells.duration.CancelFor(spell.SPELL_DRAIN_MANA, it)
			s.spells.duration.CancelFor(spell.SPELL_FORCE_OF_NATURE, it)
			s.spells.duration.CancelFor(spell.SPELL_GREATER_HEAL, it)
			s.spells.duration.CancelFor(spell.SPELL_CHANNEL_LIFE, it)
			s.spells.duration.CancelFor(spell.SPELL_CHARM, it)
			s.spells.duration.CancelFor(spell.SPELL_BLINK, it)
			s.spells.duration.CancelFor(spell.SPELL_SWAP, it)
			s.spells.duration.CancelFor(spell.SPELL_TURN_UNDEAD, it)
			s.spells.duration.CancelFor(spell.SPELL_PLASMA, it)
			s.spells.duration.CancelFor(spell.SPELL_SUMMON_BAT, it)
		}
	}
}

//export nox_setImaginaryCaster
func nox_setImaginaryCaster() C.int {
	s := noxServer
	nox_xxx_imagCasterUnit_1569664 = s.newObjectByTypeID("ImaginaryCaster").AsUnit()
	if nox_xxx_imagCasterUnit_1569664 == nil {
		return 0
	}
	s.createObjectAt(nox_xxx_imagCasterUnit_1569664, nil, types.Pointf{X: 128 * common.GridStep, Y: 128 * common.GridStep})
	return 1
}

//export nox_script_readWriteZzz_541670
func nox_script_readWriteZzz_541670(cpath, cpath2, cdst *C.char) C.int {
	if cpath == nil {
		return 0
	}
	if cpath2 == nil {
		return 0
	}
	if cdst == nil {
		return 0
	}
	path, path2, dst := GoString(cpath), GoString(cpath2), GoString(cdst)
	f1, err := ifs.Open(path)
	if err != nil {
		return 0
	}
	defer f1.Close()

	f2, err := ifs.Open(path2)
	if err != nil {
		return 0
	}
	defer f2.Close()

	v6, _ := binfile.FileSize(f1)
	if v6 == 0 {
		ifs.Remove(dst)
		ifs.Rename(path2, dst)
		return 1
	}

	v7, _ := binfile.FileSize(f2)
	if v7 == 0 {
		ifs.Remove(dst)
		ifs.Rename(path, dst)
		return 1
	}

	df, err := ifs.OpenFile(dst, os.O_RDWR)
	if err != nil {
		return 0
	}
	defer df.Close()
	C.nox_script_readWriteWww_5417C0(newFileHandle(binfile.NewFile(f1)), newFileHandle(binfile.NewFile(f2)), newFileHandle(binfile.NewFile(df)))
	return 1
}

func (s *noxScript) actRun() {
	scripts := s.scripts()
	s.s.Activators.EachTriggered(s.s.Frame(), func(it server.ActivatorArgs) {
		if scripts[it.Callback].size_28 != 0 {
			s.PushU32(it.Arg)
		}
		if err := s.callByIndex(it.Callback, asObjectS(it.Caller), asObjectS(it.Trigger)); err != nil {
			scriptLog.Println(err)
		}
	})
}

func (s *noxScript) actResolveObjs() {
	s.s.Activators.ResolveObjs(func(id int) *server.Object {
		return s.scriptToObject(id).SObj()
	})
}

func (s *noxScript) NoxScript() ns.Implementation {
	return s.s.NoxScript()
}

func (s *noxScript) PopWallGroupNS() ns.WallGroupObj {
	g := s.PopGroup()
	if g == nil || g.Type() != mapGroupWalls {
		return nil
	}
	return nsWallGroup{s.s, g}
}

func (s *noxScript) PopWaypointNS() ns.WaypointObj {
	wp := s.s.getWaypointByInd(int(s.PopI32()))
	if wp == nil {
		return nil
	}
	return wp
}

func (s *noxScript) PopWpGroupNS() ns.WaypointGroupObj {
	g := s.PopGroup()
	if g == nil || g.Type() != mapGroupWaypoints {
		return nil
	}
	return nsWpGroup{s.s, g}
}

func (s *noxScript) PopObjectNS() ns.Obj {
	obj := s.scriptToObject(int(nox_script_pop()))
	if obj == nil {
		return nil
	}
	return nsObj{obj}
}

func (s *noxScript) PopObjGroupNS() ns.ObjGroup {
	g := s.PopGroup()
	if g == nil || g.Type() != mapGroupObjects {
		return nil
	}
	return nsObjGroup{s.s, g}
}

func (s *noxScript) PushHandleNS(obj ns.Handle) {
	if obj == nil {
		s.PushI32(0)
	} else {
		s.PushI32(int32(obj.ScriptID()))
	}
}

//export nox_xxx_scriptCallByEventBlock_502490
func nox_xxx_scriptCallByEventBlock_502490(a1 unsafe.Pointer, a2, a3 unsafe.Pointer, eventCode int32) unsafe.Pointer {
	return noxServer.scriptCallback((*server.ScriptCallback)(a1), asObject(a2), asObject(a3), noxEventType(eventCode))
}

type noxScriptCallback struct {
	Block    *server.ScriptCallback
	Caller4  *Object
	Trigger8 *Object
}

var (
	noxScriptCallbacks []noxScriptCallback
)

func (s *Server) scriptPushCallback(b *server.ScriptCallback, caller, trigger *Object) {
	noxScriptCallbacks = append(noxScriptCallbacks, noxScriptCallback{
		Block: b, Caller4: caller, Trigger8: trigger,
	})
}

func (s *Server) scriptPopCallback(b *server.ScriptCallback, caller, trigger *Object) {
	for i := 0; i < len(noxScriptCallbacks); i++ {
		it := &noxScriptCallbacks[i]
		if it.Block == b && it.Caller4 == caller && it.Trigger8 == trigger {
			noxScriptCallbacks = append(noxScriptCallbacks[:i], noxScriptCallbacks[i+1:]...)
		}
	}
}

func (s *Server) scriptCallback(b *server.ScriptCallback, caller, trigger *Object, eventCode noxEventType) unsafe.Pointer {
	s.nox_script_callByEventcgo(eventCode, unsafe.Pointer(caller.CObj()), unsafe.Pointer(trigger.CObj()))
	*memmap.PtrUint32(0x5D4594, 1599076) = 0
	flags := b.Flags
	if flags&0x1 != 0 {
		return nil
	}
	sind := int(b.Func)
	if sind == -1 {
		return nil
	}
	if flags&0x2 != 0 {
		b.Flags |= 0x1
	}
	if s.noxScript.saveStack() != 0 {
		s.scriptPushCallback(b, caller, trigger)
		return memmap.PtrOff(0x5D4594, 1599076)
	}
	if err := s.noxScript.callByIndexObj(sind, caller, trigger); err != nil {
		scriptLog.Println(err)
	}
	scripts := s.noxScript.scripts()
	if scripts[sind].stack_size != 0 {
		*memmap.PtrUint32(0x5D4594, 1599076) = uint32(s.noxScript.PopI32())
	}
	s.noxScript.resetStack()
	if C.nox_script_strings_xxx < C.nox_script_strings_cnt {
		for i := int(C.nox_script_strings_xxx); i < int(C.nox_script_strings_cnt); i++ {
			StrFree(C.nox_script_strings[i])
			C.nox_script_strings[i] = nil
		}
	}
	C.nox_script_strings_cnt = C.nox_script_strings_xxx
	s.scriptPopCallback(b, caller, trigger)
	if len(noxScriptCallbacks) > 0 {
		s.scriptCallback(noxScriptCallbacks[0].Block, noxScriptCallbacks[0].Caller4, noxScriptCallbacks[0].Trigger8, 0)
	}
	return memmap.PtrOff(0x5D4594, 1599076)
}

//export nox_script_callByIndex_507310
func nox_script_callByIndex_507310(index int, a2 unsafe.Pointer, a3 unsafe.Pointer) {
	if err := noxServer.noxScript.callByIndexObj(index, asObject(a2), asObject(a3)); err != nil {
		scriptLog.Println(err)
	}
}

func (s *noxScript) callByIndex(index int, caller, trigger noxObject) error {
	return s.callByIndexObj(index, toObject(caller), toObject(trigger))
}

func (s *noxScript) callByIndexObj(index int, caller, trigger *Object) error {
	if C.nox_script_arr_xxx_1599636 == nil {
		return errors.New("no map scripts")
	}
	scripts := s.scripts()
	if index < 0 || index >= len(scripts) {
		return fmt.Errorf("invalid function index: %d (%x)", index, index)
	}
	sc := scripts[index]

	s.vm.caller = caller
	s.vm.trigger = trigger

	args := sc.field28()
	for i := 0; i < int(sc.size_28); i++ {
		args[i] = s.PopU32()
	}
	bstack := s.saveStack()
	data := sc.data
	readInt := func(p *unsafe.Pointer) int32 {
		v := *(*int32)(*p)
		*p = unsafe.Add(*p, 4)
		return v
	}
	readFloat := func(p *unsafe.Pointer) float32 {
		v := readInt(p)
		return math.Float32frombits(uint32(v))
	}
	for {
		var (
			sa1 int32   = 0
			sa2 int32   = 0
			sa3 int32   = 0
			sa4 int32   = 0
			sf1 float32 = 0
			sf2 float32 = 0
		)
		switch asm.Op(readInt(&data)) {
		case asm.OpLoadVarInt, asm.OpLoadVarString:
			sa1 = readInt(&data)
			sa2 = readInt(&data)
			if sa1 != 0 {
				if scripts[1].field_24 != nil && scripts[1].field_28 != nil {
					sa2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_24), 4*uintptr(sa2))))
					sa2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_28), 4*uintptr(sa2))))
				} else {
					sa2 = 0
				}
				s.PushI32(sa2)
			} else if sa2 < 0 {
				sa2 = -sa2
				sa2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_24), 4*uintptr(sa2))))
				sa2 = int32(*(*uint32)(unsafe.Add(trigger.Field190, sa2*4)))
				s.PushI32(sa2)
			} else if sc.field_24 != nil && sc.field_28 != nil {
				sa2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(sc.field_24), 4*uintptr(sa2))))
				sa2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(sc.field_28), 4*uintptr(sa2))))
				s.PushI32(sa2)
			} else {
				s.PushI32(0)
			}
			continue
		case asm.OpLoadVarFloat:
			sa1 = readInt(&data)
			sa2 = readInt(&data)
			if sa1 != 0 {
				if scripts[1].field_24 != nil && scripts[1].field_28 != nil {
					sa2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_24), 4*uintptr(sa2))))
					sf1 = *(*float32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_28), 4*uintptr(sa2)))))
				} else {
					sf1 = 0
				}
			} else if sa2 < 0 {
				sa2 = -sa2
				sa2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_24), 4*uintptr(sa2))))
				sf1 = *(*float32)(unsafe.Add(trigger.Field190, sa2*4))
			} else if sc.field_24 != nil && sc.field_28 != nil {
				sa2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(sc.field_24), 4*uintptr(sa2))))
				sf1 = *(*float32)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(sc.field_28))) + uint32(sa2*4))))
			} else {
				sf1 = 0
			}
			s.PushF32(sf1)
			continue
		case asm.OpLoadVarPtr:
			sa1 = readInt(&data)
			sa2 = readInt(&data)
			if sa1 != 0 {
				sa3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_20), 4*uintptr(sa2))))
				sa4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_24), 4*uintptr(sa2))))
			} else if sa2 < 0 {
				sa2 = -sa2
				sa3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_20), 4*uintptr(sa2))))
				sa4 = int32(-*(*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_24), 4*uintptr(sa2))))
			} else if sc.field_20 != nil && sc.field_24 != nil {
				sa3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(sc.field_20), 4*uintptr(sa2))))
				sa4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(sc.field_24), 4*uintptr(sa2))))
			} else {
				sa3 = 0
				sa4 = 0
			}
			if sa3 > 1 {
				s.PushI32(sa3)
			}
			s.PushI32(sa1)
			s.PushI32(sa4)
			continue
		case asm.OpPushInt, asm.OpPushString:
			sa1 = readInt(&data)
			s.PushI32(sa1)
			continue
		case asm.OpPushFloat:
			sf1 = readFloat(&data)
			s.PushF32(sf1)
			continue
		case asm.OpIntAdd:
			sa1 = s.PopI32()
			sa2 = s.PopI32()
			s.PushI32(sa2 + sa1)
			continue
		case asm.OpFloatAdd:
			sf1 = s.PopF32()
			sf2 = s.PopF32()
			sf1 += sf2
			s.PushF32(sf1)
			continue
		case asm.OpIntSub:
			sa1 = s.PopI32()
			sa2 = s.PopI32()
			s.PushI32(sa2 - sa1)
			continue
		case asm.OpFloatSub:
			sf1 = s.PopF32()
			sf2 = s.PopF32()
			sf2 -= sf1
			s.PushF32(sf2)
			continue
		case asm.OpIntMul:
			sa1 = s.PopI32()
			sa2 = s.PopI32()
			s.PushI32(sa1 * sa2)
			continue
		case asm.OpFloatMul:
			sf1 = s.PopF32()
			sf2 = s.PopF32()
			sf1 *= sf2
			s.PushF32(sf1)
			continue
		case asm.OpIntDiv:
			sa1 = s.PopI32()
			sa2 = s.PopI32()
			s.PushI32(sa2 / sa1)
			continue
		case asm.OpFloatDiv:
			sf1 = s.PopF32()
			sf2 = s.PopF32()
			sf2 /= sf1
			s.PushF32(sf2)
			continue
		case asm.OpIntMod:
			sa1 = s.PopI32()
			sa2 = s.PopI32()
			s.PushI32(sa2 % sa1)
			continue
		case asm.OpIntAnd:
			sa1 = s.PopI32()
			sa2 = s.PopI32()
			s.PushI32(sa2 & sa1)
			continue
		case asm.OpIntOr:
			sa1 = s.PopI32()
			sa2 = s.PopI32()
			s.PushI32(sa2 | sa1)
			continue
		case asm.OpIntXOr:
			sa1 = s.PopI32()
			sa2 = s.PopI32()
			s.PushI32(sa2 ^ sa1)
			continue
		case asm.OpJump:
			data = unsafe.Pointer(uintptr(uint32(uintptr(sc.data)) + uint32(readInt(&data)*4)))
			continue
		case asm.OpJumpIf:
			sa1 = readInt(&data)
			if s.PopBool() {
				data = unsafe.Pointer(uintptr(uint32(uintptr(sc.data)) + uint32(sa1*4)))
			}
			continue
		case asm.OpJumpIfNot:
			sa1 = readInt(&data)
			if !s.PopBool() {
				data = unsafe.Pointer(uintptr(uint32(uintptr(sc.data)) + uint32(sa1*4)))
			}
			continue
		case asm.OpStoreInt, asm.OpStoreString:
			sa1 = s.PopI32()
			sa2 = s.PopI32()
			if s.PopI32() != 0 {
				if scripts[1].field_28 != nil {
					*(*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_28), 4*uintptr(sa2))) = uint32(sa1)
				}
			} else if sa2 < 0 {
				*(*uint32)(unsafe.Add(trigger.Field190, -sa2*4)) = uint32(sa1)
			} else if sc.field_28 != nil {
				*(*uint32)(unsafe.Add(unsafe.Pointer(sc.field_28), 4*uintptr(sa2))) = uint32(sa1)
			}
			s.PushI32(sa1)
			continue
		case asm.OpStoreFloat:
			sf1 = s.PopF32()
			sa2 = s.PopI32()
			if s.PopI32() != 0 {
				if scripts[1].field_28 != nil {
					*(*float32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_28), 4*uintptr(sa2))))) = sf1
				}
			} else if sa2 < 0 {
				*(*float32)(unsafe.Add(trigger.Field190, -sa2*4)) = sf1
			} else if sc.field_28 != nil {
				*(*float32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(sc.field_28), 4*uintptr(sa2))))) = sf1
			}
			s.PushF32(sf1)
			continue
		case asm.OpStoreIntMul:
			sa1 = s.PopI32()
			sa2 = s.PopI32()
			if s.PopI32() != 0 {
				if scripts[1].field_28 != nil {
					*(*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_28), 4*uintptr(sa2))) *= uint32(sa1)
					s.PushI32(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_28), 4*uintptr(sa2)))))
				} else {
					s.PushI32(0)
				}
			} else if sa2 < 0 {
				*(*uint32)(unsafe.Add(trigger.Field190, -sa2*4)) *= uint32(sa1)
				s.PushI32(int32(*(*uint32)(unsafe.Add(trigger.Field190, -sa2*4))))
			} else if sc.field_28 != nil {
				*(*uint32)(unsafe.Add(unsafe.Pointer(sc.field_28), 4*uintptr(sa2))) *= uint32(sa1)
				s.PushI32(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(sc.field_28), 4*uintptr(sa2)))))
			} else {
				s.PushI32(0)
			}
			continue
		case asm.OpStoreFloatMul:
			sf1 = s.PopF32()
			sa1 = s.PopI32()
			if s.PopI32() != 0 {
				if scripts[1].field_28 != nil {
					*(*float32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_28), 4*uintptr(sa1))))) *= sf1
					s.PushF32(*(*float32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_28), 4*uintptr(sa1))))))
				} else {
					s.PushF32(0)
				}
			} else if sa1 < 0 {
				*(*float32)(unsafe.Add(trigger.Field190, -sa1*4)) *= sf1
				s.PushF32(*(*float32)(unsafe.Add(trigger.Field190, -sa1*4)))
			} else if sc.field_28 != nil {
				*(*float32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(sc.field_28), 4*uintptr(sa1))))) *= sf1
				s.PushF32(*(*float32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(sc.field_28), 4*uintptr(sa1))))))
			} else {
				s.PushF32(0)
			}
			continue
		case asm.OpStoreIntDiv:
			sa1 = s.PopI32()
			sa2 = s.PopI32()
			if s.PopI32() != 0 {
				if scripts[1].field_28 != nil {
					*(*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_28), 4*uintptr(sa2))) /= uint32(sa1)
					s.PushI32(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_28), 4*uintptr(sa2)))))
				} else {
					s.PushI32(0)
				}
			} else if sa2 < 0 {
				*(*uint32)(unsafe.Add(trigger.Field190, -sa2*4)) /= uint32(sa1)
				s.PushI32(int32(*(*uint32)(unsafe.Add(trigger.Field190, -sa2*4))))
			} else if sc.field_28 != nil {
				*(*uint32)(unsafe.Add(unsafe.Pointer(sc.field_28), 4*uintptr(sa2))) /= uint32(sa1)
				s.PushI32(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(sc.field_28), 4*uintptr(sa2)))))
			} else {
				s.PushI32(0)
			}
			continue
		case asm.OpStoreFloatDiv:
			sf1 = s.PopF32()
			sa1 = s.PopI32()
			if s.PopI32() != 0 {
				if scripts[1].field_28 != nil {
					*(*float32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_28), 4*uintptr(sa1))))) /= sf1
					s.PushF32(*(*float32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_28), 4*uintptr(sa1))))))
				} else {
					s.PushF32(0)
				}
			} else if sa1 < 0 {
				*(*float32)(unsafe.Add(trigger.Field190, -sa1*4)) /= sf1
				s.PushF32(*(*float32)(unsafe.Add(trigger.Field190, -sa1*4)))
			} else if sc.field_28 != nil {
				*(*float32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(sc.field_28), 4*uintptr(sa1))))) /= sf1
				s.PushF32(*(*float32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(sc.field_28), 4*uintptr(sa1)))))))
			} else {
				s.PushF32(0)
			}
			continue
		case asm.OpStoreIntAdd:
			sa1 = s.PopI32()
			sa2 = s.PopI32()
			if s.PopI32() != 0 {
				if scripts[1].field_28 != nil {
					*(*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_28), 4*uintptr(sa2))) += uint32(sa1)
					s.PushI32(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_28), 4*uintptr(sa2)))))
				} else {
					s.PushI32(0)
				}
			} else if sa2 < 0 {
				*(*uint32)(unsafe.Add(trigger.Field190, -sa2*4)) += uint32(sa1)
				s.PushI32(int32(*(*uint32)(unsafe.Add(trigger.Field190, -sa2*4))))
			} else if sc.field_28 != nil {
				*(*uint32)(unsafe.Add(unsafe.Pointer(sc.field_28), 4*uintptr(sa2))) += uint32(sa1)
				s.PushI32(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(sc.field_28), 4*uintptr(sa2)))))
			} else {
				s.PushI32(0)
			}
			continue
		case asm.OpStoreFloatAdd:
			sf1 = s.PopF32()
			sa1 = s.PopI32()
			if s.PopI32() != 0 {
				if scripts[1].field_28 != nil {
					*(*float32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_28), 4*uintptr(sa1))))) += sf1
					s.PushF32(*(*float32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_28), 4*uintptr(sa1))))))
				} else {
					s.PushF32(0)
				}
			} else if sa1 < 0 {
				*(*float32)(unsafe.Add(trigger.Field190, -sa1*4)) += sf1
				s.PushF32(*(*float32)(unsafe.Add(trigger.Field190, -sa1*4)))
			} else if sc.field_28 != nil {
				*(*float32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(sc.field_28), 4*uintptr(sa1))))) += sf1
				s.PushF32(*(*float32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(sc.field_28), 4*uintptr(sa1))))))
			} else {
				s.PushF32(0)
			}
			continue
		case asm.OpStoreStringAdd:
			sa1 = s.PopI32()
			sa2 = s.PopI32()
			sa3 = s.PopI32()
			if sa3 != 0 {
				if scripts[1].field_28 != nil {
					sa4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_28), 4*uintptr(sa2))))
				} else {
					sa4 = 0
				}
			} else if sa2 < 0 {
				sa4 = int32(*(*uint32)(unsafe.Add(trigger.Field190, -sa2*4)))
			} else {
				sa4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(sc.field_28), 4*uintptr(sa2))))
			}
			sa4 = int32(s.AddString(s.GetString(uint32(sa4)) + s.GetString(uint32(sa1))))
			if sa3 != 0 {
				if scripts[1].field_28 != nil {
					*(*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_28), 4*uintptr(sa2))) = uint32(sa4)
				}
			} else if sa2 < 0 {
				*(*uint32)(unsafe.Add(trigger.Field190, -sa2*4)) = uint32(sa4)
			} else if sc.field_28 != nil {
				*(*uint32)(unsafe.Add(unsafe.Pointer(sc.field_28), 4*uintptr(sa2))) = uint32(sa4)
			}
			s.PushI32(sa4)
			continue
		case asm.OpStoreIntSub:
			sa1 = s.PopI32()
			sa2 = s.PopI32()
			if s.PopI32() != 0 {
				if scripts[1].field_28 != nil {
					*(*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_28), 4*uintptr(sa2))) -= uint32(sa1)
					s.PushI32(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_28), 4*uintptr(sa2)))))
				} else {
					s.PushI32(0)
				}
			} else if sa2 < 0 {
				*(*uint32)(unsafe.Add(trigger.Field190, -sa2*4)) -= uint32(sa1)
				s.PushI32(int32(*(*uint32)(unsafe.Add(trigger.Field190, -sa2*4))))
			} else if sc.field_28 != nil {
				*(*uint32)(unsafe.Add(unsafe.Pointer(sc.field_28), 4*uintptr(sa2))) -= uint32(sa1)
				s.PushI32(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(sc.field_28), 4*uintptr(sa2)))))
			} else {
				s.PushI32(0)
			}
			continue
		case asm.OpStoreFloatSub:
			sf1 = s.PopF32()
			sa1 = s.PopI32()
			if s.PopI32() != 0 {
				if scripts[1].field_28 != nil {
					*(*float32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_28), 4*uintptr(sa1))))) -= sf1
					s.PushF32(*(*float32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_28), 4*uintptr(sa1))))))
				} else {
					s.PushF32(0)
				}
			} else if sa1 < 0 {
				*(*float32)(unsafe.Add(trigger.Field190, -sa1*4)) -= sf1
				s.PushI32(int32(*(*uint32)(unsafe.Add(trigger.Field190, -sa1*4))))
			} else if sc.field_28 != nil {
				*(*float32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(sc.field_28), 4*uintptr(sa1))))) -= sf1
				s.PushF32(*(*float32)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(sc.field_28))) + uint32(sa1*4)))))
			} else {
				s.PushF32(0)
			}
			continue
		case asm.OpStoreIntMod:
			sa1 = s.PopI32()
			sa2 = s.PopI32()
			if s.PopI32() != 0 {
				if scripts[1].field_28 != nil {
					*(*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_28), 4*uintptr(sa2))) %= uint32(sa1)
					s.PushI32(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_28), 4*uintptr(sa2)))))
				} else {
					s.PushI32(0)
				}
			} else if sa2 < 0 {
				*(*uint32)(unsafe.Add(trigger.Field190, -sa2*4)) %= uint32(sa1)
				s.PushI32(int32(*(*uint32)(unsafe.Add(trigger.Field190, -sa2*4))))
			} else if sc.field_28 != nil {
				*(*uint32)(unsafe.Add(unsafe.Pointer(sc.field_28), 4*uintptr(sa2))) %= uint32(sa1)
				s.PushI32(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(sc.field_28), 4*uintptr(sa2)))))
			} else {
				s.PushI32(0)
			}
			continue
		case asm.OpIntEq:
			sa1 = s.PopI32()
			sa2 = s.PopI32()
			s.PushBool(sa2 == sa1)
			continue
		case asm.OpFloatEq:
			sf1 = s.PopF32()
			sf2 = s.PopF32()
			s.PushBool(sf2 == sf1)
			continue
		case asm.OpStringEq:
			sa1 = s.PopI32()
			sa2 = s.PopI32()
			s.PushBool(s.GetString(uint32(sa2)) == s.GetString(uint32(sa1)))
			continue
		case asm.OpIntLSh:
			sa1 = s.PopI32()
			sa2 = s.PopI32()
			s.PushI32(sa2 << sa1)
			continue
		case asm.OpIntRSh:
			sa1 = s.PopI32()
			sa2 = s.PopI32()
			s.PushI32(sa2 >> sa1)
			continue
		case asm.OpIntLt:
			sa1 = s.PopI32()
			sa2 = s.PopI32()
			s.PushBool(sa2 < sa1)
			continue
		case asm.OpFloatLt:
			sf1 = s.PopF32()
			sf2 = s.PopF32()
			s.PushBool(sf2 < sf1)
			continue
		case asm.OpStringLt:
			sa1 = s.PopI32()
			sa2 = s.PopI32()
			s.PushBool(s.GetString(uint32(sa2)) < s.GetString(uint32(sa1)))
			continue
		case asm.OpIntGt:
			sa1 = s.PopI32()
			sa2 = s.PopI32()
			s.PushBool(sa2 > sa1)
			continue
		case asm.OpFloatGt:
			sf1 = s.PopF32()
			sf2 = s.PopF32()
			s.PushBool(sf2 > sf1)
			continue
		case asm.OpStringGt:
			sa1 = s.PopI32()
			sa2 = s.PopI32()
			s.PushBool(s.GetString(uint32(sa2)) > s.GetString(uint32(sa1)))
			continue
		case asm.OpIntLte:
			sa1 = s.PopI32()
			sa2 = s.PopI32()
			s.PushBool(sa2 <= sa1)
			continue
		case asm.OpFloatLte:
			sf1 = s.PopF32()
			sf2 = s.PopF32()
			s.PushBool(sf2 <= sf1)
			continue
		case asm.OpStringLte:
			sa1 = s.PopI32()
			sa2 = s.PopI32()
			s.PushBool(s.GetString(uint32(sa2)) <= s.GetString(uint32(sa1)))
			continue
		case asm.OpIntGte:
			sa1 = s.PopI32()
			sa2 = s.PopI32()
			s.PushBool(sa2 >= sa1)
			continue
		case asm.OpFloatGte:
			sf1 = s.PopF32()
			sf2 = s.PopF32()
			s.PushBool(sf2 >= sf1)
			continue
		case asm.OpStringGte:
			sa1 = s.PopI32()
			sa2 = s.PopI32()
			s.PushBool(s.GetString(uint32(sa2)) >= s.GetString(uint32(sa1)))
			continue
		case asm.OpIntNeq:
			sa1 = s.PopI32()
			sa2 = s.PopI32()
			s.PushBool(sa2 != sa1)
			continue
		case asm.OpFloatNeq:
			sf1 = s.PopF32()
			sf2 = s.PopF32()
			s.PushBool(sf2 != sf1)
			continue
		case asm.OpStringNeq:
			sa1 = s.PopI32()
			sa2 = s.PopI32()
			s.PushBool(s.GetString(uint32(sa2)) != s.GetString(uint32(sa1)))
			continue
		case asm.OpBoolAnd:
			y := s.PopBool()
			x := s.PopBool()
			s.PushBool(x && y)
			continue
		case asm.OpBoolOr:
			y := s.PopBool()
			x := s.PopBool()
			s.PushBool(x || y)
			continue
		case asm.OpStoreIntLSh:
			sa1 = s.PopI32()
			sa2 = s.PopI32()
			if s.PopI32() != 0 {
				if scripts[1].field_28 != nil {
					*(*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_28), 4*uintptr(sa2))) <<= uint32(sa1)
					s.PushI32(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_28), 4*uintptr(sa2)))))
				} else {
					s.PushI32(0)
				}
			} else if sa2 < 0 {
				*(*uint32)(unsafe.Add(trigger.Field190, -sa2*4)) <<= uint32(sa1)
				s.PushI32(int32(*(*uint32)(unsafe.Add(trigger.Field190, -sa2*4))))
			} else if sc.field_28 != nil {
				*(*uint32)(unsafe.Add(unsafe.Pointer(sc.field_28), 4*uintptr(sa2))) <<= uint32(sa1)
				s.PushI32(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(sc.field_28), 4*uintptr(sa2)))))
			} else {
				s.PushI32(0)
			}
			continue
		case asm.OpStoreIntRSh:
			sa1 = s.PopI32()
			sa2 = s.PopI32()
			if s.PopI32() != 0 {
				if scripts[1].field_28 != nil {
					*(*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_28), 4*uintptr(sa2))) >>= uint32(sa1)
					s.PushI32(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_28), 4*uintptr(sa2)))))
				} else {
					s.PushI32(0)
				}
			} else if sa2 < 0 {
				*(*uint32)(unsafe.Add(trigger.Field190, -sa2*4)) >>= uint32(sa1)
				s.PushI32(int32(*(*uint32)(unsafe.Add(trigger.Field190, -sa2*4))))
			} else if sc.field_28 != nil {
				*(*uint32)(unsafe.Add(unsafe.Pointer(sc.field_28), 4*uintptr(sa2))) >>= uint32(sa1)
				s.PushI32(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(sc.field_28), 4*uintptr(sa2)))))
			} else {
				s.PushI32(0)
			}
			continue
		case asm.OpStoreIntAnd:
			sa1 = s.PopI32()
			sa2 = s.PopI32()
			if s.PopI32() != 0 {
				if scripts[1].field_28 != nil {
					*(*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_28), 4*uintptr(sa2))) &= uint32(sa1)
					s.PushI32(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_28), 4*uintptr(sa2)))))
				} else {
					s.PushI32(0)
				}
			} else if sa2 < 0 {
				*(*uint32)(unsafe.Add(trigger.Field190, -sa2*4)) &= uint32(sa1)
				s.PushI32(int32(*(*uint32)(unsafe.Add(trigger.Field190, -sa2*4))))
			} else if sc.field_28 != nil {
				*(*uint32)(unsafe.Add(unsafe.Pointer(sc.field_28), 4*uintptr(sa2))) &= uint32(sa1)
				s.PushI32(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(sc.field_28), 4*uintptr(sa2)))))
			} else {
				s.PushI32(0)
			}
			continue
		case asm.OpStoreIntOr:
			sa1 = s.PopI32()
			sa2 = s.PopI32()
			if s.PopI32() != 0 {
				if scripts[1].field_28 != nil {
					*(*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_28), 4*uintptr(sa2))) |= uint32(sa1)
					s.PushI32(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_28), 4*uintptr(sa2)))))
				} else {
					s.PushI32(0)
				}
			} else if sa2 < 0 {
				*(*uint32)(unsafe.Add(trigger.Field190, -sa2*4)) |= uint32(sa1)
				s.PushI32(int32(*(*uint32)(unsafe.Add(trigger.Field190, -sa2*4))))
			} else if sc.field_28 != nil {
				*(*uint32)(unsafe.Add(unsafe.Pointer(sc.field_28), 4*uintptr(sa2))) |= uint32(sa1)
				s.PushI32(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(sc.field_28), 4*uintptr(sa2)))))
			} else {
				s.PushI32(0)
			}
			continue
		case asm.OpStoreIntXOr:
			sa1 = s.PopI32()
			sa2 = s.PopI32()
			if s.PopI32() != 0 {
				if scripts[1].field_28 != nil {
					*(*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_28), 4*uintptr(sa2))) ^= uint32(sa1)
					s.PushI32(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_28), 4*uintptr(sa2)))))
				} else {
					s.PushI32(0)
				}
			} else if sa2 < 0 {
				*(*uint32)(unsafe.Add(trigger.Field190, -sa2*4)) ^= uint32(sa1)
				s.PushI32(int32(*(*uint32)(unsafe.Add(trigger.Field190, -sa2*4))))
			} else if sc.field_28 != nil {
				*(*uint32)(unsafe.Add(unsafe.Pointer(sc.field_28), 4*uintptr(sa2))) ^= uint32(sa1)
				s.PushI32(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(sc.field_28), 4*uintptr(sa2)))))
			} else {
				s.PushI32(0)
			}
			continue
		case asm.OpBoolNot:
			x := s.PopBool()
			s.PushBool(!x)
			continue
		case asm.OpIntNot:
			sa1 = s.PopI32()
			s.PushI32(^sa1)
			continue
		case asm.OpIntNeg:
			sa1 = s.PopI32()
			s.PushI32(-sa1)
			continue
		case asm.OpFloatNeg:
			sf1 = -s.PopF32()
			s.PushF32(sf1)
			continue
		case asm.OpIndexInt:
			sa1 = s.PopI32()
			sa2 = s.PopI32()
			sa3 = s.PopI32()
			sa4 = s.PopI32()
			failed := sa1 < 0 || sa1 >= sa4
			if sa3 != 0 {
				if scripts[1].field_28 != nil {
					sa4 = *(*int32)(unsafe.Add(unsafe.Pointer(scripts[1].field_28), 4*uintptr(sa1+sa2)))
				} else {
					sa4 = 0
				}
			} else if sa2 < 0 {
				sa4 = *(*int32)(unsafe.Add(trigger.Field190, (sa2-sa1)*4))
			} else if sc.field_28 != nil {
				sa4 = *(*int32)(unsafe.Add(unsafe.Pointer(sc.field_28), 4*uintptr(sa1+sa2)))
			} else {
				sa4 = 0
			}
			s.PushI32(sa4)
			if failed {
				return fmt.Errorf("noxscript: index out of bounds: [%d:%d]", sa1, sa4)
			}
			continue
		case asm.OpIndexFloat:
			sa1 = s.PopI32()
			sa2 = s.PopI32()
			sa3 = s.PopI32()
			sa4 = s.PopI32()
			failed := sa1 < 0 || sa1 > sa4
			if sa3 != 0 {
				if scripts[1].field_28 != nil {
					sf1 = *(*float32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(scripts[1].field_28), 4*uintptr(sa1+sa2)))))
				} else {
					sf1 = 0
				}
			} else if sa2 < 0 {
				sf1 = *(*float32)(unsafe.Add(trigger.Field190, (sa2-sa1)*4))
			} else if sc.field_28 != nil {
				sf1 = *(*float32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(sc.field_28), 4*uintptr(sa1+sa2)))))
			} else {
				sf1 = 0
			}
			s.PushF32(sf1)
			if failed {
				return fmt.Errorf("noxscript: index out of bounds: [%d:%d]", sa1, sa4)
			}
			continue
		case asm.OpIndexPtr:
			sa1 = s.PopI32()
			sa2 = s.PopI32()
			sa3 = s.PopI32()
			sa4 = s.PopI32()
			failed := sa1 < 0 || sa1 > sa4
			s.PushI32(sa3)
			if sa2 < 0 {
				s.PushI32(sa2 - sa1)
			} else {
				s.PushI32(sa2 + sa1)
			}
			if failed {
				return fmt.Errorf("noxscript: index out of bounds: [%d:%d]", sa1, sa4)
			}
			continue
		case asm.OpCallBuiltin:
			fi := asm.Builtin(readInt(&data))
			if err := s.callBuiltin(index, fi); err != nil {
				return fmt.Errorf("in %q: %w", sc.Name(), err)
			}
		case asm.OpCallScript:
			ind := int(readInt(&data))
			if err := s.callByIndexObj(ind, caller, trigger); err != nil {
				return fmt.Errorf("in %q: %w", sc.Name(), err)
			}
		case asm.OpStringAdd:
			sa1 = s.PopI32()
			sa2 = s.PopI32()
			sa3 = int32(s.AddString(s.GetString(uint32(sa2)) + s.GetString(uint32(sa1))))
			s.PushI32(sa3)
		default:
			s.adjustStack(bstack, int(sc.stack_size))
			return nil
		}
	}
}
