package main

import (
    "flag"
    "fmt"
    "math"
    "runtime"
    "sync"
    "time"
)

var SEMAPHORE_ENABLED bool

func main() {
	flag.BoolVar(&SEMAPHORE_ENABLED, "semaphore-enabled", false, "enable semaphore")
	flag.Parse()
	logMsg(fmt.Sprintf("SEMAPHORE_ENABLED=%v", SEMAPHORE_ENABLED))

    runtime.GOMAXPROCS(1)

    taskCounts := []int{1, 2, 4, 8}

    for _, count := range taskCounts {
       runBenchmark(count)
    }
}

func runBenchmark(taskCount int) {
    logMsg(fmt.Sprintf("================ START: %d tasks ================", taskCount))

    var wg sync.WaitGroup
    wg.Add(taskCount)

	var cpuSemaphore chan struct{}

	if SEMAPHORE_ENABLED {
		cpuSemaphore = make(chan struct{}, 1)
		cpuSemaphore <- struct{}{} // permit
	}

    for i := 1; i <= taskCount; i++ {
        taskID := fmt.Sprintf("task-%d", i)

        go func(id string) {
            defer wg.Done()

			if SEMAPHORE_ENABLED {
				queuedAt := time.Now()
				<-cpuSemaphore
				queueWait := time.Since(queuedAt)
				logMsg(fmt.Sprintf("[%s] Queue Wait Time: %d ms", id, queueWait.Milliseconds()))

				func() {
					defer func() {
						cpuSemaphore <- struct{}{}
					}()

					cpuTask(id)
				}()
			} else {
				cpuTask(id)
			}
		}(taskID)
	}

    wg.Wait()
	logMsg(fmt.Sprintf("================ END: %d tasks ================", taskCount))
}

func cpuTask(id string) {
    taskStart := time.Now()

    _ = simulateCpuCode()

	taskDuration := time.Since(taskStart)
	logMsg(fmt.Sprintf("[%s] Finished! Task duration: %d ms", id, taskDuration.Milliseconds()))
}

func simulateCpuCode() float64 {
    var dummy float64 = 0
    for i := int64(1); i <= 400_000_000; i++ {
       dummy += math.Sqrt(float64(i))
    }
    return dummy
}

func logMsg(msg string) {
    fmt.Printf("%s : %s\n", time.Now().Format("15:04:05.000"), msg)
}