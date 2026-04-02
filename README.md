# Gopher
learning go


## Phase 1 — Setup & Basics

### 1. What is Go & why use it

- [x] What is Go — compiled, statically typed, garbage collected
- [x] Why Go — simplicity, performance, concurrency, fast compilation
- [x] Go vs Python, Go vs Java, Go vs Rust — tradeoffs
- [x] Where Go is used — cloud infra, CLIs, microservices, DevOps tools
- [x] Notable projects built in Go — Docker, Kubernetes, Terraform, Hugo

### 2. Installing Go & workspace setup

- [x] Download and install Go from golang.org
- [x] `GOPATH` vs `GOROOT` — what each means
- [x] Go modules vs old GOPATH workflow
- [x] Setting up VS Code with Go extension (or GoLand)
- [x] Verifying install: `go version`, `go env`

### 3. Go CLI basics

- [x] `go run file.go` — compile and run
- [x] `go build` — compile to binary
- [x] `go install` — build and install to bin
- [x] `go fmt` — auto-format code
- [x] `go vet` — static analysis
- [x] `go clean` — remove build artifacts
- [x] `go doc` — view documentation

### 4. First Go program

- [x] `package main` — entry point package
- [x] `import` — importing packages
- [x] `func main()` — program entry point
- [x] `fmt.Println` — printing to stdout
- [x] How Go files are structured
- [x] Semicolons — why you don't write them (compiler inserts)

---

## Phase 2 — Language Fundamentals

### 5. Variables & constants

- [x] `var` declaration — explicit type
- [x] Short declaration `:=` — type inference
- [x] Multiple variable declaration
- [x] Zero values — default value per type
- [x] `const` — compile-time constants
- [x] `iota` — auto-incrementing const values
- [x] Typed vs untyped constants

### 6. Basic data types

- [x] Integers — `int`, `int8`, `int16`, `int32`, `int64`
- [x] Unsigned integers — `uint`, `uint8`, `uint32`, `uint64`
- [x] Floats — `float32`, `float64`
- [x] Complex numbers — `complex64`, `complex128`
- [x] `bool` — true/false
- [x] `string` — UTF-8 encoded, immutable
- [x] `byte` (alias for uint8) and `rune` (alias for int32)

### 7. Operators

- [x] Arithmetic — `+`, `-`, `*`, `/`, `%`
- [x] Comparison — `==`, `!=`, `<`, `>`, `<=`, `>=`
- [x] Logical — `&&`, `||`, `!`
- [x] Bitwise — `&`, `|`, `^`, `<<`, `>>`
- [x] Assignment operators — `+=`, `-=`, `*=`, `/=`
- [x] No ternary operator in Go — use if/else

### 8. Control flow — if / else / switch

- [x] Basic `if` / `else if` / `else`
- [x] If with init statement: `if x := val; x > 0`
- [x] `switch` statement — no fallthrough by default
- [x] `switch` with no condition — acts like if-else chain
- [x] `fallthrough` keyword
- [ ] `case` with multiple values

### 9. Loops

- [ ] `for` — only loop keyword in Go
- [ ] C-style for: `for i := 0; i < 10; i++`
- [ ] While-style: `for condition {}`
- [ ] Infinite loop: `for {}`
- [ ] `range` — iterate over slices, maps, strings, channels
- [ ] `break` and `continue`
- [ ] Labeled break — breaking outer loops

### 10. Functions

- [ ] Function declaration syntax
- [x] Multiple return values
- [x] Named return values
- [x] Variadic functions — `func sum(nums ...int)`
- [ ] Functions as values — assigning to variables
- [ ] Anonymous functions
- [ ] Immediately invoked function expressions (IIFE)
- [ ] `defer` — deferred execution, LIFO order
- [ ] `init()` function — runs before main

---

## Phase 3 — Composite Types

### 11. Arrays

- [x] Fixed-size, same-type elements
- [x] Declaration: `var a [5]int`
- [x] Array literal: `a := [3]int{1, 2, 3}`
- [x] `[...]` — compiler counts elements
- [x] Arrays are value types — copied on assignment
- [x] Iterating with range

### 12. Slices

