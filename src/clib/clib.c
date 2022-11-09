#include <stdio.h>
#include "clib.h"

void cPrintMessage(char* message)
{
    printf("Message: %s\n", message);
}

void cGoBroker()
{
    printf ("** Completed running C code from GO! **\n\n");
}