#include <math.h>
#include <time.h>

#include "common__system__team.h"
#include "server__script__script.h"
#include "server__system__server.h"
#include "server__script__activator.h"

#include "GAME1.h"
#include "GAME1_1.h"
#include "GAME1_2.h"
#include "GAME1_3.h"
#include "GAME2.h"
#include "GAME2_1.h"
#include "GAME2_2.h"
#include "GAME3_2.h"
#include "GAME3_3.h"
#include "GAME4.h"
#include "GAME4_1.h"
#include "GAME4_2.h"
#include "GAME4_3.h"
#include "GAME5.h"
#include "GAME5_2.h"
#include "client__gui__guiquit.h"
#include "client__video__draw_common.h"
#include "common/fs/nox_fs.h"
#include "common__log.h"
#include "common__magic__speltree.h"
#include "common__net_list.h"
#include "common__random.h"
#include "common__strman.h"
#include "common__crypt.h"
#include "common__system__settings.h"
#include "common__telnet__telnetd.h"
#include "input_common.h"
#include "operators.h"
#include "server__network__mapsend.h"
#include "server__network__playback.h"
#include "server__network__sdecode.h"
#include "server__xfer__savegame__savegame.h"

extern uint32_t nox_xxx_respawnAllow_587000_205200;
extern uint32_t dword_5d4594_1569656;
extern uint32_t dword_5d4594_1563064;
extern void* dword_5d4594_1548532;
extern uint32_t dword_5d4594_1548524;
extern uint32_t dword_5d4594_1568028;
extern uint32_t dword_5d4594_600116;
extern uint32_t dword_5d4594_1556144;
extern uint32_t dword_5d4594_2386164;
extern uint32_t dword_5d4594_1569660;
extern uint32_t dword_5d4594_1568300;
extern uint32_t dword_5d4594_1548480;
extern uint32_t dword_5d4594_528260;
extern uint32_t dword_5d4594_1548704;
extern uint32_t dword_5d4594_1563096;
extern uint32_t dword_5d4594_1563052;
extern void* nox_alloc_execAbil_1569644;
extern uint32_t dword_5d4594_825768;
extern void* nox_alloc_springs_2386568;
extern uint32_t dword_5d4594_608316;
extern uint32_t dword_5d4594_1599060;
extern uint32_t dword_5d4594_2386160;
extern uint32_t dword_5d4594_251744;
extern uint32_t dword_5d4594_528252;
extern uint64_t qword_581450_10392;
extern uint64_t qword_581450_10256;
extern uint32_t dword_5d4594_1569756;
extern nox_object_t* nox_xxx_host_player_unit_3843628;
extern uint32_t dword_5d4594_2495920;
extern uint32_t dword_5d4594_1568024;
extern uint32_t dword_5d4594_2388648;
extern uint32_t dword_5d4594_1548476;
extern uint32_t dword_5d4594_1599656;
extern uint32_t dword_5d4594_2650652;

extern unsigned int nox_gameFPS;
extern unsigned int nox_frame_xxx_2598000;
extern nox_net_struct_t* nox_net_struct_arr[NOX_NET_STRUCT_MAX];

uint32_t nox_xxx_resetMapInit_1569652 = 0;

// monster AI?
nox_server_xxx nox_server_xxx_1599716[NOX_SERVER_XXX_SIZE * NOX_SERVER_XXX_SIZE] = {0};

void nullsub_21(void) {}
void nullsub_25(uint32_t a1) {}

//----- (00426060) --------------------------------------------------------
void sub_426060() {
	char* v0;     // eax
	char* i;      // eax
	char* v2;     // eax
	char* j;      // esi
	void* result; // eax
	char* v5;     // edi
	char* v6;     // esi
	int v7;       // [esp-4h] [ebp-4h]

	dword_5d4594_608316 = 0;
	dword_5d4594_600116 = time(0);
	v7 = sub_5545A0();
	v0 = sub_554230();
	sub_4282D0(v0, v7);
	if (!nox_common_gameFlags_check_40A5C0(0x2000) || nox_common_gameFlags_check_40A5C0(4096)) {
		result = (void*)nox_common_gameFlags_check_40A5C0(4096);
		if (result) {
			result = (void*)nox_common_gameFlags_check_40A5C0(4096);
			if (result) {
				v5 = sub_416640();
				v6 = nox_xxx_cliGamedataGet_416590(0);
				*getMemU16Ptr(0x5D4594, 739396) = sub_40A770();
				*getMemU32Ptr(0x5D4594, 739400) = *((uint32_t*)v5 + 10);
				*getMemU32Ptr(0x5D4594, 739404) = sub_4200E0();
				*getMemU8Ptr(0x5D4594, 739412) = (v6[53] & 0xC0) != 0;
				*getMemU32Ptr(0x5D4594, 739408) = 5;
				strncpy((char*)getMemAt(0x5D4594, 739676), v6 + 9, 0xFu);
				*getMemU8Ptr(0x5D4594, 739691) = 0;
				strncpy((char*)getMemAt(0x5D4594, 739420), v6, 8u);
				*getMemU8Ptr(0x5D4594, 739428) = 0;
				sub_4289D0((void**)getMemAt(0x5D4594, 739396));
			}
		}
	} else {
		for (i = nox_common_playerInfoGetFirst_416EA0(); i; i = nox_common_playerInfoGetNext_416EE0((int)i)) {
			*((uint32_t*)i + 1162) = -1;
		}
		v2 = nox_common_playerInfoFromNum_417090(31);
		if (v2) {
			sub_425F10((int)v2);
		}
		for (j = nox_common_playerInfoGetFirst_416EA0(); j; j = nox_common_playerInfoGetNext_416EE0((int)j)) {
			if (j[2064] != 31) {
				sub_425F10((int)j);
			}
		}
		sub_426150();
		sub_428810((int)getMemAt(0x5D4594, 599476), 0);
		*getMemU16Ptr(0x5D4594, 599482) = 0;
	}
}

//----- (00446040) --------------------------------------------------------
int sub_446040() { return dword_5d4594_825768; }

int sub_46DCB0(void) { return sub_46DCC0(); }

//----- (004D0CF0) --------------------------------------------------------
char* nox_xxx_getSomeMapName_4D0CF0() {
	int v0; // eax
	int v1; // eax
	int v2; // edi
	int v4; // esi

	v0 = nox_common_gameFlags_getVal_40A5B0();
	v1 = sub_4D0D50(v0);
	v2 = *getMemU32Ptr(0x5D4594, 1548428 + 4 * v1);
	if (!v2) {
		return 0;
	}
	v4 = *getMemU32Ptr(0x5D4594, 1548452 + 4 * v1);
	if (v4 > v2) {
		*getMemU32Ptr(0x5D4594, 1548452 + 4 * v1) = 0;
		v4 = 0;
	}
	*getMemU32Ptr(0x5D4594, 1548452 + 4 * v1) = (*getMemU32Ptr(0x5D4594, 1548452 + 4 * v1) + 1) % v2;
	return (char*)getMemAt(0x5D4594, 1529228 + 128 * (v4 + 20 * v1 + 5 * v1));
}

//----- (004D10F0) --------------------------------------------------------
void sub_4D10F0(char* a1) {
	int v1;            // esi
	const char* i;     // edi
	int v3;            // edi
	int v4;            // eax
	int v5;            // ebx
	int v6;            // edx
	unsigned char* v7; // ecx

	if (a1) {
		v1 = 0;
		if (dword_5d4594_1548476 > 0) {
			for (i = (const char*)getMemAt(0x5D4594, 1525136); _strcmpi(i, a1); i += 32) {
				if (++v1 >= *(int*)&dword_5d4594_1548476) {
					return;
				}
			}
			v3 = dword_5d4594_1548476;
			v4 = 32 * v1;
			v5 = dword_5d4594_1548480;
			*getMemU32Ptr(0x587000, 191880) = v1;
			*getMemU32Ptr(0x5D4594, 1525156 + v4) = *getMemU32Ptr(0x5D4594, 1525156 + 32 * v1) + 1;
			v6 = 0;
			*getMemU32Ptr(0x5D4594, 1525160 + v4) = v5;
			if (v3 > 0) {
				v7 = getMemAt(0x5D4594, 1525156);
				do {
					if (*((uint32_t*)v7 - 6) == *getMemU32Ptr(0x5D4594, 1525132 + 32 * v1) && v1 != v6) {
						++*(uint32_t*)v7;
						*((uint32_t*)v7 + 1) = v5;
					}
					++v6;
					v7 += 32;
				} while (v6 < v3);
			}
			dword_5d4594_1548480 = v5 + 1;
		}
	}
}

//----- (004D15F0) --------------------------------------------------------
void nox_xxx____setargv_14_4D15F0() { *getMemU32Ptr(0x5D4594, 1548520) = 1; }

//----- (004D2160) --------------------------------------------------------
void sub_4D2160() {
	char* result; // eax
	char* i;      // ebp
	int v2;       // edi
	char* j;      // esi
	int v4;       // eax
	int v5;       // [esp+4h] [ebp-4h]

	v5 = nox_common_playerInfoCount_416F40();
	result = nox_server_teamFirst_418B10();
	for (i = result; result; i = result) {
		if (!*(uint32_t*)(*((uint32_t*)i + 19) + 492) && sub_418BC0((int)i)) {
			do {
				v2 = nox_common_randomInt_415FA0(0, v5 - 1);
				for (j = nox_common_playerInfoGetFirst_416EA0(); j; j = nox_common_playerInfoGetNext_416EE0((int)j)) {
					if (!v2) {
						break;
					}
					--v2;
				}
				v4 = *((uint32_t*)j + 514);
			} while (!v4 || !nox_xxx_teamCompare2_419180(v4 + 48, i[57]));
			sub_4F3400(*((uint32_t*)j + 514), *((uint32_t*)i + 19), 1);
		}
		result = nox_server_teamNext_418B60((int)i);
	}
}

