package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

var nox_drawable_count int32 = 0
var nox_alloc_drawable *nox_alloc_class = nil
var nox_drawable_head_unk1 *nox_drawable = nil
var nox_drawable_head_unk2 *nox_drawable = nil
var nox_drawable_head_unk3 *nox_drawable = nil
var nox_drawable_head_unk4 *nox_drawable = nil

func nox_get_drawable_count() int32 {
	return nox_drawable_count
}
func nox_xxx_sprite_45A030() int32 {
	return int32(uintptr(unsafe.Pointer(nox_drawable_head_unk2)))
}
func nox_alloc_drawable_init(cnt int32) int32 {
	nox_alloc_drawable = nox_new_alloc_class(CString("drawableClass"), int32(unsafe.Sizeof(nox_drawable{})), cnt)
	return bool2int(nox_alloc_drawable != nil)
}
func nox_drawable_free() {
	nox_free_alloc_class(nox_alloc_drawable)
	nox_alloc_drawable = nil
	nox_drawable_head_unk2 = nil
	nox_drawable_head_unk1 = nil
	nox_drawable_head_unk3 = nil
	nox_drawable_head_unk4 = nil
	nox_drawable_count = 0
}
func sub_495B00(dr *nox_drawable) {
	var (
		a1 int32 = int32(uintptr(unsafe.Pointer(dr)))
		v1 *uint32
		v2 *uint32
	)
	v1 = *(**uint32)(unsafe.Pointer(uintptr(a1 + 456)))
	if v1 != nil {
		for {
			v2 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*16)))))
			sub_495B50(v1)
			nox_alloc_class_free_obj_first((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(memmap.PtrOff(6112660, 1203868)))), unsafe.Pointer(v1))
			v1 = v2
			if v2 == nil {
				break
			}
		}
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 456))) = 0
	} else {
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 456))) = 0
	}
}
func nox_xxx_spriteDelete_45A4B0(dr *nox_drawable) int32 {
	sub_495B00(dr)
	nox_alloc_class_free_obj_first(nox_alloc_drawable, unsafe.Pointer(dr))
	return func() int32 {
		p := &nox_drawable_count
		*p--
		return *p
	}()
}

