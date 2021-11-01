#include "MixPatch.h"
#include "GAME1.h"
#include "GAME3_3.h"
#include "GameEx.h"

#ifndef NOX_CGO
short asc_9800B0[] = {
	0,   0,   0,   0,   0,   0,   49,  33,  0,   50,  64,  0,   51,  35,  0,   52,  36,  0,   53,  37,  0,   54,  94,
	0,   55,  38,  0,   56,  42,  0,   57,  40,  0,   48,  41,  0,   45,  95,  0,   61,  43,  0,   8,   8,   0,   9,
	9,   0,   233, 201, 0,   246, 214, 0,   243, 211, 0,   234, 202, 0,   229, 197, 0,   237, 205, 0,   227, 195, 0,
	248, 216, 0,   249, 217, 0,   231, 199, 0,   245, 213, 0,   250, 218, 0,   10,  10,  0,   0,   0,   0,   244, 212,
	0,   251, 219, 0,   226, 194, 0,   224, 192, 0,   239, 207, 0,   240, 208, 0,   238, 206, 0,   235, 203, 0,   228,
	196, 0,   230, 198, 0,   253, 221, 0,   96,  126, 0,   0,   0,   0,   92,  124, 0,   255, 223, 0,   247, 215, 0,
	241, 209, 0,   236, 204, 0,   232, 200, 0,   242, 210, 0,   252, 220, 0,   225, 193, 0,   254, 222, 0,   47,  63,
	0,   0,   0,   0,   42,  42,  0,   0,   0,   0,   32,  32,  0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,
	0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,
	0,   0,   0,   0,   0,   0,   55,  55,  0,   56,  56,  0,   57,  57,  0,   45,  45,  0,   52,  52,  0,   53,  53,
	0,   54,  54,  0,   43,  43,  0,   49,  49,  0,   50,  50,  0,   51,  51,  0,   48,  48,  0,   46,  46,  0,   0,
	0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   49,  33,  0,   50,  64,  0,   51,
	35,  0,   52,  36,  0,   53,  37,  0,   54,  94,  0,   55,  38,  0,   56,  42,  0,   57,  40,  0,   48,  41,  0,
	45,  95,  0,   61,  43,  0,   8,   8,   0,   9,   9,   0,   233, 201, 0,   246, 214, 0,   243, 211, 0,   234, 202,
	0,   229, 197, 0,   237, 205, 0,   227, 195, 0,   248, 216, 0,   249, 217, 0,   231, 199, 0,   245, 213, 0,   250,
	218, 0,   10,  10,  0,   0,   0,   0,   244, 212, 0,   251, 219, 0,   226, 194, 0,   224, 192, 0,   239, 207, 0,
	240, 208, 0,   238, 206, 0,   235, 203, 0,   228, 196, 0,   230, 198, 0,   253, 221, 0,   96,  126, 0,   0,   0,
	0,   92,  124, 0,   255, 223, 0,   247, 215, 0,   241, 209, 0,   236, 204, 0,   232, 200, 0,   242, 210, 0,   252,
	220, 0,   225, 193, 0,   254, 222, 0,   47,  63,  0,   0,   0,   0,   42,  42,  0,   0,   0,   0,   32,  32,  0,
	0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,
	0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   55,  55,  0,   56,  56,  0,   57,
	57,  0,   45,  45,  0,   52,  52,  0,   53,  53,  0,   54,  54,  0,   43,  43,  0,   49,  49,  0,   50,  50,  0,
	51,  51,  0,   48,  48,  0,   46,  46,  0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0};

void init_data_mix() { GameEx_DllMain(1); }

void OnLibraryNotice_stub(int a1, ...) { /* TODO: STUB */ }
#endif // NOX_CGO

void sub_980523(nox_object_t* unit) {
	if (!unit)
		return;
	for (nox_object_t* it = unit->field_126; it; it = it->field_124) {
		if ((it->obj_class & 0x2000000) && (it->field_4 & 0x100)) {
			if (nox_xxx_unitArmorInventoryEquipFlags_415C70(it) & 0x3000000) {
				// TODO: it appears that it reuses some other field; this might make the game unstable
				*(uint32_t*)(*(uint32_t*)(((uint32_t)unit->field_187) + 276) + 2500) = it;
			}
		}
	}
}

nox_object_t* sub_9805EB(nox_object_t* unit) {
	if (!unit) {
		return 0;
	}
	for (nox_object_t* it = unit->field_126; it; it = it->field_124) {
		if ((it->obj_class & 0x2000000) && (it->field_4 == 16)) {
			return it;
		}
	}
	return 0;
}

int mix_recvfrom(nox_socket_t s, char* buf, int len, struct nox_net_sockaddr* from) {
	int result = nox_net_recvfrom(s, buf, len, from);
	if (*(uint16_t*)buf != 0xF13A) { // extension packet code
		return result;
	}
	return MixRecvFromReplacer(s, buf, len, from);
}