//----- (004D2230) --------------------------------------------------------
void sub_4D2230() {
	short v0; // ax
	short v1; // ax
	short v2; // ax

	if (nox_common_gameFlags_check_40A5C0(49152)) {
		v0 = nox_common_gameFlags_getVal_40A5B0();
		if (!nox_xxx_servGamedataGet_40A020(v0)) {
			v1 = nox_common_gameFlags_getVal_40A5B0();
			if (!sub_40A180(v1)) {
				v2 = nox_common_gameFlags_getVal_40A5B0();
				sub_409FB0_settings(v2, 0xFu);
			}
		}
	} else if (nox_common_gameFlags_check_40A5C0(1024) && !nox_xxx_servGamedataGet_40A020(1024)) {
		sub_409FB0_settings(1024, 0xFu);
	}
}

//----- (004D22B0) --------------------------------------------------------
void sub_4D22B0() {
	char* result; // eax
	int i;        // esi
	int v2;       // eax
	int v3;       // edi
	int v4;       // ecx
	int v5;       // eax
	int v6;       // [esp-8h] [ebp-10h]

	result = nox_common_playerInfoGetFirst_416EA0();
	for (i = (int)result; result; i = (int)result) {
		v2 = *(uint32_t*)(i + 2056);
		if (v2) {
			v3 = *(uint32_t*)(v2 + 748);
			if (*(uint32_t*)(v3 + 280)) {
				nox_xxx_shopCancelSession_510DC0(*(uint32_t**)(v3 + 280));
			}
			*(uint32_t*)(v3 + 280) = 0;
			nox_xxx_playerCmd_51AC30(*(unsigned char*)(i + 2064));
			if (nox_common_gameFlags_check_40A5C0(4096)) {
				v4 = *(uint32_t*)(i + 4676);
				*(uint32_t*)(i + 4676) = 0;
				*(uint32_t*)(i + 4680) = v4;
			}
			if (!nox_xxx_playerSetState_4FA020(*(uint32_t**)(i + 2056), 13) ||
				!nox_common_gameFlags_check_40A5C0(512)) {
				*(uint32_t*)(i + 4700) = 0;
				v6 = ((unsigned int)sub_4CFE00() >> 1) & 1;
				v5 = nox_xxx_gameIsSwitchToSolo_4DB240();
				nox_xxx_playerMakeDefItems_4EF7D0(*(uint32_t*)(i + 2056), v5 == 0, v6);
			}
			if (*(uint8_t*)(i + 3680) & 0x20) {
				nox_xxx_playerLeaveObserver_0_4E6AA0(i);
				nox_xxx_playerCameraUnlock_4E6040(*(uint32_t*)(i + 2056));
			}
			*(uint32_t*)(*(uint32_t*)(i + 2056) + 136) = nox_frame_xxx_2598000;
		}
		result = nox_common_playerInfoGetNext_416EE0(i);
	}
}

//----- (004D2DA0) --------------------------------------------------------
int nox_xxx_updateServer_4D2DA0(long long a1) {
	int result;         // eax
	char* v2;           // eax
	int v3;             // ebx
	char* v4;           // esi
	double v5;          // st7
	int v6;             // esi
	int v7;             // eax
	char* v8;           // eax
	float v9;           // [esp+0h] [ebp-10h]
	unsigned char* v10; // [esp+4h] [ebp-Ch]

	if (dword_5d4594_528252 == 1 && nox_frame_xxx_2598000 == dword_5d4594_528260) {
		nox_xxx_reconAttempt_41E390();
	}
	sub_5096F0();
	result = nox_common_gameFlags_check_40A5C0(8);
	if (result) {
		return result;
	}
	if (nox_common_gameFlags_check_40A5C0(0x2000)) {
		sub_416720();
		if (!nox_common_gameFlags_check_40A5C0(128)) {
			if (sub_409F40(0x2000)) {
				sub_4D7CC0();
			}
		}
	}
	if (!nox_common_gameFlags_check_40A5C0(1024)) {
		goto LABEL_31;
	}
	if (sub_40AA00()) {
		if (!dword_5d4594_1548704) {
			sub_4D2FF0();
		}
	} else {
		dword_5d4594_1548704 = 0;
	}
	if (nox_common_gameFlags_check_40A5C0(0x4000000)) {
		goto LABEL_31;
	}
	if (sub_40A300()) {
		goto LABEL_31;
	}
	v2 = nox_common_playerInfoGetFirst_416EA0();
	if (!v2) {
		goto LABEL_31;
	}
	while (*((int*)v2 + 535) <= 0) {
		v2 = nox_common_playerInfoGetNext_416EE0((int)v2);
		if (!v2) {
			goto LABEL_31;
		}
	}
	if (!sub_40AA00()) {
		goto LABEL_31;
	}
	if (nox_xxx_CheckGameplayFlags_417DA0(4)) {
		v3 = nox_xxx_getTeamCounter_417DD0();
		if (v3 >= sub_40AA40()) {
			goto LABEL_31;
		}
		v4 = nox_server_teamFirst_418B10();
		if (v4) {
			while (nox_xxx_countNonEliminatedPlayersInTeam_40A830((int)v4) != 1) {
				v4 = nox_server_teamNext_418B60((int)v4);
				if (!v4) {
					goto LABEL_31;
				}
			}
			v10 = getMemAt(0x587000, 197200);
			v5 = nox_xxx_gamedataGetFloat_419D40("SuddenDeathCountdown");
			v9 = v5;
			v7 = nox_float2int(v9);
			nox_xxx_servStartCountdown_40A2A0(v7, (const char*)v10);
		}
	} else {
		v6 = sub_40A770();
		if (v6 < sub_40AA40()) {
			v10 = getMemAt(0x587000, 197256);
			v5 = nox_xxx_gamedataGetFloat_419D40("SuddenDeathCountdown");
			v9 = v5;
			v7 = nox_float2int(v9);
			nox_xxx_servStartCountdown_40A2A0(v7, (const char*)v10);
		}
	}
LABEL_31:
	if (sub_40A6B0()) {
		v8 = sub_416640();
		sub_4D9700(*(uint32_t*)(v8 + 66));
		sub_40A6A0(0);
	}
	if ((unsigned long long)(a1 - *getMemU64Ptr(0x5D4594, 1548692)) > 0x1F4) {
		nox_xxx_netReportAllLatency_4D3050();
		*getMemU64Ptr(0x5D4594, 1548692) = a1;
	}
	if (nox_common_gameFlags_check_40A5C0(128) && !sub_40A740() && nox_xxx_getTeamCounter_417DD0() &&
		!nox_xxx_CheckGameplayFlags_417DA0(2)) {
		sub_4183C0();
	}
	result = nox_common_gameFlags_check_40A5C0(4096);
	if (result) {
		sub_4D7150();
		sub_4D71F0();
		nox_server_checkWarpGate_4D7600();
		sub_4DCE00();
		result = sub_4D7A80();
	}
	return result;
}

//----- (004D7140) --------------------------------------------------------
int sub_4D7140(int a1) {
	int result; // eax

	result = a1;
	dword_5d4594_1556144 = a1;
	return result;
}

//----- (004DB160) --------------------------------------------------------
char* sub_4DB160() { return getMemAt(0x5D4594, 1557900); }