- [x] Slice vs array — dynamic size, reference type
- [x] Creating with `make([]int, len, cap)`
- [x] Slice literal: `s := []int{1, 2, 3}`
- [x] `append` — adding elements, growth behavior
- [x] `len` and `cap`
- [x] Slicing a slice: `s[1:3]`
- [ ] Copy: `copy(dst, src)`
- [ ] Nil slice vs empty slice
- [ ] 2D slices

### 13. Maps

- [x] Key-value store, hash map under the hood
- [x] Declaration: `make(map[string]int)`
- [x] Map literal: `m := map[string]int{"a": 1}`
- [x] Adding, updating, reading keys
- [x] Deleting keys: `delete(m, key)`
- [x] Checking key existence: `v, ok := m[key]`
- [x] Iterating with range
- [x] Maps are reference types
- [x] Nil map — reading ok, writing panics

### 14. Structs

- [x] Defining a struct: `type Person struct {}`
- [x] Struct literals — named and positional
- [x] Accessing fields with dot notation
- [x] Pointer to struct — `&Person{}`
- [x] Anonymous structs
- [x] Struct embedding — composition over inheritance
- [x] Struct tags — `json:"name"`, `db:"name"`
- [x] Comparing structs — comparable if all fields are comparable

### 15. Pointers

- [x] What is a pointer — memory address
- [x] `&` — address-of operator
- [x] `*` — dereference operator
- [x] `new(T)` — allocate zeroed value, return pointer
- [x] Pointer to struct — auto-dereference with `.`
- [x] No pointer arithmetic in Go
- [x] When to use pointers — mutation, large structs, optional values
- [x] Nil pointer — dereferencing causes panic

---

## Phase 4 — Methods & Interfaces

### 16. Methods

- [x] Method vs function — receiver argument
- [x] Value receiver: `func (p Person) Greet()`
- [x] Pointer receiver: `func (p *Person) SetName()`
- [x] When to use pointer receiver — mutation, large structs
- [x] Methods on non-struct types
- [x] Method expressions and method values

### 17. Interfaces

- [x] What is an interface — set of method signatures
- [x] Implicit implementation — no `implements` keyword
- [x] Defining an interface: `type Stringer interface { String() string }`
- [x] Interface satisfaction — struct implements if it has all methods
- [x] Empty interface: `interface{}` / `any`
- [x] Interface as function parameter — polymorphism
- [x] Interfaces are reference types

### 18. Interface internals & patterns

- [x] Interface value = (type, value) pair
- [x] Nil interface vs interface holding nil pointer
- [x] Type assertion: `v, ok := i.(T)`
- [x] Type switch: `switch v := i.(type)`
- [x] Composing interfaces: `type ReadWriter interface { Reader; Writer }`
- [x] Common standard library interfaces — `io.Reader`, `io.Writer`, `fmt.Stringer`, `error`

### 19. Embedding & composition

- [x] Struct embedding — promoting fields and methods
- [x] Interface embedding — combining interfaces
- [x] Accessing embedded fields directly
- [x] Overriding embedded methods
- [x] Composition pattern — prefer over inheritance

---

## Phase 5 — Error Handling

### 20. Errors in Go

- [ ] `error` interface — `Error() string`
- [ ] Returning errors as second return value
- [ ] `errors.New("message")` — simple error
- [ ] `fmt.Errorf("context: %w", err)` — wrapping errors
- [ ] Checking error: `if err != nil`
- [ ] Ignoring errors with `_` — when not to do this

### 21. Custom error types

- [ ] Defining error struct: `type NotFoundError struct {}`
- [ ] Implementing `error` interface on custom type
- [ ] Adding context to errors with fields
- [ ] Type asserting to custom error for extra info
- [ ] Sentinel errors — `var ErrNotFound = errors.New(...)`

### 22. Error wrapping & unwrapping

- [ ] `%w` verb in `fmt.Errorf` — wraps error
- [ ] `errors.Is(err, target)` — checks chain for target
- [ ] `errors.As(err, &target)` — extracts type from chain
- [ ] `errors.Unwrap(err)` — one level unwrap
- [ ] Building meaningful error chains

### 23. Panic & recover

