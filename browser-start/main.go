// browser-start launches Chrome with remote debugging enabled.
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"
)

const (
	chromePath          = "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome"
	debugPort           = "9222"
	cacheDir            = ".cache/browser-tools"
	startupTimeout      = 15 * time.Second
	startupPollInterval = 500 * time.Millisecond
)

func main() {
	// Check platform - this tool only supports macOS
	if runtime.GOOS != "darwin" {
		fmt.Println("✗ This tool only supports macOS")
		os.Exit(1)
	}

	profile := flag.Bool("profile", false, "Copy your default Chrome profile (cookies, logins)")
	kill := flag.Bool("kill", false, "Kill existing Chrome instances before starting")
	flag.Usage = func() {
		fmt.Println("Usage: browser-start [--profile] [--kill]")
		fmt.Println("\nOptions:")
		fmt.Println("  --profile  Copy your default Chrome profile (cookies, logins)")
		fmt.Println("  --kill     Kill existing Chrome instances before starting")
		fmt.Println("\nExamples:")
		fmt.Println("  browser-start            # Start with fresh profile")
		fmt.Println("  browser-start --profile  # Start with your Chrome profile")
		fmt.Println("  browser-start --kill     # Kill existing Chrome and start fresh")
	}
	flag.Parse()

	// Check if Chrome is already running on debug port
	debugURL := fmt.Sprintf("http://localhost:%s/json/version", debugPort)
	resp, err := http.Get(debugURL)
	if err == nil {
		resp.Body.Close()
		if !*kill {
			fmt.Println("✓ Chrome already running on :" + debugPort)
			return
		}
	}

	// Get home directory
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("✗ Could not determine home directory:", err)
		os.Exit(1)
	}

	userDataDir := filepath.Join(home, cacheDir)

	// Kill existing Chrome only if --kill flag is set
	if *kill {
		_ = exec.Command("killall", "Google Chrome").Run()
		time.Sleep(1 * time.Second)
	}

	// Setup profile directory
	if err := os.MkdirAll(userDataDir, 0755); err != nil {
		fmt.Println("✗ Could not create cache directory:", err)
		os.Exit(1)
	}

	// Copy profile if requested
	if *profile {
		chromeProfile := filepath.Join(home, "Library/Application Support/Google/Chrome")
		cmd := exec.Command("rsync", "-a", "--delete", chromeProfile+"/", userDataDir+"/")
		if err := cmd.Run(); err != nil {
			fmt.Println("✗ Could not sync Chrome profile:", err)
			os.Exit(1)
		}
	}

	// Start Chrome in background
	cmd := exec.Command(chromePath,
		"--remote-debugging-port="+debugPort,
		"--user-data-dir="+userDataDir,
	)
	if err := cmd.Start(); err != nil {
		fmt.Println("✗ Could not start Chrome:", err)
		os.Exit(1)
	}

	// Detach the process
	if err := cmd.Process.Release(); err != nil {
		fmt.Println("✗ Could not detach Chrome process:", err)
		os.Exit(1)
	}

	// Wait for Chrome to be ready
	deadline := time.Now().Add(startupTimeout)
	for time.Now().Before(deadline) {
		resp, err := http.Get(debugURL)
		if err == nil {
			resp.Body.Close()
			msg := fmt.Sprintf("✓ Chrome started on :%s", debugPort)
			if *profile {
				msg += " with your profile"
			}
			fmt.Println(msg)
			return
		}
		time.Sleep(startupPollInterval)
	}

	fmt.Println("✗ Chrome failed to start within timeout")
	os.Exit(1)
}