//----- (004DBA30) --------------------------------------------------------
void sub_4DBA30(int a1) {
	char* result; // eax
	char* v2;     // ebx
	int v3;       // ebp
	int v4;       // esi
	int v5;       // edi
	int v6;       // eax
	int v7;       // esi
	int v8;       // ebx
	int* v9;      // ebp
	int v10;      // eax
	int v11;      // eax
	int v12;      // eax
	int v13;      // ebx
	int* v14;     // edi
	int v15;      // eax
	int v16;      // ebx
	int v17;      // eax
	int v18;      // esi
	int v19;      // edi
	int i;        // esi
	int v23;      // eax
	char* v24;    // [esp+Ch] [ebp-8h]
	int v25;      // [esp+10h] [ebp-4h]
	int v26;      // [esp+18h] [ebp+4h]

	result = nox_common_playerInfoFromNum_417090(31);
	v2 = result;
	v3 = 0;
	v24 = result;
	if (result) {
		v4 = *((uint32_t*)result + 514);
		v25 = *((uint32_t*)result + 514);
		if (v4) {
			if (!*getMemU32Ptr(0x5D4594, 1563128)) {
				*getMemU32Ptr(0x5D4594, 1563128) = nox_xxx_getNameId_4E3AA0("SaveGameLocation");
				*getMemU32Ptr(0x5D4594, 1563132) = nox_xxx_getNameId_4E3AA0("Glyph");
			}
			if (a1 == 1) {
				v5 = nox_server_getFirstObject_4DA790();
				if (v5) {
					do {
						v26 = nox_server_getNextObject_4DA7A0(v5);
						if (*(unsigned short*)(v5 + 4) == *getMemU32Ptr(0x5D4594, 1563128)) {
							nox_xxx_unitMove_4E7010(*((uint32_t*)v2 + 514), (float2*)(v5 + 56));
							*(uint32_t*)(*((uint32_t*)v2 + 514) + 44) = *(uint32_t*)(v5 + 44);
							*(uint32_t*)(v5 + 44) = 0;
							nox_xxx_delayedDeleteObject_4E5CC0(v5);
						} else {
							v6 = *(uint32_t*)(v5 + 8);
							if (v6 & 2) {
								v7 = *(uint32_t*)(v5 + 748);
								v8 = 0;
								if (*(uint8_t*)(v7 + 1129)) {
									v9 = (int*)(v7 + 1132);
									do {
										v10 = sub_4ECF10(*v9);
										*v9 = v10;
										if (!v10) {
											*(uint8_t*)(v7 + 1129) = 0;
										}
										++v8;
										++v9;
									} while (v8 < *(unsigned char*)(v7 + 1129));
									v3 = 0;
								}
								*(uint32_t*)(v7 + 1196) = sub_4ECF10(*(uint32_t*)(v7 + 1196));
								if ((int)*(uint32_t*)(v5 + 16) >= 0) {
									sub_52BAF0(v5);
								}
								v11 = sub_4ECF10(*(uint32_t*)(v7 + 392));
								if (v11) {
									*(uint32_t*)(v7 + 392) = *(uint32_t*)(v11 + 36);
								} else {
									*(uint32_t*)(v7 + 392) = 0;
								}
								v12 = sub_4ECF10(*(uint32_t*)(v7 + 1200));
								if (v12) {
									*(uint32_t*)(v7 + 1200) = *(uint32_t*)(v12 + 36);
								} else {
									*(uint32_t*)(v7 + 1200) = 0;
								}
								*(uint32_t*)(v7 + 1216) = sub_4ECF10(*(uint32_t*)(v7 + 1216));
								v13 = 0;
								if (*(uint8_t*)(v7 + 2172)) {
									v14 = (int*)(v7 + 2140);
									do {
										v15 = sub_4ECF10(*v14);
										if (v15) {
											*v14 = *(uint32_t*)(v15 + 36);
										} else {
											*(uint8_t*)(v7 + 2172) = 0;
										}
										++v13;
										++v14;
									} while (v13 < *(unsigned char*)(v7 + 2172));
								}
								v2 = v24;
							} else if (v6 & 0x4000) {
								if (*(uint32_t*)(*(uint32_t*)(v5 + 748) + 16)) {
									nox_xxx_unitNeedSync_4E44F0(v5);
								}
							} else if ((v6 & 0x8000) == 0) {
								if ((v6 & 0x80u) != 0 && *(uint32_t*)(*(uint32_t*)(v5 + 748) + 12) !=
															 *(uint32_t*)(*(uint32_t*)(v5 + 748) + 4)) {
									nox_xxx_unitAddToUpdatable_4DA8D0(v5);
								}
							} else {
								v23 = *(uint32_t*)(*(uint32_t*)(v5 + 748) + 4);
								if (v23 && *(uint32_t*)(*(uint32_t*)(v23 + 748) + 16)) {
									nox_xxx_unitNeedSync_4E44F0(v5);
								}
							}
						}
						v5 = v26;
					} while (v26);
				}
				nox_script_activatorResolveObjs_51B0C0();
				sub_516FC0();
				if (dword_5d4594_1563096) {
					v16 = nox_server_getFirstObject_4DA790();
					if (v16) {
						do {
							v3 = nox_server_getNextObject_4DA7A0(v16);
							if ((int)*(uint32_t*)(v16 + 16) >= 0 && nox_xxx_isUnit_4E5B50(v16)) {
								if (*(uint8_t*)(v16 + 8) & 2) {
									v17 = *(uint32_t*)(v16 + 12);
									if (v17 & 0x2000) {
										v18 = nox_xxx_inventoryGetFirst_4E7980(v16);
										if (v18) {
											do {
												v19 = nox_xxx_inventoryGetNext_4E7990(v18);
												if (*(unsigned short*)(v18 + 4) == *getMemU32Ptr(0x5D4594, 1563132)) {
													nox_xxx_delayedDeleteObject_4E5CC0(v18);
												}
												v18 = v19;
											} while (v19);
										}
									}
								}
								nox_xxx_delayedDeleteObject_4E5CC0(v16);
							}
							v16 = v3;
						} while (v3);
					}
					nox_object_t* obj = nox_xxx_getFirstUpdatable2Object_4DA840();
					if (obj != v3) {
						do {
							nox_object_t* v21 = nox_xxx_getNextUpdatable2Object_4DA850(obj);
							if ((int)obj->obj_flags >= 0) {
								if (sub_4E5B80(obj)) {
									nox_xxx_delayedDeleteObject_4E5CC0(obj);
								}
							}
							obj = v21;
						} while (obj != v3);
					}
					v2 = v24;
				}
				v4 = v25;
			}
			for (i = *(uint32_t*)(v4 + 516); i != v3; i = *(uint32_t*)(i + 512)) {
				if (*(uint8_t*)(i + 8) & 2) {
					if (*(uint8_t*)(*(uint32_t*)(i + 748) + 1440) & 0x80) {
						nox_xxx_netReportAcquireCreature_4D91A0((unsigned char)v2[2064], i);
						nox_xxx_netMarkMinimapObject_417190((unsigned char)v2[2064], i, 1);
						continue;
					}
					if ((*(uint8_t*)(i + 12) & 0x80)) {
						nox_xxx_netMonitorCreature_4D9250((unsigned char)v2[2064], i);
						nox_xxx_netMarkMinimapObject_417190((unsigned char)v2[2064], i, 1);
						continue;
					}
				}
			}
			nox_xxx_gameSetSwitchSolo_4DB220(v3);
			dword_5d4594_1563096 = v3;
			nox_ticks_reset_416D40();
		}
	}
}

//----- (004DCC00) --------------------------------------------------------
int sub_4DCC00() { return dword_5d4594_1563064; }

//----- (004DCC80) --------------------------------------------------------
int nox_xxx_mapLoadRequired_4DCC80() { return *getMemU32Ptr(0x5D4594, 1563072); }

//----- (004DEE00) --------------------------------------------------------
uint32_t* nox_xxx_netUseMap_4DEE00(const char* a1, int a2) {
	int v2;            // edx
	uint32_t* result;  // eax
	uint32_t* i;       // esi
	int v5;            // edi
	unsigned char* v6; // eax
	char v7[41];       // [esp+8h] [ebp-2Ch]

	v7[0] = 43;
	v2 = nox_frame_xxx_2598000;
	strcpy(&v7[1], a1);
	*(uint32_t*)&v7[37] = v2;
	*(uint32_t*)&v7[33] = a2;
	result = (uint32_t*)nox_xxx_getFirstPlayerUnit_4DA7C0();
	for (i = result; result; i = result) {
		v5 = i[187];
		nox_netlist_addToMsgListCli_40EBC0(*(unsigned char*)(*(uint32_t*)(v5 + 276) + 2064), 1, v7, 41);
		nox_xxx_netPlayerObjSend_518C30((int)i, i, 0, 0);
		if (!nox_common_gameFlags_check_40A5C0(2) || *(uint8_t*)(*(uint32_t*)(v5 + 276) + 2064) != 31) {
			v6 = nox_netlist_copyPacketList_40ED60(*(unsigned char*)(*(uint32_t*)(v5 + 276) + 2064), 1u, &a1);
			if (v6) {
				nox_xxx_netSendSock_552640(*(unsigned char*)(*(uint32_t*)(v5 + 276) + 2064) + 1, v6, (signed int)a1,
										   NOX_NET_SEND_NO_LOCK | NOX_NET_SEND_FLAG2);
			}
		}
		result = (uint32_t*)nox_xxx_getNextPlayerUnit_4DA7F0((int)i);
	}
	return result;
}

//----- (004E4170) --------------------------------------------------------
void sub_4E4170() {
	float v0; // [esp+0h] [ebp-4h]

	if (!(nox_frame_xxx_2598000 % (unsigned int)(5 * nox_gameFPS))) {
		v0 = sub_4E3CA0();
		sub_4E3D50();
		if (sub_4E3CA0() != v0) {
			sub_4E3DD0();
		}
	}
}

//----- (004E5E20) --------------------------------------------------------
void nox_xxx_unitsUpdateDeletedList_4E5E20() {
	uint32_t* result; // eax
	uint32_t* v1;     // edi
	uint32_t* v2;     // esi

	result = *(uint32_t**)getMemAt(0x5D4594, 1565588);
	v1 = 0;
	if (*getMemU32Ptr(0x5D4594, 1565588)) {
		do {
			v2 = (uint32_t*)result[113];
			if (result[114] == nox_frame_xxx_2598000) {
				result[113] = v1;
				v1 = result;
				nox_xxx_unitRemoveFromUpdatable_4DA920(result);
			} else {
				nox_xxx_unitDeleteFinish_4E5E80(result);
			}
			result = v2;
		} while (v2);
		*getMemU32Ptr(0x5D4594, 1565588) = v1;
	} else {
		*getMemU32Ptr(0x5D4594, 1565588) = 0;
	}
}

//----- (004E76C0) --------------------------------------------------------
int sub_4E76C0() {
	int v0;     // edi
	int result; // eax
	int i;      // esi

	v0 = 0;
	result = nox_server_getFirstObject_4DA790();
	for (i = result; result; i = result) {
		nullsub_25(i);
		v0 ^= sub_4E7700(i);
		result = nox_server_getNextObject_4DA7A0(i);
	}
	return result;
}
// 4E76F0: using guessed type void  nullsub_25(uint32_t);

