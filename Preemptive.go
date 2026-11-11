package main

import (
    "fmt"
    "math"
    "runtime"
    "sync"
    "time"
)

const TargetIterations int64 = 400_000_000

func main() {
    runtime.GOMAXPROCS(1)

    taskCounts := []int{1, 2, 4, 8}

    for _, count := range taskCounts {
       runBenchmark(count, TargetIterations)
    }
}

func runBenchmark(taskCount int, iterations int64) {
    logMsg(fmt.Sprintf("================ START: %d tasks ================", taskCount))
    //benchStart := time.Now()

    var wg sync.WaitGroup
    wg.Add(taskCount)

    for i := 1; i <= taskCount; i++ {
       taskID := fmt.Sprintf("task-%d", i)
       go func(id string) {
          defer wg.Done()

          taskStart := time.Now()

          _ = simulateCpuCode(iterations)

          taskDuration := time.Since(taskStart)
          logMsg(fmt.Sprintf("[%s] Finished! Task duration: %v ms", id, taskDuration.Milliseconds()))
       }(taskID)
    }

    wg.Wait()
}

func simulateCpuCode(iterations int64) float64 {
    var dummy float64 = 0
    for i := int64(1); i <= iterations; i++ {
       dummy += math.Sqrt(float64(i))
    }
    return dummy
}

func logMsg(msg string) {
    fmt.Printf("%s : %s\n", time.Now().Format("15:04:05.000"), msg)
}