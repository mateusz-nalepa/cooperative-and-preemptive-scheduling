# Cooperative vs Preemptive Scheduling
Ultra basic examples. Two requests at the same time.
Only one cpu-core.

# Cooperative
```sh
javac Cooperative.java
java Cooperative
```

```java
14:22:37.385514 : Thread[#3,main,5,main]         :  Start benchmark...
14:22:37.397663 : VirtualThread[#26]/runnable@ForkJoinPool-1-worker-1    :  [task-1(IO)] QUEUE WAIT TIME: 3 ms
14:22:37.398205 : VirtualThread[#26]/runnable@ForkJoinPool-1-worker-1    :  [task-1] Start I/O (5s sleep)...
14:22:42.404838 : VirtualThread[#26]/runnable@ForkJoinPool-1-worker-1    :  [task-1] End I/O.
14:22:42.405306 : VirtualThread[#26]/runnable@ForkJoinPool-1-worker-1    :  [task-1] Start Cpu Code!
14:22:47.405772 : VirtualThread[#26]/runnable@ForkJoinPool-1-worker-1    :  [task-1] End CPU work!

second task waited in a queue :(       
14:22:47.406368 : VirtualThread[#29]/runnable@ForkJoinPool-1-worker-1    :  [task-2(IO)] QUEUE WAIT TIME: 4008 ms
14:22:47.406434 : VirtualThread[#29]/runnable@ForkJoinPool-1-worker-1    :  [task-2] Start I/O (5s sleep)...
14:22:52.411679 : VirtualThread[#29]/runnable@ForkJoinPool-1-worker-1    :  [task-2] End I/O.
14:22:52.411916 : VirtualThread[#29]/runnable@ForkJoinPool-1-worker-1    :  [task-2] Start Cpu Code!
14:22:57.412209 : VirtualThread[#29]/runnable@ForkJoinPool-1-worker-1    :  [task-2] End CPU work!
14:22:57.412609 : Thread[#3,main,5,main]         :  =========================================
// total execution time is 20s, not 16s        
14:22:57.413569 : Thread[#3,main,5,main]         :  Total execution time: 20024 ms
14:22:57.413615 : Thread[#3,main,5,main]         :  =========================================
```

# Preemptive
```sh
go run Preemptive.go
```

```go
// higher number of tasks at the same time -> slower task execution
14:23:35.203 : ================ START: 1 tasks ================
14:23:35.351 : [task-1] Finished! Task duration: 147 ms
14:23:35.351 : ================ START: 2 tasks ================
14:23:35.580 : [task-1] Finished! Task duration: 217 ms
14:23:35.589 : [task-2] Finished! Task duration: 238 ms
14:23:35.589 : ================ START: 4 tasks ================
14:23:36.003 : [task-1] Finished! Task duration: 412 ms
14:23:36.038 : [task-3] Finished! Task duration: 397 ms
14:23:36.056 : [task-4] Finished! Task duration: 467 ms
14:23:36.059 : [task-2] Finished! Task duration: 444 ms
14:23:36.059 : ================ START: 8 tasks ================
14:23:36.865 : [task-1] Finished! Task duration: 796 ms
14:23:36.903 : [task-3] Finished! Task duration: 783 ms
14:23:36.936 : [task-5] Finished! Task duration: 768 ms
14:23:36.964 : [task-8] Finished! Task duration: 904 ms
14:23:36.983 : [task-2] Finished! Task duration: 888 ms
14:23:36.987 : [task-4] Finished! Task duration: 843 ms
14:23:36.994 : [task-6] Finished! Task duration: 800 ms
14:23:36.999 : [task-7] Finished! Task duration: 780 ms
```