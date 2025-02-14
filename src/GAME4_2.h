#ifndef NOX_PORT_GAME4_2
#define NOX_PORT_GAME4_2

#include "defs.h"

int sub_51DA70(int a1, int a2, int a3, int a4, int a5);
void sub_51DD50(int a1, int a2, int a3, int a4);
int sub_51DE30(uint32_t* a1, uint32_t* a2, uint32_t* a3);
int nox_xxx_mapCountWallsMB_51DEA0(int a1);
int sub_51DED0();
int nox_xxx_xfer_saveObj_51DF90(nox_object_t* a1);
int nox_xxx_mapSaveMap_51E010(char* a1, int a2);
int nox_xxx_map_51E140();
int nox_xxx_mapGenSpellIdByName_51E1D0(const char* a1);
int nox_xxx_mapGenReadTheme_51E260(int* a1, int a2);
int nox_xxx_mapGenReadLine_51E540(FILE* a1, uint8_t* a2);
int sub_51E570(FILE* a1, uint8_t* a2);
int sub_51E630(FILE* a1);
int sub_51E670(FILE* a1);
int sub_51E720(FILE* a1);
int sub_51E780(FILE* a1);
int sub_51E800(int a1, uint32_t* a2);
int sub_51EAF0(int a1, uint32_t* a2);
int nox_xxx_genReadAlgData_51EBB0(int a1, FILE* a2);
int nox_xxx_genReadSpellSet_51EFB0(int a1, FILE* a2);
int nox_xxx_genReadWeaponSet_51F030(int a1, FILE* a2);
void nox_xxx_mapGenFreeStr_51F1F0(void* lpMem);
int sub_51F230(int a1, FILE* a2);
int nox_xxx_genReadArmorSet_51F640(int a1, FILE* a2);
int nox_xxx_genReadExit_51F800(int a1, FILE* a2);
int nox_xxx_genReadDecor_51F9F0(uint32_t* a1, FILE* a2);
char* nox_xxx_genDecorReadWallFloor_51FE00(int a1, FILE* a2);
int sub_51FEC0(int a1, int a2, FILE* a3);
int nox_xxx_genDecorReadDecorSet_51FFA0(int a1, FILE* a2);
uint32_t* nox_xxx_gen_520380(FILE* a1);
uint32_t* nox_xxx_gen_5205B0(FILE* a1);
int nox_xxx_genDecorReadCopy_520660(uint32_t* a1, const char* a2, FILE* a3);
int nox_xxx_genDecorReadOccurConstraint_520810(int a1, FILE* a2);
int nox_xxx_genDecorReadOccurLimit_5208D0(int a1, FILE* a2);
int nox_xxx_genDecorReadFrequency_520910(int a1, FILE* a2);
int nox_xxx_genDecorReadRoomSizeCon_5209F0(int a1, FILE* a2);
int nox_xxx_genDecorReadDoor_520A90(int a1, FILE* a2);
int nox_xxx_genDecorReadDoubleDoor_520AB0(int a1, FILE* a2);
int nox_xxx_mapgenCheckSettings_520AD0(int* a1);
int nox_xxx_genReadPrefab_520BF0(int a1, FILE* a2);
char* sub_520CE0(int a1, FILE* a2);
uint32_t* sub_520D50(uint32_t* a1);
long long nox_xxx_mapGenRoundFloatToPtr_520DF0(float2* a1, uint32_t* a2);
int sub_520E60(int2* a1);
int sub_520EA0(int a1);
void sub_520F80();
int sub_521100(int a1);
int sub_521180(int a1);
int sub_521200(int a1);
int sub_521290(int2* a1);
int sub_5212B0(int a1, uint32_t* a2);
int nox_xxx_mapgenAllocBuffer_5213E0();
void nox_xxx_mapgenFreeBuffer_521400();
void* nox_xxx_mapGenGetTopRoom_521710();
int sub_521720(int a1);
int nox_xxx_mapGenAddNewRoom_521730(uint32_t* a1);
int sub_521760(int a1);
int sub_5217A0(int a1, int a2);
int sub_521820(int a1, int a2);
int nox_xxx_mapGenUpdateRoomRect_521850(int a1);
int nox_xxx_mapGenSetRoomPos_521880(uint32_t* a1, float2* a2);
int sub_5218B0(int a1, int a2);
int sub_521900(int a1, int a2, int a3);
float* nox_xxx_mapGenMakeRoomStruct_521940(int a1, int a2);
float* nox_xxx_mapGenPrepareRoom_521990(int a1);
void sub_521A10(void* lpMem);
uint32_t* nox_xxx_mapGenFreeTopRoom_521A40();
int sub_521A70(int a1, int a2, int a3);
int sub_521AA0(uint32_t* a1, int a2);
double sub_521B00(int a1, int a2);
double sub_521B30(int a1, int a2);
double sub_521B60(int a1, int a2);
double sub_521B90(int a1, int a2);
float* sub_521BC0(int a1, float2* a2, float a3, float a4);
uint32_t* sub_521C10(int a1);
int nox_xxx_mapgen_521C60(int a1, int a2);
int sub_521CB0(int a1, int a2, int a3, int a4);
int sub_521EB0(float* a1, float* a2);
int sub_521F10(int a1, float* a2);
void nox_xxx_mapgen_521FE0(int a1, int a2, uint32_t* a3);
uint32_t* nox_xxx_mapGenMakeSpellbook_5220E0(int a1, const char* a2);
uint32_t* nox_xxx_mapGenMakeEnchantedItem_5221A0(char* a1, char* a2, int a3);
uint32_t* sub_522300(int a1, uint32_t* a2);
void nox_xxx_mapgen_522340(int a1, int a2);
void nox_xxx_mapgen_522370(int a1, int a2, int* a3);
float* nox_xxx_mapgen_5224B0(int a1, int a2, int a3);
int sub_5226D0(int a1, float a2, int a3);
int sub_5227B0(int a1, float* a2);
int nox_xxx_mapGenMakeMonsterInRoom_522810(float* a1, char* a2);
float* sub_522C80(float* a1);
char sub_522CA0(int a1, float* a2);
float* sub_522D30(int a1);
unsigned char* nox_xxx_mapGenTryNextRoom_522F40(uint32_t* a1);
int nox_xxx_netSendPointFx_522FF0(char a1, float2* a2);
int nox_xxx_netSendFxAllCli_523030(float2* a1, const void* a2, int a3);
int sub_523150(char a1, char a2, float* a3);
int nox_xxx_netSparkExplosionFx_5231B0(float* a1, char a2);
void nox_xxx_sendGeneratorBreakFX_523200(float* a1, char a2);
int nox_xxx_netSendVampFx_523270(char a1, short* a2, short a3);
int nox_xxx_servCode_523340(int* a1, const void* a2, int a3);
int nox_xxx_netClientPredictLinear_523530(int a1);
void nox_xxx_mapGenSetFlags_5235F0(char a1);
int nox_xxx_netSendShieldFx_523670(int a1, float* a2);
int nox_xxx_sendSummonStartFX_5236F0(short a1, float* a2, char a3, short a4, short a5);
int nox_xxx_sendSummonCancelFX_523760(short a1);
int nox_xxx_netSendFxGreenBolt_523790(int4* a1, short a2);
void nox_xxx_sendGeneratorSpawnFX_523830(int4* a1, short a2);
void nox_xxx_sendArrowTrapFX_5238A0(float* a1, char a2);
int nox_xxx_mapGenCheckRoomType_5238F0(int* a1);
int sub_523920(int a1);
int sub_523960(int a1);
int sub_523970(int a1);
int sub_5239B0(int a1);
int sub_523A10(int a1, float* a2);
int sub_523C30(int a1, int a2);
int sub_523CB0(int a1, int a2);
float* sub_523D30(float* a1, float* a2);
float* sub_523E30(int a1, int a2, int a3);
float* nox_xxx_mapGenMakeHall_523EC0(int a1, int a2, int a3);
int sub_524070(int a1, int a2);
int sub_524090(int a1, int* a2);
int nox_xxx_mapGenDecorChkConstaint_5241C0(int a1, int a2);
int nox_xxx_mapGenChkDecorFillsRoom_5241F0(int a1, int a2);
char nox_xxx_mapGenDecorChkLimit_524220(int* a1, int a2);
int nox_xxx_mapGenMakeRooms_524310(int a1);
int sub_5244D0(int a1);
int sub_524500(float2* a1, int a2);
int sub_524550(int* a1, int a2);
float* sub_5245A0(int a1, float* a2, int a3, int a4);
float* sub_524610(int a1, float* a2, int a3);
int sub_524660(float a1, float a2);
int nox_xxx_gen_524680(int a1, int a2, int a3);
int sub_524950(int a1, int a2, float* a3, int* a4);
int sub_5249C0(int a1, int a2, float* a3, int* a4);
void sub_524B50(int a1, int a2, float* a3, int* a4);
void nox_xxx_gen_524E00(int a1, int a2);
int sub_524FB0(int a1, int a2, int a3);
float2* sub_525330(float2* a1, int a2);
float2* sub_525370(float2* a1, int a2);
int sub_5253B0(float* a1);
void nox_xxx_mapgen_525510(int a1, int a2);
void nox_xxx_mapgen_525570(int a1, int a2, int a3, int a4);
int nox_xxx_mapgen_525690(int a1, float2* a2, int a3);
int nox_xxx_mapgen_525740(int a1, float2* a2, int a3);
int nox_xxx_mapgen_525830(int a1, float2* a2, int a3);
int nox_xxx_mapgen_5258E0(int a1, float2* a2, int a3);
void sub_5259F0(int a1, int a2, float a3);
float* sub_525AF0(int a1);
void sub_525BF0(int a1);
float* sub_525C90();
int nox_xxx_mapGen_InPrefab1_525D20(int a1);
void sub_525DF0(float a1);
int nox_xxx_mapGenPrefabMkRoom_526100(int a1, int a2);
long long sub_526260(int a1, float* a2);
int sub_5262F0(int a1, int a2);
int sub_526550(int a1, int a2);
int nox_xxx_mapGen_InPrefab2_5266F0(int a1);
int nox_xxx_mapGenPlacePrefabs_526830(int a1);
int sub_5268F0(const char* a1);
char* sub_526950();
char* sub_526AA0(int a1);
void nox_xxx_mapGenSetRngSeed_526AB0(unsigned int a1);
signed int nox_xxx_mapGenRandFunc_526AC0(int a1, signed int a2);
int nox_xxx_mapGenRandFunc2_526B00(int a1, signed int a2);
double sub_526BC0(float a1, float a2);
int sub_526C40(int a1);
int sub_526C80(int a1);
int sub_526CA0(char* a1);
int sub_526D50(int a1);
int sub_526DD0(float2* a1, int* a2);
int sub_526E60(float* a1);
int sub_527030(float2* a1);
int sub_527380(float* a1);
int sub_527450(uint32_t* a1);
int nox_xxx_mapGenGetObjID_527940(char* a1);
float* nox_xxx_mapGenPlaceObj_5279B0(float2* a1);
float* nox_xxx_mapGenMoveObject_527A10(float* a1, float2* a2);
int nox_xxx_mapGenOrientObj_527C60(int a1, int a2);
int nox_xxx_mapGenFinishSpellbook_527DB0(int a1, char a2);
int* sub_527E00(nox_object_t* a1);
int nox_xxx_netUpdateObjectSpecial_527E50(nox_object_t* a1, nox_object_t* a2);
short sub_528030(int a1);
int nox_xxx_checkIsKillable_528190(nox_object_t* a1);
int nox_xxx_frameCounterSetCopyToNextFrame_5281D0();
int nox_xxx_frameCounterSetCopy_5281E0();
void nox_xxx_unitUpdateSightMB_5281F0(nox_object_t* a1);
int nox_xxx_aiLostSight_528560(int a1, int a2);
void sub_528610(int a1);
void nox_xxx_monsterUpdateSeenEnemies_5286D0(int a1, int a2);
void nox_xxx_monsterVisionSeeEnemy_5287B0(int a1, int a2);
int sub_528910(int a1, int a2);
int sub_528950(int a1, int a2);
int sub_528990(nox_object_t* a1);
void nox_xxx_netReportDestroyObject_5289D0(nox_object_t* a1);
int nox_xxx_netObjectOutOfSight_528A60(int a1, uint32_t* a2);
int nox_xxx_netObjectInShadows_528A90(int a1, uint32_t* a2);
int nox_xxx_monsterCmdSend_528BD0(int unit, int source, const char* command, short a4);
int nox_xxx_destroyEveryChatMB_528D60();
int nox_xxx_XFerMonster_528DB0(int a1);
void nox_xxx_monsterOnSpawnSpellcaster_529BC0(int a1);
int nox_xxx_XFer_ActionData_529CE0(int a1);
int nox_xxx_AssignIfGreater_52A420(int* a1, int a2);
int sub_52A440(int a1, int a2, int a3);
void nox_xxx_XFer_WriteShopItem_52A5F0(void* a1);
void nox_xxx_XFer_ReadShopItem_52A840(void* a1, int a2);
int nox_xxx_XFer_ReadMonsterBuffs_52AAB0(uint32_t* a1);
size_t nox_xxx_readNPCVoiceSet_52AD10(int a1);
int nox_xxx_XFerNPC_52ADE0(int a1);
int sub_52BA70(int a1);
int sub_52BAF0(int a1);
int nox_xxx_castCounterSpell_52BBB0(int a1, int a2, int a3, int a4);
void nox_xxx_cspellRemoveSpell_52BC90(int a1);
int nox_xxx_deathBallCreateFragments_52BD30(int* a1);
int sub_52BDB0(int a1, int a2);
void nox_xxx_changeOwner_52BE40(int a1, int a2);
int sub_52BEB0(int a1, int a2, int a3, int a4);
int nox_xxx_castSpellWinkORrestoreHealth_52BF20(int a1, int a2, int a3, int a4, int* a5);
int sub_52BF50(int a1, int a2, int a3, int a4, int* a5);
int nox_xxx_castPull_52BFA0(int a1, int a2, int a3, int a4, int a5, int a6);
int nox_xxx_castPush_52C000(int a1, int a2, int a3, int a4, int a5, int a6);
int nox_xxx_castFumble_52C060(int a1, int a2, int a3, int a4, int* a5);
int nox_xxx_castConfuse_52C1E0(int a1, int a2, int a3, int a4, int* a5, char a6);
int nox_xxx_castStun_52C2C0(int a1, int a2, int a3, int a4, int* a5, char a6);
int nox_xxx_castBurn_52C3E0(int a1, int a2, int a3, int a4, int a5);
int nox_xxx_useShock_52C5A0(int a1, int a2, int a3, int a4, int* a5, int a6);
int nox_xxx_castPoison_52C720(int a1, int a2, int a3, int a4, int* a5, int a6);
int nox_xxx_castFireball_52C790(int a1, int a2, int a3, int a4, int a5, int a6);
int sub_52CA80(int a1, int a2, int a3, int a4);
int sub_52CBD0(int a1, int a2, int a3, int a4);
int sub_52CCD0(int a1, int a2, int a3);
int nox_xxx_castCurePoison_52CDB0(int a1, int a2, int a3, int a4, int* a5, int a6);
void sub_52CE60(int a1);
int nox_xxx_castLock_52CE90(int a1, int a2, int a3, int a4);
void sub_52CF90(int a1, int a2);
void sub_52D060(int a1, int a2);
int nox_xxx_castTelekinesis_52D330(int a1, int a2, int a3, int a4, int* a5, char a6);
int nox_xxx_castFist_52D3C0(int a1, int a2, int a3, int a4, int a5, int a6);
int nox_xxx_castMeteorShower_52D8A0(int a1, int a2, int a3, int a4, int a5, int a6);
int nox_xxx_castMeteor_52D9D0(int a1, int a2, int a3, int a4, int a5, int a6);
int nox_xxx_castToxicCloud_52DB60(int a1, int a2, int a3, int a4, int a5);
int nox_xxx_spellArachna_52DC80(int a1, int a2, int a3, int a4, int a5);
int sub_52DD50(int a1, int a2, int a3, int a4, void* a5);
int nox_xxx_castEquake_52DE40(int a1, int a2, int a3, int a4, int a5, int a6);
short nox_xxx_equakeDamage_52DEC0(int a1, int a2);
void nox_xxx_objectApplyForce_52DF80(float* vec, nox_object_t* obj, float force);
unsigned int nox_xxx_isObjectMovable_52E020(int a1);
void nox_xxx_mapPushUnitsAround_52E040(int a1, float a2, int a3, float a4, int a5, int a6, int a7);
void nox_xxx_unitPushAroundFn_52E0E0(int a1, int** a2);
int nox_xxx_spellDrainMana_52E210(float a1);
int sub_52E450(int a1, int a2, int a3);
int sub_52E610(int* a1, int a2);
void sub_52E660(int a1, int a2);
int sub_52E7C0(int a1);
int nox_xxx_spellEnergyBoltStop_52E820(int a1);
int nox_xxx_spellEnergyBoltTick_52E850(float a1);
void nox_xxx_spellEnergyBoltSetTarget_52EC60(int a1, int a2);
int nox_xxx_firewalkTick_52ED40(float* a1);
int sub_52EF30(int a1);
int sub_52EFD0(int a1);
int sub_52F1D0(int a1);
int sub_52F220(int* a1);
int sub_52F2E0(float a1);
int sub_52F460(float a1);
int nox_xxx_castShield1_52F5A0(uint32_t* a1);
int sub_52F650(int a1);
int sub_52F670(int a1);
char nox_xxx_unitShield_52F690(int a1, int a2);
void nox_xxx_unitShieldReduceDamage_52F710(int a1, int* a2, int a3, int a4);
int nox_xxx_onStartLightning_52F820(int a1);

#endif // NOX_PORT_GAME4_2
