#include <stdio.h>
#include "libcoins.h"

int main() {
    GoString key = {"Hello world!", 12};
    GoString msg = {"Hello world!", 12};
    GoString sig = BitcoinSign(key, msg);
    Print(sig);
}
