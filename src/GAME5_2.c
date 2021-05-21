#include "client__system__parsecmd.h"
#include "common__random.h"

#include "client__gui__guiinv.h"
#include "client__shell__mainmenu.h"
#include "common__system__group.h"
#include "comw32__comlib.h"

#include "common/fs/nox_fs.h"
#include "nox_net.h"
#include "proto.h"

#include <float.h>
#include <time.h>

extern _DWORD dword_5d4594_2496988;
extern _DWORD dword_5d4594_2516352;
extern _DWORD dword_5d4594_3843632;
extern _DWORD dword_5d4594_2513924;
extern _DWORD dword_5d4594_2523888;
extern _DWORD dword_5d4594_2496472;
extern _DWORD dword_5d4594_2523904;
extern _DWORD dword_5d4594_2516380;
extern _DWORD dword_5d4594_2523804;
extern _DWORD dword_5d4594_2516372;
extern _DWORD dword_5d4594_2523764;
extern _DWORD dword_5d4594_2513932;
extern _DWORD dword_5d4594_2523760;
extern _DWORD nox_xxx_warriorMaxSpeed_587000_312796;
extern _DWORD nox_xxx_conjurerSpeed_587000_312812;
extern _DWORD nox_xxx_wizardSpeed_587000_312828;
extern _DWORD nox_xxx_wizardStrength_587000_312824;
extern _DWORD nox_xxx_warriorMaxStrength_587000_312792;
extern _DWORD dword_5d4594_2523776;
extern _DWORD nox_xxx_conjurerStrength_587000_312808;
extern void* nox_alloc_groupInfo_2523892;
extern _DWORD dword_5d4594_2516356;
extern _DWORD dword_587000_311372;
extern void* nox_alloc_debugData_2523908;
extern _DWORD dword_5d4594_2513916;
extern void* nox_alloc_itemGroupElem_2523896;
extern _DWORD nox_xxx_warriorMaxMana_587000_312788;
extern _DWORD nox_xxx_warriorMaxHealth_587000_312784;
extern _DWORD nox_xxx_conjurerMaxHealth_587000_312800;
extern _DWORD nox_xxx_wizardMaxHealth_587000_312816;
extern _DWORD nox_xxx_conjurerMaxMana_587000_312804;
extern _DWORD dword_5d4594_3844304;
extern _DWORD nox_xxx_wizardMaximumMana_587000_312820;
extern _QWORD qword_581450_9544;
extern _DWORD dword_5d4594_2523912;
extern _DWORD dword_5d4594_2523780;
extern _DWORD dword_5d4594_2495920;
extern _DWORD dword_5d4594_2516344;
extern _DWORD dword_5d4594_2523900;
extern _DWORD dword_5d4594_2523756;
extern _DWORD dword_5d4594_2523752;
extern _DWORD dword_5d4594_2516328;
extern _DWORD dword_5d4594_2516348;
extern _DWORD dword_5d4594_2650652;
extern nox_alloc_class* nox_alloc_gQueue_3844300;
extern unsigned int nox_frame_xxx_2598000;

int (*nox_client_onLobbyServer_2513928)(const char*, uint16_t, const char*, const char*) = 0;

nox_net_struct_t* nox_net_struct_arr[NOX_NET_STRUCT_MAX] = {0};
nox_net_struct2_t nox_net_struct2_arr[NOX_NET_STRUCT_MAX] = {0};
nox_socket_t nox_xxx_sockLocalBroadcast_2513920 = 0;

//----- (005528B0) --------------------------------------------------------
int  nox_xxx_netSendReadPacket_5528B0(unsigned int a1, char a2) {
	unsigned int v2;     // ecx
	unsigned int v5;     // eax
	int v7;              // edi
	DWORD v10;           // eax
	int v11;             // eax
	unsigned int v12;    // eax
	char* v13;           // eax
	int v14;             // ebp
	unsigned int v15;    // [esp+10h] [ebp-Ch]
	unsigned int v18;    // [esp+20h] [ebp+4h]

	v2 = a1;
	if (a1 >= NOX_NET_STRUCT_MAX)
		return -3;
	nox_net_struct_t* ns = nox_net_struct_arr[a1];
	if (!ns)
		return -3;
	if (ns->id == -1) {
		if (ns->field_19) {
			SetEvent(ns->field_34);
			return 0;
		}
		v5 = 0;
		v18 = NOX_NET_STRUCT_MAX;
	} else {
		v5 = a1;
		v18 = a1 + 1;
		v2 = ns->id;
	}
	v15 = v2;
	v7 = v18;
	for (int j = v5; j < v18; j++) {
		nox_net_struct_t* ns2 = nox_net_struct_arr[j];
		if (!ns2 || ns2->id != v15) {
			continue;
		}
		nox_xxx_netSend_5552D0(j, 0, 0);
		v10 = WaitForSingleObject(ns2->mutex_yyy, 0x3E8u);
		v7 = v10;
		if (v10 == -1 || v10 == 258)
			return -16;
		if (!(a2 & 1)) {
			v11 = ns2->func_xxx(j, ns2->data_2_xxx, (int)(ns2->data_2_end) - (int)(ns2->data_2_xxx), ns2->data_3);
			v7 = v11;
			if (v11 > 0) {
				v12 = v11 + (_DWORD)(ns2->data_2_xxx);
				if (v12 < ns2->data_2_end)
					ns2->data_2_xxx = v12;
			}
		}
		v13 = ns2->data_2_base;
		v14 = (_DWORD)(ns2->data_2_xxx) - (_DWORD)v13;
		if (v14 > 2) {
			v7 = nox_xxx_sendto_551F90(ns->sock, v13, v14, &ns2->addr);
			if (v7 == -1)
				return -1;
			sub_553F40(v14, 1);
			nox_xxx_netCountData_554030(v14, j);
			ns2->data_2_xxx = ns2->data_2_yyy;
		}
		if (!ReleaseMutex(ns2->mutex_yyy))
			ReleaseMutex(ns2->mutex_yyy);
	}
	return v7;
}

//----- (00552A80) --------------------------------------------------------
int  nox_xxx_servNetInitialPackets_552A80(unsigned int id, char flags) {
	struct nox_net_sockaddr_in to;
	char buf[256];

	if (id >= NOX_NET_STRUCT_MAX) {
		return -3;
	}
	nox_net_struct_t* ns = nox_net_struct_arr[id];
	if (!ns) {
		return -3;
	}
	nox_net_struct_t* ns2 = ns;
	unsigned int argp = 0;
	if (flags & 1) {
		if (nox_net_recv_available(ns->sock, &argp) == -1) {
			return -1;
		}
		if (!argp) {
			return -1;
		}
	} else {
		argp = 1;
	}
	char v26 = 1;
	while (1) {
		int n = nox_xxx_netRecv_552020(ns->sock, ns->data_1_xxx, (int)(ns->data_1_end) - (int)(ns->data_1_xxx), &to);
		if (n == -1) {
			return -1;
		}
		sub_553FC0(n, 1);
		if (n < 3) {
			ns->data_1_yyy = ns->data_1_base;
			ns->data_1_xxx = ns->data_1_base;
			if (!(flags & 1) || (flags & 4)) {
				return n;
			}
			argp = 0;
			if (nox_net_recv_available(ns->sock, &argp) == -1) {
				return -1;
			} else if (!argp) {
				return n;
			}
			continue;
		}
		unsigned char* v7 = ns->data_1_yyy;
		ns->data_1_xxx = &ns->data_1_xxx[n];
		unsigned int  id2 = v7[0];
		unsigned char v9  = v7[1];
		unsigned char op = v7[2];
		if (debugNet)
			printf("servNetInitialPackets: op=%d\n", op);
		if (op == 12) {
			// received a lobby info request from the client
			if (!nox_xxx_check_flag_aaa_43AF70()) {
				// send server info packet
				n = nox_server_makeServerInfoPacket_554040(ns->data_1_yyy, (int)(ns->data_1_xxx) - (int)(ns->data_1_yyy), buf);
				if (n > 0) {
					n = nox_net_sendto(ns->sock, buf, n, &to);
					sub_553F40(n, 1);
				}
			}
			ns->data_1_yyy = ns->data_1_base;
			ns->data_1_xxx = ns->data_1_base;
			if (!(flags & 1) || (flags & 4)) {
				return n;
			}
			argp = 0;
			if (nox_net_recv_available(ns->sock, &argp) == -1) {
				return -1;
			} else if (!argp) {
				return n;
			}
			continue;
		}
		if (op >= 14 && op <= 20) {
			v26 = 1;
		} else {
			if (id2 == 255) {
				if (v26 != 1) {
					goto LABEL_48;
				}
			} else {
				v26 = 0;
				if (!sub_551E00(id2 & 127, (int)&to)) {
					goto LABEL_48;
				}
				v26 = 1;
			}
			if (ns->id == -1) {
				ns2 = nox_net_struct_arr[id2 & 127];
			}
			if ((id2 & NOX_NET_STRUCT_MAX) == 0) {
				if (!ns2) {
					goto LABEL_48;
				}
				if (v9 != ns2->field_28_1) {
					sub_5551F0(id2, v9, 1);
					sub_555360(id2, v9, 1);
					ns2->field_28_1 = v9;
					bool v20 = 0;
					if (nox_xxx_netRead2Xxx_551EB0(id, id2, v9, ns->data_1_yyy, (int)(ns->data_1_xxx) - (int)(ns->data_1_yyy)) == 1)
						v20 = 0;
					else
						v20 = 1;
					buf[0] = 38;
					buf[1] = ns2->field_28_1;
					ns->func_yyy(id2, buf, 2, ns2->data_3);
					if (!v20) {
						goto LABEL_48;
					}
				}
			} else if (id2 == 255) {
				if (op == 0) {
					memcpy(buf, &to, sizeof(to));
					n = nox_xxx_netBigSwitch_553210(id, ns->data_1_yyy, (int)(ns->data_1_xxx) - (int)(ns->data_1_yyy), (int)buf);
					if (n > 0) {
						n = nox_xxx_sendto_551F90(ns->sock, buf, n, &to);
						sub_553F40(n, 1);
					}
					goto LABEL_48;
				}
			} else {
				ns->data_1_yyy[0] &= 127;
				id2 = ns->data_1_yyy[0];
				if (!ns2) {
					goto LABEL_48;
				}
				if (ns2->data_2_base[1] != v9) {
					goto LABEL_48;
				}
				ns2->data_2_base[1]++;
				if (nox_xxx_netRead2Xxx_551EB0(id, id2, v9, ns->data_1_yyy, (int)(ns->data_1_xxx) - (int)(ns->data_1_yyy)) == 1) {
					goto LABEL_48;
				}
			}
		}
		if (op < 32) {
			memcpy(buf, &to, sizeof(to));
			n = nox_xxx_netBigSwitch_553210(id, ns->data_1_yyy, (int)(ns->data_1_xxx) - (int)(ns->data_1_yyy), (int)buf);
			if (n > 0) {
				n = nox_xxx_sendto_551F90(ns->sock, buf, n, &to);
				sub_553F40(n, 1);
			}
		} else {
			if (ns2 && !(flags & 2)) {
				ns->func_yyy(id2, &ns->data_1_yyy[2], n - 2, ns2->data_3);
			}
		}
	LABEL_48:
		ns->data_1_yyy = ns->data_1_base;
		ns->data_1_xxx = ns->data_1_base;
		if (!(flags & 1) || (flags & 4)) {
			return n;
		}
		argp = 0;
		if (nox_net_recv_available(ns->sock, &argp) == -1) {
			return -1;
		} else if (!argp) {
			return n;
		}
	}
	// unreachable
}

//----- (00552E70) --------------------------------------------------------
int  sub_552E70(unsigned int a1) {
	unsigned int v3;     // esi
	unsigned int v4;     // edi
	unsigned int v5;     // ebp
	char v9[5];          // [esp+0h] [ebp-8h]

	if (a1 >= NOX_NET_STRUCT_MAX)
		return -3;
	nox_net_struct_t* ns = nox_net_struct_arr[a1];
	if (!ns)
		return -3;
	if (ns->id == -1) {
		v3 = 0;
		v4 = NOX_NET_STRUCT_MAX;
		v5 = a1;
	} else {
		v5 = ns->id;
		v3 = a1;
		v4 = a1 + 1;
	}
	v9[0] = 6;
	for (; v3 < v4; v3++) {
		nox_net_struct_t* ns2 = nox_net_struct_arr[v3];
		if (ns2 && ns2->id == v5) {
			int v8 = dword_5d4594_2495920;
			ns2->field_22 = dword_5d4594_2495920;
			ns2->field_23 = dword_5d4594_2495920;
			*(_DWORD*)&v9[1] = v8;
			nox_xxx_netSendSock_552640(v3, v9, 5, NOX_NET_SEND_FLAG2);
		}
	}
	return 0;
}

//----- (00552F20) --------------------------------------------------------
int  sub_552F20(unsigned int a1) {
	unsigned int v3;     // esi
	unsigned int v4;     // edi
	unsigned int v5;     // ebp
	char v9[256];        // [esp+0h] [ebp-100h]

	if (a1 >= NOX_NET_STRUCT_MAX)
		return -3;
	nox_net_struct_t* ns = nox_net_struct_arr[a1];
	if (!ns)
		return -3;
	if (ns->id == -1) {
		v3 = 0;
		v4 = NOX_NET_STRUCT_MAX;
		v5 = a1;
	} else {
		v5 = ns->id;
		v3 = a1;
		v4 = a1 + 1;
	}
	v9[0] = 5;
	for (; v3 < v4; v3++) {
		nox_net_struct_t* ns2 = nox_net_struct_arr[v3];
		if (ns2 && ns2->id == v5) {
			int v8 = ns2->field_25 + 1;
			ns2->field_25 = v8;
			ns2->field_26 = dword_5d4594_2495920;
			*(_DWORD*)&v9[1] = v8;
			nox_xxx_netSendSock_552640(v3, v9, 256, NOX_NET_SEND_FLAG2);
		}
	}
	return 0;
}

//----- (00552FD0) --------------------------------------------------------
unsigned int sub_552FD0(int a1) {
	unsigned int argp = 0;
	if (nox_net_recv_available(nox_net_struct_arr[a1]->sock, &argp)) {
		return -1;
	}
	return argp;
}

//----- (00553D60) --------------------------------------------------------
int nox_xxx_netHandlerDefXxx_553D60(unsigned int a1, char* a2, int a3, void* a4) { return 0; }
//----- (00553D70) --------------------------------------------------------
int nox_xxx_netHandlerDefYyy_553D70(unsigned int a1, char* a2, int a3, void* a4) { return 0; }
//----- (00553000) --------------------------------------------------------
nox_net_struct_t* nox_xxx_makeNewNetStruct_553000(nox_net_struct_arg_t* arg) {
	nox_net_struct_t* net = malloc(sizeof(nox_net_struct_t));
	if (!net) {
		return 0;
	}
	memset(net, 0, sizeof(nox_net_struct_t));

	HANDLE my = CreateMutexA(0, 0, 0);
	if (!my) {
		free(net);
		return 0;
	}
	net->mutex_yyy = my;

	HANDLE mx = CreateMutexA(0, 0, 0);
	if (!mx) {
		free(net);
		return 0;
	}
	net->mutex_xxx = mx;
	if (arg->data_3_size > 0) {
		void* p = malloc(arg->data_3_size);
		if (!p) {
			free(net);
			return 0;
		}
		memset(p, 0, arg->data_3_size);
		net->data_3 = p;
	}
	int dsz = arg->data_size;
	if (dsz > 0) {
		LOBYTE(dsz) = dsz & 0xFC;
		arg->data_size = dsz;
	} else {
		arg->data_size = 1024;
	}
	char* data1 = (char*)malloc(arg->data_size + 2);
	if (!data1) {
		free(net->data_3);
		free(net);
		return 0;
	}
	net->data_1_base = data1;
	net->data_1_xxx = &data1[0];
	net->data_1_yyy = &data1[0];
	net->data_1_end = &data1[arg->data_size + 2];

	char* data2 = (char*)malloc(arg->data_size + 2);
	if (!data2) {
		free(net->data_1_base);
		free(net->data_3);
		free(net);
		return 0;
	}
	memset(data2, 0, arg->data_size + 2);
	data2[0] = -1;
	net->data_2_base = data2;
	net->data_2_xxx = &data2[2];
	net->data_2_yyy = &data2[2];
	net->data_2_end = &data2[arg->data_size + 2];

	net->field_20 = arg->field_4;
	if (arg->func_xxx)
		net->func_xxx = arg->func_xxx;
	else
		net->func_xxx = nox_xxx_netHandlerDefXxx_553D60;
	if (arg->func_yyy)
		net->func_yyy = arg->func_yyy;
	else
		net->func_yyy = nox_xxx_netHandlerDefYyy_553D70;
	net->field_28_1 = -1;
	net->xor_key = 0;
	return net;
}

//----- (005531C0) --------------------------------------------------------
void  nox_xxx_netStructFree_5531C0(nox_net_struct_t* ns) {
	if (ns->data_3)
		free(ns->data_3);
	free(ns->data_1_base);
	free(ns->data_2_base);
	CloseHandle(ns->mutex_yyy);
	CloseHandle(ns->mutex_xxx);
	free(ns);
}

//----- (00553210) --------------------------------------------------------
int  nox_xxx_netBigSwitch_553210(unsigned int id, unsigned char* packet, int packetSz, int a4) {
	int out = a4;
	// dhexdump(packet, packetSz);
	int pid = (char)packet[0];
	char p1 = packet[1];

	unsigned char* packetEnd = &packet[packetSz];

	unsigned __int8 v74[8];
	*(_DWORD*)&v74[0] = *(_DWORD*)(a4 + 0);
	*(_DWORD*)&v74[4] = *(_DWORD*)(a4 + 4);
	int v75 = *(_DWORD*)(a4 + 8);
	int v76 = *(_DWORD*)(a4 + 12);

	if (packetSz <= 2) {
		return 0;
	}
	unsigned char* packetCur = &packet[2];

	if (id > NOX_NET_STRUCT_MAX) {
		printf("nox_net_struct_arr overflow (1): %d\n", (int)(id));
		abort();
	}
	nox_net_struct_t* ns1 = nox_net_struct_arr[id];
	unsigned int pidb = pid;
	while (2) {
		int op = packetCur[0];
		packetCur = &packetCur[1];
		if (debugNet)
			printf("nox_xxx_netBigSwitch_553210: op=%d [%d]\n", op, packetSz);
		switch (op) {
		case 0:
			{
			if (nox_common_gameFlags_check_40A5C0(1) && nox_common_gameFlags_check_40A5C0(8))
				return 0;
			*(_BYTE*)(out + 0) = 0;
			*(_BYTE*)(out + 1) = p1;
			if (ns1->field_21 >= nox_xxx_servGetPlrLimit_409FA0() + (unsigned int)getMemByte(0x5D4594, 2500076) - 1) {
				*(_BYTE*)(out + 2) = 2;
				return 3;
			}
			if (pid != -1) {
				// pid in the request must be -1 (0xff); fail if it's not
				*(_BYTE*)(out + 2) = 2;
				return 3;
			}
			for (int i = 0; i < NOX_NET_STRUCT_MAX; i++) {
				nox_net_struct_t* ns9 = nox_net_struct_arr[i];
				if (!ns9) {
					pid = i;
					break;
				}
				if (*(_WORD*)&v74[2] == ns9->addr.sin_port && *(_DWORD*)&v74[4] == ns9->addr.sin_addr) {
					printf("%d %d\n", *(_WORD*)&v74[2], *(_DWORD*)&v74[4]);
					*(_BYTE*)(out + 2) = 4; // already joined?
					return 3;
				}
			}
			if (pid == -1) {
				*(_BYTE*)(out + 2) = 2;
				return 3;
			}
			nox_net_struct_arg_t narg;
			memset(&narg, 0, sizeof(nox_net_struct_arg_t));
			narg.data_3_size = 4;
			narg.data_size = (int)(ns1->data_2_end) - (int)(ns1->data_2_base);
			nox_net_struct_t* ns10 = nox_xxx_makeNewNetStruct_553000(&narg);
			nox_net_struct_arr[pid] = ns10;
			if (!ns10) {
				// cannot allocate - fail
				*(_BYTE*)(out + 2) = 2;
				return 3;
			}
			ns1->field_21++;
			ns10->data_2_base[0] = id;
			int v62 = ns10->data_2_base;
			char v63 = *(_BYTE*)(v62 + 1);
			if (v63 == p1)
				*(_BYTE*)(v62 + 1) = v63 + 1;
			ns10->id = id;
			ns10->sock = ns1->sock;
			ns10->func_xxx = ns1->func_xxx;
			ns10->func_yyy = ns1->func_yyy;
			memset(getMemAt(0x5D4594, 32 * id + 2508788), 0, 32);
			*getMemU32Ptr(0x5D4594, 32 * id + 2508816) = 1;
			char key = nox_common_randomInt_415FA0(1, 255);
			if (nox_net_no_xor) {
				key = 0;
			}
			ns10->xor_key = 0; // send this packet without xor encoding

			int v66 = &ns10->addr;
			*(_QWORD*)v66 = *(_QWORD*)v74;
			*(_DWORD*)(v66 + 8) = v75;
			*(_DWORD*)(v66 + 12) = v76;

			*(_BYTE*)(out + 0) = 31;
			*(_BYTE*)(out + 1) = p1;
			*(_BYTE*)(out + 2) = 1;
			*(_DWORD*)(out + 3) = pid;
			*(_BYTE*)(out + 7) = key;
			char v67 = nox_xxx_netSendSock_552640(pid, out, 8, NOX_NET_SEND_NO_LOCK | NOX_NET_SEND_FLAG2);

			ns10->xor_key = key;
			ns10->field_38 = 1;
			ns10->data_39[0] = v67;
			ns10->field_40 = nox_frame_xxx_2598000;
			return 0;
			}
		case 1:
			{
			int v11 = *(_DWORD*)packetCur;
			_BYTE* v12 = ns1->data_2_base;
			_BYTE* v13 = packetCur + 4;
			ns1->id = v11;
			*v12 = v11;
			packetCur = v13 + 1;
			ns1->xor_key = (_BYTE)*v13;
			dword_5d4594_3844304 = 1;
			if ((unsigned int)packetCur >= packetEnd) {
				return 0;
			}
			break;
			}
		case 2:
			ns1->id = -18;
			dword_5d4594_3844304 = 1;
			if ((unsigned int)packetCur >= packetEnd) {
				return 0;
			}
			break;
		case 3: // ack?
			ns1->id = -12;
			dword_5d4594_3844304 = 1;
			if ((unsigned int)packetCur >= packetEnd) {
				return 0;
			}
			break;
		case 4:
			ns1->id = -13;
			dword_5d4594_3844304 = 1;
			if ((unsigned int)packetCur >= packetEnd) {
				return 0;
			}
			break;
		case 5:
			*(_BYTE*)(out + 0) = ns1->xor_key;
			*(_BYTE*)(out + 2) = 7;
			*(_DWORD*)(out + 3) = *(_DWORD*)packetCur;
			return 7;
		case 6:
			{
			if (pidb > NOX_NET_STRUCT_MAX) {
				printf("nox_net_struct_arr overflow (2): %d\n", (int)(pidb));
				abort();
			}
			nox_net_struct_t* ns2 = nox_net_struct_arr[pidb];
			nox_net_struct_t* ns3 = nox_net_struct_arr[pid];
			_BYTE a2b3 = 37;
			ns1->func_yyy(pid, &a2b3, 1, ns3->data_3);
			int v18 = 0;
			if (ns1->id == -1) {
				*(_BYTE*)(out + 0) = ns2->id;
				v18 = ns2->data_2_base;
			} else {
				*(_BYTE*)(out + 0) = ns1->id;
				v18 = ns1->data_2_base;
			}
			*(_BYTE*)(out + 1) = *(_BYTE*)(v18 + 1);
			*getMemU32Ptr(0x5D4594, 32 * pid + 2508792) = *(_DWORD*)packetCur;
			*(_BYTE*)(out + 2) = 8;
			*(_DWORD*)(out + 3) = *(_DWORD*)packetCur;
			return 7;
			}
		case 7:
			{
			if (pidb > NOX_NET_STRUCT_MAX) {
				printf("nox_net_struct_arr overflow (2): %d\n", (int)(pidb));
				abort();
			}
			nox_net_struct_t* ns4 = nox_net_struct_arr[pidb];
			if (!ns4->field_25)
				return 0;
			int v31 = dword_5d4594_2495920 - (int)(ns4->field_26) - (int)(ns4->field_24);
			int v32 = -1;
			if (v31 >= 1)
				v32 = 256000 / v31;
			*(_BYTE*)(out + 0) = 35;
			*(_DWORD*)(out + 1) = v32;
			if (ns1->id == -1)
				ns1->func_yyy(pid, out, 5, ns4->data_3);
			else
				ns1->func_yyy(pid, out, 5, ns1->data_3);
			return 0;
			}
		case 8:
			{
			if (pidb > NOX_NET_STRUCT_MAX) {
				printf("nox_net_struct_arr overflow (2): %d\n", (int)(pidb));
				abort();
			}
			nox_net_struct_t* ns5 = nox_net_struct_arr[pidb];
			if (*(_DWORD*)packetCur != ns5->field_22)
				return 0;
			ns5->field_24 = dword_5d4594_2495920 - ns5->field_23;
			*(_BYTE*)(out + 0) = 36;
			*(_DWORD*)(out + 1) = ns5->field_24;
			int v19 = 0;
			if (ns1->id == -1)
				v19 = ns5->data_3;
			else
				v19 = ns1->data_3;
			ns1->func_yyy(pid, out, 5, v19);
			*(_BYTE*)(out + 0) = ns1->data_2_base[0];
			*(_BYTE*)(out + 1) = ns5->data_2_base[1];
			*(_BYTE*)(out + 2) = 9;
			*(_DWORD*)(out + 3) = dword_5d4594_2495920;
			return 7;
			}
		case 9:
			{
			int v21 = 32 * pid;
			int v22 = *(_DWORD*)packetCur - *getMemU32Ptr(0x5D4594, 32 * pid + 2508792);
			if (v22 <= 0 || v22 >= 1000)
				return 0;
			int v23 = *getMemU32Ptr(0x5D4594, 32 * pid + 2508788);
			int v24 = v23 + 8 * pid;
			int v25 = 5;
			*getMemU32Ptr(0x5D4594, 4 * v24 + 2508796) = v22;
			int v26 = (v23 + 1) % 5;
			int v27 = v26;
			if (!v26) {
				unsigned __int8* v28 = getMemAt(0x5D4594, v21 + 2508796);
				do {
					int v29 = *(_DWORD*)v28;
					v28 += 4;
					v26 += v29;
					--v25;
				} while (v25);
				*getMemU32Ptr(0x5D4594, v21 + 2508816) = v26 / 5;
			}
			*getMemU32Ptr(0x5D4594, v21 + 2508788) = v27;
			return 0;
			}
		case 10:
			{
			if (pid == 255)
				return 0;
			nox_net_struct_t* ns6 = nox_net_struct_arr[pid];
			if (ns6->field_38 == 1)
				return 0;
			_BYTE a2b2 = 34;
			ns1->func_yyy(pid, &a2b2, 1, ns6->data_3);
			memset(getMemAt(0x5D4594, 32 * id + 2508788), 0, 0x20u);
			int* v69 = nox_xxx_findPlayerID_5541D0(pid);
			if (v69) {
				sub_425920((_DWORD**)v69);
				free(v69);
				--*getMemU8Ptr( 0x5D4594, 2500076);
			}
			nox_xxx_netStructReadPackets_5545B0(pid);
			return 0;
			}
		case 11:
			{
			nox_net_struct_t* ns7 = nox_net_struct_arr[pid];
			_BYTE a2b = 33;
			ns1->func_yyy(pid, &a2b, 1, ns7->data_3);
			sub_554A50(id);
			return 0;
			}
		case 14: // join game request?
			{
			int v43 = 0;
			char* v78 = sub_416640();
			nox_xxx_cliGamedataGet_416590(0);
			bool a4a = 0;
			*(_BYTE*)(out + 0) = 0;
			*(_BYTE*)(out + 1) = p1;
			if (!nox_xxx_check_flag_aaa_43AF70()) {
				char* v45 = nox_common_playerInfoGetFirst_416EA0();
				while (v45) {
					if (v45[2135] == packet[98]) {
						if (!strcmp(v45 + 2112, (const char*)packet + 56)) {
							*(_BYTE*)(out + 2) = 19;
							*(_BYTE*)(out + 3) = 12;
							return 4;
						}
						v43 = 0;
					}
					v45 = nox_common_playerInfoGetNext_416EE0((int)v45);
				}
			}
			if (*((_DWORD*)packet + 20) != 66458) {
				*(_BYTE*)(out + 2) = 19;
				*(_BYTE*)(out + 3) = 13;
				return 4;
			}
			if (ns1->field_21 >= (unsigned int)(nox_xxx_servGetPlrLimit_409FA0() - 1))
				a4a = 1;
			if (sub_40A740()) {
				int v46 = nox_xxx_countObserverPlayers_425BF0();
				if (!*((_DWORD*)packet + 21)) {
					if (v46 >= (unsigned __int8)v78[53]) {
						*(_BYTE*)(out + 2) = 19;
						*(_BYTE*)(out + 3) = 11;
						return 4;
					}
				} else if (sub_418AE0(*((_DWORD*)packet + 21))) {
					if (v46 > 0) {
						v43 = 1;
					}
				} else {
					if ((unsigned __int8)sub_417DE0() >= (unsigned __int8)v78[52]) {
						if (v46 >= (unsigned __int8)v78[53]) {
							*(_BYTE*)(out + 2) = 19;
							*(_BYTE*)(out + 3) = 11;
							return 4;
						}
					} else if (v46 > 0) {
						v43 = 1;
					}
				}
			}
			if (a4a) {
				if (!v43 || !*(_DWORD*)(v78 + 54)) {
					*(_BYTE*)(out + 2) = 19;
					*(_BYTE*)(out + 3) = 11;
					return 4;
				}
				for (char* i = nox_xxx_firstReplaceablePlayer_425C40(); i; i = nox_xxx_nextReplaceablePlayer_425C70((int)i)) {
					if (!nox_xxx_findPlayerID_5541D0((unsigned __int8)i[2064] + 1)) {
						nox_xxx_playerCallDisconnect_4DEAB0((unsigned __int8)i[2064], 4);
						_DWORD* v50 = malloc(0x10u);
						v50[3] = (unsigned __int8)i[2064] + 1;
						nox_common_list_append_4258E0((int)getMemAt(0x5D4594, 2495908), v50);
						++*getMemU8Ptr( 0x5D4594, 2500076);
						*(_BYTE*)(out + 2) = 21;
						return 3;
					}
				}
			}
			if (v78[100] & 0x10) {
				int* v48 = sub_4168E0();
				if (!v48) {
					*(_BYTE*)(out + 2) = 19;
					*(_BYTE*)(out + 3) = 4;
					return 4;
				}
				while (_nox_wcsicmp((const wchar_t*)v48 + 6, (const wchar_t*)(packet + 4))) {
					v48 = sub_4168F0(v48);
					if (!v48) {
						*(_BYTE*)(out + 2) = 19;
						*(_BYTE*)(out + 3) = 4;
						return 4;
					}
				}
			} else {
				for (int* j = sub_416900(); j; j = sub_416910(j)) {
					if (!strcmp((const char*)j + 72, "0")) {
						if (!_nox_wcsicmp((const wchar_t*)j + 6, (const wchar_t*)packet + 2)) {
							*(_BYTE*)(out + 2) = 19;
							*(_BYTE*)(out + 3) = 5;
							return 4;
						}
					} else if (!_strcmpi((const char*)j + 72, (const char*)packet + 56)) {
						*(_BYTE*)(out + 2) = 19;
						*(_BYTE*)(out + 3) = 5;
						return 4;
					}
				}
			}
			char* v35 = v78;
			char v52 = v78[100];
			if (v52 && (unsigned __int8)(1 << packet[54]) & (unsigned __int8)v52) {
				*(_BYTE*)(out + 2) = 19;
				*(_BYTE*)(out + 3) = 7;
				return 4;
			}
			if (v52 & 0x20) {
				*(_BYTE*)(out + 2) = 15;
				return 3;
			}
			if (*(short*)(v78 + 105) == -1 && *(short*)(v35 + 107) == -1) {
				*(_BYTE*)(out + 2) = 20; // OK
				return 3;
			}
			int id53 = sub_553D10();
			if (id53 < 0) {
				*(_BYTE*)(out + 2) = 20; // OK
				return 3;
			}
			nox_net_struct2_t* nx = &nox_net_struct2_arr[id53];
			nx->field_0 = 1;
			nx->field_1_1 = 0;
			nx->field_1_0 = 0;
			*(_QWORD*)(&nx->addr) = *(_QWORD*)v74;
			*((_DWORD*)(&nx->addr) + 1) = v75;
			*((_DWORD*)(&nx->addr) + 2) = v76;
			return nox_xxx_makePacketTime_552340(id53, out);
			}
		case 17:
			{
			char* v33 = sub_416640();
			char* v35 = v33;
			*(_BYTE*)(out + 0) = 0;
			*(_BYTE*)(out + 1) = p1;
			if (nox_wcscmp((const wchar_t*)(packet + 4), (const wchar_t*)v33 + 39)) {
				*(_BYTE*)(out + 2) = 19;
				*(_BYTE*)(out + 3) = 6;
				return 4;
			}
			if (*(short*)(v35 + 105) == -1 && *(short*)(v35 + 107) == -1) {
				*(_BYTE*)(out + 2) = 20;
				return 3;
			}
			int id53 = sub_553D10();
			if (id53 < 0) {
				*(_BYTE*)(out + 2) = 20;
				return 3;
			}
			nox_net_struct2_t* nx1 = &nox_net_struct2_arr[id53];
			nx1->field_0 = 1;
			nx1->field_1_1 = 0;
			nx1->field_1_0 = 0;
			*(_QWORD*)(&nx1->addr) = *(_QWORD*)v74;
			*((_DWORD*)(&nx1->addr) + 1) = v75;
			*((_DWORD*)(&nx1->addr) + 2) = v76;
			return nox_xxx_makePacketTime_552340(id53, out);
			}
		case 18:
			{
			int v39 = nox_platform_get_ticks() - *((_DWORD*)packet + 1);
			int id40 = sub_553D30((int)v74);
			if (id40 < 0)
				return 0;
			nox_net_struct2_t* nx1 = &nox_net_struct2_arr[id40];
			if (*((unsigned __int8*)packet + 3) != nx1->field_1_1)
				return 0;
			nx1->field_6[nx1->field_1_1] = v39;
			nx1->field_1_1++;
			if (nx1->field_1_1 >= 10)
				return 0;
			return nox_xxx_makePacketTime_552340(id40, out);
			}
		case 31:
			{
			if (pidb > NOX_NET_STRUCT_MAX) {
				printf("nox_net_struct_arr overflow (2): %d\n", (int)(pidb));
				abort();
			}
			nox_net_struct_t* ns8 = nox_net_struct_arr[pidb];
			_BYTE v14 = packetCur[0];
			_BYTE v15 = ns8->field_28_1;
			packetCur = &packetCur[1];
			_BYTE a4b = v14;
			printf("foo 0x%x 0x%x\n", v14, v15);
			if (v14 != v15) {
				sub_5551F0(pid, a4b, 1);
				sub_555360(pid, a4b, 1);
				ns8->field_28_1 = a4b;
				*(_BYTE*)(out + 0) = 38;
				*(_BYTE*)(out + 1) = ns8->field_28_1;
				ns1->func_yyy(pid, out, 2, ns8->data_3);
			}
			if ((unsigned int)packetCur >= packetEnd) {
				return 0;
			}
			break;
			}
		default:
			return 0;
		}
	}
}

