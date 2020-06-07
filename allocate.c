#include "allocate.h"
#include <stdio.h>
#include <stdlib.h>
#include <time.h>

#define SIZE 1024
#define RESP_SIZE 20
#define NANOSECONDS 1000000000

void junk_func(char junk) {
    junk += 1;
}

double allocate(int iterations) {
    clock_t t;
    t = clock();
    for (int i = 0; i < iterations; i++) {
        char *buf;
        buf = (char *)calloc(1, SIZE);
        junk_func(buf[1]);
        free(buf);
    }

    t = clock() - t;
    double elapsed = (((double) t)/CLOCKS_PER_SEC);
    return elapsed;
}