//----- (004EC720) --------------------------------------------------------
void sub_4EC720() {
	int v1;       // eax
	int v2;       // edi
	int v3;       // ecx
	int v4;       // eax
	uint32_t* v5; // eax
	int v6;       // edi
	short v7;     // ax
	uint8_t* v8;  // eax
	int v9;       // eax
	int v10;      // eax
	double v11;   // st7
	double v12;   // st6
	int v13;      // eax
	int v14;      // ecx
	uint8_t* v15; // eax
	int v16;      // eax
	int v17;      // ecx
	int v18;      // eax

	if (!dword_5d4594_1568028) {
		dword_5d4594_1568028 = nox_xxx_getNameId_4E3AA0("Crown");
	}
	if (nox_common_gameFlags_check_40A5C0(4608)) {
		return;
	}
	nox_xxx_respawnAllow_587000_205200 = 0;
	for (int v0 = dword_5d4594_1568024; v0; v0 = *(uint32_t*)(v0 + 52)) {
		if (!*(uint32_t*)(v0 + 24)) {
			v1 = *(uint32_t*)(v0 + 4);
			v2 = 0;
			if (!v1) {
				*(uint32_t*)(v0 + 24) = 1;
				*(uint32_t*)(v0 + 20) = nox_frame_xxx_2598000 + 30 * nox_gameFPS;
			} else {
				v3 = *(uint32_t*)(v1 + 8);
				if (v3 & 2) {
					v4 = *(uint32_t*)(v1 + 16);
					if (v4 & 0x20) {
						*(uint32_t*)(v0 + 4) = 0;
						*(uint32_t*)(v0 + 24) = 1;
						*(uint32_t*)(v0 + 20) = nox_frame_xxx_2598000 + 30 * nox_gameFPS;
					} else if ((v4 & 0x8000) != 0) {
						*(uint32_t*)(v0 + 24) = 1;
						*(uint32_t*)(v0 + 20) = nox_frame_xxx_2598000 + 30 * nox_gameFPS;
					}
				} else if (*(uint8_t*)(v1 + 16) & 0x20) {
					if (nox_xxx_getUnitDefDd10_4E3BA0(*(unsigned short*)(v1 + 4))) {
						v2 = 1;
					}
					*(uint32_t*)(v0 + 4) = 0;
					if (v2) {
						*(uint32_t*)(v0 + 24) = 1;
						*(uint32_t*)(v0 + 20) = nox_frame_xxx_2598000 + 30 * nox_gameFPS;
					}
				} else if (v3 & 0x3001000 || *(unsigned short*)(v1 + 4) == dword_5d4594_1568028) {
					if (*(uint32_t*)(v1 + 492) || !nox_xxx_getUnitDefDd10_4E3BA0(*(unsigned short*)(v1 + 4))) {
						v18 = *(uint32_t*)(v0 + 4);
						if (*(uint32_t*)(v18 + 492) && nox_xxx_getUnitDefDd10_4E3BA0(*(unsigned short*)(v18 + 4)) &&
							*(unsigned short*)(*(uint32_t*)(v0 + 4) + 4) != dword_5d4594_1568028 && sub_409F40(2)) {
							*(uint32_t*)(v0 + 24) = 1;
							*(uint32_t*)(v0 + 20) = nox_frame_xxx_2598000 + 30 * nox_gameFPS;
						}
					} else {
						v10 = *(uint32_t*)(v0 + 4);
						if (nox_frame_xxx_2598000 > (unsigned int)(5 * nox_gameFPS + *(uint32_t*)(v10 + 128))) {
							v11 = *(float*)(v0 + 8) - *(float*)(v10 + 56);
							v12 = *(float*)(v0 + 12) - *(float*)(v10 + 60);
							if (v12 * v12 + v11 * v11 > 2500.0) {
								nox_xxx_netSendPointFx_522FF0(129, (float2*)(v10 + 56));
								nox_xxx_audCreate_501A30(283, (float2*)(*(uint32_t*)(v0 + 4) + 56), 0, 0);
								nox_xxx_unitMove_4E7010(*(uint32_t*)(v0 + 4), (float2*)(v0 + 8));
								v13 = *(uint32_t*)(v0 + 4);
								v14 = *(uint32_t*)(v13 + 8);
								if (v14 & 0x1000) {
									nox_xxx_rechargeItem_53C520(v13, 100);
								} else if (v14 & 0x1000000 &&
										   nox_xxx_weaponInventoryEquipFlags_415820(*(uint32_t*)(v0 + 4)) & 0x82) {
									v15 = *(uint8_t**)(*(uint32_t*)(v0 + 4) + 736);
									v15[1] = *(uint8_t*)(v0 + 48);
									*v15 = *(uint8_t*)(v0 + 49);
								}
								v16 = *(uint32_t*)(v0 + 4);
								v17 = *(uint32_t*)(v16 + 556);
								if (v17) {
									nox_xxx_unitSetHP_4E4560(v16, *(uint16_t*)(v17 + 4));
								}
								nox_xxx_netSendPointFx_522FF0(129, (float2*)(v0 + 8));
								nox_xxx_audCreate_501A30(283, (float2*)(v0 + 8), 0, 0);
							}
						}
					}
				} else if (*(uint32_t*)(v1 + 492)) {
					*(uint32_t*)(v0 + 24) = 1;
					*(uint32_t*)(v0 + 20) = nox_frame_xxx_2598000 + 30 * nox_gameFPS;
				}
			}
			if (!*(uint32_t*)(v0 + 24)) {
				continue;
			}
		}
		if (!(nox_frame_xxx_2598000 >= *(int*)(v0 + 20) && nox_xxx_getUnitDefDd10_4E3BA0(*(uint32_t*)v0))) {
			continue;
		}
		v5 = nox_xxx_newObjectWithTypeInd_4E3450(*(uint32_t*)v0);
		v6 = (int)v5;
		if (v5) {
			nox_xxx_createAt_4DAA50((int)v5, 0, *(float*)(v0 + 8), *(float*)(v0 + 12));
			nox_xxx_netSendPointFx_522FF0(129, (float2*)(v0 + 8));
			v7 = *(uint16_t*)(v0 + 16);
			*(uint16_t*)(v6 + 124) = v7;
			*(uint16_t*)(v6 + 126) = v7;
			if (*(uint32_t*)(v6 + 8) & 0x13001000) {
				nox_xxx_modifSetItemAttrs_4E4990(v6, (int*)(v0 + 28));
			}
			if (*(uint32_t*)(v6 + 8) & 0x1000000 && nox_xxx_weaponInventoryEquipFlags_415820(v6) & 0x82) {
				v8 = *(uint8_t**)(v6 + 736);
				v8[1] = *(uint8_t*)(v0 + 48);
				*v8 = *(uint8_t*)(v0 + 49);
			}
			nox_xxx_aud_501960(283, v6, 0, 0);
		}
		v9 = *(uint32_t*)(v0 + 4);
		if (v9) {
			if (*(uint8_t*)(v9 + 8) & 2) {
				nox_xxx_delayedDeleteObject_4E5CC0(*(uint32_t*)(v0 + 4));
			}
		}
		*(uint32_t*)(v0 + 24) = 0;
		*(uint32_t*)(v0 + 4) = v6;
	}
}

//----- (004EDD70) --------------------------------------------------------
int sub_4EDD70() {
	int result;   // eax
	int i;        // edi
	uint32_t* v2; // esi
	uint32_t* v3; // ebx
	float2 a3;    // [esp+4h] [ebp-8h]

	result = nox_xxx_getFirstPlayerUnit_4DA7C0();
	for (i = result; result; i = result) {
		v2 = *(uint32_t**)(i + 504);
		if (v2) {
			do {
				v3 = (uint32_t*)v2[124];
				if (v2[2] & 0x10000000) {
					sub_4ED970(50.0, (float2*)(i + 56), &a3);
					nox_xxx_drop_4ED790(i, v2, &a3);
				}
				v2 = v3;
			} while (v3);
		}
		result = nox_xxx_getNextPlayerUnit_4DA7F0(i);
	}
	return result;
}

//----- (004EF660) --------------------------------------------------------
int sub_4EF660(nox_object_t* a1p) {
	int a1 = a1p;
	int v1; // esi
	int i;  // eax

	v1 = *(uint32_t*)(a1 + 748);
	*(uint32_t*)(v1 + 116) = 0;
	*(uint32_t*)(v1 + 120) = 0;
	*(uint32_t*)(v1 + 124) = 0;
	*(uint32_t*)(v1 + 128) = 0;
	*(uint32_t*)(v1 + 308) = 0;
	for (i = 4796; i < 4816; *(uint32_t*)(*(uint32_t*)(v1 + 276) + i - 4) = 0) {
		i += 4;
	}
	if (!nox_common_gameFlags_check_40A5C0(2048)) {
		*(uint8_t*)(v1 + 244) = sub_4EF6F0(a1);
	}
	*(uint32_t*)(v1 + 264) = 0;
	*(uint32_t*)(a1 + 520) = 0;
	return sub_422140(*(uint32_t*)(v1 + 276));
}

