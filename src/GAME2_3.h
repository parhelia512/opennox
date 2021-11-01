#ifndef NOX_PORT_GAME2_3
#define NOX_PORT_GAME2_3

#include "defs.h"


int sub_48C4D0();
int sub_48C560();
void sub_48C580(pixel8888* a1, int num);
short sub_48C5B0(short* a1, int a2);
int sub_48C5E0(int a1, int a2);
short sub_48C610();
uint32_t* sub_48C650(int a1, int a2, int a3, uint32_t* a4, uint32_t* a5);
unsigned int sub_48C690(int a1, int a2, int a3, int a4);
unsigned int sub_48C6B0(int a1, int a2);
double sub_48C6D0(float a1, float a2, float a3, float a4);
double sub_48C700(float a1, float a2);
int nox_client_winVerGetMajor_48C870(char* printOut);
int nox_xxx_showObserverWindow_48CA70(int a1);
int sub_48CA90();
int sub_48CAB0();
int sub_48CAD0();
int sub_48D000();
int nox_xxx_guiKick_48D0A0(int a1, int a2, int* a3, int a4);
int sub_48D120();
char* nox_xxx_voteSend_48D260(wchar_t* a1);
char* nox_xxx_netSendRenameMb_48D2D0(wchar_t* a1);
int sub_48D340();
int sub_48D3B0();
int nox_xxx_clientVote_48D3E0();
uint32_t* sub_48D410();
int sub_48D450();
int sub_48D4A0();
int sub_48D4B0(int a1);
int sub_48D4E0();
int sub_48D4F0(unsigned short a1, unsigned short a2);
int sub_48D560(unsigned short a1);
uint32_t* sub_48D5A0(int a1);
int sub_48D660();
int sub_48D740();
void sub_48D760();
int* sub_48D7B0();
int nox_xxx_chatInit_48D7D0();
int sub_48D800();
int sub_48D830(nox_drawable* dr);
int nox_xxx_netCode2ChatBubble_48D850(int a1);
wchar_t* nox_xxx_createTextBubble_48D880(int a1, wchar_t* a2);
void sub_48D990(uint32_t* a1);
void sub_48DCF0(uint32_t* a1);
bool sub_48E000(int4* a1, uint32_t* a2);
char sub_48E240(int a1, uint32_t* a2);
int sub_48E480(uint32_t* a1, uint32_t* a2);
int sub_48E530(int a1, int a2);
int sub_48E5C0(uint32_t* a1, int a2, int a3);
int* sub_48E6A0(char a1, uint32_t* a2, uint32_t* a3, int* a4, int* a5);
void sub_48E8E0(int a1);
void sub_48E940();
uint32_t* nox_xxx_spriteCreate_48E970(int a1, unsigned int a2, int a3, int a4);
char* sub_4947E0(int a1);
int sub_4948B0(int a1);
int nox_xxx_netCliProcUpdateStream_494A60(unsigned char* a1, int a2, uint32_t* a3);
unsigned char* nox_xxx_netCliUpdateStream2_494C30(unsigned char* a1, int a2, int* a3);
int nox_netlist_receiveCli_494E90(int ind);
int sub_494F00();
char* sub_494FF0();
char* sub_495020(int a1);
int sub_495060(int a1, short a2, short a3);
int sub_4950C0(int a1);
int sub_4950F0(int a1, char a2);
int sub_495120(int a1, short a2, short a3);
int sub_495150(int a1, short a2);
int sub_495180(int a1, uint16_t* a2, uint16_t* a3, uint8_t* a4);
char* sub_4951C0();
int nox_xxx_unitSpriteCheckAlly_4951F0(int a1);
int sub_495210(int a1);
int sub_495430();
int* sub_495500(int* a1);
int sub_4958F0();
int nox_xxx_allocClassListFriends_495980();
void sub_4959B0();
int sub_4959D0();
uint32_t* nox_xxx_cliAddObjFriend_4959F0(int a1);
void sub_495A20(int a1);
int sub_495A80(int a1);
int nox_xxx_allocArrayDrawableFX_495AB0();
int sub_495AE0();
void sub_495B00(nox_drawable* dr);
uint32_t* sub_495B50(uint32_t* a1);
uint32_t* sub_495BB0(nox_drawable* dr, nox_draw_viewport_t* vp);
int sub_495BF0(int a1, int a2, int a3);
int sub_495D00(uint32_t* a1, int a2, uint32_t* a3);
void sub_495F30(int a1, int a2);
void sub_495F70(int a1);
uint32_t* sub_495FC0(uint32_t* a1, int a2);
int sub_496020(int a1, int a2);
void sub_496050(int a1);
int sub_4960B0();
int sub_496120();
int nox_xxx_drawBlack_496150(nox_draw_viewport_t* a1p);
int sub_497180(float* a1, float* a2, float* a3);
int sub_497260(int* a1);
void sub_4974B0(int a1);
int sub_497650(float a1);
int sub_4977C0(int a1);
char sub_497B80(float* a1, int* a2);
int nox_xxx_drawBlackofWall_497C40(int a1, int a2, char a3);
int sub_497F60(int a1, int a2, char a3, int a4, int a5);
int sub_498030(uint32_t* a1);
int sub_498110();
int sub_498130(int a1);
int sub_4981D0(int a1);
int sub_498220(int a1, int a2);
int sub_498290(int a1);
int sub_4982E0(int a1);
int sub_498330(int a1);
void sub_498380(int a1, int a2);
int nox_xxx_client_4984B0_drawable(nox_drawable* dr);
void sub_4989A0();
void sub_498AE0();
int sub_498B50(int a1, int a2, int a3, int a4);
int sub_498C20(int2* a1, int2* a2, int a3);
char sub_4990D0(uint32_t* a1, uint32_t* a2);
int sub_499130(int* a1);
int sub_499160(int2* a1, int2* a2, int2* a3);
void sub_4991E0(uint32_t* a1);
uint64_t sub_499290(int a1);
int sub_4992B0(int a1, int a2);
int nox_xxx_loadReflSheild_499360();
int sub_499450();
int nox_xxx_drawShield_499810(nox_draw_viewport_t* vp, nox_drawable* dr);
uint32_t* nox_xxx_fxDrawTurnUndead_499880(short* a1);
void nox_xxx_drawPointMB_499B70(int xLeft, int yTop, int a3);
void nox_xxx_bookRewardCli_499CF0(int* a1, int a2, int a3);
uint32_t* sub_499F60(int a1, int a2, int a3, short a4, char a5, char a6, char a7, char a8, char a9, int a10);
void nox_alloc_npcs();
nox_npc* nox_new_npc(int id);
nox_npc* nox_npc_by_id(int id);
int nox_init_npc(nox_npc* ptr, int id);
nox_npc* nox_npc_set_328(int id, int a2);
char* nox_xxx_clientEquip_49A3D0(char a1, int a2, int a3, int a4);
int nox_xxx_allocArrayHealthChanges_49A5F0();
void sub_49A630();
uint16_t* nox_xxx_cliAddHealthChange_49A650(int a1, short a2);
void sub_49A6A0(nox_draw_viewport_t* vp, nox_drawable* dr);
void sub_49A880(int a1);
int sub_49A8C0();
int sub_49A8E0_init();
int sub_49A950_free();
int nox_xxx_sprite_49A9B0_drawable(nox_drawable* dr);
void nox_xxx_sprite_49AA00_drawable(nox_drawable* dr);
void nox_xxx_updateSpritePosition_49AA90(nox_drawable* dr, int a2, int a3);
#ifndef NOX_CGO
void nox_xxx_forEachSprite_49AB00(int4* a1, void(* fnc)(nox_drawable*, int), int a3);
#else // NOX_CGO
void  nox_xxx_forEachSprite_49AB00(int4*, void*, int);
#endif // NOX_CGO
nox_drawable* nox_drawable_find_49ABF0(nox_point* pt, int r);
int sub_49AD20(uint32_t* a1, int a2);
int nox_xxx_unused_49ADD0(int a1);
int sub_49AEA0();
void sub_49AEE0();
int* nox_xxx_unused_49AF80(int a1, int a2, int a3);
int nox_xxx_unused_49B0A0(int a1, int a2, int a3);
int sub_49B1A0(int a1);
int sub_49B1D0(int a1, int a2);
int sub_49B260();
int sub_49B370();
int sub_49B380();
int sub_49B3A0();
int sub_49B3B0();
int sub_49B3C0();
int sub_49B3E0();
int sub_49B420(int a1, int a2, int* a3, int a4);
int sub_49B490();
int sub_49B6B0();
void nox_xxx_consoleEsc_49B7A0();
uint32_t* nox_xxx_spriteTransparentDecay_49B950(uint32_t* a1, int a2);
uint32_t* nox_xxx_sprite_49BA10(nox_drawable* dr);
int sub_49BA70();
uint32_t* nox_xxx_spriteToSightDestroyList_49BAB0_drawable(uint32_t* a1);
uint32_t* sub_49BAF0(nox_drawable* dr);
void sub_49BB40();
void* sub_49BB80(char a1);
void sub_49BBB0();
void sub_49BBC0();
uint32_t* nox_xxx_spriteToList_49BC80_drawable(uint32_t* a1);
uint32_t* sub_49BCD0(nox_drawable* dr);
int nox_xxx_getSomeSprite_49BD40();
int nox_xxx_getSprite178_49BD50(int a1);
void sub_49BD70(nox_draw_viewport_t* a1p);
uint32_t* nox_xxx_clientAddRayEffect_49C160(int a1);
void nox_xxx_clientRemoveRayEffect_49C450(int a1);
void nox_xxx_spriteDeleteSomeList_49C4B0();
void nox_xxx_sprite_49C4F0();
int sub_49C520(int a1);
int nox_xxx_wnd_49C760(int a1, int a2, int* a3, int a4);
int sub_49C7A0();
int sub_49C810();
int sub_49CA60(int a1, int a2, int* a3, int a4);
int sub_49CB40();
int nox_video_initRectDrawingFuncs_49CB50();
void nox_client_drawBorderLines_49CC70(int xLeft, int yTop, int a3, int a4);
void sub_49CD30(int xLeft, int yTop, int a3, int a4, int a5, int a6);
void nox_client_drawRectFilledOpaque_49CE30(int xLeft, int yTop, int a3, int a4);
int4* nox_client_drawRectFilledAlpha_49CF10(int xLeft, int yTop, int a3, int a4);
int4* sub_49CFB0(int xLeft, int yTop, int a3, int a4);
int4* sub_49D050(int xLeft, int yTop, int a3, int a4);
int4* nox_client_drawRectFadingScreen_49D0F0(int xLeft, int yTop, int a3, int a4);
int nox_client_drawRectStringSize_49D190(int a1, int a2, int a3, int a4, int a5);
int sub_49D1C0(void* a1, int a2, int a3);
void sub_49D1E0(int a1, int a2, unsigned int a3, int a4);
void nox_xxx_draw_49D270_MBRect_49D270(int a1, int a2, unsigned int a3, int a4);
int sub_49D2F0(int a1, int a2, int a3, int a4);
void sub_49D370(int a1, int a2, int a3, int a4);
int sub_49D540(int a1, int a2, int a3, int a4);
int sub_49D680(int a1, int a2, int a3, int a4);
short sub_49D6F0(int a1, int a2, unsigned int a3, int a4);
void sub_49D770(int a1, int a2, unsigned int a3, int a4);
void sub_49D880(int a1, int a2, int a3, int a4);
int sub_49D8E0(int a1, int a2, int a3, int a4);
int sub_49D9A0(int a1, int a2, int a3, int a4);
unsigned char sub_49DA90(int a1, int a2, int a3, int a4);
unsigned char sub_49DB20(int a1, int a2, int a3, int a4);
int sub_49DBB0(int a1, int a2, int a3, int a4);
int sub_49DC70(int a1, int a2, int a3, int a4);
int4* sub_49DD60(int a1, int a2, int a3, int a4, int a5);
int4* sub_49E060(int a1, int a2, int a3, int a4, int a5);
void sub_49E380(uint32_t* a1, int a2, unsigned int a3);
void sub_49E3C0(uint32_t* a1, int a2, unsigned int a3);
int nox_video_initLineDrawingFuncs_49E3F0();
int nox_client_drawLineFromPoints_49E4B0();
int sub_49E4C0(int a1, int a2, int a3, int a4);
int sub_49E4F0(int a1);
int sub_49E510(int a1);
int sub_49E540(int a1);
int sub_49E6C0(int a1);
int sub_49E930(int a1);
int sub_49EAB0(int a1);
int sub_49ED80(unsigned char a1, int a2);
int sub_49EFA0(int a1, int a2);
int sub_49EFC0(int a1, int a2);
int sub_49F010(int a1, int a2);
void sub_49F060(int a1, int a2, int a3);
unsigned char* sub_49F0F0(int a1, int a2, int a3);
void sub_49F180(int a1, int a2, int a3);
short sub_49F210(int a1, int a2, int a3);
void sub_49F3A0(int a1, int a2, int a3);
void sub_49F420(int a1, int a2, int a3);
int sub_49F4A0();
void sub_49F4D0();
int nox_client_drawAddPoint_49F500(int a1, int a2);
int nox_xxx_rasterPointRel_49F570(int a1, int a2);
int sub_49F5B0(uint32_t* a1, uint32_t* a2, int a3);
int sub_49F610();
int sub_49F6D0(int a1);
int sub_49F6E0();
int4* nox_client_copyRect_49F6F0(int xLeft, int yTop, int a3, int a4);
int4* sub_49F780(int xLeft, int a2);
int4* sub_49F7C0(int a1);
int nox_xxx_wndDraw_49F7F0();
int sub_49F860();
bool sub_49F8E0(int x, int y, int d);
int4* nox_xxx_utilRect_49F930(int4* a1, int4* a2, int4* a3);
int sub_49F990(int* a1, int* a2, int* a3, int* a4);
int sub_49FC20(int* a1, int* a2, int* a3, int* a4);
void sub_49FDB0(int a1);
uint32_t* sub_49FF20();
char* nox_client_getChatMap_49FF40(short* a1);
int* sub_49FFA0(int a1);
char* sub_4A0020();
int nox_wol_servers_addResult_4A0030(nox_gui_server_ent_t* srv);
void nox_wol_servers_sortBtnHandler_4A0290(int id);
int sub_4A0330(int* a1);
int* sub_4A0360();
int* sub_4A0390();
int sub_4A0410(const char* a1, short a2);
int* sub_4A0490(int a1);
int* sub_4A04C0(int a1);
int* sub_4A04F0(char* a1);
void sub_4A0540(void* lpMem);
int nox_xxx_getConnResult_4A0560();
int nox_gui_parseColor_4A0570(unsigned int* out, char* buf);
int nox_gui_parseColorTo_4A05E0(unsigned int* out, FILE* f, char* buf);
int nox_gui_parseWindowBgColor_4A0650(nox_window_data* data, const char* buf);
int nox_gui_parseWindowEnColor_4A0690(nox_window_data* data, const char* buf);
int nox_gui_parseWindowDisColor_4A06D0(nox_window_data* data, const char* buf);
int nox_gui_parseWindowHlColor_4A0710(nox_window_data* data, const char* buf);
int nox_gui_parseWindowSelColor_4A0750(nox_window_data* data, const char* buf);
int nox_gui_parseWindowTextColor_4A0790(nox_window_data* data, const char* buf);
int nox_gui_parseWindowFont_4A07D0(nox_window_data* data, const char* buf);
int nox_gui_parseWindowImgOffs_4A0830(nox_window_data* data, const char* buf);
int nox_gui_parseWindowBgImage_4A0870(nox_window_data* data, const char* buf);
int nox_gui_parseWindowEnImage_4A08C0(nox_window_data* data, const char* buf);
int nox_gui_parseWindowDisImage_4A0910(nox_window_data* data, const char* buf);
int nox_gui_parseWindowSelImage_4A0960(nox_window_data* data, const char* buf);
int nox_gui_parseWindowHlImage_4A09B0(nox_window_data* data, const char* buf);
int nox_gui_parseWindowStatus_4A0A00(nox_window_data* data, const char* buf);
int nox_gui_parseWindowStyle_4A0A30(nox_window_data* data, const char* buf);
int nox_gui_parseWindowGroup_4A0A60(nox_window_data* data, const char* buf);
void nox_gui_resetWidgetData_4A0D10();
int nox_gui_parseFont_4A0D40(int* out, FILE* f, char* buf);
int nox_xxx_guiParse_4A1780(int a1, FILE* a2, char* a3);
#ifndef NOX_CGO
nox_window*  nox_new_window_from_file(const char* name, int (*fnc)(int, int, int, int));
void nox_gui_resetWidgetData_4A0D10();
int nox_gui_parseFont_4A0D40(int* out, FILE* f, char* buf);
nox_window*  nox_gui_parseWindowRoot_4A0D80(FILE* f, char* buf, int (*fnc)(int, int, int, int));
nox_window*  nox_gui_parseWindowOrWidget_4A1440(const char* typ, int id, int a3, int px, int py, int w, int h, nox_window_data* drawData, void* data, int (*fnc)(int, int, int, int));
nox_window*  nox_gui_parseWidget_4A1510(const char* typ, nox_window* parent, int a3, int px, int py, int w, int h, uint32_t* drawData, void* data);
int  nox_xxx_guiParse_4A1780(int a1, FILE* a2, char* a3);
#else // NOX_CGO
nox_window* nox_new_window_from_file(char* cname, void* fnc);
#endif // NOX_CGO

unsigned sub_48C730(unsigned int a1);

#endif // NOX_PORT_GAME2_3
