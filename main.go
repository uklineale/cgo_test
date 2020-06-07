package main

// #cgo CFLAGS: -g -Wall
// #include <stdlib.h>
// #include "allocate.h"
import "C"
import (
	"log"
	"unsafe"
	"flag"
	"time"
)

const SIZE = 1024
const MILLISECONDS = 1000

func main() {
	iters := flag.Int("iterations", 1000000, "number of memory allocations iterations")
	
	flag.Parse()

	cMalloc(*iters)
	cGoMalloc(*iters)
	goAssign(*iters)
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start);
	log.Printf("%s took %10s", name, elapsed);
}

func cMalloc(iterations int) {
	iters := C.int(iterations)
	duration := float64(C.allocate(iters)) // in seconds
	log.Printf("C malloc took %10fms\n", duration * MILLISECONDS)
}

func cGoMalloc(iterations int) {
	defer timeTrack(time.Now(), "cGoMalloc")
	for i := 0; i < iterations; i++ {
		ptr := C.calloc(1 , C.sizeof_char * SIZE)
		C.free(unsafe.Pointer(ptr))
	}
}

func goAssign(iterations int) {
	defer timeTrack(time.Now(), "goAssign")
	for i := 0; i < iterations; i++ {
		sl := make([]byte, SIZE)
		switch {
		case 1==1:
			// Do nothing
		case 1!=1:
			log.Printf("%s", sl)
		}
	}
}