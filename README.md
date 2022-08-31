# Deprecation

Goal is to gently migrate a library using `init()` to a new version whitout `init()` and without breaking the code.

## Test

3 version of the code importing the library `./pkg/run`:

### No import

1. Linter succeeds:

   ```shell
   golangci-lint run ./cmd/v1/... --enable-all
   ```

2. Execution fails:
   ```shell
   go run cmd/v1/main.go
   ```
   ```
   panic: not yet initialized

   goroutine 1 [running]:
   github.com/holyhope/sandbox/run.Run()
         ./pkg/run/run.go:15 +0x78
   main.main()
         ./cmd/v1/main.go:8 +0x17
   exit status 2
   ```

### Old import

1. Linter fails:

   ```shell
   golangci-lint run ./cmd/v2/... --enable-all
   ```
   ```
   cmd/v2/main.go:4:4: SA1019: "github.com/holyhope/sandbox/auto-init" is deprecated: use github.com/holyhope/auto-init/v2.Init() instead. (staticcheck)
        _ "github.com/holyhope/sandbox/auto-init"
   ```

2. Execution succeeds:
   ```shell
   go run cmd/v2/main.go
   ```

### New import without Init()

1. Linter succeeds:

   ```shell
   golangci-lint run ./cmd/v3/... --enable-all
   ```

2. Execution fails:
   ```shell
   go run cmd/v3/main.go
   ```
   ```
   panic: not yet initialized

   goroutine 1 [running]:
   github.com/holyhope/sandbox/run.Run()
         ./pkg/run/run.go:15 +0x78
   main.main()
         ./cmd/v1/main.go:8 +0x17
   exit status 2
   ```

### New import with Init()

1. Linter succeeds:

   ```shell
   golangci-lint run ./cmd/v4/... --enable-all
   ```

2. Execution succeeds:
   ```shell
   go run cmd/v4/main.go
   ```
