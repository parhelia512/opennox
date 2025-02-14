package opennox

import (
	"github.com/noxworld-dev/opennox-lib/script"
	"github.com/noxworld-dev/opennox-lib/script/noxscript/ns"

	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
)

var _ ns.Game = noxScriptImpl{}

func (s *Server) NoxScript() ns.Implementation {
	return noxScriptNS{s}
}

func (s noxScriptImpl) NoxScript() ns.Implementation {
	return s.s.NoxScript()
}

type noxScriptNS struct {
	s *Server
}

func (s noxScriptNS) NewTimer(dt script.Duration, fnc ns.Func, args ...any) ns.Timer {
	panic("TODO: implement me")
	//id := s.s.noxScript.newScriptTimer(s.s.AsFrames(dt), fnc, args)
	//return nsTimer{s: s.s, id: id}
}

func (s noxScriptNS) RandomFloat(min float32, max float32) float32 {
	return float32(noxRndCounter1.FloatClamp(float64(min), float64(max)))
}

func (s noxScriptNS) Random(min int, max int) int {
	return noxRndCounter1.IntClamp(min, max)
}

func (s noxScriptNS) StopScript(val any) {
	panic(val)
}

func (s noxScriptNS) AutoSave() {
	if noxflags.HasGame(noxflags.GameModeCoop) {
		sub_4DB130("AUTOSAVE")
		sub_4DB170(1, nil, 0)
	}
}

func (s noxScriptNS) StartupScreen(which int) {
	//TODO implement me
	panic("implement me")
}

func (s noxScriptNS) DeathScreen(which int) {
	*memmap.PtrUint32(0x5D4594, 2386832) = 0
	sub5165D0(which)
}

type nsTimer struct {
	s  *Server
	id uint32
}

func (t nsTimer) ScriptID() int {
	return int(t.id)
}

func (t nsTimer) Cancel() bool {
	return t.s.Activators.Cancel(t.id)
}
