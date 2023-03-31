#ifndef NOX_PORT_GAME5
#define NOX_PORT_GAME5

#include "defs.h"

char nox_xxx_mobActionRandomWalk_545020(int a1);
int sub_545090(int a1);
char nox_xxx_mobActionConfuse_545140(int a1);
char sub_545210(int a1);
char sub_545240(int a1, float* a2);
char sub_545300(int a1);
char sub_545340(int a1);
char sub_5453E0(int a1);
char nox_xxx_mobActionRetreat_545440(int a1);
int nox_xxx_monsterCanResumeAttack_545520(int a1);
int sub_545580(int a1);
int nox_xxx_monsterCanCast2_5455B0(int a1);
int* nox_xxx_mobRetreatCheckEdibles_5455E0(int a1);
int sub_5456B0(int a1);
int sub_5456C0(int a1);
void sub_5456D0(int a1);
long long sub_545790(int a1);
int sub_5457C0(int a1);
char nox_xxx_mobActionRoam_5457E0(int* a1);
void sub_545B00(int a1, int a2);
int sub_545B60(int a1, unsigned char a2);
int nox_xxx_monsterRoamDeadEnd_545BB0(int a1, int a2);
int sub_545C60(int a1, int a2, unsigned char a3);
int sub_545DA0(int a1);
int sub_545E60(nox_object_t* a1);
int* nox_xxx_mobActionGuard_546010(int a1);
int sub_546410(int a1);
int sub_546420(int a1);
char nox_xxx_mobActionEscort_546430(int* a1);
int nox_xxx_monsterGetObjEscortName_546600(int a1);
int nox_xxx_monsterLookAtDamager_5466B0(int a1);
int sub_5466F0(nox_object_t* a1);
void nox_xxx_mobAction_5469B0(nox_object_t* a1);
short nox_xxx_unitIsDangerous_547120(nox_object_t* a1, nox_object_t* a2);
void nox_xxx_monsterPopAttackActions_5471B0(int a1);
void nox_xxx_monsterMainAIFn_547210(nox_object_t* a1);
int nox_xxx_monsterCheckDodgeables_547C50(int a1);
int sub_547DB0(int a1, float2* a2);
int sub_547EE0(int a1, unsigned char a2);
int sub_547F10();
int sub_547F20(int a1, float* a2);
int nox_xxx_BuildWaypointPath_547F70(uint32_t* a1, int a2, uint32_t* a3, int a4);
void sub_548100(int2* a1, int a2);
void sub_5481C0(int a1);
void sub_548220(int* a1, float* a2);
int sub_548360(int a1, int a2);
int sub_5485B0(int a1, int a2);
void sub_548600(nox_object_t* a1, float a2, float a3);
void nox_xxx_collSysAddCollision_548630(int a1, unsigned int a2, float2* a3);
void nox_xxx_allocHitArray_5486D0();
void sub_548830(int a1);
void nox_xxx_collide_548740();
void sub_548860(int a1, short a2);
void sub_5488B0(int* a1, float* a2, int a3);
void sub_548B60();
void nox_xxx_script_forcedialog_548CD0(nox_object_t* a1, nox_object_t* a2);
void sub_548FE0(int a1, const char* a2);
int nox_xxx_monsterLoadStrikeFn_549040(int a1, char* a2);
int nox_xxx_monsterLoadDieFn_5490E0(int a1, char* a2);
int nox_xxx_monsterLoadDeadFn_549180(int a1, char* a2);
int nox_xxx_strikeOgre_549220(float a1);
void sub_549270(int a1, int a2);
int nox_xxx_strikeMonsterDefault_549380(float a1);
int nox_xxx_monsterPickMeleeTarget_549440(int a3, int a2);
void sub_5494C0(float* a1, int a2);
int nox_xxx_strikeScorpion_5495B0(float a1);
int sub_549690(int a1, int a2);
int nox_xxx_strikeVileZombie_549700(float a1);
int nox_xxx_strikeStoneGolem_5497E0(float a1);
int nox_xxx_sendEquakeAfterGolem_549800(float a1);
void nox_xxx_monsterAttackAreaDamage_549860(int a1, float a2);
int nox_xxx_strikeMechGolem_549960(float a1);
int nox_xxx_strikeWasp_549980(float a1);
int nox_xxx_strikeGhost_549A60(float a1);
int nox_xxx_strikeBomber_549BB0();
int nox_xxx_strikeSpider_549BC0(float a1);
int nox_xxx_strikeSpittingSpider_549CA0(float a1);
int sub_549D80(int a1);
int sub_549E00(int a1);
int sub_549E70(int a1);
int sub_549E90(int a1);
int sub_549FA0(int a1);
int nox_bomberDead_54A150(nox_object_t* a1);
int sub_54A250(int a1);
int nox_xxx_monsterDeadTroll_54A270(int a1);
int sub_54A310(int a1);
void sub_54A390(int a1, char* a2, const char* a3, const char* a4, const char* a5, const char* a6, int a7);
void sub_54A4C0(int a1);
int sub_54A750(int a1);
int sub_54A7D0(int a1);
int sub_54A850(int a1);
int sub_54A890(int a1);
int sub_54A900(int a1);
int sub_54A950(int a1);
void sub_54AD50(int a1, int a2, int a3);
double sub_54A990(float2* a1, float a2, int a3, float2* a4);
nox_object_t* nox_xxx_findObjectAtCursor_54AF40(nox_object_t* a1);
void nox_xxx_playerCursorScanFn_54AFB0(int a1, float* a2);
int sub_54B2D0(int* a1, int a2, uint32_t* a3);
int sub_54B810(int a1, int a2, int* a3, int2* a4, int a5);
int sub_54BA60(int a1, int a2, int a3, int a4);
int sub_54BB20(int a1, int a2, int* a3, uint32_t* a4, int a5);
int sub_54BD90(int a1, int a2, int* a3, int* a4, int a5);
int sub_54BF20(int a1, int a2, int* a3, int* a4, int a5);
short nox_xxx_monsterAutoSpells_54C0C0(int a1);
char nox_xxx_monsterCreateFn_54C480(int a1);
int nox_xxx_createWeapon_54C710(int a1);
uint32_t* sub_54C950(int a1);
int nox_xxx_createFnObelisk_54CA10(int a1);
int* nox_xxx_createFnAnim_54CA50(int a1);
uint8_t* nox_xxx_createTrigger_54CA60(int a1);
uint32_t* nox_xxx_createMonsterGen_54CA90(int a1);
uint32_t* nox_xxx_createRewardMarker_54CAC0(int a1);
int nox_xxx_dieImpEgg_54CAE0(int a1);
void nox_xxx_diePolyp_54CB10(int a1);
void nox_xxx_diePotion_54CBB0(int a1);
void sub_54CBD0(int a1);
void sub_54CC40(int a1);
void sub_54CD30(int a1);
void sub_54CE00(int a1);
void sub_54CEE0(int a1);
char sub_54CFB0(int a1);
void sub_54D080(int a1);
int nox_xxx_diePlayer_54D2B0(int a1);
void nox_xxx_playerHandleElimDeath_54D7A0(int a1, int a2);
void nox_xxx_playerUpdateScore_54D980(int a1, int a2, int a3, int a4);
void nox_xxx_playerHandleKotrDeath_54DC40(int a1, int a2);
void nox_xxx_netNotifyPlayerDied_54DF00(int a1);
void nox_xxx_dieGlyph_54DF30(nox_object_t* a1);
void nox_xxx_dieBarrel_54DFA0(int a1);
void nox_xxx_dieCreateObject_54E010(int a1);
short nox_xxx_dieSpawnObject_54E070(int a1);
void nox_xxx_dieMarker_54E460(int a1);
void nox_xxx_dieBoulder_54E4B0(int a1);
int nox_xxx_dieGameBall_54E620(int a1);
void nox_xxx_dieMonsterGen_54E630(int a1);
int sub_54E6F0(int a1, int a2);
int sub_54E730(int a1, int a2);
int sub_54E810(int a1, float2* a2, int a3);
void sub_54E850(int a1, int a2);
char nox_xxx_updateMonsterGenerator_54E930(uint32_t* a1);
int nox_xxx_mobGeneratorPick_54EBA0(uint32_t* a1, float2* a2, int a4);
int nox_xxx_mgenSetCreaturePos_54ED50(int a1, float2* a2, int a3, int a4);
int sub_54EF00(float* a3);
void sub_54EF60(float* a1, int a2);
int sub_54EF90(float a1, int a2, int a3, int a4);
uint32_t* nox_xxx_mobGeneratorSpawn_54F070(int a1, int a2, int a3);
void nox_xxx_unitCreatureCopyUC_54F2B0(int a1, int a2);
void nox_xxx_updateHarpoon_54F380(nox_object_t* a1);
void nox_xxx_unitUpdateMover_54F740(int a1);
int nox_xxx_updateShootingTrap_54F9A0(int a1);
void nox_xxx_createArrowTrapProjectile_54FA80(int a1, int a2);
void sub_54FBB0(int a1);
int sub_54FBF0(int a3);
void nox_xxx_unitIsAttackReachable_54FC50(int a1, int a2);
void nox_xxx_collideTrigger_54FCD0(int a1, int a2);
float* nox_xxx_createSpark_54FD80(float a1, float a2, int a3, int a4, float a5, float a6, float a7, int a8);
void sub_54FEF0(int a2);
int sub_54FFC0(int2* a1, int a2);
int sub_550280(int a1, float a2, float a3, int a4, int a5, int a6, int a7);
int sub_5502F0(float2* a1, float a2, float a3, int a4, int a5, float2* a6, float2* a7);
int sub_550380(int a1, int a2, float2* a3);
int sub_550480(int a1);
void sub_5504B0(int a2);
int sub_550580(int2* a1, float* a2);
int sub_550760(int a1, float2* a2, float2* a3, float4* a4, float2* a5, float a6);
int sub_550A10(int a1, float2* a2, float2* a3, float4* a4, float2* a5, float a6);
char sub_550CB0(float2* a1, float2* a2);
void nox_xxx_collisionCheckCircleCircle_550D00(int a1, int a2);
void sub_550F80(float* a1, int a2);
void sub_551250(unsigned int a1, float* a2, int a3);
int sub_5516A0(float4* a1, float4* a2, float2* a3, int a4, int a6);
int sub_551780(float4* a1, float a2, float a3, float a4, float2* a5, int a6);
int sub_551870(float4* a1, float a2, float a3, float a4, float2* a5, int a6);
int sub_551960(float4* a1, float4* a2, float4* a3, float2* a4);
int sub_551A90(float2* a1, float4* a2);
void sub_551AE0(int a1, int a2, int a3);
void sub_551BF0();
void sub_551C40(int a1, int a2);
int nox_xxx_netSendSock_552640(int id, unsigned char* buf, int sz, int flags);

#endif // NOX_PORT_GAME5