//----- (004F1F20) --------------------------------------------------------
void sub_4F1F20() {
	int v0;       // ebp
	int v1;       // eax
	int v2;       // esi
	int v3;       // ebx
	int v4;       // eax
	uint32_t* v5; // edi
	uint32_t* v6; // eax
	uint32_t* v7; // esi
	uint32_t* v8; // edi
	int v9;       // [esp+0h] [ebp-10h]
	int v10;      // [esp+4h] [ebp-Ch]
	float2 a3;    // [esp+8h] [ebp-8h]

	v9 = nox_game_getQuestStage_4E3CC0();
	if (!dword_5d4594_1568300) {
		dword_5d4594_1568300 = nox_xxx_getNameId_4E3AA0("RewardMarker");
		*getMemU32Ptr(0x5D4594, 1568304) = nox_xxx_getNameId_4E3AA0("RewardMarkerPlus");
	}
	sub_4F2110();
	sub_4F2210();
	v0 = nox_server_getFirstObject_4DA790();
	if (v0) {
		do {
			v10 = nox_server_getNextObject_4DA7A0(v0);
			v1 = *(unsigned short*)(v0 + 4);
			if ((unsigned short)v1 == dword_5d4594_1568300 || v1 == *getMemU32Ptr(0x5D4594, 1568304)) {
				if (*(uint8_t*)(*(uint32_t*)(v0 + 692) + 216) & 0x80) {
					v7 = nox_server_rewardgen_activateMarker_4F0720(v0, v9);
					if (v7) {
						nox_xxx_createAt_4DAA50((int)v7, 0, *(float*)(v0 + 56), *(float*)(v0 + 60));
						if (v7[2] & 0x1000000) {
							if (v7[3] & 0xC) {
								v8 = nox_xxx_newObjectByTypeID_4E3810("Quiver");
								if (v8) {
									a3 = *((float2*)v7 + 7);
									sub_4ED970(30.0, (float2*)v7 + 7, &a3);
									nox_xxx_createAt_4DAA50((int)v8, 0, a3.field_0, a3.field_4);
								}
							}
						}
					}
				}
				nox_xxx_delayedDeleteObject_4E5CC0(v0);
			} else if (*(int* (**)(int))(v0 + 688) == nox_xxx_initChest_4F0400) {
				v2 = nox_xxx_inventoryGetFirst_4E7980(v0);
				if (v2) {
					do {
						v3 = nox_xxx_inventoryGetNext_4E7990(v2);
						v4 = *(unsigned short*)(v2 + 4);
						if ((unsigned short)v4 == dword_5d4594_1568300 || v4 == *getMemU32Ptr(0x5D4594, 1568304)) {
							v5 = nox_server_rewardgen_activateMarker_4F0720(v2, v9 + 1);
							sub_4ED0C0(v0, (int*)v2);
							nox_xxx_delayedDeleteObject_4E5CC0(v2);
							if (v5) {
								nox_xxx_inventoryPutImpl_4F3070(v0, (int)v5, 0);
								if (v5[2] & 0x1000000) {
									if (v5[3] & 0xC) {
										v6 = nox_xxx_newObjectByTypeID_4E3810("Quiver");
										if (v6) {
											nox_xxx_inventoryPutImpl_4F3070(v0, (int)v6, 0);
										}
									}
								}
							}
						}
						v2 = v3;
					} while (v3);
				}
			}
			v0 = v10;
		} while (v10);
	}
}

//----- (004FBEE0) --------------------------------------------------------
void nox_xxx_abilUpdateMB_4FBEE0() {
	for (char* p = nox_common_playerInfoGetFirst_416EA0(); p; p = nox_common_playerInfoGetNext_416EE0(p)) {
		if (*((uint32_t*)p + 514) && !p[2251]) {
			for (int i = 0; i < 6; ++i) {
				int v2 = i + 6 * (unsigned char)p[2064];
				int v3 = *getMemU32Ptr(0x5D4594, 1568876 + 4 * v2);
				if (v3) {
					*getMemU32Ptr(0x5D4594, 1568876 + 4 * v2) = v3 - 1;
					if (!*getMemU32Ptr(0x5D4594, 1568876 + 4 * (i + 6 * (unsigned char)p[2064]))) {
						nox_xxx_netAbilRepotState_4D8100(*((uint32_t*)p + 514), i, 1);
					}
				}
			}
		}
	}
	int* next;
	for (int* p = *(int**)getMemAt(0x5D4594, 1569648); p; p = next) {
		next = (int*)p[4];
		if (!(*(uint32_t*)(p[1] + 16) & 0x8020)) {
			if (nox_frame_xxx_2598000 <= (unsigned int)p[2]) {
				continue;
			}
			int v11 = p[1];
			int v8 = sub_425230(*p, 2);
			nox_xxx_aud_501960(v8, v11, 0, 0);
			sub_4FC3C0(p[1], *p, 0);
			if (*p == 1) {
				nox_xxx_playerSetState_4FA020((uint32_t*)p[1], 13);
			}
		}
		int v5 = p[4];
		if (v5) {
			*(uint32_t*)(v5 + 20) = p[5];
		}
		int v7 = p[5];
		if (v7) {
			*(uint32_t*)(v7 + 16) = p[4];
		} else {
			*getMemU32Ptr(0x5D4594, 1569648) = p[4];
		}
		nox_alloc_class_free_obj_first(*(unsigned int**)&nox_alloc_execAbil_1569644, p);
	}
}

//----- (004FC680) --------------------------------------------------------
void sub_4FC680() {
	int v0; // eax

	if (nox_common_gameFlags_check_40A5C0(2048) && !nox_common_gameFlags_check_40A5C0(0x80000) &&
		dword_5d4594_1569660) {
		v0 = nox_xxx_getFirstPlayerUnit_4DA7C0();
		if (v0) {
			nox_xxx_playerExecuteAbil_4FBB70(v0, *(int*)&dword_5d4594_1569660);
			dword_5d4594_1569660 = 0;
		}
	}
}

//----- (004FC6D0) --------------------------------------------------------
void nox_server_xxxInitPlayerUnits_4FC6D0() {
	char* v1;            // eax
	int i;               // edi
	int v3;              // esi
	int v4;              // eax
	int v5;              // ebp
	int j;               // esi
	int v7;              // eax
	char FileName[1024]; // [esp+Ch] [ebp-400h]

	if (nox_xxx_resetMapInit_1569652 != 1 && dword_5d4594_1569656 != 1) {
		return;
	}
	if (!nox_xxx_getFirstPlayerUnit_4DA7C0()) {
		return;
	}
	if (nox_common_gameFlags_check_40A5C0(4096)) {
		if (nox_game_getQuestStage_4E3CC0() == 1) {
			nox_game_sendQuestStage_4D6960(255);
			sub_4D7440(1);
			sub_4D60B0();
		} else if (!sub_4D6F30() || sub_4D7430()) {
			if (sub_4D76F0() == 1) {
				sub_4D6880(255, 1);
				sub_4D76E0(0);
				sub_4D7440(1);
				sub_4D60B0();
			} else {
				v1 = nox_fs_root();
				nox_sprintf(FileName, "%s\\Save\\_temp_.dat", v1);
				for (i = nox_xxx_getFirstPlayerUnit_4DA7C0(); i; i = nox_xxx_getNextPlayerUnit_4DA7F0(i)) {
					v3 = *(uint32_t*)(i + 748);
					v4 = *(uint32_t*)(v3 + 276);
					if (*(uint32_t*)(v4 + 4792) == 1 && !*(uint32_t*)(v3 + 552) &&
						nox_xxx_playerSaveToFile_41A140(FileName, *(unsigned char*)(v4 + 2064))) {
						v5 = sub_419EE0(*(uint8_t*)(*(uint32_t*)(v3 + 276) + 2064));
						nox_xxx_sendGauntlet_4DCF80(*(unsigned char*)(*(uint32_t*)(v3 + 276) + 2064), 1);
						if (!sub_41CFA0(FileName, *(unsigned char*)(*(uint32_t*)(v3 + 276) + 2064)) && !v5) {
							nox_xxx_sendGauntlet_4DCF80(*(unsigned char*)(*(uint32_t*)(v3 + 276) + 2064), 0);
						}
						nox_fs_remove(FileName);
					}
					sub_4D6770(*(unsigned char*)(*(uint32_t*)(v3 + 276) + 2064));
				}
				sub_4D6880(255, 0);
				sub_4D7440(1);
				sub_4D60B0();
			}
		} else {
			nox_game_sendQuestStage_4D6960(255);
			sub_4D7440(1);
			sub_4D60B0();
		}
	} else {
		nox_xxx_netMsgFadeBegin_4D9800(1, 1);
	}
	if (nox_common_gameFlags_check_40A5C0(0x2000)) {
		if (!nox_common_gameFlags_check_40A5C0(128)) {
			for (j = nox_xxx_getFirstPlayerUnit_4DA7C0(); j; j = nox_xxx_getNextPlayerUnit_4DA7F0(j)) {
				v7 = *(uint32_t*)(*(uint32_t*)(j + 748) + 276);
				if (*(uint8_t*)(v7 + 2064) != 31 && !(*(uint8_t*)(v7 + 3680) & 1)) {
					nox_xxx_buffApplyTo_4FF380(j, 23, 0, 5);
				}
			}
		}
	}
}

//----- (005003B0) --------------------------------------------------------
int nox_xxx_playerSomeWallsUpdate_5003B0(nox_object_t* obj) {
	int a1 = obj;
	int v1;            // edi
	int result;        // eax
	int v3;            // ebp
	unsigned char* v4; // esi
	uint8_t* v5;       // eax

	if (*(int*)&dword_5d4594_1569756 <= 0) {
		v1 = a1;
	} else {
		v1 = nox_xxx_spellCastedFirst_4FE930();
		if (!v1) {
			return 0;
		}
		while (*(uint32_t*)(v1 + 4) != 132 || *(uint32_t*)(v1 + 16) != a1 || *(uint8_t*)(v1 + 88) & 1) {
			result = nox_xxx_spellCastedNext_4FE940(v1);
			v1 = result;
			if (!result) {
				return result;
			}
		}
		if (!v1) {
			return 0;
		}
		*(uint32_t*)(v1 + 92) = nox_xxx_spellWallCreate_4FFA90;
		*(uint32_t*)(v1 + 96) = nox_xxx_spellWallUpdate_500070;
		*(uint32_t*)(v1 + 100) = nox_xxx_spellWallDestroy_500080;
		*(uint32_t*)(v1 + 48) = 0;
	}
	v3 = 0;
	if (dword_5d4594_1569756 > 0) {
		v4 = getMemAt(0x5D4594, 1569764);
		do {
			v5 = (uint8_t*)nox_server_getWallAtGrid_410580(*v4, v4[1]);
			if (v5) {
				*v5 = v4[13];
			} else {
				v5 = nox_xxx_wallCreateAt_410250(*v4, v4[1]);
				if (!v5) {
					return 0;
				}
				v5[4] |= 8u;
				v5[1] = v4[11];
				v5[2] = v4[12];
				*v5 = v4[13];
				v5[7] = v4[14];
			}
			nox_xxx_netWallCreate_4FFE80(v1, v5, *((uint32_t*)v4 + 1), v4[8], v4[9], v4[10]);
			++v3;
			v4 += 16;
		} while (v3 < *(int*)&dword_5d4594_1569756);
	}
	dword_5d4594_1569756 = 0;
	return 1;
}