//----- (00553D10) --------------------------------------------------------
int sub_553D10() {
	for (int i = 0; i < NOX_NET_STRUCT_MAX; i++) {
		if (!nox_net_struct2_arr[i].field_0)
			return i;
	}
	return -1;
}

//----- (00553D30) --------------------------------------------------------
int  sub_553D30(int a1) {
	for (int i = 0; i < NOX_NET_STRUCT_MAX; i++) {
		nox_net_struct2_t* nx = &nox_net_struct2_arr[i];
		if (nx->addr.sin_addr == *(_DWORD*)(a1 + 4) && nx->addr.sin_port == *(_WORD*)(a1 + 2))
			return i;
	}
	return -1;
}

//----- (00553D80) --------------------------------------------------------
int sub_553D80() { return *getMemU32Ptr(0x5D4594, 2495944); }

//----- (00553D90) --------------------------------------------------------
int sub_553D90() { return *getMemU32Ptr(0x5D4594, 2495952); }

//----- (00553DA0) --------------------------------------------------------
int sub_553DA0() { return *getMemU32Ptr(0x5D4594, 2495948); }

//----- (00553DB0) --------------------------------------------------------
int sub_553DB0() { return *getMemU32Ptr(0x5D4594, 2495956); }

//----- (00553DC0) --------------------------------------------------------
__int64  sub_553DC0(int a1) {
	int v1;          // eax
	unsigned int v2; // ecx
	int v3;          // edx

	v1 = dword_5d4594_2496472;
	v2 = 0;
	v3 = 127;
	do {
		v2 += *getMemU32Ptr(0x5D4594, 4 * v1 + 2495960);
		v1 = (v1 + 1) % 128;
		--v3;
	} while (v3);
	return (__int64)((double)v2 / (128.0 / (double)a1));
}

//----- (00553E10) --------------------------------------------------------
unsigned int  sub_553E10(int a1) {
	int v1;          // eax
	unsigned int v2; // ecx
	int v3;          // edx

	v1 = *getMemU32Ptr(0x5D4594, 2497504);
	v2 = 0;
	v3 = 127;
	do {
		v2 += *getMemU32Ptr(0x5D4594, 4 * v1 + 2496992);
		v1 = (v1 + 1) % 128;
		--v3;
	} while (v3);
	return v2 / (128 / a1);
}

//----- (00553E50) --------------------------------------------------------
unsigned int  sub_553E50(int a1) {
	int v1;          // eax
	unsigned int v2; // ecx
	int v3;          // edx

	v1 = dword_5d4594_2496988;
	v2 = 0;
	v3 = 127;
	do {
		v2 += *getMemU32Ptr(0x5D4594, 4 * v1 + 2496476);
		v1 = (v1 + 1) % 128;
		--v3;
	} while (v3);
	return v2 / (128 / a1);
}

//----- (00553E90) --------------------------------------------------------
unsigned int  sub_553E90(int a1) {
	int v1;          // eax
	unsigned int v2; // ecx
	int v3;          // edx

	v1 = *getMemU32Ptr(0x5D4594, 2498020);
	v2 = 0;
	v3 = 127;
	do {
		v2 += *getMemU32Ptr(0x5D4594, 4 * v1 + 2497508);
		v1 = (v1 + 1) % 128;
		--v3;
	} while (v3);
	return v2 / (128 / a1);
}

//----- (00553ED0) --------------------------------------------------------
int  sub_553ED0(int a3) {
	unsigned __int64 v1; // rax
	__int64 v2;          // rdi
	int v4;              // edi

	v1 = nox_platform_get_ticks();
	LODWORD(v2) = *getMemU32Ptr(0x5D4594, 8 * a3 + 2499052);
	HIDWORD(v2) = *getMemU32Ptr(0x5D4594, 8 * a3 + 2499056);
	if (v1 < v2 + 1000)
		return *getMemU32Ptr(0x5D4594, 4 * a3 + 2498536);
	v4 = *getMemU32Ptr(0x5D4594, 4 * a3 + 2498024);
	*getMemU32Ptr(0x5D4594, 8 * a3 + 2499052) = v1;
	LODWORD(v1) = v4;
	*getMemU32Ptr(0x5D4594, 4 * a3 + 2498536) = v4;
	*getMemU32Ptr(0x5D4594, 4 * a3 + 2498024) = 0;
	*getMemU32Ptr(0x5D4594, 8 * a3 + 2499056) = HIDWORD(v1);
	return 0;
}

//----- (00553F40) --------------------------------------------------------
void  sub_553F40(int a1, int a2) {
	*getMemU32Ptr(0x5D4594, 2495952) += a1;
	*getMemU32Ptr(0x5D4594, 2495956) += a2;
	*getMemU32Ptr(0x5D4594, 4 * *getMemU32Ptr(0x5D4594, 2497504) + 2496992) = a1;
	*getMemU32Ptr(0x5D4594, 4 * *getMemU32Ptr(0x5D4594, 2498020) + 2497508) = a2;
	*getMemU32Ptr(0x5D4594, 2497504) = (dword_5d4594_2496472 + 1) % 128;
	*getMemU32Ptr(0x5D4594, 2498020) = (dword_5d4594_2496988 + 1) % 128;
}

//----- (00553FC0) --------------------------------------------------------
void  sub_553FC0(int a1, int a2) {
	int v2; // edx
	int v3; // eax

	*getMemU32Ptr(0x5D4594, 2495944) += a1;
	*getMemU32Ptr(0x5D4594, 2495948) += a2;
	v2 = dword_5d4594_2496472;
	*getMemU32Ptr(0x5D4594, 4 * dword_5d4594_2496472 + 2495960) = a1;
	v3 = dword_5d4594_2496988;
	*getMemU32Ptr(0x5D4594, 4 * dword_5d4594_2496988 + 2496476) = a2;
	dword_5d4594_2496472 = (v2 + 1) % 128;
	dword_5d4594_2496988 = (v3 + 1) % 128;
}

//----- (00554030) --------------------------------------------------------
void  nox_xxx_netCountData_554030(int a1, int a2) { *getMemU32Ptr(0x5D4594, 4 * a2 + 2498024) += a1; }

//----- (00554040) --------------------------------------------------------
unsigned int  nox_server_makeServerInfoPacket_554040(const char* inBuf, int inSz, char* out) {
	char buf[72];

	char* v3 = sub_416640();
	char* game = nox_xxx_cliGamedataGet_416590(0);
	if (!sub_43AF30() || sub_4D6F30())
		return 0;
	char playerLimit = nox_xxx_servGetPlrLimit_409FA0();
	char playerCount = nox_common_playerInfoCount_416F40();
	if (nox_common_getEngineFlag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING)) {
		--playerCount;
		--playerLimit;
	}
	char* srvName = nox_xxx_serverOptionsGetServername_40A4C0();
	buf[0] = 0;
	buf[1] = 0;
	buf[2] = 13;
	buf[3] = playerCount;
	buf[4] = playerLimit;
	buf[5] = v3[101] & 0xF;
	buf[6] = ((unsigned __int8)v3[101]) >> 4;
	*(_DWORD*)&buf[7] = *((_DWORD*)game + 11);
	strcpy(&buf[10], nox_xxx_mapGetMapName_409B40());
	buf[19] = v3[102] | sub_43BE50_get_video_mode_id();
	buf[20] = v3[100];
	buf[21] = v3[100] & 0x10;
	*(_DWORD*)&buf[24] = *((_DWORD*)v3 + 12);
	unsigned int gameFlags = nox_common_gameFlags_getVal_40A5B0();
	if (sub_4D6F50()) {
		gameFlags = (gameFlags & 0xFFFFFF7Fu) | 0x1000u;
		*(_WORD*)&buf[68] = nox_game_getQuestStage_4E3CC0();
	}
	*(_DWORD*)&buf[28] = gameFlags;
	*(_DWORD*)&buf[32] = *((_DWORD*)game + 12);
	*(_WORD*)&buf[36] = *(_WORD*)(v3 + 105);
	*(_WORD*)&buf[38] = *(_WORD*)(v3 + 107);
	*(_DWORD*)&buf[40] = *((_DWORD*)v3 + 11);
	*(_DWORD*)&buf[44] = *(_DWORD*)(&inBuf[8]); // timestamp of the packet
	memcpy(&buf[48], game + 24, 20);
	memcpy(&out[0], buf, 72);
	strcpy(&out[72], srvName);
	return 72 + strlen(srvName)+1;
}

//----- (005541D0) --------------------------------------------------------
int*  nox_xxx_findPlayerID_5541D0(int a1) {
	int* result; // eax

	result = nox_common_list_getFirstSafe_425890(getMemIntPtr(0x5D4594, 2495908));
	if (!result)
		return 0;
	while (result[3] != a1) {
		result = nox_common_list_getNextSafe_4258A0(result);
		if (!result)
			return 0;
	}
	return result;
}

//----- (00554200) --------------------------------------------------------
int  nox_xxx_net_getIP_554200(unsigned int a1) {
	if (a1 > NOX_NET_STRUCT_MAX)
		return 0;
	if (!a1)
		return dword_5d4594_3843632;
	nox_net_struct_t* ns = nox_net_struct_arr[a1];
	if (!ns)
		return 0;
	return ns->addr.sin_addr;
}

//----- (00554230) --------------------------------------------------------
char* sub_554230() { return (char*)getMemAt(0x5D4594, 3843644); }

unsigned int sub_554290() {
	unsigned int v0;     // edi
	int v1;              // ebx
	char* v2;            // esi
	unsigned int v3;     // eax
	unsigned int result; // eax

	v0 = -1;
	v1 = 0;
	v2 = nox_common_playerInfoGetFirst_416EA0();
	if (!v2)
		goto LABEL_13;
	do {
		if (v2[2064] != 31 && sub_554240((unsigned __int8)v2[2064]) > 0) {
			v3 = sub_554240((unsigned __int8)v2[2064]);
			if (v3 < v0)
				v0 = v3;
			++v1;
		}
		v2 = nox_common_playerInfoGetNext_416EE0((int)v2);
	} while (v2);
	if (v1)
		result = v0;
	else
	LABEL_13:
		result = 0;
	return result;
}

int sub_554300() {
	int v0;     // ebx
	int v1;     // edi
	char* v2;   // esi
	int result; // eax

	v0 = 0;
	v1 = 0;
	v2 = nox_common_playerInfoGetFirst_416EA0();
	if (!v2)
		goto LABEL_11;
	do {
		if (v2[2064] != 31 && (int)sub_554240((unsigned __int8)v2[2064]) > 0) {
			v0 += sub_554240((unsigned __int8)v2[2064]);
			++v1;
		}
		v2 = nox_common_playerInfoGetNext_416EE0((int)v2);
	} while (v2);
	if (v1)
		result = v0 / v1;
	else
	LABEL_11:
		result = 0;
	return result;
}

//----- (00554240) --------------------------------------------------------
int  sub_554240(int a1) {
	char* v1;   // eax
	int result; // eax

	if (a1 != 31)
		return *getMemU32Ptr(0x5D4594, 32 * a1 + 2508848);
	v1 = sub_416640();
	switch (*(_DWORD*)(v1 + 66)) {
	case 1:
		result = sub_554290();
		break;
	case 2:
		result = sub_554300();
		break;
	case 3:
		result = *(_DWORD*)(v1 + 70);
		break;
	default:
		result = 0;
		break;
	}
	return result;
}

//----- (00554380) --------------------------------------------------------
int nox_xxx_netInit_554380(nox_net_struct_arg_t* narg) {
	int v2;                 // ebx
	__int16 v9;             // cx
	struct hostent* v10;    // eax
	uint16_t v11;            // [esp-4h] [ebp-1B8h]
	uint16_t v12;            // [esp-4h] [ebp-1B8h]

	if (!narg)
		return -2;
	if (narg->field_0)
		return -5;
	if (narg->field_4 > 128)
		return -2;
	*getMemU8Ptr(0x5D4594, 3843644) = 0;
	dword_5d4594_3843632 = 0;
	*getMemU8Ptr(0x5D4594, 3843660) = 0;
	v2 = -1;
	for (int i = 0; i < NOX_NET_STRUCT_MAX; i++) {
		if (!nox_net_struct_arr[i]) {
			v2 = i;
			break;
		}
	}
	if (v2 == -1)
		return -8;
	nox_net_struct_t* ns = nox_xxx_makeNewNetStruct_553000(narg);
	if (!ns)
		return -1;
	nox_net_struct_arr[v2] = ns;
	ns->data_2_base[0] = v2;
	ns->id = -1;
	if (nox_net_init() == -1) {
		return -1;
	}
	nox_socket_t sock = nox_net_socket_udp();
	ns->sock = sock;
	if (sock == -1) {
		nox_net_stop();
		return -1;
	}
	int port = narg->port;
	if (port < 1024 || port > 0x10000)
		narg->port = 18590;
	v11 = (_WORD)(narg->port);

	struct nox_net_sockaddr_in name;
	name.sin_family = NOX_AF_INET;
	name.sin_port = 0;
	name.sin_addr = 0;
	memset(name.sin_zero, 0, 8);

	v9 = (_WORD)(narg->port);
	name.sin_port = htons(v11);
	name.sin_addr = 0;
	*getMemU16Ptr(0x5D4594, 3843636) = v9;
	while (nox_net_bind(ns->sock, &name) == -1) {
		if (nox_net_error(ns->sock) != NOX_NET_EADDRINUSE) {
			nox_net_stop();
			return -1;
		}
		v12 = (_WORD)(narg->port) + 1;
		++narg->port;
		name.sin_port = htons(v12);
	}
	if (gethostname((char*)getMemAt(0x5D4594, 3843660), 128) != -1) {
		v10 = gethostbyname((const char*)getMemAt(0x5D4594, 3843660));
		if (v10) {
			dword_5d4594_3843632 = **(_DWORD**)v10->h_addr_list;
			strcpy((char*)getMemAt(0x5D4594, 3843644), nox_net_ip2str(*(nox_net_in_addr*)&dword_5d4594_3843632));
		}
	}
	return v2;
}

//----- (005545A0) --------------------------------------------------------
__int16 sub_5545A0() { return *getMemU16Ptr(0x5D4594, 3843636); }

//----- (005545B0) --------------------------------------------------------
int  nox_xxx_netStructReadPackets_5545B0(unsigned int a1) {
	unsigned int v1; // ebx
	int v4;          // ebp
	char v8;         // [esp+Fh] [ebp-1h]
	unsigned int v9; // [esp+14h] [ebp+4h]

	v1 = a1;
	v8 = 11;
	if (a1 >= NOX_NET_STRUCT_MAX)
		return -3;
	nox_net_struct_t* ns = nox_net_struct_arr[a1];
	if (!ns)
		return 0;
	v4 = ns->id;
	unsigned int v5;
	if (v4 == -1) {
		v5 = 0;
		v9 = NOX_NET_STRUCT_MAX;
		v4 = v1;
	} else {
		v5 = a1;
		v9 = a1 + 1;
		nox_net_struct_t* ns2 = nox_net_struct_arr[v4];
		if (!ns2 || ns2->id != -1) {
			nox_xxx_netStructFree_5531C0(ns);
			nox_net_struct_arr[v1] = 0;
			return 0;
		}
	}
	for (int i = v5; i < v9; i++) {
		nox_net_struct_t* ns2 = nox_net_struct_arr[i];
		if (ns2 && ns2->id == v4) {
			nox_xxx_netSendReadPacket_5528B0(i, 1);
			nox_xxx_netSendSock_552640(i, &v8, 1, 0);
			nox_xxx_netSendReadPacket_5528B0(i, 1);
			nox_net_struct_arr[v4]->field_21--;
			sub_555360(v1, 0, 2);
			nox_xxx_netStructFree_5531C0(ns2);
			nox_net_struct_arr[i] = 0;
		}
	}
	return 0;
}

//----- (005546A0) --------------------------------------------------------
int  nox_server_netClose_5546A0(unsigned int a1) {
	if (a1 >= NOX_NET_STRUCT_MAX)
		return -3;
	nox_net_struct_t* ns = nox_net_struct_arr[a1];
	if (ns) {
		nox_net_close(ns->sock);
		nox_xxx_netStructFree_5531C0(ns);
		nox_net_struct_arr[a1] = 0;
		nox_net_stop();
	}
	return 0;
}

//----- (005546F0) --------------------------------------------------------
int nox_xxx_netPreStructToFull_5546F0(nox_net_struct_arg_t* narg) {
	if (!narg)
		return -2;
	if (narg->field_0)
		return -5;
	int ind = -1;
	for (int i = 0; i < NOX_NET_STRUCT_MAX; i++) {
		if (!nox_net_struct_arr[i]) {
			ind = i;
			break;
		}
	}
	if (ind == -1)
		return -8;
	nox_net_struct_t* ns = nox_xxx_makeNewNetStruct_553000(narg);
	if (!ns)
		return -1;
	nox_net_struct_arr[ind] = ns;
	return ind;
}

//----- (00554760) --------------------------------------------------------
int  sub_554760(int a1, char* cp, int hostshort, int a4, int a5) {
	int v7;                 // eax
	unsigned int v8;        // ebx
	struct hostent* v9;     // eax
	int v10;                // esi
	char v11;               // al
	char v12;               // [esp+12h] [ebp-1B2h]
	struct nox_net_sockaddr_in name;   // [esp+14h] [ebp-1B0h]
	int v15;                // [esp+28h] [ebp-19Ch]

	if ((unsigned int)a1 >= NOX_NET_STRUCT_MAX)
		return -3;
	nox_net_struct_t* ns = nox_net_struct_arr[a1];
	if (!ns)
		return -3;
	if (!cp)
		return -4;
	if (hostshort < 1024 || hostshort > 0x10000)
		return -15;
	if (nox_net_init() == -1)
		return -21;
	v7 = nox_net_socket_udp();
	ns->sock = v7;
	if (v7 == -1) {
		nox_net_stop();
		return -22;
	}
	if ((unsigned __int8)*cp < 0x30u || (unsigned __int8)*cp > 0x39u) {
		v9 = gethostbyname(cp);
		if (!v9) {
			nox_net_stop();
			return -4;
		}
		v8 = **(_DWORD**)v9->h_addr_list;
	} else {
		v8 = inet_addr(cp);
	}
	v15 = 0;

	ns->addr.sin_family = NOX_AF_INET;
	ns->addr.sin_port = htons(hostshort);
	ns->addr.sin_addr = v8;
	memset(ns->addr.sin_zero, 0, 8);

	v10 = nox_client_getClientPort_40A420();
	name.sin_family = NOX_AF_INET;
	name.sin_port = htons(v10);
	name.sin_addr = 0;
	memset(name.sin_zero, 0, 8);
	while (nox_net_bind(ns->sock, &name) == -1) {
		if (nox_net_error(ns->sock) != NOX_NET_EADDRINUSE) {
			nox_net_stop();
			return -1;
		}
		name.sin_port = htons(++v10);
	}
	dword_5d4594_3844304 = 0;
	v12 = 0;
	v11 = nox_xxx_netSendSock_552640(a1, &v12, 1, NOX_NET_SEND_NO_LOCK | NOX_NET_SEND_FLAG2);
	if (nox_xxx_cliWaitServerResponse_5525B0(a1, v11, 60, 6))
		return -23;
	printf("bar %d\n", dword_5d4594_3844304);
	if (dword_5d4594_3844304 && ns->id >= 0) {
		memset(getMemAt(0x5D4594, 2512892), 0, 0x400u);
		*getMemU8Ptr(0x5D4594, 2512892) = 31;
		*getMemU8Ptr(0x5D4594, 2512892 + 1) = ns->data_2_base[1];
		*getMemU8Ptr(0x5D4594, 2512892 + 2) = 32;
		if (a4)
			memcpy(getMemAt(0x5D4594, 2512892 + 3), (const void*)a4, a5);
		nox_xxx_netSendSock_552640(a1, getMemAt(0x5D4594, 2512892), a5 + 3, NOX_NET_SEND_NO_LOCK | NOX_NET_SEND_FLAG2);
	}
	return ns->id;
}

//----- (005549F0) --------------------------------------------------------
int  sub_5549F0(unsigned int a1) {
	char v2; // [esp+7h] [ebp-1h]

	v2 = 10;
	if (a1 >= NOX_NET_STRUCT_MAX)
		return -3;
	if (nox_net_struct_arr[a1]) {
		nox_xxx_netSendReadPacket_5528B0(a1, 0);
		nox_xxx_netSendSock_552640(a1, &v2, 1, 0);
		nox_xxx_netSendReadPacket_5528B0(a1, 0);
		sub_554A50(a1);
	}
	return 0;
}

//----- (00554A50) --------------------------------------------------------
int  sub_554A50(unsigned int a1) {
	if (a1 >= NOX_NET_STRUCT_MAX)
		return -3;
	nox_net_struct_t* ns = nox_net_struct_arr[a1];
	if (ns) {
		nox_net_close(ns->sock);
		OnLibraryNotice_259(ns);
		nox_xxx_netStructFree_5531C0(nox_net_struct_arr[a1]);
		nox_net_struct_arr[a1] = 0;
		nox_net_stop();
	}
	return 0;
}

//----- (00554AA0) --------------------------------------------------------
#ifdef NOX_CGO
void nox_client_OnServerSearch(int hostshort, void* data, int size);
#endif // NOX_CGO
void  nox_xxx_lobbyMakePacket_554AA0(uint16_t hostshort, const char* payload, int payloadSz, unsigned int ticks) {
	int dataSz = 12;
	char* data = (char*)malloc(12 + payloadSz);
	data[0] = 0;
	data[1] = 0;
	data[2] = 12;
	// data[3] = ???
	*((_WORD*)&data[4]) = 0;
	// data[6] = ???
	// data[7] = ???
	*((_DWORD*)&data[8]) = ticks;
	if (payload && payloadSz > 0) {
		*((_WORD*)&data[4]) = payloadSz;
		memcpy(&data[12], payload, payloadSz);
		dataSz += payloadSz;
	}
	dword_5d4594_3844304 = 0;
	nox_xxx_sendLobbyPacket_554C80(hostshort, data, dataSz);
#ifndef NOX_CGO
	OnLibraryNotice_260(hostshort, data, dataSz);
#else // NOX_CGO
	nox_client_OnServerSearch(hostshort, data, dataSz);
#endif // NOX_CGO
	free(data);
}

//----- (00554B40) --------------------------------------------------------
int  nox_xxx_createSocketLocal_554B40(uint16_t hostshort) {
	if (dword_5d4594_2513916 == 1)
		return -14;
	if (nox_net_init() == -1) {
		return -1;
	}
	nox_xxx_sockLocalBroadcast_2513920 = nox_net_socket_udp_broadcast();
	if (nox_xxx_sockLocalBroadcast_2513920 == -1) {
		nox_net_stop();
		return -1;
	}
	dword_5d4594_2513924 = nox_net_socket_udp();
	if (*(int*)&dword_5d4594_2513924 == -1) {
		nox_net_stop();
		return -1;
	}
	struct nox_net_sockaddr_in name;
	name.sin_family = NOX_AF_INET;
	name.sin_port = htons(hostshort);
	name.sin_addr = 0;
	memset(name.sin_zero, 0, 8);
	if (nox_net_bind(nox_xxx_sockLocalBroadcast_2513920, &name) == -1) {
		nox_net_stop();
		return -1;
	}
	nox_game_SetCliDrawFunc((int)sub_554FF0);
	dword_5d4594_2513916 = 1;
	return 0;
}

//----- (00554C80) --------------------------------------------------------
int  nox_xxx_sendLobbyPacket_554C80(uint16_t hostshort, char* buf, int a3) {
	int v3 = 0;
	if (!dword_5d4594_2513916) {
		return -17;
	}

	struct nox_net_sockaddr_in to;
	to.sin_family = NOX_AF_INET;
	to.sin_port = htons(hostshort);
	to.sin_addr = -1;
	memset(to.sin_zero, 0, 8);

	int result = 0;
	if (!buf || (unsigned __int16)a3 < 2u ||
		(result = nox_net_sendto(nox_xxx_sockLocalBroadcast_2513920, buf, (unsigned __int16)a3, &to), v3 = result,
		 result != -1)) {
		result = v3;
	}
	return result;
}

//----- (00554D10) --------------------------------------------------------
int sub_554D10() {
	if (dword_5d4594_2513916) {
		nox_net_close(nox_xxx_sockLocalBroadcast_2513920);
		nox_net_close(dword_5d4594_2513924);
		nox_xxx_sockLocalBroadcast_2513920 = 0;
		dword_5d4594_2513924 = 0;
		dword_5d4594_2513916 = 0;
		nox_game_SetCliDrawFunc(0);
		nox_client_setOnLobbyServer_555000(0);
		nox_net_stop();
	}
	return 0;
}

//----- (00554D70) --------------------------------------------------------
#ifndef NOX_CGO
int  sub_554D70(char a1) {
	int result;           // eax
	int v2;               // ebp
	uint16_t v7;           // ax
	int v8;               // [esp+0h] [ebp-230h]
	int v11;              // [esp+18h] [ebp-218h]
	char buf[256];        // [esp+30h] [ebp-200h]
	char in[256];         // [esp+130h] [ebp-100h]

	if (!dword_5d4594_2513916)
		return -17;
	v11 = a1 & 1;
	unsigned int argp = 0;
	if (a1 & 1) {
		result = nox_net_recv_available(nox_xxx_sockLocalBroadcast_2513920, &argp);
		if (result == -1)
			return result;
		if (!argp) {
			return -1;
		}
	} else {
		argp = 1;
	}
	while (1) {
		struct nox_net_sockaddr from;
		v2 = mix_recvfrom(nox_xxx_sockLocalBroadcast_2513920, buf, 256, &from);
		if (v2 == -1)
			break;
		unsigned __int8 op = buf[2];
		if (op < 32) {
			memcpy(in, &from, sizeof(struct nox_net_sockaddr));
			if (op == 13 || nox_client_getServerAddr_43B300() == *(_DWORD*)&from.sa_data[2]) {
				switch (op) {
				case 13:;
					char* saddr = nox_net_ip2str(*(nox_net_in_addr*)&in[4]);
					if (&v8 != (int*)-120) {
						if (saddr) {
							if (nox_client_onLobbyServer_2513928) {
								OnLibraryNotice_262(&v8);
								uint16_t port = ntohs(*(uint16_t*)&in[2]);
								if (nox_client_onLobbyServer_2513928(saddr, port, &buf[72], buf) == 1)
									nox_client_setOnLobbyServer_555000(0);
							}
						}
					}
					break;
				case 15:
					if (sub_43B6D0())
						sub_43AF90(5);
					break;
				case 16:
					if (sub_43B6D0()) {
						sub_43AF90(4);
						buf[2] = 18;
						v7 = htons(*(uint16_t*)from.sa_data);
						nox_client_sendToServer_555010(*(int*)&from.sa_data[2], v7, buf, 8);
					}
					break;
				/*case 19:
				  if ( sub_43B6D0() )
					sub_43AFA0((unsigned __int8)buf[3]);
				  break;*/
				case 19:
				case 20:
					if (sub_43B6D0() && sub_43AF80() == 3)
						sub_43AF90(7);
					break;
				case 21:
					if (sub_43B6D0())
						sub_43AF90(8);
					break;
				default:
					break;
				}
			}
		}
		if (!v11 || (a1 & 4)) {
			return v2;
		}
		argp = 0;
		if (nox_net_recv_available(nox_xxx_sockLocalBroadcast_2513920, &argp) == -1)
			return -1;
		if (!argp) {
			return v2;
		}
	}
	return -1;
}
#endif // NOX_CGO
// 554EC5: variable 'v6' is possibly undefined

//----- (00554FF0) --------------------------------------------------------
int sub_554FF0() {
	sub_554D70(1);
	return 1;
}

//----- (00555000) --------------------------------------------------------
void nox_client_setOnLobbyServer_555000(int (*fnc)(const char*, uint16_t, const char*, const char*)) {
	nox_client_onLobbyServer_2513928 = fnc;
}

//----- (00555010) --------------------------------------------------------
#ifndef NOX_CGO
int  nox_client_sendToServer_555010(unsigned int addr, uint16_t port, char* buf, int sz) {
	if (!dword_5d4594_2513916)
		return -17;
	if (!buf || sz < 2)
		return 0;

	struct nox_net_sockaddr_in to;
	to.sin_family = NOX_AF_INET;
	to.sin_port = htons(port);
	to.sin_addr = addr;
	memset(to.sin_zero, 0, 8);
	return nox_net_sendto(nox_xxx_sockLocalBroadcast_2513920, buf, sz, &to);
}

//----- (005550A0) --------------------------------------------------------
int  nox_client_sendJoinGame_5550A0(unsigned int addr, uint16_t port, char* data) {
	data[0] = 0;
	data[1] = 0;
	data[2] = 14;
	return nox_client_sendToServer_555010(addr, port, data, 100);
}

//----- (005550D0) --------------------------------------------------------
int  sub_5550D0(int a1, uint16_t hostshort, char* buf) {
	buf[0] = 0;
	buf[1] = 0;
	buf[2] = 17;
	return nox_client_sendToServer_555010(a1, hostshort, buf, 22);
}
#endif // NOX_CGO

//----- (00555100) --------------------------------------------------------
int sub_555100() { return dword_5d4594_2513916; }

//----- (00555130) --------------------------------------------------------
int  sub_555130(unsigned int a1, const void* a2, signed int a3) {
	_DWORD* v5; // eax

	if (a3 > *getMemIntPtr(0x5D4594, 2512884))
		return -1;
	if (!a2)
		return -1;
	if (a1 >= NOX_NET_STRUCT_MAX)
		return -3;
	nox_net_struct_t* ns = nox_net_struct_arr[a1];
	if (!ns)
		return -3;
	v5 = nox_alloc_class_new_obj_zero(nox_alloc_gQueue_3844300);
	if (!v5)
		return -1;
	*v5 = ns->field_29;
	ns->field_29 = v5;
	v5[3] = 1;
	*((_BYTE*)v5 + 20) = ns->data_2_base[0] | 0x80;
	*((_BYTE*)v5 + 21) = ns->field_28_0++;
	v5[4] = a3 + 2;
	memcpy((char*)v5 + 22, a2, a3);
	return *((unsigned __int8*)v5 + 21);
}

//----- (005551F0) --------------------------------------------------------
int  sub_5551F0(unsigned int a1, char a2, int a3) {
	int* i; // eax

	if (a1 >= NOX_NET_STRUCT_MAX)
		return -3;
	nox_net_struct_t* ns = nox_net_struct_arr[a1];
	if (!ns)
		return -3;
	for (i = ns->field_29; i; i = (int*)*i) {
		if (a3) {
			if (*((_BYTE*)i + 21) == a2) {
				i[3] = 1;
				continue;
			}
		} else if (i[1] <= (int)dword_5d4594_2495920) {
			i[3] = 1;
			continue;
		}
	}
	return 0;
}

//----- (00555250) --------------------------------------------------------
int  sub_555250(unsigned int a1, _DWORD* a2) {
	int v3;     // eax
	int v4;     // ecx
	int result; // eax

	if (a1 >= NOX_NET_STRUCT_MAX)
		return 0;
	nox_net_struct_t* ns = nox_net_struct_arr[a1];
	if (!ns)
		return 0;
	v3 = ns->field_29;
	if (!v3)
		return 0;
	v4 = *(_DWORD*)(v3 + 16);
	result = v3 + 22;
	*a2 = v4;
	dword_5d4594_2513932 = *(_DWORD*)(result - 22);
	return result;
}

//----- (00555290) --------------------------------------------------------
int  sub_555290(unsigned int a1, _DWORD* a2) {
	int result; // eax

	if (!dword_5d4594_2513932 || a1 >= NOX_NET_STRUCT_MAX || !nox_net_struct_arr[a1])
		return 0;
	result = dword_5d4594_2513932 + 22;
	*a2 = *(_DWORD*)(dword_5d4594_2513932 + 16);
	dword_5d4594_2513932 = *(_DWORD*)(result - 22);
	return result;
}

//----- (005552D0) --------------------------------------------------------
int  nox_xxx_netSend_5552D0(unsigned int a1, char a2, int a3) {
	int* i; // esi
	int v6; // eax

	if (a1 >= NOX_NET_STRUCT_MAX)
		return -3;
	nox_net_struct_t* ns = nox_net_struct_arr[a1];
	if (!ns)
		return -3;
	for (i = ns->field_29; i; i = (int*)*i) {
		if (a3) {
			if (*((_BYTE*)i + 21) == a2)
				goto LABEL_10;
		} else if (i[3]) {
		LABEL_10:
			v6 = i[4];
			i[3] = 0;
			i[1] = dword_5d4594_2495920 + 2000;
			if (nox_xxx_sendto_551F90(ns->sock, (char*)i + 20, v6, &ns->addr) == -1)
				return 0;
			continue;
		}
	}
	return 0;
}

//----- (00555360) --------------------------------------------------------
int  sub_555360(unsigned int a1, unsigned __int8 a2, int a3) {
	char* v5;    // esi
	_QWORD* v6;  // eax
	_DWORD* v7;  // ecx
	char v8[24]; // [esp+4h] [ebp-18h]

	if (a1 >= NOX_NET_STRUCT_MAX)
		return -3;
	nox_net_struct_t* ns = nox_net_struct_arr[a1];
	if (!ns)
		return -3;
	*(_DWORD*)v8 = ns->field_29;
	v5 = v8;
	while (*(_DWORD*)v5) {
		if (a3) {
			if (a3 == 1) {
				if (a2 < 0x20u || a2 > 0xE0u) {
					if (*(_BYTE*)(*(_DWORD*)v5 + 21) < (char)a2)
						goto LABEL_17;
				} else if (*(_BYTE*)(*(_DWORD*)v5 + 21) < a2) {
					goto LABEL_17;
				}
			} else if (a3 == 2) {
			LABEL_17:
				v6 = *(_QWORD**)v5;
				v7 = ns->field_29;
				if (*(_DWORD**)v5 == v7)
					ns->field_29 = *v7;
				*(_DWORD*)v5 = **(_DWORD**)v5;
				*(_DWORD*)v6 = 0;
				nox_alloc_class_free_obj(nox_alloc_gQueue_3844300, v6);
				continue;
			}
		} else if (*(_BYTE*)(*(_DWORD*)v5 + 21) == a2) {
			goto LABEL_17;
		}
		v5 = *(char**)v5;
	}
	return 0;
}

//----- (00565360) --------------------------------------------------------
int*  sub_565360(int a1, _WORD* a2, int* a3, unsigned int a4, int a5, int a6) {
	int* v6;               // edx
	int v7;                // ebp
	_WORD* v8;             // esi
	int* result;           // eax
	int* v10;              // edi
	int v11;               // ecx
	_WORD* v12;            // eax
	int v13;               // ebx
	int* v14;              // esi
	int v15;               // edi
	unsigned __int16* v16; // eax
	int v17;               // esi
	unsigned __int16 v18;  // cx
	unsigned __int16 v19;  // bx
	unsigned __int16* v20; // eax
	unsigned __int16 v21;  // cx
	unsigned __int16 v22;  // bx
	int* v23;              // [esp+Ch] [ebp-4h]
	int v24;               // [esp+1Ch] [ebp+Ch]
	int* v25;              // [esp+24h] [ebp+14h]
	int* v26;              // [esp+28h] [ebp+18h]

	v6 = a3;
	v7 = 2 * a6;
	v8 = a2;
	result = (int*)((char*)a3 + 2 * 2 * a6 * a5);
	v10 = a3;
	v23 = (int*)((char*)a3 + 2 * 2 * a6 * a5);
	v26 = a3;
	v25 = &a3[2 * a4];
	if (a3 < result) {
		while (1) {
			result = (int*)(*v8 & 0x1FFF);
			v11 = *v8 & 0xE000;
			v24 = (int)(v8 + 1);
			if (*v8 & 0xE000) {
				if (v11 == 0x2000) {
					v12 = (_WORD*)(a1 + 16 * (*v8 & 0x1FFF));
					v13 = 2;
					do {
						v14 = v6;
						v15 = 4;
						do {
							if (*v12 != -32768)
								*(_WORD*)v14 = *v12;
							v14 = (int*)((char*)v14 + 2);
							++v12;
							--v15;
						} while (v15);
						v6 = (int*)((char*)v6 + v7);
						--v13;
					} while (v13);
					v10 = v26;
					result = (int*)(2 * v7);
					v6 = (int*)((char*)v6 + 8 - 2 * v7);
				} else if (v11 == 0x4000) {
					v6 += 2;
				}
			} else {
				v16 = (unsigned __int16*)(a1 + 16 * (*v8 & 0x1FFF));
				v17 = 2;
				do {
					v18 = v16[1];
					v19 = *v16;
					v20 = v16 + 2;
					*v6 = v19 | (v18 << 16);
					v21 = v20[1];
					v22 = *v20;
					v16 = v20 + 2;
					v6[1] = v22 | (v21 << 16);
					v6 = (int*)((char*)v6 + v7);
					--v17;
				} while (v17);
				result = (int*)(2 * v7);
				v6 = (int*)((char*)v6 + 8 - 2 * v7);
			}
			if (v6 == v25) {
				v10 = (int*)((char*)v10 + 2 * v7);
				result = (int*)a4;
				v26 = v10;
				v6 = v10;
				v25 = &v10[2 * a4];
			}
			if (v6 >= v23)
				break;
			v8 = (_WORD*)v24;
		}
	}
	return result;
}

