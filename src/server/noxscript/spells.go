package noxscript

import (
	"time"

	"github.com/noxworld-dev/opennox-lib/script"
	asm "github.com/noxworld-dev/opennox-lib/script/noxscript/noxasm"
	"github.com/noxworld-dev/opennox-lib/script/noxscript/ns/effect"
	"github.com/noxworld-dev/opennox-lib/script/noxscript/ns/enchant"
	nsp "github.com/noxworld-dev/opennox-lib/script/noxscript/ns/spell"
)

func init() {
	Register(asm.BuiltinEffect, nsEffect)
	Register(asm.BuiltinCastSpellObjectObject, nsCastSpellObjObj)
	Register(asm.BuiltinCastSpellObjectLocation, nsCastSpellObjPos)
	Register(asm.BuiltinCastSpellLocationObject, nsCastSpellPosObj)
	Register(asm.BuiltinCastSpellLocationLocation, nsCastSpellPosPos)
	Register(asm.BuiltinHasEnchant, nsHasEnchant)
	Register(asm.BuiltinEnchant, nsEnchant)
	Register(asm.BuiltinEnchantOff, nsEnchantOff)
	Register(asm.BuiltinAwardSpell, nsAwardSpell)
}

func nsEffect(s VM) int {
	pos2 := s.PopPointf()
	pos := s.PopPointf()
	name := s.PopString()
	dpos := s.DPosf()
	pos = pos.Add(dpos)
	pos2 = pos2.Add(dpos)
	s.NoxScript().Effect(effect.Effect(name), pos, pos2)
	return 0
}

func nsHasEnchant(s VM) int {
	name := enchant.Enchant(s.PopString())
	obj := s.PopObjectNS()
	s.PushBool(obj != nil && obj.HasEnchant(name))
	return 0
}

func nsEnchant(s VM) int {
	sec := s.PopF32()
	name := s.PopString()
	obj := s.PopObjectNS()
	if obj != nil {
		obj.Enchant(enchant.Enchant(name), script.Time(time.Duration(float64(sec)*float64(time.Second))))
	}
	return 0
}

func nsEnchantOff(s VM) int {
	name := s.PopString()
	obj := s.PopObjectNS()
	if obj != nil {
		obj.EnchantOff(enchant.Enchant(name))
	}
	return 0
}

func nsCastSpellObjObj(s VM) int {
	targ := s.PopObjectNS()
	caster := s.PopObjectNS()
	name := s.PopString()
	var source, target script.Positioner
	if caster != nil {
		source = caster
	}
	if targ != nil {
		target = targ
	}
	s.NoxScript().CastSpell(nsp.Spell(name), source, target)
	return 0
}

func nsCastSpellObjPos(s VM) int {
	targPos := s.PopPointf()
	caster := s.PopObjectNS()
	name := s.PopString()
	var source script.Positioner
	if caster != nil {
		source = caster
	}
	s.NoxScript().CastSpell(nsp.Spell(name), source, targPos)
	return 0
}

func nsCastSpellPosObj(s VM) int {
	targ := s.PopObjectNS()
	srcPos := s.PopPointf()
	name := s.PopString()
	var target script.Positioner
	if targ != nil {
		target = targ
	}
	s.NoxScript().CastSpell(nsp.Spell(name), srcPos, target)
	return 0
}

func nsCastSpellPosPos(s VM) int {
	targPos := s.PopPointf()
	srcPos := s.PopPointf()
	name := s.PopString()
	s.NoxScript().CastSpell(nsp.Spell(name), srcPos, targPos)
	return 0
}

func nsAwardSpell(s VM) int {
	name := s.PopString()
	obj := s.PopObjectNS()
	if obj != nil {
		s.PushBool(obj.AwardSpell(nsp.Spell(name)))
	} else {
		s.PushBool(false)
	}
	return 0
}
