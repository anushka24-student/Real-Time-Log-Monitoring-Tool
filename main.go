package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"time"
)

// LogEvent represents a structured log event
type LogEvent struct {
	Timestamp string
	Level     string
	Source    string
	Message   string
}

// Monitor manages log watching and synchronization
type Monitor struct {
	events chan LogEvent
	wg     sync.WaitGroup
}

// watchFile monitors a log file for new lines.
func (m *Monitor) watchFile(path string, done chan struct{}) {
	defer m.wg.Done()

	filename := filepath.Base(path)
	file, err := os.Open(path)
	if err != nil {
		log.Printf("Error opening file %s: %v", path, err)
		return
	}
	defer file.Close()

	log.Printf("Monitoring: %s", path)
	reader := bufio.NewReader(file)

	for {
		select {
		case <-done:
			return
		default:
			line, err := reader.ReadString('\n')
			if err != nil {
				time.Sleep(500 * time.Millisecond) // Wait for new logs to be written
				continue
			}

			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}

			// Extract log level (INFO, WARNING, ERROR, CRITICAL, FAILURE)
			level := "INFO"
			parts := strings.SplitN(line, ":", 2)
			if len(parts) > 1 {
				level = strings.TrimSpace(parts[0])
			}

			m.events <- LogEvent{
				Timestamp: time.Now().Format("15:04:05"),
				Level:     level,
				Source:    filename,
				Message:   line,
			}
		}
	}
}

// processEvents listens on the events channel and formats output
func processEvents(events chan LogEvent, done chan struct{}) {
	for {
		select {
		case event, ok := <-events:
			if !ok {
				return
			}
			fmt.Printf("\n[%s] %s | %s | %s\n", event.Timestamp, event.Level, event.Source, event.Message)
		case <-done:
			return
		}
	}
}

func main() {
	// Files to monitor (missing.log ko hata diya gaya hai)
	files := []string{
		"sample-logs/app.log",
		"sample-logs/server.log",
	}

	monitor := &Monitor{
		events: make(chan LogEvent, 100),
	}
	done := make(chan struct{})

	// Start event processor
	go processEvents(monitor.events, done)

	// Start watcher goroutines for each log file
	for _, path := range files {
		monitor.wg.Add(1)
		go monitor.watchFile(path, done)
	}

	fmt.Println("===================================")
	fmt.Println("      LOGPULSE MONITOR STARTED")
	fmt.Println("===================================")
	fmt.Println("Monitoring multiple log files...")
	fmt.Println("Press Ctrl+C to stop and generate reports.")

	// Wait for Ctrl+C signal
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	<-signalChannel

	fmt.Println("\nStopping LogPulse...")

	// Signal goroutines to stop and wait for worker completion
	close(done)
	monitor.wg.Wait()

	// Safely close the channel after all workers finish
	close(monitor.events)

	fmt.Println("LogPulse stopped.")
}