//----- (00500510) --------------------------------------------------------
void sub_500510(const char* a1) {
	if (a1) {
		strcpy((char*)getMemAt(0x5D4594, 1570008), a1);
	}
}

//----- (00502100) --------------------------------------------------------
void sub_502100() {
	nox_alloc_class_free_all(*(uint32_t**)getMemAt(0x5D4594, 1599056));
	dword_5d4594_1599060 = 0;
}

//----- (00506F30) --------------------------------------------------------
void nox_xxx_voteUptate_506F30() {
	uint32_t* result; // eax
	uint32_t* v1;     // esi

	result = *(uint32_t**)&dword_5d4594_1599656;
	if (dword_5d4594_1599656) {
		do {
			v1 = (uint32_t*)result[11];
			switch (*result) {
			case 0:
			case 1:
				sub_506F80((int)result);
				break;
			case 2:
				sub_507090((int)result);
				break;
			case 3:
				sub_507100((int)result);
				break;
			default:
				break;
			}
			result = v1;
		} while (v1);
	}
}

//----- (0050AFA0) --------------------------------------------------------
void sub_50AFA0() {
	int i;              // edi
	int v1;             // eax
	int v2;             // esi
	int v3;             // eax
	int v4;             // esi
	int v5;             // eax
	int v6;             // esi
	int v7;             // eax
	int v8;             // esi
	int v9;             // eax
	int v10;            // eax
	double v11;         // st7
	int v12;            // ebx
	int v13;            // esi
	int v14;            // eax
	int v15;            // ecx
	unsigned char* v16; // ebp
	float* v17;         // esi
	short v18;          // ax
	float* result;      // eax
	float* k;           // edi
	int v21;            // esi
	int v22;            // eax
	float v23;          // [esp+0h] [ebp-3Ch]
	float v24;          // [esp+0h] [ebp-3Ch]
	float v25;          // [esp+0h] [ebp-3Ch]
	float v26;          // [esp+0h] [ebp-3Ch]
	float v27;          // [esp+0h] [ebp-3Ch]
	float v28;          // [esp+0h] [ebp-3Ch]
	float v29;          // [esp+0h] [ebp-3Ch]
	float v30;          // [esp+0h] [ebp-3Ch]
	float v31;          // [esp+0h] [ebp-3Ch]
	float v32;          // [esp+0h] [ebp-3Ch]
	float v33;          // [esp+0h] [ebp-3Ch]
	float v34;          // [esp+0h] [ebp-3Ch]
	float v35;          // [esp+4h] [ebp-38h]
	float v36;          // [esp+4h] [ebp-38h]
	int j;              // [esp+14h] [ebp-28h]
	int v38;            // [esp+18h] [ebp-24h]
	int v39;            // [esp+1Ch] [ebp-20h]
	int v40;            // [esp+20h] [ebp-1Ch]
	float v41;          // [esp+24h] [ebp-18h]
	float v42;          // [esp+28h] [ebp-14h]
	float2 a2;          // [esp+2Ch] [ebp-10h]
	int v44;            // [esp+34h] [ebp-8h]

	dword_5d4594_2386160 = 0;
	memset(nox_server_xxx_1599716, 0, sizeof(nox_server_xxx_1599716));
	dword_5d4594_2386164 = 0;
	for (i = nox_server_getFirstObject_4DA790(); i; i = nox_server_getNextObject_4DA7A0(i)) {
		nox_xxx_aiTestUnitDangerous_50B2C0(i);
		v1 = *(uint32_t*)(i + 8);
		if ((v1 & 0x80u) == 0) {
			if (v1 & 0x800) {
				v23 = *(float*)(i + 56) * 0.043478262;
				v2 = nox_float2int(v23);
				v24 = *(float*)(i + 60) * 0.043478262;
				v3 = nox_float2int(v24);
				nox_server_xxx_1599716[v3 + (v2 << 8)].field_8 |= 0x10;
			} else if (v1 & 0x400) {
				v25 = *(float*)(i + 56) * 0.043478262;
				v4 = nox_float2int(v25);
				v26 = *(float*)(i + 60) * 0.043478262;
				v5 = nox_float2int(v26);
				nox_server_xxx_1599716[v5 + (v4 << 8)].field_8 |= 0x20;
			} else if (v1 & 0x4000) {
				v27 = *(float*)(i + 56) * 0.043478262;
				v6 = nox_float2int(v27);
				v28 = *(float*)(i + 60) * 0.043478262;
				v7 = nox_float2int(v28);
				nox_server_xxx_1599716[v7 + (v6 << 8)].field_8 |= 0x4;
			} else if ((v1 & 0x8000) == 0) {
				if (!(*(uint8_t*)(i + 16) & 0x49)) {
					if (v1 & 0x400000) {
						v31 = *(float*)(i + 232) * 0.043478262;
						v10 = nox_float2int(v31);
						v11 = *(float*)(i + 236) * 0.043478262;
						v12 = v10;
						v44 = v10;
						v32 = v11;
						v13 = nox_float2int(v32);
						v33 = *(float*)(i + 240) * 0.043478262;
						v39 = nox_float2int(v33);
						v34 = *(float*)(i + 244) * 0.043478262;
						v14 = nox_float2int(v34);
						v15 = v13;
						v40 = v14;
						for (j = v13; v15 <= v14; j = v15) {
							v38 = v12;
							if (v12 <= v39) {
								v16 = (uint8_t*)&(nox_server_xxx_1599716[v15 + (v12 << 8)].field_8);
								do {
									v17 = getMemFloatPtr(0x587000, 234108);
									while (1) {
										v41 = (double)v38 * 23.0;
										a2.field_0 = v41 + *(v17 - 1);
										v42 = (double)j * 23.0;
										a2.field_4 = v42 + *v17;
										if (sub_547DB0(i, &a2)) {
											break;
										}
										v17 += 2;
										if ((int)v17 >= (int)getMemAt(0x587000, 234180)) {
											goto LABEL_22;
										}
									}
									v18 = *(uint16_t*)v16;
									LOBYTE(v18) = *(unsigned short*)v16 | 1;
									*(uint16_t*)v16 = v18;
									if (!(*(uint8_t*)(i + 16) & 0x10)) {
										LOBYTE(v18) = v18 | 2;
										*(uint16_t*)v16 = v18;
									}
								LABEL_22:
									++v12;
									v16 += sizeof(nox_server_xxx) * NOX_SERVER_XXX_SIZE;
									v38 = v12;
								} while (v12 <= v39);
								v15 = j;
								v14 = v40;
								v12 = v44;
							}
							++v15;
						}
					}
				}
			} else {
				v29 = *(float*)(i + 56) * 0.043478262;
				v8 = nox_float2int(v29);
				v30 = *(float*)(i + 60) * 0.043478262;
				v9 = nox_float2int(v30);
				nox_server_xxx_1599716[v9 + (v8 << 8)].field_8 |= 0x8;
			}
		}
	}
	result = (float*)nox_xxx_waypointGetList_579860();
	for (k = result; result; k = result) {
		if (sub_579EE0((int)k, 0x80u)) {
			v35 = k[2] * 0.043478262;
			v21 = nox_float2int(v35);
			v36 = k[3] * 0.043478262;
			v22 = nox_float2int(v36);
			nox_server_xxx_1599716[v22 + (v21 << 8)].field_8 |= 0x40;
		}
		result = (float*)nox_xxx_waypointNext_579870((int)k);
	}
}

//----- (0050D890) --------------------------------------------------------
unsigned int sub_50D890() {
	unsigned int result; // eax

	if (!(nox_frame_xxx_2598000 % (unsigned int)(5 * nox_gameFPS))) {
		sub_50D8D0();
	}
	result = nox_frame_xxx_2598000 / 0xFu;
	if (!(nox_frame_xxx_2598000 % 0xFu)) {
		result = sub_50D960();
	}
	return result;
}

//----- (00511590) --------------------------------------------------------
void nox_xxx_addDebugEntry_511590(void* a1p, void* a2p) {
	int a1 = a1p;
	int a2 = a2p;
	int v2;         // ebx
	uint32_t* v3;   // eax
	int v4;         // esi
	double v5;      // st7
	double v6;      // st6
	char v7;        // dl
	long double v8; // st7
	float v9;       // [esp+10h] [ebp+4h]

	v2 = a1;
	if (a1) {
		if (a2) {
			if (!(*(uint8_t*)(a1 + 16) & 0x20) && !(*(uint8_t*)(a2 + 16) & 0x20)) {
				v3 = nox_alloc_class_new_obj_zero(*(uint32_t**)&nox_alloc_springs_2386568);
				v4 = (int)v3;
				if (v3) {
					v3[2] = a1;
					v3[3] = a2;
					v3[4] = 1114636288;
					v9 = nox_xxx_objectGetMass_4E4A70(a1);
					if (nox_xxx_objectGetMass_4E4A70(a2) >= v9) {
						nox_xxx_objectGetMass_4E4A70(a2);
					} else {
						nox_xxx_objectGetMass_4E4A70(v2);
					}
					*(uint32_t*)(v4 + 20) = 0;
					v5 = *(float*)(a2 + 56) - *(float*)(v2 + 56);
					v6 = *(float*)(a2 + 60) - *(float*)(v2 + 60);
					v7 = *(uint8_t*)(v4 + 40) | 1;
					*(uint32_t*)(v4 + 28) = 1132462080;
					*(uint8_t*)(v4 + 40) = v7;
					v8 = sqrt(v6 * v6 + v5 * v5);
					*(float*)(v4 + 32) = v8;
					*(float*)(v4 + 36) = v8;
					*(float*)(v4 + 24) = v8;
					sub_511560(v4);
				}
			}
		}
	}
}