//----- (005654A0) --------------------------------------------------------
unsigned int  sub_5654A0(int a1, unsigned __int8* a2, int* a3, unsigned int a4, int a5, int a6) {
	int v6;                // edi
	int v7;                // ecx
	int* v8;               // ebx
	unsigned __int8* v9;   // esi
	unsigned int result;   // eax
	__int16 v11;           // ax
	__int16 v12;           // cx
	unsigned int v13;      // ecx
	int* v14;              // edi
	unsigned __int16* v15; // ecx
	int v16;               // eax
	int v17;               // esi
	unsigned __int16 v18;  // dx
	unsigned __int16 v19;  // bp
	unsigned __int16* v20; // ecx
	unsigned __int16 v21;  // dx
	unsigned __int16 v22;  // bp
	unsigned __int8* v23;  // ecx
	int* v24;              // edx
	unsigned __int16* v25; // eax
	int v26;               // edi
	unsigned __int16 v27;  // cx
	unsigned __int16 v28;  // bp
	unsigned __int16* v29; // eax
	unsigned __int16 v30;  // cx
	unsigned __int16 v31;  // bp
	int v32;               // edx
	unsigned __int16* v33; // eax
	int v34;               // edx
	unsigned __int16 v35;  // cx
	unsigned __int16 v36;  // bp
	unsigned __int16* v37; // eax
	unsigned __int16 v38;  // cx
	unsigned __int16 v39;  // bp
	int v40;               // ebp
	__int16* v41;          // eax
	int* v42;              // edi
	__int16 v43;           // dx
	int* v44;              // ecx
	int v45;               // esi
	int v46;               // ebp
	int* v47;              // ecx
	int v48;               // eax
	int* v49;              // ecx
	int* v50;              // esi
	int v51;               // edi
	int v52;               // eax
	_DWORD* v53;           // esi
	int v54;               // edi
	_WORD* v55;            // eax
	int v56;               // ebp
	int* v57;              // edx
	int v58;               // esi
	int* v59;              // [esp+Ch] [ebp-Ch]
	int* v60;              // [esp+10h] [ebp-8h]
	unsigned int v61;      // [esp+14h] [ebp-4h]
	int v62;               // [esp+20h] [ebp+8h]
	unsigned int v63;      // [esp+24h] [ebp+Ch]
	int v64;               // [esp+24h] [ebp+Ch]
	int v65;               // [esp+24h] [ebp+Ch]
	int v66;               // [esp+2Ch] [ebp+14h]
	unsigned int v67;      // [esp+2Ch] [ebp+14h]
	unsigned int v68;      // [esp+30h] [ebp+18h]

	v6 = 2 * a6;
	v7 = 2 * a6 * a5;
	v8 = a3;
	v9 = a2;
	v68 = 2 * a6;
	result = (unsigned int)a3 + 2 * v7;
	v61 = result;
	v59 = a3;
	v60 = &a3[2 * a4];
	if ((unsigned int)a3 < result) {
		do {
			v11 = *(_WORD*)v9;
			v9 += 2;
			v12 = v11;
			result = v11 & 0x1FFF;
			v13 = v12 & 0xE000;
			v66 = (int)v9;
			if (v13 > 0x6000) {
				switch (v13) {
				case 0x8000u:
					v55 = (_WORD*)(a1 + 16 * (unsigned __int16)result);
					v56 = 2;
					do {
						v57 = v8;
						v58 = 4;
						do {
							if (*v55 != -32768)
								*(_WORD*)v57 = *v55;
							v57 = (int*)((char*)v57 + 2);
							++v55;
							--v58;
						} while (v58);
						v8 = (int*)((char*)v8 + v6);
						--v56;
					} while (v56);
					result = 2 * v6;
					v8 = (int*)((char*)v8 + 8 - 2 * v6);
					goto LABEL_52;
				case 0xA000u:
					goto LABEL_16;
				case 0xC000u:
					v40 = *v9;
					v66 = (int)(v9 + 1);
					v41 = (__int16*)(a1 + 16 * (unsigned __int16)result);
					v62 = 2;
					do {
						v42 = v8;
						v65 = 4;
						do {
							v43 = *v41;
							if (*v41 != -32768) {
								v44 = v42;
								if (v40) {
									v45 = v40;
									do {
										*(_WORD*)v44 = v43;
										v44 += 2;
										--v45;
									} while (v45);
								}
							}
							v42 = (int*)((char*)v42 + 2);
							++v41;
							--v65;
						} while (v65);
						v8 = (int*)((char*)v8 + v68);
						--v62;
					} while (v62);
					v6 = v68;
					result = 2 * v68;
					v8 = (int*)((char*)v8 + 8 * v40 - 2 * v68);
					goto LABEL_52;
				}
			} else if (v13 == 24576) {
				v33 = (unsigned __int16*)(a1 + 16 * (unsigned __int16)result);
				v34 = 2;
				do {
					v35 = v33[1];
					v36 = *v33;
					v37 = v33 + 2;
					*v8 = v36 | (v35 << 16);
					v38 = v37[1];
					v39 = *v37;
					v33 = v37 + 2;
					v8[1] = v39 | (v38 << 16);
					v8 = (int*)((char*)v8 + v6);
					--v34;
				} while (v34);
				result = 2 * v6;
				v8 = (int*)((char*)v8 + 8 - 2 * v6);
			} else if (v13) {
				if (v13 == 0x2000) {
					v32 = (((unsigned __int16)result >> 7) & 0x3E) + 2;
					if ((((unsigned __int16)result >> 7) & 0x3E) == -2) {
					LABEL_16:
						result = (unsigned __int16)result;
						v32 = *v9;
						v66 = (int)(v9 + 1);
					} else {
						result = (unsigned __int8)result;
					}
					v46 = 2;
					v47 = (int*)(a1 + 16 * result);
					do {
						v48 = *v47;
						v49 = v47 + 1;
						v50 = v8;
						if (v32) {
							v51 = v32;
							do {
								*v50 = v48;
								v50 += 2;
								--v51;
							} while (v51);
							v6 = v68;
						}
						v52 = *v49;
						v47 = v49 + 1;
						v53 = v8 + 1;
						if (v32) {
							v54 = v32;
							do {
								*v53 = v52;
								v53 += 2;
								--v54;
							} while (v54);
							v6 = v68;
						}
						v8 = (int*)((char*)v8 + v6);
						--v46;
					} while (v46);
					result = 2 * v6;
					v8 = (int*)((char*)v8 + 8 * v32 - 2 * v6);
				LABEL_52:
					v9 = (unsigned __int8*)v66;
					goto LABEL_53;
				}
				if (v13 == 0x4000) {
					v14 = v8;
					v15 = (unsigned __int16*)(a1 + 16 * *(v9 - 2));
					v63 = ((result >> 7) & 0x3E) + 2;
					v16 = 2;
					v17 = 4 * (v68 >> 2);
					do {
						v18 = v15[1];
						v19 = *v15;
						v20 = v15 + 2;
						*v14 = v19 | (v18 << 16);
						v21 = v20[1];
						v22 = *v20;
						v15 = v20 + 2;
						v14[1] = v22 | (v21 << 16);
						v14 = (int*)((char*)v14 + v17);
						--v16;
					} while (v16);
					result = v63;
					v23 = (unsigned __int8*)v66;
					v8 += 2;
					if (v63) {
						v67 = v63;
						do {
							v24 = v8;
							v64 = (int)(v23 + 1);
							v25 = (unsigned __int16*)(a1 + 16 * *v23);
							v26 = 2;
							do {
								v27 = v25[1];
								v28 = *v25;
								v29 = v25 + 2;
								*v24 = v28 | (v27 << 16);
								v30 = v29[1];
								v31 = *v29;
								v25 = v29 + 2;
								v24[1] = v31 | (v30 << 16);
								v24 = (int*)((char*)v24 + v17);
								--v26;
							} while (v26);
							v23 = (unsigned __int8*)v64;
							v8 += 2;
							result = --v67;
						} while (v67);
					}
					v6 = v68;
					v66 = (int)v23;
					goto LABEL_52;
				}
			} else {
				result = (unsigned __int8)result;
				v8 += 2 * (unsigned __int8)result;
			}
		LABEL_53:
			if (v8 == v60) {
				v8 = (int*)((char*)v59 + 2 * v6);
				result = a4;
				v59 = v8;
				v60 = &v8[2 * a4];
			}
		} while ((unsigned int)v8 < v61);
	}
	return result;
}

//----- (0056F1C0) --------------------------------------------------------
int sub_56F1C0() {
	int v0;     // eax
	int result; // eax

	v0 = time(0);
	sub_56FF00(v0);
	dword_5d4594_2516352 = 0;
	dword_5d4594_2516348 = nox_frame_xxx_2598000;
	dword_5d4594_2516344 = 0;
	*getMemU16Ptr(0x587000, 311204) = 0;
	dword_5d4594_2516356 = 657757279;
	dword_5d4594_2516348 ^= nox_xxx_protect_56F240();
	dword_5d4594_2516328 = ~dword_5d4594_2516348;
	*getMemU32Ptr(0x5D4594, 2516340) = nox_xxx_protectionCreateInt_56F400(0);
	sub_56F250();
	result = nox_xxx_protectionCreateInt_56F400(1);
	*getMemU32Ptr(0x5D4594, 2516332) = result;
	return result;
}

//----- (0056F240) --------------------------------------------------------
int nox_xxx_protect_56F240() { return sub_56FF80(1, -1); }

//----- (0056F250) --------------------------------------------------------
int sub_56F250() {
	int v0;     // esi
	int result; // eax

	v0 = 7;
	do {
		result = nox_xxx_protectionCreateStructForInt_56F280(*(int*)&dword_5d4594_2516356, 0);
		--v0;
		++dword_5d4594_2516356;
	} while (v0);
	return result;
}

//----- (0056F280) --------------------------------------------------------
int  nox_xxx_protectionCreateStructForInt_56F280(int a1, int a2) {
	_DWORD* v2; // eax
	int v3;     // ecx
	int v4;     // ecx

	v2 = calloc(1u, 0x10u);
	if (!v2)
		return 0;
	v2[3] = 0;
	v2[2] = 0;
	v3 = a1 ^ dword_5d4594_2516348;
	*v2 = a1 ^ dword_5d4594_2516348;
	dword_5d4594_2516328 ^= v3;
	v2[1] = a2;
	v4 = a2 ^ dword_5d4594_2516348;
	v2[1] = a2 ^ dword_5d4594_2516348;
	dword_5d4594_2516328 ^= v4;
	return sub_56F2F0(v2);
}

//----- (0056F2F0) --------------------------------------------------------
int  sub_56F2F0(_DWORD* a1) {
	int v1;     // esi
	__int16 v2; // di
	int result; // eax
	__int16 i;  // ax
	int v5;     // eax

	v1 = dword_5d4594_2516344;
	v2 = 0;
	if (*getMemU16Ptr(0x587000, 311204)) {
		for (i = nox_common_randomInt_415FA0(0, *getMemU16Ptr(0x587000, 311204) - 1); v1; ++v2) {
			if (v2 == i)
				break;
			v1 = *(_DWORD*)(v1 + 8);
		}
		a1[3] = *(_DWORD*)(v1 + 12);
		if (!dword_5d4594_2516348)
			nullsub_31(1);
		a1[2] = v1;
		*(_DWORD*)(v1 + 12) = a1;
		v5 = a1[3];
		if (v5) {
			*(_DWORD*)(v5 + 8) = a1;
			++*getMemU16Ptr(0x587000, 311204);
		} else {
			++*getMemU16Ptr(0x587000, 311204);
			dword_5d4594_2516344 = a1;
		}
		result = 1;
	} else {
		++*getMemU16Ptr(0x587000, 311204);
		dword_5d4594_2516352 = a1;
		dword_5d4594_2516344 = a1;
		result = 1;
	}
	return result;
}
// 560840: using guessed type void  nullsub_31(_DWORD);

//----- (0056F3B0) --------------------------------------------------------
_DWORD* sub_56F3B0() {
	_DWORD* result; // eax
	_DWORD* v1;     // esi

	result = *(_DWORD**)&dword_5d4594_2516344;
	if (dword_5d4594_2516344) {
		do {
			v1 = (_DWORD*)result[2];
			free(result);
			result = v1;
		} while (v1);
	}
	dword_5d4594_2516328 = 0;
	*getMemU16Ptr(0x587000, 311204) = 0;
	dword_5d4594_2516348 = 0;
	dword_5d4594_2516352 = 0;
	dword_5d4594_2516344 = 0;
	return result;
}

//----- (0056F400) --------------------------------------------------------
int  nox_xxx_protectionCreateInt_56F400(int a1) {
	if (nox_xxx_protectionCreateStructForInt_56F280(*(int*)&dword_5d4594_2516356, a1))
		return (dword_5d4594_2516356)++;
	nullsub_31(1);
	return 0;
}
// 560840: using guessed type void  nullsub_31(_DWORD);

//----- (0056F440) --------------------------------------------------------
int  nox_xxx_protectionCreateFloat_56F440(int a1) {
	if (nox_xxx_protectionCreateStructForFloat_56F480(*(int*)&dword_5d4594_2516356, a1))
		return (dword_5d4594_2516356)++;
	nullsub_31(1);
	return 0;
}
// 560840: using guessed type void  nullsub_31(_DWORD);

//----- (0056F480) --------------------------------------------------------
int  nox_xxx_protectionCreateStructForFloat_56F480(int a1, int a2) {
	_DWORD* v2; // eax
	int v3;     // ecx
	int v4;     // ecx

	v2 = calloc(1u, 0x10u);
	if (!v2)
		return 0;
	v2[3] = 0;
	v2[2] = 0;
	v3 = a1 ^ dword_5d4594_2516348;
	*v2 = a1 ^ dword_5d4594_2516348;
	dword_5d4594_2516328 ^= v3;
	v2[1] = a2;
	v4 = a2 ^ dword_5d4594_2516348;
	v2[1] = a2 ^ dword_5d4594_2516348;
	dword_5d4594_2516328 ^= v4;
	return sub_56F2F0(v2);
}

//----- (0056F4F0) --------------------------------------------------------
int  sub_56F4F0(int* a1) {
	int result; // eax

	result = sub_56F510(*a1);
	if (result)
		*a1 = 0;
	return result;
}

//----- (0056F510) --------------------------------------------------------
int  sub_56F510(int a1) {
	_DWORD* v1; // eax
	int v2;     // ecx
	int v3;     // ecx
	int v4;     // ecx
	int v5;     // ecx

	v1 = sub_56F590(a1);
	if (!v1)
		return 0;
	v2 = v1[3];
	if (v2)
		*(_DWORD*)(v2 + 8) = v1[2];
	else
		dword_5d4594_2516344 = v1[2];
	v3 = v1[2];
	if (v3)
		*(_DWORD*)(v3 + 12) = v1[3];
	else
		dword_5d4594_2516352 = v1[3];
	v4 = *v1 ^ dword_5d4594_2516328;
	dword_5d4594_2516328 = v4;
	v5 = v1[1] ^ v4;
	--*getMemU16Ptr(0x587000, 311204);
	dword_5d4594_2516328 = v5;
	free(v1);
	return 1;
}

//----- (0056F590) --------------------------------------------------------
_DWORD*  sub_56F590(int a1) {
	_DWORD* result; // eax

	result = *(_DWORD**)&dword_5d4594_2516344;
	if (dword_5d4594_2516344) {
		while (*result != (a1 ^ dword_5d4594_2516348)) {
			result = (_DWORD*)result[2];
			if (!result)
				goto LABEL_4;
		}
	} else {
	LABEL_4:
		nullsub_31(1);
		result = 0;
	}
	return result;
}
// 560840: using guessed type void  nullsub_31(_DWORD);

//----- (0056F5C0) --------------------------------------------------------
int nox_xxx_protectData_56F5C0() {
	int v0;          // ebx
	int v1;          // edi
	int v2;          // ebx
	int v3;          // edi
	unsigned int v4; // eax
	int i;           // ebp
	int v6;          // esi
	int v7;          // eax
	int* v8;         // eax
	int* v9;         // eax
	int v10;         // ecx
	int v11;         // ebp
	bool v12;        // zf
	int result;      // eax
	int* v14;        // [esp-14h] [ebp-14h]

	if (!dword_5d4594_2516348)
		nullsub_31(1);
	v0 = nox_frame_xxx_2598000;
	v1 = dword_5d4594_2516348;
	v2 = nox_xxx_protect_56F240() ^ v0;
	v3 = v2 ^ v1;
	dword_5d4594_2516328 = ~v2;
	v4 = *getMemU16Ptr(0x587000, 311204);
	for (i = 0; i < (*getMemU16Ptr(0x587000, 311204) >> 2); v4 = *getMemU16Ptr(0x587000, 311204)) {
		v6 = nox_common_randomInt_415FA0(0, v4 >> 1);
		v7 = nox_common_randomInt_415FA0((*getMemU16Ptr(0x587000, 311204) >> 1) + 1,
										 *getMemU16Ptr(0x587000, 311204) - 1);
		if (v6 != v7) {
			v14 = sub_56F6F0(v7);
			v8 = sub_56F6F0(v6);
			sub_56F720(v8, v14);
		}
		++i;
	}
	v9 = *(int**)&dword_5d4594_2516344;
	dword_5d4594_2516348 = 0;
	if (dword_5d4594_2516344) {
		do {
			v10 = v3 ^ *v9;
			v11 = v3 ^ v9[1];
			*v9 = v10;
			v9[1] = v11;
			dword_5d4594_2516328 ^= v10;
			dword_5d4594_2516328 ^= v9[1];
			v9 = (int*)v9[2];
		} while (v9);
	}
	result = v2 ^ dword_5d4594_2516348;
	v12 = v2 == dword_5d4594_2516348;
	++*getMemU32Ptr(0x5D4594, 2516364);
	dword_5d4594_2516348 ^= v2;
	if (v12)
		nullsub_31(1);
	return result;
}
// 560840: using guessed type void  nullsub_31(_DWORD);

//----- (0056F6F0) --------------------------------------------------------
_DWORD*  sub_56F6F0(int a1) {
	_DWORD* result; // eax
	int v2;         // ecx

	result = *(_DWORD**)&dword_5d4594_2516344;
	v2 = 0;
	if (dword_5d4594_2516344) {
		while (v2 != a1) {
			result = (_DWORD*)result[2];
			++v2;
			if (!result)
				goto LABEL_4;
		}
	} else {
	LABEL_4:
		nullsub_31(1);
		result = 0;
	}
	return result;
}
// 560840: using guessed type void  nullsub_31(_DWORD);

//----- (0056F720) --------------------------------------------------------
void  sub_56F720(int* a1, int* a2) {
	int v2; // edx
	int v3; // esi

	if (!a1 || !a2) {
		nullsub_31(1);
		return;
	}
	v2 = *a1;
	v3 = a1[1];
	*a1 = *a2;
	a1[1] = a2[1];
	*a2 = v2;
	a2[1] = v3;
	++*getMemU32Ptr(0x5D4594, 2516360);
	if (!dword_5d4594_2516348)
		nullsub_31(1);
}
// 560840: using guessed type void  nullsub_31(_DWORD);

//----- (0056F780) --------------------------------------------------------
_DWORD*  sub_56F780(int a1, int a2) {
	_DWORD* result; // eax
	int v3;         // ecx

	result = (_DWORD*)a1;
	if (a1 >= 657757279) {
		result = sub_56F590(a1);
		if (result) {
			dword_5d4594_2516328 ^= result[1];
			v3 = a2 ^ dword_5d4594_2516348;
			result[1] = a2 ^ dword_5d4594_2516348;
			dword_5d4594_2516328 ^= v3;
			result = (_DWORD*)nox_xxx_protectData_56F5C0();
		}
	}
	return result;
}

//----- (0056F7D0) --------------------------------------------------------
_DWORD*  nox_xxx_playerUpdateNetBuffs_56F7D0(int a1, int a2) {
	_DWORD* result; // eax
	int v3;         // ecx

	result = (_DWORD*)a1;
	if (a1 >= 657757279) {
		result = sub_56F590(a1);
		if (result) {
			dword_5d4594_2516328 ^= result[1];
			v3 = a2 ^ dword_5d4594_2516348;
			result[1] = a2 ^ dword_5d4594_2516348;
			dword_5d4594_2516328 ^= v3;
			result = (_DWORD*)nox_xxx_protectData_56F5C0();
		}
	}
	return result;
}

//----- (0056F820) --------------------------------------------------------
_DWORD*  sub_56F820(int a1, unsigned __int8 a2) {
	_DWORD* result; // eax
	int v3;         // ecx

	result = (_DWORD*)a1;
	if (a1 >= 657757279) {
		result = sub_56F590(a1);
		if (result) {
			dword_5d4594_2516328 ^= result[1];
			v3 = dword_5d4594_2516348 ^ a2;
			result[1] = v3;
			dword_5d4594_2516328 ^= v3;
			result = (_DWORD*)nox_xxx_protectData_56F5C0();
		}
	}
	return result;
}

//----- (0056F870) --------------------------------------------------------
_DWORD*  nox_xxx_protectPlayerHPMana_56F870(int a1, unsigned __int16 a2) {
	_DWORD* result; // eax
	int v3;         // ecx

	result = (_DWORD*)a1;
	if (a1 >= 657757279) {
		result = sub_56F590(a1);
		if (result) {
			dword_5d4594_2516328 ^= result[1];
			v3 = dword_5d4594_2516348 ^ a2;
			result[1] = v3;
			dword_5d4594_2516328 ^= v3;
			result = (_DWORD*)nox_xxx_protectData_56F5C0();
		}
	}
	return result;
}

//----- (0056F8C0) --------------------------------------------------------
_DWORD*  sub_56F8C0(int a1, float a2) {
	_DWORD* result; // eax
	_DWORD* v3;     // esi
	int v4;         // eax

	result = (_DWORD*)a1;
	if (a1 >= 657757279) {
		result = sub_56F590(a1);
		v3 = result;
		if (result) {
			dword_5d4594_2516328 ^= result[1];
			v4 = dword_5d4594_2516348 ^ (unsigned __int64)(__int64)a2;
			v3[1] = v4;
			dword_5d4594_2516328 ^= v4;
			result = (_DWORD*)nox_xxx_protectData_56F5C0();
		}
	}
	return result;
}

//----- (0056F920) --------------------------------------------------------
_DWORD*  sub_56F920(int a1, int a2) {
	_DWORD* result; // eax
	int v3;         // ecx

	result = (_DWORD*)a1;
	if (a1 >= 657757279) {
		result = sub_56F590(a1);
		if (result) {
			dword_5d4594_2516328 ^= result[1];
			v3 = dword_5d4594_2516348 ^ (a2 + (dword_5d4594_2516348 ^ result[1]));
			result[1] = v3;
			dword_5d4594_2516328 ^= v3;
			result = (_DWORD*)nox_xxx_protectData_56F5C0();
		}
	}
	return result;
}

//----- (0056F9E0) --------------------------------------------------------
_DWORD*  nox_xxx_protectMana_56F9E0(int a1, __int16 a2) {
	_DWORD* result; // eax
	int v3;         // ecx

	result = (_DWORD*)a1;
	if (a1 >= 657757279) {
		result = sub_56F590(a1);
		if (result) {
			dword_5d4594_2516328 ^= result[1];
			v3 = dword_5d4594_2516348 ^ (a2 + (dword_5d4594_2516348 ^ result[1]));
			result[1] = v3;
			dword_5d4594_2516328 ^= v3;
			result = (_DWORD*)nox_xxx_protectData_56F5C0();
		}
	}
	return result;
}

//----- (0056FA40) --------------------------------------------------------
_DWORD*  sub_56FA40(int a1, float a2) {
	_DWORD* result; // eax
	_DWORD* v3;     // esi
	int v4;         // eax

	result = (_DWORD*)a1;
	if (a1 >= 657757279) {
		result = sub_56F590(a1);
		v3 = result;
		if (result) {
			dword_5d4594_2516328 ^= result[1];
			v4 = dword_5d4594_2516348 ^
				 (unsigned __int64)(__int64)((double)(unsigned int)(dword_5d4594_2516348 ^ result[1]) + a2);
			v3[1] = v4;
			dword_5d4594_2516328 ^= v4;
			result = (_DWORD*)nox_xxx_protectData_56F5C0();
		}
	}
	return result;
}

//----- (0056FAC0) --------------------------------------------------------
int  nox_xxx_protectionStringCRC_56FAC0(int* a1, unsigned int a2) {
	int* v2;        // ecx
	int result;     // eax
	unsigned int i; // edx
	int v5;         // esi

	v2 = a1;
	result = 0;
	for (i = a2 >> 2; i; --i) {
		v5 = *v2;
		++v2;
		result ^= v5;
	}
	return result;
}

//----- (0056FAE0) --------------------------------------------------------
int  nox_xxx_protectionStringCRCLen_56FAE0(int* a1, unsigned int a2) {
	int result; // eax

	result = 0;
	if (a1)
		result = nox_xxx_protectionStringCRC_56FAC0(a1, a2);
	return result;
}

//----- (0056FB00) --------------------------------------------------------
int  sub_56FB00(int* a1, unsigned int a2, int a3) {
	_DWORD* v3; // esi

	if (a3 >= 657757279) {
		v3 = sub_56F590(a3);
		if (v3 && (dword_5d4594_2516348 ^ nox_xxx_protectionStringCRCLen_56FAE0(a1, a2)) == v3[1])
			return 1;
		nullsub_31(1);
	}
	return 0;
}
// 560840: using guessed type void  nullsub_31(_DWORD);

//----- (0056FB60) --------------------------------------------------------
int  sub_56FB60(int* a1) {
	int result; // eax
	int v2;     // ebx
	int v3;     // ebx
	int v4;     // ebx
	int* v5;    // edi
	int v6;     // eax
	int* v7;    // eax

	result = 0;
	if (a1) {
		v2 = sub_4E4C00((int)a1);
		v3 = (unsigned __int16)nox_xxx_unitGetHP_4EE780((int)a1) ^ v2;
		v4 = sub_4E4C10((int)a1) ^ v3;
		v5 = (int*)sub_4E4C30((int)a1);
		v6 = sub_4E4C50((int)a1);
		if (v5 && v6 > 0)
			v4 ^= nox_xxx_protectionStringCRC_56FAC0(v5, v6);
		v7 = (int*)sub_4E4C80(a1);
		if (v7) {
			if (strlen((const char*)v7))
				v4 ^= nox_xxx_protectionStringCRC_56FAC0(v7, strlen((const char*)v7));
		}
		result = v4;
	}
	return result;
}

//----- (0056FBF0) --------------------------------------------------------
int  nox_xxx_protect_56FBF0(int a1, int* a2) {
	int result; // eax
	_DWORD* v3; // eax
	_DWORD* v4; // esi
	int v5;     // ecx

	result = a1;
	if (a1 >= 657757279) {
		v3 = sub_56F590(a1);
		v4 = v3;
		if (v3) {
			dword_5d4594_2516328 ^= v3[1];
			v5 = sub_56FB60(a2) ^ v3[1];
			v4[1] = v5;
			result = v5 ^ dword_5d4594_2516328;
			dword_5d4594_2516328 ^= v5;
		} else {
			nullsub_31(1);
		}
	}
	return result;
}
// 560840: using guessed type void  nullsub_31(_DWORD);

//----- (0056FC50) --------------------------------------------------------
int  nox_xxx_protect_56FC50(int a1, int* a2) {
	int result; // eax
	_DWORD* v3; // eax
	_DWORD* v4; // esi
	int v5;     // ecx

	result = a1;
	if (a1 >= 657757279) {
		v3 = sub_56F590(a1);
		v4 = v3;
		if (v3) {
			dword_5d4594_2516328 ^= v3[1];
			v5 = sub_56FB60(a2) ^ v3[1];
			v4[1] = v5;
			result = v5 ^ dword_5d4594_2516328;
			dword_5d4594_2516328 ^= v5;
		} else {
			nullsub_31(1);
		}
	}
	return result;
}
// 560840: using guessed type void  nullsub_31(_DWORD);

//----- (0056FCB0) --------------------------------------------------------
int  sub_56FCB0(int a1, int a2) {
	int result; // eax

	result = 0;
	if (a2)
		result = 1 << (a1 % 32);
	return result;
}

//----- (0056FCE0) --------------------------------------------------------
int  nox_xxx_playerAwardSpellProtection_56FCE0(int a1, int a2, int a3) {
	int result; // eax
	_DWORD* v4; // eax
	_DWORD* v5; // esi

	result = a1;
	if (a1 >= 657757279) {
		v4 = sub_56F590(a1);
		v5 = v4;
		if (v4) {
			dword_5d4594_2516328 ^= v4[1];
			result = dword_5d4594_2516348 ^ (dword_5d4594_2516348 ^ v4[1] | sub_56FCB0(a2, a3));
			v5[1] = result;
			dword_5d4594_2516328 ^= result;
		} else {
			nullsub_31(1);
		}
	}
	return result;
}
// 560840: using guessed type void  nullsub_31(_DWORD);

//----- (0056FD50) --------------------------------------------------------
int  sub_56FD50(int a1, int a2, int a3) {
	int v3;     // ebp
	_DWORD* v4; // eax
	int v5;     // esi
	int* v6;    // edi
	_DWORD* v8; // [esp+8h] [ebp+4h]

	v3 = 0;
	if (a1 >= 657757279) {
		v4 = sub_56F590(a1);
		v8 = v4;
		if (v4) {
			v5 = 1;
			if (a3 > 1) {
				v6 = (int*)(a2 + 4);
				do {
					v3 |= sub_56FCB0(v5++, *v6);
					++v6;
				} while (v5 < a3);
				v4 = v8;
			}
			if ((v3 ^ dword_5d4594_2516348) == v4[1])
				return 1;
		}
		nullsub_31(1);
	}
	return 0;
}
// 560840: using guessed type void  nullsub_31(_DWORD);

//----- (0056FDD0) --------------------------------------------------------
#ifndef NOX_CGO
void nox_xxx_cryptXor_56FDD0(char key, unsigned char* p, int n) {
	if (!p)
		return;
	for (int i = 0; i < n; i++) {
		p[i] ^= key;
	}
}

//----- (0056FE00) --------------------------------------------------------
void nox_xxx_cryptXorDst_56FE00(char key, unsigned char* src, int n, unsigned char* dst) {
	if (!src || !n || !dst)
		return;
	for (int i = 0; i < n; i++) {
		dst[i] = key ^ src[i];
	}
}
#endif // NOX_CGO

//----- (0056FE30) --------------------------------------------------------
double nox_xxx_unkDoubleSmth_56FE30() {
	double v0; // st7

	*getMemU64Ptr(0x5D4594, 2516412) = *getMemU64Ptr(0x5D4594, 2516404);
	*getMemU64Ptr(0x5D4594, 2516404) = *getMemU64Ptr(0x5D4594, 2516396);
	*getMemU32Ptr(0x5D4594, 2516396) = *getMemU32Ptr(0x5D4594, 2516388);
	*getMemU32Ptr(0x5D4594, 2516400) = *getMemU32Ptr(0x5D4594, 2516392);
	v0 = *getMemDoublePtr(0x5D4594, 2516388) * *getMemDoublePtr(0x581450, 11376) +
		 *getMemDoublePtr(0x5D4594, 2516404) * *getMemDoublePtr(0x581450, 11368) +
		 *getMemDoublePtr(0x5D4594, 2516412) * *getMemDoublePtr(0x581450, 11360) +
		 *getMemDoublePtr(0x5D4594, 2516412) * *getMemDoublePtr(0x581450, 11352) + *getMemDoublePtr(0x5D4594, 2516420);
	floor(v0);
	*getMemDoublePtr(0x5D4594, 2516388) = v0 - v0;
	*getMemDoublePtr(0x5D4594, 2516420) = v0 * *getMemDoublePtr(0x581450, 11344);
	return *getMemDoublePtr(0x5D4594, 2516388);
}

//----- (0056FF00) --------------------------------------------------------
void  sub_56FF00(int a1) {
	int v1;              // eax
	unsigned __int8* v2; // ecx
	unsigned int v3;     // eax
	int v4;              // esi

	v1 = a1;
	if (!a1)
		v1 = -1;
	v2 = getMemAt(0x5D4594, 2516388);
	do {
		v2 += 8;
		v3 = (((v1 << 13) ^ (unsigned int)v1) >> 17) ^ (v1 << 13) ^ v1;
		v1 = (32 * v3) ^ v3;
		*((double*)v2 - 1) = (double)(unsigned int)v1 * *getMemDoublePtr(0x581450, 11344);
	} while ((int)v2 < (int)getMemAt(0x5D4594, 2516428));
	v4 = 19;
	do {
		nox_xxx_unkDoubleSmth_56FE30();
		--v4;
	} while (v4);
	dword_5d4594_2516380 = 0;
	*getMemU32Ptr(0x5D4594, 2516376) = 99;
	dword_5d4594_2516372 = 100;
}

//----- (0056FF80) --------------------------------------------------------
int  sub_56FF80(int a1, int a2) {
	__int64 v2; // rax
	int result; // eax

	*getMemU32Ptr(0x5D4594, 2516376) = a2;
	dword_5d4594_2516380 = a1;
	dword_5d4594_2516372 = a2 - a1 + 1;
	v2 = (__int64)(nox_xxx_unkDoubleSmth_56FE30() * (double)*(unsigned int*)&dword_5d4594_2516372);
	if ((unsigned int)v2 < *(int*)&dword_5d4594_2516372)
		result = dword_5d4594_2516380 + v2;
	else
		result = dword_5d4594_2516372 + dword_5d4594_2516380;
	return result;
}

//----- (00578AC0) --------------------------------------------------------
unsigned int  nox_xxx_netGetUnitCodeServ_578AC0(_DWORD* a1) {
	unsigned int result; // eax

	if (!a1)
		return 0;
	if (a1[9] >= 0x8000u)
		return 0;
	result = a1[10];
	if (result >= 0x8000)
		return 0;
	if (!(a1[2] & 0x20400000))
		return a1[9];
	BYTE1(result) |= 0x80u;
	return result;
}

//----- (00578B00) --------------------------------------------------------
unsigned int  nox_xxx_netGetUnitCodeCli_578B00(int a1) {
	unsigned int result; // eax

	if (!a1)
		return 0;
	result = *(_DWORD*)(a1 + 128);
	if (result >= 0x8000)
		return 0;
	if (*(_DWORD*)(a1 + 112) & 0x20400000)
		BYTE1(result) |= 0x80u;
	return result;
}

