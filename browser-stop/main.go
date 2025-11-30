// browser-stop shuts down the CDP-controlled Chrome instance.
package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"time"
)

const debugPort = "9222"

func main() {
	debugURL := fmt.Sprintf("http://localhost:%s/json/version", debugPort)

	// Check if Chrome is running on debug port
	resp, err := http.Get(debugURL)
	if err != nil {
		fmt.Println("✗ No Chrome running on :" + debugPort)
		os.Exit(1)
	}
	resp.Body.Close()

	// Kill Chrome processes with remote debugging enabled
	if err := exec.Command("pkill", "-f", "remote-debugging-port="+debugPort).Run(); err != nil {
		fmt.Println("✗ Could not stop Chrome:", err)
		os.Exit(1)
	}

	// Verify Chrome stopped
	time.Sleep(500 * time.Millisecond)
	resp, err = http.Get(debugURL)
	if err == nil {
		resp.Body.Close()
		fmt.Println("✗ Chrome is still running")
		os.Exit(1)
	}

	fmt.Println("✓ Chrome stopped")
}
