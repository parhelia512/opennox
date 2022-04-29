package opennox

/*
#include "defs.h"
void nox_xxx_updateFallLogic_51B870(nox_object_t* a1);
void sub_51B810(nox_object_t* a1);
void sub_537770(nox_object_t* a1);
void nox_xxx_itemApplyUpdateEffect_4FA490(nox_object_t* a1);
nox_object_t* nox_xxx_findObjectAtCursor_54AF40(nox_object_t* a1);
*/
import "C"
import (
	"github.com/noxworld-dev/opennox-lib/object"
	"github.com/noxworld-dev/opennox-lib/types"
)

func (s *Server) updateUnitsAAA() { // nox_xxx_updateUnits_51B100_A
	for _, u := range s.getPlayerUnits() {
		ud := u.updateDataPlayer()
		nox_xxx_itemApplyUpdateEffect_4FA490(u.CObj())
		ud.cursor_obj = nox_xxx_findObjectAtCursor_54AF40(u.CObj())
	}
}

func (s *Server) updateUnitsCallUpdate() { // nox_xxx_updateUnits_51B100_callUpdate
	for obj := firstServerObjectUpdatable2(); obj != nil; obj = obj.Next() {
		if obj.Flags().HasAny(object.FlagNoUpdateMask) {
			continue
		}
		obj.callUpdate()
		nox_xxx_updateFallLogic_51B870(obj.CObj())
		sub_51B810(obj.CObj())
		sub_537770(obj.CObj())
		prev := obj.Pos()
		obj.field_31_0 = obj.direction
		npos := obj.newPos()
		obj.setPrevPos(prev)
		obj.setPos(npos)
		obj.setForce(types.Pointf{})
		nox_xxx_moveUpdateSpecial_517970(obj.CObj())
	}
}