//----- (00578B30) --------------------------------------------------------
int  nox_xxx_netClearHighBit_578B30(__int16 a1) { return a1 & 0x7FFF; }

//----- (00578B40) --------------------------------------------------------
int  nox_xxx_packetDynamicUnitCode_578B40(int a1) {
	int result; // eax
	int v2;     // eax

	result = a1;
	if ((a1 & 0x8000) == 0x8000) {
		BYTE1(result) &= 0x7Fu;
		v2 = nox_xxx_netGetUnitByExtent_4ED020(result);
		if (v2)
			result = *(_DWORD*)(v2 + 36);
		else
			result = 0;
	}
	return result;
}

//----- (00578B70) --------------------------------------------------------
unsigned int  nox_xxx_netTestHighBit_578B70(unsigned int a1) { return (a1 >> 15) & 1; }

//----- (00578B80) --------------------------------------------------------
_DWORD* sub_578B80() {
	_DWORD* v0; // eax
	_DWORD* v2; // esi
	int v3;     // eax

	v0 = operator_new(0x18u);
	if (!v0)
		return 0;
	v2 = v0;
	*v0 = operator_new(0x10000u);
	v2[1] = 0;
	sub_57DF00(v2 + 2);
	v2[5] = operator_new(0x10000u);
	v3 = 0;
	do {
		v3 += 2;
		*(_WORD*)(v3 + v2[5] - 2) = -1;
	} while (v3 < 0x10000);
	return v2;
}
// 5667CB: using guessed type void * operator_new(unsigned int);

//----- (00578BA0) --------------------------------------------------------
unsigned int  sub_578BA0(unsigned int a1) { return (a1 >> 1) + a1 + 32; }

//----- (00578BB0) --------------------------------------------------------
int  sub_578BB0(void** a1, int a2, unsigned __int8* a3, int a4) { return sub_57D1C0(a1, a2, a3, a4); }

//----- (00578BD0) --------------------------------------------------------
void  sub_578BD0(LPVOID lpMem) {
	if (lpMem) {
		sub_57D150((LPVOID*)lpMem);
		operator_delete(lpMem);
	}
}

//----- (00578BF0) --------------------------------------------------------
_DWORD* sub_578BF0() {
	_DWORD* v0;     // eax
	_DWORD* result; // eax

	v0 = operator_new(0x98u);
	if (v0)
		result = sub_57E9A0(v0);
	else
		result = 0;
	return result;
}
// 5667CB: using guessed type void * operator_new(unsigned int);

//----- (00578C10) --------------------------------------------------------
int  nox_xxx_nxzDecompress_578C10(_DWORD* a1, _BYTE* a2, _DWORD* a3, unsigned int a4, _DWORD* a5) {
	return nox_xxx_nxzDecompressImpl_57EA80(a1, a2, a3, a4, a5);
}

//----- (00578C30) --------------------------------------------------------
int  sub_578C30(int a1) { return sub_57EA60(a1); }

//----- (00578C40) --------------------------------------------------------
void  sub_578C40(LPVOID lpMem) {
	if (lpMem) {
		sub_57EA00((LPVOID*)lpMem);
		operator_delete(lpMem);
	}
}

//----- (00578C60) --------------------------------------------------------
int sub_578C60() {
	if (sub_44E560()) {
		nox_client_lockScreenBriefing_450160(255, 1, 0);
		sub_4A2530();
	}
	return 1;
}

//----- (00578C90) --------------------------------------------------------
int  sub_578C90(int a1) {
	dword_587000_311372 = a1;
	*getMemU8Ptr( 0x5D4594, 2516476) |= 1 << a1;
	nox_xxx_cliPlayMapIntro_44E0B0(1);
	sub_413960();
	sub_477530(0);
	return nox_xxx_quit_4460C0();
}

//----- (00578CD0) --------------------------------------------------------
int sub_578CD0() {
	int result;         // eax
	unsigned __int8 v1; // dl
	char* v2;           // edi
	char v3[16];        // [esp+0h] [ebp-90h]
	char v4[128];       // [esp+10h] [ebp-80h]

	result = dword_587000_311372;
	if (*(int*)&dword_587000_311372 != -1) {
		v1 = getMemByte(0x587000, 311380);
		strcpy(v3, *(const char**)getMemAt(0x587000, 4 * dword_587000_311372 + 29456));
		v2 = &v3[strlen(v3)];
		*(_DWORD*)v2 = *getMemU32Ptr(0x587000, 311376);
		v2[4] = v1;
		if (nox_game_setMovieFile_4CB230(v3, v4)) {
			sub_4B0300(v4);
			sub_4B0640(sub_578C60);
			result = nox_client_drawGeneral_4B0340(1);
		} else {
			result = sub_578C60();
		}
	}
	return result;
}
// 578CD0: using guessed type char var_90[16];

//----- (00578D80) --------------------------------------------------------
char* nox_xxx_GetEndgameDialog_578D80() {
	if (dword_587000_311372) {
		if (dword_587000_311372 == 1) {
			if (!(getMemByte(0x5D4594, 2516476) & 1))
				return (char*)getMemAt(0x587000, 311416);
			if (!(getMemByte(0x5D4594, 2516476) & 4))
				return (char*)getMemAt(0x587000, 311432);
		} else if (dword_587000_311372 == 2) {
			if (!(getMemByte(0x5D4594, 2516476) & 2))
				return (char*)getMemAt(0x587000, 311448);
			if (!(getMemByte(0x5D4594, 2516476) & 1))
				return (char*)getMemAt(0x587000, 311464);
		}
	} else {
		if (!(getMemByte(0x5D4594, 2516476) & 2))
			return (char*)getMemAt(0x587000, 311384);
		if (!(getMemByte(0x5D4594, 2516476) & 4))
			return (char*)getMemAt(0x587000, 311400);
	}
	return 0;
}

//----- (00578DE0) --------------------------------------------------------
char  sub_578DE0(char a1) {
	char result; // al

	result = a1;
	*getMemU8Ptr(0x5D4594, 2516476) = a1;
	return result;
}

//----- (00578DF0) --------------------------------------------------------
unsigned __int8 sub_578DF0() { return getMemByte(0x5D4594, 2516476); }

//----- (00578E00) --------------------------------------------------------
void sub_578E00() { dword_587000_311372 = -1; }

//----- (00579860) --------------------------------------------------------
LPVOID nox_xxx_waypointGetList_579860() { return *(LPVOID*)&dword_5d4594_2523752; }

//----- (00579870) --------------------------------------------------------
int  nox_xxx_waypointNext_579870(int a1) {
	int result; // eax

	if (a1)
		result = *(_DWORD*)(a1 + 484);
	else
		result = 0;
	return result;
}

//----- (00579890) --------------------------------------------------------
int sub_579890() { return dword_5d4594_2523756; }

//----- (005798A0) --------------------------------------------------------
int  sub_5798A0(int a1) {
	int result; // eax

	if (a1)
		result = *(_DWORD*)(a1 + 484);
	else
		result = 0;
	return result;
}

//----- (005798C0) --------------------------------------------------------
unsigned int nox_xxx_waypoint_5798C0() {
	_DWORD* v0;          // ecx
	unsigned int result; // eax

	v0 = *(_DWORD**)&dword_5d4594_2523752;
	result = 1;
	if (dword_5d4594_2523752) {
		do {
			if (result <= *v0)
				result = *v0 + 1;
			v0 = (_DWORD*)v0[121];
		} while (v0);
	}
	return result;
}

//----- (005798F0) --------------------------------------------------------
_DWORD*  nox_xxx_waypointNew_5798F0(float a1, float a2) {
	_DWORD* v2;      // esi
	unsigned int v3; // eax
	int v4;          // edx

	v2 = calloc(1u, 0x204u);
	v3 = nox_xxx_waypoint_5798C0();
	v4 = v2[120];
	*v2 = v3;
	*((float*)v2 + 2) = a1;
	*((float*)v2 + 3) = a2;
	v2[120] = v4 | 1;
	v2[121] = dword_5d4594_2523752;
	if (dword_5d4594_2523752)
		*(_DWORD*)(dword_5d4594_2523752 + 488) = v2;
	dword_5d4594_2523752 = v2;
	if (nox_common_gameFlags_check_40A5C0(1))
		nox_xxx_waypointMapRegister_5179B0((int)v2);
	return v2;
}

//----- (00579970) --------------------------------------------------------
float*  nox_xxx_waypointNewNotMap_579970(int a1, float a2, float a3) {
	float* result; // eax
	int v4;        // ecx

	result = (float*)calloc(1u, 0x204u);
	*(_DWORD*)result = a1;
	result[3] = a3;
	v4 = *((_DWORD*)result + 120) | 1;
	result[2] = a2;
	*((_DWORD*)result + 120) = v4;
	result[121] = *(float*)&dword_5d4594_2523756;
	dword_5d4594_2523756 = result;
	return result;
}

//----- (005799C0) --------------------------------------------------------
char* nox_xxx_waypoint_5799C0() {
	int v0; // esi
	int v1; // edi

	v0 = dword_5d4594_2523756;
	if (dword_5d4594_2523756) {
		do {
			v1 = *(_DWORD*)(v0 + 484);
			*(_DWORD*)(v0 + 484) = dword_5d4594_2523752;
			if (dword_5d4594_2523752)
				*(_DWORD*)(dword_5d4594_2523752 + 488) = v0;
			dword_5d4594_2523752 = v0;
			if (nox_common_gameFlags_check_40A5C0(1))
				nox_xxx_waypointMapRegister_5179B0(v0);
			v0 = v1;
		} while (v1);
	}
	dword_5d4594_2523756 = 0;
	return sub_579A30();
}

//----- (00579A30) --------------------------------------------------------
char* sub_579A30() {
	char* result; // eax
	char* i;      // esi
	int v2;       // eax
	char* v3;     // ecx
	char v4;      // dl
	char* j;      // eax
	int v6;       // edi
	_BYTE* v7;    // ecx

	result = (char*)nox_xxx_waypointGetList_579860();
	for (i = result; result; i = result) {
		i[477] = 0;
		v2 = 0;
		if (i[476]) {
			v3 = i + 96;
			do {
				v4 = *v3;
				v3 += 8;
				i[477] |= v4;
				++v2;
			} while (v2 < (unsigned __int8)i[476]);
		}
		for (j = (char*)nox_xxx_waypointGetList_579860(); j; j = (char*)nox_xxx_waypointNext_579870((int)j)) {
			v6 = 0;
			if (j[476]) {
				v7 = j + 96;
				do {
					if (*((char**)v7 - 1) == i)
						i[477] |= *v7;
					++v6;
					v7 += 8;
				} while (v6 < (unsigned __int8)j[476]);
			}
		}
		result = (char*)nox_xxx_waypointNext_579870((int)i);
	}
	return result;
}

//----- (00579AD0) --------------------------------------------------------
float*  sub_579AD0(float a1, float a2) {
	int v2;    // ecx
	int v3;    // edx
	double v4; // st7
	double v5; // st6
	double v6; // st5
	float i;   // [esp+0h] [ebp-4h]

	v2 = dword_5d4594_2523752;
	v3 = 0;
	for (i = 100.0; v2; v2 = *(_DWORD*)(v2 + 484)) {
		v4 = *(float*)(v2 + 8) - a1;
		v5 = *(float*)(v2 + 12) - a2;
		v6 = v5 * v5 + v4 * v4;
		if (v6 < i) {
			i = v6;
			v3 = v2;
		}
	}
	return (float*)v3;
}

//----- (00579B30) --------------------------------------------------------
void  sub_579B30(LPVOID lpMem) {
	int i;      // ecx
	int v2;     // eax
	int v3;     // edx
	LPVOID* v4; // esi
	int v5;     // esi
	_DWORD* v6; // eax
	int v7;     // eax
	int v8;     // eax

	for (i = dword_5d4594_2523752; i; i = *(_DWORD*)(i + 484)) {
		v2 = 0;
		v3 = *(unsigned __int8*)(i + 476);
		if (v3 > 0) {
			v4 = (LPVOID*)(i + 92);
			while (*v4 != lpMem) {
				++v2;
				v4 += 2;
				if (v2 >= v3)
					goto LABEL_11;
			}
			v5 = v2;
			if (v2 < v3 - 1) {
				v6 = (_DWORD*)(i + 8 * v2 + 92);
				do {
					++v5;
					*v6 = v6[2];
					v6[1] = v6[3];
					v6 += 2;
				} while (v5 < *(unsigned __int8*)(i + 476) - 1);
			}
			--*(_BYTE*)(i + 476);
		}
	LABEL_11:;
	}
	v7 = *((_DWORD*)lpMem + 121);
	if (v7)
		*(_DWORD*)(v7 + 488) = *((_DWORD*)lpMem + 122);
	v8 = *((_DWORD*)lpMem + 122);
	if (v8)
		*(_DWORD*)(v8 + 484) = *((_DWORD*)lpMem + 121);
	else
		dword_5d4594_2523752 = *((_DWORD*)lpMem + 121);
	if (nox_common_gameFlags_check_40A5C0(1))
		sub_517A70((int)lpMem);
	free(lpMem);
}

//----- (00579C00) --------------------------------------------------------
_DWORD* sub_579C00() {
	_DWORD* result; // eax
	_DWORD* v1;     // esi

	result = *(_DWORD**)&dword_5d4594_2523752;
	if (dword_5d4594_2523752) {
		do {
			v1 = (_DWORD*)result[121];
			free(result);
			result = v1;
		} while (v1);
		dword_5d4594_2523752 = 0;
	} else {
		dword_5d4594_2523752 = 0;
	}
	return result;
}

//----- (00579C40) --------------------------------------------------------
_DWORD*  nox_server_getWaypointById_579C40(int a1) {
	_DWORD* result; // eax

	result = *(_DWORD**)&dword_5d4594_2523752;
	if (!dword_5d4594_2523752)
		return 0;
	while (*result != a1) {
		result = (_DWORD*)result[121];
		if (!result)
			return 0;
	}
	return result;
}

//----- (00579C60) --------------------------------------------------------
int  sub_579C60(int a1) {
	int result; // eax

	result = dword_5d4594_2523756;
	if (!dword_5d4594_2523756)
		return 0;
	while (*(_DWORD*)(result + 4) != a1) {
		result = *(_DWORD*)(result + 484);
		if (!result)
			return 0;
	}
	return result;
}

//----- (00579C80) --------------------------------------------------------
_DWORD*  sub_579C80(int a1) {
	_DWORD* result; // eax

	result = *(_DWORD**)&dword_5d4594_2523756;
	if (!dword_5d4594_2523756)
		return 0;
	while (*result != a1) {
		result = (_DWORD*)result[121];
		if (!result)
			return 0;
	}
	return result;
}

//----- (00579CA0) --------------------------------------------------------
int sub_579CA0() {
	_DWORD* v0; // eax
	_DWORD* v1; // esi
	int v2;     // ebp
	int* v3;    // ebx
	int* v4;    // edi
	int v5;     // eax

	v0 = *(_DWORD**)&dword_5d4594_2523756;
	if (dword_5d4594_2523756) {
		do {
			v0[1] = *v0;
			v0 = (_DWORD*)v0[121];
		} while (v0);
		v0 = *(_DWORD**)&dword_5d4594_2523756;
	}
	v1 = v0;
	if (!v0)
		return 1;
	while (1) {
		v2 = 0;
		if (*((_BYTE*)v1 + 476))
			break;
	LABEL_9:
		v1 = (_DWORD*)v1[121];
		if (!v1)
			return 1;
	}
	v3 = v1 + 23;
	v4 = v1 + 87;
	while (1) {
		v5 = sub_579C60(*v4);
		*v3 = v5;
		if (!v5)
			return 0;
		++v2;
		++v4;
		v3 += 2;
		if (v2 >= *((unsigned __int8*)v1 + 476))
			goto LABEL_9;
	}
}

//----- (00579D20) --------------------------------------------------------
int sub_579D20() {
	unsigned int v0;  // eax
	unsigned int* v1; // ecx
	unsigned int v2;  // edx
	unsigned int* v3; // esi
	int v4;           // ebx
	int* v5;          // ebp
	int* v6;          // edi
	int v7;           // eax
	char v9;          // [esp+0h] [ebp-4h]

	v0 = nox_xxx_waypoint_5798C0();
	v1 = *(unsigned int**)&dword_5d4594_2523756;
	if (dword_5d4594_2523756) {
		do {
			v2 = *v1;
			*v1 = v0;
			v1[1] = v2;
			v1 = (unsigned int*)v1[121];
			++v0;
		} while (v1);
		v1 = *(unsigned int**)&dword_5d4594_2523756;
	}
	v3 = v1;
	if (!v1)
		return 1;
	do {
		v4 = 0;
		v9 = 0;
		if (*((_BYTE*)v3 + 476)) {
			v5 = (int*)(v3 + 23);
			v6 = (int*)(v3 + 87);
			do {
				v7 = sub_579C60(*v6);
				*v5 = v7;
				if (v7) {
					v5 += 2;
					++v9;
				}
				++v4;
				++v6;
			} while (v4 < *((unsigned __int8*)v3 + 476));
		}
		*((_BYTE*)v3 + 476) = v9;
		v3 = (unsigned int*)v3[121];
	} while (v3);
	return 1;
}

//----- (00579DD0) --------------------------------------------------------
void nox_xxx_waypointDeleteAll_579DD0() {
	_DWORD* v0; // esi
	_DWORD* v1; // edi

	v0 = *(_DWORD**)&dword_5d4594_2523752;
	if (dword_5d4594_2523752) {
		do {
			v1 = (_DWORD*)v0[121];
			if (nox_common_gameFlags_check_40A5C0(1))
				sub_517A70((int)v0);
			free(v0);
			v0 = v1;
		} while (v1);
	}
	dword_5d4594_2523752 = 0;
	dword_5d4594_2523756 = 0;
	dword_5d4594_2523760 = 0;
}

//----- (00579E30) --------------------------------------------------------
const char*  nox_xxx_waypointByName_579E30(const char* a1) {
	const char* i; // esi

	for (i = (const char*)nox_xxx_waypointGetList_579860(); i; i = (const char*)nox_xxx_waypointNext_579870((int)i)) {
		if (nox_server_strcmpWithoutMapname_4DA3F0(i + 16, a1))
			break;
	}
	return i;
}

//----- (00579E70) --------------------------------------------------------
_DWORD* sub_579E70() {
	_DWORD* result; // eax

	result = calloc(1u, 0x204u);
	if (result)
		result[120] |= 0x1000000u;
	return result;
}

//----- (00579E90) --------------------------------------------------------
void  sub_579E90(int a1) {
	*(_DWORD*)(a1 + 480) |= 0x1000000u;
	*(_DWORD*)(a1 + 484) = dword_5d4594_2523756;
	if (dword_5d4594_2523756)
		*(_DWORD*)(dword_5d4594_2523756 + 488) = a1;
	dword_5d4594_2523756 = a1;
	if (nox_common_gameFlags_check_40A5C0(1))
		nox_xxx_waypointMapRegister_5179B0(a1);
}

//----- (00579EE0) --------------------------------------------------------
BOOL  sub_579EE0(int a1, unsigned __int8 a2) { return (a2 & *(_BYTE*)(a1 + 477)) != 0; }

//----- (00579F00) --------------------------------------------------------
int  nox_xxx_waypoint_579F00(_DWORD* a1, int a2) {
	int v2;     // ebp
	int v3;     // esi
	_BYTE* i;   // esi
	int v5;     // edi
	LPVOID v6;  // esi
	float* v7;  // edx
	float v8;   // eax
	float v9;   // ecx
	float v10;  // edx
	float4 v12; // [esp+10h] [ebp-10h]

	v2 = 0;
	if (nox_common_gameFlags_check_40A5C0(32)) {
		if (a2) {
			v3 = nox_server_getFirstObject_4DA790();
			if (v3) {
				while (!(*(_DWORD*)(v3 + 8) & 0x10000000) || nox_xxx_servCompareTeams_419150(a2 + 48, v3 + 48)) {
					v3 = nox_server_getNextObject_4DA7A0(v3);
					if (!v3)
						goto LABEL_9;
				}
				v2 = v3;
			}
		}
	}
LABEL_9:
	dword_5d4594_2523760 = 0;
	for (i = nox_xxx_waypointGetList_579860(); i; i = (_BYTE*)nox_xxx_waypointNext_579870((int)i)) {
		if (sub_579EE0((int)i, 0x80u) && i[480] & 1)
			++dword_5d4594_2523760;
	}
	if (!dword_5d4594_2523760)
		return 0;
	v5 = nox_common_randomInt_415FA0(0, dword_5d4594_2523760 - 1);
	v6 = nox_xxx_waypointGetList_579860();
	if (!v6)
		return 0;
	while (1) {
		if (!sub_579EE0((int)v6, 0x80u))
			goto LABEL_24;
		if (!(*((_BYTE*)v6 + 480) & 1))
			goto LABEL_24;
		if (nox_common_gameFlags_check_40A5C0(32)) {
			if (v2) {
				if (a2) {
					v7 = *(float**)(v2 + 748);
					v12.field_0 = *v7;
					v8 = *((float*)v6 + 3);
					v9 = v7[1];
					v10 = *((float*)v6 + 2);
					v12.field_4 = v9;
					v12.field_8 = v10;
					v12.field_C = v8;
					if (nox_xxx_mapTraceRay_535250(&v12, 0, 0, 9) == 1)
						goto LABEL_24;
				}
			}
		}
		if (!v5)
			break;
		--v5;
	LABEL_24:
		v6 = (LPVOID)nox_xxx_waypointNext_579870((int)v6);
		if (!v6)
			return 0;
	}
	*a1 = *((_DWORD*)v6 + 2);
	a1[1] = *((_DWORD*)v6 + 3);
	return 1;
}

//----- (0057A160) --------------------------------------------------------
int  nox_xxx_playerCanTalkMB_57A160(int a1) {
	int result; // eax

	if (a1 && nox_common_gameFlags_check_40A5C0(2))
		result = (*(_DWORD*)(a1 + 3680) >> 3) & 1;
	else
		result = 0;
	return result;
}

//----- (0057A190) --------------------------------------------------------
int  nox_xxx_giant_57A190(int a1) {
	int result; // eax

	if (a1)
		result = (*(_DWORD*)(a1 + 3680) >> 2) & 1;
	else
		result = 0;
	return result;
}

//----- (0057A1B0) --------------------------------------------------------
char*  sub_57A1B0(__int16 a1) {
	int v1;              // ecx
	unsigned __int8* v2; // eax

	v1 = 0;
	v2 = getMemAt(0x587000, 312212);
	while ((a1 & 0x17F0) != *(_DWORD*)v2) {
		v2 += 8;
		++v1;
		if ((int)v2 >= (int)getMemAt(0x587000, 312268))
			return 0;
	}
	return *(char**)getMemAt(0x587000, 8 * v1 + 312208);
}

//----- (0057A1E0) --------------------------------------------------------
char  sub_57A1E0(int* a1, const char* a2, int* a3, char a4, __int16 a5) {
	int* v5;             // esi
	int v6;              // ebx
	unsigned __int8 v7;  // dl
	char* v8;            // edi
	int v9;              // ecx
	int v10;             // eax
	char* v11;           // edi
	unsigned __int8 v12; // cl
	char result;         // al
	char v14[256];       // [esp+14h] [ebp-200h]
	char v15[256];       // [esp+114h] [ebp-100h]

	v5 = a3;
	v6 = a5 & 0x17F0;
	if (a3)
		sub_57ADF0(a3);
	a1[6] = -1;
	a1[7] = -1;
	a1[8] = -1;
	a1[9] = -1;
	a1[10] = -1;
	a1[11] = -1;
	a1[12] = -1;
	if (a4 & 3) {
		v15[0] = 0;
		strcpy(v14, "maps\\");
		strncat(v14, (const char*)a1, 8u);
		*(_WORD*)&v14[strlen(v14)] = *getMemU16Ptr(0x587000, 312376);
		if (a4 & 2) {
			strcpy(v15, v14);
			if (a2) {
				strcat(v15, a2);
			} else {
				v7 = getMemByte(0x587000, 312388);
				v8 = &v15[strlen(v15) + 1];
				v9 = *getMemU32Ptr(0x587000, 312384);
				*(_DWORD*)--v8 = *getMemU32Ptr(0x587000, 312380);
				*((_DWORD*)v8 + 1) = v9;
				v8[8] = v7;
			}
			v10 = sub_57A3F0(v15, (int)a1, (int)a3, v6);
			v5 = a3;
		} else {
			v10 = 0;
		}
		if (a4 & 1 && !v10) {
			strncat(v14, (const char*)a1, 8u);
			v11 = &v14[strlen(v14) + 1];
			v12 = getMemByte(0x587000, 312396);
			*(_DWORD*)--v11 = *getMemU32Ptr(0x587000, 312392);
			v11[4] = v12;
			sub_57A3F0(v14, (int)a1, (int)v5, v6);
		}
	}
	if (dword_5d4594_2650652 && a4 & 4)
		sub_57A3F0((char*)getMemAt(0x587000, 312400), (int)a1, (int)v5, v6);
	result = a5;
	if (a5 & 0x40)
		result = sub_453FA0((int)(a1 + 6), 132, 0);
	return result;
}
// 57A1E0: using guessed type char var_100[256];

//----- (0057A3F0) --------------------------------------------------------
int  sub_57A3F0(char* a1, int a2, int a3, int a4) {
	FILE* v4;        // eax
	FILE* v5;        // esi
	char* v6;        // eax
	char v8[256];    // [esp+4h] [ebp-300h]
	wchar_t v9[256]; // [esp+104h] [ebp-200h]

	dword_5d4594_2523764 = 6128;
	v4 = nox_fs_open_text(a1);
	v5 = v4;
	if (!v4)
		return 0;
	if (!nox_fs_feof(v4)) {
		do {
			memset(v8, 0, sizeof(v8));
			nox_fs_fgets(v5, v8, 256);
			v6 = strchr(v8, 10);
			if (v6)
				*v6 = 0;
			if (v8[0]) {
				nox_swprintf(v9, L"%S", v8);
				sub_57A4D0(v9, a2, a3, a4);
			}
		} while (!nox_fs_feof(v5));
	}
	nox_fs_close(v5);
	return 1;
}

//----- (0057A4D0) --------------------------------------------------------
void sub_57A4D0(wchar_t* a1, int a2, int a3, int a4) {
	unsigned __int8 v4;    // bl
	int v5;                // edi
	const wchar_t* result; // eax
	const wchar_t* v7;     // esi
	int v8;                // ecx
	const wchar_t* v9;     // eax
	wchar_t* v10;          // esi
	unsigned __int8 v11;   // [esp+10h] [ebp-284h]
	int v12[32];           // [esp+14h] [ebp-280h]
	wchar_t v13[256];      // [esp+94h] [ebp-200h]

	v4 = 0;
	v5 = 0;
	v11 = 0;
	sub_416580();
	result = nox_wcscpy(v13, a1);
	if (true) { // TODO: was if (v13)
		if (v13[0] == 34) {
			result = nox_wcstok(&v13[1], L"\"\n\r");
			v7 = result;
			v5 = 1;
		} else {
			result = nox_wcstok(v13, L" \n\t\r");
			v7 = result;
		}
		if (v7) {
			do {
				v8 = v11;
				v11 = ++v4;
				v12[v8] = (int)v7;
				v9 = &v7[nox_wcslen(v7) + 1];
				if (v5)
					++v9;
				if (*v9 == 34) {
					result = nox_wcstok((wchar_t*)v9 + 1, L"\"\n\r");
					v7 = result;
					v5 = 1;
				} else {
					result = nox_wcstok(0, L" \n\t\r");
					v7 = result;
					v5 = 0;
				}
			} while (v7);
			if (v4) {
				result = (const wchar_t*)sub_57A620(v4, (const wchar_t**)v12, a2, a4);
				if (!result) {
					if (a3) {
						v10 = (wchar_t*)malloc(0x20Cu);
						nox_wcscpy(v10 + 6, a1);
						nox_common_list_append_4258E0(a3, v10);
					}
				}
			}
		}
	}
}
// 57A4D0: using guessed type int var_280[32];

//----- (0057A620) --------------------------------------------------------
BOOL  sub_57A620(unsigned __int8 a1, const wchar_t** a2, int a3, int a4) {
	const wchar_t** v4;    // ebp
	unsigned __int8 v5;    // dl
	int v6;                // esi
	unsigned __int16* v8;  // eax
	int v9;                // eax
	int v10;               // ebx
	unsigned __int8 v11;   // cl
	int v12;               // eax
	unsigned __int16* v13; // eax
	int v14;               // esi
	int v15;               // ecx
	char v16[100];         // [esp+10h] [ebp-64h]
	unsigned __int8 v17;   // [esp+7Ch] [ebp+8h]
	unsigned __int8 v18;   // [esp+7Ch] [ebp+8h]

	v4 = a2;
	nox_sprintf(v16, "%S", *a2);
	v5 = 0;
	v17 = 0;
	do {
		if (!strcmp(*(const char**)getMemAt(0x587000, 8 * v17 + 312208), v16)) {
			dword_5d4594_2523764 = *getMemU32Ptr(0x587000, 8 * v17 + 312212);
			return a4 == dword_5d4594_2523764;
		}
		v17 = ++v5;
	} while (v5 < 7u);
	if (!(dword_5d4594_2523764 & a4) || _nox_wcsicmp(*v4, L"set") || a1 <= 1u)
		return 0;
	if (_nox_wcsicmp(v4[1], L"spell")) {
		if (_nox_wcsicmp(v4[1], L"weapon")) {
			if (_nox_wcsicmp(v4[1], L"armor"))
				return 0;
			if (a1 != 4)
				return 0;
			if (!nox_common_gameFlags_check_40A5C0(1))
				return 0;
			nox_sprintf(v16, "%S", v4[2]);
			v13 = (unsigned __int16*)sub_415EC0(v16);
			if (!v13)
				return 0;
			v14 = sub_415D10((char*)*v13);
			if (!v14)
				return 0;
			if (_nox_wcsicmp(v4[3], L"off"))
				v15 = v14 | *(_DWORD*)(a3 + 48);
			else
				v15 = ~v14 & *(_DWORD*)(a3 + 48);
			*(_DWORD*)(a3 + 48) = v15;
		} else {
			if (a1 != 4)
				return 0;
			if (!nox_common_gameFlags_check_40A5C0(1))
				return 0;
			nox_sprintf(v16, "%S", v4[2]);
			v8 = (unsigned __int16*)sub_415A30(v16);
			if (!v8)
				return 0;
			v9 = nox_xxx_ammoCheck_415880((char*)*v8);
			v10 = v9;
			if (!v9)
				return 0;
			v11 = 0;
			v18 = 0;
			if (v9 > 0) {
				do {
					v12 = v10 >> 8;
					if (v10 >> 8 > 0)
						v10 >>= 8;
					++v11;
				} while (v12 > 0);
				v18 = v11;
			}
			if (_nox_wcsicmp(v4[3], L"off"))
				*(_BYTE*)(v18 + a3 + 43) |= v10;
			else
				*(_BYTE*)(v18 + a3 + 43) &= ~(_BYTE)v10;
		}
	} else {
		if (a1 != 4)
			return 0;
		nox_sprintf(v16, "%S", v4[2]);
		v6 = nox_xxx_spellNameToN_4243F0(v16);
		if (!v6) {
			v6 = sub_424960((wchar_t*)v4[2]);
			if (!v6)
				return 0;
		}
		if (!nox_xxx_spellGetValidMB_424B50(v6))
			return 0;
		if (nox_xxx_spellGetFlags_424A70(v6) & 0x7000000) {
			if (!_nox_wcsicmp(v4[3], L"off"))
				sub_453FA0(a3 + 24, v6, 0);
		}
	}
	if (a4 == dword_5d4594_2523764)
		return 1;
	return 0;
}

//----- (0057A950) --------------------------------------------------------
int  sub_57A950(char* a1) {
	char v2[256]; // [esp+Ch] [ebp-100h]

	strcpy(v2, "maps\\");
	strncat(v2, a1, (strlen(a1) - 4 < 256 ? strlen(a1) - 4 : 256));
	*(_WORD*)&v2[strlen(v2)] = *getMemU16Ptr(0x587000, 312564);
	strcat(v2, a1);
	return sub_4D0550(v2);
}

//----- (0057A9F0) --------------------------------------------------------
int  sub_57A9F0(const char* a1, const char* a2) {
	CHAR FileName[256]; // [esp+Ch] [ebp-100h]

	strcpy(FileName, "maps\\");
	strcat(FileName, a1);
	*(_WORD*)&FileName[strlen(FileName)] = *getMemU16Ptr(0x587000, 312576);
	strcat(FileName, a2);
	return remove(FileName);
}

//----- (0057AAA0) --------------------------------------------------------
char  sub_57AAA0(const char* a1, char* a2, int* a3) {
	FILE* v3;           // eax
	FILE* v4;           // ebp
	int* i;             // esi
	char* v6;           // eax
	int v7;             // esi
	int v8;             // edi
	char* v9;           // eax
	int v10;            // esi
	int v11;            // edi
	char* v12;          // eax
	int v13;            // ebx
	int v14;            // esi
	char* v15;          // edi
	char* v16;          // eax
	int v18;            // [esp+10h] [ebp-27Ch]
	char v19[24];       // [esp+14h] [ebp-278h]
	char v20[36];       // [esp+2Ch] [ebp-260h]
	char v21[24];       // [esp+50h] [ebp-23Ch]
	char v22[36];       // [esp+68h] [ebp-224h]
	CHAR v23[256];      // [esp+8Ch] [ebp-200h]
	CHAR FileName[256]; // [esp+18Ch] [ebp-100h]

	LOBYTE(v3) = a2[52];
	if ((char)v3 >= 0) {
		strcpy(FileName, "maps\\");
		strncat(FileName, a2, 8u);
		*(_WORD*)&FileName[strlen(FileName)] = *getMemU16Ptr(0x587000, 312588);
		strcat(FileName, a1);
		_chmod(FileName, 128);
		v3 = nox_fs_create_text(FileName);
		v4 = v3;
		if (v3) {
			if (dword_5d4594_2650652) {
				strcpy(v21, a2);
				strcpy(v19, a2);
				sub_57A1E0((int*)v21, 0, 0, 4, *((_WORD*)a2 + 26));
				sub_57A1E0((int*)v19, 0, 0, 3, *((_WORD*)a2 + 26));
			}
			if (a3) {
				for (i = nox_common_list_getFirstSafe_425890(a3); i; i = nox_common_list_getNextSafe_4258A0(i)) {
					nox_sprintf(v23, "%S\n", i + 3);
					nox_fs_fputs(v4, v23);
				}
			}
			v6 = sub_57A1B0(*((_WORD*)a2 + 26));
			nox_fs_fputs(v4, v6);
			nox_fs_fputs(v4, "\n");
			v7 = 1;
			v8 = 136;
			do {
				if (nox_xxx_spellGetValidMB_424B50(v7) && !sub_454000((int)(a2 + 24), v7) && nox_xxx_spellGetFlags_424A70(v7) & 0x7000000 &&
					(!dword_5d4594_2650652 || sub_454000((int)v22, v7) || !sub_454000((int)v20, v7))) {
					v9 = nox_xxx_spellNameByN_424870(v7);
					nox_sprintf(v23, "%s %s \"%s\" %s\n", getMemAt(0x587000, 312616), getMemAt(0x587000, 312608), v9,
								getMemAt(0x587000, 312604));
					nox_fs_fputs(v4, v23);
				}
				++v7;
				--v8;
			} while (v8);
			v10 = 1;
			v11 = 26;
			do {
				if (!(v10 & *((_DWORD*)a2 + 12))) {
					v12 = sub_415E40((char*)v10);
					if (v12) {
						nox_sprintf(v23, "%s %s \"%s\" %s\n", getMemAt(0x587000, 312648), getMemAt(0x587000, 312640), v12,
									getMemAt(0x587000, 312636));
						nox_fs_fputs(v4, v23);
					}
				}
				v10 *= 2;
				--v11;
			} while (v11);
			v13 = 1;
			v14 = 1;
			v18 = 27;
			v15 = a2 + 44;
			do {
				if (!((unsigned __int8)v13 & (unsigned __int8)*v15)) {
					v16 = sub_4159B0((char*)v14);
					if (v16) {
						nox_sprintf(v23, "%s %s \"%s\" %s\n", getMemAt(0x587000, 312680), getMemAt(0x587000, 312672), v16,
									getMemAt(0x587000, 312668));
						nox_fs_fputs(v4, v23);
					}
				}
				if (v13 == 128) {
					v13 = 1;
					++v15;
				} else {
					v13 *= 2;
				}
				v14 *= 2;
				--v18;
			} while (v18);
			LOBYTE(v3) = 0;
			nox_fs_close(v4);
		}
	}
	return (char)v3;
}

