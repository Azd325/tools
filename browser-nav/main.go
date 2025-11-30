// browser-nav navigates the browser to a URL.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
)

const (
	debugURL = "http://localhost:9222"
	timeout  = 30 * time.Second
)

func main() {
	newTab := flag.Bool("new", false, "Open in new tab")
	flag.Usage = func() {
		fmt.Println("Usage: browser-nav <url> [--new]")
		fmt.Println("\nExamples:")
		fmt.Println("  browser-nav https://example.com       # Navigate current tab")
		fmt.Println("  browser-nav https://example.com --new # Open in new tab")
	}
	flag.Parse()

	url := flag.Arg(0)
	if url == "" {
		flag.Usage()
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	allocCtx, allocCancel := chromedp.NewRemoteAllocator(ctx, debugURL)
	defer allocCancel()

	if *newTab {
		navigateNewTab(allocCtx, url)
	} else {
		navigateCurrentTab(allocCtx, url)
	}
}

func navigateNewTab(allocCtx context.Context, url string) {
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	if err := chromedp.Run(ctx, chromedp.Navigate(url)); err != nil {
		fmt.Println("✗ Could not open new tab:", err)
		fmt.Println("  Run: browser-start")
		os.Exit(1)
	}
	fmt.Println("✓ Opened:", url)
}

func navigateCurrentTab(allocCtx context.Context, url string) {
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

	if err := chromedp.Run(tabCtx, chromedp.Navigate(url)); err != nil {
		fmt.Println("✗ Could not navigate:", err)
		os.Exit(1)
	}
	fmt.Println("✓ Navigated to:", url)
}

func findLastPageTarget(targets []*target.Info) target.ID {
	for i := len(targets) - 1; i >= 0; i-- {
		if targets[i].Type == "page" {
			return targets[i].TargetID
		}
	}
	return ""
}
