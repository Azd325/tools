# Browser Tools

A collection of standalone CLI tools for controlling Google Chrome via the Chrome DevTools Protocol (CDP).

## Tools

| Tool | Description |
|------|-------------|
| `browser-start` | Launch Chrome with remote debugging enabled |
| `browser-stop` | Shut down the CDP-controlled Chrome instance |
| `browser-nav` | Navigate to a URL (current tab or new tab) |
| `browser-screenshot` | Capture a screenshot of the current page |
| `browser-eval` | Evaluate JavaScript in the page context |
| `browser-cookies` | List all cookies for the current page |

## Requirements

- macOS (Chrome path is hardcoded to `/Applications/Google Chrome.app`)
- Google Chrome installed
- Nix (for building) or Go 1.24+

## Installation

### With Nix

```bash
nix build .#browser-start
nix build .#browser-stop
nix build .#browser-nav
nix build .#browser-screenshot
nix build .#browser-eval
nix build .#browser-cookies
```

### Development

```bash
nix develop                          # Enter dev shell
cd browser-start && go run main.go   # Run directly
```

## Usage

### Typical Workflow

```bash
# 1. Start Chrome with debugging
browser-start

# 2. Navigate to a page
browser-nav https://example.com

# 3. Interact with the page
browser-screenshot                    # Capture screenshot
browser-eval "document.title"         # Get page title
browser-cookies                       # List cookies

# 4. Open another page in a new tab
browser-nav https://other-site.com --new

# 5. When done, stop Chrome
browser-stop
```

### browser-start

Launches Chrome with remote debugging on port 9222.

```bash
browser-start              # Start with fresh profile
browser-start --profile    # Start with your Chrome profile (cookies, extensions)
browser-start --kill       # Kill existing Chrome and start fresh
```

### browser-stop

Shuts down the CDP-controlled Chrome instance.

```bash
browser-stop
```

### browser-nav

Navigates to a URL.

```bash
browser-nav https://example.com        # Navigate current tab
browser-nav https://example.com --new  # Open in new tab
```

### browser-screenshot

Captures a screenshot of the current page and saves it to the temp directory.

```bash
browser-screenshot
# Output: /var/folders/.../screenshot-2025-01-15T10-30-45.png
```

### browser-eval

Evaluates JavaScript in the current page context.

```bash
browser-eval "document.title"
browser-eval "document.querySelectorAll('a').length"
browser-eval "Array.from(document.querySelectorAll('h1')).map(h => h.textContent)"
```

### browser-cookies

Lists all cookies for the current page.

```bash
browser-cookies
```

Output format:
```
session_id: abc123
  domain: .example.com
  path: /
  httpOnly: true
  secure: true
```

## License

MIT