//----- (0057ADF0) --------------------------------------------------------
int*  sub_57ADF0(int* a1) {
	int* result; // eax
	int* v2;     // esi
	int* v3;     // edi

	result = nox_common_list_getFirstSafe_425890(a1);
	v2 = result;
	if (result) {
		do {
			v3 = nox_common_list_getNextSafe_4258A0(v2);
			sub_425920((_DWORD**)v2);
			free(v2);
			v2 = v3;
		} while (v3);
	}
	return result;
}

//----- (0057AE30) --------------------------------------------------------
int  sub_57AE30(const char* a1) {
	int v1;          // ebp
	const char** v2; // edi

	v1 = 0;
	v2 = (const char**)getMemAt(0x587000, 312208);
	while (strcmp(*v2, a1)) {
		v2 += 2;
		++v1;
		if ((int)v2 >= (int)getMemAt(0x587000, 312264))
			return 0;
	}
	return *getMemU32Ptr(0x587000, 8 * v1 + 312212);
}

//----- (0057AEA0) --------------------------------------------------------
int  nox_xxx_playerCheckSpellClass_57AEA0(int a1, int a2) {
	int v2;     // eax
	int result; // eax
	int v4;     // ecx

	v2 = nox_xxx_spellGetFlags_424A70(a2);
	if (a1 == 1) {
		v4 = 0x2000000;
	} else {
		if (a1 != 2)
			return 9;
		v4 = 0x4000000;
	}
	if (v2 & 0x1000000 || v4 & v2)
		result = 0;
	else
		result = 9;
	return result;
}

//----- (0057AEE0) --------------------------------------------------------
BOOL  sub_57AEE0(int a1, int a2) { return a1 < 75 || a1 > 114 || nox_xxx_countControlledCreatures_500D10(a2) <= 4; }

//----- (0057AF20) --------------------------------------------------------
int nox_xxx_get_57AF20() { return dword_5d4594_2523804; }

//----- (0057B0A0) --------------------------------------------------------
void sub_57B0A0() {
	int result; // eax
	_DWORD* v1; // ecx

	result = dword_5d4594_2523804;
	if (!result) {
		return;
	}
	v1 = *(_DWORD**)&dword_5d4594_2523780;
	if (dword_5d4594_2523780 && (!*getMemU32Ptr(0x5D4594, 2523772) || *getMemU32Ptr(0x5D4594, 2523772) == 1)) {
		nox_xxx_netSendPointFx_522FF0(154, (float2*)(dword_5d4594_2523780 + 56));
		v1 = *(_DWORD**)&dword_5d4594_2523780;
	}
	if (dword_5d4594_2523776) {
		nox_xxx_delayedDeleteObject_4E5CC0(*(int*)&dword_5d4594_2523776);
		v1 = *(_DWORD**)&dword_5d4594_2523780;
	}
	dword_5d4594_2523776 = 0;
	if (v1)
		nox_xxx_playerSetState_4FA020(v1, 13);
	dword_5d4594_2523780 = 0;
	if (!sub_45D9B0())
		sub_413A00(0);
	dword_5d4594_2523804 = 0;
}

//----- (0057B180) --------------------------------------------------------
__int64 nox_xxx___Getcvt_57B180() { return *getMemU64Ptr(0x5D4594, 2523788); }

//----- (0057B190) --------------------------------------------------------
int  sub_57B190(unsigned __int16 a1, unsigned __int16 a2) {
	int result; // eax
	double v3;  // st7
	double v4;  // st6

	if (!a2)
		return 4;
	if (a1 == a2)
		return 0;
	v3 = (double)a1;
	v4 = (double)a2;
	if (v3 >= v4 * *(double*)&qword_581450_9544)
		return 1;
	result = 2;
	if (v3 < v4 * *getMemDoublePtr(0x581450, 9608))
		result = 3;
	return result;
}

//----- (0057B200) --------------------------------------------------------
void nox_xxx_loadBaseValues_57B200() {
	*getMemFloatPtr(0x5D4594, 2523812) = nox_xxx_gamedataGetFloat_419D40(getMemAt(0x587000, 312832));
	*getMemFloatPtr(0x5D4594, 2523816) = nox_xxx_gamedataGetFloat_419D40(getMemAt(0x587000, 312844));
	*getMemFloatPtr(0x5D4594, 2523824) = nox_xxx_gamedataGetFloat_419D40(getMemAt(0x587000, 312856));
	*getMemFloatPtr(0x5D4594, 2523820) = nox_xxx_gamedataGetFloat_419D40(getMemAt(0x587000, 312872));
	*getMemFloatPtr(0x5D4594, 2523828) = nox_xxx_gamedataGetFloat_419D40(getMemAt(0x587000, 312884)) * *(float*)&nox_xxx_warriorMaxHealth_587000_312784;
	*getMemFloatPtr(0x5D4594, 2523832) = nox_xxx_gamedataGetFloat_419D40(getMemAt(0x587000, 312904)) * *(float*)&nox_xxx_warriorMaxMana_587000_312788;
	*getMemFloatPtr(0x5D4594, 2523840) = nox_xxx_gamedataGetFloat_419D40(getMemAt(0x587000, 312920)) * *(float*)&nox_xxx_warriorMaxStrength_587000_312792;
	*getMemFloatPtr(0x5D4594, 2523836) = nox_xxx_gamedataGetFloat_419D40(getMemAt(0x587000, 312940)) * *(float*)&nox_xxx_warriorMaxSpeed_587000_312796;
	*getMemFloatPtr(0x5D4594, 2523860) = nox_xxx_gamedataGetFloat_419D40(getMemAt(0x587000, 312956)) * *(float*)&nox_xxx_conjurerMaxHealth_587000_312800;
	*getMemFloatPtr(0x5D4594, 2523864) = nox_xxx_gamedataGetFloat_419D40(getMemAt(0x587000, 312976)) * *(float*)&nox_xxx_conjurerMaxMana_587000_312804;
	*getMemFloatPtr(0x5D4594, 2523872) = nox_xxx_gamedataGetFloat_419D40(getMemAt(0x587000, 312992)) * *(float*)&nox_xxx_conjurerStrength_587000_312808;
	*getMemFloatPtr(0x5D4594, 2523868) = nox_xxx_gamedataGetFloat_419D40(getMemAt(0x587000, 313012)) * *(float*)&nox_xxx_conjurerSpeed_587000_312812;
	*getMemFloatPtr(0x5D4594, 2523844) = nox_xxx_gamedataGetFloat_419D40(getMemAt(0x587000, 313032)) * *(float*)&nox_xxx_wizardMaxHealth_587000_312816;
	*getMemFloatPtr(0x5D4594, 2523848) = nox_xxx_gamedataGetFloat_419D40(getMemAt(0x587000, 313048)) * *(float*)&nox_xxx_wizardMaximumMana_587000_312820;
	*getMemFloatPtr(0x5D4594, 2523856) = nox_xxx_gamedataGetFloat_419D40(getMemAt(0x587000, 313064)) * *(float*)&nox_xxx_wizardStrength_587000_312824;
	*getMemFloatPtr(0x5D4594, 2523852) = nox_xxx_gamedataGetFloat_419D40(getMemAt(0x587000, 313084)) * *(float*)&nox_xxx_wizardSpeed_587000_312828;
}

//----- (0057B350) --------------------------------------------------------
float* sub_57B350() { return getMemFloatPtr(0x5D4594, 2523812); }

//----- (0057B360) --------------------------------------------------------
float*  nox_xxx_plrGetMaxVarsPtr_57B360(int a1) { return getMemFloatPtr(0x5D4594, 16 * a1 + 2523828); }

//----- (0057B370) --------------------------------------------------------
char  sub_57B370(int a1, unsigned __int8 a2, int a3) {
	_DWORD* v3;  // eax
	char result; // al

	if (!(a1 & 0x3001010))
		return -1;
	if (a1 & 0x1001000) {
		v3 = nox_xxx_getProjectileClassById_413250(a3);
		goto LABEL_4;
	}
	if (a1 & 0x2000000) {
		v3 = nox_xxx_equipClothFindDefByTT_413270(a3);
	LABEL_4:
		if (v3)
			result = *((_BYTE*)v3 + 62);
		else
			result = 0;
		return result;
	}
	if (a1 & 0x10)
		result = ~(a2 >> 5) | 0xFE;
	else
		result = a3;
	return result;
}

//----- (0057B3D0) --------------------------------------------------------
BOOL  nox_xxx_playerClassCanUseItem_57B3D0(int a1, char a2) {
	return ((unsigned __int8)(1 << a2) &
			(unsigned __int8)sub_57B370(*(_DWORD*)(a1 + 8), *(_DWORD*)(a1 + 12), *(unsigned __int16*)(a1 + 4))) != 0;
}

//----- (0057B400) --------------------------------------------------------
int  nox_xxx_client_57B400(int a1) {
	int v1; // eax

	v1 = *getMemU32Ptr(0x5D4594, 2523876);
	if (!*getMemU32Ptr(0x5D4594, 2523876)) {
		v1 = nox_xxx_getTTByNameSpriteMB_44CFC0("Glyph");
		*getMemU32Ptr(0x5D4594, 2523876) = v1;
	}
	if (!*getMemU32Ptr(0x5D4594, 2618908))
		return 0;
	if (*(_DWORD*)(a1 + 108) != v1 || *(_BYTE*)(*getMemU32Ptr(0x5D4594, 2618908) + 2251) == 1)
		return 1;
	return 0;
}

//----- (0057B450) --------------------------------------------------------
BOOL  sub_57B450(int* a1) {
	int v1;             // eax
	unsigned __int8 v2; // bl

	v1 = *getMemU32Ptr(0x5D4594, 2523880);
	if (!*getMemU32Ptr(0x5D4594, 2523880)) {
		v1 = nox_xxx_getTTByNameSpriteMB_44CFC0("Glyph");
		*getMemU32Ptr(0x5D4594, 2523880) = v1;
	}
	if (!a1 || !*getMemU32Ptr(0x5D4594, 2614252) || !*getMemU32Ptr(0x5D4594, 2618908) ||
		a1[27] == v1 && *(_BYTE*)(*getMemU32Ptr(0x5D4594, 2618908) + 2251) != 1) {
		return 0;
	}
	v2 = 1 << *(_BYTE*)(*getMemU32Ptr(0x5D4594, 2618908) + 2251);
	return (v2 & (unsigned __int8)sub_57B370(a1[28], a1[29], a1[27])) != 0;
}

//----- (0057B4D0) --------------------------------------------------------
int  sub_57B4D0(int a1) {
	int result; // eax

	result = a1;
	if (a1) {
		*getMemU32Ptr(0x5D4594, 2523884) = a1;
		dword_5d4594_2523888 = 1;
	} else {
		dword_5d4594_2523888 = 0;
	}
	return result;
}

//----- (0057B500) --------------------------------------------------------
char  sub_57B500(int a1, int a2, char a3) {
	int v3;      // eax
	char v4;     // cl
	int v5;      // esi
	int v6;      // edi
	char result; // al
	int v8;      // eax
	int v9;      // ecx
	bool v10;    // sf
	bool v11;    // cc

	if (a1 < 0)
		return -1;
	if (a1 >= 256)
		return -1;
	if (a2 < 0)
		return -1;
	if (a2 >= 256)
		return -1;
	v3 = nox_xxx_wall_4105E0(a1, a2);
	if (!v3)
		return -1;
	v4 = *(_BYTE*)(v3 + 4);
	if (v4 & 0x10) {
		if (!(a3 & 0x10))
			return -1;
		v5 = *(_DWORD*)(v3 + 28);
		if (!v5)
			return -1;
		v6 = *(_DWORD*)(v5 + 748);
		if (a3 & 8) {
			if (!*(_BYTE*)(v6 + 1))
				return -1;
			if (dword_5d4594_2523888 && nox_xxx_doorGetSomeKey_4E8910(*getMemIntPtr(0x5D4594, 2523884), v5)) {
				dword_5d4594_2523888 = 0;
				return -1;
			}
		}
		if (a3 >= 0 && *(_BYTE*)(v5 + 12) & 4)
			return -1;
		v8 = *(_DWORD*)(v6 + 12);
		if (v8 != *(_DWORD*)(v6 + 4))
			return -1;
		v9 = *getMemIntPtr(0x587000, 8 * v8 + 196184);
		v10 = v9 < 0;
		v11 = v9 <= 0;
		if (v9 > 0) {
			if (*getMemIntPtr(0x587000, 8 * v8 + 196188) > 0)
				return 1;
			v10 = v9 < 0;
			v11 = v9 <= 0;
		}
		if (v10) {
			if (*getMemIntPtr(0x587000, 8 * v8 + 196188) < 0)
				return 1;
			v11 = v9 <= 0;
			if (v9 < 0) {
				if (*getMemIntPtr(0x587000, 8 * v8 + 196188) > 0)
					return 0;
				v11 = v9 <= 0;
			}
		}
		if (v11 || *getMemIntPtr(0x587000, 8 * v8 + 196188) >= 0)
			return -1;
		result = 0;
	} else {
		if (!(a3 & 0x40) && v4 & 0x40 || v4 & 4 && *(_BYTE*)(*(_DWORD*)(v3 + 28) + 22) > 0xBu)
			return -1;
		result = *(_BYTE*)v3;
	}
	return result;
}

//----- (0057B630) --------------------------------------------------------
char  sub_57B630(int a1, int a2, int a3) {
	int v3;   // eax
	char v4;  // cl
	int v5;   // eax
	int v6;   // esi
	int v7;   // ecx
	int v8;   // edx
	bool v9;  // sf
	bool v10; // cc
	char v11; // bl
	int v13;  // ecx

	if (a2 >= 0 && a2 < 256 && (a3 >= 0 || a3 < 256)) {
		v3 = nox_xxx_wall_4105E0(a2, a3);
		if (v3) {
			v4 = *(_BYTE*)(v3 + 4);
			if (v4 & 0x10) {
				v5 = *(_DWORD*)(v3 + 28);
				if (v5) {
					v6 = *(_DWORD*)(v5 + 748);
					v7 = *(_DWORD*)(v6 + 12);
					if (v7 == *(int*)(v6 + 4)) {
						v8 = *getMemIntPtr(0x587000, 8 * v7 + 196184);
						v9 = v8 < 0;
						v10 = v8 <= 0;
						if (v8 > 0) {
							if (*getMemIntPtr(0x587000, 8 * v7 + 196188) > 0) {
								v11 = 1;
								goto LABEL_22;
							}
							v9 = v8 < 0;
							v10 = v8 <= 0;
						}
						if (v9) {
							if (*getMemIntPtr(0x587000, 8 * v7 + 196188) < 0) {
								v11 = 1;
								goto LABEL_22;
							}
							v10 = v8 <= 0;
							if (v8 < 0) {
								if (*getMemIntPtr(0x587000, 8 * v7 + 196188) > 0) {
								LABEL_21:
									v11 = 0;
								LABEL_22:
									if (*(_DWORD*)(v5 + 508)) {
										if (nox_common_randomInt_415FA0(0, 100) >= 50)
											return v11;
									} else if (*(_BYTE*)(v6 + 1) && !nox_xxx_doorGetSomeKey_4E8910(a1, v5)) {
										return v11;
									}
									return -1;
								}
								v10 = v8 <= 0;
							}
						}
						if (!v10 && *getMemIntPtr(0x587000, 8 * v7 + 196188) < 0)
							goto LABEL_21;
					}
				}
			} else if (!(*(_DWORD*)(a1 + 16) & 0x4000) || !(v4 & 0x40)) {
				if (!(v4 & 4))
					return *(_BYTE*)v3;
				v13 = *(_DWORD*)(v3 + 28);
				if (!(*(_BYTE*)(v13 + 20) & 2) && *(_BYTE*)(v13 + 22) <= 0xBu)
					return *(_BYTE*)v3;
			}
		}
	}
	return -1;
}

//----- (0057B770) --------------------------------------------------------
float2*  sub_57B770(float2* a1, float2* a2) {
	float2* result; // eax
	long double v3; // st5
	long double v4; // st7
	double v5;      // st6
	long double v6; // st5
	float v7;       // [esp+0h] [ebp-18h]
	float v8;       // [esp+4h] [ebp-14h]
	float v9;       // [esp+Ch] [ebp-Ch]
	float v10;      // [esp+10h] [ebp-8h]
	float v11;      // [esp+20h] [ebp+8h]

	result = a2;
	v9 = a2->field_0;
	v3 = sqrt(a2->field_0 * a2->field_0 + a2->field_4 * a2->field_4);
	v4 = v3 + 0.1;
	v5 = -a2->field_4;
	v6 = (a1->field_0 * a2->field_0 + a1->field_4 * a2->field_4) / (v3 + 0.1);
	v11 = (a2->field_0 * a1->field_4 + v5 * a1->field_0) / v4;
	v7 = v6 * result->field_0 / v4;
	v8 = v6 * result->field_4 / v4;
	v10 = v11 * v5 / v4;
	a1->field_0 = v10 - v7;
	a1->field_4 = v11 * v9 / v4 - v8;
	return result;
}

//----- (0057B810) --------------------------------------------------------
int  nox_xxx_collideReflect_57B810(float* a1, int a2) {
	int result; // eax
	double v3;  // st7
	int v4;     // ecx

	result = a2;
	v3 = *(float*)a2;
	if (a1[1] * *a1 <= 0.0) {
		v4 = *(_DWORD*)(a2 + 4);
		*(float*)(a2 + 4) = *(float*)a2;
		*(_DWORD*)a2 = v4;
	} else {
		*(float*)a2 = -*(float*)(a2 + 4);
		*(float*)(a2 + 4) = -v3;
	}
	return result;
}

//----- (0057B850) --------------------------------------------------------
BOOL  nox_xxx_map_57B850(float2* a1, float* a2, float2* a3) {
	BOOL result; // eax
	float v4;    // [esp+0h] [ebp-10h]
	float v5;    // [esp+4h] [ebp-Ch]
	float v6;    // [esp+8h] [ebp-8h]
	float v7;    // [esp+Ch] [ebp-4h]

	v4 = a2[5] + a1->field_0;
	v5 = a2[6] + a1->field_4;
	result = 0;
	if ((v5 - v4 + a3->field_0 - a3->field_4) * 0.70709997 < 0.0 &&
		(a2[8] + a1->field_4 - (a2[7] + a1->field_0) + a3->field_0 - a3->field_4) * 0.70709997 > 0.0) {
		v6 = a2[9] + a1->field_0;
		v7 = a2[10] + a1->field_4;
		if ((v7 + v6 - a3->field_0 - a3->field_4) * 0.70709997 > 0.0 &&
			(v5 + v4 - a3->field_0 - a3->field_4) * 0.70709997 < 0.0) {
			result = 1;
		}
	}
	return result;
}

//----- (0057B920) --------------------------------------------------------
int  sub_57B920(void* a1) {
	int result; // eax

	result = 0;
	memset(a1, 0, 0x7F8u);
	return result;
}

//----- (0057B930) --------------------------------------------------------
char  sub_57B930(int a1, int a2, int a3, unsigned int a4) {
	int v4; // eax
	int v5; // edx

	v4 = (unsigned __int8)a2;
	v5 = (unsigned __int8)a2;
	if ((unsigned __int8)a2 == 255 || !(_BYTE)a2) {
		v4 = 1;
		v5 = 1;
	}
	while (*(unsigned __int16*)(a1 + 8 * v4) != a2 || *(unsigned __int16*)(a1 + 8 * v4 + 2) != a3) {
		if (++v4 == 255)
			v4 = 1;
		if (v4 == v5)
			goto LABEL_11;
	}
	if (*(_DWORD*)(a1 + 8 * v4 + 4) >= a4)
		return v4;
LABEL_11:
	LOBYTE(v4) = -1;
	return v4;
}

//----- (0057B9A0) --------------------------------------------------------
char  nox_xxx_cliGenerateAlias_57B9A0(int a1, int a2, int a3, unsigned int a4) {
	int v4; // eax
	int v5; // edx

	v4 = (unsigned __int8)a2;
	v5 = (unsigned __int8)a2;
	if ((unsigned __int8)a2 == 255 || !(_BYTE)a2) {
		v4 = 1;
		v5 = 1;
	}
	while ((*(unsigned __int16*)(a1 + 8 * v4) != a2 || *(unsigned __int16*)(a1 + 8 * v4 + 2) != a3) &&
		   *(_DWORD*)(a1 + 8 * v4 + 4) >= a4) {
		if (++v4 == 255)
			v4 = 1;
		if (v4 == v5) {
			LOBYTE(v4) = -1;
			return v4;
		}
	}
	return v4;
}

//----- (0057BA10) --------------------------------------------------------
int  sub_57BA10(int a1, __int16 a2, __int16 a3, int a4) {
	int result; // eax

	result = a1;
	*(_WORD*)a1 = a2;
	*(_WORD*)(a1 + 2) = a3;
	*(_DWORD*)(a1 + 4) = a4;
	return result;
}

//----- (0057BA30) --------------------------------------------------------
int  sub_57BA30(int2* a1, int2* a2, int4* a3) {
	int v3;     // eax
	int v4;     // edx
	int v5;     // ebx
	int v6;     // edx
	int v7;     // ebp
	int v8;     // ebx
	int result; // eax
	int v10;    // edx
	int v11;    // edi
	BOOL v12;   // [esp+10h] [ebp-4h]

	v12 = 0;
	while (1) {
		if (a1->field_0 >= a3->field_0)
			v3 = a1->field_0 <= a3->field_8 ? 0 : 4;
		else
			v3 = 8;
		v4 = a1->field_4;
		v5 = a3->field_4;
		if (v4 >= v5) {
			if (v4 > a3->field_C)
				LOBYTE(v3) = v3 | 1;
		} else {
			LOBYTE(v3) = v3 | 2;
		}
		if (a2->field_0 >= a3->field_0)
			v6 = a2->field_0 <= a3->field_8 ? 0 : 4;
		else
			v6 = 8;
		v7 = a2->field_4;
		if (v7 >= v5) {
			if (v7 > a3->field_C)
				v6 |= 1u;
		} else {
			v6 |= 2u;
		}
		if (v6 & v3) {
			result = 0;
			goto LABEL_30;
		}
		if (!v3)
			break;
	LABEL_20:
		if (v3 & 1) {
			a1->field_0 += (a3->field_C - a1->field_4) * (a2->field_0 - a1->field_0) / (a2->field_4 - a1->field_4);
			a1->field_4 = a3->field_C;
		} else if (v3 & 2) {
			a1->field_0 += (a3->field_4 - a1->field_4) * (a2->field_0 - a1->field_0) / (a2->field_4 - a1->field_4);
			a1->field_4 = a3->field_4;
		} else if (v3 & 4) {
			a1->field_4 += (a3->field_8 - a1->field_0) * (a2->field_4 - a1->field_4) / (a2->field_0 - a1->field_0);
			a1->field_0 = a3->field_8;
		} else if (v3 & 8) {
			a1->field_4 += (a2->field_4 - a1->field_4) * (a3->field_0 - a1->field_0) / (a2->field_0 - a1->field_0);
			a1->field_0 = a3->field_0;
		}
	}
	if (v6) {
		v3 = a1->field_0;
		v8 = a1->field_4;
		*a1 = *a2;
		a2->field_0 = v3;
		a2->field_4 = v8;
		LOBYTE(v3) = v6;
		v12 = !v12;
		goto LABEL_20;
	}
	result = 1;
LABEL_30:
	if (v12) {
		v10 = a1->field_0;
		v11 = a1->field_4;
		*a1 = *a2;
		a2->field_0 = v10;
		a2->field_4 = v11;
	}
	return result;
}

//----- (0057BBC0) --------------------------------------------------------
#ifndef NOX_CGO
int  nox_xxx_getToken_57BBC0(FILE* f, char* buf, int bufSz) {
	bool tab = true;
	for (int i = 0; i < bufSz; i++) {
		char c = nox_fs_fgetc(f);
		if (!isspace(c)) {
			if (c == ';') {
				buf[i] = 0;
				return c;
			}
			tab = false;
			buf[i] = c;
		} else {
			if (!tab) {
				buf[i] = ' ';
			} else {
				i--; // stay
			}
		}
	}
	buf[bufSz - 1] = 0;
	return bufSz;
}
#endif // NOX_CGO

//----- (0057BC50) --------------------------------------------------------
int  nox_xxx_mapNxzDecompress_57BC50(char* a1, char* a2) {
	FILE* v2;   // eax
	FILE* v3;   // edi
	int v4;     // esi
	char* v5;   // esi
	char* v6;   // ebx
	char* v7;   // eax
	char* v8;   // ebp
	_DWORD* v9; // edi
	FILE* v10;  // esi
	size_t v12; // [esp+10h] [ebp-8h]
	size_t v13; // [esp+14h] [ebp-4h]

	v12 = 0;
	if (!a1)
		return 0;
	if (!a2)
		return 0;
	v2 = nox_fs_open(a1);
	v3 = v2;
	if (!v2)
		return 0;
	v4 = nox_fs_fsize(v3);
	v5 = (char*)(v4 - 4);
	nox_binfile_fread2_40ADD0((char*)&v12, 1u, 4u, v3);
	v6 = (char*)malloc((size_t)v5);
	v7 = (char*)malloc(v12);
	v8 = v7;
	if (!v6 || !v7)
		return 0;
	nox_binfile_fread2_40ADD0(v6, 1u, (size_t)v5, v3);
	nox_fs_close(v3);
	a1 = v5;
	v13 = v12;
	v9 = sub_578BF0();
	if (nox_xxx_nxzDecompress_578C10(v9, &v8[v12 - v13], &v13, (unsigned int)&v5[v6 - a1], &a1)) {
		while ((int)a1 > 0 && nox_xxx_nxzDecompress_578C10(v9, &v8[v12 - v13], &v13, (unsigned int)&v5[v6 - a1], &a1))
			;
	}
	sub_578C40(v9);
	v10 = nox_fs_create(a2);
	if (!v10)
		return 0;
	nox_fs_fwrite(v10, v8, v12);
	nox_fs_close(v10);
	free(v6);
	free(v8);
	return 1;
}

//----- (0057BDD0) --------------------------------------------------------
int  nox_xxx_mapFile_57BDD0(LPVOID lpMem, int a2) {
	size_t v2;        // ebx
	FILE* v3;         // eax
	FILE* v4;         // esi
	char* v5;         // edi
	unsigned int v6;  // eax
	void* v7;         // eax
	void** v8;        // ebp
	size_t v9;        // eax
	size_t i;         // esi
	unsigned int v11; // eax
	FILE* v12;        // eax
	FILE* v13;        // esi
	size_t v15;       // [esp+Ch] [ebp-4h]
	void* lpMema;     // [esp+14h] [ebp+4h]

	v2 = 0;
	v15 = 0;
	if (!lpMem)
		return 0;
	if (!a2)
		return 0;
	v3 = nox_fs_open(lpMem);
	v4 = v3;
	if (!v3)
		return 0;
	v15 = nox_fs_fsize(v4);
	v5 = (char*)malloc(v15);
	v6 = sub_578BA0(v15);
	v7 = malloc(v6);
	lpMema = v7;
	if (!v5 || !v7)
		return 0;
	nox_binfile_fread2_40ADD0(v5, 1u, v15, v4);
	nox_fs_close(v4);
	v8 = (void**)sub_578B80();
	v9 = v15;
	for (i = 0; i < v15; i += 500000) {
		v11 = v9 - i;
		if (v11 > 0x7A120)
			v11 = 500000;
		v2 += sub_578BB0(v8, (int)lpMema + v2, (unsigned __int8*)&v5[i], v11);
		v9 = v15;
	}
	sub_578BD0(v8);
	v12 = nox_fs_create(a2);
	v13 = v12;
	if (!v12)
		return 0;
	nox_fs_fwrite(v12, &v15, 4);
	nox_fs_fwrite(v13, lpMema, v2);
	nox_fs_close(v13);
	free(v5);
	free(lpMema);
	return 1;
}

//----- (0057BF20) --------------------------------------------------------
BOOL  sub_57BF20(int a1, int a2) { return a1 && a2; }

//----- (0057BF40) --------------------------------------------------------
BOOL  sub_57BF40(int a1, int a2) { return a1 && a2; }

//----- (0057BF60) --------------------------------------------------------
int sub_57BF60() { return 0; }

//----- (0057BF70) --------------------------------------------------------
int sub_57BF70() { return dword_5d4594_2523904; }

//----- (0057BF80) --------------------------------------------------------
unsigned int sub_57BF80() {
	unsigned int v0; // esi
	int i;           // eax
	unsigned int v2; // ecx

	v0 = 0;
	for (i = nox_server_getFirstMapGroup_57C080(); i; i = nox_server_getNextMapGroup_57C090(i)) {
		v2 = *(_DWORD*)(i + 4);
		if (v2 >= v0)
			v0 = v2 + 1;
	}
	return v0;
}

//----- (0057BFB0) --------------------------------------------------------
int nox_xxx_allocGroupRelatedArrays_57BFB0() {
	dword_5d4594_2523904 = 0;
	char* result = nox_new_alloc_class("ItemGroupInfo", 96, 512);
	nox_alloc_groupInfo_2523892 = result;
	if (!result) {
		return 0;
	}
	nox_alloc_itemGroupElem_2523896 = nox_new_alloc_class("ItemGroupElement", 16, 5000);
	return nox_alloc_itemGroupElem_2523896 != 0;
}

//----- (0057C000) --------------------------------------------------------
void sub_57C000() {
	dword_5d4594_2523904 = 0;
	nox_alloc_class_yyy_4144D0(*(_DWORD**)&nox_alloc_itemGroupElem_2523896);
	nox_alloc_class_yyy_4144D0(*(_DWORD**)&nox_alloc_groupInfo_2523892);
	dword_5d4594_2523900 = 0;
}

//----- (0057C030) --------------------------------------------------------
int sub_57C030() {
	dword_5d4594_2523904 = 0;
	if (nox_alloc_groupInfo_2523892) {
		nox_free_alloc_class(*(LPVOID*)&nox_alloc_groupInfo_2523892);
		nox_alloc_groupInfo_2523892 = 0;
	}
	if (nox_alloc_itemGroupElem_2523896) {
		nox_free_alloc_class(*(LPVOID*)&nox_alloc_itemGroupElem_2523896);
		nox_alloc_itemGroupElem_2523896 = 0;
	}
	dword_5d4594_2523900 = 0;
	return 1;
}

//----- (0057C080) --------------------------------------------------------
int nox_server_getFirstMapGroup_57C080() { return dword_5d4594_2523900; }

//----- (0057C090) --------------------------------------------------------
int  nox_server_getNextMapGroup_57C090(int a1) {
	int result; // eax

	if (a1)
		result = *(_DWORD*)(a1 + 88);
	else
		result = 0;
	return result;
}

//----- (0057C0A0) --------------------------------------------------------
int  nox_server_scriptGetGroup_57C0A0(int a1) {
	int result; // eax

	result = dword_5d4594_2523900;
	if (!dword_5d4594_2523900)
		return 0;
	while (*(_DWORD*)(result + 4) != a1) {
		result = *(_DWORD*)(result + 88);
		if (!result)
			return 0;
	}
	return result;
}

//----- (0057C0C0) --------------------------------------------------------
int  nox_server_mapLoadAddGroup_57C0C0(char* a1, int a2, char a3) {
	int result; // eax
	int v4;     // esi
	int v5;     // ecx

	result = (int)sub_57C330();
	v4 = result;
	if (result) {
		*(_DWORD*)(result + 4) = a2;
		*(_BYTE*)result = a3;
		strncpy((char*)(result + 8), a1, 0x4Cu);
		*(_BYTE*)(v4 + 83) = 0;
		*(_DWORD*)(v4 + 84) = 0;
		v5 = dword_5d4594_2523900;
		*(_DWORD*)(v4 + 92) = 0;
		*(_DWORD*)(v4 + 88) = v5;
		if (dword_5d4594_2523900)
			*(_DWORD*)(dword_5d4594_2523900 + 92) = v4;
		dword_5d4594_2523900 = v4;
		result = 1;
	}
	return result;
}

//----- (0057C130) --------------------------------------------------------
int  sub_57C130(_DWORD* a1, int a2) {
	int result; // eax
	char* v3;   // esi
	char v4;    // bl
	int v5;     // edx
	int v6;     // ecx

	if (!a1)
		return 0;
	v3 = *(char**)&dword_5d4594_2523900;
	if (!dword_5d4594_2523900)
		return 0;
	while (*((_DWORD*)v3 + 1) != a2) {
		v3 = (char*)*((_DWORD*)v3 + 22);
		if (!v3)
			return 0;
	}
	if (!v3)
		return 0;
	v4 = *v3;
	result = (int)sub_57C360();
	if (result) {
		if (v4 && v4 != 1 && v4 != 3) {
			if (v4 != 2) {
				nox_alloc_class_free_obj(*(unsigned int**)&nox_alloc_itemGroupElem_2523896, (_QWORD*)result);
				return 0;
			}
			*(_DWORD*)result = *a1;
			*(_DWORD*)(result + 4) = a1[1];
		} else {
			*(_DWORD*)result = *a1;
		}
		v5 = *((_DWORD*)v3 + 21);
		*(_DWORD*)(result + 12) = 0;
		*(_DWORD*)(result + 8) = v5;
		v6 = *((_DWORD*)v3 + 21);
		if (v6)
			*(_DWORD*)(v6 + 12) = result;
		*((_DWORD*)v3 + 21) = result;
		result = 1;
	}
	return result;
}

//----- (0057C1E0) --------------------------------------------------------
int  sub_57C1E0(_DWORD* a1, int a2) {
	int v2;   // ebx
	char* v3; // ebp
	int i;    // esi
	char v5;  // al
	int v7;   // eax

	v2 = 0;
	v3 = *(char**)&dword_5d4594_2523900;
	if (!dword_5d4594_2523900)
		return 0;
	while (*(_DWORD*)(dword_5d4594_2523900 + 4) != a2)
		;
	for (i = *(_DWORD*)(dword_5d4594_2523900 + 84); i; i = *(_DWORD*)(i + 8)) {
		v5 = *v3;
		if (*v3 && v5 != 1 && v5 != 3) {
			if (v5 == 2 && *(_DWORD*)i == *a1 && *(_DWORD*)(i + 4) == a1[1]) {
				v2 = 1;
			LABEL_15:
				if (*(_DWORD*)(i + 12))
					*(_DWORD*)(i + 12) = *(_DWORD*)(i + 8);
				v7 = *(_DWORD*)(i + 8);
				if (v7)
					*(_DWORD*)(v7 + 12) = *(_DWORD*)(i + 12);
				sub_57C390((_QWORD*)i);
				continue;
			}
		} else if (*(_DWORD*)i == *a1) {
			v2 = 1;
			goto LABEL_15;
		}
		if (v2 == 1)
			goto LABEL_15;
	}
	return v2;
}