var nox_parse_thing_draw_funcs [69]nox_parse_thing_draw_funcs_t = [69]nox_parse_thing_draw_funcs_t{{name: CString("NoDraw"), draw: nil, kind: 0, parse_fnc: nil}, {name: CString("DebugDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_debug_draw)), kind: 0, parse_fnc: nil}, {name: CString("StaticDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_static_draw)), kind: 1, parse_fnc: nox_things_static_draw_parse}, {name: CString("StaticRandomDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_static_random_draw)), kind: 2, parse_fnc: nox_things_static_random_draw_parse}, {name: CString("DoorDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_door_draw)), kind: 2, parse_fnc: nox_things_door_draw_parse}, {name: CString("AnimateDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_animate_draw)), kind: 3, parse_fnc: nox_things_animate_draw_parse}, {name: CString("ConditionalAnimateDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_cond_animate_draw)), kind: 4, parse_fnc: nox_things_cond_animate_draw_parse}, {name: CString("MonsterGeneratorDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_monster_gen_draw)), kind: 4, parse_fnc: nox_things_cond_animate_draw_parse}, {name: CString("VectorAnimateDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_vector_animate_draw)), kind: 5, parse_fnc: nox_things_vector_animate_draw_parse}, {name: CString("MonsterDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_monster_draw)), kind: 7, parse_fnc: nox_things_monster_draw_parse}, {name: CString("MaidenDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_maiden_draw)), kind: 7, parse_fnc: nox_things_maiden_draw_parse}, {name: CString("AnimateStateDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_animate_state_draw)), kind: 8, parse_fnc: nox_things_animate_state_draw_parse}, {name: CString("PlayerDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_player_draw)), kind: 6, parse_fnc: nox_things_player_draw_parse}, {name: CString("SlaveDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_slave_draw)), kind: 2, parse_fnc: nox_things_slave_draw_parse}, {name: CString("TriggerDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_trigger_draw)), kind: 0, parse_fnc: nil}, {name: CString("PressurePlateDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_pressure_plate_draw)), kind: 0, parse_fnc: nil}, {name: CString("LightningDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_lightning_draw)), kind: 0, parse_fnc: nil}, {name: CString("ChainLightningBoltDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_chain_lightning_bolt_draw)), kind: 0, parse_fnc: nil}, {name: CString("EnergyBoltDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_energy_bolt_draw)), kind: 0, parse_fnc: nil}, {name: CString("GreenBoltDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_green_bolt_draw)), kind: 0, parse_fnc: nil}, {name: CString("PlasmaDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_plasma_draw)), kind: 0, parse_fnc: nil}, {name: CString("YellowSparkDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_yellow_spark_draw)), kind: 0, parse_fnc: nil}, {name: CString("RedSparkDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_red_spark_draw)), kind: 0, parse_fnc: nil}, {name: CString("BlueSparkDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_blue_spark_draw)), kind: 0, parse_fnc: nil}, {name: CString("CyanSparkDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_cyan_spark_draw)), kind: 0, parse_fnc: nil}, {name: CString("GreenSparkDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_green_spark_draw)), kind: 0, parse_fnc: nil}, {name: CString("VioletSparkDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_violet_spark_draw)), kind: 0, parse_fnc: nil}, {name: CString("WhiteSparkDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_white_spark_draw)), kind: 0, parse_fnc: nil}, {name: CString("BlueRainSparkDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_blue_rain_spark_draw)), kind: 0, parse_fnc: nil}, {name: CString("DeathBallSparkDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_death_ball_spark_draw)), kind: 0, parse_fnc: nil}, {name: CString("PixieDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_pixie_draw)), kind: 0, parse_fnc: nil}, {name: CString("PixieDustDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_pixie_dust_draw)), kind: 0, parse_fnc: nil}, {name: CString("ParticleDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_particle_draw)), kind: 0, parse_fnc: nil}, {name: CString("BubbleDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_bubble_draw)), kind: 0, parse_fnc: nil}, {name: CString("VortexDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_vortex_draw)), kind: 0, parse_fnc: nil}, {name: CString("BlackPowderDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_black_powder_draw)), kind: 0, parse_fnc: nil}, {name: CString("SpiderSpitDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_spider_spit_draw)), kind: 0, parse_fnc: nil}, {name: CString("GlyphDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_glyph_draw)), kind: 3, parse_fnc: nox_things_animate_draw_parse}, {name: CString("BoulderDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_boulder_draw)), kind: 2, parse_fnc: nox_things_slave_draw_parse}, {name: CString("DrainManaDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_drain_mana_draw)), kind: 0, parse_fnc: nil}, {name: CString("GlowOrbDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_glow_orb_draw)), kind: 0, parse_fnc: nil}, {name: CString("GlowOrbMoveDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_glow_orb_move_draw)), kind: 0, parse_fnc: nil}, {name: CString("MagicDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_magic_draw)), kind: 0, parse_fnc: nil}, {name: CString("MagicMissileDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_magic_missle_draw)), kind: 0, parse_fnc: nil}, {name: CString("MagicTailLinkDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_magic_tail_link_draw)), kind: 0, parse_fnc: nil}, {name: CString("MagicMissileTailLinkDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_magic_missle_tail_link_draw)), kind: 0, parse_fnc: nil}, {name: CString("MagicSparkleDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_magic_sparkle_draw)), kind: 0, parse_fnc: nil}, {name: CString("PlayerWaypointDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_player_waypoint_draw)), kind: 0, parse_fnc: nil}, {name: CString("WeaponDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_weapon_draw)), kind: 1, parse_fnc: nox_things_static_draw_parse}, {name: CString("ArmorDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_armor_draw)), kind: 1, parse_fnc: nox_things_static_draw_parse}, {name: CString("WeaponAnimateDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_weapon_animate_draw)), kind: 3, parse_fnc: nox_things_animate_draw_parse}, {name: CString("ArmorAnimateDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_armor_animate_draw)), kind: 3, parse_fnc: nox_things_animate_draw_parse}, {name: CString("FlagDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_flag_draw)), kind: 3, parse_fnc: nox_things_animate_draw_parse}, {name: CString("BaseDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_base_draw)), kind: 1, parse_fnc: nox_things_static_draw_parse}, {name: CString("NPCDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_npc_draw)), kind: 0, parse_fnc: nil}, {name: CString("SphericalShieldDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_spherical_shield_draw)), kind: 3, parse_fnc: nox_things_animate_draw_parse}, {name: CString("SummonEffectDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_summon_effect_draw)), kind: 3, parse_fnc: nox_things_animate_draw_parse}, {name: CString("ReleasedSoulDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_released_soul_draw)), kind: 5, parse_fnc: nox_things_vector_animate_draw_parse}, {name: CString("UndeadKillerDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_undead_killer_draw)), kind: 0, parse_fnc: nil}, {name: CString("ArrowDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_arrow_draw)), kind: 2, parse_fnc: nox_things_slave_draw_parse}, {name: CString("WeakArrowDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_weak_arrow_draw)), kind: 2, parse_fnc: nox_things_slave_draw_parse}, {name: CString("ArrowTailLinkDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_arrow_tail_link_draw)), kind: 0, parse_fnc: nil}, {name: CString("WeakArrowTailLinkDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_weak_arrow_tail_link_draw)), kind: 0, parse_fnc: nil}, {name: CString("BlueRainDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_blue_rain_draw)), kind: 0, parse_fnc: nil}, {name: CString("LevelUpDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_levelup_draw)), kind: 0, parse_fnc: nil}, {name: CString("OblivionUpDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_oblivion_up_draw)), kind: 0, parse_fnc: nil}, {name: CString("RainOrbDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_rain_orb_draw)), kind: 0, parse_fnc: nil}, {name: CString("HarpoonDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_harpoon_draw)), kind: 2, parse_fnc: nox_things_slave_draw_parse}, {name: CString("HarpoonRopeDraw"), draw: unsafe.Pointer(cgoFuncAddr(nox_thing_harpoon_rope_draw)), kind: 0, parse_fnc: nil}}
var nox_parse_thing_draw_funcs_cnt int32 = int32(uint32(unsafe.Sizeof(nox_parse_thing_draw_funcs)) / uint32(unsafe.Sizeof(nox_parse_thing_draw_funcs_t{})))

func nox_xxx_spriteFromCache_45A330_drawable() unsafe.Pointer {
	if nox_drawable_head_unk4 == nil {
		return nil
	}
	nox_xxx_spriteDeleteStatic_45A4E0_drawable(nox_drawable_head_unk4)
	return nox_alloc_class_new_obj_zero(nox_alloc_drawable)
}
func nox_new_drawable_for_thing(i int32) *nox_drawable {
	var (
		v4 int32
		v1 *nox_drawable = (*nox_drawable)(nox_alloc_class_new_obj_zero(nox_alloc_drawable))
	)
	if v1 == nil {
		v1 = (*nox_drawable)(nox_xxx_spriteFromCache_45A330_drawable())
	}
	if v1 == nil {
		return nil
	}
	if nox_drawable_link_thing(v1, i) == 0 {
		return nil
	}
	var draw_fnc func(*uint32, *nox_drawable) int32
	draw_fnc = v1.draw_func
	if cgoFuncAddr(draw_fnc) == cgoFuncAddr(nox_thing_static_random_draw) {
		v4 = randomIntMinMax(0, int32(*(*uint8)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(v1.field_76)), 8)))))-1)
		nox_xxx_spriteSetFrameMB_45AB80(int32(uintptr(unsafe.Pointer(v1))), v4)
	} else if cgoFuncAddr(draw_fnc) == cgoFuncAddr(nox_thing_red_spark_draw) || cgoFuncAddr(draw_fnc) == cgoFuncAddr(nox_thing_blue_spark_draw) || cgoFuncAddr(draw_fnc) == cgoFuncAddr(nox_thing_yellow_spark_draw) || cgoFuncAddr(draw_fnc) == cgoFuncAddr(nox_thing_green_spark_draw) || cgoFuncAddr(draw_fnc) == cgoFuncAddr(nox_thing_cyan_spark_draw) {
		*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v1))), unsafe.Sizeof(uint16(0))*(26*2)))) = 35
		*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v1))), 74*4))) = 2
	} else {
		nox_xxx_spriteSetFrameMB_45AB80(int32(uintptr(unsafe.Pointer(v1))), 0)
	}
	v1.field_79 = nox_frame_xxx_2598000
	v1.field_85 = nox_frame_xxx_2598000
	nox_drawable_count++
	v1.field_120 = 0
	v1.field_121 = 0
	return v1
}
