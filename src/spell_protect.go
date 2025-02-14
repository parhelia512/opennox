package opennox

import (
	"github.com/noxworld-dev/opennox-lib/spell"

	"github.com/noxworld-dev/opennox/v1/server"
)

func castProtectElectricity(spellID spell.ID, _, _, _ *Unit, args *spellAcceptArg, lvl int) int {
	return castBuffSpell(spellID, server.ENCHANT_PROTECT_FROM_ELECTRICITY, lvl, asUnitC(args.Obj), spellBuffConf{
		DurOpt: "ProtectElectricityEnchantDuration",
	})
}

func castProtectPoison(spellID spell.ID, _, _, _ *Unit, args *spellAcceptArg, lvl int) int {
	return castBuffSpell(spellID, server.ENCHANT_PROTECT_FROM_POISON, lvl, asUnitC(args.Obj), spellBuffConf{
		DurOpt: "ProtectPoisonEnchantDuration",
	})
}

func castProtectFire(spellID spell.ID, _, _, _ *Unit, args *spellAcceptArg, lvl int) int {
	return castBuffSpell(spellID, server.ENCHANT_PROTECT_FROM_FIRE, lvl, asUnitC(args.Obj), spellBuffConf{
		DurOpt: "ProtectFireEnchantDuration",
	})
}
