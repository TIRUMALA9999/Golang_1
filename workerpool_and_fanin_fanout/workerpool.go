/* Problem Without Worker Pool
Imagine you have 100 jobs to do (like resizing images, sending emails, or processing payments).
If you start 100 goroutines (one per job), your CPU may get overloaded.
Some jobs might be very fast, some slow → resources are wasted.
You need a way to control concurrency (limit how many run at the same time). */

/*Worker Pool Idea

A Worker Pool is like a team of workers:
You create a fixed number of workers (say 3).
These workers are always waiting for jobs.
Jobs are sent into a queue (channel).
Workers pick jobs one by one, do them, and send back results.
So instead of 100 goroutines, you have 3 goroutines that process all jobs efficiently and in parallel.
*/

/* Worker Pool in Go
In Go:
Jobs channel → tasks are put here.
Workers → goroutines that read from jobs channel.
Results channel → workers put processed results here.
WaitGroup → used to wait until all workers finish. */

/*Benefits of Worker Pools

✅ Control concurrency (don’t overload CPU/memory).
✅ Efficient resource usage (fixed workers handle unlimited jobs).
✅ Works well for tasks like web scraping, file processing, DB queries. */

/* ⚡ In Simple Telugu + English Mix:

👉 Worker pool ante: okka queue lo jobs pedathaam, konni limited goroutines (workers) create chesthaam. 
Avi jobs ni line lo pick chesukoni parallel ga complete chesthai. Idhi CPU efficient ga untundi, 
overload kakunda. */