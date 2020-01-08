#include <stdio.h>
#include "libcoins.h"


int main() {
    char* key = "22a47fa09a223f2aa079edf85a7c2d4f8720ee63e502ee2869afab7de234b80c";
    char* msg = "Hello world!";
    char* sig = BitcoinSign(key, msg);
    printf("%s", sig);
}