- [ ] `panic(value)` — stops normal execution
- [ ] When to panic — truly unrecoverable situations
- [ ] `recover()` — catch a panic inside `defer`
- [ ] Pattern: `defer func() { if r := recover(); r != nil {} }()`
- [ ] Panic vs error — when to use each
- [ ] Panics from nil pointer, out-of-bounds, type assertion failure

---

## Phase 6 — Packages & Modules

### 24. Packages

- [x] Every `.go` file belongs to a package
- [x] Package naming convention — lowercase, short
- [x] `package main` vs library packages
- [x] Exported vs unexported — uppercase = exported
- [x] Package-level variables and `init()`
- [x] Blank import: `import _ "pkg"` — side effects only
- [x] Dot import: `import . "pkg"` — avoid in production

### 25. Go modules

- [ ] `go mod init module-name` — create module
- [ ] `go.mod` file — module name, Go version, dependencies
- [ ] `go.sum` file — cryptographic checksums
- [ ] `go get package@version` — add/update dependency
- [ ] `go mod tidy` — clean up unused dependencies
- [ ] `go mod vendor` — vendor dependencies locally
- [ ] Module versioning — `v1`, `v2+` in import path

### 26. Organizing code

- [ ] Project layout conventions — `cmd/`, `internal/`, `pkg/`
- [ ] `internal/` package — restricted import
- [ ] `cmd/` — multiple binaries in one module
- [ ] Circular imports — Go forbids them
- [ ] Flat vs nested package structure
- [ ] Standard project layout (not official but widely used)

### 27. Working with standard library

- [ ] `fmt` — formatting, printing, scanning
- [ ] `os` — file system, env vars, process
- [ ] `io` — readers and writers
- [ ] `bufio` — buffered I/O
- [ ] `strings` — string manipulation
- [ ] `strconv` — string conversions
- [ ] `math` — math functions
- [ ] `sort` — sorting slices
- [ ] `time` — time and duration
- [ ] `log` — basic logging
- [ ] `path/filepath` — file paths
- [ ] `encoding/json` — JSON encode/decode

---

## Phase 7 — Concurrency

### 28. Goroutines

- [ ] What is a goroutine — lightweight thread managed by Go runtime
- [ ] `go func()` — launching a goroutine
- [ ] Goroutines vs OS threads — M:N scheduling
- [ ] Main goroutine exits → all goroutines die
- [ ] GOMAXPROCS — number of OS threads
- [ ] Goroutine stack — starts small, grows dynamically

### 29. Channels

- [ ] What is a channel — typed pipe for goroutine communication
- [ ] Creating: `ch := make(chan int)`
- [ ] Sending: `ch <- value`
- [ ] Receiving: `value := <-ch`
- [ ] Buffered channels: `make(chan int, 5)`
- [ ] Unbuffered vs buffered — blocking behavior
- [ ] Closing a channel: `close(ch)`
- [ ] Receiving from closed channel — zero value, ok=false
- [ ] Range over channel: `for v := range ch`

### 30. Select statement

- [ ] `select` — wait on multiple channel operations
- [ ] First ready case executes
- [ ] `default` case — non-blocking select
- [ ] Timeout pattern: `select { case <-time.After(1s) }`
- [ ] Done channel pattern — signal cancellation
- [ ] Select in a loop

### 31. Sync package

- [ ] `sync.WaitGroup` — wait for goroutines to finish
- [ ] `wg.Add(n)`, `wg.Done()`, `wg.Wait()`
- [ ] `sync.Mutex` — mutual exclusion lock
- [ ] `mu.Lock()`, `mu.Unlock()`
- [ ] `defer mu.Unlock()` pattern
- [ ] `sync.RWMutex` — multiple readers, one writer
- [ ] `sync.Once` — run function exactly once
- [ ] `sync.Map` — concurrent-safe map
- [ ] `sync.Pool` — reusable object pool

### 32. Common concurrency patterns

- [ ] Fan-out — one input, multiple workers
- [ ] Fan-in — multiple inputs merged to one channel
- [ ] Pipeline — chain of goroutine stages
- [ ] Worker pool — fixed goroutines processing job queue
- [ ] Done channel — propagate cancellation
- [ ] Semaphore with buffered channel
- [ ] Timeout with `time.After`

### 33. Race conditions & data races

