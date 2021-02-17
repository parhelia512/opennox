#include "common__random.h"
#include "proto.h"

//----- (00552020) --------------------------------------------------------
int  nox_xxx_netRecv_552020(SOCKET s, char* buf, int len, int flags, struct sockaddr_in* from, int* fromlen) {
	int n = mix_recvfrom(s, buf, len, flags, from, fromlen);
	if (n != -1) {
		nox_net_struct_t* ns = nox_xxx_netStructByAddr_551E60(from);
		if (ns) {
			char key = ns->data_37[0];
			if (key) {
				nox_xxx_cryptXor_56FDD0(key, buf, n);
			}
		}
	}
	if (nox_common_gameFlags_check_40A5C0(1))
		return n;

	int r = nox_common_randomIntMinMax_415FF0(1, 99, "C:\\NoxPost\\src\\comw32\\comlib.c", 450);
	if (r >= *getMemIntPtr(0x5D4594, 2495940))
		return n;
	return 0;
}
