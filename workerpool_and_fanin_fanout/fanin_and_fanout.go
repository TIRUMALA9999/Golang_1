/* 🔹 1. Fan-out
👉 Fan-out ante: One producer (goroutine) creates jobs and multiple workers (goroutines) 
take those jobs from the same channel and process them in parallel.
It’s like saying:
Producer: "Here are 100 tasks!"
Workers: "We’ll divide and work together." */


/* 🔹 2. Fan-in
👉 Fan-in ante: Multiple goroutines produce results and send them into one channel(single collector).
It’s like saying:
Workers: "We all finished tasks!"
Collector: "I’ll gather all your results into one place."  */

/* 🔹 3. Together (Fan-out + Fan-in)
👉 Common pattern:
Fan-out workers take jobs from one channel.
Fan-in sends their results into one results channel. */

/* 🔥 So, in short:
Fan-out = One producer, many workers.
Fan-in = Many workers, one collector. */


/* 🔹 1. Fan-out / Fan-in (as patterns)
Fan-out = multiple goroutines read jobs from one input channel (jobs are spread).
Fan-in = multiple goroutines write results into one output channel (results are merged).
These are basic building blocks.
No fixed "pool size" idea — you can spin up any number of goroutines, no tracking of when they finish (unless you add WaitGroup manually).
👉 Purely about direction of data flow (jobs spread → results merge). */


/* 🔹 2. Worker Pool
Worker Pool = a structured concurrency pattern that uses fan-out + fan-in together + synchronization (like sync.WaitGroup).
You decide a fixed number of workers → controls concurrency (e.g., 3 workers handle 100 jobs).
Ensures all jobs are processed and allows you to know when all work is done (using WaitGroup or closing channels properly).
Used in real production systems (web servers, job schedulers, etc.).
👉 It’s like a disciplined version of fan-out/fan-in with concurrency limits + lifecycle management. */


/* 🔹 Diagram Difference
Fan-out + Fan-in (basic idea)
Producer ---> Jobs ---> [Worker1]
                     ---> [Worker2]
                     ---> [Worker3]
Workers ---> Results ---> Collector  */

/* Worker Pool
Jobs Queue ---> [Fixed Workers] ---> Results Queue ---> Collector
             + WaitGroup / control to know when finished

✔️ Controlled concurrency
✔️ All jobs processed
✔️ Synchronization guaranteed */


/*🔹 Real-life Analogy
Fan-out/Fan-in → Like saying "hey, whoever is free, just grab some tasks and throw results back." (chaotic but works).
Worker Pool → Like a manager assigning tasks to a fixed team of workers and checking when everyone’s done before 
closing the office. (organized, controlled). */

/* ✅ In short:
Fan-out/Fan-in = low-level concepts (spread + merge).
Worker Pool = high-level pattern that uses both concepts + concurrency control. */

/* 🔹 Worker Pool = Fan-out + Fan-in
1. Fan-out Part
We had one jobs channel.
Multiple workers (worker1, worker2, worker3) were reading from the same jobs channel.
That means → jobs are fanned out (distributed) across workers.
for w := 1; w <= 3; w++ {
    go worker(w, jobs, results, &wg) // Fan-out
}  */


/* 2. Fan-in Part
Each worker, after finishing, sends result into the same results channel.
That means → all outputs are fanned in (collected) back into one channel.
results <- j * 2 // Fan-in */

/* 3. Collector
Main goroutine listens on results channel → gathers everything.
for r := range results {
    fmt.Println("Result:", r)
} */

/* 🔹 Diagram of Worker Pool
                FAN-OUT                          FAN-IN
Jobs ---> [ Worker 1 ] ----\
       ---> [ Worker 2 ] ----> Results ---> Main collects
       ---> [ Worker 3 ] ----/

Jobs channel distributes work → Fan-out
Results channel merges work back → Fan-in  */

/*🔹 So…

👉 Worker Pool is Fan-out + Fan-in combined with a WaitGroup to know when everything is finished.
👉 It’s the most practical concurrency pattern in Go for real projects.

⚡ Small Summary in Telugu + English:
Worker pool lo kuda fan-out (jobs distribute ayyi workers ki) and fan-in (workers anni results okkate channel ki pampadam) untayi.
So yes ra, worker pool = fan-out + fan-in! 🚀 */