- [ ] What is a data race — concurrent unsynchronized access
- [ ] `go run -race` — race detector
- [ ] `go test -race` — test with race detector
- [ ] Common race condition patterns
- [ ] Fixing races — channels or mutexes
- [ ] Atomic operations: `sync/atomic` package

### 34. Context package

- [ ] `context.Background()` and `context.TODO()`
- [ ] `context.WithCancel` — cancel propagation
- [ ] `context.WithTimeout` — deadline by duration
- [ ] `context.WithDeadline` — deadline by time
- [ ] `context.WithValue` — passing request-scoped data
- [ ] Passing context as first function argument (convention)
- [ ] Checking cancellation: `ctx.Done()`, `ctx.Err()`
- [ ] Context in HTTP requests and DB queries

---

## Phase 8 — Standard Library Deep Dive

### 35. `encoding/json`

- [ ] `json.Marshal` — Go struct → JSON bytes
- [ ] `json.Unmarshal` — JSON bytes → Go struct
- [ ] Struct tags: `json:"name,omitempty"`
- [ ] `json.NewEncoder` / `json.NewDecoder` — streaming
- [ ] Handling `null` — pointer fields
- [ ] Custom marshal/unmarshal — `MarshalJSON`, `UnmarshalJSON`
- [ ] `json.RawMessage` — defer parsing

### 36. `net/http` — HTTP client

- [ ] `http.Get`, `http.Post` — simple requests
- [ ] `http.NewRequest` + `http.Client.Do` — full control
- [ ] Setting headers, query params
- [ ] Reading response body — `io.ReadAll`, `defer body.Close()`
- [ ] HTTP client timeout
- [ ] Custom `http.Client` — transport, timeout, redirect policy
- [ ] Handling JSON responses

### 37. `net/http` — HTTP server

- [ ] `http.ListenAndServe`
- [ ] `http.HandleFunc` — register handler
- [ ] `http.Handler` interface — `ServeHTTP(w, r)`
- [ ] `http.ServeMux` — router/multiplexer
- [ ] Request — method, URL, headers, body
- [ ] Response writer — writing status, headers, body
- [ ] Middleware pattern — wrapping handlers
- [ ] `http.FileServer` — serving static files

### 38. `os` and `io` packages

- [ ] `os.Open`, `os.Create`, `os.OpenFile`
- [ ] Reading files — `os.ReadFile`, `bufio.Scanner`
- [ ] Writing files — `os.WriteFile`, `bufio.Writer`
- [ ] `os.Remove`, `os.Rename`, `os.MkdirAll`
- [ ] `os.Getenv`, `os.Setenv`, `os.LookupEnv`
- [ ] `os.Args` — CLI arguments
- [ ] `io.Reader` and `io.Writer` interfaces
- [ ] `io.Copy`, `io.ReadAll`, `io.TeeReader`

### 39. `time` package

- [ ] `time.Now()`, `time.Since()`, `time.Until()`
- [ ] `time.Duration` — nanoseconds under the hood
- [ ] Duration constants: `time.Second`, `time.Minute`
- [ ] `time.Sleep`
- [ ] `time.After`, `time.Tick`, `time.NewTimer`, `time.NewTicker`
- [ ] Parsing time: `time.Parse(layout, value)`
- [ ] Formatting time: `t.Format(layout)`
- [ ] Go's reference time — `Mon Jan 2 15:04:05 MST 2006`
- [ ] Time zones: `time.LoadLocation`

### 40. `strings` and `strconv`

- [ ] `strings.Contains`, `HasPrefix`, `HasSuffix`
- [ ] `strings.Split`, `strings.Join`
- [ ] `strings.Replace`, `strings.ReplaceAll`
- [ ] `strings.TrimSpace`, `strings.Trim`, `strings.TrimLeft/Right`
- [ ] `strings.ToUpper`, `strings.ToLower`
- [ ] `strings.Builder` — efficient string concatenation
- [ ] `strconv.Itoa`, `strconv.Atoi`
- [ ] `strconv.ParseInt`, `strconv.ParseFloat`, `strconv.ParseBool`
- [ ] `strconv.FormatInt`, `strconv.FormatFloat`

---

## Phase 9 — Testing

### 41. Unit testing basics

