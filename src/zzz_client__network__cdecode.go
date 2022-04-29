package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"math"
	"unsafe"
)

func nox_xxx_netOnPacketRecvCli_48EA70(a1 int32, data *uint8, sz int32) int32 {
	var (
		v5   int64
		k    int32
		v9   int32
		v10  int32
		v11  uint16
		v12  uint32
		v13  int32
		v14  int32
		v15  *uint8
		v16  int32
		v17  int32
		v18  int32
		v19  int32
		v20  int32
		v21  int32
		v22  int32
		v24  int32
		v25  int32
		v26  int32
		v27  *uint32
		v28  uint8
		v29  int32
		v30  uint8
		v31  uint32
		v32  int32
		v33  int32
		v34  uint32
		v35  *byte
		j    *byte
		v37  *byte
		v38  *libc.WChar
		v39  *byte
		v40  uint32
		v41  *libc.WChar
		v42  int32
		v43  *byte
		v44  int32
		v45  int32
		v46  int32
		v47  int32
		v48  uint32
		v49  int32
		v50  int32
		v51  int32
		v52  int32
		v53  uint32
		v54  int32
		v55  int32
		v56  int32
		v57  int32
		v58  int32
		v59  int32
		v60  int32
		v62  int32
		v63  int32
		v64  int32
		v65  int32
		v66  int8
		v67  int32
		v68  int32
		v69  int32
		v70  int32
		v71  int32
		v72  int32
		v73  *uint32
		v74  int32
		v75  int32
		v76  uint32
		v77  int32
		v78  int32
		v79  int32
		v80  int32
		v81  *int32
		v82  int32
		v83  int32
		v84  int32
		v85  int32
		v86  int32
		v87  int8
		v88  int32
		v89  int32
		v90  int32
		v91  int32
		v92  int32
		v93  int32
		v94  int32
		v95  int32
		v96  int32
		v97  int32
		v98  int32
		v99  int32
		v100 int32
		v101 int32
		v102 int32
		v103 int32
		v104 int32
		v105 int32
		v106 int32
		v107 int32
		v108 *byte
		v109 int32
		v110 int32
		v111 *uint32
		v112 *uint32
		v113 int32
		v114 *uint32
		v115 int32
		v116 int32
		v117 uint32
		v118 uint16
		v119 *byte
		v120 int32
		v121 int32
		v122 int32
		v123 int32
		v124 int32
		v125 int32
		v126 *byte
		v127 *byte
		v128 *libc.WChar
		v129 uint32
		v130 int32
		v132 int32
		v133 int32
		v134 int32
		v135 int32
		v137 int32
		v138 int32
		v139 int32
		v140 int32
		v141 int32
		v142 *byte
		v143 *libc.WChar
		v144 *byte
		v145 *uint32
		v146 *libc.WChar
		v147 int32
		v148 *libc.WChar
		v149 *libc.WChar
		v150 int32
		v151 *libc.WChar
		v152 int32
		v153 int32
		v154 *byte
		v155 *libc.WChar
		v156 *libc.WChar
		v157 *libc.WChar
		v158 *libc.WChar
		v159 *libc.WChar
		v160 *byte
		v161 *libc.WChar
		v162 *uint32
		v163 *libc.WChar
		v164 int32
		v165 *libc.WChar
		v166 *libc.WChar
		v167 *libc.WChar
		v168 *libc.WChar
		v169 *libc.WChar
		v170 *libc.WChar
		v171 *libc.WChar
		v172 *byte
		v173 *libc.WChar
		v174 *uint32
		v175 *libc.WChar
		v176 int32
		v177 *libc.WChar
		v178 int32
		v179 int32
		v180 int32
		v181 int32
		v182 int32
		v183 int32
		v184 uint32
		v185 uint32
		v186 uint16
		v187 uint16
		v188 *uint32
		v189 int32
		v190 int32
		v191 uint16
		v192 uint16
		v193 uint32
		v194 uint32
		v195 *uint32
		v196 int32
		v197 int32
		v198 int32
		v199 int32
		v200 int32
		v201 *libc.WChar
		v202 int32
		v203 *libc.WChar
		v204 *libc.WChar
		v205 *libc.WChar
		v206 *byte
		v207 *byte
		v208 int32
		v213 *libc.WChar
		v214 int32
		v215 int32
		v216 *byte
		v217 *libc.WChar
		v218 *uint32
		v219 int32
		v220 *libc.WChar
		v221 uint32
		v222 int32
		v223 int32
		v224 int32
		v225 int32
		v226 int32
		v227 int32
		v228 int32
		v229 int32
		v230 int32
		v231 int32
		v232 *uint32
		v233 int8
		v234 int32
		v235 int32
		v236 int32
		v237 int32
		v238 *uint32
		v239 int32
		v240 int32
		v241 int32
		v242 int32
		v243 int32
		v244 int32
		v245 int32
		v246 *uint32
		v247 int8
		v248 int16
		v249 int32
		v250 bool
		v251 int32
		v252 uint32
		v253 int32
		v254 int32
		v255 int32
		v256 int32
		v257 int32
		v258 int32
		v259 int32
		v260 int32
		v261 int32
		v262 int32
		v263 int32
		v264 int32
		v265 int32
		v266 int32
		v267 int32
		v268 int32
		v269 int32
		v270 int32
		v271 *libc.WChar
		v272 *byte
		v273 *uint32
		v274 uint32
		v275 uint32
		v276 int32
		v277 int32
		v278 *uint32
		v279 int32
		v280 *libc.WChar
		v281 uint32
		v282 uint32
		v283 int32
		v284 int32
		v285 uint16
		v286 uint16
		v287 int32
		v288 *byte
		v289 *uint8
		v290 *libc.WChar
		v291 int32
		v292 int32
		v293 int32
		v294 int32
		v295 int32
		v296 int32
		v297 int32
		v298 int32
		v299 *byte
		v300 *byte
		v301 int32
		v302 *uint32
		v303 *byte
		v304 *libc.WChar
		v305 *libc.WChar
		v306 *byte
		v307 *libc.WChar
		v308 *libc.WChar
		v309 *byte
		v310 *libc.WChar
		v311 *libc.WChar
		v312 *libc.WChar
		v314 *libc.WChar
		v315 int8
		v316 uint8
		v317 int32
		v318 int32
		v319 uint16
		v320 int32
		v321 int32
		v322 int32
		v323 int8
		v324 uint32
		v325 int8
		v326 int8
		v327 int32
		v328 float32
		v329 int32
		v330 int32
		v331 int32
		v332 float32
		v333 int32
		v334 *libc.WChar
		v335 *libc.WChar
		v336 *libc.WChar
		v337 int32
		v338 int32
		v339 int32
		v340 float32
		v341 int32
		v342 int8
		v343 *byte
		v344 uint32
		v345 int32
		v346 int8
		v347 int8
		v348 int32
		v349 int32
		v350 int32
		v351 int32
		v352 int32
		v353 int32
		v354 *libc.WChar
		v355 int32
		v356 uint32
		v357 int32
		v358 *int32
		v359 int8
		v360 int32
		i    int32
		v362 int32
		v363 int32
		v364 uint32
		v365 [6]byte
		v367 int32
		v368 int32
		v370 int32
		v371 int32
		v372 int32  = 0
		v373 uint32 = 0
		v374 int32
		v376 int32
		v377 int2
		v378 int2
		v379 int2
		v380 int2
		v381 [20]uint8
		v382 [8]int32
		v383 [5]int32
		v384 [2]int32
		v385 [20]byte
		v386 [20]libc.WChar
		v387 [20]libc.WChar
		v388 [64]libc.WChar
		v389 [128]libc.WChar
		v390 [128]libc.WChar
		v391 [128]libc.WChar
		v392 [128]libc.WChar
		v393 [128]libc.WChar
		v394 [128]libc.WChar
		v395 [64]libc.WChar
		v396 [64]libc.WChar
		v397 [100]libc.WChar
		v398 [128]libc.WChar
		v400 [256]libc.WChar
		v401 [128]libc.WChar
		v402 [128]libc.WChar
		v403 [256]libc.WChar
		v404 [256]libc.WChar
		v405 [256]libc.WChar
		v406 [256]libc.WChar
	)
	if sz != 0 {
	}
	var end *uint8 = (*uint8)(unsafe.Add(unsafe.Pointer(data), sz))
	v364 = 0
	sub_470A80()
	if uintptr(unsafe.Pointer(data)) >= uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), sz)))) {
		return 1
	}
	for uintptr(unsafe.Pointer(data)) < uintptr(unsafe.Pointer(end)) {
		var (
			old *uint8 = data
			op  int32  = int32(*(*uint8)(unsafe.Add(unsafe.Pointer(data), 0)))
		)
		switch op {
		case 33:
		case 39:
			fallthrough
		case 40:
			fallthrough
		case 42:
			v9 = 1
			if op == 40 {
				nox_frame_xxx_2598000 = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))
				*memmap.PtrUint32(6112660, 1200800) = nox_frame_xxx_2598000
				v364 = nox_frame_xxx_2598000
				*memmap.PtrUint32(6112660, 1200808) = uint32(int32(uint16(nox_frame_xxx_2598000)) >> 14)
				if *memmap.PtrUint32(0x8531A0, 2576) != 0 {
					nox_xxx_playerUnsetStatus_417530((*nox_playerInfo)(unsafe.Pointer(uintptr(*memmap.PtrInt32(0x8531A0, 2576)))), 64)
				}
				sub_43C650()
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 5))
				break
			} else if op != 39 {
				v14 = int32(*(*uint16)(unsafe.Pointer(uintptr(v373 + 1))))
				if uint32(uint16(int16(v14))) < (*memmap.PtrUint32(6112660, 1200800) + uint32(sub_43C790())) {
					nox_frame_xxx_2598000 = uint32(v14)
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 5))
				break
			}
			v373 = uint32(uintptr(unsafe.Pointer(data)))
			if *memmap.PtrUint32(0x8531A0, 2576) != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 3680))))&64 != 0 {
				return 1
			}
			v10 = int32(nox_frame_xxx_2598000)
			v11 = *(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))
			v12 = uint32(v11) + (nox_frame_xxx_2598000 & 0xFFFF0000)
			v13 = int32(v11) >> 14
			if uint32(v13) != *memmap.PtrUint32(6112660, 1200808) {
				if v13 == ((int32(*memmap.PtrUint8(6112660, 1200808)) + 1) & 3) {
					*memmap.PtrUint32(6112660, 1200808) = uint32(v13)
					if v13 == 0 {
						v12 += 0x10000
					}
				} else {
					v9 = 0
				}
			}
			if v12 < nox_frame_xxx_2598000 {
				v9 = 0
			}
			if !noxflags.HasGame(noxflags.GameHost) && v9 == 1 {
				nox_frame_xxx_2598000 = v12
				*memmap.PtrUint32(6112660, 1200800) = v12
			}
			v364 = v12
			if !noxflags.HasGame(noxflags.GameHost) && v9 == 0 {
				nox_perfmon_latePackets_2618900--
				*memmap.PtrUint32(0x85B3FC, 120)++
				return 1
			}
			if nox_frame_xxx_2598000 > uint32(v10+1) {
				nox_perfmon_latePackets_2618900 += nox_frame_xxx_2598000 - uint32(v10)
			}
			sub_43C650()
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
		case 43:
			k = int32(dword_5d4594_1200804)
			if *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 37)))) > uint32(*(*int32)(unsafe.Pointer(&dword_5d4594_1200804))) {
				nox_xxx_setMapCRC_40A360(int32(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 37))))))
				nox_xxx_gameClearAll_467DF0(1)
				nox_xxx_gameSetMapPath_409D70((*byte)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))
				nox_xxx_mapSetCrcMB_409B10(int32(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 33))))))
				if !noxflags.HasGame(noxflags.GameHost) {
					noxflags.UnsetGame(noxflags.GameFlag4)
					if dword_5d4594_2650652 != 0 {
						sub_41D6C0()
					}
				}
				noxflags.SetGame(noxflags.GameFlag24)
				dword_5d4594_1200804 = nox_frame_xxx_2598000
				nox_xxx_gameSetCliConnected_43C720(0)
				sub_49C7A0()
				nox_xxx_guiServerOptionsHide_4597E0(0)
				sub_44A400()
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 41))
		case 44:
			v42 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			nox_player_netCode_85319C = uint32(v42)
			v43 = (*byte)(unsafe.Pointer(nox_common_playerInfoNew_416F60(v42)))
			if v43 != nil {
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v43))), unsafe.Sizeof(uint32(0))*517))) = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))
				*memmap.PtrUint32(0x8531A0, 2576) = uint32(uintptr(unsafe.Pointer(v43)))
			}
			if !noxflags.HasGame(noxflags.GameHost) {
				sub_57B920(memmap.PtrOff(6112660, 1198020))
			}
			dword_5d4594_1200804 = 0
			nox_perfmon_latePackets_2618900 = 0
			*memmap.PtrUint32(0x85B3FC, 120) = 0
			nox_perfmon_ping_2614264 = 0
			dword_5d4594_1200832 = 0
			nox_xxx_cliSetSettingsAcquired_4169D0(0)
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 7))
		case 45:
			var (
				playerID int32           = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
				pl       *nox_playerInfo = nox_common_playerInfoNew_416F60(playerID)
			)
			if pl == nil {
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 129))
				break
			}
			if !noxflags.HasGame(noxflags.GameHost) {
				pl.netCode = uint32(playerID)
				pl.field_2136 = uint32(*(*int16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 100)))))
				pl.field_2140 = uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 102)))))
				pl.field_0 = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 104))))
				pl.field_4 = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 108))))
				pl.field_2152 = uint32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 116))))
				pl.field_2156 = uint32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 117))))
				libc.StrCpy(&pl.field_2096[0], (*byte)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 119)))))
				pl.field_3680 |= *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 112))))
				libc.MemCpy(unsafe.Pointer(&pl.info), unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))), 97)
				nox_swprintf(&pl.name_final[0], libc.CWString("%s%s"), &pl.info.name[0], &pl.info.name_suff[0])
				if dword_5d4594_2650652 != 0 {
					pl.field_2108 = 0
					sub_41D670(&pl.field_2096[0])
				}
				nox_xxx_playerInitColors_461460(pl)
			}
			sub_457140(playerID, &pl.name_final[0])
			sub_455920(int32(uintptr(unsafe.Pointer(&pl.name_final[0]))))
			v213 = strMan.GetStringInFile("PlayerJoined", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
			var fnameBuf [100]libc.WChar
			nox_swprintf(&fnameBuf[0], v213, &pl.name_final[0])
			if nox_xxx_gameGetPlayState_4356B0() == 3 {
				nox_xxx_printCentered_445490(&fnameBuf[0])
			}
			if uint32(playerID) == nox_player_netCode_85319C && nox_wcscmp((*libc.WChar)(memmap.PtrOff(0x85B3FC, 12204)), &pl.name_final[0]) != 0 {
				dword_5d4594_1200832 = 1
			}
			OnLibraryNotice_263(uint32(uintptr(unsafe.Pointer(pl))))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 129))
		case 46:
			if nox_client_isConnected_43C700() != 0 {
				v214 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
				v215 = v214
				v216 = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(v214)))
				if v216 != nil {
					sub_456DF0(v215)
					sub_455950((*libc.WChar)(unsafe.Add(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(v216))), unsafe.Sizeof(libc.WChar(0))*2352)))
					v217 = strMan.GetStringInFile("PlayerLeft", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
					nox_swprintf(&v397[0], v217, (*byte)(unsafe.Add(unsafe.Pointer(v216), 4704)))
					*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v216))), unsafe.Sizeof(uint32(0))*523))) = 0
					v218 = nox_xxx_objGetTeamByNetCode_418C80(v215)
					v219 = int32(uintptr(unsafe.Pointer(v218)))
					OnLibraryNotice_264(uint32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v216), 4704))))))
					if v218 != nil && nox_xxx_servObjectHasTeam_419130(int32(uintptr(unsafe.Pointer(v218)))) != 0 {
						nox_xxx_netChangeTeamMb_419570(v219, v215)
					}
				} else {
					v220 = strMan.GetStringInFile("UnknownLeft", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
					nox_swprintf(&v397[0], v220)
				}
				if nox_xxx_gameGetPlayState_4356B0() == 3 {
					nox_xxx_printCentered_445490(&v397[0])
				}
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
		case 47:
			nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			if nox_client_isConnected_43C700() != 0 {
				if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ADDITIONAL_NETWORK_TEST)) {
					nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
				}
				*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_xxx_spriteCreate_48E970(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))), uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))), int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 5))))), int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 7)))))))))
				if uint32(int32(v5)) != 0 {
					k = int32(nox_frame_xxx_2598000)
					*(*uint32)(unsafe.Pointer(uintptr(v5 + 288))) = nox_frame_xxx_2598000
				}
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 9))
		case 48:
			nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			if nox_client_isConnected_43C700() == 0 {
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 11))
				break
			}
			if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ADDITIONAL_NETWORK_TEST)) {
				nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			}
			if int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))) != 0 || int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))) != 0 {
				*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_xxx_spriteCreate_48E970(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))), uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))), int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 5))))), int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 7)))))))))
				v29 = int32(v5)
				if uint32(int32(v5)) != 0 {
					*(*uint32)(unsafe.Pointer(uintptr(v5 + 288))) = nox_frame_xxx_2598000
					nox_xxx_spriteSetFrameMB_45AB80(int32(v5), int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 10)))))
					*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v5))), 0)) = uint8(int8((int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 9)))) >> 4) & 7))
					*(*uint8)(unsafe.Pointer(uintptr(v29 + 297))) = uint8(int8(v5))
					if int32(uint8(int8(v5))) > 3 {
						*(*uint8)(unsafe.Pointer(uintptr(v29 + 297))) = uint8(int8(v5 + 1))
					}
					v30 = uint8(int8(int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 9)))) & 15))
					*((*uint8)(unsafe.Add(unsafe.Pointer(data), 9))) = v30
					if *(*uint32)(unsafe.Pointer(uintptr(v29 + 276))) != uint32(v30) {
						*(*uint32)(unsafe.Pointer(uintptr(v29 + 316))) = nox_frame_xxx_2598000
						*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*1)) = uint32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 9))))
						*(*uint32)(unsafe.Pointer(uintptr(v29 + 276))) = *(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*1))
					}
				}
				k = int32(nox_player_netCode_85319C)
				if uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))) == nox_player_netCode_85319C && sub_416120(9) {
					nox_xxx_cliUpdateCameraPos_435600(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 5))))), int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 7))))))
				}
			} else {
				nox_xxx_cliUpdateCameraPos_435600(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 5))))), int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 7))))))
				nox_xxx_setKeybTimeout_4160D0(9)
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 11))
		case 49:
			v44 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v45 = v44
			if nox_client_isConnected_43C700() == 0 {
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
				break
			}
			if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ADDITIONAL_NETWORK_TEST)) {
				nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			}
			*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(func() *uint32 {
				if nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))) != 0 {
					return nox_xxx_netSpriteByCodeStatic_45A720(v45)
				}
				return nox_xxx_netSpriteByCodeDynamic_45A6F0(v45)
			}())))
			v46 = int32(v5)
			if uint32(int32(v5)) == 0 {
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
				break
			}
			if cgoFuncAddr(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(v5 + 300))), (*func(*int32, int32) int32)(nil))) == cgoFuncAddr(nox_thing_animate_draw) {
				v47 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 304))))
				if v47 != 0 {
					if *(*uint32)(unsafe.Pointer(uintptr(v47 + 12))) == 1 {
						data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
						break
					}
				}
			}
			v48 = nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v331 = v46
			if v48 == 0 {
				goto LABEL_210
			}
			nox_xxx_cliDestroyObj_45A9A0(v46)
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
		case 50:
			v49 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v50 = v49
			if nox_client_isConnected_43C700() == 0 {
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
				break
			}
			if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ADDITIONAL_NETWORK_TEST)) {
				nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			}
			*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(func() *uint32 {
				if nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))) != 0 {
					return nox_xxx_netSpriteByCodeStatic_45A720(v50)
				}
				return nox_xxx_netSpriteByCodeDynamic_45A6F0(v50)
			}())))
			v51 = int32(v5)
			if uint32(int32(v5)) == 0 {
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
				break
			}
			if cgoFuncAddr(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(v5 + 300))), (*func(*int32, int32) int32)(nil))) == cgoFuncAddr(nox_thing_animate_draw) {
				v52 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 304))))
				if v52 != 0 {
					if *(*uint32)(unsafe.Pointer(uintptr(v52 + 12))) == 1 {
						data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
						break
					}
				}
			}
			if uint32(v51) == *memmap.PtrUint32(0x852978, 8) {
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
				break
			}
			v53 = nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			goto LABEL_208
		case 51:
			v54 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v55 = v54
			if nox_client_isConnected_43C700() == 0 {
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
				break
			}
			if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_NET_DEBUG)) {
				nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			}
			*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(func() *uint32 {
				if nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))) != 0 {
					return nox_xxx_netSpriteByCodeStatic_45A720(v55)
				}
				return nox_xxx_netSpriteByCodeDynamic_45A6F0(v55)
			}())))
			v51 = int32(v5)
			if uint32(int32(v5)) == 0 {
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
				break
			}
			k = 1
			*(*uint32)(unsafe.Pointer(uintptr(v5 + 484))) = 1
			*(*uint32)(unsafe.Pointer(uintptr(v5 + 480))) = 1
			*(*uint32)(unsafe.Pointer(uintptr(v5 + 488))) = 1
			if nox_client_fadeObjects_80836 != 0 {
				if uint32(int32(v5)) != *memmap.PtrUint32(0x852978, 8) {
					nox_xxx_spriteTransparentDecay_49B950((*uint32)(unsafe.Pointer(uintptr(v5))), int32(nox_gameFPS))
				}
			} else if cgoFuncAddr(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(v5 + 300))), (*func(*int32, int32) int32)(nil))) != cgoFuncAddr(nox_thing_animate_draw) || (func() int32 {
				v56 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 304))))
				return v56
			}()) == 0 || *(*uint32)(unsafe.Pointer(uintptr(v56 + 12))) != 1 {
				if uint32(v51) != *memmap.PtrUint32(0x852978, 8) {
					v53 = nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
				LABEL_208:
					v331 = v51
					if v53 != 0 {
						nox_xxx_cliDestroyObj_45A9A0(v51)
					} else {
					LABEL_210:
						nox_xxx_spriteDeleteStatic_45A4E0_drawable((*nox_drawable)(unsafe.Pointer(uintptr(v331))))
					}
				}
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
		case 52:
			v57 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v58 = v57
			if nox_client_isConnected_43C700() != 0 {
				nox_xxx_cliAddObjFriend_4959F0(v58)
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
		case 53:
			v59 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v60 = v59
			if nox_client_isConnected_43C700() != 0 {
				sub_495A20(v60)
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
		case 54:
			if nox_client_isConnected_43C700() != 0 {
				sub_4959B0()
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 1))
		case 55:
			v62 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v63 = v62
			if nox_client_isConnected_43C700() != 0 {
				if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_NET_DEBUG)) {
					nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
				}
				if nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))) != 0 {
					*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_xxx_netSpriteByCodeStatic_45A720(v63))))
				} else {
					*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_xxx_netSpriteByCodeDynamic_45A6F0(v63))))
				}
				if uint32(int32(v5)) != 0 {
					*(*uint32)(unsafe.Pointer(uintptr(v5 + 120))) |= 0x1000000
				}
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
		case 56:
			v64 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v65 = v64
			if nox_client_isConnected_43C700() != 0 {
				if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_NET_DEBUG)) {
					nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
				}
				if nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))) != 0 {
					*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_xxx_netSpriteByCodeStatic_45A720(v65))))
				} else {
					*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_xxx_netSpriteByCodeDynamic_45A6F0(v65))))
				}
				if uint32(int32(v5)) != 0 {
					if *(*uint32)(unsafe.Pointer(uintptr(v5 + 112)))&0x40000 != 0 {
						*(*uint32)(unsafe.Pointer(uintptr(v5 + 300))) = 0
					}
					*(*uint32)(unsafe.Pointer(uintptr(v5 + 120))) &= 0xFEFFFFFF
				}
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
		case 57:
			v16 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v17 = v16
			if nox_client_isConnected_43C700() == 0 {
				goto LABEL_764
			}
			*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(func() *uint32 {
				if nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))) != 0 {
					return nox_xxx_netSpriteByCodeStatic_45A720(v17)
				}
				return nox_xxx_netSpriteByCodeDynamic_45A6F0(v17)
			}())))
			if uint32(int32(v5)) == 0 {
				goto LABEL_764
			}
			nox_xxx_spriteSetFrameMB_45AB80(int32(v5), int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
		case 58:
			if nox_client_isConnected_43C700() != 0 && !noxflags.HasGame(noxflags.GameHost) {
				nox_xxx_wallDestroyedByWallid_410520(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
		case 59:
			if nox_client_isConnected_43C700() != 0 {
				if !noxflags.HasGame(noxflags.GameHost) {
					*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(sub_410550(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))))
					if uint32(int32(v5)) != 0 {
						if int32(*(*uint8)(unsafe.Pointer(uintptr(v5 + 4))))&4 != 0 {
							v67 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 28))))
							*(*uint8)(unsafe.Pointer(uintptr(v67 + 22))) = 23
							*(*uint8)(unsafe.Pointer(uintptr(v67 + 21))) = 3
						}
					}
				}
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
		case 60:
			if nox_client_isConnected_43C700() != 0 {
				if !noxflags.HasGame(noxflags.GameHost) {
					*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(sub_410550(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))))
					if uint32(int32(v5)) != 0 {
						if int32(*(*uint8)(unsafe.Pointer(uintptr(v5 + 4))))&4 != 0 {
							v68 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 28))))
							if v68 != 0 {
								*(*uint8)(unsafe.Pointer(uintptr(v68 + 22))) = 0
								*(*uint8)(unsafe.Pointer(uintptr(v68 + 21))) = 1
							}
						}
					}
				}
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
		case 61:
			if nox_client_isConnected_43C700() == 0 {
				goto LABEL_916
			}
			*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(nox_server_getWallAtGrid_410580(int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 4)))), int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 5)))))))
			if uint32(int32(v5)) == 0 {
				*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(nox_xxx_wallCreateAt_410250(int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 4)))), int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 5)))))))
				if uint32(int32(v5)) == 0 {
					goto LABEL_916
				}
			}
			v66 = int8(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 6))
			*(*uint8)(unsafe.Pointer(uintptr(v5 + 1))) = uint8(v66)
			*(*uint8)(unsafe.Pointer(uintptr(v5))) = *((*uint8)(unsafe.Add(unsafe.Pointer(data), -4)))
			*(*uint8)(unsafe.Pointer(uintptr(v5 + 2))) = *((*uint8)(unsafe.Add(unsafe.Pointer(data), -3)))
		case 62:
			if nox_client_isConnected_43C700() != 0 {
				*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(nox_server_getWallAtGrid_410580(int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))), int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 2)))))))
				if uint32(int32(v5)) != 0 {
					nox_xxx_mapDelWallAtPt_410430(int32(*(*uint8)(unsafe.Pointer(uintptr(v5 + 5)))), int32(*(*uint8)(unsafe.Pointer(uintptr(v5 + 6)))))
				}
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
		case 65:
			nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			if nox_client_isConnected_43C700() == 0 {
				goto LABEL_764
			}
			if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_NET_DEBUG)) {
				nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			}
			if nox_xxx_unitSpriteCheckAlly_4951F0(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))) == 0 {
				goto LABEL_764
			}
			sub_495150(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))), int16(int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))*2))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
		case 66:
			nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			if nox_client_isConnected_43C700() != 0 && nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_NET_DEBUG)) {
				nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			}
			nox_xxx_cliAddHealthChange_49A650(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))), int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 5))
		case 67:
			if nox_client_isConnected_43C700() != 0 {
				sub_470CB0(int32(*(*int16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
		case 68:
			v92 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v93 = v92
			if nox_client_isConnected_43C700() == 0 {
				goto LABEL_1149
			}
			if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_NET_DEBUG)) {
				nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			}
			sub_4675E0(v93, int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))), int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 5))))))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 7))
		case 69:
			v96 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v97 = v96
			if nox_client_isConnected_43C700() == 0 {
				goto LABEL_1163
			}
			if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_NET_DEBUG)) {
				nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			}
			if uint32(v97) != nox_player_netCode_85319C {
				goto LABEL_1163
			}
			nox_xxx_cliSetMana_470D10(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 5))
		case 71:
			if nox_client_isConnected_43C700() == 0 {
				goto LABEL_966
			}
			sub_470D20(int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))), *memmap.PtrInt32(0x587000, 157092))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 2))
		case 72:
			v98 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v99 = int32(*memmap.PtrUint32(0x8531A0, 2576))
			v100 = v98
			if nox_client_isConnected_43C700() != 0 {
				if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_NET_DEBUG)) {
					nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
				}
				if uint32(v100) == nox_player_netCode_85319C {
					v101 = bool2int(noxflags.HasGame(noxflags.GameHost))
					if v101 == 0 && v99 != 0 {
						*(*uint32)(unsafe.Pointer(uintptr(v99 + 2247))) = uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))))
						*(*uint32)(unsafe.Pointer(uintptr(v99 + 2243))) = uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 5)))))
						*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v101))), unsafe.Sizeof(uint16(0))*0)) = *(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 9))))
						v332 = float32(float64(v101))
						*(*uint32)(unsafe.Pointer(uintptr(v99 + 2235))) = uint32(int(v332))
						*(*uint32)(unsafe.Pointer(uintptr(v99 + 2239))) = uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 11)))))
						*(*uint16)(unsafe.Pointer(uintptr(v99 + 3652))) = *(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 7))))
						*(*uint8)(unsafe.Pointer(uintptr(v99 + 3684))) = *((*uint8)(unsafe.Add(unsafe.Pointer(data), 13)))
					}
					nox_xxx_j_inventoryNameSignInit_467460()
				}
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 14))
		case 73:
			if nox_client_isConnected_43C700() == 0 {
				goto LABEL_1163
			}
			sub_467450(int32(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 5))
		case 74:
			v102 = int32(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))
			if nox_client_isConnected_43C700() == 0 {
				goto LABEL_1163
			}
			sub_467490(v102)
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 5))
		case 75:
			v103 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v104 = v103
			if nox_client_isConnected_43C700() == 0 {
				goto LABEL_1163
			}
			if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_NET_DEBUG)) {
				nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			}
			*(*uint32)(unsafe.Pointer(&v381[0])) = 0
			*(*uint16)(unsafe.Pointer(&v381[16])) = math.MaxUint16
			*(*uint32)(unsafe.Pointer(&v381[4])) = 0
			*(*uint16)(unsafe.Pointer(&v381[18])) = math.MaxUint16
			*(*uint32)(unsafe.Pointer(&v381[8])) = 0
			*(*uint32)(unsafe.Pointer(&v381[12])) = 0
			if nox_xxx_spritePickup_461660(v104, int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))), unsafe.Pointer(&v381[0])) != 0 {
				goto LABEL_1163
			}
			nox_xxx_send2ServInvenFail_461630(int16(v104))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 5))
		case 76:
			v105 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v106 = v105
			if nox_client_isConnected_43C700() == 0 {
				goto LABEL_735
			}
			if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_NET_DEBUG)) {
				nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			}
			v107 = 0
			v108 = &v385[0]
			for {
				*(*uint32)(unsafe.Pointer(v108)) = uint32(uintptr(unsafe.Pointer(nox_xxx_modifGetDescById_413330(int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), func() int32 {
					p := &v107
					x := *p
					*p++
					return x
				}()))), 5))))))))
				v108 = (*byte)(unsafe.Add(unsafe.Pointer(v108), 4))
				if v107 >= 4 {
					break
				}
			}
			*(*uint16)(unsafe.Pointer(&v385[16])) = math.MaxUint16
			*(*uint16)(unsafe.Pointer(&v385[18])) = math.MaxUint16
			if nox_xxx_spritePickup_461660(v106, int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))), unsafe.Pointer(&v385[0])) != 0 {
				goto LABEL_735
			}
			nox_xxx_send2ServInvenFail_461630(int16(v106))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 9))
		case 77:
			v109 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v110 = v109
			if nox_client_isConnected_43C700() == 0 {
				goto LABEL_1163
			}
			if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_NET_DEBUG)) {
				nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			}
			sub_461A80(v110)
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 5))
		case 78:
			v124 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v125 = v124
			v126 = nox_xxx_cliGamedataGet_416590(0)
			if nox_client_isConnected_43C700() != 0 {
				v127 = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(v125)))
				if v127 != nil {
					if !noxflags.HasGame(noxflags.GameHost) {
						*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v127))), unsafe.Sizeof(uint32(0))*534))) = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))
						*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v127))), unsafe.Sizeof(uint32(0))*535))) = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 7))))
						*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v127))), unsafe.Sizeof(uint32(0))*536))) = nox_frame_xxx_2598000
					}
					if noxflags.HasGame(noxflags.GameModeElimination) {
						k = int32(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 7)))))
						if k >= int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v126))), unsafe.Sizeof(uint16(0))*27)))) {
							v128 = strMan.GetStringInFile("Eliminated", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
							nox_swprintf(&v395[0], v128, (*byte)(unsafe.Add(unsafe.Pointer(v127), 4704)))
							clientPlaySoundSpecial(312, 100)
							nox_xxx_printCentered_445490(&v395[0])
						}
					}
				}
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 11))
		case 79:
			fallthrough
		case 80:
			if nox_client_isConnected_43C700() == 0 {
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 7))
				break
			}
			*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint16(0))*0)) = *(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))
			v367 = -1
			*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))) = uint16(int16(v5 & math.MaxInt16))
			if ((uint32(int32(v5)) >> 15) & 1) == 1 {
				nox_xxx_clientEquipWeaponArmor_417AA0(int8(*data), int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))), int32(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))), int32(uintptr(unsafe.Pointer(&v367))))
			} else {
				nox_xxx_clientEquip_49A3D0(int8(*data), int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))), int32(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))), int32(uintptr(unsafe.Pointer(&v367))))
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 7))
		case 81:
			fallthrough
		case 82:
			*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(nox_client_isConnected_43C700())
			if uint32(int32(v5)) == 0 {
				goto LABEL_437
			}
			*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint16(0))*0)) = *(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))
			v116 = int32(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))))
			v117 = uint32(int32(v5))
			v118 = uint16(int16(v5 & math.MaxInt16))
			*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))) = v118
			v333 = int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 7)))))
			if ((v117 >> 15) & 1) == 1 {
				nox_xxx_clientEquipWeaponArmor_417AA0(int8(*data), int32(v118), v116, v333)
			} else {
				nox_xxx_clientEquip_49A3D0(int8(*data), int32(v118), v116, v333)
			}
		LABEL_437:
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 11))
		case 83:
			fallthrough
		case 84:
			if nox_client_isConnected_43C700() == 0 {
				goto LABEL_1149
			}
			sub_417B80(int8(*data), int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))), int32(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 7))
		case 85:
			v140 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v141 = v140
			if nox_client_isConnected_43C700() != 0 {
				if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_NET_DEBUG)) {
					nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
				}
				v142 = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(v141)))
				if v142 != nil {
					if !noxflags.HasGame(noxflags.GameHost) && nox_frame_xxx_2598000 > *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v142))), unsafe.Sizeof(uint32(0))*540))) {
						*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint16(0))*1)) = 0
						*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v142))), unsafe.Sizeof(uint32(0))*538))) = uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))))
						*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint16(0))*2)) = *(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 5))))
						*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v142))), unsafe.Sizeof(uint32(0))*539))) = uint32(*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint16(0))*2)))
						*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v142))), unsafe.Sizeof(uint32(0))*540))) = nox_frame_xxx_2598000
					}
					k = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v142))), unsafe.Sizeof(uint32(0))*539))) - 1)
					if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v142))), unsafe.Sizeof(uint32(0))*538))) == uint32(k) {
						v143 = strMan.GetStringInFile("SH_NearVictory", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
						nox_swprintf(&v406[0], v143, (*byte)(unsafe.Add(unsafe.Pointer(v142), 4704)))
						nox_xxx_printCentered_445490(&v406[0])
					}
				}
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 7))
		case 86:
			v144 = (*byte)(unsafe.Pointer(nox_xxx_clientGetTeamColor_418AB0(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))))
			if nox_client_isConnected_43C700() == 0 {
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 8))
				break
			}
			*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*1)) = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 4))))
			if (*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*1))) <= uint32(*(*int32)(unsafe.Pointer(&dword_5d4594_1200804))) {
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 8))
				break
			}
			if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_NET_DEBUG)) {
				nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			}
			clientPlaySoundSpecial(309, 100)
			if !noxflags.HasGame(noxflags.GameHost) {
				noxflags.SetGame(noxflags.GameFlag4)
			}
			nox_xxx_initTime_435570()
			if !nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING)) {
				sub_470510()
			}
			v145 = nox_xxx_objGetTeamByNetCode_418C80(int32(nox_player_netCode_85319C))
			if v145 != nil && nox_xxx_teamCompare2_419180(int32(uintptr(unsafe.Pointer(v145))), *((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))) != 0 {
				v146 = strMan.GetStringInFile("TeamWon", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
				nox_swprintf(&v400[0], v146)
				v147 = 0
			} else {
				v148 = strMan.GetStringInFile("teamformat", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
				nox_swprintf(&v400[0], v148, v144)
				v149 = strMan.GetStringInFile("FB_Victory", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
				nox_swprintf(&v400[0], v149, &v400[0])
				v147 = 1
			}
			sub_435700(&v400[0], v147)
			nox_xxx_guiServerOptionsHide_4597E0(0)
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 8))
		case 87:
			v172 = (*byte)(unsafe.Pointer(nox_xxx_clientGetTeamColor_418AB0(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))))
			if nox_client_isConnected_43C700() == 0 {
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 8))
				break
			}
			k = int32(dword_5d4594_1200804)
			if *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 4)))) <= uint32(*(*int32)(unsafe.Pointer(&dword_5d4594_1200804))) {
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 8))
				break
			}
			v393[0] = 0
			if int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))) == 1 {
				v173 = strMan.GetStringInFile("TimeLimitReached", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
				nox_swprintf(&v393[0], v173)
			}
			clientPlaySoundSpecial(309, 100)
			if !noxflags.HasGame(noxflags.GameHost) {
				noxflags.SetGame(noxflags.GameFlag4)
			}
			nox_xxx_initTime_435570()
			if !nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING)) {
				sub_470510()
			}
			v174 = nox_xxx_objGetTeamByNetCode_418C80(int32(nox_player_netCode_85319C))
			if v172 != nil {
				v175 = strMan.GetStringInFile("CTF_Victory", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
				nox_swprintf(&v398[0], v175, v172)
				if v174 == nil || nox_xxx_teamCompare2_419180(int32(uintptr(unsafe.Pointer(v174))), uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v172), 57)))) == 0 {
					v176 = 1
					goto LABEL_600
				}
			} else {
				v177 = strMan.GetStringInFile("CTF_Tie", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
				nox_swprintf(&v398[0], v177)
			}
			v176 = 0
		LABEL_600:
			nox_wcscat(&v393[0], &v398[0])
			sub_435700(&v393[0], v176)
			nox_xxx_guiServerOptionsHide_4597E0(0)
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 8))
		case 88:
			v150 = 1
			v389[0] = 0
			v392[0] = 0
			v391[0] = 0
			if nox_client_isConnected_43C700() == 0 {
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 8))
				break
			}
			*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*1)) = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 4))))
			if (*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*1))) <= uint32(*(*int32)(unsafe.Pointer(&dword_5d4594_1200804))) {
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 8))
				break
			}
			v391[0] = 0
			if int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))) == 1 {
				v151 = strMan.GetStringInFile("TimeLimitReached", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
				nox_swprintf(&v391[0], v151)
			}
			clientPlaySoundSpecial(309, 100)
			if !noxflags.HasGame(noxflags.GameHost) {
				noxflags.SetGame(noxflags.GameFlag4)
			}
			if !nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING)) {
				sub_470510()
			}
			nox_xxx_initTime_435570()
			if int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))) != 0 {
				v152 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
				v372 = v152
			}
			v153 = v372
			v154 = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(v372)))
			if !noxflags.HasGame(noxflags.GameModeElimination) {
				if int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))) != 0 {
					if v154 == nil {
						goto LABEL_559
					}
					if int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))) == 0 {
						sub_4947E0(int32(uintptr(unsafe.Pointer(v154))))
					}
					if uint32(v153) != nox_player_netCode_85319C {
						v159 = strMan.GetStringInFile("DM_Loss", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
						nox_swprintf(&v389[0], v159, (*byte)(unsafe.Add(unsafe.Pointer(v154), 4704)))
						v150 = 1
						goto LABEL_559
					}
					if *(*byte)(unsafe.Add(unsafe.Pointer(v154), 2252)) == 0 {
						v314 = strMan.GetStringInFile("DM_MaleVictory", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
						nox_swprintf(&v389[0], v314)
						goto LABEL_558
					}
					v155 = strMan.GetStringInFile("DM_FemaleVictory", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
				} else {
					v155 = strMan.GetStringInFile("DM_Tie", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
				}
			LABEL_557:
				nox_swprintf(&v389[0], v155)
				goto LABEL_558
			}
			if int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))) == 0 {
				v155 = strMan.GetStringInFile("HL_Tie", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
				goto LABEL_557
			}
			v156 = strMan.GetStringInFile("HL_Header", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
			nox_swprintf(&v389[0], v156)
			if uint32(v153) == nox_player_netCode_85319C {
				v334 = strMan.GetStringInFile("HL_You", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
				v157 = strMan.GetStringInFile("HL_Victory", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
				nox_swprintf(&v392[0], v157, v334)
				nox_wcscat(&v389[0], &v392[0])
			LABEL_558:
				v150 = 0
			} else {
				if v154 != nil {
					v158 = strMan.GetStringInFile("HL_Victory", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
					nox_swprintf(&v392[0], v158, (*byte)(unsafe.Add(unsafe.Pointer(v154), 4704)))
					if int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))) == 0 {
						sub_4947E0(int32(uintptr(unsafe.Pointer(v154))))
					}
				}
				nox_wcscat(&v389[0], &v392[0])
				v150 = 1
			}
		LABEL_559:
			nox_wcscat(&v391[0], &v389[0])
			sub_435700(&v391[0], v150)
			nox_xxx_guiServerOptionsHide_4597E0(0)
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 8))
		case 89:
			v160 = (*byte)(unsafe.Pointer(nox_xxx_clientGetTeamColor_418AB0(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))))
			if nox_client_isConnected_43C700() == 0 {
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 8))
				break
			}
			k = int32(dword_5d4594_1200804)
			if *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 4)))) <= uint32(*(*int32)(unsafe.Pointer(&dword_5d4594_1200804))) {
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 8))
				break
			}
			v394[0] = 0
			if int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))) == 1 {
				v161 = strMan.GetStringInFile("TimeLimitReached", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
				nox_swprintf(&v394[0], v161)
			}
			clientPlaySoundSpecial(309, 100)
			if !noxflags.HasGame(noxflags.GameHost) {
				noxflags.SetGame(noxflags.GameFlag4)
			}
			nox_xxx_initTime_435570()
			if !nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING)) {
				sub_470510()
			}
			v162 = nox_xxx_objGetTeamByNetCode_418C80(int32(nox_player_netCode_85319C))
			if noxflags.HasGame(noxflags.GameModeElimination) {
				if v160 == nil {
					v163 = strMan.GetStringInFile("HL_Tie", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
					nox_swprintf(&v390[0], v163)
					v164 = 0
					goto LABEL_585
				}
				v165 = strMan.GetStringInFile("HL_Header", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
				nox_swprintf(&v390[0], v165)
				if v162 != nil && nox_xxx_teamCompare2_419180(int32(uintptr(unsafe.Pointer(v162))), uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v160), 57)))) != 0 {
					v335 = strMan.GetStringInFile("HL_YourTeam", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
					v166 = strMan.GetStringInFile("HL_Victory", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
					nox_swprintf(&v388[0], v166, v335)
					nox_wcscat(&v390[0], &v388[0])
					v164 = 0
				} else {
					v336 = strMan.GetStringInFile("Team", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
					v167 = strMan.GetStringInFile("HL_Victory", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
					nox_swprintf(&v388[0], v167, v336)
					nox_swprintf(&v396[0], &v388[0], v160)
					nox_wcscat(&v390[0], &v396[0])
					v164 = 1
				}
				if int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))) == 0 {
					goto LABEL_584
				}
			} else {
				if v160 == nil {
					v168 = strMan.GetStringInFile("DM_Tie", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
					nox_swprintf(&v390[0], v168)
					v164 = 0
					goto LABEL_585
				}
				if v162 != nil && nox_xxx_teamCompare2_419180(int32(uintptr(unsafe.Pointer(v162))), uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v160), 57)))) != 0 {
					v169 = strMan.GetStringInFile("DM_TeamVictory", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
					nox_swprintf(&v390[0], v169)
					v164 = 0
				} else {
					v170 = strMan.GetStringInFile("Team", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
					nox_swprintf(&v388[0], v170, v160)
					v171 = strMan.GetStringInFile("DM_Loss", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
					nox_swprintf(&v390[0], v171, &v388[0])
					v164 = 1
				}
				if int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))) == 0 {
				LABEL_584:
					sub_4948B0(int32(uintptr(unsafe.Pointer(v160))))
					goto LABEL_585
				}
			}
		LABEL_585:
			nox_wcscat(&v394[0], &v390[0])
			sub_435700(&v394[0], v164)
			nox_xxx_guiServerOptionsHide_4597E0(0)
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 8))
		case 90:
			v132 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v133 = v132
			if nox_client_isConnected_43C700() == 0 {
				goto LABEL_1149
			}
			if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_NET_DEBUG)) {
				nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			}
			*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(func() *uint32 {
				if nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))) != 0 {
					return nox_xxx_netSpriteByCodeStatic_45A720(v133)
				}
				return nox_xxx_netSpriteByCodeDynamic_45A6F0(v133)
			}())))
			v134 = int32(v5)
			if uint32(int32(v5)) == 0 {
				goto LABEL_1149
			}
			*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(bool2int(nox_xxx_spriteCheckFlag31_4356C0((*nox_drawable)(unsafe.Pointer(uintptr(v5))), 15)))
			*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*1)) = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))
			v135 = int32(v5)
			*(*uint32)(unsafe.Pointer(uintptr(v134 + 124))) = *(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*1))
			if uint32(v134) == *memmap.PtrUint32(0x852978, 8) {
				sub_467410(int32(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))))
			}
			if v135 != 1 || nox_xxx_spriteCheckFlag31_4356C0((*nox_drawable)(unsafe.Pointer(uintptr(v134))), 15) || uint32(v134) == *memmap.PtrUint32(0x852978, 8) && int32(sub_467430())&8 != 0 {
				goto LABEL_1149
			}
			var v136 *nox_thing = nox_get_thing(int32(*(*uint32)(unsafe.Pointer(uintptr(v134 + 108)))))
			nox_xxx_spriteChangeIntensity_484D70_light_intensity(v134+136, v136.light_intensity)
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 7))
		case 91:
			if nox_client_isConnected_43C700() != 0 {
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v129))), 0)) = sub_467430()
				v130 = int32((v129 >> 3) & 1)
				sub_467420(int8(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))
				if v130 == 1 && (int32(sub_467430())&8) == 0 && *memmap.PtrUint32(0x852978, 8) != 0 && !nox_xxx_spriteCheckFlag31_4356C0((*nox_drawable)(unsafe.Pointer(uintptr(*memmap.PtrInt32(0x852978, 8)))), 15) {
					var v131 *nox_thing = nox_get_thing(int32(*(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x852978, 8) + 108)))))
					nox_xxx_spriteChangeIntensity_484D70_light_intensity(int32(*memmap.PtrUint32(0x852978, 8)+136), v131.light_intensity)
				}
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 2))
		case 92:
			v88 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v89 = v88
			if nox_client_isConnected_43C700() == 0 {
				goto LABEL_902
			}
			if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_NET_DEBUG)) {
				nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			}
			*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(func() *uint32 {
				if nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))) != 0 {
					return nox_xxx_netSpriteByCodeStatic_45A720(v89)
				}
				return nox_xxx_netSpriteByCodeDynamic_45A6F0(v89)
			}())))
			if uint32(int32(v5)) != 0 {
				nox_xxx_spriteChangeLightColor_484BE0((*uint32)(unsafe.Pointer(uintptr(v5+136))), int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))), int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 4)))), int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 5)))))
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 6))
			} else {
			LABEL_902:
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 6))
			}
		case 93:
			v84 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v85 = v84
			if nox_client_isConnected_43C700() == 0 {
				goto LABEL_948
			}
			if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_NET_DEBUG)) {
				nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			}
			*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(func() *uint32 {
				if nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))) != 0 {
					return nox_xxx_netSpriteByCodeStatic_45A720(v85)
				}
				return nox_xxx_netSpriteByCodeDynamic_45A6F0(v85)
			}())))
			if uint32(int32(v5)) == 0 {
				goto LABEL_948
			}
			nox_xxx_spriteChangeIntensity_484D70_light_intensity(int32(v5+136), *(*float32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 7))
		case 94:
			v178 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v179 = v178
			if nox_client_isConnected_43C700() != 0 {
				*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(func() *uint32 {
					if nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))) != 0 {
						return nox_xxx_netSpriteByCodeStatic_45A720(v179)
					}
					return nox_xxx_netSpriteByCodeDynamic_45A6F0(v179)
				}())))
				if uint32(int32(v5)) != 0 {
					*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint16(0))*2)) = uint16(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))
					*(*uint16)(unsafe.Pointer(uintptr(v5 + 104))) = *(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint16(0))*2))
				}
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
		case 95:
			v180 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v181 = v180
			if nox_client_isConnected_43C700() != 0 {
				*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(func() *uint32 {
					if nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))) != 0 {
						return nox_xxx_netSpriteByCodeStatic_45A720(v181)
					}
					return nox_xxx_netSpriteByCodeDynamic_45A6F0(v181)
				}())))
				if uint32(int32(v5)) != 0 {
					*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint16(0))*2)) = uint16(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))
					*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*1)) = -(*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*1)))
					*(*uint16)(unsafe.Pointer(uintptr(v5 + 104))) = *(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint16(0))*2))
				}
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
		case 96:
			v120 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v121 = v120
			if nox_client_isConnected_43C700() != 0 {
				if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_NET_DEBUG)) {
					nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
				}
				sub_462040(v121)
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
		case 97:
			v122 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v123 = v122
			if nox_client_isConnected_43C700() != 0 {
				if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_NET_DEBUG)) {
					nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
				}
				sub_4624D0(v123)
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
		case 100:
			v182 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v183 = v182
			if nox_client_isConnected_43C700() == 0 {
				goto LABEL_638
			}
			sub_467930(v183, int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))), int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 4)))))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 5))
		case 101:
			v71 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v72 = v71
			if nox_client_isConnected_43C700() != 0 {
				if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_NET_DEBUG)) {
					nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
				}
				if nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))) != 0 {
					*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_xxx_netSpriteByCodeStatic_45A720(v72))))
				} else {
					*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_xxx_netSpriteByCodeDynamic_45A6F0(v72))))
				}
				v73 = (*uint32)(unsafe.Pointer(uintptr(v5)))
				if uint32(int32(v5)) != 0 {
					v74 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 280))))
					*(*uint32)(unsafe.Add(unsafe.Pointer(v73), unsafe.Sizeof(uint32(0))*70)) = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))
					k = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v73), unsafe.Sizeof(uint32(0))*28)))
					if uint32(k)&0x20000 != 0 {
						if (v74 & 1024) == 0 {
							*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*1)) = *(*uint32)(unsafe.Add(unsafe.Pointer(v73), unsafe.Sizeof(uint32(0))*70)) & 1024
							if (*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*1))) == 1024 {
								sub_4BC720(int32(uintptr(unsafe.Pointer(v73))))
							}
						}
						v75 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v73), unsafe.Sizeof(uint32(0))*70)))
						if v75&2048 != 0 {
							k = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v73), unsafe.Sizeof(uint32(0))*28)) & 0xFFF7FFFF)
							v76 = *(*uint32)(unsafe.Add(unsafe.Pointer(v73), unsafe.Sizeof(uint32(0))*30)) & 0xDFFFFFFF
							*(*uint32)(unsafe.Add(unsafe.Pointer(v73), unsafe.Sizeof(uint32(0))*28)) = uint32(k)
							*(*uint32)(unsafe.Add(unsafe.Pointer(v73), unsafe.Sizeof(uint32(0))*30)) = v76
						}
					}
				}
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 7))
		case 102:
			if nox_client_isConnected_43C700() != 0 && *memmap.PtrUint32(0x852978, 8) != 0 {
				*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*1)) = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))
				*(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x852978, 8) + 120))) = *(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*1))
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 5))
		case 103:
			v77 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v78 = v77
			if nox_client_isConnected_43C700() != 0 {
				if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_NET_DEBUG)) {
					nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
				}
				if nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))) != 0 {
					*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_xxx_netSpriteByCodeStatic_45A720(v78))))
				} else {
					*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_xxx_netSpriteByCodeDynamic_45A6F0(v78))))
				}
				if uint32(int32(v5)) != 0 {
					v79 = int32(v5 + 432)
					if uint32(int32(v5)) != 0xFFFFFE50 {
						v80 = 0
						v81 = (*int32)(unsafe.Pointer(uintptr(v5 + 432)))
						for {
							*v81 = int32(uintptr(unsafe.Pointer(nox_xxx_modifGetDescById_413330(int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), func() int32 {
								p := &v80
								x := *p
								*p++
								return x
							}()))), 3))))))))
							v81 = (*int32)(unsafe.Add(unsafe.Pointer(v81), unsafe.Sizeof(int32(0))*1))
							if v80 >= 4 {
								break
							}
						}
						*(*uint16)(unsafe.Pointer(uintptr(v79 + 16))) = math.MaxUint16
						*(*uint16)(unsafe.Pointer(uintptr(v79 + 18))) = math.MaxUint16
					}
				}
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 7))
		case 104:
			v82 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v83 = v82
			if nox_client_isConnected_43C700() == 0 {
				goto LABEL_1057
			}
			if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_NET_DEBUG)) {
				nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			}
			if uint32(v83) == nox_player_netCode_85319C {
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&k))), 0)) = *((*uint8)(unsafe.Add(unsafe.Pointer(data), 7)))
				sub_467470(k, *(*float32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))))
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 8))
			} else {
			LABEL_1057:
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 8))
			}
		case 105:
			if nox_client_isConnected_43C700() != 0 {
				v285 = *(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))
				v286 = v285
				v285 &= math.MaxInt16
				*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))) = v285
				v287 = int32(v286) >> 15
				v288 = (*byte)(unsafe.Pointer(nox_npc_by_id(int32(v285))))
				if v288 != nil {
					*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(nox_init_npc((*nox_npc)(unsafe.Pointer(v288)), int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))))
				} else {
					*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_new_npc(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))))))
					v288 = (*byte)(unsafe.Pointer(uintptr(v5)))
				}
				if v288 != nil {
					k = int32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v288), 8)))))
					v289 = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
					v358 = (*int32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v288), 8))))
					v360 = 6
					for {
						*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v5))), 4)) = *(*uint8)(unsafe.Add(unsafe.Pointer(v289), 1))
						*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v5))), 0)) = *v289
						*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&k))), 0)) = *((*uint8)(unsafe.Add(unsafe.Pointer(v289), -1)))
						v289 = (*uint8)(unsafe.Add(unsafe.Pointer(v289), 3))
						*v358 = int32(nox_color_rgb_4344A0(k, int32(v5), *(*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(&v5))), unsafe.Sizeof(int32(0))*1))))
						k = int32(uintptr(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer(v358), unsafe.Sizeof(int32(0))*1)))))
						*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(v360 - 1)
						v250 = v360 == 1
						v358 = (*int32)(unsafe.Add(unsafe.Pointer(v358), unsafe.Sizeof(int32(0))*1))
						v360--
						if v250 {
							break
						}
					}
					*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v288))), unsafe.Sizeof(uint32(0))*328))) = uint32(v287)
				}
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 21))
		case 106:
			v119 = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))))
			if v119 == nil {
				goto LABEL_1070
			}
			if !noxflags.HasGame(noxflags.GameHost) {
				nox_xxx_playerUnsetStatus_417530((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v119)))))), 1059)
				nox_xxx_netNeedTimestampStatus_4174F0(int32(uintptr(unsafe.Pointer(v119))), int32(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))&1059))
			}
			if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING)) || (func() bool {
				*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*1)) = uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))
				return uint32(*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint16(0))*2))) != nox_player_netCode_85319C
			}()) {
			LABEL_1070:
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 7))
			} else {
				if *(*byte)(unsafe.Add(unsafe.Pointer(v119), 3680))&1 != 0 {
					nox_client_renderGUI_80828 = 0
				} else if nox_xxx_xxxRenderGUI_587000_80832 == 1 {
					nox_client_renderGUI_80828 = 1
				}
				sub_470C40(int32((*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v119))), unsafe.Sizeof(uint32(0))*920))) >> 10) & 1))
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 7))
			}
		case 107:
			v69 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v70 = v69
			if nox_client_isConnected_43C700() == 0 {
				goto LABEL_948
			}
			if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_NET_DEBUG)) {
				nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			}
			*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(func() *uint32 {
				if nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))) != 0 {
					return nox_xxx_netSpriteByCodeStatic_45A720(v70)
				}
				return nox_xxx_netSpriteByCodeDynamic_45A6F0(v70)
			}())))
			if uint32(int32(v5)) == 0 {
				goto LABEL_948
			}
			nox_xxx_spriteSetFrameMB_45AB80(int32(v5), int32(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 7))
		case 108:
			v184 = uint32(nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))))
			v185 = v184
			if nox_client_isConnected_43C700() == 0 {
				goto LABEL_638
			}
			v186 = *(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))
			v187 = v186
			v186 &= math.MaxInt16
			*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))) = v186
			nox_xxx_cliSummonCreat_4C2E50(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))), int32(v186), int32(v187)&0x8000)
			if nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))) != 0 {
				v188 = nox_xxx_netSpriteByCodeStatic_45A720(int32(v185))
			} else {
				v188 = nox_xxx_netSpriteByCodeDynamic_45A6F0(int32(v185))
			}
			if v188 != nil || (func() *uint32 {
				v188 = nox_xxx_spriteCreate_48E970(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))), v185, 0, 0)
				return v188
			}()) != nil {
				sub_459DD0((*nox_drawable)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v188)))))), 1)
			}
			sub_495060(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))), 0, 0)
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 5))
		case 109:
			v189 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v190 = v189
			if nox_client_isConnected_43C700() == 0 {
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
				break
			}
			v191 = *(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))
			v192 = v191
			v191 &= math.MaxInt16
			*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))) = v191
			nox_xxx_cliSummonOnDieOrBanish_4C3140(int32(v191), unsafe.Pointer(uintptr(int32(v192)&0x8000)))
			sub_4950C0(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v337 = v190
			if nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))) == 0 {
				goto LABEL_642
			}
			*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_xxx_netSpriteByCodeStatic_45A720(v190))))
			goto LABEL_643
		case 110:
			if nox_client_isConnected_43C700() == 0 {
				goto LABEL_1163
			}
			sub_467440(int32(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 5))
		case 111:
			v139 = 0
			if nox_client_isConnected_43C700() != 0 {
				if int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))&128 != 0 {
					v139 = 1
				}
				nox_xxx_netSpellRewardCli_45CFE0(int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))), int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 2)))), int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))&math.MaxInt8, v139)
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
		case 112:
			sub_49BB80(int8(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 2))
		case 113:
			sub_467CA0()
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 1))
		case 124:
			v236 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 8))))))
			v237 = v236
			if nox_client_isConnected_43C700() != 0 {
				if nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 8)))))) != 0 {
					v238 = nox_xxx_netSpriteByCodeStatic_45A720(v237)
				} else {
					v238 = nox_xxx_netSpriteByCodeDynamic_45A6F0(v237)
				}
				*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&k))), unsafe.Sizeof(uint16(0))*1)) = 0
				v382[0] = int32(uintptr(unsafe.Pointer(v238)))
				*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&k))), unsafe.Sizeof(uint16(0))*0)) = *(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2))))
				*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint16(0))*1)) = 0
				v382[5] = int32(uint16(int16(k)))
				*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint16(0))*2)) = *(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 4))))
				v382[6] = int32(*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint16(0))*2)))
				v382[7] = int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 6)))))
				switch *((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))) {
				case 0:
					nox_xxx_ParticleFxT0_4AEEA0((**uint8)(unsafe.Pointer(&v382[0])))
				case 1:
					nox_xxx_ParticleFxT1_4AEF80(&v382[0])
				case 2:
					nox_xxx_ParticleFxT2_4AF0F0((**uint8)(unsafe.Pointer(&v382[0])))
				case 3:
					nox_xxx_ParticleFxT3_4AF2A0((**uint8)(unsafe.Pointer(&v382[0])))
				case 4:
					nox_xxx_ParticleFxT4_4AF3D0((*uint32)(unsafe.Pointer(&v382[0])))
				case 5:
					nox_xxx_ParticleFxT5_4AF450(&v382[0])
				case 6:
					nox_xxx_ParticleFxT6_4AF5A0(&v382[0])
				}
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 14))
		case 125:
			fallthrough
		case 140:
			fallthrough
		case 141:
			fallthrough
		case 142:
			if nox_client_isConnected_43C700() == 0 {
				goto LABEL_732
			}
			nox_xxx_netDrawRays_49BDD0(data)
			*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&k))), unsafe.Sizeof(uint16(0))*1)) = 0
			v379.field_0 = int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))
			v379.field_4 = int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))))
			*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&k))), unsafe.Sizeof(uint16(0))*0)) = *(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 5))))
			*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint16(0))*1)) = 0
			v377.field_0 = int32(uint16(int16(k)))
			*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint16(0))*2)) = *(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 7))))
			v377.field_4 = int32(*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint16(0))*2)))
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v5))), 0)) = *data
			if int32(*data) == 140 || int32(uint8(int8(v5))) == 142 {
				nox_xxx_makeLightningParticles_4999D0(*(*int32)(unsafe.Pointer(&dword_5d4594_1200776)), &v379, &v377)
			} else if int32(uint8(int8(v5))) == 125 {
				nox_xxx_drawEnergyBolt_499710(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 5))))), int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 7))))), 10, *(*int32)(unsafe.Pointer(&dword_5d4594_1200776)))
			}
		LABEL_732:
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 9))
		case 126:
			if nox_client_isConnected_43C700() != 0 {
				nox_xxx_netHandleSummonPacket_4B7C40(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 5))))), (*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))), *(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 7)))), *((*uint8)(unsafe.Add(unsafe.Pointer(data), 9))), int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 10))))))
			}
		LABEL_1130:
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 12))
		case math.MaxInt8:
			if nox_client_isConnected_43C700() != 0 {
				sub_4B7EE0(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
		case 128:
			if nox_client_isConnected_43C700() != 0 {
				nox_xxx_fxShield_4B8090(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))), int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))))
			}
			goto LABEL_764
		case 129:
			if nox_client_isConnected_43C700() == 0 {
				goto LABEL_797
			}
			nox_xxx_makePointFxCli_499610(*(*int32)(unsafe.Pointer(&dword_5d4594_1200776)), 50, 1000, 30, int32(*(*int16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))), int32(*(*int16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 5))
		case 130:
			if nox_client_isConnected_43C700() == 0 {
				goto LABEL_797
			}
			nox_xxx_makePointFxCli_499610(*memmap.PtrInt32(6112660, 1200780), 25, 500, 25, int32(*(*int16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))), int32(*(*int16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 5))
		case 131:
			if nox_client_isConnected_43C700() == 0 {
				goto LABEL_797
			}
			nox_xxx_makePointFxCli_499610(*memmap.PtrInt32(6112660, 1200784), 25, 500, 25, int32(*(*int16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))), int32(*(*int16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 5))
		case 132:
			if nox_client_isConnected_43C700() != 0 {
				nox_xxx_makePointFxCli_499610(*(*int32)(unsafe.Pointer(&dword_5d4594_1200796)), 25, 500, 25, int32(*(*int16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))), int32(*(*int16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))))
			}
		LABEL_797:
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 5))
		case 133:
			fallthrough
		case 134:
			fallthrough
		case 135:
			fallthrough
		case 136:
			fallthrough
		case 137:
			fallthrough
		case 138:
			fallthrough
		case 139:
			fallthrough
		case 163:
			if nox_client_isConnected_43C700() != 0 {
				k = 0
				switch op {
				case 133:
					v257 = int32(*memmap.PtrUint32(6112660, 1200872))
					if *memmap.PtrUint32(6112660, 1200872) == 0 {
						v257 = nox_xxx_getTTByNameSpriteMB_44CFC0("FireBoom")
						*memmap.PtrUint32(6112660, 1200872) = uint32(v257)
					}
					v258 = int32(*(*int16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))))
					goto LABEL_858
				case 134:
					v257 = int32(*memmap.PtrUint32(6112660, 1200876))
					if *memmap.PtrUint32(6112660, 1200876) == 0 {
						v257 = nox_xxx_getTTByNameSpriteMB_44CFC0("MediumFireBoom")
						*memmap.PtrUint32(6112660, 1200876) = uint32(v257)
					}
					v258 = int32(*(*int16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))))
					goto LABEL_858
				case 135:
					v257 = int32(*memmap.PtrUint32(6112660, 1200880))
					if *memmap.PtrUint32(6112660, 1200880) == 0 {
						v257 = nox_xxx_getTTByNameSpriteMB_44CFC0("CounterspellBoom")
						*memmap.PtrUint32(6112660, 1200880) = uint32(v257)
					}
					v258 = int32(*(*int16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))))
					goto LABEL_858
				case 136:
					v257 = int32(*memmap.PtrUint32(6112660, 1200884))
					if *memmap.PtrUint32(6112660, 1200884) == 0 {
						v257 = nox_xxx_getTTByNameSpriteMB_44CFC0("ThinFireBoom")
						*memmap.PtrUint32(6112660, 1200884) = uint32(v257)
					}
					v258 = int32(*(*int16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))))
					goto LABEL_858
				case 137:
					v257 = int32(*memmap.PtrUint32(6112660, 1200888))
					if *memmap.PtrUint32(6112660, 1200888) == 0 {
						v257 = nox_xxx_getTTByNameSpriteMB_44CFC0("TeleportPoof")
						*memmap.PtrUint32(6112660, 1200888) = uint32(v257)
					}
					goto LABEL_857
				case 138:
					if *memmap.PtrUint32(6112660, 1200900) == 0 {
						*memmap.PtrUint32(6112660, 1200900) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("Smoke"))
						*memmap.PtrUint32(6112660, 1200896) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("Puff"))
					}
					v259 = int32(uintptr(unsafe.Pointer(nox_xxx_spriteLoadAdd_45A360_drawable(*memmap.PtrInt32(6112660, 1200900), int32(*(*int16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))), int32(*(*int16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))))))))
					if v259 != 0 {
						*(*uint16)(unsafe.Pointer(uintptr(v259 + 104))) = 20
						nox_xxx_sprite_45A110_drawable((*nox_drawable)(unsafe.Pointer(uintptr(v259))))
					}
					v260 = 6
					for {
						v341 = int32(*(*int16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))) + randomIntMinMax(-15, 15)
						v261 = randomIntMinMax(-15, 15)
						v262 = int32(uintptr(unsafe.Pointer(nox_xxx_spriteLoadAdd_45A360_drawable(*memmap.PtrInt32(6112660, 1200896), int32(*(*int16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))+v261, v341))))
						if v262 != 0 {
							*(*uint16)(unsafe.Pointer(uintptr(v262 + 104))) = uint16(int16(randomIntMinMax(5, 25)))
							nox_xxx_sprite_45A110_drawable((*nox_drawable)(unsafe.Pointer(uintptr(v262))))
						}
						v260--
						if v260 == 0 {
							break
						}
					}
					goto LABEL_868
				case 139:
					v257 = int32(*memmap.PtrUint32(6112660, 1200892))
					if *memmap.PtrUint32(6112660, 1200892) == 0 {
						v257 = nox_xxx_getTTByNameSpriteMB_44CFC0("DamagePoof")
						*memmap.PtrUint32(6112660, 1200892) = uint32(v257)
					}
				LABEL_857:
					v258 = int32(*(*int16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))) + 2
				LABEL_858:
					*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_xxx_spriteLoadAdd_45A360_drawable(v257, int32(*(*int16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))), v258))))
					if uint32(int32(v5)) == 0 {
						goto LABEL_868
					}
					nox_xxx_sprite_45A110_drawable((*nox_drawable)(unsafe.Pointer(uintptr(v5))))
					data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 5))
				case 163:
					v340 = float32(nox_xxx_gamedataGetFloat_419D40(CString("ManaBombOutRadius")))
					v252 = uint32(int(v340))
					v357 = 150
					v253 = int32(v252 >> 2)
					for i = int32(v252 >> 2); ; v253 = i {
						v254 = v253 + randomIntMinMax(0, int32(v252))
						if v254 > int32(v252) {
							v254 = int32(v252)
						}
						v255 = randomIntMinMax(0, math.MaxUint8)
						*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_xxx_spriteLoadAdd_45A360_drawable(*memmap.PtrInt32(6112660, 1200784), int32(*(*int16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))+v254**memmap.PtrInt32(0x587000, uint32(v255*8)+0x2EE58)/16, int32(*(*int16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))))+v254**memmap.PtrInt32(0x587000, uint32(v255*8)+0x2EE5C)/16))))
						v256 = int32(v5)
						if uint32(int32(v5)) != 0 {
							*(*uint32)(unsafe.Pointer(uintptr(v5 + 432))) = *(*uint32)(unsafe.Pointer(uintptr(v5 + 12))) << 12
							*(*uint32)(unsafe.Pointer(uintptr(v5 + 436))) = *(*uint32)(unsafe.Pointer(uintptr(v5 + 16))) << 12
							*(*uint8)(unsafe.Pointer(uintptr(v5 + 299))) = 0
							*(*uint32)(unsafe.Pointer(uintptr(v5 + 440))) = 0
							*(*uint32)(unsafe.Pointer(uintptr(v5 + 448))) = nox_frame_xxx_2598000 + uint32(randomIntMinMax(30, 40))
							*(*uint32)(unsafe.Pointer(uintptr(v256 + 444))) = nox_frame_xxx_2598000
							*(*uint16)(unsafe.Pointer(uintptr(v256 + 104))) = 0
							*(*uint8)(unsafe.Pointer(uintptr(v256 + 296))) = uint8(int8(randomIntMinMax(4, 10)))
							nox_xxx_sprite_45A110_drawable((*nox_drawable)(unsafe.Pointer(uintptr(v256))))
						}
						if func() int32 {
							p := &v357
							*p--
							return *p
						}() == 0 {
							break
						}
					}
					data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 5))
				default:
					goto LABEL_868
				}
			} else {
			LABEL_868:
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 5))
			}
		case 143:
			fallthrough
		case 144:
			fallthrough
		case 145:
			if nox_client_isConnected_43C700() != 0 {
				nox_xxx_netDrawRays_49BDD0(data)
			}
		LABEL_735:
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 9))
		case 147:
			if *memmap.PtrUint32(6112660, 1200852) == 0 {
				*memmap.PtrUint32(6112660, 1200852) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("Spark"))
				*memmap.PtrUint32(6112660, 1200856) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("MediumFireBoom"))
				*memmap.PtrUint32(6112660, 1197380) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("FireBoom"))
			}
			if nox_client_isConnected_43C700() == 0 {
				goto LABEL_916
			}
			v378.field_0 = int32(*(*int16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))
			v378.field_4 = int32(*(*int16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))))
			sub_49A150(&v378, *memmap.PtrInt32(6112660, 1200852), *((*uint8)(unsafe.Add(unsafe.Pointer(data), 5))))
			v235 = int32(*memmap.PtrUint32(6112660, 1197380))
			if int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 5)))) <= 170 {
				v235 = int32(*memmap.PtrUint32(6112660, 1200856))
			}
			*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_xxx_spriteLoadAdd_45A360_drawable(v235, int32(*(*int16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))), int32(*(*int16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))))))))
			goto LABEL_829
		case 148:
			if nox_client_isConnected_43C700() != 0 {
				v241 = int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 5))))) - int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))
				v242 = int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 7))))) - int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))))
				*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(nox_double2int(math.Sqrt(float64(v241*v241+v242*v242))) + 1)
				v371 = int32(v5)
				if int32(v5) >= 0 {
					v360 = 0
					v363 = 0
					i = v242 * 2
					v374 = v241 * 2
					v356 = uint32(int32(v5+2)) >> 1
					for {
						v243 = int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))) + v363/v371
						v244 = int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))) + v360/v371
						v245 = int32(uintptr(unsafe.Pointer(nox_xxx_spriteLoadAdd_45A360_drawable(*(*int32)(unsafe.Pointer(&dword_5d4594_1200796)), v243, v244))))
						v362 = v245
						if v245 != 0 {
							v246 = (*uint32)(unsafe.Pointer(uintptr(v245 + 432)))
							if v245 != -432 {
								*(*uint32)(unsafe.Pointer(uintptr(v245 + 432))) = uint32(v243 << 12)
								*(*uint32)(unsafe.Pointer(uintptr(v245 + 436))) = uint32(v244 << 12)
								v247 = int8(randomIntMinMax(0, math.MaxUint8))
								*(*uint8)(unsafe.Pointer(uintptr(v362 + 299))) = uint8(v247)
								*(*uint32)(unsafe.Add(unsafe.Pointer(v246), unsafe.Sizeof(uint32(0))*2)) = uint32(randomIntMinMax(1, 200))
								*(*uint32)(unsafe.Add(unsafe.Pointer(v246), unsafe.Sizeof(uint32(0))*4)) = nox_frame_xxx_2598000 + uint32(randomIntMinMax(20, 40))
								*(*uint32)(unsafe.Add(unsafe.Pointer(v246), unsafe.Sizeof(uint32(0))*3)) = nox_frame_xxx_2598000
							}
							v248 = int16(randomIntMinMax(15, 30))
							v249 = v362
							*(*uint16)(unsafe.Pointer(uintptr(v362 + 104))) = uint16(v248)
							*(*uint8)(unsafe.Pointer(uintptr(v249 + 296))) = uint8(int8(randomIntMinMax(-4, 4)))
							nox_xxx_sprite_45A110_drawable((*nox_drawable)(unsafe.Pointer(uintptr(v249))))
						}
						k = i
						*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*1)) = uint32(i + v360)
						v250 = v356 == 1
						v363 += v374
						v360 += i
						v356--
						if v250 {
							break
						}
					}
				}
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 9))
		case 149:
			if nox_client_isConnected_43C700() != 0 {
				sub_4C5020(int32(uintptr(unsafe.Pointer(data))))
				if randomIntMinMax(0, 100) < 25 {
					if *memmap.PtrUint32(0x852978, 8) != 0 {
						v225 = int32(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 5))))) - *(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x852978, 8) + 12))))
						v226 = nox_double2int(math.Sqrt(float64(uint32(v225*v225) + (uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 7)))))-*(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x852978, 8) + 16))))*(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 7)))))-*(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x852978, 8) + 16)))))))
						if v226 < 600 {
							clientPlaySoundSpecial(297, (600-v226)*100/600)
						}
					}
				}
				if nox_xxx_checkGameFlagPause_413A50() == 0 {
					v227 = int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 5))))) - int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))
					v228 = int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 7))))) - int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))))
					v229 = nox_double2int(math.Sqrt(float64(v227*v227 + v228*v228)))
					if v229 == 0 {
						v229 = 1
					}
					v230 = int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 5))))) - v227*4/v229
					v231 = int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 7))))) - v228*4/v229
					*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_xxx_spriteLoadAdd_45A360_drawable(*(*int32)(unsafe.Pointer(&dword_5d4594_1200796)), v230, v231))))
					v362 = int32(v5)
					if uint32(int32(v5)) != 0 {
						v232 = (*uint32)(unsafe.Pointer(uintptr(v5 + 432)))
						if uint32(int32(v5)) != 0xFFFFFE50 {
							*(*uint32)(unsafe.Pointer(uintptr(v5 + 432))) = uint32(v230 << 12)
							*(*uint32)(unsafe.Pointer(uintptr(v5 + 436))) = uint32(v231 << 12)
							v233 = int8(randomIntMinMax(0, math.MaxUint8))
							*(*uint8)(unsafe.Pointer(uintptr(v362 + 299))) = uint8(v233)
							*(*uint32)(unsafe.Add(unsafe.Pointer(v232), unsafe.Sizeof(uint32(0))*2)) = uint32(randomIntMinMax(1, 1500))
							*(*uint32)(unsafe.Add(unsafe.Pointer(v232), unsafe.Sizeof(uint32(0))*4)) = nox_frame_xxx_2598000 + uint32(randomIntMinMax(5, 20))
							*(*uint32)(unsafe.Add(unsafe.Pointer(v232), unsafe.Sizeof(uint32(0))*3)) = nox_frame_xxx_2598000
						}
						v234 = v362
						*(*uint16)(unsafe.Pointer(uintptr(v362 + 104))) = 22
						*(*uint8)(unsafe.Pointer(uintptr(v234 + 296))) = uint8(int8(randomIntMinMax(-4, 4)))
						nox_xxx_sprite_45A110_drawable((*nox_drawable)(unsafe.Pointer(uintptr(v234))))
					}
				}
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 9))
		case 150:
			if nox_client_isConnected_43C700() != 0 {
				if *memmap.PtrUint32(6112660, 1200860) == 0 {
					*memmap.PtrUint32(6112660, 1200860) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("BlueSpark"))
				}
				v239 = 5
				for {
					*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_xxx_spriteLoadAdd_45A360_drawable(*memmap.PtrInt32(6112660, 1200860), int32(*(*int16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))), int32(*(*int16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))))))))
					v240 = int32(v5)
					if uint32(int32(v5)) != 0 {
						if uint32(int32(v5)) != 0xFFFFFE50 {
							*(*uint32)(unsafe.Pointer(uintptr(v5 + 432))) = *(*uint32)(unsafe.Pointer(uintptr(v5 + 12))) << 12
							*(*uint32)(unsafe.Pointer(uintptr(v5 + 436))) = *(*uint32)(unsafe.Pointer(uintptr(v5 + 16))) << 12
							*(*uint8)(unsafe.Pointer(uintptr(v5 + 299))) = uint8(int8(randomIntMinMax(0, math.MaxUint8)))
							*(*uint32)(unsafe.Pointer(uintptr(v240 + 440))) = uint32(randomIntMinMax(1333, 4000))
							*(*uint32)(unsafe.Pointer(uintptr(v240 + 448))) = nox_frame_xxx_2598000 + uint32(randomIntMinMax(5, 20))
							*(*uint32)(unsafe.Pointer(uintptr(v240 + 444))) = nox_frame_xxx_2598000
						}
						*(*uint16)(unsafe.Pointer(uintptr(v240 + 104))) = 20
						*(*uint8)(unsafe.Pointer(uintptr(v240 + 296))) = uint8(int8(randomIntMinMax(-5, 5)))
						nox_xxx_sprite_45A110_drawable((*nox_drawable)(unsafe.Pointer(uintptr(v240))))
					}
					v239--
					if v239 == 0 {
						break
					}
				}
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 5))
		case 151:
			if nox_client_isConnected_43C700() != 0 {
				sub_4355B0(int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))) / 3)
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 2))
		case 152:
			if *memmap.PtrUint32(6112660, 1200844) == 0 {
				*memmap.PtrUint32(6112660, 1200844) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("GreenZap"))
			}
			if nox_client_isConnected_43C700() != 0 {
				*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer((*nox_drawable)(unsafe.Add(unsafe.Pointer(nox_xxx_spriteLoadAdd_45A360_drawable(*memmap.PtrInt32(6112660, 1200844), int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))+(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 5)))))-int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))/2, int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))))+(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 7)))))-int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))))/2)), unsafe.Sizeof(nox_drawable{})*432)))))
				*(*uint8)(unsafe.Pointer(uintptr(v5))) = 0
				*(*uint32)(unsafe.Pointer(uintptr(v5 + 5))) = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))
				*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*1)) = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 5))))
				*(*uint32)(unsafe.Pointer(uintptr(v5 + 9))) = *(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*1))
				k = int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 9)))))
				*(*uint32)(unsafe.Pointer(uintptr(v5 + 1))) = uint32(k)
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 11))
		case 154:
			if nox_client_isConnected_43C700() != 0 {
				nox_xxx____setargv_11_473920()
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 5))
		case 155:
			fallthrough
		case 156:
			fallthrough
		case 157:
			if nox_client_isConnected_43C700() != 0 {
				nox_xxx_mapGenClientText_4A9D00(data)
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
		case 158:
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v376))), 0)) = *((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))
			switch uint8(int8(v376)) {
			case 1:
				fallthrough
			case 2:
				fallthrough
			case 3:
				fallthrough
			case 4:
				fallthrough
			case 5:
				fallthrough
			case 6:
				fallthrough
			case 7:
				if nox_client_isConnected_43C700() != 0 {
					nox_xxx_clientAddRayEffect_49C160(int32(uintptr(unsafe.Pointer(data))))
				}
				goto LABEL_1070
			case 8:
				fallthrough
			case 9:
				fallthrough
			case 10:
				fallthrough
			case 11:
				fallthrough
			case 12:
				fallthrough
			case 13:
				fallthrough
			case 14:
				if nox_client_isConnected_43C700() == 0 {
					goto LABEL_1070
				}
				nox_xxx_clientRemoveRayEffect_49C450(int32(uintptr(unsafe.Pointer(data))))
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 7))
			default:
				return 0
			}
		case 159:
			v297 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v298 = v297
			if nox_client_isConnected_43C700() != 0 {
				*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(func() *uint32 {
					if nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))) != 0 {
						return nox_xxx_netSpriteByCodeStatic_45A720(v298)
					}
					return nox_xxx_netSpriteByCodeDynamic_45A6F0(v298)
				}())))
				if uint32(int32(v5)) != 0 {
					*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(int32(v5 + 432))
					*(*uint32)(unsafe.Pointer(uintptr(v5))) = nox_frame_xxx_2598000
					i = int32(*(*byte)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 4)))))
					*(*float32)(unsafe.Pointer(uintptr(v5 + 8))) = float32(float64(i))
					k = int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))
					i = int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))
					*(*float32)(unsafe.Pointer(uintptr(v5 + 4))) = float32(float64(i))
					*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*1)) = uint32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 5))))
					i = int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 5))))
					*(*float32)(unsafe.Pointer(uintptr(v5 + 12))) = float32(float64(i))
				}
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 6))
		case 160:
			if nox_client_isConnected_43C700() == 0 {
				goto LABEL_1163
			}
			nox_xxx_fxDrawTurnUndead_499880((*int16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 5))
		case 161:
			if nox_client_isConnected_43C700() == 0 {
				goto LABEL_916
			}
			if *memmap.PtrUint32(6112660, 1200864) == 0 {
				*memmap.PtrUint32(6112660, 1200864) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("ArrowTrap1Smoke"))
				*memmap.PtrUint32(6112660, 1200868) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("ArrowTrap2Smoke"))
			}
			v251 = int32(*(*int16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))
			v339 = int32(*(*int16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))))
			if int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 5)))) == 1 {
				*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_xxx_spriteLoadAdd_45A360_drawable(*memmap.PtrInt32(6112660, 1200864), v251+15, v339))))
				if uint32(int32(v5)) != 0 {
					nox_xxx_sprite_45A110_drawable((*nox_drawable)(unsafe.Pointer(uintptr(v5))))
					data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 6))
					break
				}
			} else {
				*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_xxx_spriteLoadAdd_45A360_drawable(*memmap.PtrInt32(6112660, 1200868), v251-3, v339))))
			LABEL_829:
				if uint32(int32(v5)) != 0 {
					nox_xxx_sprite_45A110_drawable((*nox_drawable)(unsafe.Pointer(uintptr(v5))))
					data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 6))
					break
				}
			}
		LABEL_916:
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 6))
		case 162:
			if nox_client_isConnected_43C700() != 0 {
				if *memmap.PtrUint32(6112660, 1200848) == 0 {
					*memmap.PtrUint32(6112660, 1200848) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("HealOrb"))
				}
				v221 = uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 9)))))
				k = int32(v221 & 0xFFFC)
				if k <= 28 {
					v222 = int32(v221 >> 2)
				} else {
					v222 = 7
				}
				v223 = v222 + 1
				if v222+1 > 0 {
					for {
						*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v368))), 0)) = uint8(int8(randomIntMinMax(6, 12)))
						v323 = int8(v368)
						v318 = randomIntMinMax(-20, 20)
						v224 = randomIntMinMax(-20, 20)
						sub_499490(*memmap.PtrInt32(6112660, 1200848), (*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))), v224, v318, v323, 0)
						v223--
						if v223 == 0 {
							break
						}
					}
				}
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 11))
		case 164:
			v15 = (*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), nox_xxx_netCliProcUpdateStream_494A60((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)), a1, (*uint32)(unsafe.Pointer(&v384[0])))))), 1))
			if uintptr(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(v15)))))) >= uintptr(unsafe.Pointer(end)) {
				return 1
			}
			for {
				*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_xxx_netCliUpdateStream2_494C30(v15, a1, &v384[0]))))
				if int32(v5) <= 0 {
					break
				}
				v15 = (*uint8)(unsafe.Add(unsafe.Pointer(v15), v5))
				if uintptr(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(v15)))))) >= uintptr(unsafe.Pointer(end)) {
					return 1
				}
			}
			if int32(v5) < 0 {
				*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(-int32(v5))
			}
			data = (*uint8)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v15), v5))))))))
		case 166:
			fallthrough
		case 167:
			if nox_client_isConnected_43C700() == 0 {
				goto LABEL_651
			}
			v338 = int32(*(*byte)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))
			v321 = (int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2))))) >> 9) & 126
			v317 = int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2))))) & 1023
			if *(*byte)(unsafe.Pointer(data)) == 167 {
				sub_452E10(v317, v321, v338)
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
			} else {
				sub_452DC0(v317, v321, v338)
			LABEL_651:
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
			}
		case 168:
			v198 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v199 = v198
			libc.MemSet(memmap.PtrOff(6112660, 1197384), 0, 636)
			v200 = int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))
			if v200&8 != 0 {
				v201 = strMan.GetStringInFile((*byte)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 11)))), "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
				nox_wcscpy((*libc.WChar)(memmap.PtrOff(6112660, 1197384)), v201)
				v202 = 1
			} else if v200&2 != 0 {
				if v200&1 != 0 {
					v203 = strMan.GetStringInFile("Guirank.c:team", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
					nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 1197384)), libc.CWString("%s: %S"), v203, (*uint8)(unsafe.Add(unsafe.Pointer(data), 11)))
				} else {
					nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 1197384)), libc.CWString("%S"), (*uint8)(unsafe.Add(unsafe.Pointer(data), 11)))
				}
				v202 = 1
			} else {
				if v200&1 != 0 {
					v204 = strMan.GetStringInFile("Guirank.c:team", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
					nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 1197384)), libc.CWString("%s: %s"), v204, (*uint8)(unsafe.Add(unsafe.Pointer(data), 11)))
				} else {
					nox_wcscpy((*libc.WChar)(memmap.PtrOff(6112660, 1197384)), (*libc.WChar)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 11)))))
				}
				v202 = 2
			}
			if nox_client_isConnected_43C700() == 0 {
				goto LABEL_693
			}
			if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_NET_DEBUG)) {
				nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			}
			if int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))&16 != 0 {
				if nox_xxx_gameGetPlayState_4356B0() == 3 {
					nox_wcscpy((*libc.WChar)(memmap.PtrOff(6112660, 1200068)), (*libc.WChar)(memmap.PtrOff(6112660, 1197384)))
					v205 = strMan.GetStringInFile("guiserv.c:Notice", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
					NewDialogWindow(nil, (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v205)))))), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(memmap.PtrOff(6112660, 1200068)))))), 33, nil, nil)
					k = v202 * int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 8))))
					data = (*uint8)(unsafe.Add(unsafe.Pointer(data), k+11))
					break
				}
			} else if int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))) != 0 {
				if nox_xxx_gameGetPlayState_4356B0() == 3 {
					v206 = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))))
					v207 = v206
					if v206 != nil {
						if nox_xxx_playerCanTalkMB_57A160(int32(uintptr(unsafe.Pointer(v206)))) == 0 {
							nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_YELLOW)), (*libc.WChar)(memmap.PtrOff(0x587000, 158984)), (*byte)(unsafe.Add(unsafe.Pointer(v207), 4704)), memmap.PtrOff(6112660, 1197384))
							nox_xxx_createTextBubble_48D880(int32(uintptr(unsafe.Pointer(data))), (*libc.WChar)(memmap.PtrOff(6112660, 1197384)))
							k = v202 * int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 8))))
							data = (*uint8)(unsafe.Add(unsafe.Pointer(data), k+11))
							break
						}
					} else {
						nox_xxx_createTextBubble_48D880(int32(uintptr(unsafe.Pointer(data))), (*libc.WChar)(memmap.PtrOff(6112660, 1197384)))
						if nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))) != 0 {
							*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_xxx_netSpriteByCodeStatic_45A720(v199))))
						} else {
							*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_xxx_netSpriteByCodeDynamic_45A6F0(v199))))
						}
						if uint32(int32(v5)) != 0 {
							v322 = int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))
							v208 = int32(uintptr(unsafe.Pointer(nox_get_thing_pretty_name(int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 108))))))))
							nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), (*libc.WChar)(memmap.PtrOff(0x587000, 159000)), v208, v322, memmap.PtrOff(6112660, 1197384))
						}
					}
				}
			} else if nox_xxx_gameGetPlayState_4356B0() == 3 {
				nox_xxx_printCentered_445490((*libc.WChar)(memmap.PtrOff(6112660, 1197384)))
				k = v202 * int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 8))))
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), k+11))
				break
			}
		LABEL_693:
			k = v202 * int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 8))))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), k+11))
		case 169:
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), nox_client_handlePacketInform_4C9BF0(int32(uintptr(unsafe.Pointer(data))))))
		case 170:
			var n int32
			if noxflags.HasGame(noxflags.GameHost) {
				n = 5
			} else {
				n = 1
			}
			if nox_client_isConnected_43C700() != 0 {
				var v366 [5]byte
				v366[0] = 171
				if noxflags.HasGame(noxflags.GameHost) {
					*(*uint32)(unsafe.Pointer(&v366[1])) = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))
				} else {
					*(*uint32)(unsafe.Pointer(&v366[1])) = v364
				}
				nox_netlist_addToMsgListCli_40EBC0(a1, 0, (*uint8)(unsafe.Pointer(&v366[0])), 5)
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), n))
		case 171:
			if nox_client_isConnected_43C700() != 0 {
				sub_4E55A0(uint8(int8(a1)), int32(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 5))
		case 174:
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
		case 175:
			dword_5d4594_1200768 = 0
			v34 = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))
			if uint32(nox_xxx_get3512_40A350()) < v34 {
				nox_xxx_set3512_40A340(int32(v34))
				noxflags.HasGame(noxflags.GameModeElimination)
				v355 = bool2int(noxflags.HasGame(noxflags.GameModeCTF))
				v360 = bool2int(noxflags.HasGame(noxflags.GameModeCTF))
				v363 = bool2int(noxflags.HasGame(noxflags.GameModeQuest))
				if !noxflags.HasGame(noxflags.GameHost) {
					noxflags.UnsetGame(noxflags.GameModeKOTR | noxflags.GameModeCTF | noxflags.GameModeFlagBall | noxflags.GameModeChat | noxflags.GameModeArena | noxflags.GameModeSolo10 | noxflags.GameModeElimination | noxflags.GameModeCoop | noxflags.GameModeQuest | noxflags.GameOnline | noxflags.GameFlag15 | noxflags.GameFlag16 | noxflags.GameNotQuest | noxflags.GameFlag18 | noxflags.GamePause)
					noxflags.SetGame(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 9)))))
					sub_409E40(int32(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 13))))))
					nox_client_setVersion_409AE0(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 5)))))
					nox_xxx_servSetPlrLimit_409F80(int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 17)))))
					v35 = nox_xxx_cliGamedataGet_416590(0)
					if *(*byte)(unsafe.Add(unsafe.Pointer(v35), 56)) != byte(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 19)))) || int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v35))), unsafe.Sizeof(uint16(0))*27)))) != int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 18)))) {
						dword_5d4594_1200768 = 1
					}
					*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v35))), unsafe.Sizeof(uint16(0))*27))) = uint16(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 18))))
					*(*byte)(unsafe.Add(unsafe.Pointer(v35), 56)) = byte(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 19))))
					if int32(int8(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 9))))) >= 0 {
						*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v35))), unsafe.Sizeof(uint16(0))*26))) = *(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 9))))
						libc.StrCpy(v35, nox_xxx_mapGetMapName_409B40())
					}
					if v355 != 0 || !noxflags.HasGame(noxflags.GameModeCTF) {
						if !noxflags.HasGame(noxflags.GameModeCTF) {
							sub_455C10()
						}
					} else {
						sub_455A50(2)
					}
					if v360 != 0 || !noxflags.HasGame(noxflags.GameModeFlagBall) {
						if !noxflags.HasGame(noxflags.GameModeFlagBall) {
							sub_456050()
						}
					} else {
						sub_455F60()
					}
					if v363 == 1 && !noxflags.HasGame(noxflags.GameModeQuest) {
						for j = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0())); j != nil; j = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(j))))))))) {
							nox_xxx_cliPlayerRespawn_417680(int32(uintptr(unsafe.Pointer(j))), -1)
						}
					}
					if !noxflags.HasGame(noxflags.GameFlag18) && nox_xxx_gameGetPlayState_4356B0() == 3 {
						nox_xxx_setContinueMenuOrHost_43DDD0(0)
						nox_game_exit_xxx_43DE60()
					}
				}
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 20))
		case 176:
			if !noxflags.HasGame(noxflags.GameHost) {
				v37 = nox_xxx_cliGamedataGet_416590(0)
				libc.StrNCpy((*byte)(unsafe.Add(unsafe.Pointer(v37), 9)), (*byte)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))), 15)
				if libc.MemCmp(memmap.PtrOff(6112660, 1200732), unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 17))), 20) != 0 {
					dword_5d4594_1200768 = 1
				}
				if *memmap.PtrUint32(6112660, 1200752) != *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 37)))) || *memmap.PtrUint32(6112660, 1200756) != *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 41)))) {
					dword_5d4594_1200768 = 1
				}
				libc.MemCpy(unsafe.Add(unsafe.Pointer(v37), 24), unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 17))), 20)
				sub_4540E0(int32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v37), 24))))))
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v37))), unsafe.Sizeof(uint32(0))*11))) = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 37))))
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v37))), unsafe.Sizeof(uint32(0))*12))) = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 41))))
				if *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 45)))) != 0 {
					sub_40A1F0(1)
					sub_40A310(int32(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 45))))))
				} else {
					sub_40A1F0(0)
				}
				libc.MemCpy(memmap.PtrOff(6112660, 1200708), unsafe.Pointer(v37), 58)
				if nox_client_isConnected_43C700() != 0 && dword_5d4594_1200768 != 0 {
					v38 = strMan.GetStringInFile("OptionsChanged", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
					nox_xxx_printCentered_445490(v38)
					clientPlaySoundSpecial(310, 100)
				}
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 49))
		case 177:
			v39 = nox_xxx_cliGamedataGet_416590(int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v40))), 0)) = uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v39), 52)))
			var v4b int32 = int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 54))))
			libc.MemCpy(unsafe.Pointer(v39), unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2))), 58)
			if ((v40 >> 5) & 1) != uint32((v4b>>5)&1) {
				sub_4573A0()
			}
			if sub_4169C0() == 0 {
				if sub_459DA0() == 0 {
					if int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))) == 1 {
						sub_4165F0(1, 0)
					}
					if noxflags.HasGame(noxflags.GameModeChat) {
						if dword_5d4594_1200832 != 0 {
							v330 = int32(*memmap.PtrUint32(0x8531A0, 2576) + 4704)
							v41 = strMan.GetStringInFile("NameChange", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
							nox_swprintf(&v401[0], v41, v330)
							nox_xxx_printCentered_445490(&v401[0])
							dword_5d4594_1200832 = 0
						}
						if nox_server_sanctuaryHelp_54276 != 0 {
							nox_xxx_cliShowHelpGui_49C560()
						}
					}
				}
				nox_xxx_cliSetSettingsAcquired_4169D0(1)
			}
			sub_459C30()
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 60))
		case 178:
			v18 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v19 = v18
			if nox_client_isConnected_43C700() != 0 {
				*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(func() *uint32 {
					if nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))) != 0 {
						return nox_xxx_netSpriteByCodeStatic_45A720(v19)
					}
					return nox_xxx_netSpriteByCodeDynamic_45A6F0(v19)
				}())))
				if uint32(int32(v5)) != 0 {
					*(*uint8)(unsafe.Pointer(uintptr(v5 + 299))) = *((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))
				}
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
		case 179:
			v20 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v21 = v20
			if nox_client_isConnected_43C700() != 0 {
				*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(func() *uint32 {
					if nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))) != 0 {
						return nox_xxx_netSpriteByCodeStatic_45A720(v21)
					}
					return nox_xxx_netSpriteByCodeDynamic_45A6F0(v21)
				}())))
				v22 = int32(v5)
				if uint32(int32(v5)) != 0 {
					nox_xxx_spriteSetActiveMB_45A990_drawable(int32(v5))
					v328 = float32(float64(int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))) * 16 / 10))
					nox_xxx_spriteChangeIntensity_484D70_light_intensity(v22+136, v328)
					nox_xxx_spriteSetFrameMB_45AB80(v22, int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))*8/50)
					if *(*uint32)(unsafe.Pointer(uintptr(v22 + 308))) == 8 {
						*(*uint32)(unsafe.Pointer(uintptr(v22 + 308))) = 7
					}
				}
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
		case 180:
			v24 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v25 = v24
			if nox_client_isConnected_43C700() != 0 {
				*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(func() *uint32 {
					if nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))) != 0 {
						return nox_xxx_netSpriteByCodeStatic_45A720(v25)
					}
					return nox_xxx_netSpriteByCodeDynamic_45A6F0(v25)
				}())))
				v26 = int32(v5)
				if uint32(int32(v5)) != 0 {
					if int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))) != 0 {
						*(*uint32)(unsafe.Pointer(uintptr(v5 + 112))) |= 0x80000
						nox_xxx_spriteChangeIntensity_484D70_light_intensity(int32(v5+136), 41.958)
					} else {
						*(*uint32)(unsafe.Pointer(uintptr(v5 + 112))) &= 0xFFF7FFFF
						nox_xxx_spriteChangeIntensity_484D70_light_intensity(int32(v5+136), 0.0)
					}
					nox_xxx_spriteSetFrameMB_45AB80(v26, int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))))
					*(*uint32)(unsafe.Pointer(uintptr(v26 + 288))) = nox_frame_xxx_2598000
				}
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
		case 181:
			if nox_client_isConnected_43C700() != 0 && (func() uint32 {
				v329 = int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 7)))))
				v320 = int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 5)))))
				v31 = uint32(nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))))
				*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_xxx_spriteCreate_48E970(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))), v31, v320, v329))))
				return uint32(int32(v5))
			}()) != 0 {
				*(*uint16)(unsafe.Pointer(uintptr(v5 + 508))) = *(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 9))))
				*(*float32)(unsafe.Pointer(uintptr(v5 + 468))) = float32(float64(*(*byte)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 12))))) * 0.0625)
				*(*float32)(unsafe.Pointer(uintptr(v5 + 472))) = float32(float64(*(*byte)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 13))))) * 0.0625)
				*(*float32)(unsafe.Pointer(uintptr(v5 + 476))) = float32(float64(*(*byte)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 11))))) * 0.0625)
				*(*uint32)(unsafe.Pointer(uintptr(v5 + 316))) = nox_frame_xxx_2598000
				*(*uint32)(unsafe.Pointer(uintptr(v5 + 324))) = uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 5)))))
				v32 = int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 7)))))
				*(*uint32)(unsafe.Pointer(uintptr(v5 + 460))) = uint32(cgoFuncAddr(nox_xxx_sprite_4CA540))
				*(*uint32)(unsafe.Pointer(uintptr(v5 + 328))) = uint32(v32)
				nox_xxx_spriteToList_49BC80_drawable((*uint32)(unsafe.Pointer(uintptr(v5))))
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 14))
		case 184:
			if nox_xxx_mapDownloadStart_4ABAD0((*byte)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 8)))), *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 4))))) == 0 {
				nox_xxx_cliSendCancelMap_43CAB0()
				nox_xxx_mapSetDownloadInProgress_4AB560(0)
				nox_xxx_mapSetDownloadOK_4AB570(0)
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 88))
		case 185:
			nox_xxx_netMapDownloadPart_4AB7C0(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2)))), unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 6))), uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 4))))))
			k = int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 4)))))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), k+6))
		case 186:
			nox_xxx_mapDeleteFile_4AB720()
			nox_xxx_mapSetDownloadInProgress_4AB560(0)
			nox_xxx_mapSetDownloadOK_4AB570(0)
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 2))
		case 189:
			if nox_client_isConnected_43C700() != 0 {
				if int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))) == 1 {
					v290 = strMan.GetStringInFile("sysopAccess", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
				} else {
					v290 = strMan.GetStringInFile("invalidPass", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
				}
				nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v290)
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 2))
		case 194:
			v263 = int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))
			switch v263 {
			case 0:
				v342 = int8(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 136))))
				v324 = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 4))))
				v315 = int8(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 2))))
				v264 = nox_xxx_netGet_43C750()
				sub_40B5D0(uint32(v264), v315, (*byte)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 8)))), v324, v342)
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 140))
			case 1:
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v5))), 4)) = *((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))
				v343 = (*byte)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*1)))))
				v325 = int8(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 2))))
				v265 = nox_xxx_netGet_43C750()
				sub_40BFF0(v265, v325, v343)
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
			case 2:
				v365[0] = 194
				v365[1] = 3
				v365[2] = byte(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 2))))
				*(*uint16)(unsafe.Pointer(&v365[4])) = *(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 4))))
				v266 = nox_xxx_netGet_43C750()
				nox_xxx_netSendSock_552640(uint32(v266), &v365[0], 6, int8(NOX_NET_SEND_NO_LOCK|NOX_NET_SEND_FLAG2))
				v344 = uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 6)))))
				v319 = *(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 4))))
				v316 = *((*uint8)(unsafe.Add(unsafe.Pointer(data), 2)))
				v267 = nox_xxx_netGet_43C750()
				sub_40B250(v267, v316, v319, unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 8))), v344)
				*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*1)) = uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 6)))))
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), int32(*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint16(0))*2)))+8))
			case 3:
				v345 = int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 4)))))
				v326 = int8(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 2))))
				v268 = nox_xxx_netGet_43C750()
				sub_40BF60(v268, v326, v345)
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 6))
			case 4:
				v346 = int8(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 2))))
				v269 = nox_xxx_netGet_43C750()
				sub_40C030(v269, v346)
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
			case 5:
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&k))), 0)) = *((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))
				sub_40B720(k, *((*uint8)(unsafe.Add(unsafe.Pointer(data), 2))))
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
			case 6:
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v263))), 0)) = *((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))
				v347 = int8(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 2))))
				v327 = v263
				v270 = nox_xxx_netGet_43C750()
				sub_40C070(v270, v327, v347)
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
			default:
				return 0
			}
		case 195:
			nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ADDITIONAL_NETWORK_TEST)) {
				nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			}
			if int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))) != 0 || int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))) != 0 {
				v27 = nox_xxx_spriteCreate_48E970(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))), uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))), int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 5))))), int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 7))))))
				if v27 != nil {
					*(*uint32)(unsafe.Add(unsafe.Pointer(v27), unsafe.Sizeof(uint32(0))*72)) = nox_frame_xxx_2598000
					nox_xxx_spriteSetFrameMB_45AB80(int32(uintptr(unsafe.Pointer(v27))), int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 10)))))
					v28 = uint8(int8((int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 9)))) >> 4) & 7))
					*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v27))), 297))) = v28
					if int32(v28) > 3 {
						*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v27))), 297))) = uint8(int8(int32(v28) + 1))
					}
					if *(*uint32)(unsafe.Add(unsafe.Pointer(v27), unsafe.Sizeof(uint32(0))*69)) != uint32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 11)))) {
						*(*uint32)(unsafe.Add(unsafe.Pointer(v27), unsafe.Sizeof(uint32(0))*79)) = nox_frame_xxx_2598000
						k = int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 11))))
						*(*uint32)(unsafe.Add(unsafe.Pointer(v27), unsafe.Sizeof(uint32(0))*69)) = uint32(k)
					}
				}
				*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*1)) = uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))
				if uint32(*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint16(0))*2))) == nox_player_netCode_85319C && sub_416120(8) {
					nox_xxx_cliUpdateCameraPos_435600(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 5))))), int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 7))))))
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 12))
			} else {
				nox_xxx_cliUpdateCameraPos_435600(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 5))))), int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 7))))))
				nox_xxx_setKeybTimeout_4160D0(8)
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 12))
			}
		case 196:
			switch *((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))) {
			case 0:
				if nox_client_isConnected_43C700() == 0 {
					goto LABEL_888
				}
				libc.MemCpy(unsafe.Pointer(&v386[0]), unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 18))), int(int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 15))))*2))
				v386[*((*uint8)(unsafe.Add(unsafe.Pointer(data), 15)))] = 0
				if int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 17)))) != 0 {
					v271 = nox_server_teamTitle_418C20(int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 16)))))
					nox_swprintf(&v386[0], v271)
				}
				v272 = (*byte)(unsafe.Pointer(nox_xxx_clientGetTeamColor_418AB0(int32(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2))))))))
				if v272 == nil {
					v272 = (*byte)(unsafe.Pointer(nox_xxx_teamCreate_4186D0(int8(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 2)))))))
					if v272 == nil {
						goto LABEL_888
					}
				}
				sub_418800((*libc.WChar)(unsafe.Pointer(v272)), &v386[0], 0)
				sub_418830(int32(uintptr(unsafe.Pointer(v272))), int32(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 6))))))
				nox_xxx_netChangeTeamID_419090(int32(uintptr(unsafe.Pointer(v272))), int32(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 10))))))
				*(*byte)(unsafe.Add(unsafe.Pointer(v272), 56)) = byte(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 16))))
				sub_457230(&v386[0])
				if (int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 14)))) & 1) == 0 {
					goto LABEL_888
				}
				v273 = nox_xxx_objGetTeamByNetCode_418C80(int32(nox_player_netCode_85319C))
				if v273 == nil {
					goto LABEL_888
				}
				if noxflags.HasGame(noxflags.GameHost) {
					nox_xxx_createAtImpl_4191D0(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v272), 57))), unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v273))))), 1, int32(nox_player_netCode_85319C), 1)
					*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*1)) = uint32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 15))))
					data = (*uint8)(unsafe.Add(unsafe.Pointer(data), (*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*1)))*2+18))
				} else {
					sub_419900(int32(uintptr(unsafe.Pointer(v272))), int32(uintptr(unsafe.Pointer(v273))), int16(uint16(nox_player_netCode_85319C)))
				LABEL_888:
					*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*1)) = uint32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 15))))
					data = (*uint8)(unsafe.Add(unsafe.Pointer(data), (*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*1)))*2+18))
				}
			case 1:
				v274 = uint32(nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 6)))))))
				v275 = v274
				if nox_client_isConnected_43C700() != 0 {
					if nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 6)))))) != 0 {
						*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_xxx_netSpriteByCodeStatic_45A720(int32(v275)))))
					} else {
						*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_xxx_netSpriteByCodeDynamic_45A6F0(int32(v275)))))
					}
					if uint32(int32(v5)) == 0 {
						*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_xxx_spriteCreate_48E970(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 8))))), v275, 0, 0))))
					}
					v276 = int32(v5 + 24)
					if uint32(int32(v5)) != 0xFFFFFFE8 {
						*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_xxx_clientGetTeamColor_418AB0(int32(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2)))))))))
						v277 = int32(v5)
						if uint32(int32(v5)) != 0 {
							nox_xxx_createAtImpl_4191D0(*(*uint8)(unsafe.Pointer(uintptr(v5 + 57))), unsafe.Pointer(uintptr(v276)), 0, int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 6))))), 0)
							sub_4571A0(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 6))))), int32(*(*uint8)(unsafe.Pointer(uintptr(v277 + 57)))))
						}
					}
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 10))
			case 2:
				if nox_client_isConnected_43C700() != 0 {
					*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_xxx_objGetTeamByNetCode_418C80(int32(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2)))))))))
					if uint32(int32(v5)) != 0 {
						nox_xxx_netChangeTeamMb_419570(int32(v5), int32(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2))))))
						sub_4571A0(int32(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2))))), 0)
					}
				}
				goto LABEL_902
			case 3:
				if nox_client_isConnected_43C700() != 0 {
					v278 = nox_xxx_objGetTeamByNetCode_418C80(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 6))))))
					if v278 != nil {
						*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_xxx_clientGetTeamColor_418AB0(int32(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2)))))))))
						v279 = int32(v5)
						if uint32(int32(v5)) != 0 {
							if sub_4196D0(int32(uintptr(unsafe.Pointer(v278))), int32(v5), int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 6))))), 0) != 0 {
								sub_4571A0(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 6))))), int32(*(*uint8)(unsafe.Pointer(uintptr(v279 + 57)))))
							}
						}
					}
				}
				goto LABEL_908
			case 4:
				if nox_client_isConnected_43C700() != 0 {
					*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_xxx_clientGetTeamColor_418AB0(int32(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2)))))))))
					if uint32(int32(v5)) != 0 {
						nox_xxx_teamRenameMB_418CD0((*libc.WChar)(unsafe.Pointer(uintptr(v5))), (*libc.WChar)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 6)))))
					}
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 46))
			case 5:
				if nox_client_isConnected_43C700() != 0 {
					*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_xxx_clientGetTeamColor_418AB0(int32(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2)))))))))
					if uint32(int32(v5)) != 0 {
						sub_418D80(int32(v5))
					}
				}
				goto LABEL_916
			case 6:
				if nox_client_isConnected_43C700() != 0 {
					*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_xxx_clientGetTeamColor_418AB0(int32(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2)))))))))
					v280 = (*libc.WChar)(unsafe.Pointer(uintptr(v5)))
					if uint32(int32(v5)) != 0 {
						nox_wcscpy(&v387[0], (*libc.WChar)(unsafe.Pointer(uintptr(v5))))
						sub_418F20((*nox_team_t)(unsafe.Pointer(v280)), 0)
						sub_456EA0(&v387[0])
					}
				}
				goto LABEL_920
			case 7:
				if nox_client_isConnected_43C700() != 0 {
					nox_server_teamsZzz_419030(0)
					sub_456FA0()
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 2))
			case 8:
				if nox_client_isConnected_43C700() != 0 && (func() uint32 {
					*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_xxx_clientGetTeamColor_418AB0(int32(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2)))))))))
					return uint32(int32(v5))
				}()) != 0 {
					nox_xxx_netChangeTeamID_419090(int32(v5), int32(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 6))))))
					data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 10))
				} else {
				LABEL_908:
					data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 10))
				}
			case 9:
				if nox_client_isConnected_43C700() != 0 {
					nox_server_teamsResetYyy_417D00()
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 2))
			case 12:
				if nox_client_isConnected_43C700() != 0 {
					*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(noxServer.getPlayerByID(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2)))))))))
					if uint32(int32(v5)) != 0 {
						*(*uint8)(unsafe.Pointer(uintptr(v5 + 2282))) = *((*uint8)(unsafe.Add(unsafe.Pointer(data), 4)))
					}
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 5))
			}
		case 197:
			sub_43B6E0()
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 1))
		case 198:
			sub_43B750()
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 1))
		case 199:
			v359 = -56
			v33 = nox_xxx_netGet_43C750()
			nox_xxx_netSendSock_552640(uint32(v33), (*byte)(unsafe.Pointer(&v359)), 1, int8(NOX_NET_SEND_FLAG2))
			sub_446380()
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 1))
		case 201:
			switch *((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))) {
			case 1:
				if nox_client_isConnected_43C700() != 0 {
					sub_4C1590()
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 2))
			case 2:
				if nox_client_isConnected_43C700() != 0 {
					sub_479280()
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 2))
			case 3:
				if nox_client_isConnected_43C700() != 0 {
					sub_4C1BC0(int32(uintptr(unsafe.Pointer(data))))
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
			case 4:
				if nox_client_isConnected_43C700() != 0 {
					nox_xxx_tradeClientAddItem_4C1790(int32(uintptr(unsafe.Pointer(data))))
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 15))
			case 5:
				if nox_client_isConnected_43C700() == 0 {
					goto LABEL_1133
				}
				sub_4C15D0(int32(uintptr(unsafe.Pointer(data))))
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
			case 6:
				if nox_client_isConnected_43C700() == 0 {
					data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 14))
					break
				}
				sub_4C1B50(int32(uintptr(unsafe.Pointer(data))))
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 14))
			case 7:
				if nox_client_isConnected_43C700() != 0 {
					nox_xxx_prepareP2PTrade_4C1BF0()
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 2))
			case 8:
				if nox_client_isConnected_43C700() != 0 {
					sub_479300(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2))))), int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 4))))), int32(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 6))))), int16(uint16(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 10)))))), int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 14))))))
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 18))
			case 9:
				if nox_client_isConnected_43C700() == 0 {
					goto LABEL_1133
				}
				sub_479480(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2))))))
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
			case 12:
				if nox_client_isConnected_43C700() != 0 {
					nox_xxx_netP2PStartTrade_4C1320(int32(uintptr(unsafe.Pointer(data))))
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 52))
			case 13:
				if nox_client_isConnected_43C700() != 0 {
					nox_xxx_cliStartShopDlg_478FD0((*libc.WChar)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 4)))), (*byte)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 54)))), int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2))))))
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 86))
			case 27:
				if nox_client_isConnected_43C700() == 0 {
					goto LABEL_1133
				}
				sub_479520(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2))))))
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
			case 29:
				if nox_client_isConnected_43C700() != 0 {
					sub_4795E0(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2))))), int32(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 4))))))
				}
				goto LABEL_1057
			case 31:
				if nox_client_isConnected_43C700() == 0 {
					goto LABEL_1057
				}
				sub_479740(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2))))), *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 4)))))
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 8))
			default:
				return 0
			}
		case 202:
			if nox_client_isConnected_43C700() != 0 {
				if int32(*(*int16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))) == -8531 {
					sub_48E940()
				} else {
					sub_48E8E0(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
				}
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
		case 203:
			if nox_client_isConnected_43C700() != 0 {
				sub_445450()
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 1))
		case 204:
			if nox_client_isConnected_43C700() != 0 {
				sub_48D5A0(int32(uintptr(unsafe.Pointer(data))))
			}
			k = int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), k+4))
		case 205:
			v86 = 0
			if nox_client_isConnected_43C700() != 0 {
				v87 = int8(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 2))))
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v370))), 0)) = uint8(int8(int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 2)))) & math.MaxInt8))
				if int32(v87) < 0 {
					v86 = 1
				}
				nox_xxx_netAbilityRewardCli_4611E0(int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))), int32(uint8(int8(v370))), (*byte)(unsafe.Pointer(uintptr(v86))))
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
		case 206:
			if nox_client_isConnected_43C700() != 0 {
				sub_461090(int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))), int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 2)))))
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
		case 207:
			if nox_client_isConnected_43C700() != 0 {
				sub_461120(int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))), int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 2)))))
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
		case 208:
			if int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))) == 3 {
				if nox_client_isConnected_43C700() != 0 {
					sub_479D30((*libc.WChar)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 34)))), int32(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 98))))), (*byte)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2)))), (*byte)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 102)))), int8(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 134)))))
					if *memmap.PtrUint32(0x8531A0, 2576) != 0 {
						nox_xxx_netNeedTimestampStatus_4174F0(*memmap.PtrInt32(0x8531A0, 2576), 512)
					}
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 135))
			} else if int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))) == 4 {
				if nox_client_isConnected_43C700() != 0 {
					sub_47A1F0()
					if *memmap.PtrUint32(0x8531A0, 2576) != 0 {
						nox_xxx_playerUnsetStatus_417530((*nox_playerInfo)(unsafe.Pointer(uintptr(*memmap.PtrInt32(0x8531A0, 2576)))), 512)
					}
				}
			LABEL_966:
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 2))
			}
		case 209:
			if nox_client_isConnected_43C700() != 0 {
				nox_xxx_netGuideRewardCli_45D140(int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))), int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 2))))&math.MaxInt8)
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
		case 210:
			v281 = uint32(nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))))
			v282 = v281
			if nox_client_isConnected_43C700() == 0 {
				goto LABEL_948
			}
			if nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))) != 0 {
				v283 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
				*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_xxx_netSpriteByCodeStatic_45A720(v283))))
			} else {
				v284 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
				*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_xxx_netSpriteByCodeDynamic_45A6F0(v284))))
			}
			if int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 5)))) == 1 {
				if uint32(int32(v5)) != 0 || (func() uint32 {
					*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_xxx_spriteCreate_48E970(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))), v282, 0, 0))))
					return uint32(int32(v5))
				}()) != 0 {
					sub_459DD0((*nox_drawable)(unsafe.Pointer(uintptr(v5))), int8(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 6)))))
					data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 7))
					break
				}
			} else if uint32(int32(v5)) != 0 {
				nox_xxx_cliRemoveHealthbar_459E30((*nox_drawable)(unsafe.Pointer(uintptr(v5))), int8(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 6)))))
			}
		LABEL_948:
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 7))
		case 211:
			if nox_client_isConnected_43C700() == 0 {
				goto LABEL_112
			}
			if (nox_frame_xxx_2598000 - *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 9))))) >= 30 {
				v359 = -44
				nox_xxx_netClientSend2_4E53C0(31, unsafe.Pointer(&v359), 1, 0, 1)
			LABEL_112:
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 13))
			} else {
				if *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))) != 0 {
					sub_40A1F0(1)
					v5 = sub_40A310(int32(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 5))))))
				} else {
					sub_40A1F0(0)
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 13))
			}
		case 213:
			v296 = int32(*memmap.PtrUint32(0x8531A0, 2576))
			switch *((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))) {
			case 1:
				if nox_client_isConnected_43C700() != 0 {
					if v296 != 0 {
						nox_xxx_journalEntryAdd_427490(v296, (*byte)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2)))), int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 66))))))
					}
					nox_xxx_cliBuildJournalString_469BC0()
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 68))
			case 2:
				if nox_client_isConnected_43C700() != 0 {
					if v296 != 0 {
						nox_xxx_journalEntryRemove_427590(v296, (*byte)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2)))))
					}
					nox_xxx_cliBuildJournalString_469BC0()
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 68))
			case 3:
				if nox_client_isConnected_43C700() != 0 && v296 != 0 {
					nox_xxx_journalUpdateEntry_4276B0(v296, (*byte)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2)))), int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 66))))))
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 68))
			}
		case 214:
			if nox_client_isConnected_43C700() != 0 {
				nox_client_lockScreenBriefing_450160(int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))), int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 2)))), 0)
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
		case 215:
			*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(noxServer.getPlayerByID(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))))))
			if uint32(int32(v5)) != 0 {
				*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&k))), unsafe.Sizeof(uint16(0))*0)) = *(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))
				*(*uint16)(unsafe.Pointer(uintptr(v5 + 2148))) = uint16(int16(k))
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 5))
		case 216:
			if nox_client_isConnected_43C700() != 0 {
				sub_455D80(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 2))), int8(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))
				sub_4705F0(int8(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))), int8(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))), int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 4))))))
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 6))
			} else {
			LABEL_920:
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 6))
			}
		case 217:
			if nox_client_isConnected_43C700() != 0 {
				sub_456140(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))
				sub_470650(int8(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))), int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2))))))
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
		case 218:
			v137 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v138 = v137
			if nox_client_isConnected_43C700() != 0 {
				if nox_xxx_unitSpriteCheckAlly_4951F0(v138) != 0 {
					sub_4950F0(v138, int8(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))))
				}
				nox_npc_set_328(v138, int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))))
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
			} else {
			LABEL_764:
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
			}
		case 219:
			v193 = uint32(nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))))
			v194 = v193
			if nox_client_isConnected_43C700() != 0 {
				if nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))) != 0 {
					v195 = nox_xxx_netSpriteByCodeStatic_45A720(int32(v194))
				} else {
					v195 = nox_xxx_netSpriteByCodeDynamic_45A6F0(int32(v194))
				}
				if v195 != nil || (func() *uint32 {
					v195 = nox_xxx_spriteCreate_48E970(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))), v194, 0, 0)
					return v195
				}()) != nil {
					sub_459DD0((*nox_drawable)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v195)))))), 1)
				}
				sub_495060(int32(v194), 0, 0)
			}
		LABEL_638:
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 5))
		case 220:
			v196 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v197 = v196
			if nox_client_isConnected_43C700() == 0 {
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
				break
			}
			sub_4950C0(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v337 = v197
			if nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))) != 0 {
				*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_xxx_netSpriteByCodeStatic_45A720(v197))))
			} else {
			LABEL_642:
				*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_xxx_netSpriteByCodeDynamic_45A6F0(v337))))
			}
		LABEL_643:
			if uint32(int32(v5)) != 0 {
				nox_xxx_cliRemoveHealthbar_459E30((*nox_drawable)(unsafe.Pointer(uintptr(v5))), 1)
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
		case 221:
			v90 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v91 = v90
			if nox_client_isConnected_43C700() == 0 {
				goto LABEL_1149
			}
			if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_NET_DEBUG)) {
				nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			}
			if uint32(v91) == nox_player_netCode_85319C {
				nox_xxx_cliSetTotalHealth_470C80(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))), int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 5))))))
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 7))
			} else if nox_xxx_unitSpriteCheckAlly_4951F0(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))) != 0 {
				sub_495120(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))), int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))), int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 5))))))
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 7))
			} else {
			LABEL_1149:
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 7))
			}
		case 222:
			v94 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v95 = v94
			if nox_client_isConnected_43C700() == 0 {
				goto LABEL_948
			}
			if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_NET_DEBUG)) {
				nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			}
			if uint32(v95) != nox_player_netCode_85319C {
				goto LABEL_948
			}
			nox_xxx_cliSetManaAndMax_470CE0(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))), int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 5))))))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 7))
		case 223:
			sub_460EB0(int32(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))), int8(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 5)))))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 6))
		case 224:
			v291 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v292 = v291
			if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_NET_DEBUG)) {
				nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			}
			if int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))) != 0 {
				sub_467750(v292, int8(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))))
			} else {
				sub_467750(0, int8(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))))
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
		case 225:
			v293 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			sub_467740(v293)
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
		case 226:
			v294 = nox_xxx_netClearHighBit_578B30(int16(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			v295 = v294
			if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_NET_DEBUG)) {
				nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			}
			*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(sub_478080(v295))
			if uint32(int32(v5)) != 0 || (func() uint32 {
				*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(sub_4676D0(v295))
				return uint32(int32(v5))
			}()) != 0 || (func() uint32 {
				if nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))) == 0 {
					*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_xxx_netSpriteByCodeDynamic_45A6F0(v295))))
				} else {
					*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_xxx_netSpriteByCodeStatic_45A720(v295))))
				}
				return uint32(int32(v5))
			}()) != 0 {
				k = int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))
				*(*uint32)(unsafe.Pointer(uintptr(v5 + 432))) = uint32(k)
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
		case 228:
			if nox_client_isConnected_43C700() != 0 {
				if int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))) == 1 {
					if !nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING)) {
						nox_gameDisableMapDraw_5d4594_2650672 = 0
						sub_413A00(1)
						nox_client_screenFadeXxx_44DB30(nox_client_getFadeDuration(), bool2int(int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 2)))) == 1), sub_44E020)
					}
				} else if !nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING)) {
					nox_client_screenFadeTimeout_44DAB0(nox_client_getFadeDuration(), bool2int(int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 2)))) == 1), sub_44E000)
				}
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
		case 229:
			if nox_client_isConnected_43C700() != 0 {
				sub_43D9B0(int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))), int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 2)))))
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
		case 230:
			if nox_client_isConnected_43C700() != 0 {
				sub_43DA80()
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
		case 231:
			if nox_client_isConnected_43C700() != 0 {
				sub_43DAD0()
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
		case 232:
			if nox_client_isConnected_43C700() != 0 {
				*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(noxServer.getPlayerByID(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))))))
				v111 = (*uint32)(unsafe.Pointer(uintptr(v5)))
				if uint32(int32(v5)) != 0 {
					if uint32(int32(v5)) == *memmap.PtrUint32(0x8531A0, 2576) {
						sub_4BFE40()
						sub_478000()
					}
					v112 = nox_xxx_netSpriteByCodeDynamic_45A6F0(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
					if v112 != nil {
						nox_xxx_cliRemoveHealthbar_459E30((*nox_drawable)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v112)))))), 3)
					}
					if !noxflags.HasGame(noxflags.GameModeQuest) {
						*(*uint32)(unsafe.Add(unsafe.Pointer(v111), unsafe.Sizeof(uint32(0))*1)) = 0
						*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(v111), unsafe.Sizeof(uint32(0))*582)))))
						v113 = 27
						for {
							v114 = (*uint32)(unsafe.Pointer(uintptr(v5)))
							*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(int32(v5 + 24))
							*(*uint32)(unsafe.Pointer(uintptr(v5 - 4))) = 0
							v5 = int64(uint32(int32(v5)))
							*v114 = 0
							v113--
							*(*uint32)(unsafe.Add(unsafe.Pointer(v114), unsafe.Sizeof(uint32(0))*1)) = 0
							*(*uint32)(unsafe.Add(unsafe.Pointer(v114), unsafe.Sizeof(uint32(0))*2)) = 0
							*(*uint32)(unsafe.Add(unsafe.Pointer(v114), unsafe.Sizeof(uint32(0))*3)) = 0
							if v113 == 0 {
								break
							}
						}
						for k = 0; k < 624; k += 24 {
							v115 = int32(*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v111))), k))), 2972)))))
							if (v115 & 3085) == 0 {
								*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*1)) = uint32(^v115) & *v111
								*v111 = *(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*1))
								*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v111))), k))), 2972)))) = 0
								*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v111))), k))), 2976)))) = 0
								*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v111))), k))), 2980)))) = 0
								*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v111))), k))), 2984)))) = 0
								*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v111))), k))), 2988)))) = 0
							}
						}
					}
				}
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
		case 233:
			if nox_client_isConnected_43C700() != 0 {
				if noxflags.HasGame(noxflags.GameOnline) {
					k = int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))
					if uint32(uint16(int16(k))) == nox_player_netCode_85319C {
						sub_45A670(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))))
					}
				}
				if int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 8))))&1 != 0 {
					*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(noxServer.getPlayerByID(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))))))
					if uint32(int32(v5)) != 0 {
						nox_xxx_cliPlayerRespawn_417680(int32(v5), int8(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 7)))))
					}
				}
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 9))
		case 234:
			if noxflags.HasGame(noxflags.GameOnline) {
				sub_45A670(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 5))
			} else {
			LABEL_1163:
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 5))
			}
		case 235:
			if nox_client_isConnected_43C700() == 0 {
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 2))
				break
			}
			sub_4610D0(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 2))
		case 236:
			if nox_client_isConnected_43C700() != 0 {
				sub_43C7A0(int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 2))
		case 237:
			if nox_client_isConnected_43C700() != 0 {
				sub_4C1CA0(int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 2))
		case 238:
			if int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))) == 6 {
				sub_48D4B0(int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 2)))))
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
				break
			} else if int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))) == 7 {
				sub_48D4A0()
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 2))
				break
			}
			return 0
		case 239:
			if int32(*memmap.PtrUint8(0x85B3FC, 12254)) == 0 {
				nox_xxx_warriorMaxHealth_587000_312784 = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))
				nox_xxx_warriorMaxMana_587000_312788 = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 5))))
				nox_xxx_warriorMaxStrength_587000_312792 = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 9))))
				nox_xxx_warriorMaxSpeed_587000_312796 = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 13))))
			} else if int32(*memmap.PtrUint8(0x85B3FC, 12254)) == 1 {
				nox_xxx_wizardMaxHealth_587000_312816 = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))
				nox_xxx_wizardMaximumMana_587000_312820 = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 5))))
				nox_xxx_wizardStrength_587000_312824 = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 9))))
				nox_xxx_wizardSpeed_587000_312828 = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 13))))
			} else if int32(*memmap.PtrUint8(0x85B3FC, 12254)) == 2 {
				nox_xxx_conjurerMaxHealth_587000_312800 = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))
				nox_xxx_conjurerMaxMana_587000_312804 = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 5))))
				nox_xxx_conjurerStrength_587000_312808 = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 9))))
				nox_xxx_conjurerSpeed_587000_312812 = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 13))))
			}
			nox_xxx_loadBaseValues_57B200()
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 17))
		case 240:
			switch *((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))) {
			case 0:
				if nox_client_isConnected_43C700() == 1 {
					if noxflags.HasGame(noxflags.GameHost) {
						sub_460380()
						nox_xxx_cliPrepareGameplay1_460E60()
						sub_41CC00((*byte)(memmap.PtrOff(0x85B3FC, 10984)))
						nox_xxx_plrLoad_41A480((*byte)(memmap.PtrOff(0x85B3FC, 10984)))
					}
					nox_xxx_cliShowHideTubes_470AA0(1)
					sub_48D4A0()
					nox_xxx_cliPrepareGameplay2_4721D0()
					sub_4705B0()
				}
				goto LABEL_1085
			case 1:
				if nox_client_isConnected_43C700() == 1 {
					clientPlaySoundSpecial(1008, 100)
					v299 = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2))))))))
					if v299 != nil {
						if !noxflags.HasGame(noxflags.GameHost) {
							*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v299))), unsafe.Sizeof(uint32(0))*1198))) = 1
						}
					}
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
			case 2:
				if nox_client_isConnected_43C700() == 1 {
					sub_49B4B0((*uint16)(unsafe.Pointer(data)))
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 14))
			case 4:
				if nox_client_isConnected_43C700() != 0 {
					v300 = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))))))
					if v300 != nil {
						*(*byte)(unsafe.Add(unsafe.Pointer(v300), 4816)) = byte(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 2))))
					}
					*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*1)) = uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))))
					if uint32(*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint16(0))*2))) == nox_player_netCode_85319C {
						sub_463420(int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 2)))))
					}
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 5))
			case 5:
				fallthrough
			case 6:
				fallthrough
			case 7:
				fallthrough
			case 8:
				fallthrough
			case 9:
				fallthrough
			case 10:
				nox_client_isConnected_43C700()
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
			case 11:
				nox_client_isConnected_43C700()
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 16))
			case 12:
				if nox_client_isConnected_43C700() != 0 {
					nox_xxx_clientQuestWinScreen_450770(int32(uintptr(unsafe.Pointer(data))))
				}
				nox_client_gui_flag_1556112 = 0
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 90))
			case 13:
				if nox_client_isConnected_43C700() != 0 {
					nox_client_showQuestBriefing2_450980(int32(uintptr(unsafe.Pointer(data))), int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 4))))&1)
				}
				nox_client_gui_flag_1556112 = 0
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 69))
			case 14:
				if nox_client_isConnected_43C700() != 0 {
					nox_client_showQuestBriefing_450A30(int32(uintptr(unsafe.Pointer(data))), 1)
				}
				nox_client_gui_flag_1556112 = 0
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 69))
			case 15:
				if nox_client_isConnected_43C700() != 0 {
					*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(nox_xxx_netSpriteByCodeStatic_45A720(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2)))))))))
					if uint32(int32(v5)) != 0 {
						*(*uint8)(unsafe.Pointer(uintptr(v5 + 432))) = 0
					}
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
			case 16:
				if *memmap.PtrUint32(6112660, 1200904) == 0 {
					*memmap.PtrUint32(6112660, 1200904) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("GreenZap"))
				}
				if nox_client_isConnected_43C700() != 0 {
					v301 = int32(uintptr(unsafe.Pointer(nox_xxx_spriteLoadAdd_45A360_drawable(*memmap.PtrInt32(6112660, 1200904), int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 6))))), int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 8)))))))))
					*(*uint8)(unsafe.Pointer(uintptr(v301 + 432))) = 0
					*(*uint32)(unsafe.Pointer(uintptr(v301 + 437))) = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2))))
					*(*uint32)(unsafe.Pointer(uintptr(v301 + 441))) = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 6))))
					*(*uint32)(unsafe.Pointer(uintptr(v301 + 433))) = uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 10)))))
					nox_xxx_spriteTransparentDecay_49B950((*uint32)(unsafe.Pointer(uintptr(v301))), int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 10))))))
				}
				goto LABEL_1130
			case 17:
				if nox_client_isConnected_43C700() != 0 {
					sub_45D320(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2))))))
				}
				goto LABEL_1133
			case 18:
				if nox_client_isConnected_43C700() == 0 {
					goto LABEL_1133
				}
				nox_xxx_clientQuestDisableAbility_45D4A0(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2))))))
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
			case 19:
				if nox_client_isConnected_43C700() == 0 {
					goto LABEL_1133
				}
				sub_45D400(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2))))))
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
			case 20:
				if nox_client_isConnected_43C700() != 1 || sub_470580() != 0 {
				LABEL_1085:
					data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 2))
				} else {
					sub_4705B0()
					data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 2))
				}
			case 21:
				if nox_client_isConnected_43C700() != 0 {
					*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(noxServer.getPlayerByID(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 6)))))))))
					if uint32(int32(v5)) != 0 {
						*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*1)) = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2))))
						*(*uint32)(unsafe.Pointer(uintptr(v5 + 4820))) = *(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*1))
					}
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 8))
			case 22:
				if nox_client_isConnected_43C700() != 0 {
					*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(noxServer.getPlayerByID(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))))))))
					if uint32(int32(v5)) != 0 {
						*(*uint8)(unsafe.Pointer(uintptr(v5 + 4824))) = *((*uint8)(unsafe.Add(unsafe.Pointer(data), 2)))
					}
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 5))
			case 23:
				if nox_client_isConnected_43C700() != 0 {
					*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint32(0))*0)) = uint32(uintptr(unsafe.Pointer(noxServer.getPlayerByID(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))))))))
					if uint32(int32(v5)) != 0 {
						*(*uint8)(unsafe.Pointer(uintptr(v5 + 4825))) = *((*uint8)(unsafe.Add(unsafe.Pointer(data), 2)))
					}
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 5))
			case 24:
				if nox_client_isConnected_43C700() == 1 {
					sub_4BFBB0((*uint32)(unsafe.Pointer(uintptr(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 2)))))))
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
			case 25:
				if *memmap.PtrUint32(6112660, 1200908) == 0 {
					*memmap.PtrUint32(6112660, 1200908) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("GreenSpark"))
					*memmap.PtrUint32(6112660, 1200912) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("FireBoom"))
				}
				v302 = &nox_xxx_spriteLoadAdd_45A360_drawable(*memmap.PtrInt32(6112660, 1200912), int32(*(*int16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2))))), int32(*(*int16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 4)))))).field_0
				if v302 != nil {
					nox_xxx_sprite_45A110_drawable((*nox_drawable)(unsafe.Pointer(v302)))
				}
				if nox_client_isConnected_43C700() != 0 {
					v380.field_0 = int32(*(*int16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2)))))
					v380.field_4 = int32(*(*int16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 4)))))
					sub_49A150(&v380, *memmap.PtrInt32(6112660, 1200908), *((*uint8)(unsafe.Add(unsafe.Pointer(data), 6))))
				}
				goto LABEL_1149
			case 26:
				if nox_client_isConnected_43C700() != 0 {
					nox_xxx_makePointFxCli_499610(*memmap.PtrInt32(6112660, 1200788), 25, 500, 25, int32(*(*int16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2))))), int32(*(*int16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 4))))))
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 6))
			case 28:
				if nox_client_isConnected_43C700() != 0 {
					sub_41D1A0(int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 2)))))
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
			case 29:
				if nox_client_isConnected_43C700() != 0 {
					sub_465DE0(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2))))))
					data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
				} else {
				LABEL_1133:
					data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
				}
			case 30:
				v303 = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))))))
				if nox_client_isConnected_43C700() != 0 && v303 != nil {
					if v303 == *(**byte)(memmap.PtrOff(0x8531A0, 2576)) {
						v348 = int32(uintptr(unsafe.Pointer(nox_xxx_spellTitle_424930(int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 2))))))))
						v304 = strMan.GetStringInFile("plyrspel.c:AwardSpell", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
						nox_swprintf(&v405[0], v304, v348)
					} else {
						v349 = int32(uintptr(unsafe.Pointer(nox_xxx_spellTitle_424930(int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 2))))))))
						v305 = strMan.GetStringInFile("plyrspel.c:AwardSpellToOther", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
						nox_swprintf(&v405[0], v305, (*byte)(unsafe.Add(unsafe.Pointer(v303), 4704)), v349)
					}
					nox_xxx_printCentered_445490(&v405[0])
				}
				goto LABEL_1163
			case 31:
				v306 = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))))))
				if nox_client_isConnected_43C700() == 0 || v306 == nil {
					goto LABEL_1163
				}
				if v306 == *(**byte)(memmap.PtrOff(0x8531A0, 2576)) {
					v350 = nox_xxx_guiCreatureGetName_427240(int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 2)))))
					v307 = strMan.GetStringInFile("PlyrGide.c:AwardGuide", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
					nox_swprintf(&v403[0], v307, v350)
				} else {
					v351 = nox_xxx_guiCreatureGetName_427240(int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 2)))))
					v308 = strMan.GetStringInFile("PlyrGide.c:AwardGuideToOther", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
					nox_swprintf(&v403[0], v308, (*byte)(unsafe.Add(unsafe.Pointer(v306), 4704)), v351)
				}
				nox_xxx_printCentered_445490(&v403[0])
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 5))
			case 32:
				v309 = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))))))
				if nox_client_isConnected_43C700() == 0 || v309 == nil {
					goto LABEL_1163
				}
				if v309 == *(**byte)(memmap.PtrOff(0x8531A0, 2576)) {
					v352 = nox_xxx_abilityGetName_0_425260(int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 2)))))
					v310 = strMan.GetStringInFile("ComAblty.c:AwardAbility", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
					nox_swprintf(&v404[0], v310, v352)
				} else {
					v353 = nox_xxx_abilityGetName_0_425260(int32(*((*uint8)(unsafe.Add(unsafe.Pointer(data), 2)))))
					v311 = strMan.GetStringInFile("ComAblty.c:AwardAbilityToOther", "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
					nox_swprintf(&v404[0], v311, (*byte)(unsafe.Add(unsafe.Pointer(v309), 4704)), v353)
				}
				nox_xxx_printCentered_445490(&v404[0])
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 5))
			case 33:
				if nox_client_isConnected_43C700() != 0 {
					v383[0] = int32(uintptr(memmap.PtrOff(0x587000, 160948)))
					v383[1] = int32(uintptr(memmap.PtrOff(0x587000, 160988)))
					v383[2] = int32(uintptr(memmap.PtrOff(0x587000, 161028)))
					v383[3] = int32(uintptr(memmap.PtrOff(0x587000, 161068)))
					v383[4] = int32(uintptr(memmap.PtrOff(0x587000, 161112)))
					v354 = strMan.GetStringInFile((*byte)(unsafe.Pointer(uintptr(v383[*((*uint8)(unsafe.Add(unsafe.Pointer(data), 51)))]))), "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
					v312 = strMan.GetStringInFile((*byte)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2)))), "C:\\NoxPost\\src\\Client\\Network\\cdecode.c")
					nox_swprintf(&v402[0], v312, v354)
					nox_xxx_printCentered_445490(&v402[0])
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 52))
			default:
				return 0
			}
		default:
			return 0
		}
		noxOnCliPacketDebug(op, old, int32(uintptr(unsafe.Pointer(data))-uintptr(unsafe.Pointer(old))))
	}
	return 1
}
