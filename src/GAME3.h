#ifndef NOX_PORT_GAME3
#define NOX_PORT_GAME3

#include "defs.h"


void sub_4A19D0();
int sub_4A1A40(int a1);
int sub_4A1A60();
int sub_4A1BD0();
int sub_4A1BE0(int a1);
int nox_client_drawGeneralCallback_4A2200();
BOOL sub_4A2490(int a1, int a2, int a3, int a4);
int nox_client_guiXxxDestroy_4A24A0();
BOOL sub_4A2560(_DWORD* a1, int a2);
int sub_4A25C0(_DWORD* a1, int* a2);
int sub_4A2610(int a1, _DWORD* a2, int* a3);
_DWORD* sub_4A2830(int a1, int a2, _DWORD* a3);
int sub_4A2890();
BOOL sub_4A28B0();
int sub_4A28C0(int a1);
int nox_xxx_wndListboxProcWithoutData10_4A28E0(_DWORD* a1, int a2, unsigned int a3, int a4);
int nox_xxx_wndListBox_4A2D10(int a1, int a2, int a3);
int nox_xxx_wndListboxProcWithData10_4A2DE0(int a1, int a2, unsigned int a3, int a4);
__int16* sub_4A3090(__int16* a1, int a2);
int nox_xxx_wndListboxProcPre_4A30D0(int a1, unsigned int a2, wchar_t* a3, int a4);
int nox_xxx_wndListBox_4A3A70(int a1);
int nox_xxx_wndListBoxAddLine_4A3AC0(wchar_t* a1, int a2, _DWORD* a3);
void nox_xxx_wndListboxInit_4A3C00(int a1, int a2);
int nox_xxx_wndListboxDrawNoImage_4A3C50(_DWORD* a1, int a2);
int nox_xxx_wndListboxDrawWithImage_4A3FC0(_DWORD* a1, int a2);
int sub_4A4800(int a1);
int nox_game_showSelClass_4A4840();
int sub_4A4970();
int sub_4A49A0();
int sub_4A49D0(int yTop, int a2);
int sub_4A4CB0(const void* a1, const void* a2);
int sub_4A50A0();
int sub_4A50D0();
int sub_4A5690(_BYTE* a1);
void* nox_xxx_clearAdminFlag();
void nox_xxx_setAdminFlag();
int sub_4A5C70();
int sub_4A5E90();
unsigned __int8* sub_4A61E0(_DWORD* a1, int a2, unsigned __int8* a3);
void sub_4A62B0();
int sub_4A6890();
int sub_4A6B50(wchar_t* a1);
int sub_4A6C90();
int sub_4A6D20(int a1, int a2);
int sub_4A6DC0(_DWORD* a1, int a2);
int sub_4A7270(int a1, int a2, unsigned int a3, int a4);
_DWORD* sub_4A72D0(unsigned __int16 a1);
int sub_4A7330(int a1, int a2, int* a3, unsigned int a4);
_DWORD* sub_4A7530(unsigned __int16 a1);
int sub_4A7580(int a1, int a2);
int sub_4A7A60(int a1);
int sub_4A7A70(int a1);
int sub_4A7A80(const char* a1);
int sub_4A7AC0(const char* a1);
int sub_4A7B00(const char* a1);
int sub_4A7B40(char* a1);
int sub_4A7BA0(char* a1);
int sub_4A7BC0(const char* a1);
int sub_4A7C00(const char* a1);
int sub_4A7C40(char* a1);
int sub_4A7C60(char* a1);
int sub_4A7CE0(char* a1);
int sub_4A7D00(const char* a1);
int sub_4A7D50(char* a1);
int sub_4A7D70(char* a1);
int sub_4A7DF0(char* a1);
char* sub_4A7E70();
char* sub_4A7E80();
char* sub_4A7E90();
int sub_4A7EA0();
int sub_4A7EB0();
char* sub_4A7EC0();
char* sub_4A7ED0();
int sub_4A7EE0();
char* sub_4A7EF0();
int sub_4A7F00();
char* sub_4A7F10();
int sub_4A7F20();
int sub_4A7F30();
int sub_4A7F40(int a1);
int nox_xxx_wndButtonProc_4A7F50(nox_window* win, int ev, int a3, int a4);
int nox_xxx_wndButtonDrawNoImg_4A81D0(int a1, int a2);
int nox_xxx_wndButtonInit_4A8340(int a1);
int nox_xxx_wndButtonDraw_4A8380(_DWORD* a1, int a2);
int nox_xxx_wndRadioButtonProc_4A84E0(_DWORD* a1, int a2, int a3, int a4);
int nox_xxx_wndRadioButtonSetAllFn_4A87E0(int a1);
int nox_xxx_wndRadioButtonDrawNoImg_4A8820(int a1, int a2);
int nox_xxx_wndRadioButtonDraw_4A8A20(int a1, int a2);
int nox_xxx_wndCheckBoxProc_4A8C00(int a1, int a2, int a3, int a4);
int nox_xxx_wndCheckBoxInit_4A8E60(int a1);
int nox_xxx_wndDrawCheckBoxNoImg_4A8EA0(int a1, int a2);
int nox_xxx_wndDrawCheckBox_4A9050(_DWORD* a1, int a2);
nox_window* nox_gui_newButtonOrCheckbox_4A91A0(nox_window* parent, int a2, int a3, int a4, int a5, int a6, nox_window_data* draw);
int nox_xxx_wndButtonProcPre_4A9250(int a1, int a2, wchar_t* a3, int a4);
int nox_xxx_wndCheckboxProcMB_4A92C0(int a1, int a2, wchar_t* a3, int a4);
nox_window* nox_gui_newRadioButton_4A9330(nox_window* parent, int a2, int a3, int a4, int a5, int a6, nox_window_data* draw, nox_radioButton_data* data);
int nox_xxx_wndRadioButtonProcPre_4A93C0(int a1, int a2, wchar_t* a3, int a4);
int nox_xxx_loadDefColor_4A94A0();
FILE* nox_xxx_loadPal_4A96C0_video_read_palette(char* a1);
int sub_4A9A10();
int sub_4A9A30(unsigned __int8 a1, unsigned __int8 a2, unsigned __int8 a3);
__int16 sub_4A9B20(int a1);
__int16 sub_4A9B70(__int16* a1);
int nox_xxx_unused_4A9BC0(unsigned __int8 a1, unsigned __int8 a2, unsigned __int8 a3);
int sub_4A9C50(int a1);
int nox_xxx_compassGenStrings_4A9C80();
int nox_game_showOnlineOrLAN_413800();
int sub_4AA450();
int sub_4AA490();
int nox_game_showOptions_4AA6B0();
int sub_4AA9C0();
int sub_4AAA10();
_DWORD* sub_4AAA70();
int sub_4AABE0(int a1, int a2, int* a3, int a4);
int sub_4AB0C0();
int nox_game_rollNoxLogoAndStart_4AB0F0();
int nox_game_rollIntroAndStart_4AB170();
int nox_game_rollLogoAndStart_4AB1F0();
int sub_4AB260();
int sub_4AB340(int a1, int a2, int a3, int a4);
int sub_4AB390(int a1, int a2, int* a3, int a4);
int sub_4AB420(int* a1);
int sub_4AB470();
int sub_4AB4A0(int a1);
int sub_4AB4D0(int a1);
void nox_xxx_mapSetDownloadInProgress_4AB560(int a1);
void nox_xxx_mapSetDownloadOK_4AB570(int a1);
LPCSTR sub_4AB580();
void nox_xxx_gameDownloadMap_4AB5E0();
void nox_xxx_mapDeleteFile_4AB720();
void nox_xxx_netMapDownloadPart_4AB7C0(unsigned __int16 a1, void* a2, size_t a3);
int sub_4AB9B0(char* a1);
void nox_xxx_cliCancelMapDownload_4ABA90();
int nox_xxx_mapDownloadStart_4ABAD0(char* a1, unsigned int a2);
int sub_4ABDA0(int a1, __int16 a2, __int16 a3, _DWORD* a4);
int nox_xxx_spriteLoadFromMap_4AC020(int thingInd, __int16 a2, _DWORD* a3);
char* nox_xxx_mapCliReadAll_4AC2B0(char* a1);
int nox_client_mapSpecialRWObjectData_4AC610();
int nox_xxx_clientLoadSomeObject_4AC6E0(unsigned __int16 a1);
int sub_4AC7B0(int a1);
int nox_xxx_colorLightClientLoad_4AC980(int a1);
int nox_xxx_cliLoadTeamBase_4ACE00(int a1);
int sub_4ACEF0(int a1);
int sub_4AD040(int a1);
int sub_4AD570();
int nox_xxx_windowServerOptionsGeneralProc_4AD5D0(int a1, int a2, int* a3, int a4);
int sub_4AD820();
int sub_4AD9B0(int a1);
int sub_4ADA40();
int nox_game_initOptionsInGame_4ADAD0();
int sub_4ADEF0(_DWORD* a1, int a2);
int nox_xxx_windowOptionsProc_4ADF30(int a1, int a2, int* a3, int a4);
int sub_4AE3B0();
BOOL sub_4AE3D0();
int sub_4AE400();
int sub_4AE420();
int sub_4AE470();
int sub_4AE520();
void sub_4AE540();
int sub_4AE580(int a1);
int sub_4AE5C0(int a1);
BOOL sub_4AE5E0(int a1, int a2);
int sub_4AE6F0(int a1, int a2, int a3, int a4, int a5);
void sub_4AEBD0();
int sub_4AEC20(int a1, int a2);
int sub_4AEDA0(int* a1, int* a2, int a3, int a4);
void* sub_4AEDF0();
__int64 sub_4AEE30();
int sub_4AEE80();
unsigned __int8* nox_xxx_ParticleFxT0_4AEEA0(unsigned __int8** a1);
char* sub_4AEED0(int* a1, int a2, int a3);
int nox_xxx_ParticleFxT1_4AEF80(int* a1);
int sub_4AF060(int a1);
unsigned __int8* nox_xxx_ParticleFxT2_4AF0F0(unsigned __int8** a1);
int sub_4AF200(int a1);
unsigned __int8* nox_xxx_ParticleFxT3_4AF2A0(unsigned __int8** a1);
void sub_4AF2D0(int* a1, int a2, int a3);
_DWORD* nox_xxx_ParticleFxT4_4AF3D0(_DWORD* a1);
int sub_4AF400(int a1, int a2, int a3);
char* nox_xxx_ParticleFxT5_4AF450(int* a1);
int sub_4AF4C0(int a1);
char* nox_xxx_ParticleFxT6_4AF5A0(int* a1);
int nox_xxx_partFx_4AF600(int a1);
unsigned __int8* sub_4AF890(unsigned __int8** a1);
int sub_4AF8C0(_DWORD* a1);
int sub_4AF8D0();
void sub_4AF950();
char* nox_xxx_drawPartFx2_4AF990(int a1, int a2, int a3, int a4);
void sub_4AFA40(int a1);
char* sub_4AFA70(int a1);
int sub_4AFAF0(_DWORD* a1, int a2);
int sub_4AFB00(int a1, int a2);
int sub_4AFB10(int a1, int a2);
int sub_4AFB40(int a1, int a2);
int sub_4AFB50(int a1, int a2);
int sub_4AFB60(int a1, int a2);
int sub_4AFB70(int a1, int a2);
int sub_4AFB80(int a1, int a2);
int sub_4AFB90(int a1, int a2);
int sub_4AFBA0(int a1, int a2);
_DWORD* sub_4AFBB0(_DWORD* a1, int a2, int a3, int a4);
int sub_4AFBD0(int a1, int a2);
int sub_4AFBE0(int a1, int a2);
int sub_4AFBF0(int a1, int a2);
_DWORD* sub_4AFC00(_DWORD* a1, int a2, int a3, int a4);
int sub_4AFC20(int a1, int a2);
int sub_4AFC30(int a1, int a2);
int sub_4AFC40(int a1, int a2);
int sub_4AFC50(int a1, int a2);
int sub_4AFC60(int a1, int a2, int a3);
void sub_4AFC90(_DWORD* a1, int a2, int a3, int a4);
_DWORD* sub_4AFCD0(_DWORD* a1);
int nox_xxx_registerParticleFx_4AFCF0(int a1, int a2, int a3, int a4);
void sub_4AFD40();
int nox_xxx_partfxLoadParticle_4AFE20(_DWORD* a1, CHAR* a2);
int nox_xxx_drawParticlefx_4AFEB0(int* a1);
int sub_4B0020(_DWORD* a1);
void sub_4B01A0(int a1);
char* nox_xxx_partfxAllocSmth_4B01B0();
int sub_4B0220(size_t a1);
int sub_4B02D0();
void sub_4B0610(int a1);
int sub_4B0650();
void sub_4B0660();
int* sub_4B0680(unsigned __int8 a1, unsigned __int8 a2);
int sub_4B07D0(LPVOID lpMem);
void nox_video_drawImageAt2_4B0820(void* a1, int x, int y);
int sub_4B0870(int* a1);
void* nox_video_getImagePixdataInline_4B0B20(nox_video_bag_image_t* img);
int nox_video_assignCircleDrawFuncs_4B0B30();
void nox_video_drawCircle_4B0B90(int a1, int a2, int a3);
int sub_4B0BC0(int a1, int a2, int a3);
char nox_video_drawCircle8Opaque_4B0D30(int a1, int a2, int a3);
int sub_4B0F50(int a1, int a2, int a3);
__int16 nox_video_drawCircle16Opaque_4B1380(int a1, int a2, int a3);
int sub_4B15E0(int a1, int a2, int a3);
int nox_video_drawCircle8Alpha_4B1A60(int a1, int a2, int a3);
int sub_4B1E30(int a1, int a2, int a3);
_WORD* nox_video_drawCircle16Alpha_4B2480(int a1, int a2, int a3);
int sub_4B3450(int a1, int a2, int a3);
int sub_4B4860(int a1, int a2, int a3, int a4);
int nox_xxx_wndScrollBoxDraw_4B4BA0(int a1, int a2, unsigned int a3, int a4);
nox_window* nox_gui_newSlider_4B4EE0(int a1, int a2, int a3, int a4, int a5, int a6, _DWORD* a7, float* a8);
int sub_4B5010(int a1, unsigned int a2, int a3, int a4);
int sub_4B51A0(int a1);
int sub_4B51E0(int a1, int a2);
int sub_4B52C0(int a1, int a2);
int nox_xxx_wndScrollBoxProc_4B5320(int a1, unsigned int a2, int a3, unsigned int a4);
int nox_xxx_wndScrollBoxSetAllFn_4B5500(int a1);
int nox_xxx_wndScrollBoxDraw_4B5540(int a1, int a2);
int nox_xxx_wndScrollBoxDraw_4B5620(_DWORD* a1, int a2);
int nox_xxx_wndScrollBoxButtonCreate_4B5640(int a1, int a2, int a3);
void sub_4B5700(int a1, int a2, int a3, int a4, int a5, int a6);
void* sub_4B5990();
int sub_4B5AB0(int a1, int a2, int* a3, int a4);
void sub_4B5BF0();
int sub_4B5CD0();
int sub_4B63B0(int2* a1, int2* a2);
int sub_4B64C0();
void sub_4B6720(int2* a1, int a2, int a3, char a4);
int sub_4B6880(_DWORD* a1, int a2, int a3, int a4);
int sub_4B6970(_DWORD* a1, nox_drawable* dr, int a3, int a4);
__int16 sub_4B69F0(int a1);
int sub_4B6B80(int* a1, nox_drawable* dr, int a3);
int sub_4B71A0(_DWORD* a1, int a2);
_DWORD* nox_xxx_netHandleSummonPacket_4B7C40(__int16 a1, unsigned __int16* a2, unsigned __int16 a3, unsigned __int8 a4, __int16 a5);
void sub_4B7EE0(__int16 a1);
int nox_xxx_spriteShieldLoad_4B7F90();
_DWORD* nox_xxx_fxShield_4B8090(unsigned int a1, int a2);
void nox_xxx_spriteScanForShield_4B81E0(int a1, int a2);
int nox_xxx_gameDeleteSpiningCrownSkull_4B8220();
__int16 sub_4B8960(int* a1, nox_drawable* dr, int a3, _DWORD* a4, int a5, int a6);
_DWORD* sub_4B8CA0(_DWORD* a1, char* a2);
__int16 sub_4B8D40(int* a1, nox_drawable* dr, int a3, _DWORD* a4, int a5, int a6);
_DWORD* sub_4B8E10(_DWORD* a1, char* a2);
void nox_xxx_drawOtherPlayerHP_4B8EB0(_DWORD* a1, nox_drawable* dr, unsigned __int16 a3, char a4);
int sub_4B8FA0(nox_drawable* dr, int* a2, int* a3);

_DWORD* nox_xxx_partfxSwitch_4AF690(_DWORD* a1, void(* a2)(_DWORD*, _DWORD*, int));

#endif // NOX_PORT_GAME3