//----- (00519990) --------------------------------------------------------
void nox_xxx_serverLoopSendMap_519990() {
	int v0;            // edi
	unsigned char* v1; // esi

	if (dword_5d4594_2388648) {
		if (!*getMemU32Ptr(0x5D4594, 2388644)) {
			nox_xxx_netMapSendPrepair_519EB0_net_mapsend();
		}
		v0 = 0;
		v1 = getMemAt(0x5D4594, 2387180);
		do {
			if (*((uint16_t*)v1 - 15) == 1) {
				if (sub_519930(v0 + 1) > *(v1 - 4)) {
					++*(uint32_t*)v1;
				} else {
					nox_xxx_netSendMap_5199F0_net_mapsend(v1 - 32);
				}
			}
			v1 += 48;
			++v0;
		} while ((int)v1 < (int)getMemAt(0x5D4594, 2388668));
	}
}

//----- (0051A1F0) --------------------------------------------------------
void sub_51A1F0(int a1) {
	int v1;           // ebx
	int v2;           // edi
	int v3;           // ebx
	int v4;           // eax
	int v5;           // ecx
	int v6;           // esi
	double v7;        // st7
	int v8;           // eax
	int v9;           // ebx
	int v10;          // ebp
	int v11;          // esi
	int v12;          // edi
	int v13;          // eax
	int v14;          // ebx
	int v15;          // edi
	int i;            // esi
	int v17;          // esi
	int v18;          // edi
	float v19;        // [esp+0h] [ebp-20h]
	float v20;        // [esp+0h] [ebp-20h]
	unsigned int v21; // [esp+10h] [ebp-10h]
	int v22;          // [esp+14h] [ebp-Ch]
	int v23;          // [esp+18h] [ebp-8h]
	unsigned int v24; // [esp+1Ch] [ebp-4h]
	int v25;          // [esp+24h] [ebp+4h]

	v21 = 0;
	v23 = 0;
	v1 = nox_game_getQuestStage_4E3CC0();
	v22 = v1;
	v19 = nox_xxx_gamedataGetFloat_419D40("QuestHardcoreStage");
	v24 = nox_float2int(v19);
	if (!*getMemU32Ptr(0x5D4594, 2388668)) {
		*getMemU32Ptr(0x5D4594, 2388668) = nox_xxx_getNameId_4E3AA0("HecubahMarker");
		*getMemU32Ptr(0x5D4594, 2388672) = nox_xxx_getNameId_4E3AA0("NecromancerMarker");
	}
	v2 = nox_server_getFirstObject_4DA790();
	if (v2) {
		v3 = a1;
		do {
			v4 = nox_server_getNextObject_4DA7A0(v2);
			v5 = *(uint32_t*)(v2 + 8);
			v25 = v4;
			if (v5 & 0x20 && *(uint8_t*)(v2 + 12) & 1) {
				++v21;
			} else if (*(unsigned short*)(v2 + 4) == *getMemU32Ptr(0x5D4594, 2388668)) {
				++v23;
			}
			if (v5 & 0x20000) {
				v6 = *(uint32_t*)(v2 + 748);
				if (*(uint32_t*)(16 * v3 + v6)) {
					switch (*(unsigned char*)(v6 + v3 + 83)) {
					case 0u:
						v7 = nox_xxx_gamedataGetFloat_419D40("GeneratorMaxActiveCreaturesHigh");
						*(uint8_t*)(v6 + 87) = (long long)v7;
						break;
					case 1u:
						v7 = nox_xxx_gamedataGetFloat_419D40("GeneratorMaxActiveCreaturesNormal");
						*(uint8_t*)(v6 + 87) = (long long)v7;
						break;
					case 2u:
						v7 = nox_xxx_gamedataGetFloat_419D40("GeneratorMaxActiveCreaturesLow");
						*(uint8_t*)(v6 + 87) = (long long)v7;
						break;
					case 3u:
						v7 = nox_xxx_gamedataGetFloat_419D40("GeneratorMaxActiveCreaturesSingular");
						*(uint8_t*)(v6 + 87) = (long long)v7;
						break;
					default:
						break;
					}
					if (nox_game_getQuestStage_4E3CC0() >= v24 && *(uint8_t*)(v6 + v3 + 83) != 3) {
						*(uint8_t*)(v6 + 87) *= 2;
					}
					v8 = sub_51A500(*(uint32_t*)(16 * v3 + v6));
					if (v8) {
						*(uint16_t*)(v2 + 4) = v8;
					}
				} else {
					nox_xxx_delayedDeleteObject_4E5CC0(v2);
				}
			}
			v2 = v25;
		} while (v25);
		if (v21 > 1) {
			v9 = nox_common_randomInt_415FA0(0, v21 - 1);
			v10 = 0;
			v11 = nox_server_getFirstObject_4DA790();
			if (v11) {
				do {
					v12 = nox_server_getNextObject_4DA7A0(v11);
					if (*(uint8_t*)(v11 + 8) & 0x20 && *(uint8_t*)(v11 + 12) & 1) {
						if (v10 != v9) {
							nox_xxx_delayedDeleteObject_4E5CC0(v11);
						}
						++v10;
					}
					v11 = v12;
				} while (v12);
			}
		}
		v1 = v22;
	}
	sub_51A940(0);
	if (v1 >= 5) {
		v20 = nox_xxx_gamedataGetFloat_419D40("MinionsAlwaysStage");
		v13 = nox_float2int(v20);
		if (v1 == 5 || v1 >= v13 || v1 & 1 && nox_common_randomInt_415FA0(1, 100) >= 50) {
			sub_51A940(1);
			if (v23) {
				v14 = nox_common_randomInt_415FA0(1, v23);
				v15 = 0;
				for (i = nox_server_getFirstObject_4DA790(); i; i = nox_server_getNextObject_4DA7A0(i)) {
					if (*(unsigned short*)(i + 4) == *getMemU32Ptr(0x5D4594, 2388668) && ++v15 == v14) {
						nox_xxx_spawnHecubahQuest_51A5A0((int*)(i + 56));
					}
					if (*(unsigned short*)(i + 4) == *getMemU32Ptr(0x5D4594, 2388672) &&
						nox_common_randomInt_415FA0(1, 100) >= 50) {
						nox_xxx_spawnNecroQuest_51A7A0((int*)(i + 56));
					}
				}
			}
		}
	}
	v17 = nox_server_getFirstObject_4DA790();
	if (v17) {
		do {
			v18 = nox_server_getNextObject_4DA7A0(v17);
			if (*(unsigned short*)(v17 + 4) == *getMemU32Ptr(0x5D4594, 2388668)) {
				nox_xxx_delayedDeleteObject_4E5CC0(v17);
			}
			if (*(unsigned short*)(v17 + 4) == *getMemU32Ptr(0x5D4594, 2388672)) {
				nox_xxx_delayedDeleteObject_4E5CC0(v17);
			}
			v17 = v18;
		} while (v18);
	}
}

//----- (0051A920) --------------------------------------------------------
int sub_51A920(int a1) {
	int result; // eax

	result = a1;
	*getMemU32Ptr(0x5D4594, 2388660) = a1;
	return result;
}

