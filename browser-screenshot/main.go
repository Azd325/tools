// browser-screenshot captures a screenshot of the current page.
package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
)

const (
	debugURL = "http://localhost:9222"
	timeout  = 30 * time.Second
)

func main() {
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

	var buf []byte
	if err := chromedp.Run(tabCtx, chromedp.CaptureScreenshot(&buf)); err != nil {
		fmt.Println("✗ Screenshot failed:", err)
		os.Exit(1)
	}

	filename := fmt.Sprintf("screenshot-%s.png", time.Now().Format("2006-01-02T15-04-05"))
	outputPath := filepath.Join(os.TempDir(), filename)

	if err := os.WriteFile(outputPath, buf, 0644); err != nil {
		fmt.Println("✗ Could not save screenshot:", err)
		os.Exit(1)
	}

	fmt.Println("✓ Screenshot saved:", outputPath)
}

func findLastPageTarget(targets []*target.Info) target.ID {
	for i := len(targets) - 1; i >= 0; i-- {
		if targets[i].Type == "page" {
			return targets[i].TargetID
		}
	}
	return ""
}