- [ ] Test files — `_test.go` suffix
- [ ] Test functions — `func TestXxx(t *testing.T)`
- [ ] `go test ./...` — run all tests
- [ ] `t.Error`, `t.Errorf` — report failure, continue
- [ ] `t.Fatal`, `t.Fatalf` — report failure, stop test
- [ ] `t.Log`, `t.Logf` — test output
- [ ] Running specific test: `go test -run TestName`
- [ ] Verbose mode: `go test -v`

### 42. Table-driven tests

- [ ] What are table-driven tests — idiomatic Go testing
- [ ] Defining test cases as slice of structs
- [ ] Looping over test cases with `t.Run`
- [ ] Subtests — named, runnable individually
- [ ] Parallel subtests: `t.Parallel()`
- [ ] Generating test names

### 43. Test helpers & utilities

- [ ] `t.Helper()` — marks function as test helper
- [ ] `testify/assert` and `testify/require` — popular assertion library
- [ ] `t.Cleanup(func)` — deferred cleanup
- [ ] `testing.Short()` — skip long tests with `-short` flag
- [ ] `t.Skip()` / `t.Skipf()` — skip tests conditionally
- [ ] Test fixtures — loading files from `testdata/` directory

### 44. Mocking & interfaces

- [ ] Why interfaces enable testing — inject fake dependencies
- [ ] Writing manual mocks — implement interface for tests
- [ ] `gomock` — mock generation from interfaces
- [ ] `mockery` — alternative mock generator
- [ ] Dependency injection pattern for testability
- [ ] HTTP mocking — `httptest.NewServer`, `httptest.NewRecorder`

### 45. Benchmarks & profiling

- [ ] Benchmark functions: `func BenchmarkXxx(b *testing.B)`
- [ ] `b.N` — number of iterations
- [ ] `go test -bench=.` — run benchmarks
- [ ] `go test -bench=. -benchmem` — memory allocation stats
- [ ] `b.ResetTimer()` — exclude setup time
- [ ] `b.RunParallel` — parallel benchmarks
- [ ] `go test -cpuprofile cpu.prof` — CPU profile
- [ ] `go tool pprof` — analyze profiles

### 46. Coverage

- [ ] `go test -cover` — show coverage percentage
- [ ] `go test -coverprofile=coverage.out`
- [ ] `go tool cover -html=coverage.out` — visual coverage report
- [ ] What coverage measures — statement coverage
- [ ] Writing tests to improve coverage
- [ ] Coverage in CI pipelines

---

## Phase 10 — Advanced Language Features

### 47. Generics (Go 1.18+)

- [ ] What are generics — type parameters
- [ ] Generic function: `func Map[T, U any](s []T, f func(T) U) []U`
- [ ] Type constraints — `comparable`, `any`
- [ ] Custom constraints with interfaces
- [ ] Generic structs
- [ ] `constraints` package — `Ordered`, `Integer`, `Float`
- [ ] When to use generics vs interfaces
- [ ] Type inference — compiler deduces type params

### 48. Closures

- [ ] What is a closure — function capturing outer variables
- [ ] Closure captures by reference
- [ ] Common gotcha — loop variable capture
- [ ] Fix loop capture: copy variable or use arg
- [ ] Closures as function factories
- [ ] Memoization with closures

### 49. Reflection

- [ ] `reflect` package — inspect types at runtime
- [ ] `reflect.TypeOf(v)`, `reflect.ValueOf(v)`
- [ ] Kinds — `reflect.Struct`, `reflect.Slice`, etc.
- [ ] Iterating struct fields with reflection
- [ ] Setting values via reflection
- [ ] When to use reflection — frameworks, serialization
- [ ] Reflection is slow — avoid in hot paths

### 50. `unsafe` package

- [ ] What is `unsafe` — bypass Go's type system
- [ ] `unsafe.Pointer` — generic pointer
- [ ] `unsafe.Sizeof`, `unsafe.Alignof`, `unsafe.Offsetof`
- [ ] Converting between pointer types
- [ ] `uintptr` vs `unsafe.Pointer`
- [ ] When unsafe is necessary — syscalls, performance-critical code
- [ ] Risks — GC, portability, undefined behavior

### 51. Build tags & conditional compilation

