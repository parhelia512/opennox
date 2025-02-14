package opennox

/*
#include "defs.h"
void nox_xxx_cancelAllSpells_4FEE90(nox_object_t* a1);
*/
import "C"
import (
	"github.com/noxworld-dev/opennox-lib/object"
	"github.com/noxworld-dev/opennox-lib/spell"

	"github.com/noxworld-dev/opennox/v1/server"
)

func nox_xxx_warriorBerserker_53FEB0(u *Unit) {
	if u == nil {
		return
	}
	s := noxServer
	C.nox_xxx_cancelAllSpells_4FEE90(u.CObj())
	if u.HasEnchant(server.ENCHANT_CONFUSED) {
		u.Direction2 = server.Dir16(int16(int(u.Direction1) + 4*noxRndCounter1.IntClamp(-8, 8)))
	}
	if u.Class().Has(object.ClassPlayer) {
		if ud := u.UpdateDataPlayer(); ud != nil {
			ud.Field59_0 = 0
		}
	}
	nox_xxx_playerSetState_4FA020(u, 1)
	u.DisableEnchant(server.ENCHANT_INVISIBLE)
	u.DisableEnchant(server.ENCHANT_INVULNERABLE)
	s.spells.duration.CancelFor(spell.SPELL_OVAL_SHIELD, u)
	s.abilities.netAbilReportActive(u, AbilityBerserk, true)
}
