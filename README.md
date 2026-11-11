# Cooperative vs Preemptive Scheduling
Ultra basic examples. Two requests at the same time.
Only one cpu-core.

# Cooperative
```sh
javac Cooperative.java && java Cooperative
# or
javac Cooperative.java && java Cooperative offload-cpu
```

#### Cooperative without offload-cpu flag

```txt
21:19:52.212347 : Thread[#3,main,5,main]         :  Offload CPU flag: false
21:19:52.215250 : Thread[#3,main,5,main]         :  Start benchmark...
21:19:52.226926 : VirtualThread[#26]/runnable@ForkJoinPool-1-worker-1    :  [task-1(IO)] QUEUE WAIT TIME: 3 ms
21:19:52.227897 : VirtualThread[#26]/runnable@ForkJoinPool-1-worker-1    :  [task-1] Start I/O (5s sleep)...
21:19:57.235726 : VirtualThread[#26]/runnable@ForkJoinPool-1-worker-1    :  [task-1] End I/O.
21:19:57.236713 : VirtualThread[#26]/runnable@ForkJoinPool-1-worker-1    :  [task-1] Start Cpu Code!
21:20:02.237116 : VirtualThread[#26]/runnable@ForkJoinPool-1-worker-1    :  [task-1] End CPU work!

// second task waited in a queue to start I/O part:(
// code works like a synchronous one
21:20:02.237604 : VirtualThread[#29]/runnable@ForkJoinPool-1-worker-1    :  [task-2(IO)] QUEUE WAIT TIME: 4012 ms
21:20:02.237656 : VirtualThread[#29]/runnable@ForkJoinPool-1-worker-1    :  [task-2] Start I/O (5s sleep)...
21:20:07.243069 : VirtualThread[#29]/runnable@ForkJoinPool-1-worker-1    :  [task-2] End I/O.
21:20:07.243725 : VirtualThread[#29]/runnable@ForkJoinPool-1-worker-1    :  [task-2] Start Cpu Code!
21:20:12.244268 : VirtualThread[#29]/runnable@ForkJoinPool-1-worker-1    :  [task-2] End CPU work!
21:20:12.244693 : Thread[#3,main,5,main]         :  =========================================
21:20:12.245442 : Thread[#3,main,5,main]         :  Total execution time: 20029 ms
21:20:12.245484 : Thread[#3,main,5,main]         :  =========================================
```

#### Cooperative with offload-cpu flag
```java
21:20:17.706891 : Thread[#3,main,5,main]         :  Offload CPU flag: true
21:20:17.709344 : Thread[#3,main,5,main]         :  Start benchmark...
21:20:17.722821 : VirtualThread[#26]/runnable@ForkJoinPool-1-worker-1    :  [task-1(IO)] QUEUE WAIT TIME: 2 ms
21:20:17.723643 : VirtualThread[#26]/runnable@ForkJoinPool-1-worker-1    :  [task-1] Start I/O (5s sleep)...
21:20:22.725841 : VirtualThread[#26]/runnable@ForkJoinPool-1-worker-1    :  [task-1] End I/O.
21:20:22.727813 : Thread[#29,pool-1-thread-1,5,main]     :  [task-1(CPU)] QUEUE WAIT TIME: 0 ms
21:20:22.727993 : Thread[#29,pool-1-thread-1,5,main]     :  [task-1] Start Cpu Code!

// second task didn't wait in a queue to start I/O part, yay!
21:20:23.715639 : VirtualThread[#30]/runnable@ForkJoinPool-1-worker-1    :  [task-2(IO)] QUEUE WAIT TIME: 0 ms
21:20:23.715733 : VirtualThread[#30]/runnable@ForkJoinPool-1-worker-1    :  [task-2] Start I/O (5s sleep)...
21:20:27.728276 : Thread[#29,pool-1-thread-1,5,main]     :  [task-1] End CPU work!
21:20:28.718733 : VirtualThread[#30]/runnable@ForkJoinPool-1-worker-1    :  [task-2] End I/O.
21:20:28.719940 : Thread[#29,pool-1-thread-1,5,main]     :  [task-2(CPU)] QUEUE WAIT TIME: 0 ms
21:20:28.720491 : Thread[#29,pool-1-thread-1,5,main]     :  [task-2] Start Cpu Code!
21:20:33.721416 : Thread[#29,pool-1-thread-1,5,main]     :  [task-2] End CPU work!
21:20:33.721948 : Thread[#3,main,5,main]         :  =========================================
21:20:33.722753 : Thread[#3,main,5,main]         :  Total execution time: 16012 ms
21:20:33.722797 : Thread[#3,main,5,main]         :  =========================================
```