- [ ] Build constraints: `//go:build linux`
- [ ] OS and architecture constraints
- [ ] Custom tags: `go build -tags=production`
- [ ] `_test.go` files — automatic test build tag
- [ ] File-level vs function-level constraints
- [ ] `go generate` — code generation

---

## Phase 11 — Databases & Storage

### 52. `database/sql` package

- [ ] Opening connection: `sql.Open(driver, dsn)`
- [ ] `db.Ping()` — test connection
- [ ] `db.Query` — multiple rows
- [ ] `db.QueryRow` — single row
- [ ] `db.Exec` — insert/update/delete
- [ ] `rows.Scan` — read row into variables
- [ ] `defer rows.Close()`
- [ ] Connection pool settings — `SetMaxOpenConns`, `SetMaxIdleConns`

### 53. SQL drivers

- [ ] `lib/pq` — PostgreSQL driver
- [ ] `go-sql-driver/mysql` — MySQL driver
- [ ] `mattn/go-sqlite3` — SQLite driver
- [ ] Driver registration — `import _`
- [ ] DSN format per driver

### 54. Prepared statements & transactions

- [ ] `db.Prepare` — SQL injection prevention, performance
- [ ] `stmt.Query`, `stmt.Exec`
- [ ] `defer stmt.Close()`
- [ ] `db.Begin` — start transaction
- [ ] `tx.Commit()`, `tx.Rollback()`
- [ ] Defer rollback pattern
- [ ] Nested transactions — savepoints

### 55. ORMs & query builders

- [ ] GORM — full-featured ORM
- [ ] GORM models, hooks, associations
- [ ] `sqlx` — extends `database/sql` with struct scanning
- [ ] `sqlc` — generate Go code from SQL queries
- [ ] `squirrel` — query builder
- [ ] When to use ORM vs raw SQL vs query builder

### 56. Redis with Go

- [ ] `go-redis/redis` — popular Redis client
- [ ] Connecting and pinging
- [ ] String commands — `Set`, `Get`, `Del`
- [ ] Expiry — `Set` with TTL
- [ ] Lists, Sets, Hashes, Sorted Sets
- [ ] Pub/Sub
- [ ] Pipelining — batch commands
- [ ] Context-aware commands

---

## Phase 12 — Building Web Applications

### 57. HTTP routers

- [ ] `net/http` ServeMux limitations
- [ ] `gorilla/mux` — path params, method matching
- [ ] `chi` — lightweight, idiomatic, middleware support
- [ ] `gin` — fast, popular, full-featured
- [ ] `echo` — minimalist with good performance
- [ ] Path parameters, query parameters, wildcards

### 58. Middleware

- [ ] Middleware pattern — `func(http.Handler) http.Handler`
- [ ] Request logging middleware
- [ ] Authentication middleware
- [ ] CORS middleware
- [ ] Recovery middleware — catch panics
- [ ] Rate limiting middleware
- [ ] Chaining middleware

### 59. REST API patterns

- [ ] JSON request/response handling
- [ ] Input validation — `go-playground/validator`
- [ ] Error response structure
- [ ] HTTP status codes — when to use each
- [ ] Pagination patterns
- [ ] Versioning APIs — `/v1/`, header-based
- [ ] OpenAPI / Swagger with `swaggo/swag`

### 60. Authentication & Authorization

- [ ] JWT — `golang-jwt/jwt`
- [ ] Creating and signing tokens
- [ ] Validating tokens in middleware
- [ ] Refresh token pattern
- [ ] Password hashing — `bcrypt` (`golang.org/x/crypto/bcrypt`)
- [ ] Session-based auth
- [ ] OAuth2 — `golang.org/x/oauth2`

### 61. Configuration management

- [ ] `os.Getenv` — reading env vars
- [ ] `godotenv` — load `.env` files
- [ ] `viper` — full config management, multiple sources
- [ ] Config struct pattern — centralize all config
- [ ] Validating config on startup
- [ ] 12-factor app config principles

---

## Phase 13 — gRPC & Protobuf

### 62. Protocol Buffers

