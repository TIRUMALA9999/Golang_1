# Go (Golang) — Systems Programming & Concurrency Portfolio
## Core Language Fundamentals • File Handling • Concurrency • Synchronization • Worker Pools

![Go](https://img.shields.io/badge/Go-Golang%20Programming-blue)
![Concurrency](https://img.shields.io/badge/Concurrency-Goroutines%20%7C%20Channels-green)
![Systems](https://img.shields.io/badge/Systems-Backend%20Engineering-orange)


---

**Author:** Tirumala Teja Yegineni  


---

## 📌 Overview

This repository is a **comprehensive Golang learning and systems-programming portfolio** covering **core language fundamentals, file handling, concurrency patterns, synchronization primitives, and worker pool designs**.

The goal of this repository is to demonstrate:
- Strong command over **Go syntax and standard library**
- Practical **file I/O and data handling**
- Deep understanding of **concurrency using goroutines and channels**
- Correct usage of **synchronization primitives**
- Real-world **worker pool and fan-in/fan-out patterns**

This repository is highly relevant for **Backend Engineer, Platform Engineer, Distributed Systems, and Go Developer roles**.

---

## 📂 Repository Structure (High-Level)

```
Golang_1-main/
│
├── golang_notes.docx                  # Go concepts & language notes
│
├── File_Handling/
│   ├── json_handling.go
│   ├── csv_handling.go
│   ├── copying_files.go
│   ├── appending_to_existing.go
│   ├── p1.go, p2.go, p3.go
│   └── *.txt / *.csv                  # Sample files for I/O operations
│
├── struct/
│   ├── struct_basics.go
│   ├── structwithmethods.go
│
├── synchronization/
│   ├── Mutex_p1.go
│   ├── Map_p1.go
│   ├── Once_p1.go
│   ├── WaitGroupp1.go
│   └── concepts.go
│
├── workerpool_and_fanin_fanout/
│   ├── workerpool.go
│   ├── fanin_and_fanout.go
│   ├── context_p1.go
│   ├── context(cancellations_and_timeouts).go
│   └── p1.go ... p5.go
```

---

# 🧪 Go Concepts Covered 

---

## 1️⃣ Go Language Fundamentals

**File:** `golang_notes.docx`

### Topics Covered
- Go syntax & data types
- Functions & packages
- Error handling philosophy
- Go memory model
- Idiomatic Go practices

 
“What makes Go different from other languages?”

---

## 2️⃣ File Handling & Data I/O

**Directory:** `File_Handling/`

### Capabilities Demonstrated
- Reading & writing text files
- Appending to existing files
- Copying files
- Handling CSV data
- Parsing & writing JSON

### Why It Matters
File I/O is foundational for **backend services, data pipelines, and system utilities**.

---

## 3️⃣ Structs & Methods

**Directory:** `struct/`

### Concepts Covered
- Defining structs
- Attaching methods to structs
- Struct-based design patterns


“How does Go support object-oriented concepts?”

---

## 4️⃣ Concurrency with Goroutines & Channels

### Concepts Covered
- Lightweight goroutines
- Channel-based communication
- Concurrent execution patterns

 
“How does Go handle concurrency?”

---

## 5️⃣ Synchronization Primitives

**Directory:** `synchronization/`

### Tools Used
- `sync.Mutex`
- `sync.WaitGroup`
- `sync.Once`
- Concurrent-safe maps

### Concepts Demonstrated
- Race condition prevention
- Safe shared state access
- Coordinated goroutine execution

---

## 6️⃣ Worker Pools & Fan-In / Fan-Out Patterns

**Directory:** `workerpool_and_fanin_fanout/`

### Advanced Concepts
- Worker pool implementation
- Fan-in & fan-out concurrency patterns
- Context-based cancellation & timeouts
- Graceful goroutine shutdown


“How do you design scalable concurrent systems in Go?”

---

## 🧠 How This Fits Into My Portfolio

This repository complements my work in:
- Backend API development (FastAPI, Flask)
- Distributed systems concepts
- Data engineering pipelines
- Cloud-native & microservices architectures

It demonstrates my ability to **write performant, concurrent, and safe backend systems**.

---

## ⚙️ How to Run Examples

```bash
go version
go run File_Handling/json_handling.go
go run synchronization/Mutex_p1.go
go run workerpool_and_fanin_fanout/workerpool.go
```

---

## 🧾 Points 

- Developed **concurrent Go programs** using goroutines and channels, implementing worker pools and fan-in/fan-out patterns.  
- Implemented **file handling utilities in Go**, including CSV, JSON, and text file processing.  
- Used **synchronization primitives (Mutex, WaitGroup, Once)** to prevent race conditions and manage shared state safely.  
- Designed **context-aware concurrent workflows** supporting cancellation and timeouts.  
- Built strong foundations in **systems programming and backend engineering using Golang**.

---

## 🧠 I Points

- “Go uses CSP-style concurrency with goroutines and channels.”
- “Mutex and WaitGroup help coordinate shared state.”
- “Context enables cancellation and timeout control.”
- “Worker pools help scale concurrent workloads.”

---

## 👤 Author

**Tirumala Teja Yegineni**  
GitHub: https://github.com/TIRUMALA9999
