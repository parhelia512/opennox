#ifndef NOX_PORT_COMW32_COMLIB
#define NOX_PORT_COMW32_COMLIB

#ifdef __GNUC__
#include <sys/socket.h>
#else
#include <winsock.h>
#endif

#include "../defs.h"

int __cdecl sub_552020(SOCKET s, char* buf, int len, int flags, struct sockaddr* from, int* fromlen); // idb

#endif //NOX_PORT_COMW32_COMLIB