- [ ] What is protobuf — binary serialization format
- [ ] `.proto` file syntax — messages, fields, types
- [ ] Field numbers — backward compatibility
- [ ] Scalar types in proto3
- [ ] Nested messages, enums
- [ ] Installing `protoc` compiler
- [ ] `protoc-gen-go` — Go code generation

### 63. gRPC basics

- [ ] What is gRPC — RPC framework over HTTP/2
- [ ] Four service types — unary, server streaming, client streaming, bidirectional
- [ ] Defining services in `.proto`
- [ ] Generating Go code: `protoc --go_out --go-grpc_out`
- [ ] Implementing server interface
- [ ] Creating gRPC client
- [ ] Error handling with status codes

### 64. gRPC advanced

- [ ] Interceptors — server and client side (like middleware)
- [ ] Metadata — request headers equivalent
- [ ] Deadlines and cancellation with context
- [ ] gRPC reflection — server introspection
- [ ] `grpc-gateway` — REST to gRPC proxy
- [ ] Protobuf vs JSON — performance comparison
- [ ] `buf` tool — modern protobuf workflow

---

## Phase 14 — CLI Development

### 65. CLI basics with `os.Args`

- [ ] Reading `os.Args`
- [ ] `flag` package — built-in flag parsing
- [ ] Defining flags — `flag.String`, `flag.Int`, `flag.Bool`
- [ ] `flag.Parse()` — parse flags
- [ ] Subcommands with `flag.FlagSet`

### 66. Cobra — CLI framework

- [ ] What is Cobra — used by kubectl, gh, Hugo
- [ ] Root command, subcommands
- [ ] Flags — persistent vs local
- [ ] `cobra-cli` generator
- [ ] Args validation
- [ ] Shell completion generation
- [ ] `viper` integration for config

### 67. CLI UX patterns

- [ ] Progress bars — `schollz/progressbar`
- [ ] Spinners — `briandowns/spinner`
- [ ] Colored output — `fatih/color`
- [ ] Table output — `olekukonko/tablewriter`
- [ ] Interactive prompts — `AlecAivazis/survey`
- [ ] Reading from stdin / piped input
- [ ] Exit codes — 0 = success, non-zero = failure

---

## Phase 15 — DevOps & Production Patterns

### 68. Building & compiling

- [ ] Cross-compilation: `GOOS=linux GOARCH=amd64 go build`
- [ ] Build flags: `-ldflags="-X main.version=1.0.0"`
- [ ] Embedding files: `//go:embed` directive
- [ ] `embed.FS` — embedded filesystem
- [ ] CGO — when disabled and why: `CGO_ENABLED=0`
- [ ] Reducing binary size: `-ldflags="-s -w"`

### 69. Dockerizing Go apps

- [ ] Multi-stage Dockerfile for Go
- [ ] Stage 1: build with `golang:alpine`
- [ ] Stage 2: run with `scratch` or `distroless`
- [ ] Static binary for scratch image — `CGO_ENABLED=0`
- [ ] `.dockerignore` for Go projects
- [ ] Final image size comparison

### 70. Logging

- [ ] Standard `log` package limitations
- [ ] `slog` (Go 1.21+) — structured logging built-in
- [ ] `zerolog` — zero-allocation structured logger
- [ ] `zap` — Uber's fast structured logger
- [ ] Log levels — debug, info, warn, error
- [ ] Structured logging — JSON output for production
- [ ] Adding context to logs — request ID, trace ID

### 71. Metrics & observability

- [ ] `expvar` — built-in runtime metrics
- [ ] Prometheus client — `prometheus/client_golang`
- [ ] Defining counters, gauges, histograms
- [ ] Exposing `/metrics` endpoint
- [ ] OpenTelemetry — traces, metrics, logs
- [ ] Distributed tracing — trace IDs through services

### 72. Graceful shutdown

- [ ] `os/signal` — catching OS signals
- [ ] `signal.NotifyContext` — context cancelled on signal
- [ ] `http.Server.Shutdown(ctx)` — graceful HTTP shutdown
- [ ] Completing in-flight requests before exit
- [ ] Cleanup — close DB connections, flush logs
- [ ] Shutdown timeout pattern

### 73. Health checks & readiness

- [ ] `/healthz` — liveness endpoint
- [ ] `/readyz` — readiness endpoint
- [ ] Checking DB, cache, external deps in readiness
- [ ] Kubernetes liveness vs readiness probes
- [ ] Health check libraries — `alexliesenfeld/health`

