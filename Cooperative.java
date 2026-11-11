import java.time.Duration;
import java.time.LocalTime;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.concurrent.*;

public class Cooperative {

    private static boolean OFFLOAD_CPU;

    private static final ExecutorService CPU_EXECUTOR =
            Executors.newFixedThreadPool(1);

    // monitor queue wait time
    static class MonitoredCallable<T> implements Callable<T> {
        private final String id;
        private final Callable<T> delegate;
        private final long queuedAt = System.nanoTime();

        public MonitoredCallable(String id, Callable<T> delegate) {
            this.id = id;
            this.delegate = delegate;
        }

        @Override
        public T call() throws Exception {
            var millis = Duration.ofNanos(System.nanoTime() - queuedAt).toMillis();
            log(String.format("[%s] QUEUE WAIT TIME: %d ms", id, millis));
            return delegate.call();
        }
    }

    public static void main(String[] args) throws Exception {
        OFFLOAD_CPU = Arrays.asList(args).contains("offload-cpu");
        log("Offload CPU flag: " + OFFLOAD_CPU);

        System.setProperty("jdk.virtualThreadScheduler.parallelism", "1");
        System.setProperty("jdk.virtualThreadScheduler.maxPoolSize", "1");

        log("Start benchmark...");
        long totalStart = System.currentTimeMillis();

        var executor = Executors.newVirtualThreadPerTaskExecutor();
        List<Future<Integer>> futures = new ArrayList<>();

        // 1. submit task 1
        futures.add(executor.submit(new MonitoredCallable<>("task-1(IO)", () -> runIoCpuTask("task-1"))));

        // 2. sleep 6s
        Thread.sleep(6000);

        // 3. submit task 2
        futures.add(executor.submit(new MonitoredCallable<>("task-2(IO)", () -> runIoCpuTask("task-2"))));

        futures.forEach(Cooperative::throwingGetFuture);
        executor.close();
        CPU_EXECUTOR.close();

        long totalDuration = (System.currentTimeMillis() - totalStart);
        log("=========================================");
        log("Total execution time: " + totalDuration + " ms");
        log("=========================================");
    }

    private static Integer runIoCpuTask(String id) throws Exception {
        ioPart(id);
        cpuPart(id);
        return 1;
    }

    private static void ioPart(String id) throws InterruptedException {
        log("[" + id + "] Start I/O (5s sleep)...");
        Thread.sleep(Duration.ofSeconds(5));
        log("[" + id + "] End I/O.");
    }

    private static void cpuPart(String id) throws Exception {
        // off-load cpu-code somehow
        if (OFFLOAD_CPU) {
            CPU_EXECUTOR
                    .submit(new MonitoredCallable<>(id + "(CPU)", () -> {
                        simulateCpuCode(id);
                        return null;
                    }))
                    .get();

        } else {
            simulateCpuCode(id);
        }
    }

    static void simulateCpuCode(String id) throws InterruptedException {
        log("[" + id + "] Start Cpu Code!");
        long startTime = System.nanoTime();
        long totalDurationNanos = 5_000_000_000L; // 5s
        long iteration = 0L;

        while (System.nanoTime() - startTime < totalDurationNanos) {
            iteration++;
            Math.sqrt((double) iteration);
        }
        log("[" + id + "] End CPU work!");
    }

    static void throwingGetFuture(Future<Integer> future) {
        try {
            future.get();
        } catch (InterruptedException | ExecutionException e) {
            throw new RuntimeException(e);
        }
    }

    static void log(String msg) {
        System.out.println(LocalTime.now() + " : " + Thread.currentThread() + "\t :  " + msg);
    }
}

