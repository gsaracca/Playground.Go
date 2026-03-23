# Playground.Go

A collection of Go programs exploring algorithms, concurrency, and standard library features.

## Programs

### Pi via Leibniz Formula (`leibniz.go`)

Calculates π using the [Leibniz formula](https://en.wikipedia.org/wiki/Leibniz_formula_for_%CF%80):

```
π/4 = 1 - 1/3 + 1/5 - 1/7 + ...
```

Reads the number of iterations from `rounds.txt`, validates the input via a `CalculadorPi` health-check struct, runs the series, and reports both the result and elapsed time.

**Usage:**
```bash
echo 1000000 > rounds.txt
go run leibniz.go
```

---

### Sieve of Eratosthenes — Concurrent (`PrimesGo.go`)

Counts prime numbers up to a configurable limit using a concurrent segmented Sieve of Eratosthenes. Workers run in parallel via goroutines and `sync.WaitGroup`, with results validated against known prime counts.

**Flags:**

| Flag | Default | Description |
|---|---|---|
| `-limit` | `1000000` | Upper bound for the sieve |
| `-time` | `5s` | Benchmark duration |
| `-routines` | `3 × NumCPU` | Number of parallel workers |

**Usage:**
```bash
go run PrimesGo.go -limit 1000000 -time 5s -routines 8
```

---

### UUID v7 Generator (`UUID.go`)

Generates 1,000,000 UUID v7 values and writes them to `NewUUIDv7.txt`. Each UUID encodes a millisecond-precision Unix timestamp in the first 48 bits, followed by 10 cryptographically random bytes, with version (7) and variant bits set per the RFC 9562 spec.

**Usage:**
```bash
go run UUID.go
# Output written to NewUUIDv7.txt
```

---

## Requirements

- Go 1.21 or later

## Building

Each program lives in its own `package main` and can be compiled independently:

```bash
go build -o leibniz leibniz.go
go build -o primes PrimesGo.go
go build -o uuidgen UUID.go
```
