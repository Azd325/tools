# AGENTS.md

## Build & Run
- Build with Nix: `nix build .#browser-start` (or any tool name)
- Run directly: `cd browser-start && go run main.go`
- Dev shell: `nix develop`
- No tests exist; if added: `go test ./...` or `cd <tool> && go test`

## Project Structure
Each tool is a standalone Go module in its own directory with `main.go`, `go.mod`, `go.sum`, `default.nix`.

## Code Style
- **Imports**: stdlib first, then external packages, grouped by category
- **Naming**: camelCase for vars, funcs, consts, and structs (e.g., `lastPageTargetID`, `cleanMarkdown`)
- **Formatting**: Standard gofmt
- **Doc comments**: Package-level comment at file top: `// browser-foo does X.`

## Error Handling
- Check errors immediately after each call
- Print with emoji prefix: `fmt.Println("✗ Could not...", err)`
- Include helpful hint: `fmt.Println("  Run: browser-start")`
- Exit with `os.Exit(1)` on errors
- Use `defer cancel()` for all context cancellations

## CLI Patterns
- Use `flag` package with custom `flag.Usage` showing examples
- Success: `fmt.Println("✓ Chrome started...")` | Error: `fmt.Println("✗ Could not...")`