//----- (0057C280) --------------------------------------------------------
int  nox_server_scriptGetMapGroupByName_57C280(const char* a1, int a2) {
	int i; // esi

	for (i = nox_server_getFirstMapGroup_57C080(); i; i = nox_server_getNextMapGroup_57C090(i)) {
		if (a2 == nox_server_scriptGetGroupId_57C2D0((int**)i) &&
			nox_server_strcmpWithoutMapname_4DA3F0((const char*)(i + 8), a1))
			break;
	}
	return i;
}

//----- (0057C330) --------------------------------------------------------
void* sub_57C330() {
	void* v0;     // esi
	void* result; // eax

	v0 = 0;
	if (!nox_common_gameFlags_check_40A5C0(2097153) ||
		(result = nox_alloc_class_new_obj_zero(*(_DWORD**)&nox_alloc_groupInfo_2523892), (v0 = result) != 0))
		result = v0;
	return result;
}

//----- (0057C360) --------------------------------------------------------
void* sub_57C360() { return nox_alloc_class_new_obj_zero(*(_DWORD**)&nox_alloc_itemGroupElem_2523896); }

//----- (0057C370) --------------------------------------------------------
void  sub_57C370(_QWORD* a1) { nox_alloc_class_free_obj(*(unsigned int**)&nox_alloc_groupInfo_2523892, a1); }

//----- (0057C390) --------------------------------------------------------
void  sub_57C390(_QWORD* a1) { nox_alloc_class_free_obj(*(unsigned int**)&nox_alloc_itemGroupElem_2523896, a1); }

//----- (0057C3B0) --------------------------------------------------------
int  nox_server_addNewMapGroup_57C3B0(int a1) {
	int result; // eax

	result = a1;
	*(_DWORD*)(a1 + 92) = 0;
	*(_DWORD*)(a1 + 88) = dword_5d4594_2523900;
	if (dword_5d4594_2523900)
		*(_DWORD*)(dword_5d4594_2523900 + 92) = a1;
	dword_5d4594_2523900 = a1;
	return result;
}

//----- (0057C3E0) --------------------------------------------------------
int nox_xxx_getDebugData_57C3E0() { return dword_5d4594_2523912; }

//----- (0057C3F0) --------------------------------------------------------
int  nox_xxx_nextDebugObject_57C3F0(int a1) {
	int result; // eax

	if (a1)
		result = *(_DWORD*)(a1 + 336);
	else
		result = 0;
	return result;
}

//----- (0057C410) --------------------------------------------------------
BOOL nox_xxx_allocDebugDataArray_57C410() {
	nox_alloc_debugData_2523908 = nox_new_alloc_class("DebugData", 344, 256);
	return nox_alloc_debugData_2523908 != 0;
}

//----- (0057C440) --------------------------------------------------------
void sub_57C440() {
	nox_alloc_class_yyy_4144D0(*(_DWORD**)&nox_alloc_debugData_2523908);
	dword_5d4594_2523912 = 0;
}

//----- (0057C460) --------------------------------------------------------
int sub_57C460() {
	if (nox_alloc_debugData_2523908) {
		nox_free_alloc_class(*(LPVOID*)&nox_alloc_debugData_2523908);
		nox_alloc_debugData_2523908 = 0;
	}
	dword_5d4594_2523912 = 0;
	return 1;
}

//----- (0057C490) --------------------------------------------------------
int  sub_57C490(const char* a1) {
	int v1; // edi

	v1 = dword_5d4594_2523912;
	if (!dword_5d4594_2523912)
		return 0;
	while (strcmp(a1, (const char*)v1)) {
		v1 = *(_DWORD*)(v1 + 336);
		if (!v1)
			return 0;
	}
	return v1 + 80;
}

//----- (0057C500) --------------------------------------------------------
int  sub_57C500(const char* a1, const char* a2) {
	char* v2; // edx

	v2 = (char*)nox_alloc_class_new_obj_zero(*(_DWORD**)&nox_alloc_debugData_2523908);
	if (!v2)
		return 0;
	strcpy(v2, a1);
	strcpy(v2 + 80, a2);
	*((_DWORD*)v2 + 85) = 0;
	*((_DWORD*)v2 + 84) = dword_5d4594_2523912;
	if (dword_5d4594_2523912)
		*(_DWORD*)(dword_5d4594_2523912 + 340) = v2;
	dword_5d4594_2523912 = v2;
	return 1;
}

//----- (0057C5A0) --------------------------------------------------------
void  sub_57C5A0(const char* a1) {
	int v1; // edi
	int v2; // eax
	int v3; // eax

	v1 = dword_5d4594_2523912;
	if (dword_5d4594_2523912) {
		while (strcmp(a1, (const char*)v1)) {
			v1 = *(_DWORD*)(v1 + 336);
			if (!v1)
				return;
		}
		v2 = *(_DWORD*)(v1 + 340);
		if (v2)
			*(_DWORD*)(v2 + 336) = *(_DWORD*)(v1 + 336);
		else
			dword_5d4594_2523912 = *(_DWORD*)(v1 + 336);
		v3 = *(_DWORD*)(v1 + 336);
		if (v3)
			*(_DWORD*)(v3 + 340) = *(_DWORD*)(v1 + 340);
		nox_alloc_class_free_obj(*(unsigned int**)&nox_alloc_debugData_2523908, (_QWORD*)v1);
	}
}

//----- (0057C650) --------------------------------------------------------
void  sub_57C650(float* a1, float* a2, float* a3) {
	float* v3;      // ecx
	double v4;      // st7
	double v5;      // st6
	long double v6; // st7
	float* v7;      // edx
	long double v8; // st4
	double v9;      // st7
	float v10;      // [esp+0h] [ebp-10h]
	float v11;      // [esp+0h] [ebp-10h]
	float v12;      // [esp+Ch] [ebp-4h]
	float v13;      // [esp+14h] [ebp+4h]
	float v14;      // [esp+14h] [ebp+4h]
	float v15;      // [esp+18h] [ebp+8h]
	float v16;      // [esp+1Ch] [ebp+Ch]

	v3 = a1;
	v4 = a1[2] - *a1;
	v13 = v4;
	v5 = v3[3] - v3[1];
	v6 = sqrt(v5 * v5 + v4 * v4);
	if (v6 < 0.0099999998)
		v6 = 0.0099999998;
	v12 = v5;
	v7 = a3;
	v10 = v5;
	v8 = (v10 * (a2[1] - v3[1]) + v13 * (*a2 - *v3)) / (v6 * v6);
	v11 = v8 * v13;
	*a3 = v11 + *v3;
	a3[1] = v8 * v12 + v3[1];
	if (*v3 >= (double)v3[2]) {
		v15 = *v3;
		v14 = v3[2];
	} else {
		v14 = *v3;
		v15 = v3[2];
	}
	if (v3[1] >= (double)v3[3]) {
		v9 = v3[3];
		v16 = v3[1];
	} else {
		v9 = v3[1];
		v16 = v3[3];
	}
	if (*v7 >= (double)v14) {
		if (*v7 > (double)v15)
			*v7 = v15;
	} else {
		*v7 = v14;
	}
	if (v9 <= v7[1]) {
		if (v7[1] > (double)v16)
			v7[1] = v16;
	} else {
		v7[1] = v9;
	}
}

//----- (0057C790) --------------------------------------------------------
void  sub_57C790(float4* a1, float2* a2, float2* a3, float a4) {
	double v4; // st5
	double v5; // st7
	double v6; // st6
	double v7; // st7
	float v8;  // [esp+0h] [ebp-10h]
	float v9;  // [esp+4h] [ebp-Ch]
	float v10; // [esp+8h] [ebp-8h]
	float v11; // [esp+14h] [ebp+4h]
	float v12; // [esp+18h] [ebp+8h]
	float v13; // [esp+20h] [ebp+10h]

	v8 = a1->field_8 - a1->field_0;
	v4 = a1->field_C - a1->field_4;
	v9 = v4;
	v5 = v4 * (a2->field_4 - a1->field_4) + v8 * (a2->field_0 - a1->field_0);
	v6 = a4 * a4;
	v10 = v5 * v8 / v6;
	a3->field_0 = v10 + a1->field_0;
	a3->field_4 = v5 * v9 / v6 + a1->field_4;
	if (a1->field_0 >= (double)a1->field_8) {
		v12 = a1->field_0;
		v13 = a1->field_8;
	} else {
		v13 = a1->field_0;
		v12 = a1->field_8;
	}
	if (a1->field_4 >= (double)a1->field_C) {
		v7 = a1->field_C;
		v11 = a1->field_4;
	} else {
		v7 = a1->field_4;
		v11 = a1->field_C;
	}
	if (a3->field_0 >= (double)v13) {
		if (a3->field_0 > (double)v12)
			a3->field_0 = v12;
	} else {
		a3->field_0 = v13;
	}
	if (v7 <= a3->field_4) {
		if (a3->field_4 > (double)v11)
			a3->field_4 = v11;
	} else {
		a3->field_4 = v7;
	}
}

//----- (0057C8A0) --------------------------------------------------------
BOOL  nox_xxx_mathPointOnTheLine_57C8A0(float4* a1, float2* a2, float2* a3) {
	float4* v3; // ecx
	float2* v4; // edx
	double v5;  // st7
	double v6;  // st6
	double v7;  // st6
	double v8;  // st7
	float v10;  // [esp+0h] [ebp-Ch]
	float v11;  // [esp+0h] [ebp-Ch]
	float v12;  // [esp+8h] [ebp-4h]
	float v13;  // [esp+8h] [ebp-4h]
	float v14;  // [esp+10h] [ebp+4h]
	float v15;  // [esp+10h] [ebp+4h]
	float v16;  // [esp+14h] [ebp+8h]
	float v17;  // [esp+18h] [ebp+Ch]

	v3 = a1;
	v4 = a3;
	v5 = a1->field_8 - a1->field_0;
	v6 = a1->field_C - a1->field_4;
	v12 = v6;
	v10 = v6 * v12 + v5 * v5;
	v7 = (a2->field_4 - a1->field_4) * v12 + (a2->field_0 - a1->field_0) * v5;
	v14 = v7;
	v13 = v14 * v12 / v10;
	a3->field_0 = v5 * v7 / v10 + v3->field_0;
	v15 = v13 + v3->field_4;
	a3->field_4 = v15;
	if (v3->field_0 >= (double)v3->field_8) {
		v8 = v3->field_8;
		v16 = v3->field_0;
	} else {
		v8 = v3->field_0;
		v16 = v3->field_8;
	}
	if (v3->field_4 >= (double)v3->field_C) {
		v11 = v3->field_4;
		v17 = v3->field_C;
	} else {
		v17 = v3->field_4;
		v11 = v3->field_C;
	}
	return v8 <= v4->field_0 && v4->field_0 <= (double)v16 && v15 >= (double)v17 && v15 <= (double)v11;
}

//----- (0057C9A0) --------------------------------------------------------
char  nox_xxx_mapTraceRayImpl_57C9A0(int a1, int a2, float* a3, float* a4, char a5) {
	int v5;    // ebp
	char v6;   // bl
	double v7; // st7
	double v8; // st6
	int v9;    // eax
	int v10;   // eax
	float v11; // edx
	float v12; // ecx
	float v13; // eax
	float v14; // edx
	char v15;  // bl
	double v16;
	double v17;
	float v18;           // eax
	float v19;           // edx
	float v20;           // esi
	float2 a3a;          // [esp+10h] [ebp-30h]
	float2 v23;          // [esp+18h] [ebp-28h]
	float2 a2a;          // [esp+20h] [ebp-20h]
	float2 v25;          // [esp+28h] [ebp-18h]
	float v26;           // [esp+30h] [ebp-10h]
	float v27;           // [esp+34h] [ebp-Ch]
	float v28;           // [esp+38h] [ebp-8h]
	float v29;           // [esp+3Ch] [ebp-4h]
	char v30;            // [esp+44h] [ebp+4h]
	unsigned __int8 v31; // [esp+44h] [ebp+4h]
	unsigned __int8 v32; // [esp+48h] [ebp+8h]

	v5 = a1;
	if (a1 < 0 || a1 >= 256 || a2 < 0 || a2 >= 256)
		return 0;
	v6 = a5;
	if (a5 & 8) {
		v7 = (double)a1 * *getMemFloatPtr(0x587000, 313536) + 11.5 - *a3;
		v8 = (double)a2 * *getMemFloatPtr(0x587000, 313536) + 11.5 - a3[1];
		if (v8 * v8 + v7 * v7 < 3600.0)
			v6 = a5 | 4;
	}
	v30 = 16 * (~v6 & 4);
	if (v6 & 1)
		v30 = (16 * (~v6 & 4)) | 0x10;
	v31 = (unsigned char)sub_57B500(v5, a2, v30);
	if (v31 == 255)
		return 0;
	v9 = v6 & 1 ? nox_xxx_wall_4105E0(v5, a2) : nox_server_getWallAtPoint_410580(v5, a2);
	if (!v9 || v6 < 0 && *(_BYTE*)(v9 + 4) & 4 && *(_BYTE*)(*(_DWORD*)(v9 + 28) + 20) & 2)
		return 0;
	v10 = *getMemU32Ptr(0x5D4594, 12332 * *(unsigned __int8*)(v9 + 1) + 2692780);
	if (v10 & 2 || v6 & 0x40 && !(v10 & 1))
		return 0;
	if (*a3 >= (double)a3[2]) {
		v12 = a3[2];
		v28 = *a3;
		v26 = v12;
	} else {
		v11 = a3[2];
		v26 = *a3;
		v28 = v11;
	}
	if (a3[1] >= (double)a3[3]) {
		v14 = a3[3];
		v29 = a3[1];
		v27 = v14;
	} else {
		v13 = a3[3];
		v27 = a3[1];
		v29 = v13;
	}
	a2a.field_0 = (double)(23 * v5);
	a2a.field_4 = (double)(23 * a2);
	if (v31)
		sub_57CD30((float4*)a3, &a2a, &a3a);
	if (v31 != 1) {
		v25.field_0 = a2a.field_0 + 23.0;
		v25.field_4 = a2a.field_4;
		sub_57CD70((float4*)a3, &v25, &v23);
	}
	v15 = 0;
	v32 = 0;
	if (getMemByte(0x587000, 24 * v31 + 313272) && a2a.field_0 + *getMemFloatPtr(0x587000, 24 * v31 + 313276) <= a3a.field_0 &&
		a2a.field_0 + *getMemFloatPtr(0x587000, 24 * v31 + 313280) >= a3a.field_0 && a3a.field_0 >= (double)v26 &&
		a3a.field_0 <= (double)v28 && a3a.field_4 >= (double)v27 && a3a.field_4 <= (double)v29) {
		v15 = 1;
		*(float2*)a4 = a3a;
		v32 = 1;
	}
	if (getMemByte(0x587000, 24 * v31 + 313284) && a2a.field_0 + *getMemFloatPtr(0x587000, 24 * v31 + 313288) <= v23.field_0 &&
		a2a.field_0 + *getMemFloatPtr(0x587000, 24 * v31 + 313292) >= v23.field_0 && v23.field_0 >= (double)v26 &&
		v23.field_0 <= (double)v28 && v23.field_4 >= (double)v27 && v23.field_4 <= (double)v29) {
		++v15;
		*(float2*)&a4[2 * v32] = v23;
	}
	if (v15 == 2) {
		v16 = a3[1] - a3a.field_4;
		v17 = a3[1] - v23.field_4;
		if (v17 * v17 + (*a3 - v23.field_0) * (*a3 - v23.field_0) <
			v16 * v16 + (*a3 - a3a.field_0) * (*a3 - a3a.field_0)) {
			v18 = *a4;
			v19 = a4[1];
			*a4 = a4[2];
			v20 = a4[3];
			a4[2] = v18;
			a4[1] = v20;
			a4[3] = v19;
		}
	}
	return v15;
}

//----- (0057CD30) --------------------------------------------------------
float2*  sub_57CD30(float4* a1, float2* a2, float2* a3) {
	double v3;      // st7
	double v4;      // st6
	float2* result; // eax
	double v6;      // st7

	v3 = a1->field_8 - a1->field_0;
	v4 = a1->field_C - a1->field_4;
	result = a3;
	v6 = ((a2->field_4 - a1->field_4) * v3 - (a2->field_0 - a1->field_0) * v4) / (v4 - v3);
	a3->field_0 = v6 + a2->field_0;
	a3->field_4 = v6 + a2->field_4;
	return result;
}

//----- (0057CD70) --------------------------------------------------------
float2*  sub_57CD70(float4* a1, float2* a2, float2* a3) {
	double v3;      // st7
	double v4;      // st6
	float2* result; // eax
	double v6;      // st7

	v3 = a1->field_8 - a1->field_0;
	v4 = a1->field_C - a1->field_4;
	result = a3;
	v6 = ((a2->field_4 - a1->field_4) * v3 - (a2->field_0 - a1->field_0) * v4) / (-v4 - v3);
	a3->field_0 = a2->field_0 - v6;
	a3->field_4 = v6 + a2->field_4;
	return result;
}

//----- (0057CDB0) --------------------------------------------------------
int  sub_57CDB0(int2* a1, float* a2, float2* a3) {
	int2* v3;   // esi
	char v4;    // bl
	int result; // eax
	float2* v6; // eax
	float2* v7; // eax
	float2* v8; // eax
	char v9;    // [esp+10h] [ebp+4h]

	v3 = a1;
	v9 = sub_57F2A0((float2*)a2, a1->field_0, a1->field_4);
	v4 = sub_57F1D0((float2*)a2 + 1);
	switch (sub_57B500(v3->field_0, v3->field_4, 64)) {
	case 0:
		if (v9 != 1 && v9)
			goto LABEL_44;
		a3->field_0 = -0.70709997;
		a3->field_4 = -0.70709997;
		return 1;
	case 1:
		if (v9 == 1 || v9 == 2)
			goto LABEL_43;
		a3->field_0 = -0.70709997;
		a3->field_4 = 0.70709997;
		return 1;
	case 2:
		switch (v9) {
		case 0:
			goto LABEL_40;
		case 1:
			v6 = a3;
			if (!(v4 & 1))
				goto LABEL_23;
			a3->field_0 = -0.70709997;
			a3->field_4 = -0.70709997;
			result = 1;
			break;
		case 2:
			goto LABEL_30;
		case 3:
			goto LABEL_35;
		default:
			return 1;
		}
		return result;
	case 3:
		if (!v9)
			goto LABEL_40;
		if (v9 != 1)
			goto LABEL_44;
		goto LABEL_25;
	case 4:
		if (v9 == 1)
			goto LABEL_25;
		if (v9 == 2)
			goto LABEL_30;
		a3->field_0 = -0.70709997;
		a3->field_4 = 0.70709997;
		return 1;
	case 5:
		if (v9 == 2)
			goto LABEL_30;
		if (v9 == 3)
			goto LABEL_35;
		a3->field_0 = -0.70709997;
		a3->field_4 = -0.70709997;
		return 1;
	case 6:
		if (!v9)
			goto LABEL_40;
		if (v9 == 3)
			goto LABEL_35;
		v6 = a3;
	LABEL_23:
		v6->field_0 = 0.70709997;
		v6->field_4 = -0.70709997;
		return 1;
	case 7:
		if (v9 == 1) {
		LABEL_25:
			v7 = a3;
			if (!(v4 & 1))
				goto LABEL_38;
			goto LABEL_26;
		}
		v8 = a3;
		if (v4 & 1)
			goto LABEL_45;
		a3->field_0 = -0.70709997;
		a3->field_4 = 0.70709997;
		return 1;
	case 8:
		if (v9 == 2) {
		LABEL_30:
			v8 = a3;
			a3->field_0 = 0.70709997;
			if (v4 & 1)
				goto LABEL_46;
			a3->field_4 = -0.70709997;
			result = 1;
		} else {
			v8 = a3;
			a3->field_0 = -0.70709997;
			if (!(v4 & 1))
				goto LABEL_46;
			a3->field_4 = -0.70709997;
			result = 1;
		}
		return result;
	case 9:
		if (v9 == 3) {
		LABEL_35:
			v8 = a3;
			if (!(v4 & 4))
				goto LABEL_45;
			a3->field_0 = -0.70709997;
			a3->field_4 = 0.70709997;
			result = 1;
		} else {
			v7 = a3;
			if (v4 & 4) {
			LABEL_38:
				v7->field_0 = 0.70709997;
				v7->field_4 = -0.70709997;
				result = 1;
			} else {
			LABEL_26:
				v7->field_0 = -0.70709997;
				v7->field_4 = -0.70709997;
				result = 1;
			}
		}
		return result;
	case 0xA:
		if (v9) {
			if (v4 & 2) {
			LABEL_43:
				a3->field_0 = 0.70709997;
				a3->field_4 = -0.70709997;
				return 1;
			}
		LABEL_44:
			v8 = a3;
		LABEL_45:
			v8->field_0 = 0.70709997;
		} else {
		LABEL_40:
			v8 = a3;
			a3->field_0 = -0.70709997;
			if (!(v4 & 2)) {
				a3->field_4 = -0.70709997;
				return 1;
			}
		}
	LABEL_46:
		v8->field_4 = 0.70709997;
		return 1;
	default:
		return 0;
	}
}

//----- (0057D150) --------------------------------------------------------
void __thiscall sub_57D150(LPVOID* this) {
	LPVOID* v1; // esi

	v1 = this;
	operator_delete(this[5]);
	sub_57DF70(v1 + 2);
	operator_delete(*v1);
}

//----- (0057D1C0) --------------------------------------------------------
int __thiscall sub_57D1C0(void** this, int a2, unsigned __int8* a3, int a4) {
	void** v4;             // ebx
	void* v5;              // ecx
	int v6;                // eax
	unsigned __int8* v7;   // esi
	int v8;                // ecx
	int v9;                // eax
	int v10;               // edx
	int v11;               // edi
	int v12;               // ebp
	_WORD* v13;            // ecx
	unsigned int v14;      // eax
	int v15;               // esi
	_WORD* v16;            // eax
	int v17;               // ecx
	int v18;               // eax
	int v19;               // ebp
	int v20;               // eax
	unsigned __int8* i;    // eax
	int v22;               // ecx
	int v23;               // ecx
	int v24;               // eax
	int v25;               // edx
	int v26;               // eax
	int v27;               // ecx
	unsigned __int8* v28;  // eax
	int v29;               // edx
	int v30;               // eax
	unsigned __int8 v31;   // cl
	int v32;               // eax
	int v33;               // esi
	int v34;               // esi
	char* v35;             // eax
	int v36;               // eax
	int v37;               // esi
	int v38;               // ecx
	unsigned __int8* v39;  // edx
	unsigned __int8* v40;  // edi
	unsigned __int8* v41;  // eax
	unsigned int v42;      // ecx
	int v43;               // edx
	int v44;               // eax
	int v45;               // edi
	int v46;               // edx
	int v47;               // edi
	int v48;               // eax
	char* v49;             // edi
	unsigned __int8* v50;  // esi
	unsigned int v51;      // eax
	char* v52;             // edi
	unsigned __int8* v53;  // esi
	char v54;              // cl
	int v55;               // esi
	int v56;               // ecx
	int v57;               // ebp
	unsigned __int8* v58;  // edx
	int v59;               // ecx
	int v60;               // esi
	unsigned __int8* v61;  // eax
	unsigned int v62;      // ebp
	int v63;               // ecx
	int v64;               // eax
	int v65;               // esi
	int v66;               // ecx
	int v67;               // edx
	int v68;               // edi
	signed int v69;        // ebp
	int v70;               // eax
	int v71;               // esi
	int v72;               // ecx
	unsigned __int8* v73;  // edi
	unsigned __int8* v74;  // eax
	unsigned int v75;      // ecx
	int v76;               // edx
	int v77;               // eax
	int v78;               // edi
	int v79;               // edx
	int v80;               // eax
	char* v81;             // edi
	unsigned __int8* v82;  // esi
	unsigned int v83;      // eax
	char* v84;             // edi
	int v85;               // ebp
	unsigned __int8* v86;  // edx
	int v87;               // esi
	int v88;               // ecx
	unsigned __int8* v89;  // edi
	unsigned __int8* v90;  // eax
	unsigned int v91;      // ebp
	int v92;               // ecx
	int v93;               // eax
	int v94;               // edi
	int v95;               // eax
	int v96;               // ecx
	int v97;               // edx
	int v98;               // edi
	signed int v99;        // ebp
	int v100;              // eax
	unsigned __int8* v101; // esi
	int v102;              // edi
	int v103;              // edx
	unsigned __int8* v104; // edi
	unsigned int v105;     // ecx
	int v106;              // edi
	unsigned __int8* v107; // edx
	int v108;              // ebp
	signed int v110;       // [esp+10h] [ebp-5Ch]
	int v111;              // [esp+14h] [ebp-58h]
	int v112;              // [esp+18h] [ebp-54h]
	int v113;              // [esp+18h] [ebp-54h]
	int v114;              // [esp+18h] [ebp-54h]
	int v115;              // [esp+18h] [ebp-54h]
	int v116;              // [esp+18h] [ebp-54h]
	int v117;              // [esp+1Ch] [ebp-50h]
	unsigned __int8* v118; // [esp+20h] [ebp-4Ch]
	int v119;              // [esp+24h] [ebp-48h]
	int v120;              // [esp+24h] [ebp-48h]
	int v121;              // [esp+24h] [ebp-48h]
	unsigned __int16 v122; // [esp+24h] [ebp-48h]
	int v123;              // [esp+24h] [ebp-48h]
	int v124;              // [esp+24h] [ebp-48h]
	int v125;              // [esp+28h] [ebp-44h]
	int v126;              // [esp+2Ch] [ebp-40h]
	int v127;              // [esp+30h] [ebp-3Ch]
	__int16 v128;          // [esp+34h] [ebp-38h]
	int v129;              // [esp+38h] [ebp-34h]
	int v130;              // [esp+38h] [ebp-34h]
	int v131;              // [esp+38h] [ebp-34h]
	int v132;              // [esp+38h] [ebp-34h]
	void* v133;            // [esp+3Ch] [ebp-30h]
	int v134;              // [esp+40h] [ebp-2Ch]
	int v135[10];          // [esp+44h] [ebp-28h]
	int v136;              // [esp+70h] [ebp+4h]
	unsigned __int8* v137; // [esp+70h] [ebp+4h]
	unsigned __int8* v138; // [esp+78h] [ebp+Ch]
	unsigned __int8* v139; // [esp+78h] [ebp+Ch]
	unsigned __int8* v140; // [esp+78h] [ebp+Ch]
	int v141;              // [esp+78h] [ebp+Ch]
	unsigned __int8* v142; // [esp+78h] [ebp+Ch]
	unsigned __int8* v143; // [esp+78h] [ebp+Ch]

	v4 = this;
	v5 = this[2];
	v135[0] = (int)(v4 + 2);
	v135[2] = (int)(v4 + 3);
	v135[4] = a2;
	v135[3] = a2;
	v6 = 0;
	v135[1] = (int)v5;
	v135[5] = 0;
	v135[6] = 0;
	v135[9] = 0;
	v136 = 0;
	v118 = &a3[a4];
	if ((unsigned int)a4 >= 5) {
		v7 = a3;
		if (a3 < a3 + 5) {
			do {
				v8 = *v7++;
				v6 = __ROL4__(v8 ^ v6, 5);
			} while (v7 < a3 + 5);
			v136 = v6;
		}
	}
	v9 = a4;
	v111 = a4;
	if (a4 < 5)
		goto LABEL_144;
	v10 = v136;
	while (2) {
		v11 = 0;
		v126 = v9 - 5;
		if (v9 - 5 >= 64)
			v126 = 64;
		v110 = 0;
		v125 = 0;
		while (1) {
			HIWORD(v15) = 0;
			v12 = 0;
			v13 = v4[5];
			v14 = (unsigned int)(214013 * v10 + 2531011) >> 17;
			LOWORD(v15) = v13[v14];
			v16 = &v13[v14];
			v17 = (unsigned int)v4[1] & 0xFFFF;
			v117 = v17;
			*v16 = v17;
			if ((unsigned __int16)v15 == 0xFFFF || v15 == v17)
				goto LABEL_62;
			v18 = (unsigned __int16)(v17 - v15);
			v19 = v18;
			if (v18 >= v111 - v11)
				v19 = v111 - v11;
			if (v19 >= 521) {
				v18 = 521;
			} else if (v18 >= v111 - v11) {
				v18 = v111 - v11;
			}
			v112 = v18;
			v138 = &a3[v11];
			v119 = 0x10000 - v15;
			if (0x10000 - v15 < v18) {
				v22 = v15;
				if (v15 >= 0x10000) {
				LABEL_33:
					v23 = 0;
					if (v18 + v15 - 0x10000 <= 0) {
					LABEL_23:
						v12 = v18;
						goto LABEL_24;
					}
					while (*((_BYTE*)*v4 + v23) == v138[v119 + v23]) {
						v18 = v112;
						if (++v23 >= v112 + v15 - 0x10000)
							goto LABEL_23;
					}
					v12 = v119 + v23;
				} else {
					while (*((_BYTE*)*v4 + v22) == v138[v22 - v15]) {
						if (++v22 >= 0x10000) {
							v10 = v136;
							goto LABEL_33;
						}
					}
					v10 = v136;
					v12 = v22 - v15;
				}
			} else {
				v12 = 0;
				if (v18 <= 0)
					goto LABEL_23;
				while (*((_BYTE*)*v4 + v15 + v12) == a3[v11 + v12]) {
					v18 = v112;
					if (++v12 >= v112)
						goto LABEL_23;
				}
			}
		LABEL_24:
			if ((unsigned __int16)(v15 + v12) == v117) {
				v20 = v111 + -v12 - v11;
				if (521 - v12 < v20)
					v20 = 521 - v12;
				v120 = v20;
				v113 = 0;
				if (v20 > 0) {
					for (i = &a3[v11]; *i == i[v12]; ++i) {
						if (++v113 >= v120)
							break;
					}
				}
				v12 += v113;
			}
			if (v12 >= 3) {
				v114 = 521 - v12;
				if (521 - v12 >= v11)
					v114 = v11;
				v24 = (unsigned __int16)((char*)v4[1] - v15);
				if (v114 >= v24 - v12)
					v121 = (unsigned __int16)((char*)v4[1] - v15) - v12;
				else
					v121 = v114;
				v25 = 0x10000 - v24;
				if (v121 < 0x10000 - v24) {
					v25 = v114;
					if (v114 < v24 - v12) {
					LABEL_53:
						v26 = 0;
						if (v25 <= 0)
							goto LABEL_152;
						v122 = v15 - 1;
						v139 = v138 - 1;
						do {
							if (*((_BYTE*)*v4 + v122) != *v139)
								break;
							++v26;
							--v122;
							--v139;
						} while (v26 < v114);
						if (v26 <= 0) {
						LABEL_152:
							v10 = v136;
						} else {
							v11 -= v26;
							v27 = (int)v4[1] - v26;
							LOWORD(v15) = v15 - v26;
							v12 += v26;
							v4[1] = (void*)v27;
							v28 = &a3[v11];
							LOWORD(v117) = v27;
							v10 = 0;
							if (v28 < v28 + 5) {
								do {
									v29 = *v28++ ^ v10;
									v10 = __ROL4__(v29, 5);
								} while (v28 < &a3[v11 + 5]);
							}
							v136 = v10;
							v125 = 1;
						}
						goto LABEL_62;
					}
					v25 = (unsigned __int16)((char*)v4[1] - v15) - v12;
				}
				v114 = v25;
				goto LABEL_53;
			}
		LABEL_62:
			v30 = v110;
			if (v110 >= 4) {
				v115 = v11;
				if (v12 <= v110) {
					v4[1] = v133;
					v137 = &a3[v127];
					v55 = __ROL4__(v134 ^ a3[v127 + 5] ^ __ROL4__(a3[v127], 25), 5);
					sub_57E4C0((_DWORD**)v135, (unsigned int)a3, v127, v110 - 4,
							   (unsigned __int16)((_WORD)v133 - v128));
					v56 = &v118[-v127] - a3;
					v57 = v110 - 2;
					if (v110 - 2 >= v56 - 2)
						v57 = v56 - 2;
					v58 = &a3[v127];
					v59 = (int)v4[1] + 2;
					v141 = (int)v4[1] + 2;
					if (v57 > 0) {
						if (v57 <= 1024) {
							v64 = __ROL4__(v55 ^ v137[6] ^ __ROL4__(v137[1], 25), 5);
							v65 = 0;
							v130 = (int)(v137 + 7);
							while (1) {
								*((_WORD*)v4[5] + ((unsigned int)(214013 * v64 + 2531011) >> 17)) = v65 + v59;
								v66 = *(unsigned __int8*)(v130 + v65);
								v67 = __ROL4__(*(unsigned __int8*)(v130 + v65++ - 5), 25);
								v64 = __ROL4__(v64 ^ v66 ^ v67, 5);
								if (v65 >= v57)
									break;
								LOWORD(v59) = v141;
							}
							v58 = &a3[v127];
							v136 = v64;
						} else {
							v60 = 0;
							v61 = &v137[v57 + 2];
							v62 = (unsigned int)(v61 + 5);
							if (v61 < v61 + 5) {
								do {
									v63 = *v61++;
									v60 = __ROL4__(v63 ^ v60, 5);
								} while ((unsigned int)v61 < v62);
							}
							v136 = v60;
						}
					} else {
						v136 = 0;
					}
					v68 = (unsigned int)v4[1] & 0xFFFF;
					if (v68 + v110 <= 0x10000) {
						v69 = v110;
						memcpy((char*)*v4 + v68, v58, v110);
					} else {
						memcpy((char*)*v4 + v68, v58, 0x10000 - v68);
						v69 = v110;
						memcpy(*v4, &v58[0x10000 - v68], v110 - (0x10000 - v68));
					}
					v4[1] = (char*)v4[1] + v69;
					a3 += v69 + v127;
					goto LABEL_143;
				}
				sub_57E4C0((_DWORD**)v135, (unsigned int)a3, v11, v12 - 4, (unsigned __int16)(v117 - v15));
				v36 = &v118[-v11] - a3 - 1;
				if (v12 - 1 < v36)
					v36 = v12 - 1;
				v37 = 0;
				v38 = (int)v4[1] + 1;
				v39 = &a3[v11];
				v123 = v36;
				v129 = (int)v4[1] + 1;
				v140 = &a3[v11];
				v40 = &a3[v11];
				if (v36 > 0) {
					if (v36 <= 1024) {
						v44 = __ROL4__(v136 ^ v40[5] ^ __ROL4__(*v40, 25), 5);
						v45 = (int)(v40 + 1);
						if (v123 > 0) {
							while (1) {
								*((_WORD*)v4[5] + ((unsigned int)(214013 * v44 + 2531011) >> 17)) = v37 + v38;
								v46 = v44 ^ *(unsigned __int8*)(v45 + 5 + v37) ^
									  __ROL4__(*(unsigned __int8*)(v45 + v37), 25);
								++v37;
								v44 = __ROL4__(v46, 5);
								if (v37 >= v123)
									break;
								LOWORD(v38) = v129;
							}
						}
						v136 = v44;
					} else {
						v41 = &v40[v36 + 1];
						v42 = (unsigned int)(v41 + 5);
						if (v41 < v41 + 5) {
							do {
								v43 = *v41++;
								v37 = __ROL4__(v43 ^ v37, 5);
							} while ((unsigned int)v41 < v42);
						}
						v136 = v37;
					}
					v39 = v140;
				} else {
					v136 = 0;
				}
				v47 = (unsigned int)v4[1] & 0xFFFF;
				if (v47 + v12 > 0x10000) {
					v48 = 0x10000 - v47;
					memcpy((char*)*v4 + v47, v39, 0x10000 - v47);
					v49 = (char*)*v4;
					v50 = &v140[v48];
					v51 = v12 - v48;
					memcpy(*v4, v50, 4 * (v51 >> 2));
					v53 = &v50[4 * (v51 >> 2)];
					v52 = &v49[4 * (v51 >> 2)];
					v54 = v51;
				LABEL_118:
					memcpy(v52, v53, v54 & 3);
					v4[1] = (char*)v4[1] + v12;
					a3 += v115 + v12;
					goto LABEL_143;
				}
			LABEL_117:
				v84 = (char*)*v4 + v47;
				memcpy(v84, v39, 4 * ((unsigned int)v12 >> 2));
				v53 = &v39[4 * ((unsigned int)v12 >> 2)];
				v52 = &v84[4 * ((unsigned int)v12 >> 2)];
				v54 = v12;
				goto LABEL_118;
			}
			if (v12 < 4)
				goto LABEL_66;
			if (v125) {
				v115 = v11;
				sub_57E4C0((_DWORD**)v135, (unsigned int)a3, v11, v12 - 4, (unsigned __int16)(v117 - v15));
				v70 = &v118[-v11] - a3 - 1;
				if (v12 - 1 < v70)
					v70 = v12 - 1;
				v71 = 0;
				v72 = (int)v4[1] + 1;
				v39 = &a3[v11];
				v124 = v70;
				v131 = (int)v4[1] + 1;
				v142 = &a3[v11];
				v73 = &a3[v11];
				if (v70 > 0) {
					if (v70 <= 1024) {
						v77 = __ROL4__(v136 ^ v73[5] ^ __ROL4__(*v73, 25), 5);
						v78 = (int)(v73 + 1);
						if (v124 > 0) {
							while (1) {
								*((_WORD*)v4[5] + ((unsigned int)(214013 * v77 + 2531011) >> 17)) = v71 + v72;
								v79 = v77 ^ *(unsigned __int8*)(v78 + 5 + v71) ^
									  __ROL4__(*(unsigned __int8*)(v78 + v71), 25);
								++v71;
								v77 = __ROL4__(v79, 5);
								if (v71 >= v124)
									break;
								LOWORD(v72) = v131;
							}
						}
						v136 = v77;
					} else {
						v74 = &v73[v70 + 1];
						v75 = (unsigned int)(v74 + 5);
						if (v74 < v74 + 5) {
							do {
								v76 = *v74++;
								v71 = __ROL4__(v76 ^ v71, 5);
							} while ((unsigned int)v74 < v75);
						}
						v136 = v71;
					}
					v39 = v142;
				} else {
					v136 = 0;
				}
				v47 = (unsigned int)v4[1] & 0xFFFF;
				if (v47 + v12 > 0x10000) {
					v80 = 0x10000 - v47;
					memcpy((char*)*v4 + v47, v39, 0x10000 - v47);
					v81 = (char*)*v4;
					v82 = &v142[v80];
					v83 = v12 - v80;
					memcpy(*v4, v82, 4 * (v83 >> 2));
					v53 = &v82[4 * (v83 >> 2)];
					v52 = &v81[4 * (v83 >> 2)];
					v54 = v83;
					goto LABEL_118;
				}
				goto LABEL_117;
			}
			v30 = v12;
			v110 = v12;
			v128 = v15;
			v127 = v11;
			v133 = v4[1];
			v134 = v10;
		LABEL_66:
			if (v11 + 1 > v126)
				break;
			v31 = a3[v11];
			v32 = a3[v11 + 5];
			v33 = a3[v11++];
			v34 = v32 ^ __ROL4__(v33, 25);
			v35 = (char*)v4[1];
			v10 = __ROL4__(v10 ^ v34, 5);
			v4[1] = v35 + 1;
			v136 = v10;
			*((_BYTE*)*v4 + (unsigned __int16)v35) = v31;
		}
		v116 = v11;
		if (v30 < 4) {
			if (v11 + 5 >= v111 && v111 <= 64) {
				v100 = v111 - v11;
				v101 = &a3[v11];
				v102 = (unsigned int)v4[1] & 0xFFFF;
				if (v102 + v100 <= 0x10000) {
					v105 = v100;
					v104 = (unsigned __int8*)*v4 + v102;
				} else {
					v103 = 0x10000 - v102;
					memcpy((char*)*v4 + v102, v101, 0x10000 - v102);
					v104 = (unsigned __int8*)*v4;
					v105 = v100 - v103;
					v101 += v103;
				}
				memcpy(v104, v101, v105);
				v4[1] = (char*)v4[1] + v100;
				v11 = v111;
			}
			sub_57E3F0((_DWORD**)v135, (unsigned int)a3, v11);
			a3 += v11;
		} else {
			sub_57E4C0((_DWORD**)v135, (unsigned int)a3, v11, v110 - 4, (unsigned __int16)((char*)v4[1] - v128));
			v85 = v110 - 1;
			if (v110 - 1 >= &v118[-v11] - a3 - 1)
				v85 = &v118[-v11] - a3 - 1;
			v86 = &a3[v11];
			v87 = 0;
			v88 = (int)v4[1] + 1;
			v132 = (int)v4[1] + 1;
			v143 = &a3[v11];
			v89 = &a3[v11];
			if (v85 > 0) {
				if (v85 <= 1024) {
					v93 = v136 ^ v89[5] ^ __ROL4__(*v86, 25);
					v94 = (int)(v89 + 1);
					v95 = __ROL4__(v93, 5);
					while (1) {
						*((_WORD*)v4[5] + ((unsigned int)(214013 * v95 + 2531011) >> 17)) = v87 + v88;
						v96 = *(unsigned __int8*)(v94 + 5 + v87);
						v97 = __ROL4__(*(unsigned __int8*)(v94 + v87++), 25);
						v95 = __ROL4__(v95 ^ v96 ^ v97, 5);
						if (v87 >= v85)
							break;
						LOWORD(v88) = v132;
					}
					v86 = v143;
					v136 = v95;
				} else {
					v90 = &v86[v85 + 1];
					v91 = (unsigned int)(v90 + 5);
					if (v90 < v90 + 5) {
						do {
							v92 = *v90++;
							v87 = __ROL4__(v92 ^ v87, 5);
						} while ((unsigned int)v90 < v91);
					}
					v136 = v87;
				}
			} else {
				v136 = 0;
			}
			v98 = (unsigned int)v4[1] & 0xFFFF;
			if (v98 + v110 <= 0x10000) {
				v99 = v110;
				memcpy((char*)*v4 + v98, v86, v110);
			} else {
				memcpy((char*)*v4 + v98, v86, 0x10000 - v98);
				v99 = v110;
				memcpy(*v4, &v86[0x10000 - v98], v110 - (0x10000 - v98));
			}
			v4[1] = (char*)v4[1] + v99;
			a3 += v116 + v99;
		}
	LABEL_143:
		v9 = v118 - a3;
		v111 = v118 - a3;
		if (v118 - a3 >= 5) {
			v10 = v136;
			continue;
		}
		break;
	}
LABEL_144:
	if (v111) {
		v106 = (unsigned int)v4[1] & 0xFFFF;
		if (v106 + v111 <= 0x10000) {
			v108 = v111;
			memcpy((char*)*v4 + v106, a3, v111);
			v107 = a3;
		} else {
			v107 = a3;
			memcpy((char*)*v4 + v106, a3, 0x10000 - v106);
			v108 = v111;
			memcpy(*v4, &a3[0x10000 - v106], v111 - (0x10000 - v106));
		}
		v4[1] = (char*)v4[1] + v108;
		sub_57E3F0((_DWORD**)v135, (unsigned int)v107, v108);
	}
	return sub_57E7D0((_DWORD**)v135);
}
// 57D736: variable 'v133' is possibly undefined
// 57D739: variable 'v127' is possibly undefined
// 57D75C: variable 'v134' is possibly undefined
// 57D764: variable 'v128' is possibly undefined

