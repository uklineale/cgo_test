# CGo Test

A repo to experiment with C and Go interop. Heavily uses [CGo](https://blog.golang.org/cgo). 

## Memory Allocation Test

Currently the only supported test. Allocates memory using `calloc` or Go assigment for a specified number of iterations.

```
neel@neel-UX305FA:~/go/src/cgo-test$ ./cgo-test 
2020/06/07 09:15:23 C malloc took 124.793000ms
2020/06/07 09:15:24 cGoMalloc took 316.583381ms
2020/06/07 09:15:24 goAssign took 240.907226ms
```

The number of iterations can be specified using a command line flag.

```
neel@neel-UX305FA:~/go/src/cgo-test$ ./cgo-test -iterations 10000000
2020/06/07 09:14:56 C malloc took 1256.492000ms
2020/06/07 09:15:00 cGoMalloc took 3.491870767s
2020/06/07 09:15:02 goAssign took 2.372176123s
```

