package	main

import (
    "log"
    "os"
    "runtime/pprof"
)

func main() {
    f, err := os.Create("cpu.prof")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    // Start CPU profiling
    pprof.StartCPUProfile(f)
    defer pprof.StopCPUProfile()

    // Run your code here
}