void nox_xxx_updateUnits_51B100_B() {
	int v2 = nox_xxx_getFirstUpdatableObject_4DA8A0();
	if (v2) {
		int v3 = 0;
		do {
			v3 = nox_xxx_getNextUpdatableObject_4DA8B0(v2);
			if (!(*(uint8_t*)(v2 + 16) & 0x20)) {
				int v4 = *(uint32_t*)(v2 + 520);
				if (v4 && *(uint8_t*)(v4 + 16) & 0x20) {
					int v5 = *(uint32_t*)(v4 + 508);
					if (!v5 || *(uint8_t*)(v5 + 16) & 0x20) {
						*(uint32_t*)(v2 + 520) = 0;
					} else {
						*(uint32_t*)(v2 + 520) = v5;
					}
				}
				void (*v6)(int);
				v6 = *(void (**)(int))(v2 + 744);
				if (v6) {
					if (*(uint8_t*)(v2 + 8) & 4) {
						int v7 = *(uint32_t*)(v2 + 748);
						if (!(*(uint8_t*)(v2 + 16) & 2)) {
							v6(v2); // update function for the player
						}
						nox_xxx_playerCmd_51AC30(*(unsigned char*)(*(uint32_t*)(v7 + 276) + 2064));
					} else if (!(*(uint8_t*)(v2 + 16) & 2)) {
						v6(v2);
					}
				}
				if (*(uint8_t*)(v2 + 8) & 2) {
					nox_xxx_unitNeedSync_4E44F0(v2);
				}
				nox_xxx_updateFallLogic_51B870(v2);
				uint16_t* v8 = *(uint16_t**)(v2 + 556);
				if (v8) {
					v8[1] = *v8;
				}
				double v9 = *(float*)(v2 + 64) - *(float*)(v2 + 56);
				*(uint32_t*)(v2 + 16) &= 0xFFEFFFFF;
				float v58 = v9;
				float v61 = sub_419A10(v58);
				float v59 = *(float*)(v2 + 68) - *(float*)(v2 + 60);
				float v62 = sub_419A10(v59);
				// Checks if fields are within 0.01d of...something?
				if (*(float*)(v2 + 88) <= *(double*)&qword_581450_10392 ||
					*(float*)(v2 + 88) >= *(double*)&qword_581450_10256 ||
					*(float*)(v2 + 92) <= *(double*)&qword_581450_10392 ||
					*(float*)(v2 + 92) >= *(double*)&qword_581450_10256 ||
					*(float*)(v2 + 80) <= *(double*)&qword_581450_10392 ||
					*(float*)(v2 + 80) >= *(double*)&qword_581450_10256 ||
					*(float*)(v2 + 84) <= *(double*)&qword_581450_10392 ||
					*(float*)(v2 + 84) >= *(double*)&qword_581450_10256 || v61 <= *(double*)&qword_581450_10392 ||
					v61 >= *(double*)&qword_581450_10256 || v62 <= *(double*)&qword_581450_10392 ||
					v62 >= *(double*)&qword_581450_10256 || *(float*)(v2 + 104) <= *(double*)&qword_581450_10392 ||
					*(float*)(v2 + 104) >= *(double*)&qword_581450_10256 ||
					*(float*)(v2 + 108) <= *(double*)&qword_581450_10392 ||
					*(float*)(v2 + 108) >= *(double*)&qword_581450_10256) {
					nox_xxx_unitHasCollideOrUpdateFn_537610(v2);
					*(uint32_t*)(v2 + 16) &= 0xF7FFFFFF;
				} else {
					*(uint32_t*)(v2 + 88) = 0;
					*(uint32_t*)(v2 + 92) = 0;
					*(uint32_t*)(v2 + 80) = 0;
					*(uint32_t*)(v2 + 84) = 0;
					nox_xxx_unitRaise_4E46F0(v2, 0.0);
					int v10 = *(uint32_t*)(v2 + 56);
					*(uint32_t*)(v2 + 68) = *(uint32_t*)(v2 + 60);
					int v11 = *(uint32_t*)(v2 + 16);
					*(uint32_t*)(v2 + 64) = v10;
					int v12 = *(uint32_t*)(v2 + 744);
					*(uint32_t*)(v2 + 108) = 0;
					*(uint32_t*)(v2 + 16) = v11 | 0x8000000;
					if (!v12) {
						nox_xxx_unitRemoveFromUpdatable_4DA920((uint32_t*)v2);
					}
				}
			}
			v2 = v3;
		} while (v3);
	}
}

void nox_xxx_updateUnits_51B100_C() {
	for (int k = nox_xxx_getFirstUpdatableObject_4DA8A0(); k; k = nox_xxx_getNextUpdatableObject_4DA8B0(k)) {
		int v20 = *(uint32_t*)(k + 64);
		int v21 = *(uint32_t*)(k + 60);
		*(uint32_t*)(k + 72) = *(uint32_t*)(k + 56);
		int v22 = *(uint32_t*)(k + 68);
		*(uint32_t*)(k + 56) = v20; // updates unit coords
		LOBYTE(v20) = *(uint8_t*)(k + 541);
		*(uint32_t*)(k + 76) = v21;
		LOWORD(v21) = *(uint16_t*)(k + 126);
		*(uint32_t*)(k + 60) = v22;
		*(uint32_t*)(k + 88) = 0;
		*(uint32_t*)(k + 92) = 0;
		*(uint16_t*)(k + 124) = v21;
		if ((unsigned char)v20 > 4u) {
			*(uint8_t*)(k + 541) = 4;
		}
		*(float*)(k + 544) = (*(float*)(k + 552) + *(float*)(k + 548)) * (5.0 - (double)*(unsigned char*)(k + 541)) * 0.2;
		if (nox_xxx_testUnitBuffs_4FF350(k, 4)) {
			*(float*)(k + 544) = *(float*)(k + 544) * 0.5; // switch back from running
		}
		if (*(uint8_t*)(k + 541) || *(uint8_t*)(k + 540)) {
			short v23 = *(uint16_t*)(k + 542);
			if (v23 > 0) {
				short v24 = v23 - 1;
				*(uint16_t*)(k + 542) = v24;
				if (!v24) {
					char v25 = *(uint8_t*)(k + 541);
					if (v25) {
						*(uint8_t*)(k + 541) = v25 - 1;
					}
					if (*(uint8_t*)(k + 540)) {
						int v26 = *(uint32_t*)(k + 16);
						if ((v26 & 0x8000) == 0) {
							nox_xxx_updatePoison_4EE8F0(k, 1);
						}
					}
					*(uint16_t*)(k + 542) = 1000;
				}
			}
		}
		nox_xxx_updateUnitBuffs_4FF620(k);
		if (*(uint8_t*)(k + 540)) {
			uint16_t* v27 = *(uint16_t**)(k + 556);
			if (v27) {
				if (v27[2] > 0u && *v27 > 0u) {
					int v28 = nox_common_gameFlags_check_40A5C0(4096);
					int v29 = *(uint32_t*)(*(uint32_t*)(k + 556) + 16);
					int v30 = v28 + 1;
					if (!v29 || (unsigned int)(nox_frame_xxx_2598000 - v29) > 0x3C) {
						unsigned char v31 = *(uint8_t*)(k + 540);
						if (v31 > 8u || !(nox_frame_xxx_2598000 % (unsigned int)(128 >> (v31 - 1)))) {
							(*(void (**)(int, uint32_t, uint32_t, int, int))(k + 716))(k, 0, 0, v30, 5);
						}
					}
				}
			}
		}
	}
}

void nox_xxx_updateUnits_51B100_D() {
	uint8_t* v32 = nox_xxx_wallSecretGetFirstWall_410780();
	if (v32) {
		char v33 = 0;
		for (; v32; v32 = (uint8_t*)nox_xxx_wallSecretNext_410790((int*)v32)) {
			switch (v32[21]) {
			case 1:;
				char v47 = v32[20];
				if (!((v47 & 4) && ((v47 & 8) == 8))) {
					v33 = 0;
					break;
				}
				int v48 = *((uint32_t*)v32 + 6) - 1;
				*((uint32_t*)v32 + 6) = v48;
				if (v48) {
					v33 = 0;
					break;
				}
				int v49 = *((uint32_t*)v32 + 1);
				v32[21] = 4;
				int v50 = 23 * v49;
				int v51 = *((uint32_t*)v32 + 2);
				double v52 = (double)(v50 + 11);
				int v53 = *((uint32_t*)v32 + 3);
				float2 v64;
				v64.field_0 = v52;
				v64.field_4 = (double)(23 * v51 + 11);
				char* v54 = nox_xxx_wallFindOpenSound_410EE0(*(unsigned char*)(v53 + 1));
				int v55 = nox_xxx_utilFindSound_40AF50(v54);
				nox_xxx_audCreate_501A30(v55, &v64, 0, 0);
				v33 = 0;
				break;
			case 2:;
				char v36 = v32[22] - 1;
				v32[22] = v36;
				if (!v36) {
					int v37 = *((uint32_t*)v32 + 4);
					v32[21] = 1;
					*((uint32_t*)v32 + 6) = nox_gameFPS * v37;
				}
				v33 = 1;
				break;
			case 3:;
				char v38 = v32[20];
				int v39 = 0;
				if (!(!(v38 & 4) || v38 & 8 || (v39 = *((uint32_t*)v32 + 6) - 1, (*((uint32_t*)v32 + 6) = v39) != 0))) {
					int v40 = *((uint32_t*)v32 + 1);
					v32[21] = 2;
					int v41 = 23 * v40;
					int v42 = *((uint32_t*)v32 + 2);
					double v43 = (double)(v41 + 11);
					int v44 = *((uint32_t*)v32 + 3);
					float2 v63;
					v63.field_0 = v43;
					v63.field_4 = (double)(23 * v42 + 11);
					char* v45 = nox_xxx_wallFindCloseSound_410F20(*(unsigned char*)(v44 + 1));
					int v46 = nox_xxx_utilFindSound_40AF50(v45);
					nox_xxx_audCreate_501A30(v46, &v63, 0, 0);
				}
				v33 = 0;
				break;
			case 4:;
				char v34 = v32[22] + 1;
				v32[22] = v34;
				if (v34 == 23) {
					int v35 = *((uint32_t*)v32 + 4);
					v32[21] = 3;
					*((uint32_t*)v32 + 6) = nox_gameFPS * v35;
				}
				v33 = 1;
				break;
			default:
				break;
			}
			if (v33) {
				double v56 = (double)(int)(23 * *((uint32_t*)v32 + 1)) + 11.5;
				double v65 = (double)(int)(23 * *((uint32_t*)v32 + 2)) + 11.5;
				float4 v66;
				v66.field_0 = v56 - 42.5;
				v66.field_4 = v65 - 42.5;
				v66.field_8 = v56 + 42.5;
				v66.field_C = v65 + 42.5;
				nox_xxx_getUnitsInRect_517C10(&v66, sub_51B860, 0);
			}
		}
	}
}

//----- (005524C0) --------------------------------------------------------
void sub_5524C0() {
	dword_5d4594_2495920 = nox_platform_get_ticks();
	for (int i = 0; i < NOX_NET_STRUCT_MAX; i++) {
		nox_net_struct_t* ns = nox_net_struct_arr[i];
		if (ns && ns->field_38 == 1) {
			if (ns->field_40 + 300 < nox_frame_xxx_2598000) {
				nox_xxx_netStructReadPackets_5545B0(i);
			}
		}
	}
}

//----- (0057B140) --------------------------------------------------------
bool sub_57B140() {
	unsigned long long v0; // rdi
	int result;            // eax

	result = 0;
	if (*getMemU64Ptr(0x5D4594, 2523796)) {
		v0 = nox_xxx___Getcvt_57B180() + 5000;
		if (v0 < nox_platform_get_ticks()) {
			result = 1;
		}
	}
	return result;
}
