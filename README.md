# Go Core Patterns & Standard Library Cookbook

![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-blue.svg)
![Status](https://img.shields.io/badge/Status-Active_Learning-success)

A comprehensive reference repository documenting idiomatic usage of the Go Standard Library. 

This project serves as a **Pattern Library** (Cookbook) for reliable systems programming. It dissects core libraries (`net/http`, `sync`, `log`, etc.) through practical, real-world scenarios relevant to academic research and web interactions (e.g., verifying paper URLs, monitoring service uptime).

---

## 📖 Philosophy
Instead of abstract tutorials, this repository implements specific, isolated patterns to solidify understanding of Go's internal mechanics. Each module focuses on a specific library, explaining not just *how* to use a method, but *why* it behaves the way it does.

**Key Goals:**
* **Memory Management:** Understanding `defer` and closing connections.
* **Network Reliability:** Handling timeouts and status codes correctly.
* **Concurrency Safety:** Using Mutexes and Atomic operations to prevent race conditions.

---

## 📂 Repository Structure

The project is organized by standard library packages. Each folder contains a standalone `main.go` file demonstrating a specific pattern.

```text
go-core-patterns/
├── net-http/           # Web Client & Network Interactions
│   ├── basic_get.go   <-- Handling HTTP GET, Status Codes, and Body logic
│   └── client_timeouts.go
├── io-log/             # Logging & Input/Output
│   └── custom_log.go
├── sync-mutex/         # Concurrency Control
    └── safety.go
```

---

## 🛠️ Modules Overview

### 1. Networking (`net/http`) *(Coming Soon)*

**Scenario:** Automated verification of research paper availability on Google Scholar/GitHub.

**Key Concepts:**
- `http.Get()` vs `http.Client`: Understanding the default transport
- `resp.Body.Close()`: Preventing memory leaks by closing TCP streams
- `resp.StatusCode`: Programmatic validation of web resources (200 OK vs 404 Not Found)
- `io.ReadAll()`: Efficiently reading response streams

### 2. Logging & Diagnostics (`log`) *(Coming Soon)*

**Scenario:** Tracking the execution flow of long-running simulation scripts.

**Key Concepts:**
- Standard Output vs Standard Error
- Log prefixes and flags (Timestamping)
- Fatal errors vs recoverable warnings

### 3. Synchronization (`sync`) *(Coming Soon)*

**Scenario:** Managing shared counters in parallel data processing.

**Key Concepts:**
- `sync.Mutex`: Locking memory to prevent Race Conditions
- `defer mu.Unlock()`: Safe unlocking patterns

## 🚀 Usage

You can run any specific pattern directly from the root of the repository.

### Example: Running the HTTP Pattern

```bash
go run net-http/main.go
```

**Example Output:**

```
--- 1. The http.Get Method ---
Status Code: 200
Result: Success! Your profile is live.
--- 3. Reading the Body ---
Preview of Page Content:
<!DOCTYPE html>...
```

## 🧠 Knowledge Retention (Cheat Sheet)

### `http.Get(url)`
- **Use for:** Quick, simple checks where default settings (no timeout) are acceptable
- **Warning:** The default client has NO timeout. If the server hangs, your program hangs forever

### `defer resp.Body.Close()`
- **Use for:** Always cleaning up after a network call
- **Why:** Go keeps the underlying TCP connection open for re-use. If you don't close the body, the connection remains "busy" and you will eventually run out of file descriptors (crash)


## 🎯 Learning Goals

This repository is designed to help you:
- Understand Go's standard library through practical, research-oriented scenarios
- Avoid common pitfalls in network programming, concurrency, and resource management
- Build production-ready Go applications with confidence

## 👤 Author

**Armin Ahmadkhaniha**

Quantum Software Developer & PhD Student

Researching Quantum Algorithms, Machine Learning, and Distributed Systems

---

⭐ If you find this resource helpful, please consider starring the repository!