---

## Phase 16 — Performance & Optimization

### 74. Memory management & GC

- [ ] Stack vs heap allocation
- [ ] Escape analysis: `go build -gcflags="-m"`
- [ ] Reducing heap allocations — performance impact
- [ ] GC pressure — what causes frequent GC
- [ ] `GOGC` environment variable — GC tuning
- [ ] `runtime.GC()` — trigger GC manually
- [ ] `runtime.MemStats` — memory statistics

### 75. Profiling

- [ ] `pprof` — Go's built-in profiler
- [ ] CPU profiling — `runtime/pprof` or `net/http/pprof`
- [ ] Memory profiling
- [ ] Goroutine profiling
- [ ] `go tool pprof` — interactive analysis
- [ ] Flame graphs
- [ ] `net/http/pprof` — expose profiling via HTTP

### 76. Benchmarking & optimization

- [ ] Writing good benchmarks
- [ ] `benchstat` — comparing benchmark results
- [ ] Reducing allocations — reuse buffers, `sync.Pool`
- [ ] String builder over concatenation
- [ ] Pre-allocating slices and maps
- [ ] Avoiding interface overhead in hot paths
- [ ] Inlining — when Go inlines functions

### 77. Common Go performance patterns

- [ ] Use `[]byte` instead of `string` for I/O
- [ ] Prefer `strings.Builder` over `+` concatenation
- [ ] Slice tricks — avoid unnecessary copies
- [ ] Map pre-sizing with `make(map[K]V, hint)`
- [ ] Worker pools to limit goroutine count
- [ ] Batching DB queries
- [ ] Connection pooling

---

## Phase 17 — Microservices & Cloud Patterns

### 78. Microservice patterns in Go

- [ ] Service structure — cmd, internal, pkg layout
- [ ] Config from environment
- [ ] Health and metrics endpoints
- [ ] Graceful shutdown
- [ ] Structured logging with correlation IDs
- [ ] Circuit breaker — `sony/gobreaker`
- [ ] Retry with backoff — `avast/retry-go`

### 79. Message queues

- [ ] Kafka with Go — `segmentio/kafka-go` or `confluentinc/confluent-kafka-go`
- [ ] Producing and consuming messages
- [ ] Consumer groups
- [ ] RabbitMQ — `rabbitmq/amqp091-go`
- [ ] NATS — `nats-io/nats.go`
- [ ] Message serialization — JSON vs protobuf

### 80. AWS SDK for Go

- [ ] `aws/aws-sdk-go-v2` — official SDK
- [ ] Authentication — credentials chain
- [ ] S3 — upload, download, presigned URLs
- [ ] DynamoDB — put, get, query, scan
- [ ] SQS — send, receive, delete messages
- [ ] Lambda handler in Go
- [ ] Using context for cancellation/timeout

### 81. Kubernetes client

- [ ] `client-go` — official K8s Go client
- [ ] In-cluster vs out-of-cluster config
- [ ] Listing, getting, creating resources
- [ ] Informers — watch for resource changes
- [ ] Custom controllers — reconcile loop pattern
- [ ] Operator SDK / Kubebuilder for operators

---

## Hands-on Projects Checklist

- [ ] **Project 1** — CLI calculator using `flag` package (Phase 2–3)
- [ ] **Project 2** — File organizer CLI tool (Phase 3–5)
- [ ] **Project 3** — REST API with `net/http` — CRUD for a resource (Phase 6–8)
- [ ] **Project 4** — REST API with Gin + PostgreSQL + GORM (Phase 11–12)
- [ ] **Project 5** — JWT auth middleware for REST API (Phase 12)
- [ ] **Project 6** — Concurrent web scraper with goroutines + channels (Phase 7)
- [ ] **Project 7** — Worker pool processing jobs from a queue (Phase 7)
- [ ] **Project 8** — CLI tool with Cobra + config with Viper (Phase 14)
- [ ] **Project 9** — gRPC service with unary and streaming RPCs (Phase 13)
- [ ] **Project 10** — Dockerized Go microservice with metrics + graceful shutdown (Phase 15–17)