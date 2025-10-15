/* 1. What is context package?
The context package in Go is used to control the lifetime of goroutines.
It helps you cancel goroutines or set deadlines/timeouts so they don’t run forever.
Prevents resource leaks (like orphan goroutines still running even after you don’t need them). */


/* 2. Why do we need it?
Imagine:
You start a goroutine to fetch data from a DB or API.
The user cancels the request (closes browser, hits stop).
Without context, your goroutine will still run and waste CPU/Memory.
👉 With context, you can signal: "Stop now, I don’t need the result anymore." */

/*3. Types of Context
There are mainly 3 functions to create cancellable contexts:
context.WithCancel → lets you manually cancel.
context.WithTimeout → cancels automatically after given time.
context.WithDeadline → cancels at a specific time (absolute time).*/

/*6. How Context Works Internally
Context is like a signal channel (ctx.Done()).
Workers/goroutines listen to it in a select statement.
When cancellation or timeout happens → all listening goroutines stop.*/

/*7. Real-world Analogy
Think of context as a remote control.
You start multiple goroutines = TVs.
If you press the cancel button = all TVs get switched off at once.
If you set a timeout = TVs automatically switch off after some time.*/

/*✅ In short
WithCancel → manual stop button.
WithTimeout → auto-stop after X time.
WithDeadline → stop at a specific time.
Goroutines must respect ctx.Done() signal.*/