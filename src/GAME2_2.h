#ifndef NOX_PORT_GAME2_2
#define NOX_PORT_GAME2_2

#include "defs.h"

int sub_476080(unsigned char* a1);
int sub_4761B0(nox_drawable* a1p);
void sub_476270(nox_draw_viewport_t* a1p);
int sub_4765F0(nox_draw_viewport_t* a1p);
int sub_476680();
void nox_video_setCutSize_4766A0(int a1);
int nox_video_getCutSize_4766D0();
void nox_draw_setCutSize_476700(int cutPerc, int a2);
void* sub_476AE0(nox_draw_viewport_t* vp, nox_drawable* dr);
short sub_476D70(uint32_t* a1, int* a2, unsigned int a3);
int nox_client_setPhonemeFrame_476E00(int a1);
uint32_t* sub_476E20();
int sub_476E90();
int nox_xxx_clientDeleteSprite_476F10_drawable(nox_drawable* dr);
unsigned int nox_xxx_packetGetMarshall_476F40();
void nox_xxx_clientEnumHover_476FA0();
void nox_xxx_clientOnCursorHover_477050(int arg0, int a2);
void nox_video_drawCursorSelectCircle2_477470(uint32_t* a1, int a2, int a3);
int nox_xxx_guiCursor_477600();
int nox_client_setCursorType_477610(int a1);
int nox_xxx_cursorGetTypePrev_477630();
void nox_xxx_bookSaveSpellForDragDrop_477640(int a1, int a2);
void nox_xxx_bookSpellDnDclear_477660();
int nox_xxx_bookGetSpellDnDType_477670();
void nox_xxx_cursorSetDraggedItem_477690(nox_drawable* a1);
void nox_xxx_cursorResetDraggedItem_4776A0();
void nox_xxx_cursorSetTooltip_4776B0(wchar_t* a1);
int sub_478000();
int sub_478030();
int sub_478040();
int sub_478080(int a1);
char* sub_4780A0(int a1);
int sub_478110();
int sub_478480(int a1, int a2, int* a3, int a4);
int sub_478650(int a1, int a2, unsigned int a3);
void sub_478850(int a1, short a2, int a3, int a4);
int sub_478970();
int sub_478A70(int2* a1);
int sub_478C80();
int sub_478E50(int a1, int a2, unsigned int a3);
uint32_t* sub_478F10();
int sub_478F80();
char* nox_xxx_getShopPic_4790F0(int a1);
int sub_479280();
uint32_t* sub_479300(int a1, int a2, int a3, short a4, int a5);
char* sub_4793C0(int a1);
char* sub_479430();
char* sub_479480(int a1);
int sub_4794D0(int a1, int a2);
int sub_479590();
void sub_4795A0(int a1);
int sub_479690(int a1, short a2, short a3, int a4);
int nox_client_tradeXxxSellAccept_4796D0(short a1);
int sub_479700(short a1, char a2);
void sub_479810();
int sub_479820(int a1, short a2);
int sub_479840(short a1);
int sub_479870();
bool sub_479880(uint32_t* a1);
int sub_4798A0(uint32_t* a1);
int sub_479950();
int sub_4799A0();
int nox_xxx_guiDialog_479B00(int a1, int a2, int* a3, int a4);
int sub_479BE0(uint32_t* a1, int a2, unsigned int a3, int a4);
int sub_479C40(uint32_t* a1, int a2);
int sub_479CB0(int a1, int a2);
int sub_479D00();
int sub_479D10();
int sub_47A1F0();
int sub_47A260();
void nox_client_drawImageAt_47D2C0(nox_video_bag_image_t* img, int x, int y);
void sub_47D370(int a1);
int sub_47D380(int a1, int a2);
void sub_47D400(int a1, char a2);
int sub_47D5C0(int a1, uint32_t* a2, uint32_t* a3, uint32_t* a4, uint32_t* a5);
unsigned char sub_47DBC0();
int sub_47FCE0(uint32_t* a1, int a2);
uint8_t* sub_480220(uint8_t* a1, uint8_t* a2);
uint16_t* sub_480250(uint8_t* a1, uint16_t* a2);
void sub_480860(uint16_t* a1, uint16_t* a2, int a3, int* a4, int* a5);
void sub_480BE0(uint32_t* a1, int* a2, int a3, int a4, int a5);
int* nox_xxx_edgeDraw_480EF0(int a1, int a2, int a3, int* a4, int* a5, int a6, int a7, int a8, int a9, int a10);
void sub_481410();
char nox_xxx_tileDraw_4815E0(uint32_t* a1, int a2);
char* sub_481770(uint32_t* a1, int a2, unsigned short a3);
char nox_xxx_drawTexEdgesProbably_481900(uint32_t* a1, uint32_t* a2);
void nox_xxx_tileCallDrawEdges_481BF0(int a1, int a2);
int nox_xxx_tileCheckRedrawMB_482570(nox_draw_viewport_t* vp);
int nox_xxx_tileDrawImpl_4826A0(nox_draw_viewport_t* vp);
short sub_4831C0(int a1, int a2);
int sub_484450(int a1, int a2);
uint32_t* nox_xxx_spriteChangeLightColor_484BE0(uint32_t* a1, int a2, int a3, int a4);
long long sub_484C00(int a1, int a2);
long long nox_xxx_spriteChangeLightSize_484C30(int a1, int a2);
int sub_484C60(float a1);
int sub_484CE0(int a1, float a2);
int nox_xxx_spriteChangeIntensity_484D70_light_intensity(int a1, float a2);
int nox_thing_read_floor_485B30(int a1, char* a2);
int nox_thing_read_edge_485D40(int a1, char* a2);
int nox_xxx_tile_486060();
int sub_4862E0(int a3, int a4);
uint32_t* sub_486320(uint32_t* a1, int a2);
int sub_486350(void* a1, int a2);
int sub_486380(uint32_t* a1, unsigned int a2, int a3, int a4);
int sub_4863B0(unsigned int* a2);
uint32_t* sub_4864A0(uint32_t* a3);
int sub_486520(unsigned int* a2);
int sub_486550(uint8_t* a1);
int sub_486570(unsigned int* a1, uint32_t* a2);
uint32_t* sub_486620(uint32_t* a1);
int sub_486640(void* a1, int a2);
int sub_486670(int a1, int a2);
int sub_4866A0(int a1);
int sub_4866D0(uint32_t* a1, int a2);
char* sub_4866F0(const char* a1, const char* a2);
void sub_4869C0(void* lpMem);
unsigned int sub_486A10(int a1, void* a2);
int sub_486AA0(uint32_t* a1, int a2, uint32_t* a3);
int sub_486B60(int a1, int a2);
signed int sub_486DB0(int a1, char* a2, signed int a3);
FILE* sub_486E00(int a1);
int sub_486E30(int a1, uint32_t* a2);
int sub_486E90(int a1);
void* sub_486EF0();
int sub_486F30();
uint32_t* sub_486FA0(int a1);
uint32_t* sub_486FE0(int a1);
void sub_487030(void* lpMem);
void sub_487050(uint32_t* a1);
void sub_487070(void* lpMem);
uint32_t** sub_487090(uint32_t** a1);
void sub_4870A0();
int* sub_4870E0(int* a1);
int* sub_487100(int** a1);
uint32_t* sub_487150(int a1, const void* a2);
uint32_t* sub_4871C0(int a1, int a2, const void* a3);
void sub_4872C0(void* lpMem);
int sub_487310(uint32_t* a1);
int* sub_487360(int a1, int** a2, int* a3);
int sub_4873C0(int a3);
int sub_487590(int a1, const void* a2);
int* sub_4875B0(int* a1);
int* sub_4875D0(int** a1);
int sub_4875F0();
void sub_487680(void* lpMem);
void* sub_4876A0(uint32_t** a1);
uint32_t* sub_487750(int a1);
int sub_487790(int a1, int a2);
int* sub_4877D0(int a1, int* a2);
int* sub_4877F0(int** a1);
int* sub_487810(int a1, int a2);
int sub_487910(int a1, int a2);
int* sub_487970(int a1, int a2);
void sub_487C30(uint32_t* a1);
int sub_487C50(int a1, uint32_t* a2);
int sub_487C80(int a1);
int sub_487D00(uint32_t* a1);
uint32_t* sub_487D30(uint32_t* a1, int a2, int a3);
int sub_487D60(int a1);
int nox_xxx_wndEditProc_487D70(nox_window* a1, int a2, int a3, int a4);
int nox_xxx_wndEditDrawNoImage_488160(int a1, int a2);
nox_window* nox_gui_newEntryField_488500(nox_window* a1, int a2, int a3, int a4, int a5, int a6, int a7, wchar_t* a8);
int nox_xxx_wndEditProcPre_488710(int a1, unsigned int a2, wchar_t* a3, int a4);
int nox_xxx_wndEdit_488830(int a1);
int nox_xxx_wndEditDrawWithImage_488870(int a1, int a2);
int sub_488B60();
int sub_488BA0();
void nox_xxx_onChar_488BD0(unsigned short a1);
int nox_xxx_wndStaticDrawNoImage_488D00(nox_window* a1p, nox_window_data* a2p);
nox_window* nox_gui_newStaticText_489300(nox_window* a1, int a2, int a3, int a4, int a5, int a6, nox_window_data* a7p, nox_staticText_data* a8p);
int sub_4896E0();
int sub_489870();
int nox_xxx_checkSomeFlagsOnJoin_4899C0(nox_gui_server_ent_t* srv);
uint32_t* sub_489B80(int a1);
void sub_489DC0();
int nox_xxx_windowMplayFilterProc_489E70(int a1, int a2, int* a3, int a4);
int sub_489FB0();
int sub_489FF0(int a1, int a2, const void* a3);
int nox_xxx_setSomeFunc_48A210(int a1);
int sub_48B3E0(int a1);
void sub_48BD90(int a1);

#endif // NOX_PORT_GAME2_2