# Preemptive
```sh
go run Preemptive.go
# or
go run Preemptive.go -semaphore-enabled
```

#### Preemptive without -semaphore-enabled flag
```txt
```go
// higher number of tasks at the same time -> slower task execution
21:38:16.030 : SEMAPHORE_ENABLED=false
21:38:16.030 : ================ START: 1 tasks ================
21:38:16.157 : [task-1] Finished! Task duration: 126 ms
21:38:16.157 : ================ END: 1 tasks ================
21:38:16.157 : ================ START: 2 tasks ================
21:38:16.400 : [task-2] Finished! Task duration: 242 ms
21:38:16.401 : [task-1] Finished! Task duration: 232 ms
21:38:16.401 : ================ END: 2 tasks ================
21:38:16.401 : ================ START: 4 tasks ================
21:38:16.868 : [task-3] Finished! Task duration: 408 ms
21:38:16.880 : [task-4] Finished! Task duration: 479 ms
21:38:16.884 : [task-1] Finished! Task duration: 471 ms
21:38:16.885 : [task-2] Finished! Task duration: 448 ms
21:38:16.885 : ================ END: 4 tasks ================
21:38:16.885 : ================ START: 8 tasks ================
21:38:17.708 : [task-1] Finished! Task duration: 803 ms
21:38:17.744 : [task-3] Finished! Task duration: 789 ms
21:38:17.777 : [task-5] Finished! Task duration: 777 ms
21:38:17.795 : [task-8] Finished! Task duration: 909 ms
21:38:17.813 : [task-7] Finished! Task duration: 764 ms
21:38:17.824 : [task-2] Finished! Task duration: 894 ms
21:38:17.833 : [task-4] Finished! Task duration: 855 ms
21:38:17.842 : [task-6] Finished! Task duration: 816 ms
21:38:17.842 : ================ END: 8 tasks ================
```

#### #### Preemptive with -semaphore-enabled flag
```go
// task duration is stable, yay!
// there is also task queue wait time, yay!
21:39:25.122 : SEMAPHORE_ENABLED=true
21:39:25.123 : ================ START: 1 tasks ================
21:39:25.123 : [task-1] Queue Wait Time: 0 ms
21:39:25.259 : [task-1] Finished! Task duration: 135 ms
21:39:25.259 : ================ END: 1 tasks ================
21:39:25.259 : ================ START: 2 tasks ================
21:39:25.259 : [task-2] Queue Wait Time: 0 ms
21:39:25.376 : [task-2] Finished! Task duration: 117 ms
21:39:25.376 : [task-1] Queue Wait Time: 114 ms
21:39:25.494 : [task-1] Finished! Task duration: 117 ms
21:39:25.494 : ================ END: 2 tasks ================
21:39:25.494 : ================ START: 4 tasks ================
21:39:25.494 : [task-4] Queue Wait Time: 0 ms
21:39:25.619 : [task-4] Finished! Task duration: 125 ms
21:39:25.619 : [task-1] Queue Wait Time: 107 ms
21:39:25.743 : [task-1] Finished! Task duration: 124 ms
21:39:25.743 : [task-2] Queue Wait Time: 231 ms
21:39:25.860 : [task-2] Finished! Task duration: 117 ms
21:39:25.860 : [task-3] Queue Wait Time: 349 ms
21:39:25.977 : [task-3] Finished! Task duration: 116 ms
21:39:25.977 : ================ END: 4 tasks ================
21:39:25.977 : ================ START: 8 tasks ================
21:39:25.977 : [task-8] Queue Wait Time: 0 ms
21:39:26.093 : [task-8] Finished! Task duration: 115 ms
21:39:26.093 : [task-1] Queue Wait Time: 92 ms
21:39:26.209 : [task-1] Finished! Task duration: 115 ms
21:39:26.209 : [task-2] Queue Wait Time: 208 ms
21:39:26.331 : [task-2] Finished! Task duration: 122 ms
21:39:26.331 : [task-3] Queue Wait Time: 331 ms
21:39:26.453 : [task-3] Finished! Task duration: 121 ms
21:39:26.453 : [task-4] Queue Wait Time: 452 ms
21:39:26.569 : [task-4] Finished! Task duration: 116 ms
21:39:26.569 : [task-5] Queue Wait Time: 569 ms
21:39:26.687 : [task-5] Finished! Task duration: 117 ms
21:39:26.687 : [task-6] Queue Wait Time: 687 ms
21:39:26.806 : [task-6] Finished! Task duration: 118 ms
21:39:26.806 : [task-7] Queue Wait Time: 806 ms
21:39:26.924 : [task-7] Finished! Task duration: 118 ms
21:39:26.924 : ================ END: 8 tasks ================
```