//----- (0057DD70) --------------------------------------------------------
void __thiscall sub_57DD70(LPVOID* this) { operator_delete(*this); }

//----- (0057DD90) --------------------------------------------------------
_DWORD* __thiscall sub_57DD90(_DWORD* this) {
	_DWORD* v1; // esi
	void* v2;   // eax

	v1 = this;
	v2 = operator_new(0x224u);
	*v1 = v2;
	memset(v2, 0, 0x224u);
	return v1;
}
// 5667CB: using guessed type void * operator_new(unsigned int);

//----- (0057DDC0) --------------------------------------------------------
void __thiscall sub_57DDC0(LPVOID* this) { operator_delete(*this); }

//----- (0057DDD0) --------------------------------------------------------
int __thiscall sub_57DDD0(void** this) {
	int result; // eax

	result = 0;
	memset(*this, 0, 0x224u);
	return result;
}

//----- (0057DDE0) --------------------------------------------------------
unsigned int  sub_57DDE0(int a1, int a2) {
	int v2;              // ecx
	int v3;              // esi
	int i;               // ebx
	int v5;              // eax
	int v6;              // edx
	int v7;              // edi
	int v8;              // esi
	unsigned int result; // eax
	int v10;             // [esp+10h] [ebp-8h]
	int v11;             // [esp+14h] [ebp-4h]
	int v12;             // [esp+1Ch] [ebp+4h]

	v2 = a1;
	v3 = a2;
	for (i = 40; i > 0; i /= 3) {
		v5 = i + 1;
		v11 = i + 1;
		if (i + 1 <= v3) {
			do {
				v6 = 4 * v5;
				v10 = v5;
				v12 = *(_DWORD*)(4 * v5 + v2);
				if (v5 > i) {
					v7 = 4 * i;
					do {
						v8 = *(__int16*)(v6 - v7 + v2 + 2) - SHIWORD(v12);
						if (!v8)
							v8 = *(__int16*)(v6 - v7 + v2) - (__int16)v12;
						if (v8 >= 0)
							break;
						*(_DWORD*)(v6 + v2) = *(_DWORD*)(v6 - v7 + v2);
						v6 -= v7;
						v10 -= i;
					} while (v10 > i);
					v5 = v11;
				}
				++v5;
				*(_DWORD*)(v2 + 4 * v10) = v12;
				v3 = a2;
				v11 = v5;
			} while (v5 <= a2);
		}
		result = (unsigned int)(i / 3) >> 31;
	}
	return result;
}

//----- (0057DEA0) --------------------------------------------------------
int __thiscall sub_57DEA0(_DWORD* this, _WORD* a2) {
	_WORD* v2;   // esi
	int v3;      // ebx
	int i;       // eax
	__int16* v5; // edi
	__int16 v6;  // dx

	v2 = a2;
	v3 = 0;
	for (i = 0; i < 274; ++i) {
		*v2 = i;
		v2 += 2;
		*(v2 - 1) = *(_WORD*)(*this + 2 * i);
		v5 = (__int16*)(*this + 2 * i);
		v6 = *v5;
		v3 += v6;
		*v5 = v6 >> 1;
	}
	sub_57DDE0((int)(a2 - 2), 274);
	return v3;
}

//----- (0057DF00) --------------------------------------------------------
_DWORD* __thiscall sub_57DF00(_DWORD* this) {
	_DWORD* v1;          // ebx
	unsigned __int8* v2; // eax
	unsigned __int8* v3; // edi
	_DWORD* result;      // eax

	v1 = this;
	sub_57DD90(this);
	v1[1] = 4096;
	v2 = (unsigned __int8*)operator_new(0x448u);
	v3 = v2;
	v1[2] = v2;
	result = v1;
	memcpy(v3, getMemAt(0x587000, 313544), 0x448u);
	return result;
}
// 5667CB: using guessed type void * operator_new(unsigned int);

//----- (0057DF70) --------------------------------------------------------
void __thiscall sub_57DF70(LPVOID* this) {
	LPVOID* v1; // esi

	v1 = this;
	operator_delete(this[2]);
	sub_57DDC0(v1);
}

//----- (0057DFC0) --------------------------------------------------------
int __thiscall sub_57DFC0(_DWORD* this, int* a2) {
	_DWORD* v2;     // esi
	int v3;         // eax
	int v4;         // ecx
	int v5;         // edi
	int v6;         // ebp
	int v7;         // esi
	int v8;         // ebx
	int v9;         // eax
	int v10;        // edx
	int v11;        // eax
	char* v12;      // edx
	int v13;        // eax
	int* v14;       // edx
	int v15;        // edx
	int v16;        // ebx
	int v17;        // esi
	char* v18;      // eax
	int v19;        // ecx
	int v20;        // esi
	int v21;        // eax
	int v22;        // ecx
	int v23;        // eax
	char* v24;      // ecx
	int v25;        // ecx
	int v26;        // eax
	int* v27;       // edi
	int v28;        // ecx
	int* v29;       // eax
	int v30;        // ecx
	int* v31;       // eax
	int v32;        // ebx
	int v33;        // esi
	int v34;        // ecx
	int v35;        // edi
	int result;     // eax
	int v37;        // edx
	__int16 v38;    // di
	int v39;        // esi
	char* v40;      // ecx
	int v41;        // eax
	__int16 v42;    // bp
	int* v43;       // [esp+10h] [ebp-46Ch]
	int v44;        // [esp+10h] [ebp-46Ch]
	int v45;        // [esp+10h] [ebp-46Ch]
	int v46;        // [esp+14h] [ebp-468h]
	int v47;        // [esp+14h] [ebp-468h]
	int* v48;       // [esp+14h] [ebp-468h]
	int v49;        // [esp+18h] [ebp-464h]
	int v50;        // [esp+18h] [ebp-464h]
	int v51;        // [esp+18h] [ebp-464h]
	int v52;        // [esp+1Ch] [ebp-460h]
	int v53;        // [esp+1Ch] [ebp-460h]
	int v54;        // [esp+1Ch] [ebp-460h]
	int v55;        // [esp+20h] [ebp-45Ch]
	int v56;        // [esp+24h] [ebp-458h]
	int v57;        // [esp+24h] [ebp-458h]
	int v58;        // [esp+28h] [ebp-454h]
	int v59;        // [esp+2Ch] [ebp-450h]
	char v60[1100]; // [esp+30h] [ebp-44Ch]

	v2 = this;
	*(_DWORD*)v60 = this;
	v3 = sub_57DEA0(this, &v60[4]);
	v4 = 0;
	v5 = 0;
	v6 = 16;
	v56 = v3;
	v2[1] = 4096;
	v55 = 0;
	v46 = 0;
	v52 = 16;
	v43 = a2 - 1;
	while (1) {
		v7 = 0;
		v58 = 0;
		v8 = 0;
		v59 = (v56 - v55) / v6;
		while (1) {
			v49 = 0;
			v9 = 1 << v4;
			if (v5 + (1 << v4) + v7 > 274) {
				v49 = 1;
				v9 = 274 - v5;
			}
			if (v7 < v9) {
				v10 = v7 + v5;
				v11 = v9 - v7;
				v7 += v11;
				v12 = &v60[4 * v10 + 6];
				do {
					v8 += *(__int16*)v12;
					v12 += 4;
					--v11;
				} while (v11);
				v6 = v52;
			}
			if (v49 || v4 >= 8 || v8 > v59)
				break;
			v58 = v8;
			++v4;
		}
		if (v4 && abs32(v58 - v59) <= abs32(v8 - v59)) {
			v8 = v58;
			--v4;
		}
		v13 = v46;
		if (v6 < 16) {
			v14 = v43;
			do {
				if (v4 >= *v14)
					break;
				v14[1] = *v14;
				--v13;
				--v14;
			} while (v13 > 0);
		}
		a2[v13] = v4;
		v55 += v8;
		++v43;
		v5 += 1 << v4;
		--v6;
		++v46;
		v52 = v6;
		if (v6 <= 2)
			break;
		v4 = 0;
	}
	v15 = 0;
	v16 = 0;
	v17 = 0;
	v44 = 0;
	v47 = 0;
	v50 = 0x7FFFFFFF;
	v53 = 0;
	if (v5 < 274) {
		v18 = &v60[4 * v5 + 6];
		v19 = 274 - v5;
		do {
			v17 += *(__int16*)v18;
			v18 += 4;
			--v19;
		} while (v19);
		v53 = v17;
	}
	v20 = 0;
	v21 = 1;
	if (v5 + 1 <= 274) {
		do {
			if (v15 < v21) {
				v22 = v15 + v5;
				v23 = v21 - v15;
				v15 += v23;
				v24 = &v60[4 * v22 + 6];
				do {
					v16 += *(__int16*)v24;
					v24 += 4;
					--v23;
				} while (v23);
			}
			v25 = 0;
			v26 = 274 - v15 - v5;
			if (v26 > 1) {
				do
					++v25;
				while (1 << v25 < v26);
			}
			if (v20 <= 8 && v25 <= 8) {
				if (v16 * v20 + v25 * (v53 - v16) >= v50)
					break;
				v50 = v16 * v20 + v25 * (v53 - v16);
				v44 = v20;
				v47 = v25;
			}
			v21 = 1 << ++v20;
		} while (v5 + (1 << v20) + v15 <= 274);
	}
	v27 = a2;
	v28 = 14;
	v29 = a2 + 13;
	do {
		if (v44 >= *v29)
			break;
		v29[1] = *v29;
		--v28;
		--v29;
	} while (v28 > 0);
	a2[v28] = v44;
	v30 = 15;
	v31 = a2 + 14;
	do {
		if (v47 >= *v31)
			break;
		v31[1] = *v31;
		--v30;
		--v31;
	} while (v30 > 0);
	a2[v30] = v47;
	v32 = 0;
	v33 = 0;
	v51 = 0;
	v45 = 0;
	v48 = a2;
	do {
		v34 = *v27;
		v35 = 1 << *v27;
		result = 274 - v32;
		v57 = v35;
		if (v35 < 274 - v32)
			result = v35;
		v37 = 0;
		v54 = result;
		if (result > 0) {
			v38 = v34 + 4;
			v39 = v33 << v34;
			v40 = &v60[4 * v32 + 4];
			do {
				v41 = *(__int16*)v40;
				v40 += 4;
				*(_WORD*)(*(_DWORD*)(*(_DWORD*)v60 + 8) + 4 * v41) = v38;
				v42 = v37++ | v39;
				*(_WORD*)(*(_DWORD*)(*(_DWORD*)v60 + 8) + 4 * v41 + 2) = v42;
				result = v54;
			} while (v37 < v54);
			v35 = v57;
			v33 = v45;
			v32 = v51;
		}
		v32 += v35;
		++v33;
		v27 = v48 + 1;
		v51 = v32;
		v45 = v33;
		++v48;
	} while (v33 < 16);
	return result;
}

//----- (0057E2C0) --------------------------------------------------------
int __thiscall sub_57E2C0(_DWORD** this) {
	_DWORD** v1;         // esi
	_DWORD* v2;          // edi
	unsigned int v3;     // ebp
	int v4;              // eax
	int v5;              // edx
	unsigned __int16 v6; // bx
	int v7;              // ebp
	_DWORD* v8;          // edx
	_BYTE* v9;           // ebx
	int v10;             // ecx
	_DWORD* v11;         // eax
	int v12;             // edi
	char* v13;           // ebx
	_DWORD* v14;         // edx
	int v15;             // eax
	char v16;            // cl
	int v17;             // eax
	int v18;             // ebp
	unsigned int v19;    // ecx
	int v20;             // ecx
	_BYTE* v21;          // eax
	unsigned int v22;    // ecx
	_DWORD* v23;         // eax
	int v24;             // ecx
	int result;          // eax
	int v26;             // [esp+10h] [ebp-44h]
	char v27[64];        // [esp+14h] [ebp-40h]

	v1 = this;
	*this[2] = 2;
	if ((int)--*this[2] <= 0)
		sub_57E2C0(this);
	++*((_WORD*)v1[1] + 272);
	v2 = v1[6];
	v3 = (unsigned int)v1[5];
	v4 = (*v1)[2];
	v5 = *(__int16*)(v4 + 1088);
	v6 = *(_WORD*)(v4 + 1090);
	v1[6] = (_DWORD*)((char*)v2 + v5);
	v7 = (v6 << (32 - (_BYTE)v2 - v5)) | v3;
	v1[5] = (_DWORD*)v7;
	if ((int)v2 + v5 >= 16) {
		*(_BYTE*)v1[3] = HIBYTE(v7);
		v8 = v1[5];
		v9 = (char*)v1[3] + 1;
		v1[3] = v9;
		*v9 = BYTE2(v8);
		v10 = (int)(v1[6] - 4);
		v11 = (_DWORD*)((_DWORD)v1[5] << 16);
		v1[3] = (_DWORD*)((char*)v1[3] + 1);
		v1[6] = (_DWORD*)v10;
		v1[5] = v11;
	}
	sub_57DFC0(*v1, (int*)v27);
	v12 = 0;
	v13 = v27;
	v26 = 16;
	do {
		v14 = v1[6];
		v15 = *(_DWORD*)v13 - v12;
		v12 = *(_DWORD*)v13;
		v16 = 32 - (_BYTE)v14 - ++v15;
		v17 = (int)v14 + v15;
		v18 = 1 << v16;
		v19 = (unsigned int)v1[5];
		v1[6] = (_DWORD*)v17;
		v20 = v18 | v19;
		v1[5] = (_DWORD*)v20;
		if (v17 >= 16) {
			*(_BYTE*)v1[3] = HIBYTE(v20);
			v21 = (char*)v1[3] + 1;
			v22 = (unsigned int)v1[5] >> 16;
			v1[3] = v21;
			*v21 = v22;
			v23 = v1[5];
			v24 = (int)(v1[6] - 4);
			v1[3] = (_DWORD*)((char*)v1[3] + 1);
			v1[6] = (_DWORD*)v24;
			v1[5] = (_DWORD*)((_DWORD)v23 << 16);
		}
		v13 += 4;
		result = --v26;
	} while (v26);
	return result;
}

//----- (0057E3F0) --------------------------------------------------------
unsigned int __thiscall sub_57E3F0(_DWORD** this, unsigned int a2, int a3) {
	unsigned __int8* v3;  // ebx
	unsigned int result;  // eax
	_DWORD** v5;          // esi
	bool v6;              // cf
	unsigned __int16 v7;  // di
	int v8;               // eax
	_DWORD* v9;           // edi
	int v10;              // edx
	unsigned __int16 v11; // bp
	int v12;              // edx
	_DWORD* v13;          // edx
	_BYTE* v14;           // ebp
	int v15;              // ecx
	_DWORD* v16;          // eax
	unsigned int v17;     // [esp+Ch] [ebp+4h]

	v3 = (unsigned __int8*)a2;
	result = a2 + a3;
	v5 = this;
	v6 = a2 < a2 + a3;
	v17 = a2 + a3;
	if (v6) {
		do {
			v7 = *v3;
			if ((int)--*v5[2] <= 0)
				sub_57E2C0(v5);
			v8 = v7;
			++*((_WORD*)v5[1] + v7);
			v9 = v5[6];
			v10 = (*v5)[2];
			v11 = *(_WORD*)(v10 + 4 * v8 + 2);
			v12 = *(__int16*)(v10 + 4 * v8);
			v5[5] = (_DWORD*)((v11 << (32 - (_BYTE)v9 - v12)) | (unsigned int)v5[5]);
			v5[6] = (_DWORD*)((char*)v9 + v12);
			if ((int)v9 + v12 >= 16) {
				*(_BYTE*)v5[3] = (unsigned int)v5[5] >> 24;
				v13 = v5[5];
				v14 = (char*)v5[3] + 1;
				v5[3] = v14;
				*v14 = BYTE2(v13);
				v15 = (int)(v5[6] - 4);
				v16 = (_DWORD*)((_DWORD)v5[5] << 16);
				v5[3] = (_DWORD*)((char*)v5[3] + 1);
				v5[6] = (_DWORD*)v15;
				v5[5] = v16;
			}
			result = v17;
			++v3;
		} while ((unsigned int)v3 < v17);
	}
	return result;
}

//----- (0057E4C0) --------------------------------------------------------
int __thiscall sub_57E4C0(_DWORD** this, unsigned int a2, int a3, unsigned int a4, unsigned int a5) {
	_DWORD** v5;          // esi
	unsigned __int16 v6;  // di
	int v7;               // eax
	_DWORD* v8;           // edi
	unsigned int v9;      // ebp
	int v10;              // ecx
	int v11;              // edx
	unsigned __int16 v12; // bx
	unsigned __int16 v13; // bp
	unsigned int v14;     // ebx
	int v15;              // edi
	int v16;              // ecx
	_DWORD* v17;          // edx
	int v18;              // eax
	int v19;              // ebp
	unsigned int v20;     // edi
	int v21;              // ecx
	int v22;              // eax
	unsigned __int8* v23; // ebx
	unsigned __int16 v24; // bp
	_DWORD* v25;          // edx
	int v26;              // eax
	unsigned int v27;     // edi
	int v28;              // ebx
	int v29;              // eax
	_DWORD* v30;          // ecx
	_BYTE* v31;           // ebx
	int v32;              // ecx
	_DWORD* v33;          // eax
	int v34;              // edx
	unsigned int v35;     // eax
	_DWORD* v36;          // edi
	unsigned int v37;     // ebp
	int v38;              // edx
	int v39;              // ebp
	_DWORD* v40;          // edx
	_BYTE* v41;           // ebp
	int v42;              // edx
	_DWORD* v43;          // ecx
	_DWORD* v44;          // edx
	int v45;              // edi
	int result;           // eax
	_DWORD* v47;          // edi
	char v48;             // cl
	int v49;              // edx
	unsigned int v50;     // ecx
	_DWORD* v51;          // edx
	_BYTE* v52;           // ebx
	int v53;              // ecx

	v5 = this;
	sub_57E3F0(this, a2, a3);
	if (a4 >= 8) {
		if (a4 >= 0x26) {
			v23 = getMemAt(0x587000, 12 * ((a4 - 38) >> 5) + 314640);
			v24 = *getMemU16Ptr(0x587000, 12 * ((a4 - 38) >> 5) + 314640) + 4;
			if ((int)--*v5[2] <= 0)
				sub_57E2C0(v5);
			++*((_WORD*)v5[1] + v24);
			sub_57F160((int)v5, *(__int16*)((*v5)[2] + 4 * v24), *(unsigned __int16*)((*v5)[2] + 4 * v24 + 2));
			v25 = v5[6];
			v26 = *((_DWORD*)v23 + 1) + 4;
			v27 = (unsigned int)v5[5];
			v28 = (((_BYTE)a4 - 38) & 0x1F | (16 * *((unsigned __int16*)v23 + 4))) << (32 - (_BYTE)v25 - v26);
			v29 = (int)v25 + v26;
			v5[6] = (_DWORD*)v29;
			v5[5] = (_DWORD*)(v28 | v27);
			if (v29 >= 16)
				goto LABEL_14;
		} else {
			v13 = *getMemU16Ptr(0x587000, 12 * ((a4 - 8) >> 1) + 314640);
			v14 = (a4 - 8) & 1 | *getMemU16Ptr(0x587000, 12 * ((a4 - 8) >> 1) + 314648);
			v15 = *getMemU32Ptr(0x587000, 12 * ((a4 - 8) >> 1) + 314644);
			if ((int)--*v5[2] <= 0)
				sub_57E2C0(v5);
			++*((_WORD*)v5[1] + v13);
			v16 = (*v5)[2];
			v17 = v5[6];
			v18 = v15 + *(__int16*)(v16 + 4 * v13);
			v19 = *(unsigned __int16*)(v16 + 4 * v13 + 2) << v15;
			v20 = (unsigned int)v5[5];
			v21 = 32 - (_DWORD)v17 - v18;
			v22 = (int)v17 + v18;
			v5[6] = (_DWORD*)v22;
			v5[5] = (_DWORD*)(((v14 | v19) << v21) | v20);
			if (v22 >= 16)
				goto LABEL_14;
		}
	} else {
		v6 = a4 + 256;
		if ((int)--*v5[2] <= 0)
			sub_57E2C0(v5);
		v7 = v6;
		++*((_WORD*)v5[1] + v6);
		v8 = v5[6];
		v9 = (unsigned int)v5[5];
		v10 = (*v5)[2];
		v11 = *(__int16*)(v10 + 4 * v7);
		v12 = *(_WORD*)(v10 + 4 * v7 + 2);
		v5[6] = (_DWORD*)((char*)v8 + v11);
		v5[5] = (_DWORD*)((v12 << (32 - (_BYTE)v8 - v11)) | v9);
		if ((int)v8 + v11 >= 16) {
		LABEL_14:
			*(_BYTE*)v5[3] = (unsigned int)v5[5] >> 24;
			v30 = v5[5];
			v31 = (char*)v5[3] + 1;
			v5[3] = v31;
			*v31 = BYTE2(v30);
			v32 = (int)(v5[6] - 4);
			v33 = (_DWORD*)((_DWORD)v5[5] << 16);
			v5[3] = (_DWORD*)((char*)v5[3] + 1);
			v5[6] = (_DWORD*)v32;
			v5[5] = v33;
			goto LABEL_15;
		}
	}
LABEL_15:
	v34 = *getMemU32Ptr(0x587000, 8 * (a5 >> 9) + 314824) + 9;
	v35 = a5 & 0x1FF | (*getMemU16Ptr(0x587000, 8 * (a5 >> 9) + 314828) << 9);
	if (v34 <= 16) {
		v47 = v5[6];
		v48 = 32 - (_BYTE)v47 - v34;
		v49 = (int)v47 + v34;
		result = v35 << v48;
		v50 = (unsigned int)v5[5];
		v5[6] = (_DWORD*)v49;
		v5[5] = (_DWORD*)(result | v50);
		if (v49 < 16)
			return result;
		goto LABEL_21;
	}
	v36 = v5[6];
	v37 = (unsigned int)v5[5];
	v38 = *getMemU32Ptr(0x587000, 8 * (a5 >> 9) + 314824) - 7;
	v5[6] = (_DWORD*)((char*)v36 + v38);
	v39 = (v35 >> 16 << (32 - (_BYTE)v36 - v38)) | v37;
	v5[5] = (_DWORD*)v39;
	if ((int)v36 + v38 >= 16) {
		*(_BYTE*)v5[3] = HIBYTE(v39);
		v40 = v5[5];
		v41 = (char*)v5[3] + 1;
		v5[3] = v41;
		*v41 = BYTE2(v40);
		v42 = (int)(v5[6] - 4);
		v43 = (_DWORD*)((_DWORD)v5[5] << 16);
		v5[3] = (_DWORD*)((char*)v5[3] + 1);
		v5[6] = (_DWORD*)v42;
		v5[5] = v43;
	}
	v44 = v5[6];
	v45 = ((unsigned __int16)v35 << (16 - (_BYTE)v44)) | (unsigned int)v5[5];
	result = (int)(v44 + 4);
	v5[5] = (_DWORD*)v45;
	v5[6] = v44 + 4;
	if ((int)(v44 + 4) >= 16) {
	LABEL_21:
		*(_BYTE*)v5[3] = (unsigned int)v5[5] >> 24;
		v51 = v5[5];
		v52 = (char*)v5[3] + 1;
		v5[3] = v52;
		*v52 = BYTE2(v51);
		v53 = (int)(v5[6] - 4);
		result = (_DWORD)v5[5] << 16;
		v5[3] = (_DWORD*)((char*)v5[3] + 1);
		v5[6] = (_DWORD*)v53;
		v5[5] = (_DWORD*)result;
	}
	return result;
}

//----- (0057E7D0) --------------------------------------------------------
int __thiscall sub_57E7D0(_DWORD** this) {
	_DWORD** v1; // esi
	_DWORD* v2;  // edi
	int v3;      // edx
	_DWORD* v4;  // ecx
	_BYTE* v5;   // ebx
	int v6;      // ecx
	_DWORD* v7;  // eax
	_DWORD* v8;  // edi
	int v9;      // ebx

	v1 = this;
	if ((int)--*this[2] <= 0)
		sub_57E2C0(this);
	++*((_WORD*)v1[1] + 273);
	v2 = v1[6];
	v3 = *(__int16*)((*v1)[2] + 1092);
	v1[5] = (_DWORD*)((*(unsigned __int16*)((*v1)[2] + 1094) << (32 - (_BYTE)v2 - v3)) | (unsigned int)v1[5]);
	v1[6] = (_DWORD*)((char*)v2 + v3);
	if ((int)v2 + v3 >= 16) {
		*(_BYTE*)v1[3] = (unsigned int)v1[5] >> 24;
		v4 = v1[5];
		v5 = (char*)v1[3] + 1;
		v1[3] = v5;
		*v5 = BYTE2(v4);
		v6 = (int)(v1[6] - 4);
		v7 = (_DWORD*)((_DWORD)v1[5] << 16);
		v1[3] = (_DWORD*)((char*)v1[3] + 1);
		v1[6] = (_DWORD*)v6;
		v1[5] = v7;
	}
	if ((int)v1[6] > 0) {
		do {
			*(_BYTE*)v1[3] = (unsigned int)v1[5] >> 24;
			v8 = v1[5];
			v9 = (int)(v1[6] - 2);
			v1[3] = (_DWORD*)((char*)v1[3] + 1);
			v1[6] = (_DWORD*)v9;
			v1[5] = (_DWORD*)((_DWORD)v8 << 8);
		} while (v9 > 0);
	}
	return (char*)v1[3] - (char*)v1[4];
}

//----- (0057E8A0) --------------------------------------------------------
_DWORD* __thiscall sub_57E8A0(_DWORD* this) {
	_DWORD* v1;          // ebx
	unsigned __int8* v2; // eax

	v1 = this;
	sub_57DD90(this);
	v2 = (unsigned __int8*)operator_new(0x224u);
	v1[33] = v2;
	memcpy(v2, getMemAt(0x587000, 315976), 0x224u);
	memcpy(v1 + 1, getMemAt(0x587000, 315848), 0x80u);
	return v1;
}
// 5667CB: using guessed type void * operator_new(unsigned int);

//----- (0057E910) --------------------------------------------------------
void __thiscall sub_57E910(LPVOID* this) {
	LPVOID* v1; // esi

	v1 = this;
	operator_delete(this[33]);
	sub_57DDC0(v1);
}

//----- (0057E970) --------------------------------------------------------
int __thiscall sub_57E970(void** this) {
	memcpy(this[33], getMemAt(0x587000, 315976), 0x224u);
	memcpy(this + 1, getMemAt(0x587000, 315848), 0x80u);
	return sub_57DDD0(this);
}

//----- (0057E9A0) --------------------------------------------------------
_DWORD* __thiscall sub_57E9A0(_DWORD* this) {
	_DWORD* v1; // esi

	v1 = this;
	*this = operator_new(0x10000u);
	v1[1] = 0;
	sub_57E8A0(v1 + 2);
	v1[37] = 0;
	v1[36] = 0;
	return v1;
}
// 5667CB: using guessed type void * operator_new(unsigned int);

//----- (0057EA00) --------------------------------------------------------
void __thiscall sub_57EA00(LPVOID* this) {
	LPVOID* v1; // esi
	LPVOID* v2; // ecx
	LPVOID* v4; // [esp+4h] [ebp-10h]

	v1 = this;
	v4 = this;
	v2 = 0;
	if (v4)
		v2 = v1 + 2;
	sub_57E910(v2);
	operator_delete(*v1);
}

//----- (0057EA60) --------------------------------------------------------
int __thiscall sub_57EA60(int this) {
	*(_DWORD*)(this + 148) = 0;
	*(_DWORD*)(this + 144) = 0;
	*(_DWORD*)(this + 4) = 0;
	return sub_57E970((void**)(this + 8));
}

