package main

import (
	"fmt"
	"runtime/debug"
	"time"
)

/*
## Enable GC tracing
gctrace: You can enable detailed garbage collection trace logging using the GODEBUG environment variable.

```sh
go build -o <program_name>
ie. go build -o prog main.go 01_variables_and_datatypes.go 02_arithmetic.go 03_functions_and_control_structures.go 04_arrays_slices_maps_and_loops.go 99_gc_shananigans.go

GODEBUG=gctrace=1 ./<program_name> // this is the binary exe from step above

// output
Allocated 10MB
gc 1 @0.001s 2%: 0.026+0.15+0.002 ms clock, 0.32+0/0.25/0.003+0.031 ms cpu, 10->10->0 MB, 10 MB goal, 0 MB stacks, 0 MB globals, 12 P
Allocated 10MB
gc 2 @0.509s 0%: 0.010+0.13+0.002 ms clock, 0.12+0.056/0.095/0.002+0.032 ms cpu, 10->10->0 MB, 10 MB goal, 0 MB stacks, 0 MB globals, 12 P
Allocated 10MB
gc 3 @1.013s 0%: 0.014+0.15+0.004 ms clock, 0.16+0.067/0.11/0+0.051 ms cpu, 10->10->0 MB, 10 MB goal, 0 MB stacks, 0 MB globals, 12 P
Allocated 10MB
gc 4 @1.514s 0%: 0.046+0.41+0.011 ms clock, 0.55+0/0.61/0.10+0.13 ms cpu, 10->10->0 MB, 10 MB goal, 0 MB stacks, 0 MB globals, 12 P
```

## GC hinting:
Freeing Objects Early: Although Go does not provide explicit deallocation, you can influence GC by setting variables to nil or using short-lived scopes.
*/
func gcHintingMethod() {
	data := make([]byte, 1024*1024)
	fmt.Printf("data is at *p: %p\n", &data)
	data = nil // Hint to GC that data is no longer needed
}

func part99() {
	debug.SetGCPercent(20) // Set a more aggressive GC policy

	// Simulate workload
	for i := 0; i < 20; i++ {
		allocateMemory()
		time.Sleep(500 * time.Millisecond)
	}
}

func allocateMemory() {
	mem := make([]byte, 1*1024*1024) // Allocate 10MB
	fmt.Println("Allocated 1MB")
	_ = mem // Use mem to avoid compiler optimization
}
