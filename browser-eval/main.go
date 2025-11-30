// browser-eval evaluates JavaScript in the browser.
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
)

const (
	debugURL = "http://localhost:9222"
	timeout  = 30 * time.Second
)

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: browser-eval 'code'")
		fmt.Println("\nExamples:")
		fmt.Println(`  browser-eval "document.title"`)
		fmt.Println(`  browser-eval "document.querySelectorAll('a').length"`)
	}
	flag.Parse()

	code := strings.Join(flag.Args(), " ")
	if code == "" {
		flag.Usage()
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	allocCtx, allocCancel := chromedp.NewRemoteAllocator(ctx, debugURL)
	defer allocCancel()

	browserCtx, browserCancel := chromedp.NewContext(allocCtx)
	defer browserCancel()

	targets, err := chromedp.Targets(browserCtx)
	if err != nil {
		fmt.Println("✗ Could not connect to browser:", err)
		fmt.Println("  Run: browser-start")
		os.Exit(1)
	}

	targetID := findLastPageTarget(targets)
	if targetID == "" {
		fmt.Println("✗ No active tab found")
		os.Exit(1)
	}

	tabCtx, _ := chromedp.NewContext(allocCtx, chromedp.WithTargetID(targetID))

	var result any
	if err := chromedp.Run(tabCtx, chromedp.Evaluate(code, &result)); err != nil {
		fmt.Println("✗ Evaluation error:", err)
		os.Exit(1)
	}

	printResult(result)
}

func findLastPageTarget(targets []*target.Info) target.ID {
	for i := len(targets) - 1; i >= 0; i-- {
		if targets[i].Type == "page" {
			return targets[i].TargetID
		}
	}
	return ""
}

func printResult(result any) {
	if result == nil {
		fmt.Println("null")
		return
	}

	switch v := result.(type) {
	case []any:
		for i, item := range v {
			if i > 0 {
				fmt.Println()
			}
			printValue(item)
		}
	case map[string]any:
		printMap(v)
	default:
		if b, err := json.Marshal(v); err == nil {
			fmt.Println(string(b))
		} else {
			fmt.Println(v)
		}
	}
}

func printValue(v any) {
	if obj, ok := v.(map[string]any); ok {
		printMap(obj)
	} else {
		fmt.Println(v)
	}
}

func printMap(m map[string]any) {
	for key, value := range m {
		fmt.Printf("%s: %v\n", key, value)
	}
}