//----- (0057EA80) --------------------------------------------------------
int __thiscall nox_xxx_nxzDecompressImpl_57EA80(_DWORD* this, _BYTE* a2, _DWORD* a3, unsigned int a4, _DWORD* a5) {
	unsigned __int8* v5; // ebp
	_DWORD* v6;          // ebx
	unsigned int v7;     // edx
	int v8;              // esi
	int v9;              // eax
	int v10;             // ecx
	unsigned int v11;    // eax
	int v12;             // esi
	int v13;             // eax
	int v14;             // edi
	int v15;             // edx
	int v16;             // ecx
	int v17;             // eax
	int v18;             // ecx
	int v19;             // eax
	char* v20;           // ecx
	__int16 v21;         // si
	int v22;             // edi
	_DWORD* v23;         // edx
	int v24;             // esi
	int v25;             // ecx
	int v26;             // eax
	int v27;             // esi
	int v28;             // eax
	int v29;             // edi
	int v30;             // esi
	int v31;             // esi
	int v32;             // ecx
	int v33;             // edx
	int v34;             // esi
	int v35;             // ecx
	unsigned int v36;    // eax
	int v37;             // esi
	int v38;             // edi
	int v39;             // eax
	int v40;             // ecx
	unsigned int v41;    // eax
	int v42;             // edi
	int v43;             // eax
	int v44;             // edx
	_BYTE* v45;          // edi
	int v46;             // ebp
	int v47;             // eax
	int v48;             // ecx
	int v49;             // esi
	int v50;             // esi
	int v51;             // edx
	const void* v52;     // esi
	unsigned int v53;    // ecx
	int v54;             // esi
	int v55;             // edi
	_BYTE* v56;          // ebp
	int v57;             // edi
	int v58;             // edx
	void* v59;           // edi
	unsigned int v60;    // ecx
	_BYTE* v61;          // esi
	int v63;             // [esp+10h] [ebp-464h]
	int i;               // [esp+10h] [ebp-464h]
	int v65;             // [esp+10h] [ebp-464h]
	int v66;             // [esp+14h] [ebp-460h]
	int v67;             // [esp+14h] [ebp-460h]
	_BYTE* v68;          // [esp+14h] [ebp-460h]
	unsigned int v69;    // [esp+18h] [ebp-45Ch]
	int v70;             // [esp+1Ch] [ebp-458h]
	int v71;             // [esp+1Ch] [ebp-458h]
	_BYTE* v72;          // [esp+20h] [ebp-454h]
	unsigned int v73;    // [esp+24h] [ebp-450h]
	_BYTE* v74;          // [esp+28h] [ebp-44Ch]
	char v75[1096];      // [esp+2Ch] [ebp-448h]

	v5 = (unsigned __int8*)a4;
	v6 = this;
	v7 = a4 + *a5;
	v74 = a2;
	v73 = a4;
	v69 = a4 + *a5;
	v72 = &a2[*a3];
	this[37] = 0;
	if (a4 >= v7)
		return 0;
	while (1) {
		v8 = v6[37];
		if (v8 < 4) {
			if ((unsigned int)v5 >= v7) {
				v9 = -1;
				v6[37] = 0;
				v63 = -1;
				goto LABEL_9;
			}
			v10 = (*v5++ << (24 - v8)) | v6[36];
			v6[36] = v10;
			a4 = (unsigned int)v5;
			v6[37] = v8 + 8;
		}
		v11 = v6[36];
		v6[36] = 16 * v11;
		v63 = v11 >> 28;
		v6[37] -= 4;
		v9 = v11 >> 28;
	LABEL_9:
		v12 = v6[2 * v9 + 3];
		if (!v12) {
			v13 = *(__int16*)(v6[35] + 2 * v6[2 * v9 + 4]);
			goto LABEL_18;
		}
		v14 = v6[37];
		if (v14 >= v12)
			goto LABEL_15;
		if ((unsigned int)v5 < v7) {
			v16 = (*v5++ << (24 - v14)) | v6[36];
			v6[36] = v16;
			a4 = (unsigned int)v5;
			v6[37] = v14 + 8;
		LABEL_15:
			v15 = v6[36] >> (32 - v12);
			v6[36] <<= v12;
			v6[37] -= v12;
			v9 = v63;
			goto LABEL_16;
		}
		v6[37] = 0;
		v15 = -1;
	LABEL_16:
		v17 = v15 + v6[2 * v9 + 4];
		if (v17 >= 274)
			return 0;
		v13 = *(__int16*)(v6[35] + 2 * v17);
	LABEL_18:
		++*(_WORD*)(v6[2] + 2 * v13);
		// _dprintf("decompress: %d", v13);
		if (v13 < 256) {
			if (a2 < v72) {
				*a2++ = v13;
				v18 = v6[1];
				v6[1] = v18 + 1;
				*(_BYTE*)((unsigned __int16)v18 + *v6) = v13;
				goto LABEL_73;
			}
			return 0;
		}
		if (v13 == 272) {
			sub_57DEA0(v6 + 2, v75);
			v19 = 0;
			v20 = v75;
			do {
				v21 = *(_WORD*)v20;
				v20 += 4;
				*(_WORD*)(v19 + v6[35]) = v21;
				v19 += 2;
			} while (v19 < 548);
			v22 = 0;
			v23 = v6 + 4;
			v70 = 0;
			v66 = 16;
			while (1) {
				for (i = 0;; ++i) {
					v24 = v6[37];
					if (v24 >= 1)
						goto LABEL_29;
					if ((unsigned int)v5 >= v69)
						break;
					v25 = (*v5++ << (24 - v24)) | v6[36];
					v6[36] = v25;
					v6[37] = v24 + 8;
				LABEL_29:
					v26 = v6[36] >> 31;
					v27 = v6[37] - 1;
					v6[36] *= 2;
					v6[37] = v27;
					if (v26)
						goto LABEL_32;
				}
				v6[37] = 0;
			LABEL_32:
				v22 += i;
				*(v23 - 1) = v22;
				*v23 = v70;
				v23 += 2;
				v70 += 1 << v22;
				if (!--v66) {
					a4 = (unsigned int)v5;
					goto LABEL_73;
				}
			}
		}
		if (v13 == 273)
			break;
		if (v13 < 264) {
			v28 = v13 - 256;
			goto LABEL_43;
		}
		v29 = *getMemU32Ptr(0x587000, 8 * v13 + 314416);
		v30 = v6[37];
		if (v30 >= v29)
			goto LABEL_41;
		if ((unsigned int)v5 < v69) {
			v32 = (*v5++ << (24 - v30)) | v6[36];
			v6[36] = v32;
			a4 = (unsigned int)v5;
			v6[37] = v30 + 8;
		LABEL_41:
			v31 = v6[36] >> (32 - v29);
			v33 = v6[36] << v29;
			v6[37] -= v29;
			v6[36] = v33;
			goto LABEL_42;
		}
		v6[37] = 0;
		v31 = -1;
	LABEL_42:
		v28 = v31 + *getMemU32Ptr(0x587000, 8 * v13 + 314420);
	LABEL_43:
		v34 = v6[37];
		v67 = v28;
		if (v34 < 3) {
			if ((unsigned int)v5 >= v69) {
				v6[37] = 0;
				v71 = -1;
				goto LABEL_48;
			}
			v35 = (*v5++ << (24 - v34)) | v6[36];
			v6[36] = v35;
			a4 = (unsigned int)v5;
			v6[37] = v34 + 8;
		}
		v36 = v6[36];
		v6[36] = 8 * v36;
		v71 = v36 >> 29;
		v6[37] -= 3;
	LABEL_48:
		v65 = 0;
		v37 = *getMemU32Ptr(0x587000, 8 * v71 + 316592) + 9;
		if (v37 <= 8)
			goto LABEL_55;
		v38 = v6[37];
		v37 = *getMemU32Ptr(0x587000, 8 * v71 + 316592) + 1;
		if (v38 >= 8)
			goto LABEL_53;
		if ((unsigned int)v5 < v69) {
			v40 = (*v5++ << (24 - v38)) | v6[36];
			v6[36] = v40;
			a4 = (unsigned int)v5;
			v6[37] = v38 + 8;
		LABEL_53:
			v41 = v6[36];
			v6[36] = v41 << 8;
			v6[37] -= 8;
			v39 = v41 >> 24;
			goto LABEL_54;
		}
		v6[37] = 0;
		v39 = -1;
	LABEL_54:
		v65 = v39 << v37;
	LABEL_55:
		v42 = v6[37];
		if (v42 >= v37)
			goto LABEL_59;
		if ((unsigned int)v5 < v69) {
			v6[36] |= *v5 << (24 - v42);
			a4 = (unsigned int)(v5 + 1);
			v6[37] = v42 + 8;
		LABEL_59:
			v43 = v6[36] >> (32 - v37);
			v44 = v6[36] << v37;
			v6[37] -= v37;
			v6[36] = v44;
			goto LABEL_60;
		}
		v6[37] = 0;
		v43 = -1;
	LABEL_60:
		v45 = a2;
		v46 = (*getMemU32Ptr(0x587000, 8 * v71 + 316596) << 9) + (v65 | v43);
		v47 = v67 + 4;
		v68 = &a2[v67 + 4];
		// _dprintf("length: %d, distance: %d", v47, v46);
		if (v68 > v72)
			return 0;
		v48 = v6[1] - v46;
		if (v47 >= v46) {
			v50 = (unsigned __int16)v48;
			if ((unsigned __int16)v48 + v46 <= 0x10000) {
				v53 = v46;
				v52 = (const void*)(*v6 + v50);
			} else {
				v51 = 0x10000 - (unsigned __int16)v48;
				memcpy(a2, (const void*)(*v6 + (unsigned __int16)v48), 0x10000 - (unsigned __int16)v48);
				v52 = (const void*)*v6;
				v53 = v46 - v51;
				v45 = &a2[v51];
			}
			memcpy(v45, v52, v53);
			v54 = 0;
			v55 = v47 - v46;
			if (v47 - v46 > 0) {
				v56 = &a2[v46];
				do {
					++v54;
					v56[v54 - 1] = a2[v54 - 1];
				} while (v54 < v55);
			}
		} else {
			v49 = (unsigned __int16)v48;
			if (v49 + v47 <= 0x10000) {
				memcpy(a2, (const void*)(*v6 + v49), v47);
			} else {
				memcpy(a2, (const void*)(*v6 + v49), 0x10000 - v49);
				memcpy(&a2[0x10000 - v49], (const void*)*v6, v47 - (0x10000 - v49));
			}
		}
		v57 = v6[1] & 0xFFFF;
		if (v57 + v47 <= 0x10000) {
			v61 = a2;
			v60 = v47;
			v59 = (void*)(*v6 + v57);
		} else {
			v58 = 0x10000 - v57;
			memcpy((void*)(*v6 + v57), a2, 0x10000 - v57);
			v59 = (void*)*v6;
			v60 = v47 - v58;
			v61 = &a2[v58];
		}
		v5 = (unsigned __int8*)a4;
		memcpy(v59, v61, v60);
		v6[1] += v47;
		a2 = v68;
	LABEL_73:
		if ((unsigned int)v5 >= v69)
			return 0;
		v7 = v69;
	}
	if (a3)
		*a3 += v74 - a2;
	if (a5)
		*a5 += v73 - (_DWORD)v5;
	return 1;
}

//----- (0057F160) --------------------------------------------------------
_DWORD* __thiscall sub_57F160(int this, int a2, int a3) {
	_DWORD* result; // eax
	int v4;         // edx
	_BYTE* v5;      // ecx
	int v6;         // edx
	int v7;         // ecx
	int v8;         // edx

	result = (_DWORD*)this;
	v4 = *(_DWORD*)(this + 24);
	*(_DWORD*)(this + 20) |= a3 << (32 - v4 - a2);
	*(_DWORD*)(this + 24) = v4 + a2;
	if (v4 + a2 >= 16) {
		**(_BYTE**)(this + 12) = *(_DWORD*)(this + 20) >> 24;
		v5 = (_BYTE*)(*(_DWORD*)(this + 12) + 1);
		v6 = result[5] >> 16;
		result[3] = v5;
		*v5 = v6;
		v7 = result[5];
		v8 = result[6] - 16;
		++result[3];
		result[6] = v8;
		result[5] = v7 << 16;
	}
	return result;
}

//----- (0057F1D0) --------------------------------------------------------
char  sub_57F1D0(float2* a1) {
	char v1;            // bl
	int v2;             // edi
	double v3;          // st7
	double v4;          // st6
	double v5;          // st7
	unsigned __int8 v7; // [esp+Ch] [ebp-4h]

	v1 = 0;
	v2 = nox_float2int(a1->field_0);
	v7 = nox_float2int(a1->field_4) % 23;
	v3 = (double)(unsigned __int8)(v2 % 23);
	if (v3 >= 11.5) {
		v4 = (double)v7;
		if (v4 >= 11.5)
			v1 = 4;
		if (v4 <= 11.5)
			v1 |= 1u;
	}
	if (v3 <= 11.5) {
		v5 = (double)v7;
		if (v5 >= 11.5)
			v1 |= 8u;
		if (v5 <= 11.5)
			v1 |= 2u;
	}
	return v1;
}

//----- (0057F2A0) --------------------------------------------------------
int  sub_57F2A0(float2* a1, int a2, int a3) {
	int v3;     // esi
	int v4;     // eax
	int result; // eax
	int v6;     // eax
	float v7;   // [esp+0h] [ebp-Ch]
	float v8;   // [esp+0h] [ebp-Ch]

	v7 = a1->field_0 - (double)(23 * a2);
	v3 = nox_float2int(v7);
	v8 = a1->field_4 - (double)(23 * a3);
	v4 = nox_float2int(v8);
	if (v3 <= v4) {
		LOBYTE(v4) = 22 - v3 <= v4;
		v6 = v4 - 1;
		LOBYTE(v6) = v6 & 0xFD;
		result = v6 + 3;
	} else {
		LOBYTE(v4) = 22 - v3 <= v4;
		result = v4 + 1;
	}
	return result;
}

#if 0
//----- (0057FA20) --------------------------------------------------------
int __usercall sub_57FA20@<eax > (int a1@<ebp > )
{
    return sub_40E3D0((int*)(a1 - 16));
}

//----- (0057FA28) --------------------------------------------------------
int __usercall sub_57FA28@<eax > (int a1@<ebp > )
{
    return sub_40E3D0((int*)(a1 - 20));
}

//----- (0057FA30) --------------------------------------------------------
int  SEH_40E260(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 11984), a1, a2, a3, a4);
}

//----- (0057FA40) --------------------------------------------------------
int __usercall sub_57FA40@<eax > (int a1@<ebp > )
{
    return sub_40E3D0((int*)(a1 - 16));
}

//----- (0057FA48) --------------------------------------------------------
int __usercall sub_57FA48@<eax > (int a1@<ebp > )
{
    return sub_40E3D0((int*)(a1 - 20));
}

//----- (0057FA50) --------------------------------------------------------
int  SEH_40E320(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 12032), a1, a2, a3, a4);
}

//----- (0057FB20) --------------------------------------------------------
void __usercall sub_57FB20@<eax > (int a1@<ebp > )
{
    operator_delete(*(LPVOID*)(a1 - 16));
}

//----- (0057FB2B) --------------------------------------------------------
void __usercall sub_57FB2B@<eax > (int a1@<ebp > )
{
    operator_delete(*(LPVOID*)(a1 + 4));
}

//----- (0057FB36) --------------------------------------------------------
void __usercall sub_57FB36@<eax > (int a1@<ebp > )
{
    operator_delete(*(LPVOID*)(a1 + 4));
}

//----- (0057FB41) --------------------------------------------------------
void __usercall sub_57FB41@<eax > (int a1@<ebp > )
{
    operator_delete(*(LPVOID*)(a1 + 4));
}

//----- (0057FB57) --------------------------------------------------------
void __usercall sub_57FB57@<eax > (int a1@<ebp > )
{
    operator_delete(*(LPVOID*)(a1 - 16));
}

//----- (0057FB62) --------------------------------------------------------
void __usercall sub_57FB62@<eax > (int a1@<ebp > )
{
    operator_delete(*(LPVOID*)(a1 - 16));
}

//----- (0057FB6D) --------------------------------------------------------
void __usercall sub_57FB6D@<eax > (int a1@<ebp > )
{
    operator_delete(*(LPVOID*)(a1 - 16));
}

//----- (0057FBB0) --------------------------------------------------------
int __usercall sub_57FBB0@<eax > (int a1@<ebp > )
{
    return sub_5562D0(*(_DWORD * *)(a1 - 16));
}

//----- (0057FBB8) --------------------------------------------------------
int __usercall sub_57FBB8@<eax > (int a1@<ebp > )
{
    return sub_559AB0(*(_DWORD*)(a1 - 16) + 40);
}

//----- (0057FBC3) --------------------------------------------------------
int  SEH_556500(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 12456), a1, a2, a3, a4);
}

//----- (0057FBD0) --------------------------------------------------------
int __usercall sub_57FBD0@<eax > (int a1@<ebp > )
{
    return sub_5562D0(*(_DWORD * *)(a1 - 16));
}

//----- (0057FBD8) --------------------------------------------------------
int __usercall sub_57FBD8@<eax > (int a1@<ebp > )
{
    return sub_559AB0(*(_DWORD*)(a1 - 16) + 40);
}

//----- (0057FBE3) --------------------------------------------------------
int  SEH_556570(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 12504), a1, a2, a3, a4);
}

//----- (0057FBF0) --------------------------------------------------------
int __usercall sub_57FBF0@<eax > (int a1@<ebp > )
{
    return sub_5562D0(*(_DWORD * *)(a1 - 16));
}

//----- (0057FBF8) --------------------------------------------------------
int __usercall sub_57FBF8@<eax > (int a1@<ebp > )
{
    return sub_559AB0(*(_DWORD*)(a1 - 16) + 40);
}

//----- (0057FC03) --------------------------------------------------------
int  SEH_5565E0(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 12552), a1, a2, a3, a4);
}

//----- (0057FC50) --------------------------------------------------------
int __usercall sub_57FC50@<eax > (int a1@<ebp > )
{
    return sub_559AE0(*(_DWORD*)(a1 - 16) + 4);
}
// 559AE0: using guessed type int __thiscall sub_559AE0(_DWORD);

//----- (0057FC5B) --------------------------------------------------------
int  SEH_557C70(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 12648), a1, a2, a3, a4);
}

//----- (0057FC90) --------------------------------------------------------
wstring* __usercall sub_57FC90@<eax > (int a1@<ebp > )
{
    return sub_57211C(*(HIMC * *)(a1 - 16));
}

//----- (0057FC99) --------------------------------------------------------
int  SEH_5704C0(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 12728), a1, a2, a3, a4);
}

//----- (0057FCB0) --------------------------------------------------------
wstring* __usercall sub_57FCB0@<eax > (int a1@<ebp > )
{
    return sub_57211C(*(HIMC * *)(a1 - 16));
}

//----- (0057FCB9) --------------------------------------------------------
int  SEH_570570(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 12768), a1, a2, a3, a4);
}

//----- (0057FCD0) --------------------------------------------------------
void sub_57FCD0()
{
    nox_xxx___initp_misc_winxfltr_5717C0();
}

//----- (0057FCE1) --------------------------------------------------------
int  SEH_570B90(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 12808), a1, a2, a3, a4);
}

//----- (0057FD00) --------------------------------------------------------
wstring* __usercall sub_57FD00@<eax > (int a1@<ebp > )
{
    return sub_570CE0((int*)(a1 - 36));
}

//----- (0057FD09) --------------------------------------------------------
int  SEH_571810(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 12936), a1, a2, a3, a4);
}

//----- (0057FD13) --------------------------------------------------------
int* __usercall sub_57FD13@<eax > (int a1@<ebp > )
{
    return sub_573DE0((int*)(a1 - 48));
}

//----- (0057FD1C) --------------------------------------------------------
int* __usercall sub_57FD1C@<eax > (int a1@<ebp > )
{
    return sub_573DE0((int*)(a1 - 28));
}

//----- (0057FD25) --------------------------------------------------------
int* __usercall sub_57FD25@<eax > (int a1@<ebp > )
{
    return sub_573DE0((int*)(a1 - 92));
}

//----- (0057FD2E) --------------------------------------------------------
int* __usercall sub_57FD2E@<eax > (int a1@<ebp > )
{
    return sub_573DE0((int*)(a1 - 116));
}

//----- (0057FD37) --------------------------------------------------------
int  SEH_571C0E(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 12976), a1, a2, a3, a4);
}

//----- (0057FD41) --------------------------------------------------------
wstring* __usercall sub_57FD41@<eax > (int a1@<ebp > )
{
    return sub_570CE0((int*)(*(_DWORD*)(a1 - 32) + 8));
}

//----- (0057FD4D) --------------------------------------------------------
wstring* __usercall sub_57FD4D@<eax > (int a1@<ebp > )
{
    return sub_570CE0((int*)(*(_DWORD*)(a1 - 32) + 24));
}

//----- (0057FD59) --------------------------------------------------------
wstring* __usercall sub_57FD59@<eax > (int a1@<ebp > )
{
    return sub_570CE0((int*)(*(_DWORD*)(a1 - 32) + 40));
}

//----- (0057FD65) --------------------------------------------------------
int  SEH_572045(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 13040), a1, a2, a3, a4);
}

//----- (0057FD6F) --------------------------------------------------------
wstring* __usercall sub_57FD6F@<eax > (int a1@<ebp > )
{
    return sub_570CE0((int*)(*(_DWORD*)(a1 - 16) + 8));
}

//----- (0057FD7B) --------------------------------------------------------
wstring* __usercall sub_57FD7B@<eax > (int a1@<ebp > )
{
    return sub_570CE0((int*)(*(_DWORD*)(a1 - 16) + 24));
}

//----- (0057FD87) --------------------------------------------------------
wstring* __usercall sub_57FD87@<eax > (int a1@<ebp > )
{
    return sub_570CE0((int*)(*(_DWORD*)(a1 - 16) + 40));
}

//----- (0057FD93) --------------------------------------------------------
int  SEH_57211C(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 13096), a1, a2, a3, a4);
}

//----- (0057FD9D) --------------------------------------------------------
BOOL __usercall sub_57FD9D@<eax > (int a1@<ebp > )
{
    return sub_574EC0(a1 - 20);
}

//----- (0057FDA6) --------------------------------------------------------
wstring* __usercall sub_57FDA6@<eax > (int a1@<ebp > )
{
    return sub_570CE0((int*)(a1 - 40));
}

//----- (0057FDAF) --------------------------------------------------------
int  SEH_57222A(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 13152), a1, a2, a3, a4);
}

//----- (0057FDB9) --------------------------------------------------------
BOOL __usercall sub_57FDB9@<eax > (int a1@<ebp > )
{
    return sub_574EC0(a1 - 20);
}

//----- (0057FDC2) --------------------------------------------------------
wstring* __usercall sub_57FDC2@<eax > (int a1@<ebp > )
{
    return sub_570CE0((int*)(a1 - 40));
}

//----- (0057FDCB) --------------------------------------------------------
int  SEH_572333(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 13200), a1, a2, a3, a4);
}

//----- (0057FDD5) --------------------------------------------------------
BOOL __usercall sub_57FDD5@<eax > (int a1@<ebp > )
{
    return sub_574EC0(a1 - 20);
}

//----- (0057FDDE) --------------------------------------------------------
wstring* __usercall sub_57FDDE@<eax > (int a1@<ebp > )
{
    return sub_570CE0((int*)(a1 - 304));
}

//----- (0057FDEA) --------------------------------------------------------
int  SEH_572442(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 13248), a1, a2, a3, a4);
}

//----- (0057FE07) --------------------------------------------------------
BOOL __usercall sub_57FE07@<eax > (int a1@<ebp > )
{
    return sub_574EC0(a1 - 20);
}

//----- (0057FE10) --------------------------------------------------------
int  SEH_572A06(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 13336), a1, a2, a3, a4);
}

//----- (0057FE1A) --------------------------------------------------------
int* __usercall sub_57FE1A@<eax > (int a1@<ebp > )
{
    return sub_574660((int*)(a1 - 36));
}

//----- (0057FE23) --------------------------------------------------------
int* __usercall sub_57FE23@<eax > (int a1@<ebp > )
{
    return sub_574660((int*)(a1 - 60));
}

//----- (0057FE2C) --------------------------------------------------------
int  SEH_572BD6(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 13376), a1, a2, a3, a4);
}

//----- (0057FE36) --------------------------------------------------------
BOOL __usercall sub_57FE36@<eax > (int a1@<ebp > )
{
    return sub_574EC0(a1 - 20);
}

//----- (0057FE3F) --------------------------------------------------------
int** __usercall sub_57FE3F@<eax > (int a1@<ebp > )
{
    return sub_5705D0((int**)(a1 - 60));
}

//----- (0057FE48) --------------------------------------------------------
int  SEH_573153(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 13424), a1, a2, a3, a4);
}

//----- (0057FE52) --------------------------------------------------------
BOOL __usercall sub_57FE52@<eax > (int a1@<ebp > )
{
    return sub_574EC0(a1 - 20);
}

//----- (0057FE5B) --------------------------------------------------------
int** __usercall sub_57FE5B@<eax > (int a1@<ebp > )
{
    return sub_5705D0((int**)(a1 - 60));
}

//----- (0057FE64) --------------------------------------------------------
int  SEH_57330C(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 13472), a1, a2, a3, a4);
}

//----- (0057FE6E) --------------------------------------------------------
LPVOID* __usercall sub_57FE6E@<eax > (int a1@<ebp > )
{
    return sub_574930((LPVOID*)(a1 - 32));
}

//----- (0057FE77) --------------------------------------------------------
wstring* __usercall sub_57FE77@<eax > (int a1@<ebp > )
{
    return sub_570CE0((int*)(a1 - 60));
}

//----- (0057FE80) --------------------------------------------------------
wstring* __usercall sub_57FE80@<eax > (int a1@<ebp > )
{
    return sub_570CE0((int*)(a1 - 92));
}

//----- (0057FE89) --------------------------------------------------------
int  SEH_573401(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 13520), a1, a2, a3, a4);
}

//----- (0057FE93) --------------------------------------------------------
BOOL __usercall sub_57FE93@<eax > (int a1@<ebp > )
{
    return sub_574EC0(a1 - 20);
}

//----- (0057FE9C) --------------------------------------------------------
int  SEH_57366C(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 13576), a1, a2, a3, a4);
}

//----- (0057FEA6) --------------------------------------------------------
BOOL __usercall sub_57FEA6@<eax > (int a1@<ebp > )
{
    return sub_574EC0(a1 - 20);
}

//----- (0057FEAF) --------------------------------------------------------
int  SEH_57381A(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 13616), a1, a2, a3, a4);
}

//----- (0057FEB9) --------------------------------------------------------
BOOL __usercall sub_57FEB9@<eax > (int a1@<ebp > )
{
    return sub_574EC0(a1 - 20);
}

//----- (0057FEC2) --------------------------------------------------------
int  SEH_57390B(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 13656), a1, a2, a3, a4);
}

//----- (0057FECC) --------------------------------------------------------
BOOL __usercall sub_57FECC@<eax > (int a1@<ebp > )
{
    return sub_574EC0(a1 - 20);
}

//----- (0057FED5) --------------------------------------------------------
int  SEH_5739D2(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 13696), a1, a2, a3, a4);
}

//----- (0057FEE0) --------------------------------------------------------
void __usercall sub_57FEE0(int a1@<ebp > )
{
    std___Lockit__destructor_Lockit((std___Lockit*)(a1 - 16));
}
// 57F713: using guessed type void __thiscall std___Lockit__destructor_Lockit(std___Lockit *__hidden this);

//----- (0057FEE9) --------------------------------------------------------
int  SEH_574410(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 13736), a1, a2, a3, a4);
}

//----- (0057FF00) --------------------------------------------------------
void __usercall sub_57FF00(int a1@<ebp > )
{
    std___Lockit__destructor_Lockit((std___Lockit*)(a1 - 24));
}
// 57F713: using guessed type void __thiscall std___Lockit__destructor_Lockit(std___Lockit *__hidden this);

//----- (0057FF09) --------------------------------------------------------
int  SEH_5755F0(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 13776), a1, a2, a3, a4);
}

//----- (0057FF20) --------------------------------------------------------
void __usercall sub_57FF20(int a1@<ebp > )
{
    std___Lockit__destructor_Lockit((std___Lockit*)(a1 - 16));
}
// 57F713: using guessed type void __thiscall std___Lockit__destructor_Lockit(std___Lockit *__hidden this);

//----- (0057FF29) --------------------------------------------------------
int  SEH_575E60(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 13816), a1, a2, a3, a4);
}

//----- (0057FF40) --------------------------------------------------------
void __usercall sub_57FF40(int a1@<ebp > )
{
    std___Lockit__destructor_Lockit((std___Lockit*)(a1 - 32));
}
// 57F713: using guessed type void __thiscall std___Lockit__destructor_Lockit(std___Lockit *__hidden this);

//----- (0057FF49) --------------------------------------------------------
int  SEH_576A20(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 13856), a1, a2, a3, a4);
}

//----- (0057FF80) --------------------------------------------------------
void __usercall sub_57FF80(int a1@<ebp > )
{
    std___Lockit__destructor_Lockit((std___Lockit*)(a1 - 16));
}
// 57F713: using guessed type void __thiscall std___Lockit__destructor_Lockit(std___Lockit *__hidden this);

//----- (0057FF89) --------------------------------------------------------
int  SEH_576DA0(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 13936), a1, a2, a3, a4);
}

//----- (0057FFC0) --------------------------------------------------------
void __usercall sub_57FFC0(int a1@<ebp > )
{
    std___Lockit__destructor_Lockit((std___Lockit*)(a1 - 16));
}
// 57F713: using guessed type void __thiscall std___Lockit__destructor_Lockit(std___Lockit *__hidden this);

//----- (0057FFC9) --------------------------------------------------------
int  SEH_577060(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 14016), a1, a2, a3, a4);
}

//----- (0057FFE0) --------------------------------------------------------
void __usercall sub_57FFE0(int a1@<ebp > )
{
    std___Lockit__destructor_Lockit((std___Lockit*)(a1 - 16));
}
// 57F713: using guessed type void __thiscall std___Lockit__destructor_Lockit(std___Lockit *__hidden this);

//----- (0057FFE9) --------------------------------------------------------
int  SEH_577100(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 14056), a1, a2, a3, a4);
}

//----- (00580020) --------------------------------------------------------
int  SEH_577760(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 14136), a1, a2, a3, a4);
}

//----- (00580050) --------------------------------------------------------
void __usercall sub_580050(int a1@<ebp > )
{
    std___Lockit__destructor_Lockit((std___Lockit*)(a1 - 24));
}
// 57F713: using guessed type void __thiscall std___Lockit__destructor_Lockit(std___Lockit *__hidden this);

//----- (00580059) --------------------------------------------------------
int  SEH_577F90(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 14264), a1, a2, a3, a4);
}

//----- (00580070) --------------------------------------------------------
void __usercall sub_580070(int a1@<ebp > )
{
    std___Lockit__destructor_Lockit((std___Lockit*)(a1 - 16));
}
// 57F713: using guessed type void __thiscall std___Lockit__destructor_Lockit(std___Lockit *__hidden this);

//----- (00580079) --------------------------------------------------------
int  SEH_578240(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 14304), a1, a2, a3, a4);
}

//----- (00580090) --------------------------------------------------------
void __usercall sub_580090(int a1@<ebp > )
{
    std___Lockit__destructor_Lockit((std___Lockit*)(a1 - 16));
}
// 57F713: using guessed type void __thiscall std___Lockit__destructor_Lockit(std___Lockit *__hidden this);

//----- (00580099) --------------------------------------------------------
int  SEH_578370(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 14344), a1, a2, a3, a4);
}

//----- (005800B0) --------------------------------------------------------
void sub_5800B0()
{
    nox_xxx___initp_misc_winxfltr_5717C0();
}

//----- (005800C1) --------------------------------------------------------
int  SEH_578810(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 14384), a1, a2, a3, a4);
}

//----- (005800D8) --------------------------------------------------------
int __usercall sub_5800D8@<eax > (int a1@<ebp > )
{
    return sub_57DF70((LPVOID*)(*(_DWORD*)(a1 - 16) + 8));
}

//----- (005800F5) --------------------------------------------------------
int __usercall sub_5800F5@<eax > (int a1@<ebp > )
{
    return sub_57DF70((LPVOID*)(*(_DWORD*)(a1 - 16) + 8));
}

//----- (0058010A) --------------------------------------------------------
int __usercall sub_58010A@<eax > (int a1@<ebp > )
{
    return nullsub_71(a1 - 40);
}
// 57DD80: using guessed type int __thiscall nullsub_71(_DWORD);

//----- (00580112) --------------------------------------------------------
int  SEH_57D1C0(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 14520), a1, a2, a3, a4);
}

//----- (00580120) --------------------------------------------------------
void __usercall sub_580120@<eax > (int a1@<ebp > )
{
    sub_57DDC0(*(LPVOID * *)(a1 - 16));
}

//----- (00580128) --------------------------------------------------------
int  SEH_57DF00(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 14560), a1, a2, a3, a4);
}

//----- (00580132) --------------------------------------------------------
void __usercall sub_580132@<eax > (int a1@<ebp > )
{
    sub_57DDC0(*(LPVOID * *)(a1 - 16));
}

//----- (0058013A) --------------------------------------------------------
int  SEH_57DF70(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 14600), a1, a2, a3, a4);
}

//----- (00580144) --------------------------------------------------------
void __usercall sub_580144@<eax > (int a1@<ebp > )
{
    sub_57DDC0(*(LPVOID * *)(a1 - 16));
}

//----- (0058014C) --------------------------------------------------------
int  SEH_57E8A0(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 14640), a1, a2, a3, a4);
}

//----- (00580156) --------------------------------------------------------
void __usercall sub_580156@<eax > (int a1@<ebp > )
{
    sub_57DDC0(*(LPVOID * *)(a1 - 16));
}

//----- (0058015E) --------------------------------------------------------
int  SEH_57E910(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 14680), a1, a2, a3, a4);
}

//----- (0058017A) --------------------------------------------------------
int __usercall sub_58017A@<eax > (int a1@<ebp > )
{
    return sub_57DD70(*(LPVOID * *)(a1 - 16));
}

//----- (00580182) --------------------------------------------------------
int  SEH_57EA00(int a1, int a2, int a3, int a4)
{
    return __CxxFrameHandler((int)getMemAt(0x581450, 14760), a1, a2, a3, a4);
}
#endif

void nullsub_2() {}

int sub_448640(void) { return sub_44A400(); }

void nullsub_4(_DWORD a1, _DWORD a2, _DWORD a3, _DWORD a4) {}

void nox_xxx_set_sage(_DWORD a1) {}

void nullsub_5(void) {}

void nullsub_10(_DWORD a1) {}

int __thiscall sub_42CC50(LPVOID* this) { return sub_42C770(this); }

void nox_xxx_j_resetNPCRenderData_49A2E0(void) { nox_alloc_npcs(); }

void nullsub_16(void) {}

void nullsub_12(void) {}

void nullsub_17(void) {}

void nullsub_18(void) {}

void nullsub_11(void) {}

void nullsub_13(void) {}

void nullsub_15(void) {}

void nullsub_14(void) {}

void nullsub_3(void) {}

void nullsub_6(void) {}

void nullsub_19(void) {}

int nox_xxx_j_inventoryNameSignInit_467460(void) { return nox_xxx_inventoryNameSignInit_4671E0(); }

void  nox_alloc_npcs_2() { nox_alloc_npcs(); }

int nullsub_8(int a1, int a2) {
	return 0;
}
void nullsub_27(_DWORD a1) {}
void nullsub_28(_DWORD a1) {}
void nullsub_30(_DWORD a1) {}
void nullsub_36(void) {}
void nullsub_29(void) {}
void nullsub_35(_DWORD a1, _DWORD a2) {}
void nullsub_34(_DWORD a1, _DWORD a2, _DWORD a3) {}
void nullsub_20(void) {}
void nullsub_24(_DWORD a1) {}
void nullsub_31(_DWORD a1) {}
void nullsub_22(void) {}
void nullsub_23(void) {}
void nullsub_9(_DWORD a1) {}
void nullsub_33(_DWORD a1, _DWORD a2, _DWORD a3) {}
int __thiscall sub_558800(int (**this)(void)) { return this[385](); }
int __thiscall sub_558810(int (**this)(void)) { return this[386](); }
BOOL nox_xxx_testCPUID2_444D90() { return 0; }
void nox_xxx_j_allocHitArray_511840(void) { nox_xxx_allocHitArray_5486D0(); }

// ALL OK, 8062 function(s) have been successfully decompiled
